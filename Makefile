# Task
run: 
	go run server.go

build: 
	go build -trimpath -ldflags '-s -w' -o bin/go-server-poc

fmt:
	go fmt ./...

lint:
	./bin/golangci-lint run

lint-fix:
	./bin/golangci-lint run --fix

gen:
	go generate ./...

build-tools:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b ./bin/ v1.53.3
	go build -mod=mod -ldflags '-s -w' -o ./bin/oapi-codegen github.com/deepmap/oapi-codegen/cmd/oapi-codegen
	go mod tidy

compose-up:
	docker compose up -d

docker-build:
	docker build -t go-server-poc .

docker-run:
	docker run -it -p 8080:8080 go-server-poc
