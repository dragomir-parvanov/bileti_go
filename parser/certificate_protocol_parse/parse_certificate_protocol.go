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
	}
}
