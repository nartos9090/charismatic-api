package crypto

import (
	"testing"
)

var pwd = "admin"
var salt = "QG6k7XsdTD92293O"
var hashed = "dNdYiWzw85Pz62uCUw3fCei5s1mj920je8J966aiYUXCmJd9ceGG9Ba7BwcHeixYl7vMNVQliun090c+py3XkQ=="

func TestHash(t *testing.T) {
	t.Logf("hashing : %s", pwd)

	hashed := Hash(pwd)
	t.Logf("salt : %s", hashed.Salt)
	t.Logf("hashed : %s", hashed.Pwd)
}

func TestMatch(t *testing.T) {
	t.Logf("%s", hashed)
	if Match(pwd, hashed, salt) == false {
		t.Errorf("ERROR unmathched password")
	}
}
