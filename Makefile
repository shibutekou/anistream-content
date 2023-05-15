build:
	go build -o anigo cmd/main.go

run:
	go run cmd/main.go

docker:
	docker build --tag anigo .
	docker run -e KODIK_TOKEN=${KODIK_TOKEN} --network=host anigo