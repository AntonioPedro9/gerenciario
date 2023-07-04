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

Dentro do arquivo `/database/connection.go` altere a variável "connStr" com o usuário, nome, senha, host e porta do seu banco de dados.

Em seguida execute os seguintes comandos dentro do diretório do projeto:

```
$ go mod tidy
$ go run main.go
```

**OBS:** Porta padrão da API é a 8080 mas pode ser alterada no arquivo `main.go`.
