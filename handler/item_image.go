package handler

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/techforward/ECscript_server/db"
	"github.com/techforward/ECscript_server/models"
	"github.com/techforward/ECscript_server/util"
)

// GetAllItemImages is getting all ItemImages.
// @Summary get all ItemImages
// @Description get all ItemImages in a ECsite
// @Tags ItemImage
// @Accept  json
// @Produce  json
// @Success 200 {array} models.ItemImage
// @Router /item/{id}/image [get]
func GetAllItemImages(c echo.Context) error {
	itemUlid := c.Param("id")
	var itemImages []models.ItemImage

	db := db.ConnectGORM()
	db.SingularTable(true)
	if err := db.Where("item_ulid = ?", itemUlid).Find(&itemImages).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}
	defer db.Close()

	return c.JSON(http.StatusOK, itemImages)
}

// GetItemImage is getting itemImage.
// @Summary get itemImage
// @Description get itemImage in a ECsite
// @Tags ItemImage
// @Accept  json
// @Produce  json
// @Success 200 {object} models.ItemImage
// @Router /item/{id}/image/{imageUlid} [get]
func GetItemImage(c echo.Context) error {
	// itemUlid := c.Param("id")
	itemImageUlid := c.Param("imageUlid")
	var itemImage models.ItemImage

	db := db.ConnectGORM()
	db.SingularTable(true)
	if err := db.Where("ulid = ?", itemImageUlid).First(&itemImage).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}
	defer db.Close()

	return c.JSON(http.StatusOK, itemImage)
}

// CreateItemImage is creating itemImage.
// @Summary create new itemImage
// @Description create new itemImage in a ECsite
// @Tags ItemImage
// @Accept  json
// @Produce  json
// @Success 200 {object} models.ItemImage
// @Router /item/{id}/image [post]
func CreateItemImage(c echo.Context) error {
	itemUlid := c.Param("id")
	itemImage := new(models.ItemImage)

	// Source
	fileHeader, err := c.FormFile("image")
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	itemImage.Path, err = util.UploadImage(fileHeader, fileHeader.Filename)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	if itemImage.Ulid == "" {
		itemImage.Ulid = util.NewUlid()
	}
	if itemImage.ItemUlid == "" {
		itemImage.ItemUlid = itemUlid
	}
	if err := itemImage.Validation(); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	db := db.ConnectGORM()
	db.SingularTable(true)
	if err := db.Create(&itemImage).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}
	defer db.Close()

	return c.JSON(http.StatusOK, itemImage)
}

// UpdateItemImage is updating itemImage.
// @Summary Update itemImage
// @Description updating itemImage in a ECsite
// @Tags ItemImage
// @Accept  json
// @Produce  json
// @Success 200 {object} models.ItemImage
// @Router /item/{id}/image/{imageUlid} [put]
func UpdateItemImage(c echo.Context) error {
	itemImageUlid := c.Param("imageUlid")
	itemImage := new(models.ItemImage)
	itemImage.Ulid = itemImageUlid

	if err := c.Bind(itemImage); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	// Source
	fileHeader, err := c.FormFile("image")
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	itemImage.Path, err = util.UploadImage(fileHeader, fileHeader.Filename)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	if err := itemImage.Validation(); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	db := db.ConnectGORM()
	db.SingularTable(true)
	if err := db.Save(&itemImage).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}
	defer db.Close()

	return c.JSON(http.StatusOK, itemImage)
}

// DeleteItemImage is deleting itemImage.
// @Summary Delete itemImage
// @Description deleting itemImage in a ECsite
// @Tags ItemImage
// @Accept  json
// @Produce  json
// @Success 200 {string} ItemImage Deleted successfully
// @Router /item/{id}/image/{imageUlid} [delete]
func DeleteItemImage(c echo.Context) error {
	itemImageUlid := c.Param("imageUlid")
	itemImage := new(models.ItemImage)
	if itemImage.Ulid == "" {
		itemImage.Ulid = itemImageUlid
	}

	db := db.ConnectGORM()
	db.SingularTable(true)
	if err := db.Delete(&itemImage).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}
	defer db.Close()

	return c.JSON(http.StatusOK, "ItemImage Deleted successfully")
}
