package parser

import (
	utilsstring "bileti_go/utils/string"
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"regexp"
)

func ParsePermittedForFellingHTML(htmlResult io.Reader) ParsedResult {

	doc, err := goquery.NewDocumentFromReader(htmlResult)

	if err != nil {
		log.Fatal(err)
	}

	id := _ExtractId(doc)
	directiveNumber := _ExtractDirectiveNumber(doc)
	permittedFor := _ExtractPermittedFor(doc)
	typeOfFelling := _ExtractTypeOfFelling(doc)
	section := _ExtractSection(doc)
	subSection := _ExtractSubSection(doc)
	municipality := _ExtractMunicipality(doc)
	land := _ExtractLand(doc)

	return ParsedResult{
		id, directiveNumber,
		permittedFor,
		typeOfFelling,
		section,
		subSection,
		municipality,
		land}
}

func _ExtractId(doc *goquery.Document) string {
	regex := regexp.MustCompile("(?i)позволително за сеч № ([0-9]+)")
	text := doc.Text()

	matched := regex.FindStringSubmatch(text)[1]

	return utilsstring.RemoveBadSpaces(matched)
}

func _ExtractDirectiveNumber(doc *goquery.Document) string {
	directiveNumber := doc.Find("td:contains('На основание')").First().Find("b").Text()

	return utilsstring.RemoveBadSpaces(directiveNumber)
}

func _ExtractPermittedFor(doc *goquery.Document) string {
	permittedFor := doc.Find("td:contains('разрешава се на')").First().Find("b").Text()

	return utilsstring.RemoveBadSpaces(permittedFor)
}

func _ExtractTypeOfFelling(doc *goquery.Document) string {
	typeOfFelling := doc.Find("td:contains('ВИД НА СЕЧТА')").First().Find("b").Text()

	return utilsstring.RemoveBadSpaces(typeOfFelling)
}

func _ExtractSection(doc *goquery.Document) string {
	section := doc.Find("td:contains('да извърши добива в отдел №')").First().Find("b").First().Text()

	return utilsstring.RemoveBadSpaces(section)
}

func _ExtractSubSection(doc *goquery.Document) string {
	subSection := doc.Find("td:contains('подотдел')").First().Find("b").Next().Text()

	return utilsstring.RemoveBadSpaces(subSection)
}

func _ExtractMunicipality(doc *goquery.Document) string {
	municipality := doc.Find("td:contains('Община')").First().Find("b").First().Text()

	return utilsstring.RemoveBadSpaces(municipality)
}

func _ExtractLand(doc *goquery.Document) string {
	municipality := doc.Find("td:contains('Землище')").First().Find("b").Next().First().Text()

	return utilsstring.RemoveBadSpaces(municipality)
}
