package certificate_protocol_parse_fields

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
)

const volumeDifferenceDescription = `Разликите между количествата и категориите дървесина, по позволително за сеч и фактически добити в повече или в по-малко се дължът на следното:`

func ExtractVolumeDifferenceDescription(section *goquery.Selection) string {
	contains := fmt.Sprintf(`td:contains('%s')`, volumeDifferenceDescription)

	line := section.Find(contains).First().Find("b").First().Text()

	return line
}
