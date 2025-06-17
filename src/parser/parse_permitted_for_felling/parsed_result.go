package parse_permitted_for_felling

import (
	"bileti_go/src/parser/extraction_category"
	utilstime "bileti_go/src/utils/time"
	"time"
)

type ParsedResult struct {
	id                          string
	regionalForestryDirectorate string
	directiveNumber             string
	permittedFor                string
	allowedForester             string
	typeOfFelling               string
	section                     string
	subSection                  string
	accordingToTheInventoryOf   string
	inventoryOrderId            string
	cadastreId                  string
	municipality                string
	land                        string
	areaClearing                float64
	ownershipType               string
	treesMarkedBy               string
	controlMarkNumber           string
	controlMarkColor            string
	dateOfCarnetInventory       time.Time
	expectedTreeExtraction      float64
	extraction                  []TreeExtraction
	additionalRequirements      string
	deadlineLogging             utilstime.TimeRange
	deadlineMaterialsUsage      utilstime.TimeRange
	cleaningProcedure           string
	removalToTemporaryStorage   string
	issuedBy                    string
	whoReceivedThePermit        string
	issuedOn                    time.Time
	issuedByEmployee            string
	issuedCode                  string
	permitIssuePlace            PermitIssuePlace
	extension                   Extension
}

type TreeExtraction struct {
	category parser_extraction_category.ExtractionCategory
	treeType string
	value    float64
	note     string
}

type PermitIssuePlace struct {
	region         string
	municipality   string
	land           string
	address        string
	subSection     string
	gpsCoordinates string
}

type Extension struct {
	loggingDeadlineTo              time.Time
	deadlineTransportingDeadlineTo time.Time
	issuedBy                       string
}
