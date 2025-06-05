package utils_string

func IsPlaceholder(str string) bool {

	for _, char := range str {
		if char != '.' && char != '-' {
			return false
		}
	}

	return true
}
