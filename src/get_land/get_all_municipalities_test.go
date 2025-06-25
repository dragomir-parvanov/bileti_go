package get_land

import (
	"bileti_go/src/utils"
	"errors"
	"io"
	"reflect"
	"testing"
)

func makeTestFetchLandForMunicipalities(municipalitiesLength int) FetchLand {
	idCounter := 1

	return func(regionId int, municipalityId int) (io.Reader, error) {
		var ids []int
		for i := 0; i < municipalitiesLength; i++ {
			ids = append(ids, idCounter)
			idCounter++
		}

		return makeTestFetchLandForIds(ids)(regionId, municipalityId)
	}
}

func TestShouldReturnAllMunicipalities(t *testing.T) {
	getLand := GetLand{fetchLand: makeTestFetchLandForMunicipalities(2)}

	actual := utils.Must(getLand.GetAllMunicipalities([]int{1, 2}))

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
	getLand := GetLand{fetchLand: func(regionId int, municipalityId int) (io.Reader, error) {
		return nil, expectedError
	}}

	_, actual := getLand.GetAllMunicipalities([]int{1})

	if !errors.Is(actual, expectedError) {
		t.Errorf(`Expected %v, received %v`, expectedError, actual)
	}
}
