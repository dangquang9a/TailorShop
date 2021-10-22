network:
	docker network create tailorshop-network

postgres:
	docker run --name postgres14 --network tailorshop-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root tailor_shop

dropdb:
	docker exec -it postgres12 dropdb tailor_shop

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/tailor_shop?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/tailor_shop?sslmode=disable" -verbose down


sqlc:
	sqlc generate