.PHONY: start dr dc redc

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