package util

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"os"
	"path"
)

func GenerateRSAKeyPairs(tokenType string) {
	workDir, err := os.Getwd()
	if err != nil {
		log.Println("Error retrieving workDir")
		return
	}

	privateKey, err := rsa.GenerateKey(rand.Reader, 2056)
	if err != nil {
		log.Println("Error generating RSA private key: ", err)
		return
	}

	privateKeyPEM := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}

	privateKeyFile, err := os.Create(path.Join(workDir, fmt.Sprintf("/configs/%s-private.pem", tokenType)))
	if err != nil {
		log.Println("Error creating key file")
		return
	}
	defer privateKeyFile.Close()

	err = pem.Encode(privateKeyFile, privateKeyPEM)
	if err != nil {
		log.Println("Error encoding private key to PEM", err)
		return
	}

	log.Printf("Private RSA key saved to %s-private.pem\n", tokenType)

	publicKey := privateKey.PublicKey

	publicKeyDER, err := x509.MarshalPKIXPublicKey(&publicKey)
	if err != nil {
		log.Println("Error marshaling public key:", err)
		return
	}

	publicKeyPEM := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyDER,
	}

	publicKeyFile, err := os.Create(path.Join(workDir, fmt.Sprintf("/configs/%s-public.pem", tokenType)))
	if err != nil {
		log.Println("Error creating public key file")
		return
	}
	defer publicKeyFile.Close()

	err = pem.Encode(publicKeyFile, publicKeyPEM)
	if err != nil {
		log.Println("Error encoding public key to PEM", err)
		return
	}
	log.Printf("Public key saved to %s-public.pem\n", tokenType)
}
