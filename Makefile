# Task
run: 
	go run server.go

build: 
	go build -trimpath -ldflags '-s -w' -o bin/go-server-poc

fmt:
	go fmt ./...

gen:
	go generate ./...

build-tools:
	go build -mod=mod -ldflags '-s -w' -o ./bin/oapi-codegen github.com/deepmap/oapi-codegen/cmd/oapi-codegen

compose-up:
	docker compose up -d

docker-build:
	docker build -t go-server-poc .

docker-run:
	docker run -it -p 8080:8080 go-server-poc
