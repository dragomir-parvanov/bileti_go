package certificate_protocol_parse_fields

import (
	"testing"
)

func TestShouldReturnAreaOfPermit(t *testing.T) {
	html := `
	<table>
		<tr>
		  <td align="left" colspan="2">Площ по позволително за сеч (ха): <b>0.090</b>&nbsp;&nbsp;|&nbsp;&nbsp;
		Сечта е проведена на площ (ха): <b>0.090</b>
		</td>
		</tr>
	</table>
`

	actual := ExtractAreaOfPermit(GetTestSelection(html))

	expected := 0.090

	if actual != expected {
		t.Errorf("Should be %f but got %f", expected, actual)
	}
}
