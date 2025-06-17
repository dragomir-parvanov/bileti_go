package certificate_protocol_parse_fields

import (
	"github.com/PuerkitoBio/goquery"
)

func ExtractLand(selection *goquery.Selection) string {
	line := selection.Find("td:contains('землище: ')").First().Find("b").First().Text()

	return line
}
