package certificate_protocol_parse_fields


import "github.com/PuerkitoBio/goquery"

import "regexp"

func ExtractId(selection *goquery.Selection) string {
	line := selection.Find("td:contains('ПРОТОКОЛ за освидетелстване на сечище')").
		First().
		Text()

	regex := regexp.MustCompile(`ПРОТОКОЛ за освидетелстване на сечище № (.*)`)

	id := regex.FindStringSubmatch(line)[1]

	return id
}
