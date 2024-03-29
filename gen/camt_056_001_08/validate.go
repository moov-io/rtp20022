// Code generated by GoComply XSD2Go for Moov; DO NOT EDIT.
// Validations for urn:iso:std:iso:20022:tech:xsd:camt.056.001.08
package camt_056_001_08

import (
	"github.com/moov-io/base"
	"github.com/moov-io/rtp20022/pkg/rtp"
)

// XSD ComplexType validations

func (v ActiveOrHistoricCurrencyAndAmount) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "ActiveOrHistoricCurrencyAndAmount"

	rtp.AddError(&errs, baseName+".Ccy", v.Ccy.Validate())

	if errs.Empty() {
		return nil
	}
	return errs
}

func (v BranchAndFinancialInstitutionIdentification6) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "BranchAndFinancialInstitutionIdentification6"
	rtp.AddError(&errs, baseName+".FinInstnId", v.FinInstnId.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v BranchAndFinancialInstitutionIdentification6TCH) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "BranchAndFinancialInstitutionIdentification6TCH"
	rtp.AddError(&errs, baseName+".FinInstnId", v.FinInstnId.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v CancellationReason33Choice) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "CancellationReason33Choice"
	if v.Cd != nil {
		rtp.AddError(&errs, baseName+".Cd", v.Cd.Validate())
	}
	if v.Prtry != nil {
		rtp.AddError(&errs, baseName+".Prtry", v.Prtry.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v CancellationReason33ChoiceTCH) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "CancellationReason33ChoiceTCH"
	if v.Cd != nil {
		rtp.AddError(&errs, baseName+".Cd", v.Cd.Validate())
	}
	if v.Prtry != nil {
		rtp.AddError(&errs, baseName+".Prtry", v.Prtry.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v Case5) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "Case5"
	rtp.AddError(&errs, baseName+".Id", v.Id.Validate())
	rtp.AddError(&errs, baseName+".Cretr", v.Cretr.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v Case5TCH) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "Case5TCH"
	rtp.AddError(&errs, baseName+".Id", v.Id.Validate())
	rtp.AddError(&errs, baseName+".Cretr", v.Cretr.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v CaseAssignment5) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "CaseAssignment5"
	rtp.AddError(&errs, baseName+".Id", v.Id.Validate())
	rtp.AddError(&errs, baseName+".Assgnr", v.Assgnr.Validate())
	rtp.AddError(&errs, baseName+".Assgne", v.Assgne.Validate())
	rtp.AddError(&errs, baseName+".CreDtTm", v.CreDtTm.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v CaseAssignment5TCH) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "CaseAssignment5TCH"
	rtp.AddError(&errs, baseName+".Id", v.Id.Validate())
	rtp.AddError(&errs, baseName+".Assgnr", v.Assgnr.Validate())
	rtp.AddError(&errs, baseName+".Assgne", v.Assgne.Validate())
	rtp.AddError(&errs, baseName+".CreDtTm", v.CreDtTm.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v ClearingSystemMemberIdentification2) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "ClearingSystemMemberIdentification2"
	rtp.AddError(&errs, baseName+".MmbId", v.MmbId.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v ClearingSystemMemberIdentification2TCH) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "ClearingSystemMemberIdentification2TCH"
	rtp.AddError(&errs, baseName+".MmbId", v.MmbId.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v Contact4) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "Contact4"
	if v.PhneNb != nil {
		rtp.AddError(&errs, baseName+".PhneNb", v.PhneNb.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v DateAndPlaceOfBirth1) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "DateAndPlaceOfBirth1"
	rtp.AddError(&errs, baseName+".BirthDt", v.BirthDt.Validate())
	rtp.AddError(&errs, baseName+".CityOfBirth", v.CityOfBirth.Validate())
	rtp.AddError(&errs, baseName+".CtryOfBirth", v.CtryOfBirth.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v DocumentTCH) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "DocumentTCH"
	rtp.AddError(&errs, baseName+".FIToFIPmtCxlReq", v.FIToFIPmtCxlReq.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v FinancialInstitutionIdentification18) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "FinancialInstitutionIdentification18"
	rtp.AddError(&errs, baseName+".ClrSysMmbId", v.ClrSysMmbId.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v FinancialInstitutionIdentification18TCH) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "FinancialInstitutionIdentification18TCH"
	rtp.AddError(&errs, baseName+".ClrSysMmbId", v.ClrSysMmbId.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v FIToFIPaymentCancellationRequestV08) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "FIToFIPaymentCancellationRequestV08"
	rtp.AddError(&errs, baseName+".Assgnmt", v.Assgnmt.Validate())
	rtp.AddError(&errs, baseName+".Case", v.Case.Validate())
	rtp.AddError(&errs, baseName+".Undrlyg", v.Undrlyg.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v FIToFIPaymentCancellationRequestV08TCH) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "FIToFIPaymentCancellationRequestV08TCH"
	rtp.AddError(&errs, baseName+".Assgnmt", v.Assgnmt.Validate())
	rtp.AddError(&errs, baseName+".Case", v.Case.Validate())
	rtp.AddError(&errs, baseName+".Undrlyg", v.Undrlyg.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v GenericOrganisationIdentification1) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "GenericOrganisationIdentification1"
	rtp.AddError(&errs, baseName+".Id", v.Id.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v GenericOrganisationIdentification1TCH) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "GenericOrganisationIdentification1TCH"
	rtp.AddError(&errs, baseName+".Id", v.Id.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v OrganisationIdentification29) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "OrganisationIdentification29"
	if v.LEI != nil {
		rtp.AddError(&errs, baseName+".LEI", v.LEI.Validate())
	}
	if v.Othr != nil {
		rtp.AddError(&errs, baseName+".Othr", v.Othr.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v OrganisationIdentification29TCH) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "OrganisationIdentification29TCH"
	rtp.AddError(&errs, baseName+".Othr", v.Othr.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v OrganisationIdentification29TCH2) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "OrganisationIdentification29TCH2"
	rtp.AddError(&errs, baseName+".LEI", v.LEI.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v OriginalGroupHeader15) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "OriginalGroupHeader15"
	rtp.AddError(&errs, baseName+".OrgnlMsgId", v.OrgnlMsgId.Validate())
	rtp.AddError(&errs, baseName+".OrgnlMsgNmId", v.OrgnlMsgNmId.Validate())
	if v.OrgnlCreDtTm != nil {
		rtp.AddError(&errs, baseName+".OrgnlCreDtTm", v.OrgnlCreDtTm.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v OriginalTransactionReference28) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "OriginalTransactionReference28"
	if v.Dbtr != nil {
		rtp.AddError(&errs, baseName+".Dbtr", v.Dbtr.Validate())
	}
	if v.Cdtr != nil {
		rtp.AddError(&errs, baseName+".Cdtr", v.Cdtr.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v OriginalTransactionReference28TCH) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "OriginalTransactionReference28TCH"
	if v.Dbtr != nil {
		rtp.AddError(&errs, baseName+".Dbtr", v.Dbtr.Validate())
	}
	if v.Cdtr != nil {
		rtp.AddError(&errs, baseName+".Cdtr", v.Cdtr.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v Party38Choice) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "Party38Choice"
	if v.OrgId != nil {
		rtp.AddError(&errs, baseName+".OrgId", v.OrgId.Validate())
	}
	if v.PrvtId != nil {
		rtp.AddError(&errs, baseName+".PrvtId", v.PrvtId.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v Party38ChoiceTCH) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "Party38ChoiceTCH"
	if v.OrgId != nil {
		rtp.AddError(&errs, baseName+".OrgId", v.OrgId.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v Party38ChoiceTCH2) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "Party38ChoiceTCH2"
	if v.OrgId != nil {
		rtp.AddError(&errs, baseName+".OrgId", v.OrgId.Validate())
	}
	if v.PrvtId != nil {
		rtp.AddError(&errs, baseName+".PrvtId", v.PrvtId.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v Party40Choice) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "Party40Choice"
	if v.Pty != nil {
		rtp.AddError(&errs, baseName+".Pty", v.Pty.Validate())
	}
	if v.Agt != nil {
		rtp.AddError(&errs, baseName+".Agt", v.Agt.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v Party40ChoiceTCH) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "Party40ChoiceTCH"
	if v.Agt != nil {
		rtp.AddError(&errs, baseName+".Agt", v.Agt.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v Party40ChoiceTCH2) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "Party40ChoiceTCH2"
	if v.Pty != nil {
		rtp.AddError(&errs, baseName+".Pty", v.Pty.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v Party40ChoiceTCH3) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "Party40ChoiceTCH3"
	if v.Pty != nil {
		rtp.AddError(&errs, baseName+".Pty", v.Pty.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v PartyIdentification135) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "PartyIdentification135"
	if v.Nm != nil {
		rtp.AddError(&errs, baseName+".Nm", v.Nm.Validate())
	}
	if v.PstlAdr != nil {
		rtp.AddError(&errs, baseName+".PstlAdr", v.PstlAdr.Validate())
	}
	if v.Id != nil {
		rtp.AddError(&errs, baseName+".Id", v.Id.Validate())
	}
	if v.CtctDtls != nil {
		rtp.AddError(&errs, baseName+".CtctDtls", v.CtctDtls.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v PartyIdentification135TCH) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "PartyIdentification135TCH"
	if v.Nm != nil {
		rtp.AddError(&errs, baseName+".Nm", v.Nm.Validate())
	}
	if v.Id != nil {
		rtp.AddError(&errs, baseName+".Id", v.Id.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v PartyIdentification135TCH2) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "PartyIdentification135TCH2"
	rtp.AddError(&errs, baseName+".Nm", v.Nm.Validate())
	if v.PstlAdr != nil {
		rtp.AddError(&errs, baseName+".PstlAdr", v.PstlAdr.Validate())
	}
	if v.Id != nil {
		rtp.AddError(&errs, baseName+".Id", v.Id.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v PartyIdentification135TCH3) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "PartyIdentification135TCH3"
	rtp.AddError(&errs, baseName+".Nm", v.Nm.Validate())
	if v.PstlAdr != nil {
		rtp.AddError(&errs, baseName+".PstlAdr", v.PstlAdr.Validate())
	}
	if v.Id != nil {
		rtp.AddError(&errs, baseName+".Id", v.Id.Validate())
	}
	if v.CtctDtls != nil {
		rtp.AddError(&errs, baseName+".CtctDtls", v.CtctDtls.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v PaymentCancellationReason5) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "PaymentCancellationReason5"
	if v.Orgtr != nil {
		rtp.AddError(&errs, baseName+".Orgtr", v.Orgtr.Validate())
	}
	rtp.AddError(&errs, baseName+".Rsn", v.Rsn.Validate())
	if v.AddtlInf != nil {
		rtp.AddError(&errs, baseName+".AddtlInf", v.AddtlInf.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v PaymentCancellationReason5TCH) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "PaymentCancellationReason5TCH"
	if v.Orgtr != nil {
		rtp.AddError(&errs, baseName+".Orgtr", v.Orgtr.Validate())
	}
	rtp.AddError(&errs, baseName+".Rsn", v.Rsn.Validate())
	if v.AddtlInf != nil {
		rtp.AddError(&errs, baseName+".AddtlInf", v.AddtlInf.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v PaymentTransaction106) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "PaymentTransaction106"
	rtp.AddError(&errs, baseName+".OrgnlInstrId", v.OrgnlInstrId.Validate())
	if v.OrgnlEndToEndId != nil {
		rtp.AddError(&errs, baseName+".OrgnlEndToEndId", v.OrgnlEndToEndId.Validate())
	}
	rtp.AddError(&errs, baseName+".OrgnlTxId", v.OrgnlTxId.Validate())
	if v.OrgnlUETR != nil {
		rtp.AddError(&errs, baseName+".OrgnlUETR", v.OrgnlUETR.Validate())
	}
	rtp.AddError(&errs, baseName+".OrgnlClrSysRef", v.OrgnlClrSysRef.Validate())
	rtp.AddError(&errs, baseName+".OrgnlIntrBkSttlmAmt", v.OrgnlIntrBkSttlmAmt.Validate())
	rtp.AddError(&errs, baseName+".OrgnlIntrBkSttlmDt", v.OrgnlIntrBkSttlmDt.Validate())
	rtp.AddError(&errs, baseName+".CxlRsnInf", v.CxlRsnInf.Validate())
	if v.OrgnlTxRef != nil {
		rtp.AddError(&errs, baseName+".OrgnlTxRef", v.OrgnlTxRef.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v PaymentTransaction106TCH) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "PaymentTransaction106TCH"
	rtp.AddError(&errs, baseName+".OrgnlInstrId", v.OrgnlInstrId.Validate())
	if v.OrgnlEndToEndId != nil {
		rtp.AddError(&errs, baseName+".OrgnlEndToEndId", v.OrgnlEndToEndId.Validate())
	}
	rtp.AddError(&errs, baseName+".OrgnlTxId", v.OrgnlTxId.Validate())
	if v.OrgnlUETR != nil {
		rtp.AddError(&errs, baseName+".OrgnlUETR", v.OrgnlUETR.Validate())
	}
	rtp.AddError(&errs, baseName+".OrgnlClrSysRef", v.OrgnlClrSysRef.Validate())
	rtp.AddError(&errs, baseName+".OrgnlIntrBkSttlmAmt", v.OrgnlIntrBkSttlmAmt.Validate())
	rtp.AddError(&errs, baseName+".OrgnlIntrBkSttlmDt", v.OrgnlIntrBkSttlmDt.Validate())
	rtp.AddError(&errs, baseName+".CxlRsnInf", v.CxlRsnInf.Validate())
	if v.OrgnlTxRef != nil {
		rtp.AddError(&errs, baseName+".OrgnlTxRef", v.OrgnlTxRef.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v PersonIdentification13) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "PersonIdentification13"
	rtp.AddError(&errs, baseName+".DtAndPlcOfBirth", v.DtAndPlcOfBirth.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v PersonIdentification13TCH) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "PersonIdentification13TCH"
	rtp.AddError(&errs, baseName+".DtAndPlcOfBirth", v.DtAndPlcOfBirth.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v PostalAddress24) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "PostalAddress24"
	rtp.AddError(&errs, baseName+".StrtNm", v.StrtNm.Validate())
	if v.BldgNb != nil {
		rtp.AddError(&errs, baseName+".BldgNb", v.BldgNb.Validate())
	}
	rtp.AddError(&errs, baseName+".PstCd", v.PstCd.Validate())
	rtp.AddError(&errs, baseName+".TwnNm", v.TwnNm.Validate())
	rtp.AddError(&errs, baseName+".CtrySubDvsn", v.CtrySubDvsn.Validate())
	rtp.AddError(&errs, baseName+".Ctry", v.Ctry.Validate())
	if v.AdrLine != nil {
		rtp.AddError(&errs, baseName+".AdrLine", v.AdrLine.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v PostalAddress24TCH) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "PostalAddress24TCH"
	rtp.AddError(&errs, baseName+".StrtNm", v.StrtNm.Validate())
	if v.BldgNb != nil {
		rtp.AddError(&errs, baseName+".BldgNb", v.BldgNb.Validate())
	}
	rtp.AddError(&errs, baseName+".PstCd", v.PstCd.Validate())
	rtp.AddError(&errs, baseName+".TwnNm", v.TwnNm.Validate())
	rtp.AddError(&errs, baseName+".CtrySubDvsn", v.CtrySubDvsn.Validate())
	rtp.AddError(&errs, baseName+".Ctry", v.Ctry.Validate())
	if v.AdrLine != nil {
		rtp.AddError(&errs, baseName+".AdrLine", v.AdrLine.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v UnderlyingTransaction23) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "UnderlyingTransaction23"
	rtp.AddError(&errs, baseName+".OrgnlGrpInfAndCxl", v.OrgnlGrpInfAndCxl.Validate())
	rtp.AddError(&errs, baseName+".TxInf", v.TxInf.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v UnderlyingTransaction23TCH) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "UnderlyingTransaction23TCH"
	rtp.AddError(&errs, baseName+".OrgnlGrpInfAndCxl", v.OrgnlGrpInfAndCxl.Validate())
	rtp.AddError(&errs, baseName+".TxInf", v.TxInf.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

// XSD SimpleType validations

func (v ActiveOrHistoricCurrencyCode) Validate() error {
	if err := rtp.ValidatePattern(string(v), `[A-Z]{3,3}`); err != nil {
		return err
	}
	if err := rtp.ValidateEnumeration(string(v), "USD"); err != nil {
		return err
	}
	return nil
}

func (v CountryCode) Validate() error {
	if err := rtp.ValidatePattern(string(v), `[A-Z]{2,2}`); err != nil {
		return err
	}
	if err := rtp.ValidateEnumeration(string(v), "AD", "AE", "AF", "AG", "AI", "AL", "AM", "AO", "AQ", "AR", "AS", "AT", "AU", "AW", "AX", "AZ", "BA", "BB", "BD", "BE", "BF", "BG", "BH", "BI", "BJ", "BL", "BM", "BN", "BO", "BQ", "BR", "BS", "BT", "BV", "BW", "BY", "BZ", "CA", "CC", "CD", "CF", "CG", "CH", "CI", "CK", "CL", "CM", "CN", "CO", "CR", "CU", "CV", "CW", "CX", "CY", "CZ", "DE", "DJ", "DK", "DM", "DO", "DZ", "EC", "EE", "EG", "EH", "ER", "ES", "ET", "FI", "FJ", "FK", "FM", "FO", "FR", "GA", "GB", "GD", "GE", "GF", "GG", "GH", "GI", "GL", "GM", "GN", "GP", "GQ", "GR", "GS", "GT", "GU", "GW", "GY", "HK", "HM", "HN", "HR", "HT", "HU", "ID", "IE", "IL", "IM", "IN", "IO", "IQ", "IR", "IS", "IT", "JE", "JM", "JO", "JP", "KE", "KG", "KH", "KI", "KM", "KN", "KP", "KR", "KW", "KY", "KZ", "LA", "LB", "LC", "LI", "LK", "LR", "LS", "LT", "LU", "LV", "LY", "MA", "MC", "MD", "ME", "MF", "MG", "MH", "MK", "ML", "MM", "MN", "MO", "MP", "MQ", "MR", "MS", "MT", "MU", "MV", "MW", "MX", "MY", "MZ", "NA", "NC", "NE", "NF", "NG", "NI", "NL", "NO", "NP", "NR", "NU", "NZ", "OM", "PA", "PE", "PF", "PG", "PH", "PK", "PL", "PM", "PN", "PR", "PS", "PT", "PW", "PY", "QA", "RE", "RO", "RS", "RU", "RW", "SA", "SB", "SC", "SD", "SE", "SG", "SH", "SI", "SJ", "SK", "SL", "SM", "SN", "SO", "SR", "SS", "ST", "SV", "SX", "SY", "SZ", "TC", "TD", "TF", "TG", "TH", "TJ", "TK", "TL", "TM", "TN", "TO", "TR", "TT", "TV", "TW", "TZ", "UA", "UG", "UM", "US", "UY", "UZ", "VA", "VC", "VE", "VG", "VI", "VN", "VU", "WF", "WS", "YE", "YT", "ZA", "ZM", "ZW"); err != nil {
		return err
	}
	return nil
}

func (v ExternalCancellationReason1Code) Validate() error {
	if err := rtp.ValidateEnumeration(string(v), "AC03", "AM09", "CUST", "DS24", "DUPL", "FRAD", "FRTR", "TECH", "UPAY"); err != nil {
		return err
	}
	if err := rtp.ValidateMinLength(string(v), 1); err != nil {
		return err
	}
	if err := rtp.ValidateMaxLength(string(v), 4); err != nil {
		return err
	}
	return nil
}

func (v LEIIdentifier) Validate() error {
	if err := rtp.ValidatePattern(string(v), `[A-Z0-9]{18,18}[0-9]{2,2}`); err != nil {
		return err
	}
	return nil
}

func (v Max105Text) Validate() error {
	if err := rtp.ValidateMinLength(string(v), 1); err != nil {
		return err
	}
	if err := rtp.ValidateMaxLength(string(v), 105); err != nil {
		return err
	}
	return nil
}

func (v Max140Text) Validate() error {
	if err := rtp.ValidateMinLength(string(v), 1); err != nil {
		return err
	}
	if err := rtp.ValidateMaxLength(string(v), 140); err != nil {
		return err
	}
	return nil
}

func (v Max16Text) Validate() error {
	if err := rtp.ValidateMinLength(string(v), 1); err != nil {
		return err
	}
	if err := rtp.ValidateMaxLength(string(v), 16); err != nil {
		return err
	}
	return nil
}

func (v Max35Text) Validate() error {
	if err := rtp.ValidateMinLength(string(v), 1); err != nil {
		return err
	}
	if err := rtp.ValidateMaxLength(string(v), 35); err != nil {
		return err
	}
	return nil
}

func (v Max35TextTCH) Validate() error {
	if err := rtp.ValidatePattern(string(v), `M[0-9]{4}(((01|03|05|07|08|10|12)((0[1-9])|([1-2][0-9])|(3[0-1])))|((04|06|09|11)((0[1-9])|([1-2][0-9])|30))|((02)((0[1-9])|([1-2][0-9]))))[A-Z0-9]{11}.*`); err != nil {
		return err
	}
	if err := rtp.ValidateMinLength(string(v), 1); err != nil {
		return err
	}
	if err := rtp.ValidateMaxLength(string(v), 35); err != nil {
		return err
	}
	return nil
}

func (v Max35TextTCH2) Validate() error {
	if err := rtp.ValidateMinLength(string(v), 9); err != nil {
		return err
	}
	if err := rtp.ValidateMaxLength(string(v), 9); err != nil {
		return err
	}
	return nil
}

func (v Max35TextTCH3) Validate() error {
	if err := rtp.ValidateEnumeration(string(v), "WIAM", "WICT", "WIDP", "WIFD", "WIFT", "WITH"); err != nil {
		return err
	}
	if err := rtp.ValidateMinLength(string(v), 1); err != nil {
		return err
	}
	if err := rtp.ValidateMaxLength(string(v), 4); err != nil {
		return err
	}
	return nil
}

func (v Max70Text) Validate() error {
	if err := rtp.ValidateMinLength(string(v), 1); err != nil {
		return err
	}
	if err := rtp.ValidateMaxLength(string(v), 70); err != nil {
		return err
	}
	return nil
}

func (v OrigMsgName) Validate() error {
	if err := rtp.ValidateEnumeration(string(v), "pacs.008.001.06", "pacs.008.001.08", "pacs.009.001.08", "pain.013.001.07"); err != nil {
		return err
	}
	if err := rtp.ValidateMinLength(string(v), 1); err != nil {
		return err
	}
	if err := rtp.ValidateMaxLength(string(v), 35); err != nil {
		return err
	}
	return nil
}

func (v PhoneNumber) Validate() error {
	if err := rtp.ValidatePattern(string(v), `\+[0-9]{1,3}-[0-9()+\-]{1,30}`); err != nil {
		return err
	}
	return nil
}

func (v UUIDv4Identifier) Validate() error {
	if err := rtp.ValidatePattern(string(v), `[a-f0-9]{8}-[a-f0-9]{4}-4[a-f0-9]{3}-[89ab][a-f0-9]{3}-[a-f0-9]{12}`); err != nil {
		return err
	}
	return nil
}
