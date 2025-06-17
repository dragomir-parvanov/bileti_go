package certificate_protocol_parse_fields

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
)

const loggingFinishDescription = `Провеждането на сечта и почистването на сечището е извършено според определения начин в технологичния план и позволителното за сеч:`

func ExtractLoggingFinishDescription(section *goquery.Selection) string {
	contains := fmt.Sprintf(`td:contains('%s')`, loggingFinishDescription)

	line := section.Find(contains).First().Find("b").Next().Text()

	return line
}
