-include .env
export

include .env

# Server Configuration.
start:
	@GO111MODULE=on go run main.go

install: # use goose with github version 
	@go get github.com/pressly/goose/cmd/goose
	@go get ./...
	@GO111MODULE=on go mod tidy
	@GO111MODULE=on go mod vendor

new-migration:
	@echo generating new migration on database: ${DB_NAME}
	@echo format command should be make new-migrate name=init_user_table
	@goose -dir migrations mysql "${DB_USER}:${DB_PASS}@tcp(${DB_HOST})/${DB_NAME}" create $(name) sql

migrate: # Run the MySQL migration.
	@echo operating on database: ${DB_NAME}
	@goose -dir migrations mysql "${DB_USER}:${DB_PASS}@tcp(${DB_HOST})/${DB_NAME}" up

migrate-dbtest: # Run the MySQL migration.
	@echo operating on database: ${DB_TEST_NAME}
	@goose -dir migrations mysql "${DB_TEST_USER}:${DB_TEST_PASS}@tcp(${DB_TEST_HOST})/${DB_TEST_NAME}" up

rollback: # Rollback to previous migration.
	@echo operating on database: ${DB_NAME}
	@goose -dir migrations mysql "${DB_USER}:${DB_PASS}@tcp(${DB_HOST})/${DB_NAME}" down

test: # Run the test.
	@go test -v -cover -race -coverprofile=coverage.out -bench=. ./...

clean: # Clear the MySQL tmp/ directory.
	@rm -rf tmp

review:
	@go get -u github.com/kisielk/errcheck; ls -la; errcheck ./...
	@go get honnef.co/go/tools/cmd/staticcheck; staticcheck -checks all ./...
	@go vet ./...
	@go get github.com/securego/gosec/cmd/gosec; gosec ./...

