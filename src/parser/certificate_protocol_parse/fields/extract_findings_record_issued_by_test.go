package certificate_protocol_parse_fields

import "testing"

func TestShouldReturnEmptyStringIfNoFindingsRecordIssuedBy(t *testing.T) {
	html := `
	<table>
		<tr>
		  <td colspan="2" align="left">констативен протокол от: <b>...........</b></td>
		</tr>
	</table>
	`

	actual := ExtractFindingsRecordIssuedBy(GetTestSelection(html))

	expected := ""

	if actual != expected {
		t.Errorf(`Expected "%s", got "%s"`, expected, actual)
	}
}

func TestShouldReturnFindingsRecordIssuedBy(t *testing.T) {
	html := `
	<table>
		<tr>
		  <td colspan="2" align="left">констативен протокол от: <b>Бай Иван</b></td>
		</tr>
	</table>
	`

	actual := ExtractFindingsRecordIssuedBy(GetTestSelection(html))

	expected := "Бай Иван"

	if actual != expected {
		t.Errorf(`Expected "%s", got "%s"`, expected, actual)
	}
}
