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
// @Summary 商品リスト
// @Description 商品をリストで返す
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
// @Summary 商品
// @Description 商品を返す
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

// CreateUser is creating user.
// @Summary create new user
// @Description create new user in a ECsite
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

// UpdateUser is updating user.
// @Summary Update user
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

// DeleteUser is deleting user.
// @Summary Delete user
// @Description deleting user in a ECsite
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
