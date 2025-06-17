package parse_permitted_for_felling

import (
	"bileti_go/src/parser"
	parse_permitted_for_felling_fields "bileti_go/src/parser/parse_permitted_for_felling/fields"
	"bileti_go/src/utils"
	"bileti_go/src/utils/string"
	utilstime "bileti_go/src/utils/time"
	"github.com/PuerkitoBio/goquery"
	"io"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func ParsePermittedForFelling(htmlResult io.Reader) ParsedResult {

	doc := utils.Must(goquery.NewDocumentFromReader(htmlResult))

	tables := doc.Find("table")

	headerTable := tables.First()

	mainTable := tables.Next()

	categoryTable := mainTable.Next()

	additionalTable := categoryTable.Next()

	return ParsedResult{
		id:                          extractId(mainTable),
		regionalForestryDirectorate: extractRegionalForestryDirectorate(headerTable),
		directiveNumber:             extractDirectiveNumber(mainTable),
		permittedFor:                extractPermittedFor(mainTable),
		allowedForester:             extractAllowedForester(mainTable),
		typeOfFelling:               extractTypeOfFelling(mainTable),
		section:                     extractSection(mainTable),
		subSection:                  extractSubSection(mainTable),
		accordingToTheInventoryOf:   parse_permitted_for_felling_fields.ExtractAccordingToTheInventoryOf(mainTable),
		inventoryOrderId:            parse_permitted_for_felling_fields.ExtractInventoryOrderId(mainTable),
		cadastreId:                  parse_permitted_for_felling_fields.ExtractCadastreId(mainTable),
		municipality:                extractMunicipality(mainTable),
		land:                        extractLand(mainTable),
		areaClearing:                extractAreaClearing(mainTable),
		ownershipType:               extractOwnershipType(mainTable),
		treesMarkedBy:               extractTreesMarkedBy(mainTable),
		controlMarkNumber:           extractControlMarkNumber(mainTable),
		controlMarkColor:            extractMarkColor(mainTable),
		dateOfCarnetInventory:       parse_permitted_for_felling_fields.ExtractDateOfCarnetInventory(mainTable),
		expectedTreeExtraction:      extractExpectedTreeExtraction(mainTable),
		extraction:                  ParseTreeExtraction(categoryTable),
		additionalRequirements:      extractAdditionalRequirements(additionalTable),
		deadlineLogging:             extractDeadlineLogging(additionalTable),
		deadlineTransporting:        extractDeadlineTransporting(additionalTable),
		cleaningProcedure:           extractCleaningProcedure(additionalTable),
		removalToTemporaryStorage:   extractRemovalFromTemporaryStorage(additionalTable),
		issuedBy:                    extractIssuedBy(additionalTable),
		whoReceivedThePermit:        extractWhoReceivedThePermit(additionalTable),
		issuedOn:                    extractIssuedOn(additionalTable),
		issuedByEmployee:            extractIssuedByEmployee(additionalTable),
		issuedCode:                  extractIssuedCode(additionalTable),
		permitIssuePlace:            extractPermissionIssuePlace(additionalTable),
		extension:                   extractExtension(additionalTable),
	}
}

func extractId(doc *goquery.Selection) string {
	regex := regexp.MustCompile(`(?i)позволително за сеч № ([0-9]+)`)
	text := doc.Text()

	matched := regex.FindStringSubmatch(text)[1]

	return utils_string.CleanString(matched)
}

func extractRegionalForestryDirectorate(doc *goquery.Selection) string {
	regionalForestryDirectorate := doc.Find("font:contains('Регионална дирекция по горите')").Find("b").Text()

	return utils_string.CleanString(regionalForestryDirectorate)
}

func extractDirectiveNumber(doc *goquery.Selection) string {
	directiveNumber := doc.Find("td:contains('На основание')").First().Find("b").Text()

	return utils_string.CleanString(directiveNumber)
}

func extractPermittedFor(doc *goquery.Selection) string {
	permittedFor := doc.Find("td:contains('разрешава се на')").First().Find("b").Text()

	return utils_string.CleanString(permittedFor)
}

func extractAllowedForester(doc *goquery.Selection) string {
	allowedForester := doc.Find("td:contains('с нает регистриран лесовъд')").First().Find("b").Text()

	return utils_string.CleanString(allowedForester)
}

func extractTypeOfFelling(doc *goquery.Selection) string {
	typeOfFelling := doc.Find("td:contains('ВИД НА СЕЧТА')").First().Find("b").Text()

	return utils_string.CleanString(typeOfFelling)
}

func extractSection(doc *goquery.Selection) string {
	section := doc.Find("td:contains('да извърши добива в отдел №')").First().Find("b").First().Text()

	return utils_string.CleanString(section)
}

func extractSubSection(doc *goquery.Selection) string {
	subSection := doc.Find("td:contains('подотдел')").First().Find("b").Next().First().Text()

	return utils_string.CleanString(subSection)
}

func extractMunicipality(doc *goquery.Selection) string {
	municipality := doc.Find("td:contains('Община')").First().Find("b").First().Text()

	return utils_string.CleanString(municipality)
}

func extractLand(doc *goquery.Selection) string {
	land := doc.Find("td:contains('Землище')").First().Find("b").Next().First().Text()

	return utils_string.RemoveBadSpaces(land)
}

func extractAreaClearing(doc *goquery.Selection) float64 {
	areaClearing := doc.Find("td:contains('площно сечище от')").First().Find("b").Next().Next().First().Text()

	result := utils.Must(strconv.ParseFloat(areaClearing, 64))

	return result
}

func extractOwnershipType(doc *goquery.Selection) string {
	ownershipType := doc.Find("td:contains('Вид собственост')").First().Find("b").Next().Next().Next().First().Text()

	return utils_string.CleanString(ownershipType)
}

func extractTreesMarkedBy(doc *goquery.Selection) string {
	treesMarkedBy := doc.Find("td:contains('Дърветата са маркирани от')").First().Find("b").Text()

	return utils_string.CleanString(treesMarkedBy)
}

func extractControlMarkNumber(doc *goquery.Selection) string {
	controlMark := doc.Find("td:contains('с контролна горска марка №')").First().Find("b").First().Text()

	return utils_string.CleanString(controlMark)
}

func extractMarkColor(doc *goquery.Selection) string {
	markColor := doc.Find("td:contains('с контролна горска марка №')").First().Find("b").Next().First().Text()

	return utils_string.CleanString(markColor)
}

func extractExpectedTreeExtraction(doc *goquery.Selection) float64 {
	expectedTreeExtraction := doc.Find("td:contains('Очакваният добив е')").First().Find("b").Text()

	result := utils.Must(strconv.ParseFloat(expectedTreeExtraction, 64))

	return result
}

func extractAdditionalRequirements(doc *goquery.Selection) string {
	additionalRequirements := doc.Find("td:contains('Допълнителни изисквания при провеждане на сечта :')").First().Find("b").Text()

	return utils_string.CleanString(additionalRequirements)
}

func extractDeadlineLogging(doc *goquery.Selection) utilstime.TimeRange {
	deadlineLoggingLine := doc.Find("td:contains('Срок за провеждане на сечта от')").First().Find("b")

	deadlineLoggingFrom := utils_string.CleanString(deadlineLoggingLine.First().Text())

	deadlineLoggingTo := utils_string.CleanString(deadlineLoggingLine.Next().Text())

	resultFrom := utils.Must(time.ParseInLocation(parser.DateLayout, deadlineLoggingFrom, parser.GetLocation()))

	resultTo := utils.Must(time.ParseInLocation(parser.DateLayout, deadlineLoggingTo, parser.GetLocation()))

	return utilstime.TimeRange{From: resultFrom, To: resultTo}
}

func extractDeadlineTransporting(doc *goquery.Selection) utilstime.TimeRange {
	deadlineLoggingLine := doc.Find("td:contains('Срок за извозване на материалите от сечището от')").First().Find("b")

	transportingDeadlineFrom := utils_string.CleanString(deadlineLoggingLine.First().Text())

	transportingDeadlineTo := utils_string.CleanString(deadlineLoggingLine.Next().Text())

	resultFrom := utils.Must(time.ParseInLocation(parser.DateLayout, transportingDeadlineFrom, parser.GetLocation()))

	resultTo := utils.Must(time.ParseInLocation(parser.DateLayout, transportingDeadlineTo, parser.GetLocation()))

	return utilstime.TimeRange{From: resultFrom, To: resultTo}
}

func extractCleaningProcedure(doc *goquery.Selection) string {
	cleaningProcedure := doc.Find("td:contains('Начин на почистване на сечището :')").First().Find("b").Text()

	return utils_string.CleanString(cleaningProcedure)
}

func extractRemovalFromTemporaryStorage(doc *goquery.Selection) string {
	removalFromTemporaryStorage := doc.Find("td:contains('Материалите ще се извозят до временен склад :')").First().Find("b").Text()

	return utils_string.CleanString(removalFromTemporaryStorage)
}

func extractIssuedBy(doc *goquery.Selection) string {
	issuedByLine := doc.Find("td:contains('Издал:')").First().Text()

	issuedBy := strings.Replace(issuedByLine, "Издал:", "", 1)

	return utils_string.CleanString(issuedBy)
}

func extractWhoReceivedThePermit(doc *goquery.Selection) string {
	whoReceivedThePermitLine := doc.Find("td:contains('Получил позволителното:')").First().Text()

	whoReceivedThePermit := strings.Replace(whoReceivedThePermitLine, "Получил позволителното:", "", 1)

	return utils_string.CleanString(whoReceivedThePermit)
}

func extractIssuedOn(doc *goquery.Selection) time.Time {
	issuedOnLine := doc.Find("td:contains('Дата:')").Find("b").Text()

	issuedOn := utils.Must(time.ParseInLocation(parser.DateLayout, utils_string.CleanString(issuedOnLine), parser.GetLocation()))

	return issuedOn
}

func extractIssuedByEmployee(doc *goquery.Selection) string {
	issuedByEmployeeLine := doc.Find("td:contains('Дата:')").Text()
	noNewLines := strings.ReplaceAll(issuedByEmployeeLine, "\n", "")
	regex := regexp.MustCompile(`Издал служител : \|(.+)\| `)

	matched := regex.FindStringSubmatch(noNewLines)[1]

	return utils_string.CleanString(matched)
}

func extractIssuedCode(doc *goquery.Selection) string {
	issuedCodeLine := doc.Find("td:contains('Дата:')").Text()
	noNewLines := strings.ReplaceAll(issuedCodeLine, "\n", "")
	regex := regexp.MustCompile(`Код: \|(.+)\|`)

	matched := regex.FindStringSubmatch(noNewLines)[1]

	return utils_string.CleanString(matched)
}

func extractPermissionIssuePlace(doc *goquery.Selection) PermitIssuePlace {
	line := doc.Find("td:contains('Област ')").First().Text()

	regex := regexp.MustCompile("Област(.+), община(.*), землище(.*), адрес(.*), подотдел(.*), GPS координати:(.*)")

	matches := regex.FindStringSubmatch(line)

	cleanedMatches := matches[1:]

	for match := range cleanedMatches {
		cleanedMatches[match] = utils_string.CleanString(cleanedMatches[match])
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

func extractExtension(doc *goquery.Selection) Extension {

	return Extension{
		loggingDeadlineTo:      parse_permitted_for_felling_fields.ExtractLoggingDeadlineToExtension(doc),
		transportingDeadlineTo: parse_permitted_for_felling_fields.ExtractTransportingDeadlineToExtension(doc),
		issuedBy:               extractIssuedByExtension(doc),
	}
}

func extractIssuedByExtension(doc *goquery.Selection) string {
	issuedByLine := doc.Find("td:contains('Издал: ')").First().Text()

	issuedBy := strings.Replace(issuedByLine, "Издал: ", "", 1)

	return utils_string.CleanString(issuedBy)
}
