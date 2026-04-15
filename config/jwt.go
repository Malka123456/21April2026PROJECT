package config

import (
	"crypto/rsa"
	"log"

	//"io/ioutil"
	"os"

	"github.com/golang-jwt/jwt"
	//"golang.org/x/crypto/bcrypt"
)


var PrivateKey *rsa.PrivateKey
var PublicKey *rsa.PublicKey

func LoadKeys() {

	privateBytes, err := os.ReadFile("key/private.pem") 
		if err != nil {
			log.Fatal("Failed to load private key:", err)
			}

	publicBytes, err := os.ReadFile("key/public.pem") 
		if err != nil {
			log.Fatal("Failed to load public key:", err)
			}

	PrivateKey, err = jwt.ParseRSAPrivateKeyFromPEM(privateBytes)
			if err != nil {
				log.Fatal("Failed to parse private key:", err)
			}

	PublicKey, err = jwt.ParseRSAPublicKeyFromPEM(publicBytes)
			if err != nil {
				log.Fatal("Failed to parse public key:", err)
			}
}