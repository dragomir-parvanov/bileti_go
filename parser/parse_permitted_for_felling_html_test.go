package parser

import (
	"io"
	"log"
	"os"
	"testing"
	"time"
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

	testFile := GetTestFile()

	result := ParsePermittedForFellingHTML(testFile).id

	expected := "0805138"

	if result != expected {
		t.Errorf("Got %s, expected %s", result, expected)
	}
}

func TestShouldReturnDirectiveNumber(t *testing.T) {

	testFile := GetTestFile()

	result := ParsePermittedForFellingHTML(testFile).directiveNumber

	expected := "127/18.02.2021"

	if result != expected {
		t.Errorf("Got %s, expected %s", result, expected)
	}
}

func TestShouldReturnPermittedFor(t *testing.T) {

	testFile := GetTestFile()

	result := ParsePermittedForFellingHTML(testFile).permittedFor

	expected := "ЕЛЕКТРОЕНЕРГИЕН СИСТЕМЕН ОПЕРАТОР ЕАД"

	if result != expected {
		t.Errorf("Got %s, expected %s", result, expected)
	}
}

func TestShouldReturnTypeOfFelling(t *testing.T) {

	testFile := GetTestFile()

	result := ParsePermittedForFellingHTML(testFile).typeOfFelling

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

func TestShouldReturnAreaClearing(t *testing.T) {
	testFile := GetTestFile()

	result := ParsePermittedForFellingHTML(testFile).areaClearing

	expected := 2.2

	if result != expected {
		t.Errorf("Got %f, expected %f", result, expected)
	}
}

func TestShouldReturnOwnershipType(t *testing.T) {
	testFile := GetTestFile()

	result := ParsePermittedForFellingHTML(testFile).ownershipType

	expected := "ДГТ"

	if result != expected {
		t.Errorf("Got %s, expected %s", result, expected)
	}
}

func TestShouldReturnTreesMarkedBy(t *testing.T) {
	testFile := GetTestFile()

	result := ParsePermittedForFellingHTML(testFile).treesMarkedBy

	expected := ""

	if result != expected {
		t.Errorf("Got %s, expected %s", result, expected)
	}
}

func TestShouldReturnControlMarkNumber(t *testing.T) {
	testFile := GetTestFile()

	result := ParsePermittedForFellingHTML(testFile).controlMarkNumber

	expected := "Б 1401"

	if result != expected {
		t.Errorf("Got %s, expected %s", result, expected)
	}
}

func TestShouldReturnMarkColor(t *testing.T) {
	testFile := GetTestFile()

	result := ParsePermittedForFellingHTML(testFile).controlMarkColor

	expected := "оранжева"

	if result != expected {
		t.Errorf("Got %s, expected %s", result, expected)
	}
}

func TestShouldReturnDateOfCarnetInventory(t *testing.T) {
	testFile := GetTestFile()

	result := ParsePermittedForFellingHTML(testFile).dateOfCarnetInventory

	expected := ""

	if result != expected {
		t.Errorf("Got %s, expected %s", result, expected)
	}
}

func TestShouldReturnExpectedTreeExtraction(t *testing.T) {
	testFile := GetTestFile()

	result := ParsePermittedForFellingHTML(testFile).expectedTreeExtraction

	expected := 1.0

	if result != expected {
		t.Errorf("Got %f, expected %f", result, expected)
	}
}

func TestShouldReturnAdditionalRequirements(t *testing.T) {
	testFile := GetTestFile()

	result := ParsePermittedForFellingHTML(testFile).additionalRequirements

	expected := "Позволителното за сеч се издава за почистване в сервитут на електропровод без материален добив."

	if result != expected {
		t.Errorf("Got %s, expected %s", result, expected)
	}

}

func TestShouldReturnDeadlineLoggingFrom(t *testing.T) {
	testFile := GetTestFile()

	result := ParsePermittedForFellingHTML(testFile).deadlineLogging.From

	expected, err := time.Parse("02.01.2006", "27.01.2025")

	if err != nil {
		t.Errorf("Got %s, expected error", err)
	}

	if result != expected {
		t.Errorf("Got %s, expected %s", result, expected)
	}
}

func TestShouldReturnDeadlineLoggingTo(t *testing.T) {
	testFile := GetTestFile()

	result := ParsePermittedForFellingHTML(testFile).deadlineLogging.To

	expected, err := time.Parse("02.01.2006", "31.12.2025")

	if err != nil {
		t.Errorf("Got %s, expected error", err)
	}

	if result != expected {
		t.Errorf("Got %s, expected %s", result, expected)
	}
}

func TestShouldReturnDeadlineMaterialsUsageFrom(t *testing.T) {
	testFile := GetTestFile()

	result := ParsePermittedForFellingHTML(testFile).deadlineMaterialsUsage.From

	expected, err := time.Parse("02.01.2006", "27.01.2025")

	if err != nil {
		t.Errorf("Got %s, expected error", err)
	}

	if result != expected {
		t.Errorf("Got %s, expected %s", result, expected)
	}
}

func TestShouldReturnDeadlineMaterialsUsageTo(t *testing.T) {
	testFile := GetTestFile()

	result := ParsePermittedForFellingHTML(testFile).deadlineMaterialsUsage.To

	expected, err := time.Parse("02.01.2006", "31.12.2025")

	if err != nil {
		t.Errorf("Got %s, expected error", err)
	}

	if result != expected {
		t.Errorf("Got %s, expected %s", result, expected)
	}
}

func TestShouldReturnCleaningProcedure(t *testing.T) {
	testFile := GetTestFile()

	result := ParsePermittedForFellingHTML(testFile).cleaningProcedure

	expected := ""

	if result != expected {
		t.Errorf("Got %s, expected %s", result, expected)
	}
}

func TestShouldReturnRemovalToTemporaryStorage(t *testing.T) {
	testFile := GetTestFile()

	result := ParsePermittedForFellingHTML(testFile).removalToTemporaryStorage

	expected := "съгласно ТП"

	if result != expected {
		t.Errorf("Got %s, expected %s", result, expected)
	}
}

func TestShouldReturnIssuedBy(t *testing.T) {
	testFile := GetTestFile()

	result := ParsePermittedForFellingHTML(testFile).issuedBy

	expected := ""

	if result != expected {
		t.Errorf("Got %s, expected %s", result, expected)
	}
}

func TestShouldReturnWhoReceivedThePermit(t *testing.T) {
	testFile := GetTestFile()

	result := ParsePermittedForFellingHTML(testFile).whoReceivedThePermit

	expected := ""

	if result != expected {
		t.Errorf("Got %s, expected %s", result, expected)
	}
}

func TestShouldReturnIssuedOn(t *testing.T) {
	testFile := GetTestFile()

	result := ParsePermittedForFellingHTML(testFile).issuedOn

	expected, err := time.Parse("02.01.2006", "24.01.2025")

	if err != nil {
		t.Errorf("Got %s, expected error", err)
	}

	if result != expected {
		t.Errorf("Got %s, expected %s", result, expected)
	}
}

func TestShouldReturnIssuedByEmployee(t *testing.T) {
	testFile := GetTestFile()

	result := ParsePermittedForFellingHTML(testFile).issuedByEmployee

	expected := ""

	if result != expected {
		t.Errorf("Got %s, expected %s", result, expected)
	}
}

func TestShouldReturnIssuedCode(t *testing.T) {
	testFile := GetTestFile()

	result := ParsePermittedForFellingHTML(testFile).issuedCode

	expected := ""

	if result != expected {
		t.Errorf("Got %s, expected %s", result, expected)
	}
}
