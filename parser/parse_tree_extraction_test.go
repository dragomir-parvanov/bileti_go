/**
This test file is using the
parse_tree_extraction_test.html file for its tests
*/

package parser

import (
	"bileti_go/utils"
	"github.com/PuerkitoBio/goquery"
	"reflect"
	"testing"
)

const TestExtractionFileName = "parse_tree_extraction_test.html"

type filterFunc func(extraction TreeExtraction) bool

func extractTreeNames(extractions []TreeExtraction) []string {
	var names []string

	for _, extraction := range extractions {
		names = append(names, extraction.treeType)
	}

	return names
}

func extractWoodValues(extractions []TreeExtraction) []float64 {
	var values []float64

	for _, extraction := range extractions {
		values = append(values, extraction.value)
	}

	return values
}

func extractWoodNotes(extractions []TreeExtraction) []string {
	var notes []string

	for _, extraction := range extractions {
		notes = append(notes, extraction.note)
	}

	return notes
}

func filter(extractions []TreeExtraction, f filterFunc) []TreeExtraction {
	var filtered []TreeExtraction

	for i := range extractions {
		if f(extractions[i]) {
			filtered = append(filtered, extractions[i])
		}
	}

	return filtered
}

func getTestSelection() *goquery.Selection {
	testFile := GetTestFile(TestExtractionFileName)

	doc := utils.Must(goquery.NewDocumentFromReader(testFile))

	return doc.Find("body")
}

func TestShouldReturnNoLargeConstructionExtraction(t *testing.T) {
	selection := getTestSelection()

	actual := ParseTreeExtraction(selection)

	largeConstruction := filter(actual, func(extraction TreeExtraction) bool {
		return extraction.category == LargeConstructionTimber
	})

	if len(largeConstruction) != 0 {
		t.Errorf(`There should be no large construction extractions, received %v`, len(largeConstruction))
	}
}

func TestShouldReturnNoMediumExtraction(t *testing.T) {
	selection := getTestSelection()

	actual := ParseTreeExtraction(selection)

	mediumConstruction := filter(actual, func(extraction TreeExtraction) bool {
		return extraction.category == MediumConstructionTimber
	})

	if len(mediumConstruction) != 0 {
		t.Errorf(`There should be no large construction extractions, received %v`, len(mediumConstruction))
	}
}

func TestShouldReturnNoSmallExtraction(t *testing.T) {
	selection := getTestSelection()

	actual := ParseTreeExtraction(selection)

	smallConstruction := filter(actual, func(extraction TreeExtraction) bool {
		return extraction.category == SmallConstructionTimber
	})

	if len(smallConstruction) != 0 {
		t.Errorf(`There should be no large construction extractions, received %v`, len(smallConstruction))
	}
}

func TestShouldReturnWoodTreeTypes(t *testing.T) {
	selection := getTestSelection()

	actual := ParseTreeExtraction(selection)

	wood := filter(actual, func(extraction TreeExtraction) bool {
		return extraction.category == Wood
	})

	names := extractTreeNames(wood)

	expectedNames := []string{"Айлант", "Джанка", "Други широколистни", "Ясен"}

	if !reflect.DeepEqual(names, expectedNames) {
		t.Errorf(`The actual array should match the expected array`)
	}
}

func TestShouldReturnWoodValues(t *testing.T) {
	selection := getTestSelection()

	actual := ParseTreeExtraction(selection)

	wood := filter(actual, func(extraction TreeExtraction) bool {
		return extraction.category == Wood
	})

	values := extractWoodValues(wood)

	expectedNotes := []float64{0.1, 0.3, 0.1, 0}

	if !reflect.DeepEqual(values, expectedNotes) {
		t.Errorf(`The actual array should match the expected array`)
	}
}

func TestShouldReturnWoodNotes(t *testing.T) {
	selection := getTestSelection()

	actual := ParseTreeExtraction(selection)

	wood := filter(actual, func(extraction TreeExtraction) bool {
		return extraction.category == Wood
	})

	notes := extractWoodNotes(wood)

	expectedNotes := []string{"", "", "див рожков", ""}

	if !reflect.DeepEqual(notes, expectedNotes) {
		t.Errorf(`The actual array should match the expected array`)
	}
}
