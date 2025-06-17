package certificate_protocol_parse_fields

import (
	"github.com/PuerkitoBio/goquery"
	"regexp"
)

func ExtractViolationId(selection *goquery.Selection) string {
	line := selection.Find("td:contains('За допуснатите нарушения е съставен акт за установяване на административно')").
		First().
		Text()

	regex := regexp.MustCompile(`нарушение № (.*) /`)

	match := regex.FindStringSubmatch(line)[1]

	return match
}
