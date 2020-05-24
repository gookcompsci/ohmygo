package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	const bits = 2048
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		log.Fatal(err)
	}
	privateKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	})
	if err := ioutil.WriteFile("private-key.pem", privateKeyPEM, 0644); err != nil {
		log.Fatal(err)
	}
	publicKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(&privateKey.PublicKey),
	})
	if err := ioutil.WriteFile("public-key.pem", publicKeyPEM, 0644); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", publicKeyPEM)
}
