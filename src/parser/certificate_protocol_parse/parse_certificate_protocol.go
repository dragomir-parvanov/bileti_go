package certificate_protocol_parse

import (
	certificate_protocol_parse_fields2 "bileti_go/src/parser/certificate_protocol_parse/fields"
	"bileti_go/src/utils"
	"github.com/PuerkitoBio/goquery"
	"io"
)

func ParseCertificateProtocol(reader io.Reader) CertificateProtocolParseResult {

	doc := utils.Must(goquery.NewDocumentFromReader(reader))

	selection := doc.Selection

	return CertificateProtocolParseResult{
		id:                           certificate_protocol_parse_fields2.ExtractId(selection),
		regionalForestryDirectorate:  certificate_protocol_parse_fields2.ExtractRegionalForestryDirectorate(selection),
		permittedForFellingId:        certificate_protocol_parse_fields2.ExtractPermittedForFellingId(selection),
		permittedForFellingDate:      utils.Must(certificate_protocol_parse_fields2.ExtractPermittedForFellingDate(selection)),
		when:                         utils.Must(certificate_protocol_parse_fields2.ExtractWhen(selection)),
		findingsRecordIssuedBy:       certificate_protocol_parse_fields2.ExtractFindingsRecordIssuedBy(selection),
		findingsRecordWitnessedBy:    certificate_protocol_parse_fields2.ExtractFindingsRecordWitnessedBy(selection),
		municipality:                 certificate_protocol_parse_fields2.ExtractMunicipality(selection),
		land:                         certificate_protocol_parse_fields2.ExtractLand(selection),
		section:                      certificate_protocol_parse_fields2.ExtractSection(selection),
		subsection:                   certificate_protocol_parse_fields2.ExtractSubsection(selection),
		cadastreId:                   certificate_protocol_parse_fields2.ExtractCadastreId(selection),
		ownership:                    certificate_protocol_parse_fields2.ExtractOwnership(selection),
		areaOfPermit:                 certificate_protocol_parse_fields2.ExtractAreaOfPermit(selection),
		areaDone:                     certificate_protocol_parse_fields2.ExtractAreaDone(selection),
		extraction:                   ExtractExtraction(selection),
		notMarkedLoggedTreesExist:    certificate_protocol_parse_fields2.ExtractNotMarkedLoggedTreesExist(selection),
		volumeDifferenceDescription:  certificate_protocol_parse_fields2.ExtractVolumeDifferenceDescription(selection),
		violationDescription:         certificate_protocol_parse_fields2.ExtractViolationDescription(selection),
		loggingFinishDescription:     certificate_protocol_parse_fields2.ExtractLoggingFinishDescription(selection),
		violationId:                  certificate_protocol_parse_fields2.ExtractViolationId(selection),
		violationRecipient:           certificate_protocol_parse_fields2.ExtractViolationRecipient(selection),
		toImproveClearingDeadline:    utils.Must(certificate_protocol_parse_fields2.ExtractToImproveClearingDeadline(selection)),
		toImproveClearingDescription: certificate_protocol_parse_fields2.ExtractToImproveClearingDescription(selection),
		issuedBy:                     certificate_protocol_parse_fields2.ExtractIssuedBy(selection),
		witnessedBy:                  certificate_protocol_parse_fields2.ExtractWitnessedBy(selection),
	}
}
