postgres:
	docker run --name postgres12 -p 5433:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret  -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres12 dropdb simple_bank

up:
	migrate -path db/migration/ -database "postgresql://root:secret@localhost:5433/simple_bank?sslmode=disable" -verbose up

down:
	migrate -path db/migration/ -database "postgresql://root:secret@localhost:5433/simple_bank?sslmode=disable" -verbose down

.PHONY: postgres createdb dropdb up down