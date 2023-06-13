package db

import (
	"context"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/storage"
	"google.golang.org/api/option"
)

var (
	App     *firebase.App
	Storage *storage.Client
)

func InitializeFirebase() error {
	opt := option.WithCredentialsFile("keyfile.json")

	// Initialize Firebase Admin SDK.
	var err error
	App, err = firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return err
	}

	// Initialize Firebase Storage client.
	Storage, err = App.Storage(context.Background())
	if err != nil {
		return err
	}

	return nil
}
