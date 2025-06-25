package get_land

import (
	"fmt"
	"io"
	"strings"
)

func makeTestFetchLandForIds(ids []int) FetchLand {
	return func(regionId int, municipalityId int) (io.Reader, error) {
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

		for _, id := range ids {
			html = html + `\n` + fmt.Sprintf(`<option value="%v">%v</option>`, id, id)
		}

		return strings.NewReader(html), nil
	}
}
