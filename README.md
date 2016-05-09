# dali-server

Server for Project Dali


## Build

### Go

- Install [gvm](https://github.com/moovweb/gvm)
- `> gvm install go1.6 && gvm use go1.6`
- `> go get -u github.com/soroushjp/dali-server`

### PostgreSQL

- Install and run PostgreSQL: http://www.postgresql.org/download/
- `> go get bitbucket.org/liamstask/goose/cmd/goose`
- `> goose up`

### Env variables

| name         | type   | description            | default value (if any) |
|:-------------|:-------|:-----------------------|:-----------------------|
| DALI_PORT    | uint   | Server port            |                        |
| DALI_DB_NAME | string | Postgres database name |                        |
| DALI_DB_USER | string | Postgres database user |                        |
