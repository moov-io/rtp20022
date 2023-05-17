package response_return_of_funds

import (
	"encoding/xml"
	"github.com/moov-io/rtp20022/gen/camt_029_001_09"
	"github.com/moov-io/rtp20022/pkg/dt"
	"time"
)

type Writer struct {
	XMLName         xml.Name `xml:"ResponseReturnOfFunds"`
	Text            string   `xml:",chardata"`
	Head            string   `xml:"xmlns:rf,attr"`
	RsltnOfInvstgtn struct {
		Text    string     `xml:",chardata"`
		Head    string     `xml:"xmlns:rf,attr"`
		Attrs   []xml.Attr `xml:",attr"`
		Assgnmt struct {
			Id     camt_029_001_09.Max35Text `xml:"rf:Id"`
			Assgnr struct {
				Agt struct {
					FinInstnId struct {
						ClrSysMmbId struct {
							MmbId camt_029_001_09.Max35Text `xml:"rf:MmbId"`
						} `xml:"rf:ClrSysMmbId"`
					} `xml:"rf:FinInstnId"`
				} `xml:"rf:Agt"`
			} `xml:"rf:Assgnr"`
			Assgne struct {
				Agt struct {
					FinInstnId struct {
						ClrSysMmbId struct {
							MmbId camt_029_001_09.Max35Text `xml:"rf:MmbId"`
						} `xml:"rf:ClrSysMmbId"`
					} `xml:"rf:FinInstnId"`
				} `xml:"rf:Agt"`
			} `xml:"rf:Assgne"`
			CreDtTm dt.ISODateTime `xml:"rf:CreDtTm"`
		} `xml:"rf:Assgnmt"`
		Sts struct {
			Conf *camt_029_001_09.ExternalInvestigationExecutionConfirmation1Code `xml:"rf:Conf"`
		} `xml:"rf:Sts"`
		CxlDtls struct {
			OrgnlGrpInfAndSts struct {
				RslvdCase struct {
					Id    camt_029_001_09.Max35Text `xml:"rf:Id"`
					Cretr struct {
						Agt struct {
							FinInstnId struct {
								ClrSysMmbId struct {
									MmbId camt_029_001_09.Max35Text `xml:"rf:MmbId"`
								} `xml:"rf:ClrSysMmbId"`
							} `xml:"rf:FinInstnId"`
						} `xml:"rf:Agt"`
					} `xml:"rf:Cretr"`
				} `xml:"rf:RslvdCase"`
				OrgnlMsgId   camt_029_001_09.Max35Text   `xml:"rf:OrgnlMsgId"`
				OrgnlMsgNmId camt_029_001_09.OrigMsgName `xml:"rf:OrgnlMsgNmId"`
				OrgnlCreDtTm dt.ISODateTime              `xml:"rf:OrgnlCreDtTm"`
			} `xml:"rf:OrgnlGrpInfAndSts"`
		} `xml:"rf:CxlDtls"`
	} `xml:"rf:RsltnOfInvstgtn"`
}

func (r Writer) Validate() error {
	return nil
}

// WriteParams is a human-readable struct of the Writer fields
type WriteParams struct {
	CreatedOn             time.Time
	CaseAssignmentId      string
	AssignorAgentMemberId string
	AssigneeAgentMemberId string
	StsConfCd             string
	OriginalMessageId     string
	OriginalMessageNameId string
	OriginalCreatedOn     time.Time
}

const (
	Namespace = "urn:iso:std:iso:20022:tech:xsd:camt.029.001.09"
)

func NewWriter(params WriteParams) Writer {
	headAttr := xml.Attr{
		Name:  xml.Name{Local: "xmlns:rf"},
		Value: Namespace,
	}

	w := Writer{}
	w.Head = headAttr.Value

	w.RsltnOfInvstgtn.Head = headAttr.Value

	w.RsltnOfInvstgtn.Assgnmt.Id = camt_029_001_09.Max35Text(params.CaseAssignmentId)
	w.RsltnOfInvstgtn.Assgnmt.Assgnr.Agt.FinInstnId.ClrSysMmbId.MmbId = camt_029_001_09.Max35Text(params.AssignorAgentMemberId)
	w.RsltnOfInvstgtn.Assgnmt.Assgne.Agt.FinInstnId.ClrSysMmbId.MmbId = camt_029_001_09.Max35Text(params.AssigneeAgentMemberId)
	w.RsltnOfInvstgtn.Assgnmt.CreDtTm = dt.ISODateTime(params.CreatedOn)

	stsConfCd := camt_029_001_09.ExternalInvestigationExecutionConfirmation1Code(params.StsConfCd)
	w.RsltnOfInvstgtn.Sts.Conf = &stsConfCd

	w.RsltnOfInvstgtn.CxlDtls.OrgnlGrpInfAndSts.RslvdCase.Id = camt_029_001_09.Max35Text(params.CaseAssignmentId)
	w.RsltnOfInvstgtn.CxlDtls.OrgnlGrpInfAndSts.RslvdCase.Cretr.Agt.FinInstnId.ClrSysMmbId.MmbId = camt_029_001_09.Max35Text(params.AssignorAgentMemberId)
	w.RsltnOfInvstgtn.CxlDtls.OrgnlGrpInfAndSts.OrgnlMsgId = camt_029_001_09.Max35Text(params.OriginalMessageId)
	w.RsltnOfInvstgtn.CxlDtls.OrgnlGrpInfAndSts.OrgnlMsgNmId = camt_029_001_09.OrigMsgName(params.OriginalMessageNameId)
	w.RsltnOfInvstgtn.CxlDtls.OrgnlGrpInfAndSts.OrgnlCreDtTm = dt.ISODateTime(params.OriginalCreatedOn)

	return w
}
