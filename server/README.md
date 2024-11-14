## Estrutura de pastas

```
Server
├───cmd
│   └───api
│       ├───handlers
│       ├───middlewares
│       └───routes
├───internals
│   ├───database
│   ├───models
│   ├───repositories
│   └───services
└───pkg
    ├───errors
    ├───logs
    └───utils
```

## Criação do banco de dados

Crie um banco de dados Postgres da maneira que preferir, recomendo criar um usando Docker:

```
$ docker run --name=database -e POSTGRES_PASSWORD=senha_do_seu_banco_de_dados -d -p 5432:5432 postgres
```

## Variáveis de ambiente

Para se conectar ao banco de dados e assinar os tokens de autenticação é necessário criar um arquivo `.env` na raiz do projeto e definir algumas variáveis de ambiente, exemplo:

```
CONNECTION_STRING="user=postgres password=password host=localhost port=5432 dbname=postgres sslmode=disable"
ACCESS_SECRET=c1260015a47cb673a75577a7af075c9ff968ace63d30c59d2bdde25ff904ff94
REFRESH_SECRET=c1260015a47cb673a75577a7af075c9ff968ace63d30c59d2bdde25ff904ff94
```

**CONNECTION_STRING:** Substitua com as informações do seu banco de dados (usuário, senha, host e nome do banco), a porta é 5432, padrão do Postgres.

**ACCESS_SECRET e REFRESH_SECRET:** Use uma string aleatória forte, essas variáveis serão usadas para assinar os token JWT. Você pode gerar essas strings seguindo esse [passo a passo](https://mojitocoder.medium.com/generate-a-random-jwt-secret-22a89e8be00d).

## Executando a aplicação

**Compilando pelo terminal:**

```
$ go run ./cmd/api/main.go
```

**Gerando executável (Linux):**

```
$ go build -o ./cmd/api/main ./cmd/api
$ ./cmd/api/main
```

**Usando Docker:**

```
$ docker build -t server .
$ docker run --name=server -p 8080:8080 golang
```

Caso queira executar o container em segundo plano adicione a flag `-d` antes de `-p`.
