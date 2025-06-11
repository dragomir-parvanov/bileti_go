package certificate_protocol_parse_fields

import (
	"github.com/PuerkitoBio/goquery"
	"regexp"
)

func ExtractSection(section *goquery.Selection) string {

	line := section.Find("td:contains('; Отдел: ')").First().Text()

	regex := regexp.MustCompile(` Отдел: (.*); Подотдел: `)

	match := regex.FindStringSubmatch(line)[1]

	return match
}
