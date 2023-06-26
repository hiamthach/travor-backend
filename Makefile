DB_URL=postgresql://root:123456@localhost:5432/travor-backend?sslmode=disable

dev:
	air

server:
	go run main.go

postgres:
	docker run --name travor-backend -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=123456 -d postgres:12-alpine

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

init-migrate:
	migrate create -ext sql -dir db/migration -seq

swagger:
	swag init

docker-build:
	docker build . -t travor-backend 
	docker run -p 8080:8080 travor-backend --rm

.PHONY: dev server migrateup migratedown init-migrate postgres swagger