set dotenv-load

default:
    @just --list

copy-settings:
    cp .env.example .env

run:
    go run cmd/main.go

migrate-run *args:
    echo $POSTGRES_URI
    @migrate -path ./migrations -database $POSTGRES_URI {{args}}
