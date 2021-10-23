include app.env
export

network:
	docker network create tailorshop-network

postgres:
	docker run --name postgres14 --network tailorshop-network -p 5432:5432 -e POSTGRES_USER=merlin -e POSTGRES_PASSWORD=@merlin123 -d postgres:14

createdb:
	docker exec -it postgres14 createdb --username=merlin --owner=merlin tailor_shop

dropdb:
	docker exec -it postgres14 dropdb tailor_shop

migrateup:
	migrate -path db/migration -database "${DB_SOURCE}" -verbose up

migratedown:
	migrate -path db/migration -database "${DB_SOURCE}" -verbose down

server:
	go run main.go
test:
	go test -v -cover ./...
sqlc:
	sqlc generate

.PHONY: network postgres createdb dropdb migrateup migratedown sqlc server