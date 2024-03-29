// Code generated by GoComply XSD2Go for Moov; DO NOT EDIT.
// Validations for urn:iso:std:iso:20022:tech:xsd:admi.004.001.02
package admi_004_001_02

import (
	"github.com/moov-io/base"
	"github.com/moov-io/rtp20022/pkg/rtp"
)

// XSD ComplexType validations

func (v DocumentTCH) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "DocumentTCH"
	rtp.AddError(&errs, baseName+".SysEvtNtfctn", v.SysEvtNtfctn.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v Event2) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "Event2"
	rtp.AddError(&errs, baseName+".EvtCd", v.EvtCd.Validate())
	for indx := range v.EvtParam {
		rtp.AddError(&errs, baseName+".EvtParam", v.EvtParam[indx].Validate())
	}
	if v.EvtDesc != nil {
		rtp.AddError(&errs, baseName+".EvtDesc", v.EvtDesc.Validate())
	}
	rtp.AddError(&errs, baseName+".EvtTm", v.EvtTm.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v Event2TCH) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "Event2TCH"
	rtp.AddError(&errs, baseName+".EvtCd", v.EvtCd.Validate())
	for indx := range v.EvtParam {
		rtp.AddError(&errs, baseName+".EvtParam", v.EvtParam[indx].Validate())
	}
	if v.EvtDesc != nil {
		rtp.AddError(&errs, baseName+".EvtDesc", v.EvtDesc.Validate())
	}
	rtp.AddError(&errs, baseName+".EvtTm", v.EvtTm.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v SystemEventNotificationV02) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "SystemEventNotificationV02"
	rtp.AddError(&errs, baseName+".EvtInf", v.EvtInf.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v SystemEventNotificationV02TCH) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "SystemEventNotificationV02TCH"
	rtp.AddError(&errs, baseName+".EvtInf", v.EvtInf.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

// XSD SimpleType validations

func (v Max1000Text) Validate() error {
	if err := rtp.ValidateMinLength(string(v), 1); err != nil {
		return err
	}
	if err := rtp.ValidateMaxLength(string(v), 1000); err != nil {
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

func (v Max4AlphaNumericText) Validate() error {
	if err := rtp.ValidatePattern(string(v), `[a-zA-Z0-9]{1,4}`); err != nil {
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
