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
	CartDeletedMessage = "\"Cart Deleted successfully\"\n"

	createdCart models.Cart
	gotCart     models.Cart

	testCartID   string
	testBoughtAt string
)

func TestCart(t *testing.T) {
	println("--------------------------------")
	println("------------- Cart -------------")
	println("--------------------------------")
}

func TestCreateCart(t *testing.T) {
	cartUlid := testCartID
	// Setup
	e := echo.New()
	f := make(url.Values)
	f.Set("ulid", cartUlid)
	f.Set("userUlid", "01DDSX4GRR2PSRKCDVTRFA96WH")
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(f.Encode()))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/cart")

	// Assertions
	if assert.NoError(t, CreateCart(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		result := rec.Body.String()
		json.Unmarshal([]byte(result), &createdCart)
		fmt.Printf("CreateCart: %+v", result)
	}
}

func TestGetAllCarts(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/cart")

	// Assertions
	if assert.NoError(t, GetAllCarts(c)) {
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
		fmt.Printf("GetAllCarts: %v...%v", result[:number], result[(len(result)-number):])
	}
}

func TestGetCart(t *testing.T) {
	cartUlid := testCartID
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/cart/:id")
	c.SetParamNames("id")
	c.SetParamValues(cartUlid)

	// Assertions
	if assert.NoError(t, GetCart(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		result := rec.Body.String()
		json.Unmarshal([]byte(result), &gotCart)
		fmt.Printf("GetCart: %+v", gotCart)
		assert.Equal(t, createdCart.Ulid, gotCart.Ulid)
	}
}

func TestUpdateCart(t *testing.T) {
	cartUlid := testCartID
	BoughtAt := testBoughtAt

	// Setup
	e := echo.New()
	f := make(url.Values)
	f.Set("userUlid", "01DDSX4GRR2PSRKCDVTRFA96WH")
	f.Set("boughtAt", BoughtAt)
	req := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(f.Encode()))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/cart/:id")
	c.SetParamNames("id")
	c.SetParamValues(cartUlid)

	// Assertions
	if assert.NoError(t, UpdateCart(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		result := rec.Body.String()
		fmt.Printf("UpdateCart: %+v", result)
	}
}

func TestDeleteCart(t *testing.T) {
	cartUlid := testCartID

	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/cart/:id")
	c.SetParamNames("id")
	c.SetParamValues(cartUlid)

	// Assertions
	if assert.NoError(t, DeleteCart(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		result := rec.Body.String()
		fmt.Printf("DeleteCart: %v", result)
		assert.Equal(t, CartDeletedMessage, result)
	}
}
