package jwt

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	UUID_KEY = "JWT_UUID"
)

type Claims struct {
	Email string `json:"email"`
	UUID  string `json:"uuid"`

	jwt.StandardClaims
}

// GenerateJWT generate a JWT
func GenerateJWT(claims *Claims, jwtKey, issuer string, expireFrom time.Duration) (string, error) {
	currTime := time.Now()
	expireTime := currTime.Add(expireFrom)

	claims.IssuedAt = currTime.Unix()
	claims.ExpiresAt = expireTime.Unix()
	claims.Issuer = issuer

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString([]byte(jwtKey))

	return token, err
}

// ParseJWT  parse JWT
func ParseJWT(token, jwtKey string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtKey), nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
