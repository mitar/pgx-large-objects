Reproduction:

```sh
$ docker run --name pgsql --rm -d --network=host -e LOG_TO_STDOUT=1 -e PGSQL_ROLE_1_USERNAME=test -e PGSQL_ROLE_1_PASSWORD=test -e PGSQL_DB_1_NAME=test -e PGSQL_DB_1_OWNER=test registry.gitlab.com/tozd/docker/postgresql:16
$ go run ./...
```

Output:

```
2024/01/04 21:42:04 write failed: write tcp 127.0.0.1:45636->127.0.0.1:5432: write: connection reset by peer
```

Output from `docker logs pgsql`:

```
{"level":"LOG","msg":"checkpoint starting: shutdown immediate","pid":36,"time":"2024-01-04T20:41:58.626Z","service":"postgresql","stage":"run","logged":"2024-01-04T20:41:58.626Z"}
{"level":"LOG","msg":"checkpoint complete: wrote 7 buffers (0.0%); 0 WAL file(s) added, 0 removed, 0 recycled; write=0.004 s, sync=0.008 s, total=0.022 s; sync files=6, longest=0.003 s, average=0.002 s; distance=3 kB, estimate=3 kB; lsn=0/14E8610, redo lsn=0/14E8610","pid":36,"time":"2024-01-04T20:41:58.644Z","service":"postgresql","stage":"run","logged":"2024-01-04T20:41:58.644Z"}
{"level":"LOG","msg":"checkpoint starting: shutdown immediate","pid":37,"time":"2024-01-04T20:41:58.733Z","service":"postgresql","stage":"run","logged":"2024-01-04T20:41:58.733Z"}
{"level":"LOG","msg":"checkpoint complete: wrote 929 buffers (1.9%); 0 WAL file(s) added, 0 removed, 0 recycled; write=0.013 s, sync=0.025 s, total=0.047 s; sync files=305, longest=0.004 s, average=0.001 s; distance=4252 kB, estimate=4252 kB; lsn=0/190F900, redo lsn=0/190F900","pid":37,"time":"2024-01-04T20:41:58.777Z","service":"postgresql","stage":"run","logged":"2024-01-04T20:41:58.777Z"}
{"level":"LOG","msg":"starting PostgreSQL 16.1 (Ubuntu 16.1-1.pgdg22.04+1) on x86_64-pc-linux-gnu, compiled by gcc (Ubuntu 11.4.0-1ubuntu1~22.04) 11.4.0, 64-bit","pid":15,"time":"2024-01-04T20:41:58.833Z","service":"postgresql","stage":"run","logged":"2024-01-04T20:41:58.834Z"}
{"level":"LOG","msg":"listening on IPv4 address \"0.0.0.0\", port 5432","pid":15,"time":"2024-01-04T20:41:58.834Z","service":"postgresql","stage":"run","logged":"2024-01-04T20:41:58.834Z"}
{"level":"LOG","msg":"listening on IPv6 address \"::\", port 5432","pid":15,"time":"2024-01-04T20:41:58.834Z","service":"postgresql","stage":"run","logged":"2024-01-04T20:41:58.834Z"}
{"level":"LOG","msg":"listening on Unix socket \"/var/run/postgresql/.s.PGSQL.5432\"","pid":15,"time":"2024-01-04T20:41:58.839Z","service":"postgresql","stage":"run","logged":"2024-01-04T20:41:58.839Z"}
{"level":"LOG","msg":"database system was shut down at 2024-01-04 20:41:58 UTC","pid":40,"time":"2024-01-04T20:41:58.844Z","service":"postgresql","stage":"run","logged":"2024-01-04T20:41:58.844Z"}
{"level":"LOG","msg":"database system is ready to accept connections","pid":15,"time":"2024-01-04T20:41:58.850Z","service":"postgresql","stage":"run","logged":"2024-01-04T20:41:58.850Z"}
{"database":"test","level":"LOG","msg":"invalid message length","pid":44,"time":"2024-01-04T20:42:04.533Z","user":"test","service":"postgresql","stage":"run","logged":"2024-01-04T20:42:04.533Z"}
{"database":"[unknown]","level":"LOG","msg":"PID 44 in cancel request did not match any process","pid":45,"time":"2024-01-04T20:42:04.537Z","user":"[unknown]","service":"postgresql","stage":"run","logged":"2024-01-04T20:42:04.537Z"}
```
