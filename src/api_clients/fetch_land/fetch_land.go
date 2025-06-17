package api_clients

import (
	"bileti_go/src/utils"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type taskRequestName string

const (
	Population   taskRequestName = "GetPop"
	Municipality taskRequestName = "GetMun"
)

func FetchLand(baseUrl string, regionId int32, municipalityId int32) (io.Reader, error) {
	apiUrl := baseUrl + "/cgi-bin/GetZemlishte.cgi"

	form := url.Values{}

	form.Add("Task", string(getTask(municipalityId)))
	form.Add("MunicipalityID", strconv.Itoa(int(municipalityId)))
	form.Add("RegionID", strconv.Itoa(int(regionId)))

	body := strings.NewReader(form.Encode())

	request := utils.Must(http.NewRequest("POST", apiUrl, body))

	request.Header.Set("Accept", "text/html, */*; q=0.01")

	resp, err := http.DefaultClient.Do(request)

	if err != nil {
		return nil, err
	}

	return resp.Body, nil
}

func getTask(municipalityId int32) taskRequestName {
	if municipalityId > 0 {
		return Population
	}

	return Municipality
}
