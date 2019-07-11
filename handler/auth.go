package handler

import (
	"net/http"
	"time"

	"github.com/techforward/ECscript_server/models"
	"github.com/techforward/ECscript_server/util"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func setClaim(user models.User, firebaseToken string) *jwt.Token {
	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["ulid"] = user.Ulid
	claims["firebaseUid"] = user.FirebaseUID
	claims["stripeId"] = user.StripeID
	claims["cartUlid"] = user.CartUlid
	claims["addressUlid"] = user.AddressUlid
	claims["name"] = user.Name
	claims["email"] = user.Email

	claims["token"] = firebaseToken
	claims["admin"] = false
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	return token
}

// VeryfyIDToken is veryfing firebase token.
// @Summary Veryfy firebase token
// @Description veryfy firebase token in a ECsite
// @Tags Auth
// @Accept  json
// @Produce  json
// @Success 200 {string} jwt token
// @Router /auth/{id} [get]
func VeryfyIDToken(c echo.Context) error {
	firebaseToken := c.Param("token")

	// Throws unauthorized error
	isVerified, tokenStruct, _ := util.VerifyUIDToken(firebaseToken)
	if !isVerified {
		return echo.ErrUnauthorized
	}

	isExist, user := util.IsUserExist(tokenStruct.UID)
	if !isExist {
		return echo.ErrUnauthorized
	}

	// Create token
	token := setClaim(user, firebaseToken)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}
