package get_land

import (
	"bileti_go/src/utils"
	"errors"
	"io"
	"reflect"
	"testing"
)

func makeTestFetchLandFullLink() FetchLand {

	return func(regionId int, municipalityId int) (io.Reader, error) {

		if municipalityId != 0 {
			return makeTestFetchLandForIds([]int{municipalityId + 1})(regionId, municipalityId)
		}

		return makeTestFetchLandForIds([]int{regionId + 1})(regionId, municipalityId)
	}
}

func TestShouldReturnAllLandsIfOneLink(t *testing.T) {
	regionId := 1

	fetchLand := makeTestFetchLandFullLink()

	getLand := GetLand{fetchLand}

	actual := utils.Must(getLand.GetAll([]int{regionId}))

	expected := []FinalLandLink{
		{RegionId: 1,
			MunicipalityId: 2,
			Land: Land{
				id:   3,
				name: "3",
			},
		},
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %v but got %v", expected, actual)
	}
}

func TestShouldReturnErrorIfFetchingFails(t *testing.T) {
	regionId := 1

	getLand := GetLand{fetchLand: func(regionId int, municipalityId int) (io.Reader, error) {
		return nil, errors.New("fetch error")
	}}

	_, err := getLand.GetAll([]int{regionId})

	if err == nil {
		t.Error("Expected error but got nil")
	}
}

func TestShouldReturnErrorIfMunicipalitiesFetchSucceedButLandFails(t *testing.T) {
	regionId := 1

	getLand := GetLand{fetchLand: func(regionId int, municipalityId int) (io.Reader, error) {
		if municipalityId != 0 {
			return nil, errors.New("fetch error")

		}

		return makeTestFetchLandForIds([]int{1})(regionId, municipalityId)
	}}

	_, err := getLand.GetAll([]int{regionId})

	if err == nil {
		t.Error("Expected error but got nil")
	}
}
