package certificate_protocol_parse_fields

import (
	"bileti_go/src/parser"
	"bileti_go/src/utils/string"
	"github.com/PuerkitoBio/goquery"
	"time"
)

func ExtractToImproveClearingDeadline(selection *goquery.Selection) (time.Time, error) {
	line := selection.Find("td:contains('За подобряване състоянието на насаждението в срок до:')").
		First().
		Find("b").
		Next().
		First().
		Text()

	cleanedLine := utils_string.CleanString(line)

	if cleanedLine == "" {
		return time.Time{}, nil
	}

	date, err := time.ParseInLocation(parser.DateLayout, cleanedLine, parser.GetLocation())

	return date, err
}
