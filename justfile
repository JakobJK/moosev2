# Default command: show help
set shell := ["powershell.exe", "-Command"]

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

# View logs for the API specifically
logs-api:
    docker logs -f moosev2-api

# --- DATABASE COMMANDS ---

db:
    $env:PGPASSWORD='moose'; pgcli -h 127.0.0.1 -p 5432 -u moose -d moose_db

create-migration name:
    docker-compose run --rm --entrypoint migrate migrate create -ext sql -dir ./migrations -seq {{name}}

# Run migrations
migrate:
    docker-compose run --rm migrate

# Generate sqlc code using Docker
sqlc:
    docker-compose run --rm api sqlc generate
