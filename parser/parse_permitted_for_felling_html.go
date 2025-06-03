package parser

import (
	utilsstring "bileti_go/utils/string"
	utilstime "bileti_go/utils/time"
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const DateLayout = "02.01.2006"

func ParsePermittedForFellingHTML(htmlResult io.Reader) ParsedResult {

	doc, err := goquery.NewDocumentFromReader(htmlResult)

	if err != nil {
		log.Fatal(err)
	}

	return ParsedResult{
		id:                          _ExtractId(doc),
		regionalForestryDirectorate: _ExtractRegionalForestryDirectorate(doc),
		directiveNumber:             _ExtractDirectiveNumber(doc),
		permittedFor:                _ExtractPermittedFor(doc),
		allowedForester:             _ExtractAllowedForester(doc),
		typeOfFelling:               _ExtractTypeOfFelling(doc),
		section:                     _ExtractSection(doc),
		subSection:                  _ExtractSubSection(doc),
		cadastreId:                  _ExtractCadastreId(doc),
		municipality:                _ExtractMunicipality(doc),
		land:                        _ExtractLand(doc),
		areaClearing:                _ExtractAreaClearing(doc),
		ownershipType:               _ExtractOwnershipType(doc),
		treesMarkedBy:               _ExtractTreesMarkedBy(doc),
		controlMarkNumber:           _ExtractControlMarkNumber(doc),
		controlMarkColor:            _ExtractMarkColor(doc),
		dateOfCarnetInventory:       _ExtractDateOfCarnetInventory(doc),
		expectedTreeExtraction:      _ExtractExpectedTreeExtraction(doc),
		additionalRequirements:      _ExtractAdditionalRequirements(doc),
		deadlineLogging:             _ExtractDeadlineLogging(doc),
		deadlineMaterialsUsage:      _ExtractDeadlineMaterialsUsage(doc),
		cleaningProcedure:           _ExtractCleaningProcedure(doc),
		removalToTemporaryStorage:   _ExtractRemovalFromTemporaryStorage(doc),
		issuedBy:                    _ExtractIssuedBy(doc),
		whoReceivedThePermit:        _ExtractWhoReceivedThePermit(doc),
		issuedOn:                    _ExtractIssuedOn(doc),
		issuedByEmployee:            _ExtractIssuedByEmployee(doc),
		issuedCode:                  _ExtractIssuedCode(doc),
		permitIssuePlace:            _ExtractPermissionIssuePlace(doc),
	}
}

func _CleanString(str string) string {
	noBadSpaces := utilsstring.RemoveBadSpaces(str)

	isPlaceholder := utilsstring.IsPlaceholder(noBadSpaces)

	if isPlaceholder {
		return ""
	}

	return noBadSpaces
}

func _ExtractId(doc *goquery.Document) string {
	regex := regexp.MustCompile("(?i)позволително за сеч № ([0-9]+)")
	text := doc.Text()

	matched := regex.FindStringSubmatch(text)[1]

	return _CleanString(matched)
}

func _ExtractRegionalForestryDirectorate(doc *goquery.Document) string {
	regionalForestryDirectorate := doc.Find("font:contains('Регионална дирекция по горите')").Find("b").Text()

	return _CleanString(regionalForestryDirectorate)
}

func _ExtractDirectiveNumber(doc *goquery.Document) string {
	directiveNumber := doc.Find("td:contains('На основание')").First().Find("b").Text()

	return _CleanString(directiveNumber)
}

func _ExtractPermittedFor(doc *goquery.Document) string {
	permittedFor := doc.Find("td:contains('разрешава се на')").First().Find("b").Text()

	return _CleanString(permittedFor)
}

func _ExtractAllowedForester(doc *goquery.Document) string {
	allowedForester := doc.Find("td:contains('с нает регистриран лесовъд')").First().Find("b").Text()

	return _CleanString(allowedForester)
}

func _ExtractTypeOfFelling(doc *goquery.Document) string {
	typeOfFelling := doc.Find("td:contains('ВИД НА СЕЧТА')").First().Find("b").Text()

	return _CleanString(typeOfFelling)
}

func _ExtractSection(doc *goquery.Document) string {
	section := doc.Find("td:contains('да извърши добива в отдел №')").First().Find("b").First().Text()

	return _CleanString(section)
}

func _ExtractSubSection(doc *goquery.Document) string {
	subSection := doc.Find("td:contains('подотдел')").First().Find("b").Next().First().Text()

	return _CleanString(subSection)
}

func _ExtractCadastreId(doc *goquery.Document) string {
	cadastreId := doc.Find("td:contains('подотдел')").First().Find("b").Next().Next().Text()

	return _CleanString(cadastreId)
}

func _ExtractMunicipality(doc *goquery.Document) string {
	municipality := doc.Find("td:contains('Община')").First().Find("b").First().Text()

	return _CleanString(municipality)
}

func _ExtractLand(doc *goquery.Document) string {
	land := doc.Find("td:contains('Землище')").First().Find("b").Next().First().Text()

	return utilsstring.RemoveBadSpaces(land)
}

func _ExtractAreaClearing(doc *goquery.Document) float64 {
	areaClearing := doc.Find("td:contains('площно сечище от')").First().Find("b").Next().Next().First().Text()

	result, err := strconv.ParseFloat(areaClearing, 64)

	if err != nil {
		log.Fatal(err)
	}

	return result
}

func _ExtractOwnershipType(doc *goquery.Document) string {
	ownershipType := doc.Find("td:contains('Вид собственост')").First().Find("b").Next().Next().Next().First().Text()

	return _CleanString(ownershipType)
}

func _ExtractTreesMarkedBy(doc *goquery.Document) string {
	treesMarkedBy := doc.Find("td:contains('Дърветата са маркирани от')").First().Find("b").Text()

	return _CleanString(treesMarkedBy)
}

func _ExtractControlMarkNumber(doc *goquery.Document) string {
	controlMark := doc.Find("td:contains('с контролна горска марка №')").First().Find("b").First().Text()

	return _CleanString(controlMark)
}

func _ExtractMarkColor(doc *goquery.Document) string {
	markColor := doc.Find("td:contains('с контролна горска марка №')").First().Find("b").Next().First().Text()

	return _CleanString(markColor)
}

func _ExtractDateOfCarnetInventory(doc *goquery.Document) time.Time {
	dateOfCarnetInventoryLine := doc.Find("td:contains('с контролна горска марка №')").First().Find("b").Next().Next().First().Text()

	dateOfCarnetInventoryLine = _CleanString(dateOfCarnetInventoryLine)

	if dateOfCarnetInventoryLine == "" {
		return time.Time{}
	}

	dateOfCarnetInventory, err := time.Parse(DateLayout, _CleanString(dateOfCarnetInventoryLine))

	if err != nil {
		log.Fatal(err)
	}

	return dateOfCarnetInventory
}

func _ExtractExpectedTreeExtraction(doc *goquery.Document) float64 {
	expectedTreeExtraction := doc.Find("td:contains('Очакваният добив е')").First().Find("b").Text()

	result, err := strconv.ParseFloat(expectedTreeExtraction, 64)

	if err != nil {
		log.Fatal(err)
	}

	return result
}

func _ExtractAdditionalRequirements(doc *goquery.Document) string {
	additionalRequirements := doc.Find("td:contains('Допълнителни изисквания при провеждане на сечта :')").First().Find("b").Text()

	return _CleanString(additionalRequirements)
}

func _ExtractDeadlineLogging(doc *goquery.Document) utilstime.TimeRange {
	deadlineLoggingLine := doc.Find("td:contains('Срок за провеждане на сечта от')").First().Find("b")

	deadlineLoggingFrom := _CleanString(deadlineLoggingLine.First().Text())

	deadlineLoggingTo := _CleanString(deadlineLoggingLine.Next().Text())

	resultFrom, errFrom := time.Parse(DateLayout, deadlineLoggingFrom)

	if errFrom != nil {
		log.Fatal(errFrom)
	}

	resultTo, errTo := time.Parse(DateLayout, deadlineLoggingTo)

	if errTo != nil {
		log.Fatal(errTo)
	}

	return utilstime.TimeRange{From: resultFrom, To: resultTo}
}

func _ExtractDeadlineMaterialsUsage(doc *goquery.Document) utilstime.TimeRange {
	deadlineLoggingLine := doc.Find("td:contains('Срок за извозване на материалите от сечището от')").First().Find("b")

	deadlineMaterialsUsageFrom := _CleanString(deadlineLoggingLine.First().Text())

	deadlineMaterialsUsageTo := _CleanString(deadlineLoggingLine.Next().Text())

	resultFrom, errFrom := time.Parse(DateLayout, deadlineMaterialsUsageFrom)

	if errFrom != nil {
		log.Fatal(errFrom)
	}

	resultTo, errTo := time.Parse(DateLayout, deadlineMaterialsUsageTo)

	if errTo != nil {
		log.Fatal(errTo)
	}

	return utilstime.TimeRange{From: resultFrom, To: resultTo}
}

func _ExtractCleaningProcedure(doc *goquery.Document) string {
	cleaningProcedure := doc.Find("td:contains('Начин на почистване на сечището :')").First().Find("b").Text()

	return _CleanString(cleaningProcedure)
}

func _ExtractRemovalFromTemporaryStorage(doc *goquery.Document) string {
	removalFromTemporaryStorage := doc.Find("td:contains('Материалите ще се извозят до временен склад :')").First().Find("b").Text()

	return _CleanString(removalFromTemporaryStorage)
}

func _ExtractIssuedBy(doc *goquery.Document) string {
	issuedByLine := doc.Find("td:contains('Издал:')").First().Text()

	issuedBy := strings.Replace(issuedByLine, "Издал:", "", 1)

	return _CleanString(issuedBy)
}

func _ExtractWhoReceivedThePermit(doc *goquery.Document) string {
	whoReceivedThePermitLine := doc.Find("td:contains('Получил позволителното:')").First().Text()

	whoReceivedThePermit := strings.Replace(whoReceivedThePermitLine, "Получил позволителното:", "", 1)

	return _CleanString(whoReceivedThePermit)
}

func _ExtractIssuedOn(doc *goquery.Document) time.Time {
	issuedOnLine := doc.Find("td:contains('Дата:')").Find("b").Text()

	issuedOn, err := time.Parse(DateLayout, _CleanString(issuedOnLine))

	if err != nil {
		log.Fatal(err)
	}

	return issuedOn
}

func _ExtractIssuedByEmployee(doc *goquery.Document) string {
	issuedByEmployeeLine := doc.Find("td:contains('Дата:')").Text()
	noNewLines := strings.ReplaceAll(issuedByEmployeeLine, "\n", "")
	regex := regexp.MustCompile("Издал служител : \\|(.+)\\| ")

	matched := regex.FindStringSubmatch(noNewLines)[1]

	return _CleanString(matched)
}

func _ExtractIssuedCode(doc *goquery.Document) string {
	issuedCodeLine := doc.Find("td:contains('Дата:')").Text()
	noNewLines := strings.ReplaceAll(issuedCodeLine, "\n", "")
	regex := regexp.MustCompile("Код: \\|(.+)\\|")

	matched := regex.FindStringSubmatch(noNewLines)[1]

	return _CleanString(matched)
}

func _ExtractPermissionIssuePlace(doc *goquery.Document) PermitIssuePlace {
	line := doc.Find("td:contains('Област ')").First().Text()

	regex := regexp.MustCompile("Област(.+), община(.*), землище(.*), адрес(.*), подотдел(.*), GPS координати:(.*)")

	matches := regex.FindStringSubmatch(line)

	cleanedMatches := matches[1:]

	for match := range cleanedMatches {
		cleanedMatches[match] = _CleanString(cleanedMatches[match])
	}

	return PermitIssuePlace{
		region:         cleanedMatches[0],
		municipality:   cleanedMatches[1],
		land:           cleanedMatches[2],
		address:        cleanedMatches[3],
		subSection:     cleanedMatches[4],
		gpsCoordinates: cleanedMatches[5],
	}
}
