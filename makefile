#===========================================================
# edit this to match your migration dir
MIGRATION_DIR=./migrations
# ==========================================================
# edit this fields to match your database environment
# must match with app database
# user name of database
POSTGRES_USER=neptune
# password of database
POSTGRES_PASSWORD=neptune
# database to apply migrations
POSTGRES_DB=neptune
# port of database
POSTGRES_PORT=5432
# hostname of database
POSTGRES_HOST=localhost
# ============================================
# used by goose, do not edit these fields
export GOOSE_DBSTRING=user=$(POSTGRES_USER) password=$(POSTGRES_PASSWORD) host=$(POSTGRES_HOST) port=$(POSTGRES_PORT) dbname=$(POSTGRES_DB)
export GOOSE_DRIVER=postgres
# ==========================================================
# commands
# spin up database and other services on docker compose
spinup:
	bash -c "sudo docker compose --project-directory=infrastructure up -d"
# ==========================================================
# spin down
spindown:
	bash -c "sudo docker compose --project-directory=infrastructure down"
# ==========================================================
# start app on dev mode with hot reload
dev:
	bash -c "air -c .air.toml"
	# ==========================================================
# create new migration file
create:
	bash -c "goose -dir=$(MIGRATION_DIR) create $(name) sql"
# ==========================================================
# Migrate the DB to the most recent version available
up:
	bash -c "goose -dir=$(MIGRATION_DIR) up"
# ==========================================================
# Migrate the DB up by 1
up-by-one:
	bash -c "goose -dir=$(MIGRATION_DIR) up-by-one"
# ==========================================================
#  Migrate the DB to a specific VERSION
up-to:
	bash -c "goose -dir=$(MIGRATION_DIR) up-to $(ver)"
# ==========================================================
# Roll back the version by 1
down:
	bash -c "goose -dir=$(MIGRATION_DIR) down"
# ==========================================================
# Roll back to a specific VERSION
down-to:
	bash -c "goose -dir=$(MIGRATION_DIR) down-to $(ver)"
# ==========================================================
#  Re-run the latest migration
redo:
	bash -c "goose -dir=$(MIGRATION_DIR) redo"
# ==========================================================
# Roll back all migrations
reset:
	bash -c "goose -dir=$(MIGRATION_DIR) reset"
# ==========================================================
# Dump the migration status for the current DB
status:
	bash -c "goose -dir=$(MIGRATION_DIR) status"
# ==========================================================
# Print the current version of the database
version:
	bash -c "goose -dir=$(MIGRATION_DIR) version"
# ==========================================================
# Apply sequential ordering to migrations
fix:
	bash -c "goose -dir=$(MIGRATION_DIR) fix"
# ==========================================================
# Check migration files without running them
validate:
	bash -c "goose -dir=$(MIGRATION_DIR) validate"
# ==========================================================
# tests migrations on database
test_db:
	bash -c "make spinup || make up && make down"
# ==========================================================
# run unit tests
test_unit:
	bash -c "export GO_ENV="test" && go test ./..."
# ==========================================================
# run code coverage tests
test_cov:
	bash -c "export GO_ENV="test" && go test -v -coverprofile coverage/cover.out ./internal/... ./pkg/..."
	bash -c "export GO_ENV="test" && go tool cover -html coverage/cover.out -o coverage/cover.html"
	# ==========================================================
# Seeds database with data
seed:
	bash -c "bash seeds/account.sh POSTGRES_USER=$(POSTGRES_USER) POSTGRES_PASSWORD=$(POSTGRES_PASSWORD) \
                POSTGRES_HOST=$(POSTGRES_HOST) \
                POSTGRES_PORT=$(POSTGRES_PORT) \
                POSTGRES_DB=$(POSTGRES_DB)"