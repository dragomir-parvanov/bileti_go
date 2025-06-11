package parser

import (
	"strings"
	"testing"
	"time"
)

const html = `
<body data-new-gr-c-s-check-loaded="14.1237.0" data-gr-ext-installed="">
<table width="650" align="center" border="0" summary="" class="sech">
    <tbody><tr>
        <td rowspan="2" align="center" width="80" height="53"><img src="/images/logo.png"></td>
        <td align="center"><font size="+1">ИЗПЪЛНИТЕЛНА АГЕНЦИЯ ПО ГОРИТЕ</font></td>
    </tr>
    <tr>
        <td colspan="2" align="center"><font size="2px">Регионална дирекция по горите <b>Бургас</b></font></td>
    </tr>
    <tr>
        <td colspan="2" align="center"><hr></td>
    </tr>
    </tbody></table>
<table border="0" align="center" width="650" cellpadding="0" cellspacing="8" class="sech">
    <tbody><tr>
        <td colspan="2" align="left">ВИД НА СЕЧТА : <b>В сервитути&nbsp;</b>
        </td>
    </tr>
    <tr>
        <td colspan="2" align="center"><br><font size="2"><b>ПОЗВОЛИТЕЛНО за СЕЧ № 0805138</b>
        </font></td>
    </tr>
    <tr>
        <td colspan="2" align="left"><br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;На основание чл. 108 от Закона за горите, Заповед № <b>127/18.02.2021</b> г. за утвърждаване на
        </td>
    </tr>
    <tr>
        <td colspan="2" align="left">горскостопански план (горскостопанска програма)</td>
    </tr>
    <tr>
        <td colspan="2" align="left">разрешава се на <b>ЕЛЕКТРОЕНЕРГИЕН СИСТЕМЕН ОПЕРАТОР  ЕАД</b></td>
    </tr>
    <tr>
        <td colspan="2" align="left">с нает регистриран лесовъд &nbsp;<b>Регистриран лесовъд</b> , </td>
    </tr>
    <tr>
        <td colspan="2" align="left">да извърши добива в отдел № <b>450</b>; подотдел <b>1</b>; имот с (кадастрален/КВС) № <b>тестов кадастрален номер</b>, находящ се в: </td>
    </tr>
    <tr>
        <td colspan="2" align="left">Община <b>Айтос</b>; Землище <b>Айтос землище</b> площно сечище от <b>2.200</b> хектара; Вид собственост : <b>ДГТ</b></td>
    </tr>
    <tr>
        <td colspan="2" align="left"> Дърветата са маркирани от <b>Тестов Маркировач</b></td>
    </tr>
    <tr>
        <td colspan="2" align="left"> с контролна горска марка № <b>Б 1401</b>, с <b>оранжева</b> боя, дата на карнет опис: <b>&nbsp;</b> г.</td>
    </tr>
    <tr>
        <td colspan="2" align="left"> Очакваният добив е <b>1.0</b> плътни кубически метра лежаща маса, която по категории е както следва :</td>
    </tr>
    </tbody></table>
<table align="center" width="640" border="0" cellpadding="0" cellspacing="0" class="ex">
    <tbody><tr>
        <td align="center" width="20">№</td>
        <td align="center" width="200">Категория дървесина</td>
        <td align="center" width="90">Дървесен вид</td>
        <td align="center" width="50">куб. м. </td>
        <td align="center">Забележка </td>
    </tr>
    <tr>
        <td rowspan="1" align="center">1.</td>
        <td rowspan="1">&nbsp;Едра строителна дървесина</td>
        <td><b> --- </b></td>
        <td align="right"><b> --- </b></td>
        <td align="center"><font size="1"></font></td>
    </tr>
    <tr>
        <td rowspan="1" align="center">2.</td>
        <td rowspan="1">&nbsp;Средна строителна дървесина</td>
        <td><b> --- </b></td>
        <td align="right"><b> --- </b></td>
        <td align="center"><font size="1"></font></td>
    </tr>
    <tr>
        <td rowspan="1" align="center">3.</td>
        <td rowspan="1">&nbsp;Дребна строителна дървесина</td>
        <td><b> --- </b></td>
        <td align="right"><b> --- </b></td>
        <td align="center"><font size="1"></font></td>
    </tr>
    <tr>
        <td rowspan="1" align="center">4.</td>
        <td rowspan="1">&nbsp;Дърва</td>
        <td><b> --- </b></td>
        <td align="right"><b> --- </b></td>
        <td align="center"><font size="1"></font></td>
    </tr>
    <tr>
        <td rowspan="1" align="center">5.</td>
        <td rowspan="1">&nbsp;Вършина</td>
        <td><b>Келяв габър</b></td>
        <td align="right"><b>1.0</b></td>
        <td align="center"><font size="1"></font></td>
    </tr>
    </tbody></table>
<table border="0" align="center" width="650" cellpadding="0" cellspacing="8" class="sech">
    <tbody><tr>
        <td colspan="2" align="left"> Допълнителни изисквания при провеждане на сечта : <b>Позволителното за сеч се издава за почистване в сервитут на електропровод без материален добив.</b></td>
    </tr>
    <tr>
        <td colspan="2" align="left"> Срок за провеждане на сечта от <b>27.01.2025</b>&nbsp;г. до <b>31.12.2025</b>&nbsp;г. </td>
    </tr>
    <tr>
        <td colspan="2" align="left"> Срок за извозване на материалите от сечището от <b>27.01.2025</b>&nbsp;г. до <b>31.12.2025</b>&nbsp;г. </td>
    </tr>
    <tr>
        <td colspan="2" align="left"> Начин на почистване на сечището : <b>Тестово почистване</b></td>
    </tr>
    <tr>
        <td colspan="2" align="left"> Материалите ще се извозят до временен склад : <b>съгласно ТП</b></td>
    </tr>
    <tr>
        <td colspan="2">&nbsp;</td>
    </tr>
    <tr>
        <td align="left" width="315">Издал:Издал човек</td>
        <td align="left" width="335">Получил позволителното:Получил човек</td>
    </tr>
    <tr>
        <td align="left" width="325">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;/<b> .......... </b>/</td>
        <td align="left" width="325">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
            &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;/<b> .......... </b>/</td>
    </tr>
    <tr>
        <td colspan="2" align="left"> Дата: <b>24.01.2025</b> г.
            &nbsp;&nbsp;Издал служител : |Служител|
            &nbsp;&nbsp;Код: |Някакъв код|</td>
    </tr>
    <tr>
        <td>Място на издаване на позволителното за сеч:</td>	</tr>
    <tr>
        <td colspan="2" align="left">Област <b>Бургас</b>, община <b>Айтос</b>, землище <b>Айтос</b>, адрес Някакъв адрес, подотдел Някакъв подотдел, GPS координати: Някакви кординати</td>
    </tr>
    <tr>
        <td colspan="2"><b><b></b></b></td>
    </tr>
    <tr>
        <td colspan="2" align="center"><b>ЗАВЕРКИ ПРИ ПРОДЪЛЖАВАНЕ СРОКОВЕТЕ ЗА СЕЧ И ИЗВОЗ</b></td>
    </tr>
    <tr>
        <td align="left"> За провеждане на сечта до : 27.02.2025 г.</td>
        <td align="left"> За извоз на материалите до : 28.02.2025 г.</td>
    </tr>
    <tr>
        <td colspan="2" align="left">Издал:  Удължителен Човек</td>
    </tr>
    <tr>
        <td colspan="2" align="left">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;/
            &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;/</td>
    </tr>
    <tr>
        <td colspan="2" align="left">Продължаване на сроковете за сеч и извоз става само чрез информационната система на ИАГ.</td>
    </tr>
    </tbody></table>
</body>
`

