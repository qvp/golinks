up:
	docker compose -f ./configs/docker-compose.yaml up
down:
	docker compose -f ./configs/docker-compose.yaml down
build:
	docker compose -f ./configs/docker-compose.yaml build
logs:
	docker compose -f ./configs/docker-compose.yaml logs api
exec:
	docker compose -f ./configs/docker-compose.yaml exec api bash
swag:
	swag init -g ./cmd/rest/main.go -o ./docs/swagger
rest:
	go run ./cmd/rest
worker:
	go run ./cmd/worker
grpc:
	go run ./cmd/grpc
proto:
	protoc --go_out=. --go-grpc_out=. ./internal/grpc/golinks.proto

sqlc:
	sqlc generate

migrate-create:
	migrate create -ext sql -dir internal/db/migrations -seq $(1)

migrate-up:
	migrate -database postgres://golinks:golinks@localhost:5432/golinks?sslmode=disable -path internal/db/migrations up

migrate-down:
	migrate -database postgres://golinks:golinks@localhost:5432/golinks?sslmode=disable -path internal/db/migrations down 1