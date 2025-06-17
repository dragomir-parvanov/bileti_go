package certificate_protocol_parse_fields

import (
	"testing"
)

func TestShouldReturnViolationDescription(t *testing.T) {
	html := `
	<table>
		<tr>
		  <td colspan="2" align="left">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<b>2.</b> При сечта и извоза на дървесина са допуснати следните нарушения на Закона за горите:&nbsp;
		  <b>Дарамдудай</b></td>
		</tr>
	</table>
`

	actual := ExtractViolationDescription(GetTestSelection(html))

	expected := "Дарамдудай"

	if actual != expected {
		t.Errorf("Should return '%s' but got '%s'", expected, actual)
	}
}

func TestShouldReturnEmptyStringOnNoViolationDescription(t *testing.T) {
	html := `
	<table>
		<tr>
		  <td colspan="2" align="left">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<b>2.</b> При сечта и извоза на дървесина са допуснати следните нарушения на Закона за горите:&nbsp;
		  <b></b></td>
		</tr>
	</table>
`

	actual := ExtractViolationDescription(GetTestSelection(html))

	expected := ""

	if actual != expected {
		t.Errorf("Should return '%s' but got '%s'", expected, actual)
	}

}
