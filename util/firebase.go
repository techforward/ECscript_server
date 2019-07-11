package util

import (
	"context"

	"github.com/techforward/ECscript_server/db"
	"github.com/techforward/ECscript_server/models"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/labstack/echo"
)

// IsUserExist IsUserExist
func IsUserExist(firebaseUUID string) (bool, models.User) {
	var user models.User

	db := db.ConnectGORM()
	db.SingularTable(true)
	// Get first matched record
	// SELECT * FROM users WHERE name = 'jinzhu' limit 1;
	if result := db.Where("firebase_uid = ?", firebaseUUID).First(&user); result.Error == nil {
		return true, user
	}

	return false, user
}

// VerifyUIDToken Firebaseのトークン検証
func VerifyUIDToken(idToken string) (isVerified bool, token *auth.Token, err error) {
	// $env:GOOGLE_APPLICATION_CREDENTIALS="C:\Users\username\Downloads\ecsite-242111-firebase-adminsdk-cyv20-10513497c8.json"
	// export GOOGLE_APPLICATION_CREDENTIALS="/home/user/Downloads/ecsite-242111-firebase-adminsdk-cyv20-10513497c8.json"
	err = nil
	token = nil

	app, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		return false, token, err
	}

	client, err := app.Auth(context.Background())
	if err != nil {
		return false, token, err
	}

	token, err = client.VerifyIDToken(context.Background(), idToken)
	if err != nil {
		return false, token, err
	}
	return true, token, err
}

// UserTokenMiddleware UserTokenMiddleware
func UserTokenMiddleware(idToken string, c echo.Context) (bool, error) {
	status, _, _ := VerifyUIDToken(idToken)
	return status, nil
}

// AdminTokenMiddleware AdminTokenMiddleware
func AdminTokenMiddleware(idToken string, c echo.Context) (bool, error) {
	status, _, _ := VerifyUIDToken(idToken)
	return status, nil
}
