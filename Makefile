build:
	@go build -o bin/learn-go cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/learn-go