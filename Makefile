DB_URL=postgresql://root:123456@localhost:5432/travor-backend?sslmode=disable

# Run dev server
dev:
	air

# Run server
server:
	go run main.go

# Run postgres
postgres:
	docker run --name travor-backend -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=123456 -d postgres:12-alpine

# Run migrate db
migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

init-migrate:
	migrate create -ext sql -dir db/migration -seq

# Create swaggger
swagger:
	swag init

docker-build:
	docker build . -t travor-backend 
	docker run -p 8088:8088 travor-backend --rm

protoc:
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
	proto/*.proto

.PHONY: dev server migrateup migratedown init-migrate postgres swagger protoc