set dotenv-load

default:
    @just --list

copy-settings:
    cp .env.example .env

run:
    go run cmd/main.go

test:
    go clean -testcache
    docker compose stop
    docker compose -f compose.test.yaml down -v
    docker compose -f compose.test.yaml up -d
    go test -v ./integration-tests

migrate-run *args:
    echo $DATABASE_URL
    @migrate -path ./migrations -database $DATABASE_URL {{args}}

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
    @just open-api init  \
    -g ./internal/app/app.go \
     -o ./internal/openapi \
     --outputTypes go,yaml
