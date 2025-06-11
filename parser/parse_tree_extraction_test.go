package parser

import (
	"bileti_go/utils"
	"github.com/PuerkitoBio/goquery"
	"reflect"
	"strings"
	"testing"
)

const treeExtractionHtml = `
<html><head>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
    <title></title>
    <link rel="StyleSheet" type="text/css" href="/style.css">
    <script type="text/javascript" language="JavaScript" src="/js/jquery.js"></script>
    <script type="text/javascript">
        $(document).ready(function(){
            $("#fr").attr("style", "pointer-events: none; border: none;");
            $("#fr").removeAttr("hidden");
            $("#fr").contents().find("#navigationBar").remove();
            $("#fr").contents().find("#mapstore-drawermenu").remove();
            $("#fr").contents().find("#mapstore-navbar").remove();
            $("#fr").contents().find('iframe[id="navigationBar"]').remove();
        });
    </script>
</head>
<body data-new-gr-c-s-check-loaded="14.1238.0" data-gr-ext-installed="">
<table width="650" align="center" border="0" summary="" class="sech">
    <tbody><tr>
        <td rowspan="2" align="center" width="80" height="53"><img src="/images/logo.png"></td>
        <td align="center"><font size="+1">ИЗПЪЛНИТЕЛНА АГЕНЦИЯ ПО ГОРИТЕ</font></td>
    </tr>
    <tr>
        <td colspan="2" align="center"><font size="2px">Регионална дирекция по горите <b>Варна</b></font></td>
    </tr>
    <tr>
        <td colspan="2" align="center"><hr></td>
    </tr>
    </tbody></table>
<table border="0" align="center" width="650" cellpadding="0" cellspacing="8" class="sech">
    <tbody><tr>
        <td colspan="2" align="left">ВИД НА СЕЧТА : <b>по чл. 51в от Наредба 8 за сечите в горите&nbsp;</b>
        </td>
    </tr>
    <tr>
        <td colspan="2" align="center"><br><font size="2"><b>ПОЗВОЛИТЕЛНО за СЕЧ № 0800703</b>
        </font></td>
    </tr>
    <tr>
        <td colspan="2" align="left"><br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;На основание чл. 108 от Закона за горите, Заповед № <b>/&nbsp;</b> г. за утвърждаване на
        </td>
    </tr>
    <tr>
        <td colspan="2" align="left">горскостопански план (горскостопанска програма)</td>
    </tr>
    <tr>
        <td colspan="2" align="left">разрешава се на <b>ТП ДЪРЖАВНО ЛОВНО СТОПАНСТВО БАЛЧИК</b></td>
    </tr>
    <tr>
        <td colspan="2" align="left">с нает регистриран лесовъд &nbsp;<b>..........</b> , </td>
    </tr>
    <tr>
        <td colspan="2" align="left">да извърши добива в отдел № <b>53</b>; подотдел <b>10</b>; имот с (кадастрален/КВС) № <b>02508.62.38</b>, находящ се в: </td>
    </tr>
    <tr>
        <td colspan="2" align="left">Община <b>Балчик</b>; Землище <b>Балчик</b> площно сечище от <b>0.090</b> хектара; Вид собственост : <b>ДГТ</b></td>
    </tr>
    <tr>
        <td colspan="2" align="left"> Дърветата са маркирани от <b>..........</b></td>
    </tr>
    <tr>
        <td colspan="2" align="left"> с контролна горска марка № <b>А 4290</b>, с <b>оранжева</b> боя, дата на карнет опис: <b>&nbsp;</b> г.</td>
    </tr>
    <tr>
        <td colspan="2" align="left"> Очакваният добив е <b>0.5</b> плътни кубически метра лежаща маса, която по категории е както следва :</td>
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
        <td rowspan="4" align="center">4.</td>
        <td rowspan="4">&nbsp;Дърва</td>
        <td><b>Айлант</b></td>
        <td align="right"><b>0.1</b></td>
        <td align="center"><font size="1"></font></td>
    </tr>
    <tr>
        <td><b>Джанка</b></td>
        <td align="right"><b>0.3</b></td>
        <td align="center"><font size="1"></font></td>
    </tr>
    <tr>
        <td><b>Други широколистни</b></td>
        <td align="right"><b>0.1</b></td>
        <td align="center"><font size="1">див рожков</font></td>
    </tr>
    <tr>
        <td><b>Ясен</b></td>
        <td align="right"><b>0.0</b></td>
        <td align="center"><font size="1"></font></td>
    </tr>
    <tr>
        <td rowspan="1" align="center">5.</td>
        <td rowspan="1">&nbsp;Вършина</td>
        <td><b> --- </b></td>
        <td align="right"><b> --- </b></td>
        <td align="center"><font size="1"></font></td>
    </tr>
    </tbody></table>
<table border="0" align="center" width="650" cellpadding="0" cellspacing="8" class="sech">
    <tbody><tr>
        <td colspan="2" align="left"> Допълнителни изисквания при провеждане на сечта : <b>Съгласно издаден констативен протокол Серия Е00/002219/04.12.2024г. по чл 51в от НСГ от служители на РДГ Варна</b></td>
    </tr>
    <tr>
        <td colspan="2" align="left"> Срок за провеждане на сечта от <b>13.01.2025</b>&nbsp;г. до <b>31.03.2025</b>&nbsp;г. </td>
    </tr>
    <tr>
        <td colspan="2" align="left"> Срок за извозване на материалите от сечището от <b>13.01.2025</b>&nbsp;г. до <b>31.03.2025</b>&nbsp;г. </td>
    </tr>
    <tr>
        <td colspan="2" align="left"> Начин на почистване на сечището : <b>&nbsp;</b></td>
    </tr>
    <tr>
        <td colspan="2" align="left"> Материалите ще се извозят до временен склад : <b>съгласно ТП</b></td>
    </tr>
    <tr>
        <td colspan="2">&nbsp;</td>
    </tr>
    <tr>
        <td align="left" width="315">Издал:..........................................</td>
        <td align="left" width="335">Получил позволителното:..........................................</td>
    </tr>
    <tr>
        <td align="left" width="325">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;/<b> .......... </b>/</td>
        <td align="left" width="325">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
            &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;/<b> .......... </b>/</td>
    </tr>
    <tr>
        <td colspan="2" align="left"> Дата: <b>08.01.2025</b> г.
            &nbsp;&nbsp;Издал служител : |..........|
            &nbsp;&nbsp;Код: |..........|</td>
    </tr>
    <tr>
        <td>Място на издаване на позволителното за сеч: </td>	</tr>
    <tr>
        <td colspan="2" align="left">Област <b>Добрич</b>, община <b>Балчик</b>, землище <b>Балчик</b>, адрес, подотдел, GPS координати: .......................</td>
    </tr>
    <tr>
        <td colspan="2"><b><b></b></b></td>
    </tr>
    <tr>
        <td colspan="2" align="center"><b>ЗАВЕРКИ ПРИ ПРОДЪЛЖАВАНЕ СРОКОВЕТЕ ЗА СЕЧ И ИЗВОЗ</b></td>
    </tr>
    <tr>
        <td align="left"> За провеждане на сечта до : ...................... г.</td>
        <td align="left"> За извоз на материалите до : ...................... г.</td>
    </tr>
    <tr>
        <td colspan="2" align="left">Издал:  ............................................</td>
    </tr>
    <tr>
        <td colspan="2" align="left">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;/
            &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;/</td>
    </tr>
    <tr>
        <td colspan="2" align="left">Продължаване на сроковете за сеч и извоз става само чрез информационната система на ИАГ.</td>
    </tr>
    </tbody></table>


</body></html>
`

