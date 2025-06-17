package certificate_protocol_parse_fields

import (
	"bileti_go/src/utils/string"
	"github.com/PuerkitoBio/goquery"
)

func ExtractFindingsRecordWitnessedBy(selection *goquery.Selection) string {

	line := selection.Find("td:contains('в присъствието на: ')").First().Find("b").First().Text()

	return utils_string.CleanString(line)
}
