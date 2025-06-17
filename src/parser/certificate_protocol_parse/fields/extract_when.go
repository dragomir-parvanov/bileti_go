package certificate_protocol_parse_fields

import (
	"bileti_go/src/parser"
	utilsstring "bileti_go/src/utils/string"
	"github.com/PuerkitoBio/goquery"
	"time"
)

func ExtractWhen(selection *goquery.Selection) (time.Time, error) {

	loc := parser.GetLocation()

	line := selection.Find("td:contains('Днес: ')").First().Find("b").First().Text()

	cleaned := utilsstring.CleanString(line)

	result, err := time.ParseInLocation(parser.DateLayout, cleaned, loc)

	return result, err
}
