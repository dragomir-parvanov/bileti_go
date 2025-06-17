package parse_permitted_for_felling_fields

import (
	"bileti_go/src/parser"
	"bileti_go/src/utils"
	utils_string "bileti_go/src/utils/string"
	"github.com/PuerkitoBio/goquery"
	"regexp"
	"time"
)

func ExtractLoggingDeadlineToExtension(selection *goquery.Selection) time.Time {
	line := selection.Find("td:contains('За провеждане на сечта до : ')").First().Text()

	cleanedLine := utils_string.CleanString(line)

	regex := regexp.MustCompile("За провеждане на сечта до :(.*) г.")

	match := regex.FindStringSubmatch(cleanedLine)[1]

	cleanedMatch := utils_string.CleanString(match)

	if cleanedMatch == "" {
		return time.Time{}
	}

	result := utils.Must(time.ParseInLocation(parser.DateLayout, cleanedMatch, parser.GetLocation()))

	return result
}
