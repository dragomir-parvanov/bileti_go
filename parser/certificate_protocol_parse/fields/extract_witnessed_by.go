package certificate_protocol_parse_fields

import (
	utils_string "bileti_go/utils/string"
	"github.com/PuerkitoBio/goquery"
)

import "regexp"

func ExtractWitnessedBy(selection *goquery.Selection) string {
	line := selection.Find("td:contains('Присъствал: ')").
		First().
		Text()

	regex := regexp.MustCompile(`Присъствал: (.*)`)

	witnessedBy := regex.FindStringSubmatch(line)[1]

	return utils_string.CleanString(witnessedBy)
}
