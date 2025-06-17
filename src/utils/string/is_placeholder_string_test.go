package utils_string

import "testing"

func TestOnOtherStringShouldReturnFalse(t *testing.T) {
	param := "Some other string"

	actual := IsPlaceholder(param)

	expected := false

	if actual != expected {
		t.Errorf("IsPlaceholder returned %t but expected %t", actual, expected)
	}
}

func TestOnOnlyDotsShouldReturnTrue(t *testing.T) {
	param := "....."

	actual := IsPlaceholder(param)

	expected := true

	if actual != expected {
		t.Errorf("IsPlaceholder returned %t but expected %t", actual, expected)
	}
}

func TestOnOnlyHyphensShouldReturnTrue(t *testing.T) {
	param := "-----"

	actual := IsPlaceholder(param)

	expected := true

	if actual != expected {
		t.Errorf("IsPlaceholder returned %t but expected %t", actual, expected)
	}
}
