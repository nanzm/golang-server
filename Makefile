.PHONY: up down mod lint tag build
# custom define
PROJECT := dora
MAINFILE := main.go

up: ## 
	docker-compose -f ./docker-compose.dev.yml up -d

down: ## 
	docker-compose -f ./docker-compose.dev.yml down

build:
	docker build . -t nancode/dora-server

tag:
	docker tag nancode/dora-server registry.cn-hangzhou.aliyuncs.com/nancode/dora-server:latest

push:
	docker push registry.cn-hangzhou.aliyuncs.com/nancode/dora-server:latest

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