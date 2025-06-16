package permit_table_parser

import (
	"github.com/PuerkitoBio/goquery"
	"io"
)

func ParsePermitTable(reader io.Reader) ([]string, error) {

	doc, err := goquery.NewDocumentFromReader(reader)

	if err != nil {
		return nil, err
	}

	elements := doc.Find("td > a[href]")

	urls := make([]string, elements.Length())

	elements.Each(func(i int, element *goquery.Selection) {
		urls[i] = element.AttrOr("href", "")
	})

	return urls, nil
}
