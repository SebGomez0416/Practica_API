package authorization

import (
	"errors"
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

func ValidateToken(t string) (model.Claim, error) {

	token, err := jwt.ParseWithClaims(t, &model.Claim{}, verifyfunction)

	if err != nil {
		return model.Claim{}, err
	}

	if !token.Valid {
		return model.Claim{}, errors.New("token invalido")
	}

	claim, ok := token.Claims.(*model.Claim)

	if !ok {

		return model.Claim{}, errors.New("no se pudo obtener los claim")
	}

	return *claim, nil

}

func verifyfunction(t *jwt.Token) (interface{}, error) {

	return verifyKey, nil
}
