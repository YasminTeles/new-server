VERSION := 'github.com/YasminTeles/new-server/server.GitTag=$(shell git describe --tags --always)'
COMMIT := 'github.com/YasminTeles/new-server/server.GitCommit=$(shell git rev-list --oneline -1 HEAD)'
BUILD := 'github.com/YasminTeles/new-server/server.BuildData=$(shell date +%F%t%T)'

setup:
	@go mod download

run:
	@go run -ldflags "-X $(COMMIT) -X $(VERSION) -X $(BUILD)" main.go

test:
	@go test -v ./...

build:
	@go build -v -ldflags "-X $(COMMIT) -X $(VERSION) -X $(BUILD)" -o main .

docker-build:
	@docker build -t server .

docker-run:
	@docker run --name new-server -p 3000:3000 -it server

docker-kill:
	@docker kill new-server

lint:
	golangci-lint run ./... --enable-all
