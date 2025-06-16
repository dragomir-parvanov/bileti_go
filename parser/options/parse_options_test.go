package options_parser

import (
	"bileti_go/utils"
	"fmt"
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
        <option value="0">
        <<  >></option>Select * From regions.PopulatedPlaces where RegID = '1' and MunID = '2' Order by PolpulatedPlace		
        <option value="14">Бабяк</option>
        '.		<option value="13">Белица (Бл)</option>
        '.		<option value="16">Горно Краище</option>
        '.		<option value="12">Гълъбово (Бл)</option>
        '.		<option value="11">Дагоново</option>
        '.		<option value="15">Златарица (Бл)</option>
        '.		<option value="10">Краище (Бл)</option>
        '.		<option value="9">Кузьово</option>
        '.		<option value="19">Лютово</option>
        '.		<option value="18">Орцево</option>
        '.		<option value="20">Палатик</option>
        '.		<option value="17">Черешово (Бл)</option>
        '.
`

	actual := utils.Must(ParseOptions(strings.NewReader(html)))

	expected := []Option{
		{value: 14, label: "Бабяк"},
		{value: 13, label: "Белица (Бл)"},
		{value: 16, label: "Горно Краище"},
		{value: 12, label: "Гълъбово (Бл)"},
		{value: 11, label: "Дагоново"},
		{value: 15, label: "Златарица (Бл)"},
		{value: 10, label: "Краище (Бл)"},
		{value: 9, label: "Кузьово"},
		{value: 19, label: "Лютово"},
		{value: 18, label: "Орцево"},
		{value: 20, label: "Палатик"},
		{value: 17, label: "Черешово (Бл)"},
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("\nExpected: %#v\nActual:   %#v", expected, actual)
	}
}

func TestShouldReturnErrorOnEmpty(t *testing.T) {
	html := `

`

	_, err := ParseOptions(strings.NewReader(html))

	if err != nil {
		fmt.Errorf("error should be defined")
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
        <option value="0">
        <<  >></option>Select * From regions.PopulatedPlaces where RegID = '1' and MunID = '2' Order by PolpulatedPlace		
        <option value="14">Бабяк</option>
        '.		<option>Белица (Бл)</option>
        '.
`

	_, err := ParseOptions(strings.NewReader(html))

	expectedErrMessage := "option attribute not found"

	if err.Error() != expectedErrMessage {
		t.Errorf("\nExpected: %s\nActual:   %s\n", expectedErrMessage, err.Error())
	}
}
