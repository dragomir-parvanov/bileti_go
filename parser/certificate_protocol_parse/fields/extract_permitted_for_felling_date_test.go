package certificate_protocol_parse_fields

import (
	"bileti_go/parser"
	"testing"
	"time"
)

func TestShouldReturnPermittedForFellingDate(t *testing.T) {
	html := `
	<table>
		<tr>
		  <td colspan="2" align="center">Към позволително за сеч № 0800703/08.01.2025 год.</td>
		</tr>
	</table>
`

	actual, _ := ExtractPermittedForFellingDate(GetTestSelection(html))

	expected := time.Date(2025, 1, 8, 0, 0, 0, 0, parser.GetLocation())

	if actual.Compare(expected) != 0 {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}
}
