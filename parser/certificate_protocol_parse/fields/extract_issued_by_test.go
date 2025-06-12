package certificate_protocol_parse_fields

import "testing"

func TestShouldReturnIssuedBy(t *testing.T) {
	html := `
<table>
	<tr>
	  <td align="left">Съставил: Драго</td>
	  <td align="left">Присъствал: ..........................................</td>
	</tr>
</table>`

	actual := ExtractIssuedBy(GetTestSelection(html))

	expected := "Драго"

	if actual != expected {
		t.Errorf("Expected: %s, Actual: %s", expected, actual)
	}
}

func TestShouldReturnEmptyStringOnNoIssuedBy(t *testing.T) {
	html := `
<table>
	<tr>
	  <td align="left">Съставил: ............................................</td>
	  <td align="left">Присъствал: ..........................................</td>
	</tr>
</table>`

	actual := ExtractIssuedBy(GetTestSelection(html))

	expected := ""

	if actual != expected {
		t.Errorf("Expected: %s, Actual: %s", expected, actual)
	}
}
