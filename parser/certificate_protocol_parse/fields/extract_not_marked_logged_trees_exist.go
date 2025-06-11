package certificate_protocol_parse_fields

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
)

func ExtractNotMarkedLoggedTreesExist(section *goquery.Selection) bool {
	line := section.Find("td:contains('Налични неотсечени маркирани стъбла :')").First().Find("b").First().Text()

	if line == "Да" {
		return true
	}

	if line == "Не" {
		return false
	}

	errMessage := fmt.Sprintf(`Expected either "Да" or "Не", received %s`, line)

	panic(errMessage)
}
