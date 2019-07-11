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
	OrderDeletedMessage = "\"Order Deleted successfully\"\n"

	createdOrder models.Order
	gotOrder     models.Order

	testOrderID string
)

func TestOrder(t *testing.T) {
	println("--------------------------------")
	println("------------- Order ------------")
	println("--------------------------------")
}

func TestCreateOrder(t *testing.T) {
	orderUlid := testOrderID
	// Setup
	e := echo.New()
	f := make(url.Values)
	f.Set("ulid", orderUlid)
	f.Set("userUlid", "01DDSX4GRR2PSRKCDVTRFA96WH")
	f.Set("addressUlid", "01DDTFT8B2EPJ2MKJK1BDWGQGA")
	f.Set("isCanceled", "0")
	f.Set("status", "0")
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(f.Encode()))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/order")

	// Assertions
	if assert.NoError(t, CreateOrder(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		result := rec.Body.String()
		json.Unmarshal([]byte(result), &createdOrder)
		fmt.Printf("CreateOrder: %+v", result)
	}
}

func TestGetAllOrders(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/order")

	// Assertions
	if assert.NoError(t, GetAllOrders(c)) {
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
		fmt.Printf("GetAllOrders: %v...%v", result[:number], result[(len(result)-number):])
	}
}

func TestGetOrder(t *testing.T) {
	orderUlid := testOrderID
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/order/:id")
	c.SetParamNames("id")
	c.SetParamValues(orderUlid)

	// Assertions
	if assert.NoError(t, GetOrder(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		result := rec.Body.String()
		json.Unmarshal([]byte(result), &gotOrder)
		fmt.Printf("GetOrder: %+v", gotOrder)
		assert.Equal(t, createdOrder.Ulid, gotOrder.Ulid)
	}
}

func TestUpdateOrder(t *testing.T) {
	orderUlid := testOrderID
	// Setup
	e := echo.New()
	f := make(url.Values)
	f.Set("userUlid", "01DDSX4GRR2PSRKCDVTRFA96WH")
	f.Set("addressUlid", "01DDTFT8B2EPJ2MKJK1BDWGQGA")
	f.Set("isCanceled", "1")
	f.Set("status", "2")
	req := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(f.Encode()))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/order/:id")
	c.SetParamNames("id")
	c.SetParamValues(orderUlid)

	// Assertions
	if assert.NoError(t, UpdateOrder(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		result := rec.Body.String()
		fmt.Printf("UpdateOrder: %+v", result)
	}
}

func TestDeleteOrder(t *testing.T) {
	orderUlid := testOrderID

	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/order/:id")
	c.SetParamNames("id")
	c.SetParamValues(orderUlid)

	// Assertions
	if assert.NoError(t, DeleteOrder(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		result := rec.Body.String()
		fmt.Printf("DeleteOrder: %v", result)
		assert.Equal(t, OrderDeletedMessage, result)
	}
}
