postgres:
	docker container run --name postgres15 -p 5430:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:15-alpine

createdb:
	docker exec -it postgres15 createdb --username=root --owner=root go_blog

dropdb:
	docker exec -it postgres15 dropdb go_blog

migrateup:
	migrate -path db/migrations -database "postgres://root:secret@localhost:5430/go_blog?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migrations -database "postgres://root:secret@localhost:5430/go_blog?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migrations -database "postgres://root:secret@localhost:5430/go_blog?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migrations -database "postgres://root:secret@localhost:5430/go_blog?sslmode=disable" -verbose down 1

.PHONY: postgres createdb dropdb migrateup migratedown