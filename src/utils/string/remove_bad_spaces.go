package utils_string

import "strings"

func RemoveBadSpaces(str string) string {

	res := strings.Join(strings.Fields(strings.TrimSpace(str)), " ")

	return res
}
