
MIGRATION_DIR=./db/migrations
DATABASE="postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable"

test-migration:
	echo $(DATABASE)

create-migration:
	migrate create -ext sql -dir $(MIGRATION_DIR) -seq $(name)

migrate-up:
	migrate -path $(MIGRATION_DIR) -database $(DATABASE) -verbose up

migrate-force:
	migrate -path $(MIGRATION_DIR) -database $(DATABASE) force $(VERSION)
