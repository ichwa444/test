package piscine

func FirstWord(s string) string {
	var result []byte

	for i := 0; i < len(s); i++ {
		if s[i] >= 'a' && s[i] <= 'z' || s[i] >= 'A' && s[i] <= 'Z' || s[i] >= '0' && s[i] <= '9' {
			result = append(result, byte(s[i]))
		}
		if s[i] == ' ' || s[i] == '\t' || s[i] == '\n' {
			break
		}
	}

	return string(result) + "\n"
}
