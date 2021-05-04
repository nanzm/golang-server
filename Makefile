.PHONY: mod lint test
# custom define
PROJECT := dora
MAINFILE := main.go

mod: ## Get the dependencies
	@go mod download

lint: ## Lint Golang files
	@golangci-lint --version
	@golangci-lint run -D errcheck

test: ## Run tests with coverage
	go test ./... -v

coverage-html: ## show coverage by the html
	go tool cover -html=.coverprofile

generate:
	go generate ./...