package certificate_protocol_parse_fields

import (
	"testing"
)

func TestShouldReturnViolationRecipient(t *testing.T) {
	html := `
<table>
		<tr>
			<td align="left" colspan="2">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<b>4.</b> За допуснатите нарушения е съставен акт за установяване на административно<br> 
			 нарушение № <b></b> / <b>&nbsp;</b>
			 на: <b>Иван Петков</b>
			</td>
		</tr>
</table>
`

	actual := ExtractViolationRecipient(GetTestSelection(html))

	expected := "Иван Петков"

	if actual != expected {
		t.Errorf("Should return '%s' but got '%s'", expected, actual)
	}
}

func TestShouldReturnViolationRecipientOnBothIdAndRecipient(t *testing.T) {
	html := `
<table>
		<tr>
			<td align="left" colspan="2">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<b>4.</b> За допуснатите нарушения е съставен акт за установяване на административно<br> 
			 нарушение № <b>1337</b> / <b>&nbsp;</b>
			 на: <b>Иван Петков</b>
			</td>
		</tr>
</table>
`

	actual := ExtractViolationRecipient(GetTestSelection(html))

	expected := "Иван Петков"

	if actual != expected {
		t.Errorf("Should return '%s' but got '%s'", expected, actual)
	}
}

func TestShouldReturnEmptyStringOnNoViolationRecipient(t *testing.T) {
	html := `
<table>
		<tr>
			<td align="left" colspan="2">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<b>4.</b> За допуснатите нарушения е съставен акт за установяване на административно<br> 
			 нарушение № <b></b> / <b>&nbsp;</b>
			 на: <b></b>
			</td>
		</tr>
</table>
`

	actual := ExtractViolationRecipient(GetTestSelection(html))

	expected := ""

	if actual != expected {
		t.Errorf("Should return '%s' but got '%s'", expected, actual)
	}

}
