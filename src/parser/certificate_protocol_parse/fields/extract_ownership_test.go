package certificate_protocol_parse_fields

import (
	"testing"
)

func TestShouldReturnOwnership(t *testing.T) {
	html := `
	<table>
		<tr>
		  <td colspan="2" align="left">собственост на: <b>ДГТ</b>, и на всички документи</td>
		</tr>
	</table>
`

	actual := ExtractOwnership(GetTestSelection(html))

	expected := "ДГТ"

	if actual != expected {
		t.Errorf("Should be %s but got %s", expected, actual)
	}
}
