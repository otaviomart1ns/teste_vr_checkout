# Projeto VR Software - Flutter + Golang

## 📌 Descrição

Este projeto é uma aplicação fullstack desenvolvida para o desafio técnico da **VR Software**, utilizando:

- Frontend em **Flutter Web**
- Backend em **Golang**
- Banco de dados **PostgreSQL**
- Mensageria com **RabbitMQ**
- Orquestração com **Docker Compose**
- Geração de código com **SQLC**
- Scripts e automações com **Makefile**

---

## 🚀 Tecnologias Utilizadas

| Camada       | Tecnologias                                |
|--------------|---------------------------------------------|
| Frontend     | Flutter (Web), MobX, flutter_modular       |
| Backend      | Golang, Gin, SQLC, RabbitMQ           |
| Banco de Dados | PostgreSQL                               |
| DevOps       | Docker, Docker Compose, Makefile           |
| Testes       | Go Test, Flutter Test                      |

---

## 📦 Estrutura do Projeto

```
teste_vr_checkout/
├── backend/
│   ├── cmd/transaction-api/
│   ├── internal/
│   │   ├── config/
│   │   ├── domain/
│   │   ├── infra/
│   │   ├── interfaces/
│   │   ├── pkg/
│   │   └── usecases/
│   ├── Dockerfile
│   ├── go.mod
│   └── go.sum
├── frontend/
│   ├── lib/
│   │   ├── core/
│   │   │   ├── config/
│   │   │   ├── services/
│   │   │   ├── theme/
│   │   │   ├── utils/
│   │   │   └── widgets/
│   │   └── modules/
│   │       ├── home/
│   │       └── transaction/
│   │           ├── create/
│   │           ├── pending/
│   │           ├── shared/
│   │           └── view/
│   ├── pubspec.yaml
│   └── Dockerfile
├── docker-compose.yml
├── .env
├── .env.example
├── Makefile
├── sqlc.yaml
└── README.md
```

---

## ⚙️ Requisitos

- Docker e Docker Compose
- Make (GNU Make)

---

## 🔧 Configuração `.env`

Crie um arquivo `.env` na raiz do projeto com as seguintes variáveis:

```env
# Banco de Dados
POSTGRES_USER=postgres
POSTGRES_PASSWORD=postgres
POSTGRES_DB=teste_vr_checkout
POSTGRES_PORT=5432

# RabbitMQ
RABBITMQ_DEFAULT_USER=rabbitmq
RABBITMQ_DEFAULT_PASS=rabbitmq
RABBIT_MANAGEMENT_PORT=15672
RABBIT_PORT=5672

# API
API_PORT=8080
GIN_MODE=release
HOST=localhost

# Flutter Web (Nginx)
FLUTTER_PORT=80

# API externa
TREASURY_API_BASE_URL=https://api.fiscaldata.treasury.gov/services/api/fiscal_service/
TREASURY_API_ENDPOINT=v1/accounting/od/rates_of_exchange
```

---

## 🛠️ Como Rodar o Projeto

### 1. Subir os containers

```bash
make build
make up
```

Esses comandos compilam e sobem os serviços:

- **API Go** em `http://localhost:8080`
- **Flutter Web** em `http://localhost:80`
- **PostgreSQL** em `localhost:5432`
- **RabbitMQ** em `localhost:5672` (painel: `http://localhost:15672`)

### 2. Rodar as migrações

```bash
make migrationup
```

### 3. Acessar os serviços

| Serviço       | URL                            |
|---------------|---------------------------------|
| Frontend      | http://localhost:80          |
| API Backend   | http://localhost:8080          |
| RabbitMQ UI   | http://localhost:15672 (user/pasword) |
| PostgreSQL    | via cliente na porta 5432 (user/pasword)     |

---

## 🧪 Executar Testes

### Backend - Cobertura 95%

```bash
make test-backend
```

### Frontend - Cobertura 

```bash
make test-frontend
```

---

## 📚 Endpoints da API

> Base URL: `http://localhost:8080`

> Swagger URL: `http://localhost:8080/swagger/index.html`

1. **POST** `/transactions`

Cria uma nova transação.

##### Body (JSON)

```json
{
  "description": "Moto",
  "date": "2020-01-02",
  "amount_usd": 542.96
}
```

##### Response

- **202 Accepted**


2. **GET** `/transactions/:uuid`

Retorna uma transação com valor convertido para a moeda escolhida pelo seu ID.

##### Exemplo

`/transactions/2d10af24-9aa3-41f7-a5eb-f441f755f676?currency=Brazil-Real`

##### Response

```json
{
  "id": "2d10af24-9aa3-41f7-a5eb-f441f755f676",
  "description": "Moto",
  "date": "2025-03-31",
  "amount_usd": 542.97,
  "exchange_rate": 5.758,
  "amount_converted": 3126.42,
  "to_currency": "Brazil-Real",
  "rate_date": "2025-03-31"
}
```

3. **GET** `/transactions/latest?limit=5`

Retorna as últimas transações inseridas.

##### Response
```json
[
  {
    "id": "b1661298-a7ff-4c7a-82f0-8dff8c94ff50",
    "description": "Moto",
    "date": "2020-01-02T00:00:00Z",
    "amount_usd": 542.96
  }
]
```

4. **GET** `/currencies`

Retorna lista de moedas disponíveis para conversão.

##### Response
```json
[
  "Afghanistan-Afghani",
  "Albania-Lek",
  "Brazil-Real",
  "United Kingdom-Pound",
  "..."
]
```

---

## 📥 Makefile - Comandos Úteis

```bash
make up              # Sobe os containers
make down            # Derruba os containers
make build           # Builda todas as imagens
make db-up           # Sobe apenas o banco
make createdb        # Cria o banco de dados
make migrationup     # Aplica as migrações
make migrationdown   # Reverte a última migração
make sqlc            # Gera código Go com SQLC
make test            # Executa os testes do backend
make server          # Executa a API localmente (fora do Docker)
```

---

## 🐞 Problemas Comuns

### 1. Porta em uso

Altere as portas no `.env`.

---

## 🚀 Deploy

A aplicação foi implantada e está disponível publicamente. Você pode acessá-la diretamente através do seguinte endereço:

- http://18.229.68.129:80

Esse endereço aponta para o frontend da aplicação web.
