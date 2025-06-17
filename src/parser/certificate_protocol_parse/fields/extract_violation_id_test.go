package certificate_protocol_parse_fields

import "testing"

func TestShouldReturnViolationId(t *testing.T) {
	html := `
<table>
		<tr>
			<td align="left" colspan="2">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<b>4.</b> За допуснатите нарушения е съставен акт за установяване на административно<br> 
			 нарушение № <b>1337</b> / <b>&nbsp;</b>
			 на: <b></b>
			</td>
		</tr>
</table>`

	actual := ExtractViolationId(GetTestSelection(html))

	expected := "1337"

	if actual != expected {
		t.Errorf("Expected: %s, Actual: %s", expected, actual)
	}
}

func TestShouldReturnEmptyStringOnNoViolationId(t *testing.T) {
	html := `
<table>
		<tr>
			<td align="left" colspan="2">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<b>4.</b> За допуснатите нарушения е съставен акт за установяване на административно<br> 
			 нарушение № <b></b> / <b>&nbsp;</b>
			 на: <b></b>
			</td>
		</tr>
</table>`

	actual := ExtractViolationId(GetTestSelection(html))

	expected := ""

	if actual != expected {
		t.Errorf("Expected: %s, Actual: %s", expected, actual)
	}
}
