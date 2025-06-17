package certificate_protocol_parse_fields

import "testing"

func TestShouldReturnId(t *testing.T) {
	html := `
<table>
	<tr>
		<td colspan="2" align="center"><br><font size="2"><i><b>ПРОТОКОЛ за освидетелстване на сечище</b></i> № <b>0800519</b></font></td>
	</tr>
</table>`

	actual := ExtractId(GetTestSelection(html))

	expected := "0800519"

	if actual != expected {
		t.Errorf("Expected: %s, Actual: %s", expected, actual)
	}
}
