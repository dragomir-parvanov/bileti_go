package certificate_protocol_parse_fields

import (
	"github.com/PuerkitoBio/goquery"
)

func ExtractRegionalForestryDirectorate(selection *goquery.Selection) string {

	line := selection.Find("font:contains('РЕГИОНАЛНА ДИРЕНЦИЯ ПО ГОРИТЕ')").First().Find("b").Text()

	return line
}
