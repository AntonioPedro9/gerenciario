## Go Lang REST API

CRUD de usuários criado em Go no formato de API RESTful usando banco de dados Postgres.

## Setup

Crie um banco de dados Postgres na sua máquina e crie a tabela "users":

```
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);
```

Em seguida crie um arquivo `.env` e preencha as informações do seu banco de dados de produção e de testes:

```
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

Executando testes E2E:

```
$ go test ./...
```

Executando a aplicação:

```
$ go run main.go
```

**OBS:** Porta padrão da API é a 8080 mas pode ser alterada no arquivo `main.go`.
