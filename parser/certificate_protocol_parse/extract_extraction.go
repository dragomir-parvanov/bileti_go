package certificate_protocol_parse

import (
	"bileti_go/parser"
	"bileti_go/utils"
	utilsstring "bileti_go/utils/string"
	"github.com/PuerkitoBio/goquery"
	"regexp"
	"slices"
	"strconv"
)

func ExtractExtraction(selection *goquery.Selection) []CertificateProtocolTreeExtraction {
	table := selection.Find("tbody:contains('Категория дървесина')")

	strings := getTdStrings(table)

	cleanedStrings := cleanStrings(strings)

	categorizedStrings := getCategorizedStrings(cleanedStrings)

	return extractExtractions(categorizedStrings)
}

func getTdStrings(table *goquery.Selection) []string {
	var strings []string

	trs := table.Find("tr")

	trs.Each(func(i int, s *goquery.Selection) {
		var subStrings []string

		s.Find("td").
			Each(func(i int, s *goquery.Selection) {
				subStrings = append(subStrings, s.Text())
			})

		strings = append(strings, subStrings...)
	})

	return strings
}

func cleanStrings(strings []string) []string {
	var cleanedStrings []string

	for _, s := range strings {
		cleanedStrings = append(cleanedStrings, utilsstring.CleanString(s))
	}

	return cleanedStrings
}

func getSplitStrings() map[parser.ExtractionCategory][]string {
	return map[parser.ExtractionCategory][]string{
		parser.LargeConstructionTimber:  make([]string, 0),
		parser.MediumConstructionTimber: make([]string, 0),
		parser.SmallConstructionTimber:  make([]string, 0),
		parser.Wood:                     make([]string, 0),
		parser.TopHamper:                make([]string, 0),
	}
}

func getCategorizedStrings(strings []string) map[parser.ExtractionCategory][]string {
	splitStrings := getSplitStrings()

	var activeCategory parser.ExtractionCategory

	for i, str := range strings {
		if i == 0 {
			continue
		}

		if str == "Общо количество:" {
			break
		}

		_isSeparator := isSeparator(str, strings[i-1])

		if _isSeparator {
			activeCategory = getReverseMap()[str]
			continue
		}

		if activeCategory != "" && !isNumberSeparator(str) {
			splitStrings[activeCategory] = append(splitStrings[activeCategory], str)
		}
	}

	return splitStrings
}

const columnsNumber = 5

func extractExtractions(categorizedStrings map[parser.ExtractionCategory][]string) []CertificateProtocolTreeExtraction {
	var extractions []CertificateProtocolTreeExtraction

	for k, categoryStrings := range categorizedStrings {
		if isEmptyCategory(categoryStrings) {
			continue
		}

		for chunk := range slices.Chunk(categoryStrings, columnsNumber) {
			extractions = append(extractions, makeTreeExtraction(k, chunk))
		}
	}

	return extractions
}

func isEmptyCategory(strings []string) bool {
	if len(strings) != columnsNumber {
		return false
	}

	areAllEmpty := strings[0] == "" && strings[1] == "" && strings[2] == ""

	return areAllEmpty
}

func makeTreeExtraction(category parser.ExtractionCategory, chunk []string) CertificateProtocolTreeExtraction {

	return CertificateProtocolTreeExtraction{
		category:                    category,
		treeType:                    chunk[0],
		byLoggingPermitCollected:    parseFloat(chunk[1]),
		actuallyCollected:           parseFloat(chunk[2]),
		availableInTheClearing:      parseFloat(chunk[3]),
		availableInTemporaryStorage: parseFloat(chunk[4]),
	}
}

func parseFloat(str string) float64 {
	if str == "" {
		return 0
	}

	parsedValue := utils.Must(strconv.ParseFloat(str, 64))

	return parsedValue
}

func isSeparator(str string, prevString string) bool {

	if getReverseMap()[str] != "" {
		return isNumberSeparator(prevString)
	}

	return false
}

func isNumberSeparator(str string) bool {
	regex := regexp.MustCompile(`^[1-5]\.$`)

	isMatch := regex.MatchString(str)

	return isMatch
}

func getMap() map[parser.ExtractionCategory]string {
	categoryMap := map[parser.ExtractionCategory]string{
		parser.LargeConstructionTimber:  "Едра строителна дървесина",
		parser.MediumConstructionTimber: "Средна строителна дървесина",
		parser.SmallConstructionTimber:  "Дребна строителна дървесина",
		parser.Wood:                     "Дърва",
		parser.TopHamper:                "Вършина",
	}

	return categoryMap
}

func getReverseMap() map[string]parser.ExtractionCategory {
	categoryMap := getMap()

	reversedMap := make(map[string]parser.ExtractionCategory)

	for k, v := range categoryMap {
		reversedMap[v] = k
	}

	return reversedMap
}
