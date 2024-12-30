DB_URL=postgresql://postgres:postgres@localhost:5555/postgres?sslmode=disable
OPENAPI_GENERATOR := java -jar ./openapi-generator-cli.jar
CONFIG_FILE := ./config.yaml

generate-models:
	rm -rf resources/*
	$(OPENAPI_GENERATOR) generate -i docs/api.yaml -g go -o ./docs/web --additional-properties=packageName=resources
	mkdir -p resources
	find docs/web -name '*.go' -exec mv {} resources/ \;
	find resources -type f -name "*_test.go" -delete

generate-sqlc:
	sqlc generate

migrate-up:
	KV_VIPER_FILE=$(CONFIG_FILE) go build -o main main.go
	KV_VIPER_FILE=$(CONFIG_FILE) ./main migrate up

migrate-down:
	migrate -path internal/db/migration -database $(DB_URL) -verbose down

run-server:
	KV_VIPER_FILE=$(CONFIG_FILE) go build -o main main.go
	KV_VIPER_FILE=$(CONFIG_FILE) ./main run service