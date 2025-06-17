package utils_string

import "testing"

func TestOnNoDuplicatedSpacesShouldReturnStrIntact(t *testing.T) {

	str := "No duplicated spaces"

	actual := RemoveBadSpaces(str)

	expected := str

	if actual != expected {
		t.Errorf("Expected: %s, Actual: %s", expected, actual)
	}
}

func TestOnSpacesAtTheStartAndEndShouldTrimThem(t *testing.T) {

	str := " Trim Test "

	actual := RemoveBadSpaces(str)

	expected := "Trim Test"

	if actual != expected {
		t.Errorf("Expected: %s, Actual: %s", expected, actual)
	}
}

func TestOnDuplicatedSpacesShouldReturnRemovedThem(t *testing.T) {

	str := "Duplicated   Spaces"

	actual := RemoveBadSpaces(str)

	expected := "Duplicated Spaces"

	if actual != expected {
		t.Errorf("Expected: %s, Actual: %s", expected, actual)
	}
}

func TestOnDuplicatedSpacesOnTheEndShouldRemoveThem(t *testing.T) {

	str := "Duplicated Spaces  "

	actual := RemoveBadSpaces(str)

	expected := "Duplicated Spaces"

	if actual != expected {
		t.Errorf("Expected: %s, Actual: %s", expected, actual)
	}
}
