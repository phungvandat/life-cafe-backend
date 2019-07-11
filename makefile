bin:
	go build -o bin/migrator ./cmd/migrator
dev: bin
	bin/migrator up
	ENV=local go run cmd/server/main.go