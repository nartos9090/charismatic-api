package utils

import "strings"

func FilterStringSlice(s, list []string) []string {
	res := make([]string, 0)

	for _, val := range s {
		val = strings.ToUpper(val)

		for _, item := range list {
			item = strings.ToUpper(item)

			if val == item {
				res = append(res, val)
			}
		}
	}

	return res
}

func FilterValidString(s, fb string, list []string) string {
	s = strings.ToLower(s)

	for _, v := range list {
		if s == v {
			return v
		}
	}

	return fb
}
