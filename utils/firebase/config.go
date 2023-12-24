package utils_firebase

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

type firebaseConfig struct {
	TYPE                        string `json:"type"`
	PROJECT_ID                  string `json:"project_id"`
	PRIVATE_KEY_ID              string `json:"private_key_id"`
	PRIVATE_KEY                 string `json:"private_key"`
	CLIENT_EMAIL                string `json:"client_email"`
	CLIENT_ID                   string `json:"client_id"`
	AUTH_URI                    string `json:"auth_uri"`
	TOKEN_URI                   string `json:"token_uri"`
	AUTH_PROVIDER_X509_CERT_URL string `json:"auth_provider_x509_cert_url"`
	CLIENT_X509_CERT_URL        string `json:"client_x509_cert_url"`
}

func NewApp(ctx context.Context) *firebase.App {
	cj, err := getConfigJSON()
	if err != nil {
		log.Fatalln(err)
		return nil
	}

	sa := option.WithCredentialsJSON(cj)
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	return app
}

func getConfigJSON() ([]byte, error) {
	var (
		FIREBASE_TYPE                        = os.Getenv("FIREBASE_TYPE")
		FIREBASE_PROJECT_ID                  = os.Getenv("FIREBASE_PROJECT_ID")
		FIREBASE_PRIVATE_KEY_ID              = os.Getenv("FIREBASE_PRIVATE_KEY_ID")
		FIREBASE_PRIVATE_KEY                 = fmt.Sprintf("-----BEGIN PRIVATE KEY-----\n%s\n-----END PRIVATE KEY-----\n", os.Getenv("FIREBASE_PRIVATE_KEY"))
		FIREBASE_CLIENT_EMAIL                = os.Getenv("FIREBASE_CLIENT_EMAIL")
		FIREBASE_CLIENT_ID                   = os.Getenv("FIREBASE_CLIENT_ID")
		FIREBASE_AUTH_URI                    = os.Getenv("FIREBASE_AUTH_URI")
		FIREBASE_TOKEN_URI                   = os.Getenv("FIREBASE_TOKEN_URI")
		FIREBASE_AUTH_PROVIDER_X509_CERT_URL = os.Getenv("FIREBASE_AUTH_PROVIDER_X509_CERT_URL")
		FIREBASE_CLIENT_X509_CERT_URL        = os.Getenv("FIREBASE_CLIENT_X509_CERT_URL")
	)
	conf := firebaseConfig{
		TYPE:                        FIREBASE_TYPE,
		PROJECT_ID:                  FIREBASE_PROJECT_ID,
		PRIVATE_KEY_ID:              FIREBASE_PRIVATE_KEY_ID,
		PRIVATE_KEY:                 strings.ReplaceAll(FIREBASE_PRIVATE_KEY, `\n`, "\n"),
		CLIENT_EMAIL:                FIREBASE_CLIENT_EMAIL,
		CLIENT_ID:                   FIREBASE_CLIENT_ID,
		TOKEN_URI:                   FIREBASE_TOKEN_URI,
		AUTH_URI:                    FIREBASE_AUTH_URI,
		AUTH_PROVIDER_X509_CERT_URL: FIREBASE_AUTH_PROVIDER_X509_CERT_URL,
		CLIENT_X509_CERT_URL:        FIREBASE_CLIENT_X509_CERT_URL,
	}
	json, err := json.Marshal(conf)
	if err != nil {
		return json, errors.New("failed to marshal firebase config to JSON")
	}

	return json, nil
}

func FirestoreClient(ctx context.Context) *firestore.Client {
	app := NewApp(ctx)
	if app == nil {
		log.Fatalln("could not connect to firebase")
		return nil
	}

	cl, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln("could not connect to firestore::", err.Error())
		return nil
	}

	return cl
}
