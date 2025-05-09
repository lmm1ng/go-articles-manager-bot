build:
	go build -o bin/bot cmd/go-articles-manager-bot/main.go
run-dev:
	go run cmd/go-articles-manager-bot/main.go
run-prod:
	./bin/bot
