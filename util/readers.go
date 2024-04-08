package util

import (
	"crypto/hmac"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"path"

	"github.com/dgrijalva/jwt-go"
)

func ReadRSAFromPEMFile(tokenType string, mode string)[]byte{
    workDir, err := os.Getwd()
    if err != nil{
        log.Fatal(err)
    }

    pemKey, err := os.ReadFile(path.Join(workDir, fmt.Sprintf("/configs/%s-%s.pem", tokenType, mode)))
    if err != nil{
        log.Fatal(err)
    }

    return pemKey
}

func PrivateRSAFromFile(tokenType string) *rsa.PrivateKey {
    privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(ReadRSAFromPEMFile(tokenType, "private"))
    if err != nil{
        log.Fatal(err)
    }
    return privateKey
}

func PublicRSAFromFile(tokenType string) *rsa.PublicKey {
    publicKey, err := jwt.ParseRSAPublicKeyFromPEM(ReadRSAFromPEMFile(tokenType, "public"))
    if err != nil{
        log.Fatal(err)
    }

    return publicKey
}

func ReadRSAPairFromPEMFile(tokenType string) (*rsa.PrivateKey, *rsa.PublicKey){
    return PrivateRSAFromFile(tokenType), PublicRSAFromFile(tokenType)
}

func  GenerateCustomKey(userId string, tokenHash string) string {
    h := hmac.New(sha256.New, []byte(tokenHash))
    h.Write([]byte(userId))
    sha := hex.EncodeToString(h.Sum(nil))
    return sha
}
