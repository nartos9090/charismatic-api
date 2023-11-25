package crypto

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	math_rand "math/rand"
	"time"
)

const saltSize = 16

type HashedPwd struct {
	Salt string
	Pwd  string
	Raw  string
}

func generateRandomSalt(saltSize int) string {
	saltBytes := make([]byte, saltSize)

	_, err := rand.Read(saltBytes[:])

	if err != nil {
		panic(err)
	}

	return base64.StdEncoding.EncodeToString(saltBytes)[:saltSize]
}

func hashPassword(pwd string, salt string) string {
	sha512Hasher := hmac.New(sha512.New, []byte(salt))
	sha512Hasher.Write([]byte(pwd))

	var hashedPwdBytes = sha512Hasher.Sum(nil)

	var base64EncodedPwdHash = base64.StdEncoding.EncodeToString(hashedPwdBytes)

	return base64EncodedPwdHash
}

func Match(currPwd, hashedPwd, salt string) bool {
	currPwdHash := hashPassword(currPwd, salt)

	return hashedPwd == currPwdHash
}

func Hash(pwd string) *HashedPwd {
	var hashedPwd HashedPwd

	hashedPwd.Salt = generateRandomSalt(saltSize)

	hashedPwd.Pwd = hashPassword(pwd, hashedPwd.Salt)

	return &hashedPwd
}

func TempPassword() *HashedPwd {
	s := math_rand.NewSource(time.Now().UTC().UnixNano())
	r := math_rand.New(s)
	random := r.Intn(10_000)

	pwd := fmt.Sprintf(`%04d`, random)

	var hashedPwd HashedPwd

	hashedPwd.Raw = pwd

	hashedPwd.Salt = generateRandomSalt(saltSize)

	hashedPwd.Pwd = hashPassword(pwd, hashedPwd.Salt)

	return &hashedPwd
}
