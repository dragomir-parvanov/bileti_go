package api_clients

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func getZemlishteHandlerAssert(t *testing.T, returnHtml string) http.HandlerFunc {
	const expectedPath = "/cgi-bin/GetZemlishte.cgi"

	const expectedHeader = "text/html, */*; q=0.01"

	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != expectedPath {
			t.Errorf("Expected to request %s, got: %s", expectedPath, r.URL.Path)
		}
		if r.Header.Get("Accept") != expectedHeader {
			t.Errorf("Expected Accept: %s header, got: %s", expectedHeader, r.Header.Get("Accept"))
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(returnHtml))
	}
}

func TestShouldReturnListOfMunicipalitiesWhenMunicipalitiesPassed(t *testing.T) {
	html := `<html>
    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
        <title></title>
        <link href="http://iag.bg/static/css/reset-min.css" rel="stylesheet" type="text/css"/>
        <link href="http://iag.bg/static/css/style.css" rel="stylesheet" type="text/css"/>
    </head>
    <select id="MunicipalityID" name="MunicipalityID" class="textFiled" onchange="ajax_get_pop()">
        <option value="0">
        <<  >></option>
        <option value="50">Белоградчик</option>
        '.		<option value="51">Бойница</option>
        '.		<option value="52">Брегово</option>
        '.		<option value="53">Видин</option>
        '.		<option value="54">Грамада</option>
        '.		<option value="55">Димово</option>
        '.		<option value="56">Кула</option>
        '.		<option value="57">Макреш</option>
        '.		<option value="58">Ново село</option>
        '.		<option value="59">Ружинци</option>
        '.		<option value="60">Чупрене</option>
        '.
`

	server := httptest.NewServer(getZemlishteHandlerAssert(t, html))
	defer server.Close()

}
