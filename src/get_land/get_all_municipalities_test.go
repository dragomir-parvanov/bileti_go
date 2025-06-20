package get_land

import (
	"bileti_go/src/utils"
	"errors"
	"fmt"
	"io"
	"reflect"
	"strings"
	"testing"
)

func makeTestFetchLand(municipalitiesLength int32) FetchLand {
	idCounter := int32(1)
	return func(regionId int32, municipalityId int32) (io.Reader, error) {
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
`

		for i := int32(0); i < municipalitiesLength; i++ {
			id := idCounter
			idCounter++
			html = html + `\n` + fmt.Sprintf(`<option value="%v">%v</option>`, id, id)
		}

		return strings.NewReader(html), nil
	}
}

func TestShouldReturnAllMunicipalities(t *testing.T) {
	getLand := GetLand{fetchLand: makeTestFetchLand(2)}

	actual := utils.Must(getLand.GetAllMunicipalities([]int32{1, 2}))

	expected := []LandRegionPair{
		{RegionId: 1, Land: Land{id: 1, name: "1"}},
		{RegionId: 1, Land: Land{id: 2, name: "2"}},
		{RegionId: 2, Land: Land{id: 3, name: "3"}},
		{RegionId: 2, Land: Land{id: 4, name: "4"}},
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf(`Expected %v, received %v`, expected, actual)
	}
}

func TestShouldReturnErrorOnFetchError(t *testing.T) {
	expectedError := errors.New("fetch error")
	getLand := GetLand{fetchLand: func(regionId int32, municipalityId int32) (io.Reader, error) {
		return nil, expectedError
	}}

	_, actual := getLand.GetAllMunicipalities([]int32{1})

	if !errors.Is(actual, expectedError) {
		t.Errorf(`Expected %v, received %v`, expectedError, actual)
	}
}
