package authorization

import (
	"time"

	"github.com/SebGomez0416/Practica_API/model"
	"github.com/golang-jwt/jwt/v4"
)

func GenerateToken(data *model.Login) (string, error) {
	claim := model.Claim{
		Email: data.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour).Unix(),
			Issuer:    "Zion",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claim)
	signedToken, err := token.SignedString(signkey)

	if err != nil {

		return "", err
	}

	return signedToken, nil

}
