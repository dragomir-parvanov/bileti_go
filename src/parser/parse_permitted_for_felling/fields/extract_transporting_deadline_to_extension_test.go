package parse_permitted_for_felling_fields

import (
	"bileti_go/src/parser"
	"testing"
	"time"
)

func TestShouldReturnDeadlineToExtension(t *testing.T) {
	html := `
	<table>
		<tr>
			<td align="left"> За провеждане на сечта до : ...................... г.</td>
			<td align="left"> За извоз на материалите до : 28.02.2025 г.</td>
		</tr>
	</table>
`
	result := ExtractTransportingDeadlineToExtension(GetTestSelection(html))

	expected := time.Date(2025, 2, 28, 0, 0, 0, 0, parser.GetLocation())

	if result.Compare(expected) != 0 {
		t.Errorf("Got %s, expected %s", result, expected)
	}
}

func TestShouldReturnZeroTimeIfTransportingDeadlineToExtensionIsEmpty(t *testing.T) {
	html := `
	<table>
		<tr>
			<td align="left"> За провеждане на сечта до : ...................... г.</td>
			<td align="left"> За извоз на материалите до : ...................... г.</td>
		</tr>
	</table>
`
	result := ExtractTransportingDeadlineToExtension(GetTestSelection(html))

	expected := time.Time{}

	if result.Compare(expected) != 0 {
		t.Errorf("Got %s, expected %s", result, expected)
	}
}
