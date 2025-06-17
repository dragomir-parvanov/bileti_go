package certificate_protocol_parse_fields

import (
	"testing"
)

func TestShouldReturnRecordWitnessedBy(t *testing.T) {
	html := `
	<table>
		<tr>
		  <td colspan="2" align="left">в присъствието на: <b>Бай Пешо</b></td>
		</tr>
	</table>
	`

	actual := ExtractFindingsRecordWitnessedBy(GetTestSelection(html))
	expected := "Бай Пешо"

	if actual != expected {
		t.Errorf("Got %s, Expected %s", actual, expected)
	}
}

func TestShouldEmptyStringOnNoReturnRecordWitnessedBy(t *testing.T) {
	html := `
	<table>
		<tr>
		  <td colspan="2" align="left">в присъствието на: <b>.........</b></td>
		</tr>
	</table>
	`

	actual := ExtractFindingsRecordWitnessedBy(GetTestSelection(html))
	expected := ""

	if actual != expected {
		t.Errorf("Got %s, Expected %s", actual, expected)
	}
}