func TestShouldReturnId(t *testing.T) {
	result := ParsePermittedForFellingHTML(strings.NewReader(html)).id

	expected := "0805138"

	if result != expected {
		t.Errorf("Got %s, expected %s", result, expected)
	}
}

func TestShouldReturnRegionalForestryDirectorate(t *testing.T) {
	result := ParsePermittedForFellingHTML(strings.NewReader(html)).regionalForestryDirectorate

	expected := "Бургас"

	if result != expected {
		t.Errorf("Got %s, expected %s", result, expected)
	}
}

func TestShouldReturnDirectiveNumber(t *testing.T) {
	result := ParsePermittedForFellingHTML(strings.NewReader(html)).directiveNumber

	expected := "127/18.02.2021"

	if result != expected {
		t.Errorf("Got %s, expected %s", result, expected)
	}
}

func TestShouldReturnPermittedFor(t *testing.T) {
	result := ParsePermittedForFellingHTML(strings.NewReader(html)).permittedFor

	expected := "ЕЛЕКТРОЕНЕРГИЕН СИСТЕМЕН ОПЕРАТОР ЕАД"

	if result != expected {
		t.Errorf("Got %s, expected %s", result, expected)
	}
}

func TestShouldReturnAllowedForester(t *testing.T) {
	result := ParsePermittedForFellingHTML(strings.NewReader(html)).allowedForester

	expected := "Регистриран лесовъд"

	if result != expected {
		t.Errorf("Got %s, expected %s", result, expected)
	}
}

