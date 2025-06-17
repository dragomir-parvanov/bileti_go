package parse_permitted_for_felling_fields

import (
	"bileti_go/src/utils/string"
	"github.com/PuerkitoBio/goquery"
	"regexp"
)

func ExtractCadastreId(doc *goquery.Selection) string {
	line := doc.Find("td:contains('да извърши добива в отдел № ')").First().Text()

	regex := regexp.MustCompile(`; имот с \(кадастрален/КВС\) № (.*), находящ се в:`)

	match := regex.FindStringSubmatch(line)[1]

	return utils_string.RemoveBadSpaces(match)
}
