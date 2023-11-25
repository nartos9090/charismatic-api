package utils

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func GenerateTemporaryPin(n int) string {
	s := rand.NewSource(time.Now().UTC().UnixNano())
	r := rand.New(s)

	base := 1
	for i := 1; i < n; i++ {
		base *= 10
	}

	g := r.Intn(10*base-1) + base
	res := fmt.Sprintf(`%0`+strconv.Itoa(n)+`d`, g)
	return res
}
