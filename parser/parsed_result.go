package parser

import (
	utilstime "bileti_go/utils/time"
	"time"
)

type ParsedResult struct {
	id                        string
	directiveNumber           string
	permittedFor              string
	typeOfFelling             string
	section                   string
	subSection                string
	municipality              string
	land                      string
	areaClearing              float64
	ownershipType             string
	treesMarkedBy             string
	controlMarkNumber         string
	controlMarkColor          string
	dateOfCarnetInventory     string
	expectedTreeExtraction    float64
	additionalRequirements    string
	deadlineLogging           utilstime.TimeRange
	deadlineMaterialsUsage    utilstime.TimeRange
	cleaningProcedure         string
	removalToTemporaryStorage string
	issuedBy                  string
	whoReceivedThePermit      string
	issuedOn                  time.Time
	issuedByEmployee          string
	issuedCode                string
}
