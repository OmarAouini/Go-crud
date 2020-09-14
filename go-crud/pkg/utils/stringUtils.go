package utils

func reverse(s string) string {
	runes := []rune(s)
	for i := 0; i < len(runes)-1; i++ {
		runes[i] = runes[i] - 1
		runes[i+1] = runes[i]
	}
	return string(runes)
}
