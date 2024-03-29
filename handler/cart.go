package handler

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/techforward/ECscript_server/db"
	"github.com/techforward/ECscript_server/models"
	"github.com/techforward/ECscript_server/util"
)

// GetAllCarts is getting all carts.
// @Summary get all carts
// @Description get all carts in a ECsite
// @Tags Cart
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Cart
// @Router /cart [get]
func GetAllCarts(c echo.Context) error {
	var carts []models.Cart

	db := db.ConnectGORM()
	db.SingularTable(true)
	db.Find(&carts)
	defer db.Close()

	return c.JSON(http.StatusOK, carts)
}

// GetCart is getting cart.
// @Summary get cart
// @Description get cart in a ECsite
// @Tags Cart
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Cart
// @Router /cart/{id} [get]
func GetCart(c echo.Context) error {
	cartUlid := c.Param("id")
	var cart models.Cart

	db := db.ConnectGORM()
	db.SingularTable(true)
	if err := db.Where("ulid = ?", cartUlid).First(&cart).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}
	defer db.Close()

	return c.JSON(http.StatusOK, cart)
}

// CreateCart is creating cart.
// @Summary create new cart
// @Description create new cart in a ECsite
// @Tags Cart
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Cart
// @Router /cart [post]
func CreateCart(c echo.Context) error {
	cart := new(models.Cart)

	if err := c.Bind(cart); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}
	if cart.Ulid == "" {
		cart.Ulid = util.NewUlid()
	}
	if err := cart.Validation(); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	db := db.ConnectGORM()
	db.SingularTable(true)
	if err := db.Create(&cart).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}
	defer db.Close()

	return c.JSON(http.StatusOK, cart)
}

// UpdateCart is updating cart.
// @Summary Update cart
// @Description updating cart in a ECsite
// @Tags Cart
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Cart
// @Router /cart/{id} [put]
func UpdateCart(c echo.Context) error {
	cartID := c.Param("id")
	cart := new(models.Cart)
	cart.Ulid = cartID
	if err := c.Bind(cart); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}
	if err := cart.Validation(); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	db := db.ConnectGORM()
	db.SingularTable(true)
	if err := db.Save(&cart).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}
	defer db.Close()

	return c.JSON(http.StatusOK, cart)
}

// DeleteCart is deleting cart.
// @Summary Delete cart
// @Description deleting cart in a ECsite
// @Tags Cart
// @Accept  json
// @Produce  json
// @Success 200 {string} Cart Deleted successfully
// @Router /cart/{id} [delete]
func DeleteCart(c echo.Context) error {
	cartID := c.Param("id")
	cart := new(models.Cart)
	if cart.Ulid == "" {
		cart.Ulid = cartID
	}

	db := db.ConnectGORM()
	db.SingularTable(true)
	if err := db.Delete(&cart).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}
	defer db.Close()

	return c.JSON(http.StatusOK, "Cart Deleted successfully")
}
