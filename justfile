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


# --- DATABASE COMMANDS ---
db:
    $env:PGPASSWORD='moose'; pgcli -h 127.0.0.1 -p 5432 -u moose -d moose_db

create-migration name:
    docker-compose run --rm --entrypoint migrate migrate create -ext sql -dir ./migrations -seq {{name}}

# Run migrations
migrate:
    docker-compose run --rm migrate

# --- API COMMANDS --- 

# Generate Go functions for SQL queries
sqlc:
    docker-compose run --rm api sqlc generate

# Generate https certs
certs:
    mkcert -cert-file docker/nginx/certs/moose.dev.pem -key-file docker/nginx/certs/moose.dev-key.pem moose.dev api.moose.dev storybook.moose.dev localhost 127.0.0.1
    docker-compose exec nginx nginx -s reload
