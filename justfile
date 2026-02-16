# Default command: show help

default:
    @just --list

# --- DOCKER COMMANDS ---
# Spin up the whole stack (Go API, Postgres, SvelteKit)
up:
    docker-compose up --build

# Stop and remove containers
down:
    docker-compose down

# Force a clean rebuild (fixes that Vite/npm cache issue)
rebuild:
    docker-compose up --build --force-recreate


## Open interactive shell (using local pgcli)
db:
    pgcli postgres://moose:moose@127.0.0.1:5432/moose_db

create-migration name:
    docker-compose run --rm --entrypoint migrate migrate create -ext sql -dir ./migrations -seq {{name}}

# Run migrations
migrate:
    docker-compose run --rm migrate
# Wipe the entire database schema (Roll back everything)
migrate-down:
    docker-compose run --rm migrate -path=/migrations/ -database="postgres://moose:moose@db:5432/moose_db?sslmode=disable" drop -f
seed:
    docker-compose exec -T db psql "postgres://moose:moose@localhost:5432/moose_db" < ./packages/db/mockdata/0001.sql

# --- API COMMANDS --- 

# Generate Go functions for SQL queries
sqlc:
    docker-compose run --rm api sqlc generate

# Generate https certs
certs:
    mkcert -cert-file docker/nginx/certs/moose.dev.pem -key-file docker/nginx/certs/moose.dev-key.pem moose.dev api.moose.dev storybook.moose.dev localhost 127.0.0.1

