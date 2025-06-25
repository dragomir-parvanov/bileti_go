package get_land

import (
	"bileti_go/src/utils"
	"errors"
	"io"
	"reflect"
	"strings"
	"testing"
)

func TestShouldReturnLand(t *testing.T) {
	html := `
<html>
    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
        <title></title>
        <link href="http://iag.bg/static/css/reset-min.css" rel="stylesheet" type="text/css"/>
        <link href="http://iag.bg/static/css/style.css" rel="stylesheet" type="text/css"/>
    </head>
    <select id="MunicipalityID" name="MunicipalityID" class="textFiled" onchange="ajax_get_pop()">
        <option value="0">
        <<  >></option>
        <option value="1">Банско</option>
        '.		<option value="2">Белица</option>
        '.
`

	getLand := GetLand{fetchLand: func(regionId int, municipalityId int) (io.Reader, error) {
		return strings.NewReader(html), nil
	}}

	actual := utils.Must(getLand.Get(0, 0))

	expected := []Land{{id: 1, name: "Банско"}, {id: 2, name: "Белица"}}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}

func TestShouldReturnErrorOnFetchLandErr(t *testing.T) {
	fetchLandError := errors.New("fetchLand error")

	getLand := GetLand{fetchLand: func(regionId int, municipalityId int) (io.Reader, error) {
		return nil, fetchLandError
	}}

	_, err := getLand.Get(0, 0)

	if err == nil {
		t.Errorf("expected error, but got nil")
	}
}

func TestShouldReturnErrorOnParsingErr(t *testing.T) {
	// removing the "value" attribute
	// to trigger a parsing error
	html := `
<html>
    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
        <title></title>
        <link href="http://iag.bg/static/css/reset-min.css" rel="stylesheet" type="text/css"/>
        <link href="http://iag.bg/static/css/style.css" rel="stylesheet" type="text/css"/>
    </head>
    <select id="MunicipalityID" name="MunicipalityID" class="textFiled" onchange="ajax_get_pop()">
        <option value="0">
        <<  >></option>
        <option>Банско</option>
        '.		<option >Белица</option>
        '.
`

	getLand := GetLand{fetchLand: func(regionId int, municipalityId int) (io.Reader, error) {
		return strings.NewReader(html), nil
	}}

	_, err := getLand.Get(0, 0)

	if err == nil {
		t.Errorf("expected error, but got nil")
	}
}
