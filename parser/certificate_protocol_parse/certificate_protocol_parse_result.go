package certificate_protocol_parse

import (
	"bileti_go/parser"
	"time"
)

type CertificateProtocolParseResult struct {
	id                           string
	regionalForestryDirectorate  string
	permittedForFellingId        string
	permittedForFellingDate      time.Time
	when                         time.Time
	findingsRecordIssuedBy       string
	findingsRecordWitnessedBy    string
	municipality                 string
	land                         string
	section                      string
	subsection                   string
	cadastreId                   string
	ownership                    string
	areaOfPermit                 float64
	areaDone                     float64
	notMarkedLoggedTreesExist    bool
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
