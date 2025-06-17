package certificate_protocol_parse_fields

import (
	"testing"
)

func TestShouldReturnToImproveClearingDescription(t *testing.T) {
	html := `
	<table>
		<tr>
		  <td align="left" colspan="2">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<b>5.</b> За подобряване състоянието на насаждението в срок до: 
		  <b></b>
		 следва да се проведът следните мероприятия:&nbsp;
		 <b>Дарамдудай</b></td>
		</tr>
	</table>
`

	actual := ExtractToImproveClearingDescription(GetTestSelection(html))

	expected := "Дарамдудай"

	if actual != expected {
		t.Errorf("Expected: %s, Actual: %s", expected, actual)
	}
}

func TestShouldReturnToImproveClearingDescriptionOnBothDeadlineAndDescription(t *testing.T) {
	html := `
	<table>
		<tr>
		  <td align="left" colspan="2">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<b>5.</b> За подобряване състоянието на насаждението в срок до: 
		  <b>12.04.2000</b>
		 следва да се проведът следните мероприятия:&nbsp;
		 <b>Дарамдудай</b></td>
		</tr>
	</table>
`

	actual := ExtractToImproveClearingDescription(GetTestSelection(html))

	expected := "Дарамдудай"

	if actual != expected {
		t.Errorf("Expected: %s, Actual: %s", expected, actual)
	}
}

func TestShouldReturnEmptyStringOnNoToImproveClearingDescription(t *testing.T) {
	html := `
	<table>
		<tr>
		  <td align="left" colspan="2">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<b>5.</b> За подобряване състоянието на насаждението в срок до: 
		  <b>&nbsp;</b>
		 следва да се проведът следните мероприятия:&nbsp;
		 <b></b></td>
		</tr>
	</table>
`

	actual := ExtractToImproveClearingDescription(GetTestSelection(html))

	expected := ""

	if actual != expected {
		t.Errorf("Expected: %s, Actual: %s", expected, actual)
	}
}
