package crypto

import (
	"strconv"
	"testing"
)

func TestGenerateRSAKeypair(t *testing.T) {
	keypair := GenerateRSAKey()

	t.Logf(`private : %+v`, keypair.Private)
	t.Logf(`public : %+v`, keypair.Public)
}

func TestEncryptAndDecrypt(t *testing.T) {
	keypair := GenerateRSAKey()
	id := 17

	stringID := strconv.Itoa(id)

	encryptedID := EncryptWithPublicKey(stringID, keypair.Public)
	t.Logf(`encrypted id: %+v`, encryptedID)

	decryptedID := DecryptWithPrivateKey(encryptedID, keypair.Private)
	t.Logf(`decrypted id: %+v`, decryptedID)
}
