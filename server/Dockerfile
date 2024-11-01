FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o ./cmd/api/main ./cmd/api

EXPOSE 8080

CMD ["./cmd/api/main"]
