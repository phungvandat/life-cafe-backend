bin:
	go build -o bin/migrator ./cmd/migrator

dev: bin
	go build -o bin/migrator ./cmd/migrator
	bin/migrator up
	ENV=local go run cmd/server/main.go

test:
	go test ./...
