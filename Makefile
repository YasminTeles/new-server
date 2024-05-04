.PHONY: help setup run test build docker-build docker-run docker-kill lint update-dependencies

VERSION := 'github.com/YasminTeles/new-server/handlers.GitTag=$(shell git describe --tags --always)'
COMMIT := 'github.com/YasminTeles/new-server/handlers.GitCommit=$(shell git rev-list --oneline -1 HEAD)'
BUILD := 'github.com/YasminTeles/new-server/handlers.BuildData=$(shell date +%F%t%T)'

VERSION_TEST := 'github.com/YasminTeles/new-server/handlers.GitTag=1.0.1'
COMMIT_TEST := 'github.com/YasminTeles/new-server/handlers.GitCommit=feat: add new route'
BUILD_TEST := 'github.com/YasminTeles/new-server/handlers.BuildData=2024-05-04 18:39:59'

help: ## Show help.
	@printf "A set of development commands.\n"
	@printf "\nUsage:\n"
	@printf "\t make \033[36m<commands>\033[0m\n"
	@printf "\nThe Commands are:\n\n"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\t\033[36m%-30s\033[0m %s\n", $$1, $$2}'

setup: ## Setup server.
	@go mod download

run: ## Run local server.
	@go run -ldflags "-X $(COMMIT) -X $(VERSION) -X $(BUILD)" main.go

test: ## Run test.
	@go test -ldflags "-X $(COMMIT_TEST) -X $(VERSION_TEST) -X $(BUILD_TEST)" -v ./... -cover -covermode=atomic

build: ## Build server.
	@go build -v -ldflags "-X $(COMMIT) -X $(VERSION) -X $(BUILD)" -o main .

docker-build: ## Build container's Docker.
	@docker build -t server .

docker-run: ## Run container's Docker.
	@docker run --name new-server -p 4000:4000 -it server

docker-kill: ## Kill container's Docker.
	@docker kill new-server

lint: ## Run lint.
	golangci-lint run ./...

update-dependencies: ## Update all dependencies.
	@go get -u
