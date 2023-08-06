package firebase

import (
	"context"
	"log"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
)

var (
	app    *firebase.App
	client *auth.Client
)

// ClientのCloseは必要ないっぽい
func InitClient() {
	var err error
	app, err = firebase.NewApp(context.Background(), nil)
	if err != nil {
		log.Fatalf("InitClient: Failed to initialize firebase app: %v\n", err)
	}
	client, err = app.Auth(context.Background())
	if err != nil {
		log.Fatalf("InitClient: Failed to initialize firebase client: %v\n", err)
	}
}

func GetClient() *auth.Client {
	if client == nil {
		InitClient()
		log.Println("GetClient: Firebase client should be initialized when server starts.")
	}
	return client
}
