package parse_permitted_for_felling_fields

import (
	"github.com/PuerkitoBio/goquery"
	"regexp"
)

func ExtractAccordingToTheInventoryOf(selection *goquery.Selection) string {
	line := selection.Find("td:contains('да извърши добива в отдел № ')").First().Text()

	regex := regexp.MustCompile("подотдел з съгласно инвентаризацията на (.*), утвърдена със заповед №")

	matches := regex.FindStringSubmatch(line)

	if len(matches) >= 2 {
		return matches[1]
	}

	return ""
}
