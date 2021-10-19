package appMiddleware

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func CreateToken(personId int, jwtSecret string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["personId"] = personId
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix() //Token expires after 1 hour
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtSecret))

}

func ExtractTokenUserId(c echo.Context) int {
	token := c.Get("user").(*jwt.Token)
	if token != nil && token.Valid {
		claims := token.Claims.(jwt.MapClaims)
		personId := claims["personId"]
		switch personId.(type) {
		case float64:
			return int(personId.(float64))
		default:
			return personId.(int)
		}
	}
	return -1 // invalid user
}
