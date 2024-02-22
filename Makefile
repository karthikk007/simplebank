postgres:
	docker run --name postgres16 -p 5432:5432 -e POSTGRES_USER=admin -e POSTGRES_PASSWORD=admin -d postgres:16-alpine

createdb:
	docker exec -it postgres16 createdb --username=admin --owner=admin simple_bank

dropdb:
	docker exec -it postgres16 dropdb --username=admin simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://admin:admin@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://admin:admin@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go clean -testcache && go test -v -cover ./...

.PHONY: postgres createdb dropdb migrateup migratedown sqlc
