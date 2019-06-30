package handler

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/techforward/ECscript_server/db"
	"github.com/techforward/ECscript_server/models"
	"github.com/techforward/ECscript_server/util"
)

// GetAllAddresses is getting all addresses.
// @Summary get all addresses
// @Description get all addresses in a ECsite
// @Tags Address
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Address
// @Router /address [get]
func GetAllAddresses(c echo.Context) error {
	var addresses []models.Address

	db := db.ConnectGORM()
	db.SingularTable(true)
	db.Find(&addresses)
	defer db.Close()

	return c.JSON(http.StatusOK, addresses)
}

// GetAddress is getting address.
// @Summary get address
// @Description get address in a ECsite
// @Tags Address
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Address
// @Router /address/{id} [get]
func GetAddress(c echo.Context) error {
	addressUlid := c.Param("id")
	var address models.Address

	db := db.ConnectGORM()
	db.SingularTable(true)
	if err := db.Where("ulid = ?", addressUlid).First(&address).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}
	defer db.Close()

	return c.JSON(http.StatusOK, address)
}

// CreateAddress is creating address.
// @Summary create new address
// @Description create new address in a ECsite
// @Tags Address
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Address
// @Router /address [post]
func CreateAddress(c echo.Context) error {
	address := new(models.Address)

	if err := c.Bind(address); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}
	if address.Ulid == "" {
		address.Ulid = util.NewUlid()
	}
	if err := address.Validation(); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	db := db.ConnectGORM()
	db.SingularTable(true)
	if err := db.Create(&address).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}
	defer db.Close()

	return c.JSON(http.StatusOK, address)
}

// UpdateAddress is updating address.
// @Summary Update address
// @Description updating address in a ECsite
// @Tags Address
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Address
// @Router /address/{id} [put]
func UpdateAddress(c echo.Context) error {
	addressID := c.Param("id")
	address := new(models.Address)
	address.Ulid = addressID
	if err := c.Bind(address); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}
	if err := address.Validation(); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	db := db.ConnectGORM()
	db.SingularTable(true)
	if err := db.Save(&address).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}
	defer db.Close()

	return c.JSON(http.StatusOK, address)
}

// DeleteAddress is deleting address.
// @Summary Delete address
// @Description deleting address in a ECsite
// @Tags Address
// @Accept  json
// @Produce  json
// @Success 200 {string} Address Deleted successfully
// @Router /address/{id} [delete]
func DeleteAddress(c echo.Context) error {
	addressID := c.Param("id")
	address := new(models.Address)
	if address.Ulid == "" {
		address.Ulid = addressID
	}

	db := db.ConnectGORM()
	db.SingularTable(true)
	if err := db.Delete(&address).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}
	defer db.Close()

	return c.JSON(http.StatusOK, "Address Deleted successfully")
}
