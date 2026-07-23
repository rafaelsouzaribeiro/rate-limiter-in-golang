# Rate Limiter em Go (MBA desafio técnico)

## Objetivo

Rate Limiter em Go que funcione como um middleware para controlar o fluxo de requisições de um serviço web.

O sistema é capaz de limitar o tráfego com base no IP do solicitante, utilizando o Redis para persistência e orquestração.

## Variáveis de ambiente

| Variável         | Descrição                                          | Exemplo       |
|------------------|-----------------------------------------------------|---------------|
| `MAX_REQUEST`    | Número máximo de requisições permitidas             | `3`           |
| `TIME_LIMIT`     | Janela de tempo para contagem das requisições       | `60s`         |
| `BLOCK_TIME`     | Tempo de bloqueio após exceder o limite             | `60s`         |
| `HOST_REDIS`     | Host e porta do Redis                               | `localhost:6379` |
| `PASSWORD_REDIS` | Senha de acesso ao Redis                            | `123mudar`    |
| `SERVER_PORT`    | Porta em que o servidor web irá rodar               | `8080`        |

Exemplo de arquivo `.env`:

```dotenv
MAX_REQUEST=3
TIME_LIMIT=60s
BLOCK_TIME=60s
HOST_REDIS=localhost:6379
PASSWORD_REDIS=123mudar
SERVER_PORT=8080
```

## Como executar

### Opção 1: Rodando localmente (sem container)

1. Crie um arquivo `.env` dentro da pasta `cmd` com as variáveis acima.
2. Certifique-se de ter uma instância do Redis rodando localmente (ex: `redis-server` ou via Docker).
3. Acesse a pasta `cmd` e execute:

    ```sh
    cd cmd
    go run main.go
    ```

### Opção 2: Rodando com Docker Compose

1. Crie um arquivo `.env` na raiz do projeto com as variáveis acima.
2. Execute o Docker Compose:

    ```sh
    docker compose up --build
    ```

Isso irá subir dois serviços:
- **redis**: instância do Redis protegida por senha (`PASSWORD_REDIS`).
- **app**: aplicação Go que consome o Redis para controlar o rate limit.

## Testando a aplicação

Após subir a aplicação (local ou via Docker), faça requisições para o endpoint configurado:

```sh
curl -i http://localhost:8080/rate-limiter
```

Se o limite de requisições (`MAX_REQUEST`) for excedido dentro da janela de tempo (`TIME_LIMIT`), a aplicação deverá retornar um erro de limite excedido e bloquear novas requisições pelo tempo definido em `BLOCK_TIME`.