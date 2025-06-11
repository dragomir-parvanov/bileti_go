package certificate_protocol_parse_fields

import (
	"testing"
)

func TestShouldReturnLoggingFinishDescription(t *testing.T) {
	html := `
	<table>
		<tr>
		  <td align="left" colspan="2">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<b>3.</b> Провеждането на сечта и почистването на сечището е извършено според определения начин в технологичния план и позволителното за сеч:
		<b>Задоволително</b>
		</td></tr>
	</table>
`

	actual := ExtractLoggingFinishDescription(GetTestSelection(html))

	expected := "Задоволително"

	if actual != expected {
		t.Errorf("Should return '%s' but got '%s'", expected, actual)
	}
}
