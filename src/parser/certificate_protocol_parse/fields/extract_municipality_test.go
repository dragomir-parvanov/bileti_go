package certificate_protocol_parse_fields

import "testing"

func TestShouldReturnMunicipality(t *testing.T) {
	html := `
		<table>
			<tr>
			  <td colspan="2" align="left">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<b>1.</b> След съвместно обхождане и проверка на сечището в Община: <b>Балчик</b>;</td>
			</tr>
		</table>
`

	actual := ExtractMunicipality(GetTestSelection(html))

	expected := "Балчик"

	if actual != expected {
		t.Errorf("Expected: %s, Actual: %s", expected, actual)
	}
}
