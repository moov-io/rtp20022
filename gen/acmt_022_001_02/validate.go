// Code generated by GoComply XSD2Go for Moov; DO NOT EDIT.
// Validations for urn:iso:std:iso:20022:tech:xsd:acmt.022.001.02
package acmt_022_001_02

import (
	"github.com/moov-io/base"
	"github.com/moov-io/rtp20022/pkg/rtp"
)

// XSD ComplexType validations

func (v AccountIdentification4Choice) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "AccountIdentification4Choice"
	if v.Othr != nil {
		rtp.AddError(&errs, baseName+".Othr", v.Othr.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v BranchAndFinancialInstitutionIdentification5) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "BranchAndFinancialInstitutionIdentification5"
	rtp.AddError(&errs, baseName+".FinInstnId", v.FinInstnId.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v BranchAndFinancialInstitutionIdentification5TCH) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "BranchAndFinancialInstitutionIdentification5TCH"
	rtp.AddError(&errs, baseName+".FinInstnId", v.FinInstnId.Validate())
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

func (v DocumentTCH) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "DocumentTCH"
	rtp.AddError(&errs, baseName+".IdModAdvc", v.IdModAdvc.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v FinancialInstitutionIdentification8) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "FinancialInstitutionIdentification8"
	rtp.AddError(&errs, baseName+".ClrSysMmbId", v.ClrSysMmbId.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v FinancialInstitutionIdentification8TCH) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "FinancialInstitutionIdentification8TCH"
	rtp.AddError(&errs, baseName+".ClrSysMmbId", v.ClrSysMmbId.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v GenericAccountIdentification1) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "GenericAccountIdentification1"
	rtp.AddError(&errs, baseName+".Id", v.Id.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v IdentificationAssignment2) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "IdentificationAssignment2"
	rtp.AddError(&errs, baseName+".MsgId", v.MsgId.Validate())
	rtp.AddError(&errs, baseName+".CreDtTm", v.CreDtTm.Validate())
	rtp.AddError(&errs, baseName+".Assgnr", v.Assgnr.Validate())
	rtp.AddError(&errs, baseName+".Assgne", v.Assgne.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v IdentificationAssignment2TCH) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "IdentificationAssignment2TCH"
	rtp.AddError(&errs, baseName+".MsgId", v.MsgId.Validate())
	rtp.AddError(&errs, baseName+".CreDtTm", v.CreDtTm.Validate())
	rtp.AddError(&errs, baseName+".Assgnr", v.Assgnr.Validate())
	rtp.AddError(&errs, baseName+".Assgne", v.Assgne.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v IdentificationInformation2) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "IdentificationInformation2"
	rtp.AddError(&errs, baseName+".Acct", v.Acct.Validate())
	rtp.AddError(&errs, baseName+".Agt", v.Agt.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v IdentificationInformation2TCH) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "IdentificationInformation2TCH"
	rtp.AddError(&errs, baseName+".Acct", v.Acct.Validate())
	rtp.AddError(&errs, baseName+".Agt", v.Agt.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v IdentificationModification2) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "IdentificationModification2"
	rtp.AddError(&errs, baseName+".Id", v.Id.Validate())
	if v.OrgnlPtyAndAcctId != nil {
		rtp.AddError(&errs, baseName+".OrgnlPtyAndAcctId", v.OrgnlPtyAndAcctId.Validate())
	}
	rtp.AddError(&errs, baseName+".UpdtdPtyAndAcctId", v.UpdtdPtyAndAcctId.Validate())
	if v.AddtlInf != nil {
		rtp.AddError(&errs, baseName+".AddtlInf", v.AddtlInf.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v IdentificationModification2TCH) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "IdentificationModification2TCH"
	rtp.AddError(&errs, baseName+".Id", v.Id.Validate())
	if v.OrgnlPtyAndAcctId != nil {
		rtp.AddError(&errs, baseName+".OrgnlPtyAndAcctId", v.OrgnlPtyAndAcctId.Validate())
	}
	rtp.AddError(&errs, baseName+".UpdtdPtyAndAcctId", v.UpdtdPtyAndAcctId.Validate())
	if v.AddtlInf != nil {
		rtp.AddError(&errs, baseName+".AddtlInf", v.AddtlInf.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v IdentificationModificationAdviceV02) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "IdentificationModificationAdviceV02"
	rtp.AddError(&errs, baseName+".Assgnmt", v.Assgnmt.Validate())
	rtp.AddError(&errs, baseName+".OrgnlTxRef", v.OrgnlTxRef.Validate())
	rtp.AddError(&errs, baseName+".Mod", v.Mod.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v IdentificationModificationAdviceV02TCH) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "IdentificationModificationAdviceV02TCH"
	rtp.AddError(&errs, baseName+".Assgnmt", v.Assgnmt.Validate())
	rtp.AddError(&errs, baseName+".OrgnlTxRef", v.OrgnlTxRef.Validate())
	rtp.AddError(&errs, baseName+".Mod", v.Mod.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v OriginalTransactionReference18) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "OriginalTransactionReference18"
	rtp.AddError(&errs, baseName+".MsgId", v.MsgId.Validate())
	rtp.AddError(&errs, baseName+".MsgNmId", v.MsgNmId.Validate())
	rtp.AddError(&errs, baseName+".CreDtTm", v.CreDtTm.Validate())
	rtp.AddError(&errs, baseName+".OrgnlTx", v.OrgnlTx.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v OriginalTransactionReference18TCH) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "OriginalTransactionReference18TCH"
	rtp.AddError(&errs, baseName+".MsgId", v.MsgId.Validate())
	rtp.AddError(&errs, baseName+".MsgNmId", v.MsgNmId.Validate())
	rtp.AddError(&errs, baseName+".CreDtTm", v.CreDtTm.Validate())
	rtp.AddError(&errs, baseName+".OrgnlTx", v.OrgnlTx.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v Party12Choice) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "Party12Choice"
	if v.Agt != nil {
		rtp.AddError(&errs, baseName+".Agt", v.Agt.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v Party12ChoiceTCH) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "Party12ChoiceTCH"
	if v.Agt != nil {
		rtp.AddError(&errs, baseName+".Agt", v.Agt.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v PaymentIdentification4) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "PaymentIdentification4"
	rtp.AddError(&errs, baseName+".InstrId", v.InstrId.Validate())
	rtp.AddError(&errs, baseName+".EndToEndId", v.EndToEndId.Validate())
	rtp.AddError(&errs, baseName+".TxId", v.TxId.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v PaymentIdentification4TCH) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "PaymentIdentification4TCH"
	rtp.AddError(&errs, baseName+".InstrId", v.InstrId.Validate())
	rtp.AddError(&errs, baseName+".EndToEndId", v.EndToEndId.Validate())
	rtp.AddError(&errs, baseName+".TxId", v.TxId.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

// XSD SimpleType validations

func (v Max140Text) Validate() error {
	if err := rtp.ValidateMinLength(string(v), 1); err != nil {
		return err
	}
	if err := rtp.ValidateMaxLength(string(v), 140); err != nil {
		return err
	}
	return nil
}

func (v Max34Text) Validate() error {
	if err := rtp.ValidateMinLength(string(v), 1); err != nil {
		return err
	}
	if err := rtp.ValidateMaxLength(string(v), 34); err != nil {
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
	if err := rtp.ValidateMinLength(string(v), 9); err != nil {
		return err
	}
	if err := rtp.ValidateMaxLength(string(v), 9); err != nil {
		return err
	}
	return nil
}

func (v OrigMsgName) Validate() error {
	if err := rtp.ValidateEnumeration(string(v), "pacs.008.001.08", "pain.013.001.07"); err != nil {
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
