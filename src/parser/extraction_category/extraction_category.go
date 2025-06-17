package parser_extraction_category

type ExtractionCategory string

const (
	LargeConstructionTimber  ExtractionCategory = "LargeConstructionTimber"
	MediumConstructionTimber ExtractionCategory = "MediumConstructionTimber"
	SmallConstructionTimber  ExtractionCategory = "SmallConstructionTimber"
	Wood                     ExtractionCategory = "Wood"
	TopHamper                ExtractionCategory = "TopHamper"
)

func GetMap() map[ExtractionCategory]string {
	categoryMap := map[ExtractionCategory]string{
		LargeConstructionTimber:  "Едра строителна дървесина",
		MediumConstructionTimber: "Средна строителна дървесина",
		SmallConstructionTimber:  "Дребна строителна дървесина",
		Wood:                     "Дърва",
		TopHamper:                "Вършина",
	}

	return categoryMap
}

func GetReverseMap() map[string]ExtractionCategory {
	categoryMap := GetMap()

	reversedMap := make(map[string]ExtractionCategory)

	for k, v := range categoryMap {
		reversedMap[v] = k
	}

	return reversedMap
}
