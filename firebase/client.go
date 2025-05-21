package firebase

import (
	"context"
	"log"
	"os"
	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

var FirebaseApp *firebase.App

func InitFirebase() {
	keyPath := os.Getenv("FIREBASE_KEY_PATH")
	opt := option.WithCredentialsFile(keyPath)
	// opt := option.WithCredentialsFile("firebase/serviceAccountKey.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing firebase app: %v\n", err)
	}
	FirebaseApp = app
	log.Println("✅ Firebase initialized successfully")
}
