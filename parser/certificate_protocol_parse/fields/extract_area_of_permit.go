package certificate_protocol_parse_fields

import (
	"bileti_go/utils"
	"github.com/PuerkitoBio/goquery"
	"strconv"
)

func ExtractAreaOfPermit(selection *goquery.Selection) float64 {
	line := selection.Find("td:contains('Площ по позволително за сеч (ха): ')").First().Find("b").First().Text()

	value := utils.Must(strconv.ParseFloat(line, 64))

	return value
}
