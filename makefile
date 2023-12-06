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
	bash -c "docker compose --project-directory=deployment up -d"
# ==========================================================
# spin down
spindown:
	bash -c "docker compose --project-directory=deployment down"
# ==========================================================
# start app on dev mode with hot reload
start\:dev:
	bash -c "export GO_ENV=development && wgo run  ./cmd/... ./internal/..."
	# ==========================================================
# create new migration file
migration\:create:
	bash -c "goose -dir=$(MIGRATION_DIR) create $(name) sql"
# ==========================================================
# Migrate the DB to the most recent version available
migration\:up:
	bash -c "goose -dir=$(MIGRATION_DIR) up"
# ==========================================================
# Migrate the DB up by 1
migration\:up-by-one:
	bash -c "goose -dir=$(MIGRATION_DIR) up-by-one"
# ==========================================================
#  Migrate the DB to a specific VERSION
migration\:up-to:
	bash -c "goose -dir=$(MIGRATION_DIR) up-to $(ver)"
# ==========================================================
# Roll back the version by 1
migration\:down:
	bash -c "goose -dir=$(MIGRATION_DIR) down"
# ==========================================================
# Roll back to a specific VERSION
migration\:down-to:
	bash -c "goose -dir=$(MIGRATION_DIR) down-to $(ver)"
# ==========================================================
#  Re-run the latest migration
migration\:redo:
	bash -c "goose -dir=$(MIGRATION_DIR) redo"
# ==========================================================
# Roll back all migrations
migration\:reset:
	bash -c "goose -dir=$(MIGRATION_DIR) reset"
# ==========================================================
# Dump the migration status for the current DB
migration\:status:
	bash -c "goose -dir=$(MIGRATION_DIR) status"
# ==========================================================
# Print the current version of the database
migration\:version:
	bash -c "goose -dir=$(MIGRATION_DIR) version"
# ==========================================================
# Apply sequential ordering to migrations
migration\:fix:
	bash -c "goose -dir=$(MIGRATION_DIR) fix"
# ==========================================================
# Check migration files without running them
migration\:validate:
	bash -c "goose -dir=$(MIGRATION_DIR) validate"
# ==========================================================
# tests migrations on database
test\:db:
	bash -c "make spinup || make up && make down"
# ==========================================================
# run unit tests
test\:unit:
	bash -c "export GO_ENV="test" && go test ./..."
# ==========================================================
# run unit tests watch mode
test\:unit\:watch: 
	bash -c "air -c .air.toml.test"
# ==========================================================
# run code coverage tests
test\:cov:
	bash -c "export GO_ENV="test" && go test -v -coverprofile coverage/cover.out ./internal/... ./pkg/..."
	bash -c "export GO_ENV="test" && go tool cover -html coverage/cover.out -o coverage/cover.html"
# ==========================================================
# Seeds database with data
seed:
	bash -c "bash seeds/account.sh POSTGRES_USER=$(POSTGRES_USER) POSTGRES_PASSWORD=$(POSTGRES_PASSWORD) \
                POSTGRES_HOST=$(POSTGRES_HOST) \
                POSTGRES_PORT=$(POSTGRES_PORT) \
                POSTGRES_DB=$(POSTGRES_DB)"