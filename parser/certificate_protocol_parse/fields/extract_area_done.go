package certificate_protocol_parse_fields

import (
	"bileti_go/utils"
	"github.com/PuerkitoBio/goquery"
	"strconv"
)

func ExtractAreaDone(selection *goquery.Selection) float64 {
	line := selection.Find("td:contains('Сечта е проведена на площ (ха): ')").First().Find("b").Next().Text()

	value := utils.Must(strconv.ParseFloat(line, 64))

	return value
}
