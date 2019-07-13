package handler

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"

	"github.com/techforward/ECscript_server/db"
	"github.com/techforward/ECscript_server/models"
	"github.com/techforward/ECscript_server/util"
)

// GetAllUsers godoc
// @Summary ユーザーリスト
// @Description ユーザーをリストで返す
// @Tags User
// @Accept  json
// @Produce  json
// @Success 200 {object} models.User
// @Failure 404
// @Router /user [get]
func GetAllUsers(c echo.Context) error {
	var users []models.User

	db := db.ConnectGORM()
	db.SingularTable(true)
	db.Find(&users)
	defer db.Close()

	return c.JSON(http.StatusOK, users)
}

// GetUser godoc
// @Summary ユーザー
// @Description ユーザーを返す
// @Tags User
// @Accept  json
// @Produce  json
// @Success 200 {object} models.User
// @Failure 404
// @Router /user/{id} [get]
func GetUser(c echo.Context) error {
	userUlid := c.Param("id")
	var user models.User

	db := db.ConnectGORM()
	db.SingularTable(true)
	if err := db.Where("ulid = ?", userUlid).First(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}
	defer db.Close()

	return c.JSON(http.StatusOK, user)
}

// CreateUser is creating ユーザー.
// @Summary create new ユーザー
// @Description create new ユーザー in a ECsite
// @Tags User
// @Accept  json
// @Produce  json
// @Success 200 {object} models.User
// @Router /user [post]
func CreateUser(c echo.Context) error {
	user := new(models.User)

	if err := c.Bind(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	if user.Ulid == "" {
		user.Ulid = util.NewUlid()
	}
	if err := user.Validation(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	db := db.ConnectGORM()
	db.SingularTable(true)
	if err := db.Create(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}
	defer db.Close()

	return c.JSON(http.StatusOK, user)
}

// UpdateUser is updating ユーザー.
// @Summary Update ユーザー
// @Description updating user in a ECsite
// @Tags User
// @Accept  json
// @Produce  json
// @Success 200 {object} models.User
// @Router /user/{id} [put]
func UpdateUser(c echo.Context) error {
	userClaim := c.Get("user").(*jwt.Token)
	claims := userClaim.Claims.(jwt.MapClaims)
	userUlid := claims["ulid"].(string)

	userID := c.Param("id")

	if userUlid != userID {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	user := new(models.User)
	user.Ulid = userID
	if err := c.Bind(user); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}
	if err := user.Validation(); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	db := db.ConnectGORM()
	db.SingularTable(true)
	if err := db.Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}
	defer db.Close()

	return c.JSON(http.StatusOK, user)
}

// DeleteUser is deleting ユーザー.
// @Summary Delete ユーザー
// @Description deleting ユーザー in a ECsite
// @Tags User
// @Accept  json
// @Produce  json
// @Success 200 {string} User Deleted successfully
// @Router /user/{id} [delete]
func DeleteUser(c echo.Context) error {
	userID := c.Param("id")
	user := new(models.User)
	if user.Ulid == "" {
		user.Ulid = userID
	}

	db := db.ConnectGORM()
	db.SingularTable(true)
	if err := db.Delete(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err)
	}
	defer db.Close()

	return c.JSON(http.StatusOK, "User Deleted successfully")
}
