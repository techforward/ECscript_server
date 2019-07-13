package handler

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/techforward/ECscript_server/db"
	"github.com/techforward/ECscript_server/models"
	"github.com/techforward/ECscript_server/util"
)

// GetAllOrders godoc
// @Summary オーダーリスト
// @Description オーダーをリストで返す
// @Tags Order
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Order
// @Failure 404
// @Router /order [get]
func GetAllOrders(c echo.Context) error {
	var orders []models.Order

	db := db.ConnectGORM()
	db.SingularTable(true)
	db.Find(&orders)
	defer db.Close()

	return c.JSON(http.StatusOK, orders)
}

// GetOrder godoc
// @Summary オーダー
// @Description オーダーを返す
// @Tags Order
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Order
// @Failure 404
// @Router /order/{id} [get]
func GetOrder(c echo.Context) error {
	orderUlid := c.Param("id")
	var order models.Order

	db := db.ConnectGORM()
	db.SingularTable(true)
	if err := db.Where("ulid = ?", orderUlid).First(&order).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}
	defer db.Close()

	return c.JSON(http.StatusOK, order)
}

// CreateOrder is creating オーダー.
// @Summary create new オーダー
// @Description create new オーダー in a ECsite
// @Tags Order
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Order
// @Router /order [post]
func CreateOrder(c echo.Context) error {
	order := new(models.Order)

	if err := c.Bind(order); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}
	if order.Ulid == "" {
		order.Ulid = util.NewUlid()
	}
	if err := order.Validation(); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	db := db.ConnectGORM()
	db.SingularTable(true)
	if err := db.Create(&order).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}
	defer db.Close()

	return c.JSON(http.StatusOK, order)
}

// UpdateOrder is updating オーダー.
// @Summary Update オーダー
// @Description updating オーダー in a ECsite
// @Tags Order
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Order
// @Router /order/{id} [put]
func UpdateOrder(c echo.Context) error {
	orderID := c.Param("id")
	order := new(models.Order)
	order.Ulid = orderID
	if err := c.Bind(order); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}
	if err := order.Validation(); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	db := db.ConnectGORM()
	db.SingularTable(true)
	if err := db.Save(&order).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}
	defer db.Close()

	return c.JSON(http.StatusOK, order)
}

// DeleteOrder is deleting オーダー.
// @Summary Delete オーダー
// @Description deleting オーダー in a ECsite
// @Tags Order
// @Accept  json
// @Produce  json
// @Success 200 {string} Order Deleted successfully
// @Router /order/{id} [delete]
func DeleteOrder(c echo.Context) error {
	orderID := c.Param("id")
	order := new(models.Order)
	if order.Ulid == "" {
		order.Ulid = orderID
	}

	db := db.ConnectGORM()
	db.SingularTable(true)
	if err := db.Delete(&order).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}
	defer db.Close()

	return c.JSON(http.StatusOK, "Order Deleted successfully")
}
