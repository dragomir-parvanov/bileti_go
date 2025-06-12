package parse_permitted_for_felling_fields

import (
	"github.com/PuerkitoBio/goquery"
	"regexp"
)

func ExtractInventoryOrderId(selection *goquery.Selection) string {
	line := selection.Find("td:contains('да извърши добива в отдел № ')").First().Text()

	regex := regexp.MustCompile(`да извърши добива в отдел № 256; подотдел з съгласно инвентаризацията на (.*), утвърдена със заповед № (.*); имот с \(кадастрален/КВС\) № `)

	matches := regex.FindStringSubmatch(line)

	if len(matches) >= 3 {
		return matches[2]
	}

	return ""
}
