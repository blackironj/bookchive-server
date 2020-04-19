package jwt

import (
	"errors"
	"net/http"

	"github.com/blackironj/bookchive-server/env"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// CheckToken is jwt middleware
func CheckToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error
		token := c.GetHeader("Authorization")
		if token == "" {
			err = errors.New("Cannot find a token")
		} else {
			parsedToken, parseErr := ParseJWT(token, env.Conf.Auth.JWTKey)

			if parseErr != nil {
				switch parseErr.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					err = errors.New("Expired token")
				default:
					err = errors.New("Invalid token")
				}
			} else {
				c.Set(UUID_KEY, parsedToken.UUID)
			}
		}
		if err != nil {
			c.JSON(http.StatusUnauthorized, err.Error())

			c.Abort()
			return
		}
		c.Next()
	}
}
