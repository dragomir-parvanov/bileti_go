package certificate_protocol_parse_fields

import (
	"bileti_go/src/parser"
	"github.com/PuerkitoBio/goquery"
	"regexp"
	"time"
)

func ExtractPermittedForFellingDate(selection *goquery.Selection) (time.Time, error) {
	line := selection.Find("td:contains('Към позволително за сеч № ')").First().Text()

	regex := regexp.MustCompile(`Към позволително за сеч № (.*)/(.*) год.`)

	id := regex.FindStringSubmatch(line)[2]

	date, err := time.ParseInLocation(parser.DateLayout, id, parser.GetLocation())

	return date, err
}
