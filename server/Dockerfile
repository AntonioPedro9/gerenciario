FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o cmd/main ./cmd

EXPOSE 8080

CMD ["./cmd/main"]
