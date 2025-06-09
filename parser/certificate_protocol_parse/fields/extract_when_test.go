package certificate_protocol_parse_fields

import (
	"bileti_go/parser"
	"log"
	"testing"
	"time"
)

func TestOnWhenShouldReturnTheTime(t *testing.T) {
	html := `
<table>
<tr>
  	<td colspan="2" align="left">
		&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;Днес: <b>08.04.2025</b> год., на основание чл. 109 от Закона за горите се състави настоящият
	</td>
  </tr>
</table>`

	actual, err := ExtractWhen(GetTestSelection(html))

	if err != nil {
		log.Fatal(err)
	}

	expected := time.Date(
		2025, 4, 8, 0, 0, 0, 0, parser.GetLocation())

	if actual.Compare(expected) != 0 {
		t.Errorf("The time should be %s but it is %s", expected, actual)
	}
}
