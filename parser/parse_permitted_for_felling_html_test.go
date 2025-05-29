package parser

import (
	"io"
	"log"
	"os"
	"testing"
)

const TEST_FILE_NAME = "parse_bileti_test.html"

func GetTestFile() io.Reader {
	file, err := os.Open(TEST_FILE_NAME)
	if err != nil {
		log.Fatal(err)
	}

	return file
}

func TestShouldReturnId(t *testing.T) {

	test_file := GetTestFile()

	result := ParsePermittedForFellingHTML(test_file).id

	expected := "0805138"

	if result != expected {
		t.Errorf("Got %s, expected %s", result, expected)
	}
}

func TestShouldReturnDirectiveNumber(t *testing.T) {

	test_file := GetTestFile()

	result := ParsePermittedForFellingHTML(test_file).directiveNumber

	expected := "127/18.02.2021"

	if result != expected {
		t.Errorf("Got %s, expected %s", result, expected)
	}
}

func TestShouldReturnPermittedFor(t *testing.T) {

	test_file := GetTestFile()

	result := ParsePermittedForFellingHTML(test_file).permittedFor

	expected := "ЕЛЕКТРОЕНЕРГИЕН СИСТЕМЕН ОПЕРАТОР ЕАД"

	if result != expected {
		t.Errorf("Got %s, expected %s", result, expected)
	}
}

func TestShouldReturnTypeOfFelling(t *testing.T) {

	test_file := GetTestFile()

	result := ParsePermittedForFellingHTML(test_file).typeOfFelling

	expected := "В сервитути"

	if result != expected {
		t.Errorf("Got %s, expected %s", result, expected)
	}
}

func TestShouldReturnSection(t *testing.T) {

	test_file := GetTestFile()

	result := ParsePermittedForFellingHTML(test_file).section

	expected := "450"

	if result != expected {
		t.Errorf("Got %s, expected %s", result, expected)
	}
}

func TestShouldReturnSubSection(t *testing.T) {

	test_file := GetTestFile()

	result := ParsePermittedForFellingHTML(test_file).subSection

	expected := "1"

	if result != expected {
		t.Errorf("Got %s, expected %s", result, expected)
	}
}

func TestShouldReturnMunicipality(t *testing.T) {

	testFile := GetTestFile()

	result := ParsePermittedForFellingHTML(testFile).municipality

	expected := "Айтос"

	if result != expected {
		t.Errorf("Got %s, expected %s", result, expected)
	}
}

func TestShouldReturnLand(t *testing.T) {

	testFile := GetTestFile()

	result := ParsePermittedForFellingHTML(testFile).land

	expected := "Айтос землище"

	if result != expected {
		t.Errorf("Got %s, expected %s", result, expected)
	}
}
