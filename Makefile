.PHONY: start dr dc redc protoc

start:
	go build -o anistream cmd/main.go
	./anigo

dr:
	docker build --tag anistream .
	docker run -e KODIK_TOKEN=${KODIK_TOKEN} --network=host anistream

dc:
	docker compose down
	docker compose up

redc:
	docker compose down
	docker compose up --build

protoc:
	protoc --go_out=./internal/controller/grpc/pb --go_opt=paths=source_relative \
    --go-grpc_out=./internal/controller/grpc/pb --go-grpc_opt=paths=source_relative \
    --proto_path=./proto proto/content.proto