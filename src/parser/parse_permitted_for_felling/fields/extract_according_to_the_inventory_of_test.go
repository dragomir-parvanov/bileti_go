package parse_permitted_for_felling_fields

import "testing"

func TestShouldReturnAccordingToTheInventoryOf(t *testing.T) {
	html := `
	<table>
			<tr>
				<td colspan="2" align="left">да извърши добива в отдел № <b>256</b>; подотдел <b>з</b> съгласно инвентаризацията на <b>ДГС София</b>, утвърдена със <br><br>заповед № <b>1024/2019-12-19</b>г.; имот с (кадастрален/КВС) № <b>68134.4129.537</b>, находящ се в: </td>
			</tr>
	</table>
	`

	actual := ExtractAccordingToTheInventoryOf(GetTestSelection(html))

	expected := "ДГС София"

	if actual != expected {
		t.Errorf("Expected: %s, Actual: %s", expected, actual)
	}
}

func TestShouldReturnEmptyStringOnNoAccordingToTheInventoryOf(t *testing.T) {
	html := `
	<table>
			<tr>
				<td colspan="2" align="left">да извърши добива в отдел № <b>1118</b>; подотдел <b>6</b>; имот с (кадастрален/КВС) № <b>&nbsp;</b>, находящ се в: </td>
			</tr>
	</table>
	`

	actual := ExtractAccordingToTheInventoryOf(GetTestSelection(html))

	expected := ""

	if actual != expected {
		t.Errorf("Expected: %s, Actual: %s", expected, actual)
	}
}
