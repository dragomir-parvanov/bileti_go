package parse_permitted_for_felling_fields

import (
	"bileti_go/utils"
	"github.com/PuerkitoBio/goquery"
	"strings"
)

func GetTestSelection(html string) *goquery.Selection {
	doc := utils.Must(goquery.NewDocumentFromReader(strings.NewReader(html)))

	return doc.Selection
}
