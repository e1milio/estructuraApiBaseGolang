package auth

import (
	"crypto/rsa"
	"io/ioutil"
	"log"
	"sync"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var privateKey *rsa.PrivateKey
var publicKey *rsa.PublicKey
var oncePrivateKey sync.Once // Lo utilizamos para construir un singleton de privateKey
var oncePublicKey sync.Once // Lo utilizamos para construir un singleton de publicKey

// GetPrivateKey retorna la privateKey
func GetPrivateKey() *rsa.PrivateKey {

	oncePrivateKey.Do(func() { // Función nativa que hace que el código se ejecute una sola vez

		privateBytes, err := ioutil.ReadFile("private.rsa")
		if err != nil {
			log.Println("auth.go - No se pudo leer el archivo private.rsa")
		}

		privateKey, err = jwt.ParseRSAPrivateKeyFromPEM(privateBytes)
		if err != nil {
			log.Println("auth.go - No se pudo hacer el parse a privateKey")
		}

	})

	return privateKey
}

// GetPublicKey retorna la publicKey
func GetPublicKey() *rsa.PublicKey {

	oncePublicKey.Do(func() { // Función nativa que hace que el código se ejecute una sola vez

		publicBytes, err := ioutil.ReadFile("public.rsa.pub")
		if err != nil {
			log.Println("auth.go - No se pudo leer el archivo public.rsa.pub")
		}

		publicKey, err = jwt.ParseRSAPublicKeyFromPEM(publicBytes)
		if err != nil {
			log.Println("auth.go - No se pudo hacer el parse a publicKey")
		}

	})

	return publicKey
}

// CreateToken  crea y retorna un token string
func CreateToken(username string) (string, error) {

	claims := &JWTCustomClaims{
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	tokenString, err := token.SignedString(GetPrivateKey())

	return tokenString, err

}

// GetToken retorna un token struct JWTCustomClaims
func GetToken(auth_token interface{}) *JWTCustomClaims {

	token := auth_token.(*jwt.Token).Claims.(*JWTCustomClaims)

	return token
}