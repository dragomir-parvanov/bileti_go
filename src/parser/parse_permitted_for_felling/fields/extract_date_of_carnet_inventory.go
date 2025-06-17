package parse_permitted_for_felling_fields

import (
	"bileti_go/src/parser"
	"bileti_go/src/utils"
	utils_string "bileti_go/src/utils/string"
	"github.com/PuerkitoBio/goquery"
	"time"
)

func ExtractDateOfCarnetInventory(selection *goquery.Selection) time.Time {
	dateOfCarnetInventoryLine := selection.Find("td:contains('с контролна горска марка №')").First().Find("b").Next().Next().First().Text()

	dateOfCarnetInventoryLine = utils_string.CleanString(dateOfCarnetInventoryLine)

	if dateOfCarnetInventoryLine == "" {
		return time.Time{}
	}

	cleanedLine := utils_string.CleanString(dateOfCarnetInventoryLine)

	result, err := time.ParseInLocation(parser.DateLayout, cleanedLine, parser.GetLocation())

	dateOfCarnetInventory := utils.Must(result, err)

	return dateOfCarnetInventory
}
