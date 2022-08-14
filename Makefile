postgres12:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres12 dropdb simple_bank

migrateup:
	migrate -path D:/Go-code/simplebank/db/migration/ -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path D:/Go-code/simplebank/db/migration/ -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlcversion:
	docker run --rm -v "%cd%:/src" -w /src kjconroy/sqlc version

sqlcinit:
	docker run --rm -v "%cd%:/src" -w /src kjconroy/sqlc init

sqlcgenerate:
	docker run --rm -v "%cd%:/src" -w /src kjconroy/sqlc generate

test:
	go test -v -cover ./...

.PHONY: postgres12 createdb dropdb migrateup migratedown sqlcversion sqlcinit sqlcversion