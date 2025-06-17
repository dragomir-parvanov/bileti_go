package certificate_protocol_parse_fields

import (
	"testing"
)

func TestShouldReturnAreaDone(t *testing.T) {
	html := `
	<table>
		<tr>
		  <td align="left" colspan="2">Площ по позволително за сеч (ха): <b>0.090</b>&nbsp;&nbsp;|&nbsp;&nbsp;
		Сечта е проведена на площ (ха): <b>0.080</b>
		</td>
		</tr>
	</table>
`

	actual := ExtractAreaDone(GetTestSelection(html))

	expected := 0.08

	if actual != expected {
		t.Errorf("Should be %f but got %f", expected, actual)
	}
}
