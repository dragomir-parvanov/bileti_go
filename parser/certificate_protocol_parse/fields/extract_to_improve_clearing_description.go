package certificate_protocol_parse_fields

import (
	"github.com/PuerkitoBio/goquery"
)

func ExtractToImproveClearingDescription(selection *goquery.Selection) string {
	line := selection.Find("td:contains('За подобряване състоянието на насаждението в срок до:')").
		First().
		Find("b").
		Next().
		Next().
		Text()

	return line
}
