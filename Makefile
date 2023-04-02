postgres:
	docker run --name postgres-glassy-sky -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	docker exec -it postgres-glassy-sky createdb --username=root --owner=root glassy_sky

dropdb:
	docker exec -it postgres-glassy-sky dropdb glassy_sky

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/glassy_sky?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/glassy_sky?sslmode=disable" -verbose down

sqlc:
	sqlc generate

mock:
	mockgen -package mockdb -destination internal/admin/domain/mockdb/db.go github.com/amyunfei/glassy-sky/internal/admin/domain/postgresql Repository

swag:
	swag init -g ./cmd/main.go -o ./api

swagtots:
	npx swagger-typescript-api -p ./api/swagger.json -o ./web/admin/src/api -n dto.ts --no-client

test:
	go test -v -cover ./...

.PHONY: postgres createdb dropdb migrateup migratedown sqlc mock swag test swagtots