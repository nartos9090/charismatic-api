package utils

func ConvertNumberToAlphabet(n int) string {
	result := ""
	for n > 0 {
		mod := (n - 1) % 26
		result = string(rune('A'+mod)) + result
		n = (n - mod) / 26
	}

	return result
}
