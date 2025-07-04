package options_parser

import (
	utils_string "bileti_go/src/utils/string"
	"errors"
	"github.com/PuerkitoBio/goquery"
	"io"
	"strconv"
)

func ParseOptions(reader io.Reader) ([]Option, error) {
	doc, err := goquery.NewDocumentFromReader(reader)

	if err != nil {
		return nil, err
	}

	options, errOptions := extractErrorAndOptions(doc)

	return options, errors.Join(errOptions...)
}

func extractErrorAndOptions(doc *goquery.Document) ([]Option, []error) {
	var options []Option
	var optionErrs []error

	doc.Find("option").Each(func(i int, s *goquery.Selection) {
		if i == 0 {
			// the first item is the placeholder
			return
		}
		option, err := extractOption(s)

		if err != nil {
			optionErrs = append(optionErrs, err)
			return
		}

		options = append(options, option)
	})

	return options, optionErrs
}

func extractOption(s *goquery.Selection) (Option, error) {
	attr, exists := s.Attr("value")

	if !exists {
		return Option{}, errors.New(`"value" attribute not found`)
	}

	intVal, err := strconv.ParseInt(attr, 10, 64)

	option := Option{
		Label: utils_string.RemoveBadSpaces(s.Text()),
		Value: int(intVal),
	}

	return option, err
}
