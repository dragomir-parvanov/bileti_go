package certificate_protocol_parse_fields

import (
	"github.com/PuerkitoBio/goquery"
	"regexp"
)

func ExtractCadastreId(section *goquery.Selection) string {

	line := section.Find("td:contains('; Отдел: ')").First().Text()

	regex := regexp.MustCompile(`; Кад. № (.*);`)

	match := regex.FindStringSubmatch(line)[1]

	return match
}
