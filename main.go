package main

import (
	"bytes"
	"context"
	"io"
	"log"
	"strconv"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"gitlab.com/tozd/go/errors"
)

type PostgresqlLO struct {
	dbpool *pgxpool.Pool
}

func (e *PostgresqlLO) Close() errors.E {
	e.dbpool.Close()
	return nil
}

func (e *PostgresqlLO) Init() errors.E {
	ctx := context.Background()

	dbpool, err := pgxpool.New(ctx, "postgres://test:test@localhost:5432")
	if err != nil {
		return errors.WithStack(err)
	}
	var maxConnectionsStr string
	err = dbpool.QueryRow(ctx, `SHOW max_connections`).Scan(&maxConnectionsStr)
	if err != nil {
		return errors.WithStack(err)
	}
	maxConnections, err := strconv.Atoi(maxConnectionsStr)
	if err != nil {
		return errors.WithStack(err)
	}
	if maxConnections < 3 {
		return errors.New("max_connections too low")
	}
	_, err = dbpool.Exec(ctx, `CREATE SEQUENCE kv_value_seq`)
	if err != nil {
		return errors.WithStack(err)
	}
	_, err = dbpool.Exec(ctx, `CREATE TABLE kv (key BYTEA PRIMARY KEY NOT NULL, value OID NOT NULL DEFAULT nextval('kv_value_seq'))`)
	if err != nil {
		return errors.WithStack(err)
	}
	_, err = dbpool.Exec(ctx, `ALTER SEQUENCE kv_value_seq OWNED BY kv.value`)
	if err != nil {
		return errors.WithStack(err)
	}
	e.dbpool = dbpool
	return nil
}

func (e *PostgresqlLO) Put(key []byte, value []byte) (errE errors.E) {
	ctx := context.Background()

	tx, err := e.dbpool.BeginTx(ctx, pgx.TxOptions{
		IsoLevel:   pgx.Serializable,
		AccessMode: pgx.ReadWrite,
	})
	if err != nil {
		return errors.WithStack(err)
	}
	defer tx.Rollback(ctx)

	var oid uint32
	var inserted bool
	err = tx.QueryRow(ctx,
		`WITH
			existing AS (
				SELECT value FROM kv WHERE key=$1
			),
			inserted AS (
				INSERT INTO kv (key)
				SELECT $1 WHERE NOT EXISTS (SELECT FROM existing)
				RETURNING value
			)
		SELECT value, true FROM inserted
		UNION ALL
		SELECT value, false FROM existing`,
		key,
	).Scan(&oid, &inserted)
	if err != nil {
		return errors.WithStack(err)
	}

	largeObjects := tx.LargeObjects()
	if inserted {
		_, err := largeObjects.Create(ctx, oid)
		if err != nil {
			return errors.WithStack(err)
		}
	}

	lo, err := largeObjects.Open(ctx, oid, pgx.LargeObjectModeWrite)
	if err != nil {
		return errors.WithStack(err)
	}
	// We do not need to defer lo.Close() because any large object descriptors
	// that remain open at the end of a transaction are closed automatically.

	_, err = io.Copy(lo, bytes.NewReader(value))
	if err != nil {
		return errors.WithStack(err)
	}

	return errors.WithStack(tx.Commit(ctx))
}

func main() {
	engine := &PostgresqlLO{}

	errE := engine.Init()
	if errE != nil {
		log.Fatal(errE)
	}
	defer engine.Close()

	errE = engine.Put([]byte("abcd"), []byte("xxx"))
	if errE != nil {
		log.Fatal(errE)
	}

	errE = engine.Put([]byte("abcf"), make([]byte, 3*1024*1024*1024))
	if errE != nil {
		log.Fatal(errE)
	}
}
