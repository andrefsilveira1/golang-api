build:
	@go build -o bin/golang-api

run: build 
	@./bin/golang-api

test:
	@go test -v ./...