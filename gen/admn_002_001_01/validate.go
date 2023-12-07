// Code generated by GoComply XSD2Go for Moov; DO NOT EDIT.
// Validations for urn:iso:std:iso:20022:tech:xsd:admn.002.001.01
package admn_002_001_01

import (
	"github.com/moov-io/base"
	"github.com/moov-io/rtp20022/pkg/rtp"
)

// XSD ComplexType validations

func (v BranchAndFinancialInstitutionIdentification4ADMN) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "BranchAndFinancialInstitutionIdentification4ADMN"
	rtp.AddError(&errs, baseName+".FinInstnId", v.FinInstnId.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v ClearingSystemMemberIdentification2ADMN) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "ClearingSystemMemberIdentification2ADMN"
	rtp.AddError(&errs, baseName+".MmbId", v.MmbId.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v DocumentTCH) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "DocumentTCH"
	rtp.AddError(&errs, baseName+".AdmnSignOnResp", v.AdmnSignOnResp.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v FinancialInstitutionIdentification7ADMN) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "FinancialInstitutionIdentification7ADMN"
	rtp.AddError(&errs, baseName+".ClrSysMmbId", v.ClrSysMmbId.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v GrpHdr) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "GrpHdr"
	rtp.AddError(&errs, baseName+".MsgId", v.MsgId.Validate())
	rtp.AddError(&errs, baseName+".CreDtTm", v.CreDtTm.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v GrpHdrTCH) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "GrpHdrTCH"
	rtp.AddError(&errs, baseName+".MsgId", v.MsgId.Validate())
	rtp.AddError(&errs, baseName+".CreDtTm", v.CreDtTm.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v SignOnResp) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "SignOnResp"
	rtp.AddError(&errs, baseName+".OrgnlInstrId", v.OrgnlInstrId.Validate())
	rtp.AddError(&errs, baseName+".InstgAgt", v.InstgAgt.Validate())
	rtp.AddError(&errs, baseName+".InstdAgt", v.InstdAgt.Validate())
	rtp.AddError(&errs, baseName+".Sts", v.Sts.Validate())
	if v.StsRsnInf != nil {
		rtp.AddError(&errs, baseName+".StsRsnInf", v.StsRsnInf.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v SignOnResponse) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "SignOnResponse"
	rtp.AddError(&errs, baseName+".GrpHdr", v.GrpHdr.Validate())
	rtp.AddError(&errs, baseName+".SignOnResp", v.SignOnResp.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v SignOnResponseTCH) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "SignOnResponseTCH"
	rtp.AddError(&errs, baseName+".GrpHdr", v.GrpHdr.Validate())
	rtp.AddError(&errs, baseName+".SignOnResp", v.SignOnResp.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v SignOnRespTCH) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "SignOnRespTCH"
	rtp.AddError(&errs, baseName+".OrgnlInstrId", v.OrgnlInstrId.Validate())
	rtp.AddError(&errs, baseName+".InstgAgt", v.InstgAgt.Validate())
	rtp.AddError(&errs, baseName+".InstdAgt", v.InstdAgt.Validate())
	rtp.AddError(&errs, baseName+".Sts", v.Sts.Validate())
	if v.StsRsnInf != nil {
		rtp.AddError(&errs, baseName+".StsRsnInf", v.StsRsnInf.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v StatusReason6Choice) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "StatusReason6Choice"
	if v.Prtry != nil {
		rtp.AddError(&errs, baseName+".Prtry", v.Prtry.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v StatusReasonInformation8) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "StatusReasonInformation8"
	rtp.AddError(&errs, baseName+".Rsn", v.Rsn.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v StatusReasonInformation8TCH) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "StatusReasonInformation8TCH"
	rtp.AddError(&errs, baseName+".Rsn", v.Rsn.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

// XSD SimpleType validations

func (v Max35Text) Validate() error {
	if err := rtp.ValidateMinLength(string(v), 1); err != nil {
		return err
	}
	if err := rtp.ValidateMaxLength(string(v), 35); err != nil {
		return err
	}
	return nil
}

func (v Min11Max11Text) Validate() error {
	if err := rtp.ValidateMinLength(string(v), 11); err != nil {
		return err
	}
	if err := rtp.ValidateMaxLength(string(v), 11); err != nil {
		return err
	}
	return nil
}

func (v ProprietaryReasonCode) Validate() error {
	if err := rtp.ValidateEnumeration(string(v), "9946", "9948", "9964", "DS0H", "RC02"); err != nil {
		return err
	}
	return nil
}

func (v TransactionGroupStatus3CodeAdmin) Validate() error {
	if err := rtp.ValidateEnumeration(string(v), "ACTC", "RJCT"); err != nil {
		return err
	}
	return nil
}
