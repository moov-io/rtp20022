package rtp

import (
	"errors"
	"fmt"
	"github.com/moov-io/base"
	"regexp"
	"strings"
)

func AddError(errs *base.ErrorList, err error) {
	if err != nil {
		if el, ok := err.(*base.ErrorList); ok {
			if !el.Empty() {
				errs.Add(errors.New(el.Error()))
			}
		} else {
			errs.Add(err)
		}
	}
}

func ValidatePattern(value string, regex string) error {
	pat, err := regexp.Compile(regex)
	if err != nil {
		return err
	}
	if !pat.MatchString(value) {
		return fmt.Errorf("%s fails validation with pattern %s", value, regex)
	}
	return nil
}

func ValidateEnumeration(value string, enumeration ...string) error {
	for indx := range enumeration {
		if enumeration[indx] == value {
			return nil
		}
	}
	return fmt.Errorf("%s fails enumeration validation", value)
}

func ValidateLength(value string, length int) error {
	if len(value) != length {
		return fmt.Errorf("%s fails validation with length %v != required length %v", value, len(value), length)
	}
	return nil
}

func ValidateMinLength(value string, minLength int) error {
	if len(value) < minLength {
		return fmt.Errorf("%s fails validation with length %v >= required minLength %v", value, len(value), minLength)
	}
	return nil
}

func ValidateMaxLength(value string, maxLength int) error {
	if len(value) > maxLength {
		return fmt.Errorf("%s fails validation with length %v <= required maxLength %v", value, len(value), maxLength)
	}
	return nil
}

func ValidateMinInclusive(value int, minValue int) error {
	if value <= minValue {
		return fmt.Errorf("%v fails validation as it is > minInclusive %v", value, minValue)
	}
	return nil
}

func ValidateMaxInclusive(value int, maxValue int) error {
	if value >= maxValue {
		return fmt.Errorf("%v fails validation as it is < maxInclusive %v", value, maxValue)
	}
	return nil
}

func ValidateMinExclusive(value int, minValue int) error {
	if value < minValue {
		return fmt.Errorf("%v fails validation as it is >= minExclusive %v", value, minValue)
	}
	return nil
}

func ValidateMaxExclusive(value int, maxValue int) error {
	if value > maxValue {
		return fmt.Errorf("%v fails validation as it is <= maxExclusive %v", value, maxValue)
	}
	return nil
}

func ValidateFractionDigits(value string, maxValue int) error {
	splt := strings.Split(value, ".")
	if len(splt) > 1 && len(splt[1]) > maxValue {
		return fmt.Errorf("%s fails validation with length %v <= required fractionDigits %v", value, len(splt[1]), maxValue)
	}
	return nil
}

func ValidateTotalDigits(value string, maxValue int) error {
	str := strings.ReplaceAll(value, ".", "")
	if len(str) > maxValue {
		return fmt.Errorf("%s fails validation with length %v <= required totalDigits %v", value, len(str), maxValue)
	}
	return nil
}
