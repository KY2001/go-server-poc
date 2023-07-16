# デプロイ用コンテナに含まれるバイナリを作成するコンテナ
FROM golang:1.20 as deploy-builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -trimpath -ldflags '-s -w' -o app

# デプロイ用のコンテナ
FROM debian:bookworm-slim as deploy

RUN apt-get update

COPY gcp-key.json .
COPY --from=deploy-builder /app/app .

CMD ["./app"]
