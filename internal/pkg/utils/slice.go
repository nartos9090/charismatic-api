package utils

import (
	"fmt"
	"strconv"
	"strings"
)

func SliceToString[T any](slice []T) string {
	return strings.Trim(strings.Join(strings.Split(fmt.Sprint(slice), ` `), `,`), `[]`)
}

func StringToIntSlice(str string) []int {
	if len(str) == 0 {
		return make([]int, 0)
	}

	raw := strings.Split(str, `,`)
	res := make([]int, len(raw))

	for i, v := range raw {
		val, err := strconv.Atoi(v)
		if err != nil {
			res[i] = -1
		}
		res[i] = val
	}

	return res
}

func SumIntSlice(slice []int) (res int) {
	for _, v := range slice {
		res += v
	}

	return
}

func StringToStringSlice(str string) (ret []string) {
	raw := strings.Split(str, `,`)

	for _, v := range raw {
		if len(v) > 0 {
			ret = append(ret, v)
		}
	}

	return
}

func IntToCharStr(i int) string {
	return string(rune('A' - 1 + i))
}

func Prepend[T any](slice []T, val T) []T {
	slice = append(slice, val)

	copy(slice[1:], slice)

	slice[0] = val

	return slice
}
