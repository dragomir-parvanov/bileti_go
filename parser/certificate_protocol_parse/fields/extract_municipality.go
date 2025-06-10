package certificate_protocol_parse_fields

import (
	"github.com/PuerkitoBio/goquery"
)

func ExtractMunicipality(selection *goquery.Selection) string {

	line := selection.Find("td:contains('След съвместно обхождане и проверка на сечището в Община: ')").
		First().
		Find("b").
		Next().
		Text()

	return line
}
