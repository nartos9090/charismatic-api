package crypto

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"crypto/x509"
	"encoding/pem"
)

type (
	RSAKeyPair struct {
		Private string
		Public  string
	}
)

func GenerateRSAKey() RSAKeyPair {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}

	publicKey := &privateKey.PublicKey

	private := privateRSAToString(privateKey)
	public := publicRSAToString(publicKey)

	return RSAKeyPair{
		Private: private,
		Public:  public,
	}
}
func privateRSAToString(private *rsa.PrivateKey) string {
	var bytes []byte = x509.MarshalPKCS1PrivateKey(private)
	block := &pem.Block{
		Type:  `RSA PRIVATE KEY`,
		Bytes: bytes,
	}
	pem := pem.EncodeToMemory(block)
	key := string(pem)

	return key
}
func publicRSAToString(public *rsa.PublicKey) string {
	var bytes []byte = x509.MarshalPKCS1PublicKey(public)
	block := &pem.Block{
		Type:  `RSA PUBLIC KEY`,
		Bytes: bytes,
	}
	pem := pem.EncodeToMemory(block)
	key := string(pem)

	return key
}

func stringToPrivateRSA(private string) *rsa.PrivateKey {
	block, _ := pem.Decode([]byte(private))
	if block == nil {
		panic(`failed to parse PEM block containing private key`)
	}

	key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}

	return key
}
func stringToPublicRSA(public string) *rsa.PublicKey {
	block, _ := pem.Decode([]byte(public))
	if block == nil {
		panic(`failed to parse PEM block containing public key`)
	}

	key, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		panic(err)
	}

	return key
}

func EncryptWithPublicKey(id string, pub string) string {
	hash := sha512.New()
	publicKey := stringToPublicRSA(pub)
	msg := []byte(id)

	cipher, err := rsa.EncryptOAEP(hash, rand.Reader, publicKey, msg, nil)
	if err != nil {
		panic(err)
	}

	return string(cipher)
}
func DecryptWithPrivateKey(data string, pri string) string {
	hash := sha512.New()
	privateKey := stringToPrivateRSA(pri)
	cipher := []byte(data)

	plain, err := rsa.DecryptOAEP(hash, rand.Reader, privateKey, cipher, nil)
	if err != nil {
		panic(err)
	}
	id := string(plain)

	return id
}
