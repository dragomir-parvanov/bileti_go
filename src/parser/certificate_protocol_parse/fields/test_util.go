package certificate_protocol_parse_fields

import (
	"bileti_go/src/utils"
	"github.com/PuerkitoBio/goquery"
	"strings"
)

func GetTestSelection(html string) *goquery.Selection {
	doc := utils.Must(goquery.NewDocumentFromReader(strings.NewReader(html)))

	return doc.Selection
}
