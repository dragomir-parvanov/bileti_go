package certificate_protocol_parse_fields

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"strings"
)

func GetTestSelection(html string) *goquery.Selection {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))

	if err != nil {
		log.Fatal(err)
	}

	return doc.Selection
}