func TestShouldReturnTypeOfFelling(t *testing.T) {
	result := ParsePermittedForFellingHTML(strings.NewReader(html)).typeOfFelling

	expected := "В сервитути"

	if result != expected {
		t.Errorf("Got %s, expected %s", result, expected)
	}
}

func TestShouldReturnSection(t *testing.T) {
	result := ParsePermittedForFellingHTML(strings.NewReader(html)).section

	expected := "450"

	if result != expected {
		t.Errorf("Got %s, expected %s", result, expected)
	}
}

func TestShouldReturnSubSection(t *testing.T) {
	result := ParsePermittedForFellingHTML(strings.NewReader(html)).subSection

	expected := "1"

	if result != expected {
		t.Errorf("Got %s, expected %s", result, expected)
	}
}

func TestShouldReturnCadastreId(t *testing.T) {
	result := ParsePermittedForFellingHTML(strings.NewReader(html)).cadastreId

	expected := "тестов кадастрален номер"

	if result != expected {
		t.Errorf("Got %s, expected %s", result, expected)
	}
}

func TestShouldReturnMunicipality(t *testing.T) {
	result := ParsePermittedForFellingHTML(strings.NewReader(html)).municipality

	expected := "Айтос"

	if result != expected {
		t.Errorf("Got %s, expected %s", result, expected)
	}
}

func TestShouldReturnLand(t *testing.T) {
	result := ParsePermittedForFellingHTML(strings.NewReader(html)).land

	expected := "Айтос землище"

	if result != expected {
		t.Errorf("Got %s, expected %s", result, expected)
	}
}

func TestShouldReturnAreaClearing(t *testing.T) {
	result := ParsePermittedForFellingHTML(strings.NewReader(html)).areaClearing

	expected := 2.2

	if result != expected {
		t.Errorf("Got %f, expected %f", result, expected)
	}
}

func TestShouldReturnOwnershipType(t *testing.T) {
	result := ParsePermittedForFellingHTML(strings.NewReader(html)).ownershipType

	expected := "ДГТ"

	if result != expected {
		t.Errorf("Got %s, expected %s", result, expected)
	}
}

func TestShouldReturnTreesMarkedBy(t *testing.T) {
	result := ParsePermittedForFellingHTML(strings.NewReader(html)).treesMarkedBy

	expected := "Тестов Маркировач"

	if result != expected {
		t.Errorf("Got %s, expected %s", result, expected)
	}
}

func TestShouldReturnControlMarkNumber(t *testing.T) {
	result := ParsePermittedForFellingHTML(strings.NewReader(html)).controlMarkNumber

	expected := "Б 1401"

	if result != expected {
		t.Errorf("Got %s, expected %s", result, expected)
	}
}

