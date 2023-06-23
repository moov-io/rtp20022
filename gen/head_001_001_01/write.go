package head_001_001_01

import "encoding/xml"

func NewSignature() *Sgntr {
	sgntr := &Sgntr{
		Signature: &Signature{
			XMLName: xml.Name{
				Space: "http://www.w3.org/2000/09/xmldsig#",
				Local: "Signature",
			},
			SignedInfo: SignedInfo{
				CanonicalizationMethod: CanonicalizationMethod{
					Algorithm: "http://www.w3.org/2001/10/xml-exc-c14n#",
				},
				SignatureMethod: SignatureMethod{
					Algorithm: "http://www.w3.org/2001/04/xmldsig-more#rsa-sha256",
				},
				References: []Reference{
					{
						URI: "",
						Transforms: Transforms{
							Transforms: []Transform{
								{
									Algorithm: "http://www.w3.org/2000/09/xmldsig#enveloped-signature",
								},
								{
									Algorithm: "http://www.w3.org/2001/10/xml-exc-c14n#",
								},
							},
						},
						DigestMethod: DigestMethod{
							Algorithm: "http://www.w3.org/2001/04/xmlenc#sha256",
						},
					},
				},
			},
			SignatureValue: SignatureValue{},
			KeyInfo:        KeyInfo{},
		},
	}
	return sgntr
}
