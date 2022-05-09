postgres:
	docker run --name postgres14 -p 5431:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14.2-alpine

createdb:
	docker exec -it postgres14  createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres14  dropdb simple_bank

migrateup:
	migrate -path db/migrations -database "postgresql://root:secret@localhost:5431/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations -database "postgresql://root:secret@localhost:5431/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...
server:
	go run main.go

build-image-app:
	docker build -t simple_bank:latest . 

run-container-app:
	docker run --name simple_bank -e GIN_MODE=release --network bank_network -e DB_SOURCE="postgresql://root:secret@postgres14:5432/simple_bank?sslmode=disable" -p 8080:8080 simple_bank:latest

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server build-image run-container