up:
	docker compose -f ./deployments/docker-compose.yaml up
down:
	docker compose -f ./deployments/docker-compose.yaml down
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