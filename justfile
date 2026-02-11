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
    $env:PGPASSWORD='moose'; pgcli -h localhost -p 5432 -u moose -d moose_db
