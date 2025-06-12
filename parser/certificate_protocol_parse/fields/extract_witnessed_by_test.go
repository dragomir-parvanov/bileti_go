package certificate_protocol_parse_fields

import "testing"

func TestShouldReturnWitnessedBy(t *testing.T) {
	html := `
<table>
	<tr>
	  <td align="left">Съставил: ............................................</td>
	  <td align="left">Присъствал: Драго</td>
	</tr>
</table>`

	actual := ExtractWitnessedBy(GetTestSelection(html))

	expected := "Драго"

	if actual != expected {
		t.Errorf("Expected: %s, Actual: %s", expected, actual)
	}
}

func TestShouldReturnEmptyStringOnNoWitnessedBy(t *testing.T) {
	html := `
<table>
	<tr>
	  <td align="left">Съставил: ............................................</td>
	  <td align="left">Присъствал: ..........................................</td>
	</tr>
</table>`

	actual := ExtractWitnessedBy(GetTestSelection(html))

	expected := ""

	if actual != expected {
		t.Errorf("Expected: %s, Actual: %s", expected, actual)
	}
}
