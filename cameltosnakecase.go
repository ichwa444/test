package piscine

func CamelToSnakeCase(s string) string {
	if len(s) == 0 {
		return ""
	}

	var result string
	for i := 0; i < len(s); i++ {
		char := s[i]
		if char >= 'A' && char <= 'Z' {
			if i != 0 {
				result += "_"
			}
			result += string(char + 32)
		} else {
			result += string(char)
		}
	}
	return result
}
