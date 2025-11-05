run:
	go run ./cmd/bot

test:
	go test ./...

build:
	go build -o bin/trading-bot ./cmd/bot

docker:
	docker build -t trading-bot:latest .
