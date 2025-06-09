package certificate_protocol_parse_fields

import (
	"bileti_go/parser"
	"time"
)

type CertificateProtocolParseResult struct {
	id                           string
	regionalForestryDirectorate  string
	permittedForFellingId        string
	permittedForFellingDate      string
	when                         time.Time
	findingsRecordIssuedBy       string
	findingsRecordWitnessedBy    string
	municipality                 string
	land                         string
	section                      string
	subSection                   string
	cadastreId                   string
	ownership                    string
	areOfPermit                  float64
	areaDone                     float64
	loggedNotMarkedTrees         bool
	volumeDifferenceDescription  string
	violationDescription         string
	loggingFinishDescription     string
	violationId                  string
	violationRecipient           string
	toImproveClearingDeadline    time.Time
	toImproveClearingDescription string
	createdBy                    string
	witnessedBy                  string
}

type CertificateProtocolTreeExtraction struct {
	category                    parser.ExtractionCategory
	treeType                    string
	byLoggingPermitCollected    float64
	actuallyCollected           float64
	availableInTheClearing      float64
	availableInTemporaryStorage float64
}
