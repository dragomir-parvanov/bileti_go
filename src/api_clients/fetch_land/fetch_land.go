package api_clients

import (
	"bileti_go/src/utils"
	"fmt"
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

type FetchLand struct {
	BaseUrl string
}

func (f FetchLand) Fetch(regionId int, municipalityId int) (io.Reader, error) {
	apiUrl := f.BaseUrl + "/cgi-bin/GetZemlishte.cgi"

	form := url.Values{}

	form.Add("Task", string(getTask(municipalityId)))
	form.Add("MunicipalityID", strconv.Itoa(municipalityId))
	form.Add("RegionID", strconv.Itoa(regionId))

	body := strings.NewReader(form.Encode())

	request := utils.Must(http.NewRequest("POST", apiUrl, body))

	request.Header.Set("Accept", "text/html, */*; q=0.01")

	resp, err := http.DefaultClient.Do(request)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("request failed with status code %d", resp.StatusCode)
	}

	return resp.Body, nil
}

func getTask(municipalityId int) taskRequestName {
	if municipalityId > 0 {
		return Population
	}

	return Municipality
}
