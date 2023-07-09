package head_001_001_01

import (
	"encoding/xml"

	"github.com/moov-io/rtp20022/gen/xmldsig"
	"github.com/moov-io/rtp20022/pkg/rtp"
)

func NewSignature() *Sgntr {
	sgntr := &Sgntr{
		Signature: &xmldsig.Signature{
			XMLName: xml.Name{
				Space: "http://www.w3.org/2000/09/xmldsig#",
				Local: "Signature",
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
