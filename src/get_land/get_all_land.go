package get_land

type FinalLandLink struct {
	RegionId       int
	MunicipalityId int
	Land           Land
}

func (g GetLand) GetAll(regions []int) ([]FinalLandLink, error) {

	municipalities, err := g.GetAllMunicipalities(regions)

	if err != nil {
		return nil, err
	}

	var links []FinalLandLink

	for _, municipality := range municipalities {
		land, err := g.Get(municipality.RegionId, municipality.Land.id)

		if err != nil {
			return nil, err
		}

		links = append(links, makeLinks(municipality, land)...)
	}

	return links, nil
}

func makeLinks(municipality LandRegionPair, land []Land) []FinalLandLink {
	var links []FinalLandLink

	for _, l := range land {

		links = append(links, FinalLandLink{RegionId: municipality.RegionId, MunicipalityId: municipality.Land.id, Land: l})
	}

	return links
}
