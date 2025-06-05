package utils_string

func CleanString(str string) string {
	noBadSpaces := RemoveBadSpaces(str)

	isPlaceholder := IsPlaceholder(noBadSpaces)

	if isPlaceholder {
		return ""
	}

	return noBadSpaces
}
