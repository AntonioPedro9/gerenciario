## Go Lang REST API

CRUD de usuários criado em Go no formato de API RESTful usando banco de dados Postgres.

## Setup

Crie um arquivo `.env` na raiz do projeto e preencha as informações do seu banco de dados de produção e de testes:

```
PORT=8080

DB_USER=
DB_PASSWORD=
DB_HOST=
DB_PORT=
DB_NAME=

TEST_DB_USER=
TEST_DB_PASSWORD=
TEST_DB_HOST=
TEST_DB_PORT=
TEST_DB_NAME=
```

Executando testes:

```
$ go test ./...
```

Executando a aplicação:

```
$ go run main.go
```
