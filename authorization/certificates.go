package authorization

import (
	"crypto/rsa"
	"io/ioutil"

	"github.com/golang-jwt/jwt/v4"
)

var (
	singkey   *rsa.PrivateKey
	verifyKey *rsa.PublicKey
)

func loadFlies(privateFile, publicFile string) error {

	privateBytes, err := ioutil.ReadFile(privateFile)

	if err != nil {

		return err
	}
	publicBytes, err := ioutil.ReadFile(publicFile)

	if err != nil {

		return err
	}

	return parseRSA(privateBytes, publicBytes)

}

func parseRSA(privateBytes, publicBytes []byte) error {
	var err error
	singkey, err = jwt.ParseRSAPrivateKeyFromPEM(privateBytes)

	if err != nil {
		return err
	}

	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(publicBytes)
	if err != nil {
		return err
	}

	return nil

}
