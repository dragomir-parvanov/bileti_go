package parse_permitted_for_felling_fields

import (
	"bileti_go/src/parser"
	"testing"
	"time"
)

func TestShouldReturnDateOfCarnetInventory(t *testing.T) {
	html := `
	<table>
		<tr>
			<td colspan="2" align="left"> с контролна горска марка № <b>Б8813</b>, с <b>черна</b> боя, дата на карнет опис: <b>21.02.2025</b> г.</td>
		</tr>
	</table>
`
	result := ExtractDateOfCarnetInventory(GetTestSelection(html))

	expected := time.Date(2025, 2, 21, 0, 0, 0, 0, parser.GetLocation())

	if result.Compare(expected) != 0 {
		t.Errorf("Got %s, expected %s", result, expected)
	}
}

func TestShouldReturnZeroTimeIfDateOfCarnetInventoryIsEmpty(t *testing.T) {
	html := `
	<table>
		<tr>
			<td colspan="2" align="left"> с контролна горска марка № <b>Б8813</b>, с <b>черна</b> боя, дата на карнет опис: <b>&nbsp;</b> г.</td>
		</tr>
	</table>
`
	result := ExtractDateOfCarnetInventory(GetTestSelection(html))

	expected := time.Time{}

	if result.Compare(expected) != 0 {
		t.Errorf("Got %s, expected %s", result, expected)
	}
}
