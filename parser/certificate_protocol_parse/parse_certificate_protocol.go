package certificate_protocol_parse

import (
	certificate_protocol_parse_fields "bileti_go/parser/certificate_protocol_parse/fields"
	"bileti_go/utils"
	"github.com/PuerkitoBio/goquery"
	"io"
)

func ParseCertificateProtocol(reader io.Reader) CertificateProtocolParseResult {

	doc := utils.Must(goquery.NewDocumentFromReader(reader))

	selection := doc.Selection

	return CertificateProtocolParseResult{
		id:                           certificate_protocol_parse_fields.ExtractId(selection),
		regionalForestryDirectorate:  certificate_protocol_parse_fields.ExtractRegionalForestryDirectorate(selection),
		permittedForFellingId:        certificate_protocol_parse_fields.ExtractPermittedForFellingId(selection),
		permittedForFellingDate:      utils.Must(certificate_protocol_parse_fields.ExtractPermittedForFellingDate(selection)),
		when:                         utils.Must(certificate_protocol_parse_fields.ExtractWhen(selection)),
		findingsRecordIssuedBy:       certificate_protocol_parse_fields.ExtractFindingsRecordIssuedBy(selection),
		findingsRecordWitnessedBy:    certificate_protocol_parse_fields.ExtractFindingsRecordWitnessedBy(selection),
		municipality:                 certificate_protocol_parse_fields.ExtractMunicipality(selection),
		land:                         certificate_protocol_parse_fields.ExtractLand(selection),
		section:                      certificate_protocol_parse_fields.ExtractSection(selection),
		subsection:                   certificate_protocol_parse_fields.ExtractSubsection(selection),
		cadastreId:                   certificate_protocol_parse_fields.ExtractCadastreId(selection),
		ownership:                    certificate_protocol_parse_fields.ExtractOwnership(selection),
		areaOfPermit:                 certificate_protocol_parse_fields.ExtractAreaOfPermit(selection),
		areaDone:                     certificate_protocol_parse_fields.ExtractAreaDone(selection),
		notMarkedLoggedTreesExist:    certificate_protocol_parse_fields.ExtractNotMarkedLoggedTreesExist(selection),
		volumeDifferenceDescription:  certificate_protocol_parse_fields.ExtractVolumeDifferenceDescription(selection),
		violationDescription:         certificate_protocol_parse_fields.ExtractViolationDescription(selection),
		loggingFinishDescription:     certificate_protocol_parse_fields.ExtractLoggingFinishDescription(selection),
		violationId:                  certificate_protocol_parse_fields.ExtractViolationId(selection),
		violationRecipient:           certificate_protocol_parse_fields.ExtractViolationRecipient(selection),
		toImproveClearingDeadline:    utils.Must(certificate_protocol_parse_fields.ExtractToImproveClearingDeadline(selection)),
		toImproveClearingDescription: certificate_protocol_parse_fields.ExtractToImproveClearingDescription(selection),
		issuedBy:                     certificate_protocol_parse_fields.ExtractIssuedBy(selection),
	}
}
