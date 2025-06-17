package certificate_protocol_parse_fields

import (
	"github.com/PuerkitoBio/goquery"
	"regexp"
)

func ExtractPermittedForFellingId(selection *goquery.Selection) string {
	line := selection.Find("td:contains('Към позволително за сеч № ')").First().Text()

	regex := regexp.MustCompile(`Към позволително за сеч № (.*)/(.*) год.`)

	id := regex.FindStringSubmatch(line)[1]

	return id
}
