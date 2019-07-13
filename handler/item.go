package handler

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/techforward/ECscript_server/db"
	"github.com/techforward/ECscript_server/models"
	"github.com/techforward/ECscript_server/util"
)

// GetAllItems godoc
// @Summary Itemリスト
// @Description Itemをリストで返す
// @Tags Item
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Item
// @Failure 404
// @Router /item [get]
func GetAllItems(c echo.Context) error {
	var items []models.Item

	db := db.ConnectGORM()
	db.SingularTable(true)
	db.Find(&items)
	defer db.Close()

	return c.JSON(http.StatusOK, items)
}

// GetItem godoc
// @Summary 商品
// @Description 商品を返す
// @Tags Item
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Item
// @Failure 404
// @Router /item/{id} [get]
func GetItem(c echo.Context) error {
	itemUlid := c.Param("id")
	var item models.Item

	db := db.ConnectGORM()
	db.SingularTable(true)
	if err := db.Where("ulid = ?", itemUlid).First(&item).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}
	defer db.Close()

	return c.JSON(http.StatusOK, item)
}

// CreateItem is creating item.
// @Summary create new item
// @Description create new item in a ECsite
// @Tags Item
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Item
// @Router /item [post]
func CreateItem(c echo.Context) error {
	item := new(models.Item)

	if err := c.Bind(item); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}
	if item.Ulid == "" {
		item.Ulid = util.NewUlid()
	}
	if err := item.Validation(); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	db := db.ConnectGORM()
	db.SingularTable(true)
	if err := db.Create(&item).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}
	defer db.Close()

	return c.JSON(http.StatusOK, item)
}

// UpdateItem is updating item.
// @Summary Update item
// @Description updating item in a ECsite
// @Tags Item
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Item
// @Router /item/{id} [put]
func UpdateItem(c echo.Context) error {
	itemID := c.Param("id")
	item := new(models.Item)
	item.Ulid = itemID
	if err := c.Bind(item); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}
	if err := item.Validation(); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	db := db.ConnectGORM()
	db.SingularTable(true)
	if err := db.Save(&item).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}
	defer db.Close()

	return c.JSON(http.StatusOK, item)
}

// DeleteItem is deleting item.
// @Summary Delete item
// @Description deleting item in a ECsite
// @Tags Item
// @Accept  json
// @Produce  json
// @Success 200 {string} Item Deleted successfully
// @Router /item/{id} [delete]
func DeleteItem(c echo.Context) error {
	itemID := c.Param("id")
	item := new(models.Item)
	if item.Ulid == "" {
		item.Ulid = itemID
	}

	db := db.ConnectGORM()
	db.SingularTable(true)
	if err := db.Delete(&item).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}
	defer db.Close()

	return c.JSON(http.StatusOK, "Item Deleted successfully")
}
