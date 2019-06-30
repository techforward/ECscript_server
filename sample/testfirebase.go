package main

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
)

func main() {
	VerifyIdToken(`eyJhbGciOiJSUzI1NiIsImtpZCI6Ijk4MGVkMGQ3ODY2ODk1Y2E0M2MyMGRhZmM4NTlmMThjNjcwMWU3OTYiLCJ0eXAiOiJKV1QifQ.eyJuYW1lIjoiS2FzdSBHMiIsInBpY3R1cmUiOiJodHRwczovL2xoNS5nb29nbGV1c2VyY29udGVudC5jb20vLVJUaUgyT2Nrbl9RL0FBQUFBQUFBQUFJL0FBQUFBQUFBQUtzL1JxYnJEbjBXY3NVL3Bob3RvLmpwZyIsImlzcyI6Imh0dHBzOi8vc2VjdXJldG9rZW4uZ29vZ2xlLmNvbS9lY3NpdGUtMjQyMTExIiwiYXVkIjoiZWNzaXRlLTI0MjExMSIsImF1dGhfdGltZSI6MTU2MTE4MzE0NCwidXNlcl9pZCI6InNEb05RSVdTdzJWcUZaVFlmNGpJVFlWb1JhczIiLCJzdWIiOiJzRG9OUUlXU3cyVnFGWlRZZjRqSVRZVm9SYXMyIiwiaWF0IjoxNTYxMTgzMTg0LCJleHAiOjE1NjExODY3ODQsImVtYWlsIjoiYWx0c2lyLmNnYy55aHlhdnJAZ21haWwuY29tIiwiZW1haWxfdmVyaWZpZWQiOnRydWUsImZpcmViYXNlIjp7ImlkZW50aXRpZXMiOnsiZ29vZ2xlLmNvbSI6WyIxMDA0ODYxMzgxODM1NjQxMzU1NDkiXSwiZW1haWwiOlsiYWx0c2lyLmNnYy55aHlhdnJAZ21haWwuY29tIl19LCJzaWduX2luX3Byb3ZpZGVyIjoiZ29vZ2xlLmNvbSJ9fQ.PmB5LYMLeDHoaUhclmE1USLJONXGeGRwaBeMwA2yhHF9J7fjqnOCW1F_WeQudY4Yk_e2FsjJ8VtCup_6DE1Gnk2YmmOSZVEWwJ-NUGkBT4brn_k1FhOY_96yNfhrDajZKcSn4fpi79zv8Ka_6910l32s7h9SOhq37EPSbGhaYoOcmaNDveUjtU0__DyYSNKwXt3aHQpjAbN4j4ONX_2QRD6Lloih2VFcf03LpeLmIJkf3HZeJ6nWLhIIRWQRk7mfkcO4cHUPLexlBtOjsTUCEfNENx1-uIKbMbUnVjrP6vzsi5YdkqqdAYul1Fei76v8lPVlMthZsjG2Yd7yKNxaOw`)
}

func initApp() *firebase.App {
	// export GOOGLE_APPLICATION_CREDENTIALS="/home/user/Downloads/ecsite-242111-firebase-adminsdk-cyv20-10513497c8.json"
	// $env:GOOGLE_APPLICATION_CREDENTIALS="C:\Users\username\Downloads\ecsite-242111-firebase-adminsdk-cyv20-10513497c8.json"
	app, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	return app
}

func VerifyIdToken(idToken string) {
	app := initApp()
	client, err := app.Auth(context.Background())
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}

	token, err := client.VerifyIDToken(context.Background(), idToken)
	if err != nil {
		log.Fatalf("error verifying ID token: %+v\n", err)
	}

	log.Printf("Verified ID token: %+v\n", token)
}
