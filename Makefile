# Подключаем переменные окружения для локальной инфраструктуры
include infra.env
export

# Новая миграция
migration-new:
	goose -dir="./db/migrations" create $(name) sql

# Накатить последнюю доступную версию
migration-up:
	goose $(opts) -dir ./db/migrations postgres "host=$(POSTGRES_HOST) port=$(POSTGRES_PORT) user=$(POSTGRES_USER) password=$(POSTGRES_PASSWORD) dbname=$(POSTGRES_DB) sslmode=disable" up

# Откатить версию на 1
migration-down:
	goose $(opts) -dir ./db/migrations postgres "host=$(POSTGRES_HOST) port=$(POSTGRES_PORT) user=$(POSTGRES_USER) password=$(POSTGRES_PASSWORD) dbname=$(POSTGRES_DB) sslmode=disable" down
