ifneq (,$(wildcard ./.env))
  include .env
  export
endif

COMPOSE_FILE = docker-compose.yml

.PHONY: help build up down postgresql createdb migrationup migrationdown sqlc test server

help:
	@echo "Targets: build, up, down, db-up, createdb, migrationup, migrationdown, sqlc, test, server"

build: ## Builda todas as imagens Docker
	docker compose -f $(COMPOSE_FILE) --build

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

test:
	go test -v -cover ./...

server:
	go run cmd/transaction-api/main.go
