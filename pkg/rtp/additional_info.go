package rtp

import (
	"slices"
	"strings"
)

// additionalInfoRule maps vendor text in TCH pacs.002 AddtlInf onto an ISO/TCH
// reason code. TCH only allows AddtlInf for 1100, NARR, and FF02 (spec v5.0).
// Bankcore stuffs comma-separated "code:text" tokens that are not TCH codes:
// 2003 wrapper, 2005 retryable, 20008 invalid account, 5001 nested inquiry.
// Permanent rules are listed first so they win. 2005 remaps to 9912
// (ErrTemporary) so consumers that only read the network reason can retry.
type additionalInfoRule struct {
	isoCode string
	codes   []string
	needles []string
}

var additionalInfoRules = []additionalInfoRule{
	{
		isoCode: "AC03", // creditor account invalid
		codes:   []string{"20008"},
		needles: []string{"invalid account number"},
	},
	{
		isoCode: "9912", // recipient connection not available (retryable)
		codes:   []string{"2005"},
		needles: []string{"internal server error"},
	},
}

// RemapReasonFromAdditionalInfo remaps proprietary 1100 using vendor AddtlInf.
// Other reasons are returned unchanged. Permanent Bankcore rules (AC03) still
// win over retryable 2005 -> 9912.
func RemapReasonFromAdditionalInfo(reason, additionalInfo string) string {
	if strings.ToUpper(reason) != "1100" {
		return reason
	}
	if code := parseAdditionalInfoForRejection(additionalInfo); code != "" {
		return code
	}
	return reason
}

func parseAdditionalInfoForRejection(input string) string {
	lower := strings.ToLower(input)
	var vendorCodes []string
	for _, part := range strings.Split(input, ",") {
		code, _, ok := strings.Cut(strings.TrimSpace(part), ":")
		if ok {
			vendorCodes = append(vendorCodes, strings.TrimSpace(code))
		}
	}

	for _, rule := range additionalInfoRules {
		for _, code := range rule.codes {
			if slices.Contains(vendorCodes, code) {
				return rule.isoCode
			}
		}
		for _, needle := range rule.needles {
			if strings.Contains(lower, needle) {
				return rule.isoCode
			}
		}
	}

	return ""
}
