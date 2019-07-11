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
	ItemDeletedMessage = "\"Item Deleted successfully\"\n"

	createdItem models.Item
	gotItem     models.Item

	testItemID string
)

func TestItem(t *testing.T) {
	println("--------------------------------")
	println("------------- Item -------------")
	println("--------------------------------")
}

func TestCreateItem(t *testing.T) {
	itemUlid := testItemID
	// Setup
	e := echo.New()
	f := make(url.Values)
	f.Set("ulid", itemUlid)
	f.Set("name", "みかん")
	f.Set("context", "くだもの")
	f.Set("amount", "4214")
	f.Set("priority", "12")
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(f.Encode()))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/item")

	// Assertions
	if assert.NoError(t, CreateItem(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		result := rec.Body.String()
		json.Unmarshal([]byte(result), &createdItem)
		testItemID = createdItem.Ulid
		fmt.Printf("CreateItem: %+v", result)
	}
}

func TestGetAllItem(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/item")

	// Assertions
	if assert.NoError(t, GetAllItems(c)) {
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
		fmt.Printf("GetAllItems: %v...%v", result[:number], result[(len(result)-number):])
	}
}

func TestGetItem(t *testing.T) {
	itemUlid := testItemID
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/item/:id")
	c.SetParamNames("id")
	c.SetParamValues(itemUlid)

	// Assertions
	if assert.NoError(t, GetItem(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		result := rec.Body.String()
		json.Unmarshal([]byte(result), &gotItem)
		fmt.Printf("GetItem: %+v", gotItem)
		assert.Equal(t, createdItem.Ulid, gotItem.Ulid)
	}
}

func TestUpdateItem(t *testing.T) {
	itemUlid := testItemID

	// Setup
	e := echo.New()
	f := make(url.Values)
	f.Set("name", "みかん2")
	f.Set("context", "くだもの2")
	f.Set("amount", "2222")
	f.Set("priority", "22")
	req := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(f.Encode()))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/item/:id")
	c.SetParamNames("id")
	c.SetParamValues(itemUlid)

	// Assertions
	if assert.NoError(t, UpdateItem(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		result := rec.Body.String()
		fmt.Printf("UpdateItem: %+v", result)
	}
}

func TestDeleteItem(t *testing.T) {
	itemUlid := testItemID

	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/item/:id")
	c.SetParamNames("id")
	c.SetParamValues(itemUlid)

	// Assertions
	if assert.NoError(t, DeleteItem(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		result := rec.Body.String()
		fmt.Printf("DeleteItem: %v", result)
		assert.Equal(t, ItemDeletedMessage, result)
	}
}
