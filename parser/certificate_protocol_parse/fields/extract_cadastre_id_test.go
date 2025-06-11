package certificate_protocol_parse_fields

import (
	"testing"
)

func TestShouldReturnCadastreId(t *testing.T) {
	html := `
	<table>
		<tr>
		  <td colspan="2" align="left">землище: <b>Балчик</b>; Отдел: <b>53</b>; Подотдел: <b>10</b>; Кад. № <b>02508.62.38</b>;</td>
		</tr>
	</table>
`

	actual := ExtractCadastreId(GetTestSelection(html))

	expected := "02508.62.38"

	if actual != expected {
		t.Errorf("Should return section '%s' but got '%s'", expected, actual)
	}
}
