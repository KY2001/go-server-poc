# Task
run: 
	go run server.go

build: 
	go build -ldflags '-s -w' -o bin/go-server-poc

fmt:
	go fmt ./...

gen:
	go generate ./...

build-tools:
	go build -mod=mod -ldflags '-s -w' -o ./bin/oapi-codegen github.com/deepmap/oapi-codegen/cmd/oapi-codegen

compose-up:
	docker compose up -d
