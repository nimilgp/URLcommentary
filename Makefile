## help: print this help message
.PHONY:	help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

## format: format code into the go standard 
.PHONY: format
format: 
	@go fmt ./...

## build: build the api server executable
.PHONY:	build
build:
	@go build -o bin/srv cmd/main.go

## test: run all test cases
.PHONY:	test
test:
	@go test -v ./...

## run: run api server after building 
.PHONY:	run
run: build
	@./bin/srv