package certificate_protocol_parse_fields

import (
	"github.com/PuerkitoBio/goquery"
)

func ExtractViolationRecipient(section *goquery.Selection) string {
	contains := "td:contains('За допуснатите нарушения е съставен акт за установяване на административно')"

	line := section.Find(contains).First().Find("b").Next().Next().Next().Next().Text()

	return line
}
