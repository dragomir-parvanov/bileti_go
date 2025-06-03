package parser

import (
	"io"
	"log"
	"os"
	"testing"
	"time"
)

const TestFileName = "parse_bileti_test.html"

func GetTestFile() io.Reader {
	file, err := os.Open(TestFileName)
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

func TestShouldReturnRegionalForestryDirectorate(t *testing.T) {
	testFile := GetTestFile()

	result := ParsePermittedForFellingHTML(testFile).regionalForestryDirectorate

	expected := "Бургас"

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

func TestShouldReturnAllowedForester(t *testing.T) {

	testFile := GetTestFile()

	result := ParsePermittedForFellingHTML(testFile).allowedForester

	expected := "Регистриран лесовъд"

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

	testFile := GetTestFile()

	result := ParsePermittedForFellingHTML(testFile).section

	expected := "450"

	if result != expected {
		t.Errorf("Got %s, expected %s", result, expected)
	}
}

func TestShouldReturnSubSection(t *testing.T) {

	testFile := GetTestFile()

	result := ParsePermittedForFellingHTML(testFile).subSection

	expected := "1"

	if result != expected {
		t.Errorf("Got %s, expected %s", result, expected)
	}
}

func TestShouldReturnCadastreId(t *testing.T) {

	testFile := GetTestFile()

	result := ParsePermittedForFellingHTML(testFile).cadastreId

	expected := "тестов кадастрален номер"

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

	expected := "Тестов Маркировач"

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

	expected := time.Time{}

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

	expected := "Тестово почистване"

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

	expected := "Издал човек"

	if result != expected {
		t.Errorf("Got %s, expected %s", result, expected)
	}
}

func TestShouldReturnWhoReceivedThePermit(t *testing.T) {
	testFile := GetTestFile()

	result := ParsePermittedForFellingHTML(testFile).whoReceivedThePermit

	expected := "Получил човек"

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

	expected := "Служител"

	if result != expected {
		t.Errorf("Got %s, expected %s", result, expected)
	}
}

func TestShouldReturnIssuedCode(t *testing.T) {
	testFile := GetTestFile()

	result := ParsePermittedForFellingHTML(testFile).issuedCode

	expected := "Някакъв код"

	if result != expected {
		t.Errorf("Got %s, expected %s", result, expected)
	}
}

func TestShouldReturnPermitIssueRegion(t *testing.T) {
	testFile := GetTestFile()

	result := ParsePermittedForFellingHTML(testFile).permitIssuePlace.region

	expected := "Бургас"

	if result != expected {
		t.Errorf("Got %s, expected %s", result, expected)
	}
}

func TestShouldReturnPermitIssueMunicipality(t *testing.T) {
	testFile := GetTestFile()

	result := ParsePermittedForFellingHTML(testFile).permitIssuePlace.municipality

	expected := "Айтос"

	if result != expected {
		t.Errorf("Got %s, expected %s", result, expected)
	}
}

func TestShouldReturnPermitIssueLand(t *testing.T) {
	testFile := GetTestFile()

	result := ParsePermittedForFellingHTML(testFile).permitIssuePlace.land

	expected := "Айтос"

	if result != expected {
		t.Errorf("Got %s, expected %s", result, expected)
	}
}

func TestShouldReturnPermitIssueAddress(t *testing.T) {
	testFile := GetTestFile()

	result := ParsePermittedForFellingHTML(testFile).permitIssuePlace.address

	expected := "Някакъв адрес"

	if result != expected {
		t.Errorf("Got %s, expected %s", result, expected)
	}
}

func TestShouldReturnPermitIssueSubsection(t *testing.T) {
	testFile := GetTestFile()

	result := ParsePermittedForFellingHTML(testFile).permitIssuePlace.subSection

	expected := "Някакъв подотдел"

	if result != expected {
		t.Errorf("Got %s, expected %s", result, expected)
	}
}

func TestShouldReturnPermitIssueGpsCoordinates(t *testing.T) {
	testFile := GetTestFile()

	result := ParsePermittedForFellingHTML(testFile).permitIssuePlace.gpsCoordinates

	expected := "Някакви кординати"

	if result != expected {
		t.Errorf("Got %s, expected %s", result, expected)
	}
}
