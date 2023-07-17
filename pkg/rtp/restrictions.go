package rtp

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"unicode/utf8"

	"github.com/moov-io/base"
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
	var strLength = utf8.RuneCountInString(value)
	if strLength != length {
		return fmt.Errorf("%s fails validation with length %v != required length %v", value, strLength, length)
	}
	return nil
}

func ValidateMinLength(value string, minLength int) error {
	var strLength = utf8.RuneCountInString(value)
	if strLength < minLength {
		return fmt.Errorf("%s fails validation with length %v >= required minLength %v", value, strLength, minLength)
	}
	return nil
}

func ValidateMaxLength(value string, maxLength int) error {
	var strLength = utf8.RuneCountInString(value)
	if strLength > maxLength {
		return fmt.Errorf("%s fails validation with length %v <= required maxLength %v", value, strLength, maxLength)
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
	if len(splt) > 1 {
		var strLength = utf8.RuneCountInString(splt[1])
		if strLength > maxValue {
			return fmt.Errorf("%s fails validation with length %v <= required fractionDigits %v", value, strLength, maxValue)
		}
	}
	return nil
}

func ValidateTotalDigits(value string, maxValue int) error {
	str := strings.ReplaceAll(value, ".", "")
	var strLength = utf8.RuneCountInString(str)
	if strLength > maxValue {
		return fmt.Errorf("%s fails validation with length %v <= required totalDigits %v", value, strLength, maxValue)
	}
	return nil
}
