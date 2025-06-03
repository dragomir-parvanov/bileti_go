package parser

import (
	utilstime "bileti_go/utils/time"
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
}

type PermitIssuePlace struct {
	region         string
	municipality   string
	land           string
	address        string
	subSection     string
	gpsCoordinates string
}
