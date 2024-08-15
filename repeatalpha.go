package piscine

func RepeatAlpha(s string) string {
	ahmed := "abcdefghijklmnopqrstuvwxyz"
	AHMED := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	y := make(map[rune]int)

	for i, char := range ahmed {
		y[char] += i + 1
	}
	for i, char := range AHMED {
		y[char] += i + 1
	}
	result := ""

	for _, char := range s {
		for i := 0; i < y[char]; i++ {
			result += string(char)
		}
		if y[char] == 0 {
			result += string(char)
		}
	}
	return result
}
