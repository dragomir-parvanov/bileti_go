package certificate_protocol_parse_fields

import (
	"github.com/PuerkitoBio/goquery"
)

func ExtractOwnership(selection *goquery.Selection) string {
	line := selection.Find("td:contains('собственост на: ')").First().Find("b").First().Text()

	return line
}
