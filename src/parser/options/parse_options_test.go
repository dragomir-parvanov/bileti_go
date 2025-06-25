package options_parser

import (
	"bileti_go/src/utils"
	"bileti_go/src/utils/test"
	"errors"
	"reflect"
	"strings"
	"testing"
)

func TestShouldReturnOptions(t *testing.T) {

	html := `
<html>
    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
        <title></title>
        <link href="http://iag.bg/static/css/reset-min.css" rel="stylesheet" type="text/css"/>
        <link href="http://iag.bg/static/css/style.css" rel="stylesheet" type="text/css"/>
    </head>
    <select id="PopulatedPlaceID" name="PopulatedPlaceID" class="textFiled">
        <option Value="0">
        <<  >></option>Select * From regions.PopulatedPlaces where RegID = '1' and MunID = '2' Order by PolpulatedPlace		
        <option Value="14">Бабяк</option>
        '.		<option Value="13">Белица (Бл)</option>
        '.		<option Value="16">Горно Краище</option>
        '.		<option Value="12">Гълъбово (Бл)</option>
        '.		<option Value="11">Дагоново</option>
        '.		<option Value="15">Златарица (Бл)</option>
        '.		<option Value="10">Краище (Бл)</option>
        '.		<option Value="9">Кузьово</option>
        '.		<option Value="19">Лютово</option>
        '.		<option Value="18">Орцево</option>
        '.		<option Value="20">Палатик</option>
        '.		<option Value="17">Черешово (Бл)</option>
        '.
`

	actual := utils.Must(ParseOptions(strings.NewReader(html)))

	expected := []Option{
		{Value: 14, Label: "Бабяк"},
		{Value: 13, Label: "Белица (Бл)"},
		{Value: 16, Label: "Горно Краище"},
		{Value: 12, Label: "Гълъбово (Бл)"},
		{Value: 11, Label: "Дагоново"},
		{Value: 15, Label: "Златарица (Бл)"},
		{Value: 10, Label: "Краище (Бл)"},
		{Value: 9, Label: "Кузьово"},
		{Value: 19, Label: "Лютово"},
		{Value: 18, Label: "Орцево"},
		{Value: 20, Label: "Палатик"},
		{Value: 17, Label: "Черешово (Бл)"},
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("\nExpected: %#v\nActual:   %#v", expected, actual)
	}
}

func TestShouldReturnErrorOnReaderError(t *testing.T) {
	val, err := ParseOptions(&utils_test.ErrReader{Error: errors.New("reader error")})

	if err == nil {
		t.Errorf("error should be defined, instead received Value %v", val)
	}
}

func TestShouldReturnErrorOnOptionsWithNoValue(t *testing.T) {
	html := `
<html>
    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
        <title></title>
        <link href="http://iag.bg/static/css/reset-min.css" rel="stylesheet" type="text/css"/>
        <link href="http://iag.bg/static/css/style.css" rel="stylesheet" type="text/css"/>
    </head>
    <select id="PopulatedPlaceID" name="PopulatedPlaceID" class="textFiled">
        <option Value="0">
        <<  >></option>Select * From regions.PopulatedPlaces where RegID = '1' and MunID = '2' Order by PolpulatedPlace		
        <option Value="14">Бабяк</option>
        '.		<option>Белица (Бл)</option>
        '.
`

	_, err := ParseOptions(strings.NewReader(html))

	expectedErrMessage := `"value" attribute not found`

	if err.Error() != expectedErrMessage {
		t.Errorf("\nExpected: %s\nActual:   %s\n", expectedErrMessage, err.Error())
	}
}

func TestShouldReturnWithoutSpaces(t *testing.T) {
	html := `
<html>
    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
        <title></title>
        <link href="http://iag.bg/static/css/reset-min.css" rel="stylesheet" type="text/css"/>
        <link href="http://iag.bg/static/css/style.css" rel="stylesheet" type="text/css"/>
    </head>
    <select id="PopulatedPlaceID" name="PopulatedPlaceID" class="textFiled">
        <option value="0">
        <<  >></option>Select * From regions.PopulatedPlaces where RegID = '12' and MunID = '111' Order by PolpulatedPlace		
        <option value="2732"> Вършец</option>

`

	actual := utils.Must(ParseOptions(strings.NewReader(html)))

	expected := []Option{
		{Value: 2732, Label: "Вършец"},
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("\nExpected: %#v\nActual:   %#v", expected, actual)
	}

}