func TestShouldReturnMarkColor(t *testing.T) {
	result := ParsePermittedForFellingHTML(strings.NewReader(html)).controlMarkColor

	expected := "оранжева"

	if result != expected {
		t.Errorf("Got %s, expected %s", result, expected)
	}
}

func TestShouldReturnDateOfCarnetInventory(t *testing.T) {
	result := ParsePermittedForFellingHTML(strings.NewReader(html)).dateOfCarnetInventory

	expected := time.Time{}

	if result != expected {
		t.Errorf("Got %s, expected %s", result, expected)
	}
}

func TestShouldReturnExpectedTreeExtraction(t *testing.T) {
	result := ParsePermittedForFellingHTML(strings.NewReader(html)).expectedTreeExtraction

	expected := 1.0

	if result != expected {
		t.Errorf("Got %f, expected %f", result, expected)
	}
}

func TestShouldReturnAdditionalRequirements(t *testing.T) {
	result := ParsePermittedForFellingHTML(strings.NewReader(html)).additionalRequirements

	expected := "Позволителното за сеч се издава за почистване в сервитут на електропровод без материален добив."

	if result != expected {
		t.Errorf("Got %s, expected %s", result, expected)
	}

}

func TestShouldReturnDeadlineLoggingFrom(t *testing.T) {
	result := ParsePermittedForFellingHTML(strings.NewReader(html)).deadlineLogging.From

	expected := time.Date(2025, 1, 27, 0, 0, 0, 0, GetLocation())

	if result.Compare(expected) != 0 {
		t.Errorf("Got %s, expected %s", result, expected)
	}
}

func TestShouldReturnDeadlineLoggingTo(t *testing.T) {
	result := ParsePermittedForFellingHTML(strings.NewReader(html)).deadlineLogging.To

	expected := time.Date(2025, 12, 31, 0, 0, 0, 0, GetLocation())

	if result.Compare(expected) != 0 {
		t.Errorf("Got %s, expected %s", result, expected)
	}
}

func TestShouldReturnDeadlineMaterialsUsageFrom(t *testing.T) {
	result := ParsePermittedForFellingHTML(strings.NewReader(html)).deadlineMaterialsUsage.From

	expected := time.Date(2025, 1, 27, 0, 0, 0, 0, GetLocation())

	if result.Compare(expected) != 0 {
		t.Errorf("Got %s, expected %s", result, expected)
	}
}

func TestShouldReturnDeadlineMaterialsUsageTo(t *testing.T) {
	result := ParsePermittedForFellingHTML(strings.NewReader(html)).deadlineMaterialsUsage.To

	expected := time.Date(2025, 12, 31, 0, 0, 0, 0, GetLocation())

	if result.Compare(expected) != 0 {
		t.Errorf("Got %s, expected %s", result, expected)
	}
}

func TestShouldReturnCleaningProcedure(t *testing.T) {
	result := ParsePermittedForFellingHTML(strings.NewReader(html)).cleaningProcedure

	expected := "Тестово почистване"

	if result != expected {
		t.Errorf("Got %s, expected %s", result, expected)
	}
}

func TestShouldReturnRemovalToTemporaryStorage(t *testing.T) {
	result := ParsePermittedForFellingHTML(strings.NewReader(html)).removalToTemporaryStorage

	expected := "съгласно ТП"

	if result != expected {
		t.Errorf("Got %s, expected %s", result, expected)
	}
}

func TestShouldReturnIssuedBy(t *testing.T) {
	result := ParsePermittedForFellingHTML(strings.NewReader(html)).issuedBy

	expected := "Издал човек"

	if result != expected {
		t.Errorf("Got %s, expected %s", result, expected)
	}
}

func TestShouldReturnWhoReceivedThePermit(t *testing.T) {
	result := ParsePermittedForFellingHTML(strings.NewReader(html)).whoReceivedThePermit

	expected := "Получил човек"

	if result != expected {
		t.Errorf("Got %s, expected %s", result, expected)
	}
}

