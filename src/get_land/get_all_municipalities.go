package get_land

import "errors"

type LandRegionPair struct {
	RegionId int
	Land     Land
}

func (g GetLand) GetAllMunicipalities(regions []int) ([]LandRegionPair, error) {
	var pairs []LandRegionPair
	var errs []error

	for _, regionId := range regions {
		values, err := g.Get(regionId, 0)

		if err == nil {
			pairs = append(pairs, makePairs(regionId, values)...)
		}

		if err != nil {
			errs = append(errs, err)
		}
	}

	return pairs, errors.Join(errs...)
}

func makePairs(regionId int, land []Land) []LandRegionPair {
	pairs := make([]LandRegionPair, len(land))

	for i, land := range land {
		pairs[i] = LandRegionPair{
			RegionId: regionId,
			Land:     land,
		}
	}
	return pairs
}
