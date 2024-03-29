package head_001_001_01

import (
	"encoding/xml"

	"github.com/moov-io/rtp20022/gen/xmldsig"
	"github.com/moov-io/rtp20022/pkg/rtp"
)

func (r BusinessApplicationHeaderV01TCH) CopyDuplicateCode() string {
	if r.CpyDplct != nil {
		return string(*r.CpyDplct)
	}
	return ""
}

func (r BusinessApplicationHeaderV01TCH) FromMemberID() string {
	if r.Fr.FIId != nil {
		return string(r.Fr.FIId.FinInstnId.ClrSysMmbId.MmbId)
	}
	return ""
}

func (r BusinessApplicationHeaderV01TCH) ToMemberID() string {
	if r.To.FIId != nil {
		return string(r.To.FIId.FinInstnId.ClrSysMmbId.MmbId)
	}
	return ""
}

func NewSignature() *Sgntr {
	sgntr := &Sgntr{
		Signature: &xmldsig.Signature{
			Xmlns: []xml.Attr{
				{
					Name: xml.Name{
						Local: "xmlns:ds",
					},
					Value: "http://www.w3.org/2000/09/xmldsig#",
				},
			},
			SignedInfo: xmldsig.SignedInfoType{
				CanonicalizationMethod: xmldsig.CanonicalizationMethodType{
					Algorithm: "http://www.w3.org/2001/10/xml-exc-c14n#",
				},
				SignatureMethod: xmldsig.SignatureMethodType{
					Algorithm: "http://www.w3.org/2001/04/xmldsig-more#rsa-sha256",
				},
				Reference: []xmldsig.ReferenceType{
					{
						URI: rtp.Ptr(""),
						Transforms: &xmldsig.TransformsType{
							Transform: []xmldsig.TransformType{
								{
									Algorithm: "http://www.w3.org/2000/09/xmldsig#enveloped-signature",
								},
								{
									Algorithm: "http://www.w3.org/2001/10/xml-exc-c14n#",
								},
							},
						},
						DigestMethod: xmldsig.DigestMethodType{
							Algorithm: "http://www.w3.org/2001/04/xmlenc#sha256",
						},
					},
				},
			},
			SignatureValue: xmldsig.SignatureValueType{},
			KeyInfo:        &xmldsig.KeyInfoType{},
		},
	}
	return sgntr
}
