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

lint:
    @golangci-lint run
fmt:
    @golangci-lint fmt

docker-build:
	@pack build maxtet1703/pr-assignment-service --builder paketobuildpacks/builder-jammy-full \
	--buildpack paketo-buildpacks/go

open-api *args:
    docker run --rm -v $(pwd):/code ghcr.io/swaggo/swag:latest {{args}}

update-spec:
    @just open-api init --parseInternal -g ./internal/app/app.go -o ./internal/openapi
