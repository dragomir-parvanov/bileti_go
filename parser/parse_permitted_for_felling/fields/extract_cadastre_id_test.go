package parse_permitted_for_felling_fields

import "testing"

func TestShouldReturnCadastreId(t *testing.T) {
	html := `
	<table>
		<tr>
			<td colspan="2" align="left">да извърши добива в отдел № <b>182</b>; подотдел <b>и</b>; имот с (кадастрален/КВС) № <b>1337</b>, находящ се в: </td>
		</tr>
	</table>
`

	actual := ExtractCadastreId(GetTestSelection(html))

	expected := "1337"

	if actual != expected {
		t.Errorf("Expected %s, Received %s", actual, expected)
	}
}

func TestShouldReturnEmptyStringOnNoCadastreId(t *testing.T) {
	html := `
	<table>
		<tr>
			<td colspan="2" align="left">да извърши добива в отдел № <b>182</b>; подотдел <b>и</b>; имот с (кадастрален/КВС) № <b>&nbsp;</b>, находящ се в: </td>
		</tr>
	</table>
`

	actual := ExtractCadastreId(GetTestSelection(html))

	expected := ""

	if actual != expected {
		t.Errorf("Expected %s, Received %s", actual, expected)
	}
}

func TestShouldReturnCadastreIdWhenThereIsInventoryInformation(t *testing.T) {
	html := `
	<table>
		<tr>
			<td colspan="2" align="left">да извърши добива в отдел № <b>256</b>; подотдел <b>з</b> съгласно инвентаризацията на <b>ДГС София</b>, утвърдена със <br><br>заповед № <b>1024/2019-12-19</b>г.; имот с (кадастрален/КВС) № <b>1337</b>, находящ се в: </td>
		</tr>
	</table>
`

	actual := ExtractCadastreId(GetTestSelection(html))

	expected := "1337"

	if actual != expected {
		t.Errorf("Expected %s, Received %s", expected, actual)
	}
}

func TestShouldReturnEmptyStringWhenNoCadastreIdAndWhenThereIsInventoryInformation(t *testing.T) {
	html := `
	<table>
		<tr>
			<td colspan="2" align="left">да извърши добива в отдел № <b>256</b>; подотдел <b>з</b> съгласно инвентаризацията на <b>ДГС София</b>, утвърдена със <br><br>заповед № <b>1024/2019-12-19</b>г.; имот с (кадастрален/КВС) № <b>&nbsp;</b>, находящ се в: </td>
		</tr>
	</table>
`

	actual := ExtractCadastreId(GetTestSelection(html))

	expected := ""

	if actual != expected {
		t.Errorf("Expected %s, Received %s", expected, actual)
	}
}
