package certificate_protocol_parse_fields

import (
	"github.com/PuerkitoBio/goquery"
	"regexp"
)

func ExtractSubsection(section *goquery.Selection) string {

	line := section.Find("td:contains('; Отдел: ')").First().Text()

	regex := regexp.MustCompile(`; Подотдел: (.*); `)

	match := regex.FindStringSubmatch(line)[1]

	return match
}
