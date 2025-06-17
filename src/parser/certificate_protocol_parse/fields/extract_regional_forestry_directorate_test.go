package certificate_protocol_parse_fields

import (
	"testing"
)

func TestShouldReturnRegionalForestryDirectorate(t *testing.T) {

	html := `
	<tbody>
		<tr>
			<td rowspan="2" align="center" width="80" height="53"><img src="/images/logo.png"></td>
			<td align="center"><font size="+1">ИЗПЪЛНИТЕЛНА АГЕНЦИЯ ПО ГОРИТЕ</font></td>
		</tr>
		<tr>
			<td colspan="2" align="center"><font size="2px">РЕГИОНАЛНА ДИРЕНЦИЯ ПО ГОРИТЕ <b>Варна</b></font></td>
		</tr>
		<tr>
			<td colspan="2" align="center"><hr></td>
		</tr>
	</tbody>`

	actual := ExtractRegionalForestryDirectorate(GetTestSelection(html))

	expected := "Варна"

	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}
