package certificate_protocol_parse

import (
	"bileti_go/src/parser/certificate_protocol_parse/fields"
	"bileti_go/src/parser/extraction_category"
	"reflect"
	"testing"
)

const treeExtractionHtml = `
<table align="center" width="640" border="0" cellpadding="0" cellspacing="0" class="ex">
		<tbody><tr>
			<td rowspan="2" align="center" width="20">№</td>
			<td rowspan="2" align="center">Категория дървесина</td>
			<td rowspan="2" align="center">Дървесен вид</td>
			<td colspan="4" align="center">Лежаща маса в куб. м.</td>
		</tr>
		<tr>
			<td align="center">По позволително <br>за сеч</td>
			<td align="center">Действително <br>отсечено</td>
			<td align="center">Наличност в <br>сечището</td>
			<td align="center">Налично на <br>временен склад</td>
		</tr>
		<tr>
			<td align="center">1.</td>
			<td align="left">&nbsp;Едра</td>
			<td align="right">---</td>
			<td align="right">---</td>
			<td align="right">---</td>
			<td align="right">---</td>
			<td align="right">---</td>
		</tr>
		<tr>
			<td align="center">2.</td>
			<td align="left">&nbsp;Средна</td>
			<td align="right">---</td>
			<td align="right">---</td>
			<td align="right">---</td>
			<td align="right">---</td>
			<td align="right">---</td>
		</tr>
		<tr>
			<td align="center">3.</td>
			<td align="left">&nbsp;Дребна</td>
			<td align="right">---</td>
			<td align="right">---</td>
			<td align="right">---</td>
			<td align="right">---</td>
			<td align="right">---</td>
		</tr>
		<tr>
			<td rowspan="4" align="center">4.</td>
			<td rowspan="4">&nbsp;Дърва</td>
			<td align="right"><b>Ясен</b></td>
			<td align="right"><b>&nbsp;</b></td>
			<td align="right"><b>&nbsp;</b></td>
			<td align="right"><b>&nbsp;</b></td>
			<td align="right"><b>&nbsp;</b></td>
		</tr>
		<tr>
			<td align="right"><b>Други широколистни</b></td>
			<td align="right"><b>0.1</b></td>
			<td align="right"><b>&nbsp;</b></td>
			<td align="right"><b>&nbsp;</b></td>
			<td align="right"><b>&nbsp;</b></td>
		</tr>
		<tr>
			<td align="right"><b>Айлант</b></td>
			<td align="right"><b>0.1</b></td>
			<td align="right"><b>&nbsp;</b></td>
			<td align="right"><b>&nbsp;</b></td>
			<td align="right"><b>&nbsp;</b></td>
		</tr>
		<tr>
			<td align="right"><b>Джанка</b></td>
			<td align="right"><b>0.3</b></td>
			<td align="right"><b>&nbsp;</b></td>
			<td align="right"><b>&nbsp;</b></td>
			<td align="right"><b>&nbsp;</b></td>
		</tr>
		<tr>
			<td align="center">5.</td>
			<td align="left">&nbsp;Вършина</td>
			<td align="right">---</td>
			<td align="right">---</td>
			<td align="right">---</td>
			<td align="right">---</td>
			<td align="right">---</td>
		</tr>
		<tr>
			<td colspan="3" align="right"><b>Общо количество:</b></td>
			<td align="right"><b>0.5</b></td>
			<td align="right"><b>&nbsp;</b></td>
			<td align="right"><b>&nbsp;</b></td>
			<td align="right"><b>&nbsp;</b></td>
		</tr>
	</tbody></table>
`

func TestShouldReturnTreeExtraction(t *testing.T) {

	actual := ExtractExtraction(certificate_protocol_parse_fields.GetTestSelection(treeExtractionHtml))

	expected := []CertificateProtocolTreeExtraction{
		{category: parser_extraction_category.Wood, treeType: "Ясен", byLoggingPermitCollected: 0, actuallyCollected: 0, availableInTheClearing: 0, availableInTemporaryStorage: 0},
		{category: parser_extraction_category.Wood, treeType: "Други широколистни", byLoggingPermitCollected: 0.1, actuallyCollected: 0, availableInTheClearing: 0, availableInTemporaryStorage: 0},
		{category: parser_extraction_category.Wood, treeType: "Айлант", byLoggingPermitCollected: 0.1, actuallyCollected: 0, availableInTheClearing: 0, availableInTemporaryStorage: 0},
		{category: parser_extraction_category.Wood, treeType: "Джанка", byLoggingPermitCollected: 0.3, actuallyCollected: 0, availableInTheClearing: 0, availableInTemporaryStorage: 0},
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected: %v\nActual: %v", expected, actual)
	}

}
