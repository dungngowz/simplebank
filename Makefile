DB_URL=postgresql://root:secret@localhost:5431/simple_bank?sslmode=disable

postgres:
	docker run --name postgres14 -p 5431:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14.2-alpine

createdb:
	docker exec -it postgres14  createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres14  dropdb simple_bank

migrateup:
	migrate -path db/migrations -database "$(DB_URL)" -verbose up

migrateup_1:
	migrate -path db/migrations -database "$(DB_URL)" -verbose up 1

migratedown:
	migrate -path db/migrations -database "$(DB_URL)" -verbose down

migratedown_1:
	migrate -path db/migrations -database "$(DB_URL)" -verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

dev-server:
	nodemon --exec go run main.go --signal SIGTERM

build-image-app:
	docker build -t simple_bank:latest . 

run-container-app:
	docker run --name simple_bank -e GIN_MODE=release --network bank_network -e DB_SOURCE="postgresql://root:secret@postgres14:5432/simple_bank?sslmode=disable" -p 8080:8080 simple_bank:latest

db_doc:
	dbdocs build doc/db.dbml   

db_schema:
	dbml2sql --postgres -o doc/schema.sql doc/db.dbml 

.PHONY: postgres createdb dropdb migrateup migrateup_1 migratedown migratedown_1 sqlc test server build-image run-container db_doc db_schema