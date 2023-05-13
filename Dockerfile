FROM golang:1.20
LABEL authors="vgekko"

WORKDIR /app
COPY go.mod go.sum ./

RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o ani-go cmd/main.go

EXPOSE 8800
CMD ["./ani-go"]