up-db:
	@docker-compose up -d db

run: build
	@./backend/bin/go-app

build:
	@go build -o backend/bin/go-app backend/cmd/main.go
