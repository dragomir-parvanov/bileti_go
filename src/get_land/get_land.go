package get_land

import (
	options_parser "bileti_go/src/parser/options"
	"io"
)

type FetchLand func(regionId int32, municipalityId int32) (io.Reader, error)

type GetLand struct {
	fetchLand FetchLand
}

type Land struct {
	id   int32
	name string
}

func (g GetLand) Get(regionId int32, municipalityId int32) ([]Land, error) {
	reader, err := g.fetchLand(regionId, municipalityId)

	if err != nil {
		return nil, err
	}

	options, err := options_parser.ParseOptions(reader)

	if err != nil {
		return nil, err
	}

	land := make([]Land, len(options))

	for i, option := range options {
		land[i] = Land{
			id:   option.Value,
			name: option.Label,
		}
	}

	return land, nil
}
