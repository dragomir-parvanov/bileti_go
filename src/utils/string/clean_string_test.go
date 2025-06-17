package utils_string

import "testing"

func TestShouldReturnEmptyStringIfItContainsOnlyDots(t *testing.T) {
	str := "......."

	actual := CleanString(str)

	expected := ""

	if actual != expected {
		t.Errorf("Expected: %s, Actual: %s", expected, actual)
	}
}
