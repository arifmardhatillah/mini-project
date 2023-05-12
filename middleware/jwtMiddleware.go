package middleware

import (
	"net/http"
	"project_structure/constant"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func CreateToken(userId int, role string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userId"] = userId
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	byteSecret := []byte(constant.SECRET_JWT)
	return token.SignedString(byteSecret)
}

func ExtractTokenUserId(e echo.Context) (int, error) {
	user := e.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	if user.Valid {
		userId := int(claims["userId"].(float64))
		return userId, nil
	}

	if claims["role"] != "user" {
		return 0, echo.NewHTTPError(http.StatusUnauthorized, "error requesting user")
	}

	return 0, nil
}

func ExtractTokenAdminId(e echo.Context) (int, error) {
	user := e.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	if user.Valid {
		userId := int(claims["userId"].(float64))
		return userId, nil
	}

	if claims["role"] != "admin" {
		return 0, echo.NewHTTPError(http.StatusUnauthorized, "error requesting admin")
	}
	return 0, nil
}