func TestShouldReturnIssuedOn(t *testing.T) {
	result := ParsePermittedForFellingHTML(strings.NewReader(html)).issuedOn

	expected := time.Date(2025, 1, 24, 0, 0, 0, 0, GetLocation())

	if result.Compare(expected) != 0 {
		t.Errorf("Got %s, expected %s", result, expected)
	}
}

func TestShouldReturnIssuedByEmployee(t *testing.T) {
	result := ParsePermittedForFellingHTML(strings.NewReader(html)).issuedByEmployee

	expected := "Служител"

	if result != expected {
		t.Errorf("Got %s, expected %s", result, expected)
	}
}

func TestShouldReturnIssuedCode(t *testing.T) {
	result := ParsePermittedForFellingHTML(strings.NewReader(html)).issuedCode

	expected := "Някакъв код"

	if result != expected {
		t.Errorf("Got %s, expected %s", result, expected)
	}
}

func TestShouldReturnPermitIssueRegion(t *testing.T) {

	result := ParsePermittedForFellingHTML(strings.NewReader(html)).permitIssuePlace.region

	expected := "Бургас"

	if result != expected {
		t.Errorf("Got %s, expected %s", result, expected)
	}
}

func TestShouldReturnPermitIssueMunicipality(t *testing.T) {
	result := ParsePermittedForFellingHTML(strings.NewReader(html)).permitIssuePlace.municipality

	expected := "Айтос"

	if result != expected {
		t.Errorf("Got %s, expected %s", result, expected)
	}
}

func TestShouldReturnPermitIssueLand(t *testing.T) {
	result := ParsePermittedForFellingHTML(strings.NewReader(html)).permitIssuePlace.land

	expected := "Айтос"

	if result != expected {
		t.Errorf("Got %s, expected %s", result, expected)
	}
}

func TestShouldReturnPermitIssueAddress(t *testing.T) {
	result := ParsePermittedForFellingHTML(strings.NewReader(html)).permitIssuePlace.address

	expected := "Някакъв адрес"

	if result != expected {
		t.Errorf("Got %s, expected %s", result, expected)
	}
}

func TestShouldReturnPermitIssueSubsection(t *testing.T) {
	result := ParsePermittedForFellingHTML(strings.NewReader(html)).permitIssuePlace.subSection

	expected := "Някакъв подотдел"

	if result != expected {
		t.Errorf("Got %s, expected %s", result, expected)
	}
}

func TestShouldReturnPermitIssueGpsCoordinates(t *testing.T) {
	result := ParsePermittedForFellingHTML(strings.NewReader(html)).permitIssuePlace.gpsCoordinates

	expected := "Някакви кординати"

	if result != expected {
		t.Errorf("Got %s, expected %s", result, expected)
	}
}

func TestShouldReturnExtensionLoggingTo(t *testing.T) {

	result := ParsePermittedForFellingHTML(strings.NewReader(html)).extension.loggingTo

	expected := time.Date(2025, 2, 27, 0, 0, 0, 0, GetLocation())

	if result.Compare(expected) != 0 {
		t.Errorf("Got %s, expected %s", result, expected)
	}
}

func TestShouldReturnExtensionMaterialsUsageTo(t *testing.T) {
	result := ParsePermittedForFellingHTML(strings.NewReader(html)).extension.materialsUsageTo

	expected := time.Date(2025, 2, 28, 0, 0, 0, 0, GetLocation())

	if result.Compare(expected) != 0 {
		t.Errorf("Got %s, expected %s", result, expected)
	}
}

func TestShouldReturnExtensionIssuedBy(t *testing.T) {
	result := ParsePermittedForFellingHTML(strings.NewReader(html)).extension.issuedBy

	expected := "Удължителен Човек"

	if result != expected {
		t.Errorf("Got %s, expected %s", result, expected)
	}
}

func TestShouldReturnExtraction(t *testing.T) {
	result := ParsePermittedForFellingHTML(strings.NewReader(html)).extraction

	if len(result) != 1 {
		t.Errorf("Got %v, expected 1", len(result))
	}
}
