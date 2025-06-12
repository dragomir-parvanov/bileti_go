package parse_permitted_for_felling

import (
	parser_extraction_category "bileti_go/parser/extraction_category"
	"bileti_go/utils"
	utilsstring "bileti_go/utils/string"
	"github.com/PuerkitoBio/goquery"
	"regexp"
	"slices"
	"strconv"
)

func ParseTreeExtraction(selection *goquery.Selection) []TreeExtraction {
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

func getSplitStrings() map[parser_extraction_category.ExtractionCategory][]string {
	return map[parser_extraction_category.ExtractionCategory][]string{
		parser_extraction_category.LargeConstructionTimber:  make([]string, 0),
		parser_extraction_category.MediumConstructionTimber: make([]string, 0),
		parser_extraction_category.SmallConstructionTimber:  make([]string, 0),
		parser_extraction_category.Wood:                     make([]string, 0),
		parser_extraction_category.TopHamper:                make([]string, 0),
	}
}

func getCategorizedStrings(strings []string) map[parser_extraction_category.ExtractionCategory][]string {
	splitStrings := getSplitStrings()

	var activeCategory parser_extraction_category.ExtractionCategory

	for i, str := range strings {
		if i == 0 {
			continue
		}

		_isSeparator := isSeparator(str, strings[i-1])

		if _isSeparator {
			activeCategory = parser_extraction_category.GetReverseMap()[str]
			continue
		}

		if activeCategory != "" && !isNumberSeparator(str) {
			splitStrings[activeCategory] = append(splitStrings[activeCategory], str)
		}
	}

	return splitStrings
}

func extractExtractions(categorizedStrings map[parser_extraction_category.ExtractionCategory][]string) []TreeExtraction {
	var extractions []TreeExtraction

	for k, categoryStrings := range categorizedStrings {
		if isEmptyCategory(categoryStrings) {
			continue
		}

		for chunk := range slices.Chunk(categoryStrings, 3) {
			extractions = append(extractions, makeTreeExtraction(k, chunk))
		}
	}

	return extractions
}

func isEmptyCategory(strings []string) bool {
	if len(strings) != 3 {
		return false
	}

	areAllEmpty := strings[0] == "" && strings[1] == "" && strings[2] == ""

	return areAllEmpty
}

func makeTreeExtraction(category parser_extraction_category.ExtractionCategory, chunk []string) TreeExtraction {

	return TreeExtraction{
		category: category,
		treeType: chunk[0],
		value:    parseFloat(chunk[1]),
		note:     chunk[2],
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

	if parser_extraction_category.GetReverseMap()[str] != "" {
		return isNumberSeparator(prevString)
	}

	return false
}

func isNumberSeparator(str string) bool {
	regex := regexp.MustCompile(`^[1-5]\.$`)

	isMatch := regex.MatchString(str)

	return isMatch
}
