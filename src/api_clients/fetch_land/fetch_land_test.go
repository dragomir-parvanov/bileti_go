package api_clients

import (
	"bileti_go/src/utils"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func getZemlishteHandlerAssert(t *testing.T, expectedBody string) http.HandlerFunc {
	const expectedPath = "/cgi-bin/GetZemlishte.cgi"

	const expectedHeader = "text/html, */*; q=0.01"

	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("Expected to request with POST, got: %s", r.Method)
		}

		if r.URL.Path != expectedPath {
			t.Errorf("Expected to request %s, got: %s", expectedPath, r.URL.Path)
		}
		if r.Header.Get("Accept") != expectedHeader {
			t.Errorf("Expected Accept: %s header, got: %s", expectedHeader, r.Header.Get("Accept"))
		}

		body := string(utils.Must(io.ReadAll(r.Body)))

		if !reflect.DeepEqual(body, expectedBody) {
			t.Errorf("Expected to get %v, got: %v", expectedBody, body)
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Ok Response"))
	}
}

func TestShouldReturnListOfLand(t *testing.T) {

	server := httptest.NewServer(getZemlishteHandlerAssert(t, "MunicipalityID=0&RegionID=0&Task=GetMun"))
	defer server.Close()

	actual := string(utils.Must(io.ReadAll(utils.Must(FetchLand(server.URL, 0, 0)))))

	expected := "Ok Response"

	if actual != expected {
		t.Errorf("Expected to get %s, got %s", expected, actual)
	}
}

func TestShouldReturnTaskGetMunIfMunicipalityIdIsZero(t *testing.T) {

	server := httptest.NewServer(getZemlishteHandlerAssert(t, "MunicipalityID=0&RegionID=0&Task=GetMun"))
	defer server.Close()

	FetchLand(server.URL, 0, 0)

}

func TestShouldReturnTaskGetPopIfMunicipalityIdIsNotZero(t *testing.T) {

	server := httptest.NewServer(getZemlishteHandlerAssert(t, "MunicipalityID=1&RegionID=0&Task=GetPop"))
	defer server.Close()

	FetchLand(server.URL, 0, 1)
}
