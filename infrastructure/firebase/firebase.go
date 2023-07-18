package firebase

import (
	"context"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
)

var (
	app    *firebase.App
	client *auth.Client
)

func InitClient() {
	var err error
	app, err = firebase.NewApp(context.Background(), nil)
	if err != nil {
		panic("Failed to initialize firebase app.")
	}
	client, err = app.Auth(context.Background())
	if err != nil {
		panic("Failed to initialize firebase client.")
	}
}

func GetClient() *auth.Client {
	return client
}
