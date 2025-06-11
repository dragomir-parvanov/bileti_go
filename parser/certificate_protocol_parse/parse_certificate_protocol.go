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
		id:                          certificate_protocol_parse_fields.ExtractId(selection),
		regionalForestryDirectorate: certificate_protocol_parse_fields.ExtractRegionalForestryDirectorate(selection),
		permittedForFellingId:       certificate_protocol_parse_fields.ExtractPermittedForFellingId(selection),
		permittedForFellingDate:     utils.Must(certificate_protocol_parse_fields.ExtractPermittedForFellingDate(selection)),
		when:                        utils.Must(certificate_protocol_parse_fields.ExtractWhen(selection)),
		findingsRecordIssuedBy:      certificate_protocol_parse_fields.ExtractFindingsRecordIssuedBy(selection),
		findingsRecordWitnessedBy:   certificate_protocol_parse_fields.ExtractFindingsRecordWitnessedBy(selection),
		municipality:                certificate_protocol_parse_fields.ExtractMunicipality(selection),
		land:                        certificate_protocol_parse_fields.ExtractLand(selection),
		section:                     certificate_protocol_parse_fields.ExtractSection(selection),
		subsection:                  certificate_protocol_parse_fields.ExtractSubsection(selection),
		cadastreId:                  certificate_protocol_parse_fields.ExtractCadastreId(selection),
		ownership:                   certificate_protocol_parse_fields.ExtractOwnership(selection),
		areaOfPermit:                certificate_protocol_parse_fields.ExtractAreaOfPermit(selection),
	}
}
