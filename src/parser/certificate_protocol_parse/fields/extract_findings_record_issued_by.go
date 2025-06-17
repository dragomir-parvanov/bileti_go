package certificate_protocol_parse_fields

import (
	"bileti_go/src/utils/string"
	"github.com/PuerkitoBio/goquery"
)

func ExtractFindingsRecordIssuedBy(selection *goquery.Selection) string {

	line := selection.Find("td:contains('констативен протокол от: ')").First().Find("b").First().Text()

	return utils_string.CleanString(line)
}
