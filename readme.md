# Setting up Environment Variables
Make sure to set the environment variables before starting:
```bash
export DB_PORT=5432
export DB_HOST=localhost
export DB_USER=bstack_user
export DB_PASSWORD=password
export DB_NAME=bstack_db

export JWT_SECRET=partycat

```

# Golang Migrate
## Using Golang Migrate to Setup Database
```bash
migrate -path db/migrations/ -database "postgres://bstack_user:password@localhost:5432/bstack_db?sslmode=disable" -verbose up
```

## Using Golang Migrate to Create Migrations
```bash
migrate create -ext sql -dir db/migrations -seq <name_of_migration>
```

## Using Makefile to do migrations
```bash
make create-migration name=<migration_name>
```

## Using Makefile migrate up
```bash
make migrate-up
```
