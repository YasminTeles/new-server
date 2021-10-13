VERSION := $(shell git rev-list --oneline -1 HEAD)

setup:
	@go mod download

run:
	@go run -ldflags "-X 'github.com/YasminTeles/new-server/server.GitCommit=$(VERSION)'" main.go

test:
	@go test -v ./...

build:
	@go build -v -ldflags "-X 'github.com/YasminTeles/new-server/server.GitCommit=$(VERSION)'" -o main .

docker-build:
	@docker build -t server .

docker-run:
	@docker run --name new-server -p 3000:3000 -it server

docker-kill:
	@docker kill new-server
