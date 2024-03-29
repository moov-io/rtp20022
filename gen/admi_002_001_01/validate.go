// Code generated by GoComply XSD2Go for Moov; DO NOT EDIT.
// Validations for urn:iso:std:iso:20022:tech:xsd:admi.002.001.01
package admi_002_001_01

import (
	"github.com/moov-io/base"
	"github.com/moov-io/rtp20022/pkg/rtp"
)

// XSD ComplexType validations

func (v DocumentTCH) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "DocumentTCH"
	rtp.AddError(&errs, baseName+".Admi00200101", v.Admi00200101.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v MessageReference) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "MessageReference"
	rtp.AddError(&errs, baseName+".Ref", v.Ref.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v MessageReferenceTCH) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "MessageReferenceTCH"
	rtp.AddError(&errs, baseName+".Ref", v.Ref.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v MessageRejectV01) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "MessageRejectV01"
	rtp.AddError(&errs, baseName+".RltdRef", v.RltdRef.Validate())
	rtp.AddError(&errs, baseName+".Rsn", v.Rsn.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v MessageRejectV01TCH) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "MessageRejectV01TCH"
	rtp.AddError(&errs, baseName+".RltdRef", v.RltdRef.Validate())
	rtp.AddError(&errs, baseName+".Rsn", v.Rsn.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v RejectionReason2) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "RejectionReason2"
	rtp.AddError(&errs, baseName+".RjctgPtyRsn", v.RjctgPtyRsn.Validate())
	if v.AddtlData != nil {
		rtp.AddError(&errs, baseName+".AddtlData", v.AddtlData.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v RejectionReason2TCH) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "RejectionReason2TCH"
	rtp.AddError(&errs, baseName+".RjctgPtyRsn", v.RjctgPtyRsn.Validate())
	if v.AddtlData != nil {
		rtp.AddError(&errs, baseName+".AddtlData", v.AddtlData.Validate())
	}
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

func (v Max35TextTCH) Validate() error {
	if err := rtp.ValidatePattern(string(v), `[0-9]{4}(((01|03|05|07|08|10|12)((0[1-9])|([1-2][0-9])|(3[0-1])))|((04|06|09|11)((0[1-9])|([1-2][0-9])|30))|((02)((0[1-9])|([1-2][0-9]))))((([0-1][0-9])|(2[0-3]))(([0-5][0-9])){2})[A-Z0-9]{11}.*`); err != nil {
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
