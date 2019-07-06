package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/techforward/ECscript_server/models"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

var (
	UserDeletedMessage string = "\"User Deleted successfully\"\n"

	createdUser models.User
	gotUser     models.User

	testUserID      string
	testFirebaseUID string
)

func TestUser(t *testing.T) {
	println("--------------------------------")
	println("-------------- User ------------")
	println("--------------------------------")
}

func TestCreateUser(t *testing.T) {
	userUlid := testUserID
	firebaseUID := testFirebaseUID

	// Setup
	e := echo.New()
	f := make(url.Values)
	f.Set("ulid", userUlid)
	f.Set("firebaseUid", firebaseUID)
	f.Set("name", "ユーザさん")
	f.Set("email", "test@me.com")
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(f.Encode()))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/user")

	// Assertions
	if assert.NoError(t, CreateUser(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		result := rec.Body.String()
		json.Unmarshal([]byte(result), &createdUser)
		fmt.Printf("CreateUser: %+v", result)
	}
}

func TestGetAllUsers(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/user")

	// Assertions
	if assert.NoError(t, GetAllUsers(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		result := rec.Body.String()
		var number, resultLen int
		resultLen = len(result)
		switch {
		case resultLen > 50:
			number = 50
		case resultLen > 25:
			number = 25
		default:
			number = 0
		}
		fmt.Printf("GetAllUsers: %v...%v", result[:number], result[(len(result)-number):])
	}
}

func TestGetUser(t *testing.T) {
	userUlid := testUserID
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/user/:id")
	c.SetParamNames("id")
	c.SetParamValues(userUlid)

	// Assertions
	if assert.NoError(t, GetUser(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		result := rec.Body.String()
		json.Unmarshal([]byte(result), &gotUser)
		fmt.Printf("GetUser: %+v", gotUser)
		assert.Equal(t, createdUser.Ulid, gotUser.Ulid)
	}
}

func TestUpdateUser(t *testing.T) {
	userUlid := testUserID
	// Setup
	e := echo.New()
	f := make(url.Values)
	f.Set("firebaseUid", "01DDS-X4GRR2-PSRKCDV-TRFA96WH")
	f.Set("cartUlid", "01DDTKER1V853DJ7YMS4S59KD7")
	f.Set("addressUlid", "01DDTFT8B2EPJ2MKJK1BDWGQGA")
	f.Set("name", "ユーザさん2")
	f.Set("email", "test@me.com2")
	req := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(f.Encode()))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/user/:id")
	c.SetParamNames("id")
	c.SetParamValues(userUlid)

	// Assertions
	if assert.NoError(t, UpdateUser(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		result := rec.Body.String()
		fmt.Printf("UpdateUser: %+v", result)
	}
}

func TestDeleteUser(t *testing.T) {
	userUlid := testUserID

	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/user/:id")
	c.SetParamNames("id")
	c.SetParamValues(userUlid)

	// Assertions
	if assert.NoError(t, DeleteUser(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		result := rec.Body.String()
		fmt.Printf("DeleteUser: %v", result)
		assert.Equal(t, UserDeletedMessage, result)
	}
}
