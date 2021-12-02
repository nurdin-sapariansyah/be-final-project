package middlewares

import (
	"time"

	"arisan.com/arisan/constants"
	"github.com/golang-jwt/jwt"
)

type JwtCustomClaims struct {
	AnggotaId int `json:"anggotaId"`
	jwt.StandardClaims
}

func GenereteTokenJWT(anggotaId int) (string, error) {

	// Set custom claims
	claims := &JwtCustomClaims{
		anggotaId,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(constants.SECRET_JWT))

	return t, err
}
