package certificate_protocol_parse_fields

import (
	"bileti_go/parser"
	"testing"
	"time"
)

func TestShouldReturnToImproveClearingDeadline(t *testing.T) {
	html := `
	<table>
		<tr>
		  <td align="left" colspan="2">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<b>5.</b> За подобряване състоянието на насаждението в срок до: 
		  <b>12.04.2000</b>
		 следва да се проведът следните мероприятия:&nbsp;
		 <b></b></td>
		</tr>
	</table>
`

	actual, _ := ExtractToImproveClearingDeadline(GetTestSelection(html))

	expected := time.Date(2000, 4, 12, 0, 0, 0, 0, parser.GetLocation())

	if actual.Compare(expected) != 0 {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}
}

func TestShouldReturnToImproveClearingDeadlineOnBothDeadlineAndDescription(t *testing.T) {
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

	actual, _ := ExtractToImproveClearingDeadline(GetTestSelection(html))

	expected := time.Date(2000, 4, 12, 0, 0, 0, 0, parser.GetLocation())

	if actual.Compare(expected) != 0 {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}
}

func TestShouldReturnZeroTimeOnNoClearingDeadline(t *testing.T) {
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

	actual, err := ExtractToImproveClearingDeadline(GetTestSelection(html))

	if err != nil {
		t.Errorf("Did not expect an error, Actual error: <%s>", err.Error())
	}

	expected := time.Time{}

	if actual.Compare(expected) != 0 {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}
}
