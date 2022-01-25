.PHONY: createdb dropdb migrateup migratedown postgres test server

createdb: 
	docker exec -it go-api_postgres createdb --username=root --owner=root go-api

dropdb: 
	docker exec -it go-api_postgres dropdb go-api

migrateup:
	migrate -path db/migration -database "postgresql://root:password@localhost:5432/go-api?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:password@localhost:5432/go-api?sslmode=disable" -verbose down

postgres:
	docker run --name go-api_postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -e POSTGRES_DB=lendspace -d postgres:14-alpine

test:
	go test -v -cover ./internal/user/repository ./internal/space/repository

server:
	go run main.go