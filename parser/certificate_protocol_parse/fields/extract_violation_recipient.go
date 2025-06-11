package certificate_protocol_parse_fields

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
)

func ExtractViolationRecipient(section *goquery.Selection) string {
	contains := fmt.Sprintf(`td:contains('За допуснатите нарушения е съставен акт за установяване на административно')`)

	line := section.Find(contains).First().Find("b").Next().Next().Next().Next().Text()

	return line
}
