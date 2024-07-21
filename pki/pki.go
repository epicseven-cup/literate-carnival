package pki

import (
	"crypto/rand"
	"crypto/rsa"
	"literatecarnival/logger"
)

func GenKeyPair() (*rsa.PublicKey, *rsa.PrivateKey) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 256)
	if err != nil {
		logger.DefaultLogger.Panicln("Something unexpected happend when Generating Private Key", err)
	}
	err = privateKey.Validate()
	if err != nil {
		logger.DefaultLogger.Panicln("The generated private key is incorrect", err)
	}
	publicKey := privateKey.Public()
	return publicKey.(*rsa.PublicKey), privateKey
}
