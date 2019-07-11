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
	AddressDeletedMessage = "\"Address Deleted successfully\"\n"

	createdAddress models.Address
	gotAddress     models.Address

	testAddressID string
)

func TestAddress(t *testing.T) {
	println("--------------------------------")
	println("----------- Address ------------")
	println("--------------------------------")
}

func TestCreateAddress(t *testing.T) {
	addressUlid := testAddressID
	// Setup
	e := echo.New()
	f := make(url.Values)
	f.Set("ulid", addressUlid)
	f.Set("postcode", "165-0022")
	f.Set("address1", "Kita-ku, Tokyo, Jp")
	f.Set("address2", "23-3-12")
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(f.Encode()))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/address")

	// Assertions
	if assert.NoError(t, CreateAddress(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		result := rec.Body.String()
		json.Unmarshal([]byte(result), &createdAddress)
		fmt.Printf("CreateAddress: %+v", result)
	}
}

func TestGetAllAddresses(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/address")

	// Assertions
	if assert.NoError(t, GetAllAddresses(c)) {
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
		fmt.Printf("GetAllAddresses: %v...%v", result[:number], result[(len(result)-number):])
	}
}

func TestGetAddress(t *testing.T) {
	addressUlid := testAddressID
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/address/:id")
	c.SetParamNames("id")
	c.SetParamValues(addressUlid)

	// Assertions
	if assert.NoError(t, GetAddress(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		result := rec.Body.String()
		json.Unmarshal([]byte(result), &gotAddress)
		fmt.Printf("GetAddress: %+v", gotAddress)
		assert.Equal(t, createdAddress.Ulid, gotAddress.Ulid)
	}
}

func TestUpdateAddress(t *testing.T) {
	addressUlid := testAddressID

	// Setup
	e := echo.New()
	f := make(url.Values)
	f.Set("postcode", "165-0022")
	f.Set("address1", "Kita-ku, Tokyo, Jp")
	f.Set("address2", "23-3-12")
	req := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(f.Encode()))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/address/:id")
	c.SetParamNames("id")
	c.SetParamValues(addressUlid)

	// Assertions
	if assert.NoError(t, UpdateAddress(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		result := rec.Body.String()
		fmt.Printf("UpdateAddress: %+v", result)
	}
}

func TestDeleteAddress(t *testing.T) {
	addressUlid := testAddressID

	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/address/:id")
	c.SetParamNames("id")
	c.SetParamValues(addressUlid)

	// Assertions
	if assert.NoError(t, DeleteAddress(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		result := rec.Body.String()
		fmt.Printf("DeleteAddress: %v", result)
		assert.Equal(t, AddressDeletedMessage, result)
	}
}
