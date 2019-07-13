package handler

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/techforward/ECscript_server/db"
	"github.com/techforward/ECscript_server/models"
	"github.com/techforward/ECscript_server/util"
)

// GetAllOthers godoc
// @Summary その他リスト
// @Description その他をリストで返す
// @Tags Other
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Other
// @Failure 404
// @Router /other [get]
func GetAllOthers(c echo.Context) error {
	var others []models.Other

	db := db.ConnectGORM()
	db.SingularTable(true)
	db.Find(&others)
	defer db.Close()

	return c.JSON(http.StatusOK, others)
}

// GetOther godoc
// @Summary その他
// @Description その他を返す
// @Tags Other
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Other
// @Failure 404
// @Router /other/{id} [get]
func GetOther(c echo.Context) error {
	otherUlid := c.Param("id")
	var other models.Other

	db := db.ConnectGORM()
	db.SingularTable(true)
	if err := db.Where("ulid = ?", otherUlid).First(&other).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}
	defer db.Close()

	return c.JSON(http.StatusOK, other)
}

// CreateOther is creating その他.
// @Summary create new その他
// @Description create new その他 in a ECsite
// @Tags Other
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Other
// @Router /other [post]
func CreateOther(c echo.Context) error {
	other := new(models.Other)

	if err := c.Bind(other); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}
	if other.Ulid == "" {
		other.Ulid = util.NewUlid()
	}

	// Source
	fileHeader, err := c.FormFile("image")
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	other.Path, err = util.UploadImage(fileHeader, fileHeader.Filename)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	db := db.ConnectGORM()
	db.SingularTable(true)
	if err := db.Create(&other).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}
	defer db.Close()

	return c.JSON(http.StatusOK, other)
}

// UpdateOther is updating その他.
// @Summary Update その他
// @Description updating その他 in a ECsite
// @Tags Other
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Other
// @Router /other/{id} [put]
func UpdateOther(c echo.Context) error {
	otherID := c.Param("id")
	other := new(models.Other)
	other.Ulid = otherID
	if err := c.Bind(other); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	// Source
	fileHeader, err := c.FormFile("image")
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}
	_, err = util.DeleteImage(other.Path)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}
	other.Path, err = util.UploadImage(fileHeader, fileHeader.Filename+other.Ulid)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	db := db.ConnectGORM()
	db.SingularTable(true)
	if err := db.Save(&other).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}
	defer db.Close()

	return c.JSON(http.StatusOK, other)
}

// DeleteOther is deleting その他.
// @Summary Delete その他
// @Description deleting その他 in a ECsite
// @Tags Other
// @Accept  json
// @Produce  json
// @Success 200 {string} Other Deleted successfully
// @Router /other/{id} [delete]
func DeleteOther(c echo.Context) error {
	otherID := c.Param("id")
	other := new(models.Other)
	if other.Ulid == "" {
		other.Ulid = otherID
	}

	db := db.ConnectGORM()
	db.SingularTable(true)
	if err := db.Delete(&other).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}
	defer db.Close()

	return c.JSON(http.StatusOK, "Other Deleted successfully")
}