type filterFunc func(extraction TreeExtraction) bool

func extractTreeNames(extractions []TreeExtraction) []string {
	var names []string

	for _, extraction := range extractions {
		names = append(names, extraction.treeType)
	}

	return names
}

func extractWoodValues(extractions []TreeExtraction) []float64 {
	var values []float64

	for _, extraction := range extractions {
		values = append(values, extraction.value)
	}

	return values
}

func extractWoodNotes(extractions []TreeExtraction) []string {
	var notes []string

	for _, extraction := range extractions {
		notes = append(notes, extraction.note)
	}

	return notes
}

func filter(extractions []TreeExtraction, f filterFunc) []TreeExtraction {
	var filtered []TreeExtraction

	for i := range extractions {
		if f(extractions[i]) {
			filtered = append(filtered, extractions[i])
		}
	}

	return filtered
}

func getTestSelection() *goquery.Selection {

	doc := utils.Must(goquery.NewDocumentFromReader(strings.NewReader(treeExtractionHtml)))

	return doc.Find("body")
}

func TestShouldReturnNoLargeConstructionExtraction(t *testing.T) {
	selection := getTestSelection()

	actual := ParseTreeExtraction(selection)

	largeConstruction := filter(actual, func(extraction TreeExtraction) bool {
		return extraction.category == LargeConstructionTimber
	})

	if len(largeConstruction) != 0 {
		t.Errorf(`There should be no large construction extractions, received %v`, len(largeConstruction))
	}
}

func TestShouldReturnNoMediumExtraction(t *testing.T) {
	selection := getTestSelection()

	actual := ParseTreeExtraction(selection)

	mediumConstruction := filter(actual, func(extraction TreeExtraction) bool {
		return extraction.category == MediumConstructionTimber
	})

	if len(mediumConstruction) != 0 {
		t.Errorf(`There should be no large construction extractions, received %v`, len(mediumConstruction))
	}
}

func TestShouldReturnNoSmallExtraction(t *testing.T) {
	selection := getTestSelection()

	actual := ParseTreeExtraction(selection)

	smallConstruction := filter(actual, func(extraction TreeExtraction) bool {
		return extraction.category == SmallConstructionTimber
	})

	if len(smallConstruction) != 0 {
		t.Errorf(`There should be no large construction extractions, received %v`, len(smallConstruction))
	}
}

func TestShouldReturnWoodTreeTypes(t *testing.T) {
	selection := getTestSelection()

	actual := ParseTreeExtraction(selection)

	wood := filter(actual, func(extraction TreeExtraction) bool {
		return extraction.category == Wood
	})

	names := extractTreeNames(wood)

	expectedNames := []string{"Айлант", "Джанка", "Други широколистни", "Ясен"}

	if !reflect.DeepEqual(names, expectedNames) {
		t.Errorf(`The actual array should match the expected array`)
	}
}

func TestShouldReturnWoodValues(t *testing.T) {
	selection := getTestSelection()

	actual := ParseTreeExtraction(selection)

	wood := filter(actual, func(extraction TreeExtraction) bool {
		return extraction.category == Wood
	})

	values := extractWoodValues(wood)

	expectedNotes := []float64{0.1, 0.3, 0.1, 0}

	if !reflect.DeepEqual(values, expectedNotes) {
		t.Errorf(`The actual array should match the expected array`)
	}
}

func TestShouldReturnWoodNotes(t *testing.T) {
	selection := getTestSelection()

	actual := ParseTreeExtraction(selection)

	wood := filter(actual, func(extraction TreeExtraction) bool {
		return extraction.category == Wood
	})

	notes := extractWoodNotes(wood)

	expectedNotes := []string{"", "", "див рожков", ""}

	if !reflect.DeepEqual(notes, expectedNotes) {
		t.Errorf(`The actual array should match the expected array`)
	}
}
