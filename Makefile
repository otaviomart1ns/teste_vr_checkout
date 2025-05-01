ifneq (,$(wildcard ./.env))
  include .env
  export
endif

COMPOSE_FILE = docker-compose.yaml

.PHONY: help build up down postgresql createdb migrationup migrationdown sqlc test-backend test-frontend server

help:
	@echo "Targets: build, up, down, db-up, createdb, migrationup, migrationdown, sqlc, test-backend, test-frontend, server"

build: ## Builda todas as imagens Docker
	docker compose -f $(COMPOSE_FILE) build

up: ## Sobe containers (detached)
	docker compose -f $(COMPOSE_FILE) up -d

down: ## Desliga e remove containers
	docker compose -f $(COMPOSE_FILE) down --remove-orphans


db-up: ## Sobe somente o serviço de banco
	docker compose -f $(COMPOSE_FILE) up -d db

createdb: ## Cria base de dados dentro do container
	docker exec -it teste_vr_checkout-db-1 createdb \
		--username=$(POSTGRES_USER) \
		--owner=$(POSTGRES_USER) \
		$(POSTGRES_DB)

migrationup: ## Aplica migrações do banco
	migrate -path backend/internal/infra/db/migrations \
	  -database "postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@localhost:$(POSTGRES_PORT)/$(POSTGRES_DB)?sslmode=disable" \
	  -verbose up

migrationdown: ## Reverte migrações
	migrate -path backend/internal/infra/db/migrations \
	  -database "postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@localhost:$(POSTGRES_PORT)/$(POSTGRES_DB)?sslmode=disable" \
	  -verbose down

sqlc:
	docker run --rm -v $$(pwd):/src -w /src sqlc/sqlc generate

test-backend:
	@export POSTGRES_USER=$(POSTGRES_USER); \
	export POSTGRES_PASSWORD=$(POSTGRES_PASSWORD); \
	export POSTGRES_DB=$(POSTGRES_DB); \
	export POSTGRES_PORT=$(POSTGRES_PORT); \
	export RABBITMQ_DEFAULT_USER=$(RABBITMQ_DEFAULT_USER); \
	export RABBITMQ_DEFAULT_PASS=$(RABBITMQ_DEFAULT_PASS); \
	export RABBIT_PORT=$(RABBIT_PORT); \
	export API_PORT=$(API_PORT); \
	export FLUTTER_PORT=$(FLUTTER_PORT); \
	export TREASURY_API_BASE_URL=$(TREASURY_API_BASE_URL); \
	export TREASURY_API_ENDPOINT=$(TREASURY_API_ENDPOINT); \
	export GIN_MODE=$(GIN_MODE); \
	export POSTGRES_URL=postgresql://$${POSTGRES_USER}:$${POSTGRES_PASSWORD}@localhost:$${POSTGRES_PORT}/$${POSTGRES_DB}; \
	export RABBITMQ_URL=amqp://$${RABBITMQ_DEFAULT_USER}:$${RABBITMQ_DEFAULT_PASS}@localhost:$${RABBIT_PORT}/; \
	cd backend && go test -v -coverprofile=coverage.out ./...

test-frontend:
	@echo "API_URL=http://${HOST}:${API_PORT}" > frontend/.env
	cd frontend && flutter test --coverage \
	genhtml coverage/lcov.info -o coverage/html \
	xdg-open coverage/html/index.html

server:
	@export POSTGRES_USER=$(POSTGRES_USER); \
	export POSTGRES_PASSWORD=$(POSTGRES_PASSWORD); \
	export POSTGRES_DB=$(POSTGRES_DB); \
	export POSTGRES_PORT=$(POSTGRES_PORT); \
	export RABBITMQ_DEFAULT_USER=$(RABBITMQ_DEFAULT_USER); \
	export RABBITMQ_DEFAULT_PASS=$(RABBITMQ_DEFAULT_PASS); \
	export RABBIT_PORT=$(RABBIT_PORT); \
	export API_PORT=$(API_PORT); \
	export FLUTTER_PORT=$(FLUTTER_PORT); \
	export TREASURY_API_BASE_URL=$(TREASURY_API_BASE_URL); \
	export TREASURY_API_ENDPOINT=$(TREASURY_API_ENDPOINT); \
	export GIN_MODE=$(GIN_MODE); \
	export POSTGRES_URL=postgresql://$${POSTGRES_USER}:$${POSTGRES_PASSWORD}@localhost:$${POSTGRES_PORT}/$${POSTGRES_DB}; \
	export RABBITMQ_URL=amqp://$${RABBITMQ_DEFAULT_USER}:$${RABBITMQ_DEFAULT_PASS}@localhost:$${RABBIT_PORT}/; \
	cd backend && go run cmd/transaction-api/main.go
