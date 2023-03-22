include .env

run:
	go run cmd/main.go

tidy:
	go mod tidy

migrate-new-postgres:
	migrate create -ext sql -dir $(POSTGRES_MIGRATE_PATH) $(FILENAME)

migrate-up-postgres:
	migrate -path $(POSTGRES_MIGRATE_PATH) -database $(POSTGRES_CONNECTION_URL) -verbose up

migrate-down-postgres:
	migrate -path $(POSTGRES_MIGRATE_PATH) -database $(POSTGRES_CONNECTION_URL) -verbose down 

migrate-force-postgres:
	migrate -path $(POSTGRES_MIGRATE_PATH) -database $(POSTGRES_CONNECTION_URL) -verbose force $(DIRTY)

migrate-drop:
	migrate -path $(POSTGRES_MIGRATE_PATH) -database $( POSTGRES_CONNECTION_URL) -verbose drop

migrate-fresh-postgres: migrate-down-postgres migrate-up-postgres

sqlc:
	sqlc -f sqlc.yaml generate

.PHONY:	migrate-up migrate-down sqlc migrate-fresh run
