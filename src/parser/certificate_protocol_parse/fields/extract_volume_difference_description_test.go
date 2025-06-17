package certificate_protocol_parse_fields

import (
	"testing"
)

func TestShouldReturnVolumeDifferenceDescription(t *testing.T) {
	html := `
	<table>
		<tr>
		  <td colspan="2" align="left">Разликите между количествата и категориите дървесина, по позволително за сеч и фактически добити в повече или в по-малко се дължът на следното:&nbsp;
		<b>Сечта е изведена по КП на РДГ Варна съставен за установяване на опасни дървета в имот 02508.62.38.Дървесината е негодна за употреба</b></td>
		</tr>
	</table>
`

	actual := ExtractVolumeDifferenceDescription(GetTestSelection(html))

	expected := "Сечта е изведена по КП на РДГ Варна съставен за установяване на опасни дървета в имот 02508.62.38.Дървесината е негодна за употреба"

	if actual != expected {
		t.Errorf("Should return section '%s' but got '%s'", expected, actual)
	}
}
