## Go Lang REST API

CRUD de usuários criado em Go no formato de API RESTful usando banco de dados Postgres.

## Setup

Crie um arquivo `.env` na raiz do projeto e especifique a porta da aplicação e as informações do banco de dados

```
PORT=8080

DB_USER=
DB_PASSWORD=
DB_HOST=
DB_PORT=
DB_NAME=
```

Executando testes:

```
$ go test ./...
```

Executando a aplicação:

```
$ go run main.go
```
