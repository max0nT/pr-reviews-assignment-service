set dotenv-load

default:
    @just --list

copy-settings:
    cp .env.example .env

migrate-run *args:
    echo $POSTGRES_URI
    @migrate -path ./migrations -database $POSTGRES_URI {{args}}
