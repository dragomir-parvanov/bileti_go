package certificate_protocol_parse_fields

import (
	"testing"
)

func TestShouldReturnFalseOnNe(t *testing.T) {
	html := `
	<table>
		<tr>
		  <td align="left" colspan="2">Налични неотсечени маркирани стъбла : 
		<b>Не</b></td>
		</tr>
	</table>
`

	actual := ExtractNotMarkedLoggedTreesExist(GetTestSelection(html))

	expected := false

	if actual != expected {
		t.Errorf("Should return section '%v' but got '%v'", expected, actual)
	}
}

func TestShouldReturnTrueOnDa(t *testing.T) {
	html := `
	<table>
		<tr>
		  <td align="left" colspan="2">Налични неотсечени маркирани стъбла : 
		<b>Да</b></td>
		</tr>
	</table>
`

	actual := ExtractNotMarkedLoggedTreesExist(GetTestSelection(html))

	expected := true

	if actual != expected {
		t.Errorf("Should return section '%v' but got '%v'", expected, actual)
	}
}
