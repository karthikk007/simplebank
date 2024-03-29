# instructions

## design the DB

use db diagram - [LINK](https://dbdiagram.io/)

Download the SQL and diagram

## install the following

1. docker desktop - [LINK](https://hub.docker.com/)
2. table plus - [LINK](https://tableplus.com/)
3. Postgres - [LINK](https://hub.docker.com/_/postgres)
4. Create DB schema

## postgres container

### pull the image

```
docker pull postgres:16-alpine

docker images
```

### start a postgres intance

```
docker run --name postgres16 -p 5432:5432 -e POSTGRES_USER=admin -e POSTGRES_PASSWORD=admin -d postgres:16-alpine
```

### check docker ps

```
docker ps
```

### connect to container

```
docker exec -it postgres16 psql -U admin
```

### container commands or queries

```
select now();


quit - \q
```

### container logs

```
docker logs postgres16
```

## table plus

1. launch app
2. create new connection

```
name: postgres16
User: admin
Password: admin
Database: admin
```

3. load sql file
4. cmd + a
5. cmd + enter

## DB schema migration

use golang migrate library [LINK](https://github.com/golang-migrate/migrate)

CLI DOCUMENT [REFER](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)

```
brew install golang-migrate

migrate -version

mkdir -p db/migration


migrate create -ext sql -dir db/migration -seq init_schema



```

## docker commands

```
docker run
docker ps
docker ps -a
docker stop postgres16
docker start postgres16


docker rm postgres16
```

### container shell

```
docker exec -it postgres16 /bin/sh


docker exec -it postgres16 createdb --username=admin --owner=admin simple_bank


docker exec -it postgres16 psql -U admin simple_bank

```

### create db in container

```
createdb --username=admin --owner=admin simple_bank

psql simple_bank
psql simple_bank --username=admin

dropdb simple_bank
dropdb sample_bank --username=admin

exit
```

## Project setup

### make file

```
postgres:
	docker run --name postgres16 -p 5432:5432 -e POSTGRES_USER=admin -e POSTGRES_PASSWORD=admin -d postgres:16-alpine

createdb:
	docker exec -it postgres16 createdb --username=admin --owner=admin simple_bank

dropdb:
	docker exec -it postgres16 dropdb --username=admin simple_bank

.PHONY: postgres createdb
```

run in terminal

```
make postgres
```

## start the migration

```
migrate -path db/migration -database "postgresql://admin:admin@localhost:5432/simple_bank?sslmode=disable" -verbose up
```

## DB CURD LIBRATIES

-   Ddatabase/sql
-   GROM
-   SQLX
-   [SQLC](https://docs.sqlc.dev/en/latest/overview/install.html)

### INSTALL SQLC

```
brew install sqlc

add sqlc.yaml

add sqlc command to make file

make sqlc
```

the yaml file

```
version: "1"
packages:
    - name: "db"
      path: "./db/sqlc"
      queries: "./db/query/"
      schema: "./db/migration/"
      engine: "postgresql"
      emit_json_tags: true
      emit_prepared_queries: false
      emit_interface: false
      emit_exact_table_names: false
```

## init go mod

```
go mod init example.com/simplebank
go mod tidy
```

# UNIT TESTING DB

use [PQ driver](https://pkg.go.dev/github.com/lib/pq)

```
go get github.com/lib/pq
```

## checking test results

[testify](https://github.com/stretchr/testify)

```
go get github.com/stretchr/testify

```

# GITHUB WORKFLOWS

```
mkdir -p .github/workflows
touch .github/workflows/ci.yaml

```

# WEB FRAMEWORK

```
https://github.com/gin-gonic/gin

go get -u github.com/gin-gonic/gin
```

# VIPER

ENV READ

```
https://github.com/spf13/viper

override
SERVER_ADDRESS=0.0.0.0:8081 make server
```

# GO MOCK

go get github.com/golang/mock/mockgen@v1.6.0

https://github.com/golang/mock

go install github.com/golang/mock/mockgen@v1.6.0

```
ls -l ~/go/bin
vim ~/.zshrc
export PATH=$PATH:~/go/bin
which mockgen

mockgen -package mockdb -destination db/mock/store.go example.com/simplebank/db/sqlc Store
```

## packages

### UUID

https://github.com/google/uuid

go get github.com/google/uuid

### jwt

https://github.com/dgrijalva/jwt-go

go get github.com/dgrijalva/jwt-go

### PASETO

https://github.com/o1egl/paseto

go get github.com/o1egl/paseto

# REFER

[YOUTUBE](https://www.youtube.com/watch?v=phHDfOHB2PU&list=PLy_6D98if3ULEtXtNSY_2qN21VCKgoQAE&index=6)
[GITHUB](https://github.com/techschool/simplebank)

# VSCODE GO TEST VERBOSE

```
> OPEN SETTINGS JSON
"go.testFlags": ["-v", "-race"]
```

# Docker build commands

```
docker build -t simplebank:latest .
docker images
docker run --name simplebank -p 8080:8080 -e GIN_MODE=release simplebank:latest
docker rmi simplebank
docker rm simplebank

docker container inspect postgres16

docker network ls
docker network inspect bridge
```

# docker networking

```
docker network create bank-network
docker network connect bank-network postgres16
docker network inspect bank-network
docker run --name simplebank --network bank-network -p 8080:8080 -e GIN_MODE=release -e DB_SOURCE="postgresql://admin:admin@postgres16:5432/simple_bank?sslmode=disable" simplebank:latest
```
