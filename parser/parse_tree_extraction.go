package parser

import (
	utilsstring "bileti_go/utils/string"
	"github.com/PuerkitoBio/goquery"
	"log"
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

func getSplitStrings() map[ExtractionCategory][]string {
	return map[ExtractionCategory][]string{
		LargeConstructionTimber:  make([]string, 0),
		MediumConstructionTimber: make([]string, 0),
		SmallConstructionTimber:  make([]string, 0),
		Wood:                     make([]string, 0),
		TopHamper:                make([]string, 0),
	}
}

func getCategorizedStrings(strings []string) map[ExtractionCategory][]string {
	splitStrings := getSplitStrings()

	var activeCategory ExtractionCategory

	for i, str := range strings {
		if i == 0 {
			continue
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

func extractExtractions(categorizedStrings map[ExtractionCategory][]string) []TreeExtraction {
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

func makeTreeExtraction(category ExtractionCategory, chunk []string) TreeExtraction {

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

	parsedValue, err := strconv.ParseFloat(str, 64)

	if err != nil {
		log.Fatal(err)
	}

	return parsedValue
}

func isSeparator(str string, prevString string) bool {

	if getReverseMap()[str] != "" {
		return isNumberSeparator(prevString)
	}

	return false
}

func isNumberSeparator(str string) bool {
	regex := regexp.MustCompile("^[1-5]\\.$")

	isMatch := regex.MatchString(str)

	return isMatch
}

func getMap() map[ExtractionCategory]string {
	categoryMap := map[ExtractionCategory]string{
		LargeConstructionTimber:  "Едра строителна дървесина",
		MediumConstructionTimber: "Средна строителна дървесина",
		SmallConstructionTimber:  "Дребна строителна дървесина",
		Wood:                     "Дърва",
		TopHamper:                "Вършина",
	}

	return categoryMap
}

func getReverseMap() map[string]ExtractionCategory {
	categoryMap := getMap()

	reversedMap := make(map[string]ExtractionCategory)

	for k, v := range categoryMap {
		reversedMap[v] = k
	}

	return reversedMap
}
