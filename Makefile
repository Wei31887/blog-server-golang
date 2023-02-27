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

proto:
	rm -f pb/*.go
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	proto/*.proto

evans:
	evans --host localhost --port 8888 -r repl

.PHONY: postgres createdb dropdb migrateup migrateup1 migratedown migratedown1 proto evans