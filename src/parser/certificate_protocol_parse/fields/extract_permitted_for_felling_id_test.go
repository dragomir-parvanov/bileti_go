package certificate_protocol_parse_fields

import "testing"

func TestShouldReturnPermitId(t *testing.T) {
	html := `
		<table>
			<tr>
			  <td colspan="2" align="center">Към позволително за сеч № 0800703/08.01.2025 год.</td>
			</tr>
		</table>
	`

	actual := ExtractPermittedForFellingId(GetTestSelection(html))

	expected := "0800703"

	if actual != expected {
		t.Errorf("Received %s, expected %s", actual, expected)
	}
}
