package certificate_protocol_parse_fields

import (
	"bileti_go/src/utils/string"
	"github.com/PuerkitoBio/goquery"
)

import "regexp"

func ExtractIssuedBy(selection *goquery.Selection) string {
	line := selection.Find("td:contains('Съставил: ')").
		First().
		Text()

	regex := regexp.MustCompile(`Съставил: (.*)`)

	issuedBy := regex.FindStringSubmatch(line)[1]

	return utils_string.CleanString(issuedBy)
}
