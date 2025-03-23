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
	@go build -o bin/srv ./cmd/

## test: run all test cases
.PHONY:	test
test:
	@go test -v ./...

.PHONY: copyenv
copyenv: 
	cp .env ./endpoints/.env

## run: run api server after building 
.PHONY:	run
run: build copyenv
	@./bin/srv

.PHONY: dropdb
dropdb:
	@sudo -u postgres psql -c "DROP DATABASE IF EXISTS urlc;"

## newdb: replace old db with a new one
.PHONY: newdb
newdb: dropdb
	@sudo -u postgres psql -c "CREATE DATABASE urlc;"
	@sudo -u postgres psql -d urlc -f ./sqlc/schema.sql
	@sudo -u postgres psql -d urlc -f ./sqlc/index.sql
	@sudo -u postgres psql -d urlc -f ./sqlc/triggers.sql

## testdb: creates and poluates db for testing
.PHONY: testdb
testdb: newdb
	@sudo -u postgres psql -d urlc -f ./sqlc/dummy/pages.sql
	@sudo -u postgres psql -d urlc -f ./sqlc/dummy/users.sql
	@sudo -u postgres psql -d urlc -f ./sqlc/dummy/parentComments.sql
	@sudo -u postgres psql -d urlc -f ./sqlc/dummy/childComments.sql