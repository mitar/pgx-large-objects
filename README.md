Reproduction:

```sh
$ docker run --name pgsql --rm -d --network=host -e LOG_TO_STDOUT=1 -e PGSQL_ROLE_1_USERNAME=test -e PGSQL_ROLE_1_PASSWORD=test -e PGSQL_DB_1_NAME=test -e PGSQL_DB_1_OWNER=test registry.gitlab.com/tozd/docker/postgresql:16
$ go run ./...
```

Output:

```
2024/01/04 21:34:39 write failed: write tcp 127.0.0.1:56518->127.0.0.1:5432: write: connection reset by peer
exit status 1
```

Output from `docker logs pgsql`:

```
{"level":"LOG","msg":"checkpoint starting: shutdown immediate","pid":33,"time":"2024-01-04T20:34:34.050Z","service":"postgresql","stage":"run","logged":"2024-01-04T20:34:34.050Z"}
{"level":"LOG","msg":"checkpoint complete: wrote 7 buffers (0.0%); 0 WAL file(s) added, 0 removed, 0 recycled; write=0.004 s, sync=0.007 s, total=0.022 s; sync files=6, longest=0.003 s, average=0.002 s; distance=3 kB, estimate=3 kB; lsn=0/14E8610, redo lsn=0/14E8610","pid":33,"time":"2024-01-04T20:34:34.070Z","service":"postgresql","stage":"run","logged":"2024-01-04T20:34:34.070Z"}
{"level":"LOG","msg":"checkpoint starting: shutdown immediate","pid":34,"time":"2024-01-04T20:34:34.176Z","service":"postgresql","stage":"run","logged":"2024-01-04T20:34:34.176Z"}
{"level":"LOG","msg":"checkpoint complete: wrote 929 buffers (1.9%); 0 WAL file(s) added, 0 removed, 0 recycled; write=0.013 s, sync=0.024 s, total=0.046 s; sync files=305, longest=0.004 s, average=0.001 s; distance=4252 kB, estimate=4252 kB; lsn=0/190F900, redo lsn=0/190F900","pid":34,"time":"2024-01-04T20:34:34.218Z","service":"postgresql","stage":"run","logged":"2024-01-04T20:34:34.218Z"}
{"level":"LOG","msg":"starting PostgreSQL 16.1 (Ubuntu 16.1-1.pgdg22.04+1) on x86_64-pc-linux-gnu, compiled by gcc (Ubuntu 11.4.0-1ubuntu1~22.04) 11.4.0, 64-bit","pid":14,"time":"2024-01-04T20:34:34.265Z","service":"postgresql","stage":"run","logged":"2024-01-04T20:34:34.266Z"}
{"level":"LOG","msg":"listening on IPv4 address \"0.0.0.0\", port 5432","pid":14,"time":"2024-01-04T20:34:34.266Z","service":"postgresql","stage":"run","logged":"2024-01-04T20:34:34.266Z"}
{"level":"LOG","msg":"listening on IPv6 address \"::\", port 5432","pid":14,"time":"2024-01-04T20:34:34.266Z","service":"postgresql","stage":"run","logged":"2024-01-04T20:34:34.266Z"}
{"level":"LOG","msg":"listening on Unix socket \"/var/run/postgresql/.s.PGSQL.5432\"","pid":14,"time":"2024-01-04T20:34:34.270Z","service":"postgresql","stage":"run","logged":"2024-01-04T20:34:34.270Z"}
{"level":"LOG","msg":"database system was shut down at 2024-01-04 20:34:34 UTC","pid":37,"time":"2024-01-04T20:34:34.275Z","service":"postgresql","stage":"run","logged":"2024-01-04T20:34:34.275Z"}
{"level":"LOG","msg":"database system is ready to accept connections","pid":14,"time":"2024-01-04T20:34:34.281Z","service":"postgresql","stage":"run","logged":"2024-01-04T20:34:34.281Z"}
{"database":"test","level":"LOG","msg":"invalid message length","pid":41,"time":"2024-01-04T20:34:39.091Z","user":"test","service":"postgresql","stage":"run","logged":"2024-01-04T20:34:39.091Z"}
```
