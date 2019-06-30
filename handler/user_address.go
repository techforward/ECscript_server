package handler

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"

	"github.com/techforward/ECscript_server/db"
	"github.com/techforward/ECscript_server/models"
	"github.com/techforward/ECscript_server/util"
)

// GetAllUserAddresses is getting all user_addresses.
// @Summary get all user_addresses
// @Description get all user_addresses in a ECsite
// @Tags UserAddress
// @Accept  json
// @Produce  json
// @Success 200 {array} models.UserAddress
// @Router /user_address [get]
func GetAllUserAddresses(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userUlid := claims["ulid"].(string)

	var userUserAddresses []models.UserAddress

	db := db.ConnectGORM()
	db.SingularTable(true)
	db.Where("user_ulid = ?", userUlid).Find(&userUserAddresses)
	defer db.Close()

	return c.JSON(http.StatusOK, userUserAddresses)
}

// GetUserAddress is getting user_address.
// @Summary get user_address
// @Description get user_address in a ECsite
// @Tags UserAddress
// @Accept  json
// @Produce  json
// @Success 200 {object} models.UserAddress
// @Router /user_address/{id} [get]
func GetUserAddress(c echo.Context) error {
	userUserAddressUlid := c.Param("id")
	var userUserAddress models.UserAddress

	db := db.ConnectGORM()
	db.SingularTable(true)
	if err := db.Where("ulid = ?", userUserAddressUlid).First(&userUserAddress).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}
	defer db.Close()

	return c.JSON(http.StatusOK, userUserAddress)
}

// CreateUserAddress is creating user_address.
// @Summary create new user_address
// @Description create new user_address in a ECsite
// @Tags UserAddress
// @Accept  json
// @Produce  json
// @Success 200 {object} models.UserAddress
// @Router /user_address [post]
func CreateUserAddress(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userUlid := claims["ulid"].(string)

	userUserAddress := new(models.UserAddress)

	if err := c.Bind(userUserAddress); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}
	if userUserAddress.Ulid == "" {
		userUserAddress.Ulid = util.NewUlid()
	}
	if userUserAddress.UserUlid == "" {
		userUserAddress.UserUlid = userUlid
	}
	if err := userUserAddress.Validation(); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	db := db.ConnectGORM()
	db.SingularTable(true)
	if err := db.Create(&userUserAddress).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}
	defer db.Close()

	return c.JSON(http.StatusOK, userUserAddress)
}

// UpdateUserAddress is updating user_address.
// @Summary Update user_address
// @Description updating user_address in a ECsite
// @Tags UserAddress
// @Accept  json
// @Produce  json
// @Success 200 {object} models.UserAddress
// @Router /user_address/{id} [put]
func UpdateUserAddress(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userUlid := claims["ulid"].(string)

	userUserAddressID := c.Param("id")
	userUserAddress := new(models.UserAddress)
	userUserAddress.Ulid = userUserAddressID
	userUserAddress.UserUlid = userUlid

	if err := c.Bind(userUserAddress); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}
	if err := userUserAddress.Validation(); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	db := db.ConnectGORM()
	db.SingularTable(true)
	if err := db.Save(&userUserAddress).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}
	defer db.Close()

	return c.JSON(http.StatusOK, userUserAddress)
}

// DeleteUserAddress is deleting user_address.
// @Summary Delete user_address
// @Description deleting user_address in a ECsite
// @Tags UserAddress
// @Accept  json
// @Produce  json
// @Success 200 {string} UserAddress Deleted successfully
// @Router /user_address/{id} [delete]
func DeleteUserAddress(c echo.Context) error {
	userUserAddressID := c.Param("id")
	userUserAddress := new(models.UserAddress)
	if userUserAddress.Ulid == "" {
		userUserAddress.Ulid = userUserAddressID
	}

	db := db.ConnectGORM()
	db.SingularTable(true)
	if err := db.Delete(&userUserAddress).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}
	defer db.Close()

	return c.JSON(http.StatusOK, "UserAddress Deleted successfully")
}
