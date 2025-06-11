package certificate_protocol_parse_fields

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
)

func ExtractViolationDescription(section *goquery.Selection) string {
	contains := fmt.Sprintf(`td:contains('При сечта и извоза на дървесина са допуснати следните нарушения на Закона за горите:')`)

	line := section.Find(contains).First().Find("b").Next().Text()

	return line
}
