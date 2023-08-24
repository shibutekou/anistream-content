.PHONY: start dr dc redc protoc

start:
	go build -o bin/anistream-content cmd/main.go
	./bin/anistream-content

dr:
	docker build --tag anistream-content .
	docker run -e KODIK_TOKEN=${KODIK_TOKEN} --network=host anistream-content

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