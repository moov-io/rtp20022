package rtp_test

import (
	"fmt"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/moov-io/base"
	"github.com/moov-io/rtp20022/gen/pacs_008_001_08"
	"github.com/moov-io/rtp20022/pkg/rtp"
)

func TestValidateMultipleErrors(t *testing.T) {
	var errs base.ErrorList = base.ErrorList{}
	rtp.AddError(&errs, "Max35TextTCH", pacs_008_001_08.Max35TextTCH("B20230931145322200000057A11712044729M").Validate())
	require.Len(t, errs, 1)
	require.ErrorContains(t, errs.Err(), "Max35TextTCH: B20230931145322200000057A11712044729M fails validation with pattern M[0-9]{4}(((01|03|05|07|08|10|12)((0[1-9])|([1-2][0-9])|(3[0-1])))|((04|06|09|11)((0[1-9])|([1-2][0-9])|30))|((02)((0[1-9])|([1-2][0-9]))))[A-Z0-9]{11}.*")
}

func TestValidatePattern(t *testing.T) {
	pattern := `[0-9]{4}(((01|03|05|07|08|10|12)((0[1-9])|([1-2][0-9])|(3[0-1])))|((04|06|09|11)((0[1-9])|([1-2][0-9])|30))|((02)((0[1-9])|([1-2][0-9]))))((([0-1][0-9])|(2[0-3]))(([0-5][0-9])){2})[A-Z0-9]{11}.*`
	require.NoError(t, rtp.ValidatePattern("20230713145322200000057A11712044729", pattern))
	require.Error(t, rtp.ValidatePattern("20230931145322200000057A11712044729", pattern)) // invalid MMDD

	t.Run("concurrent", func(t *testing.T) {
		iterations := 1000

		var wg sync.WaitGroup
		wg.Add(iterations)

		for i := 0; i < iterations; i++ {
			go func(idx int) {
				wg.Done()

				patern := fmt.Sprintf("[0-9]{0,%d}", idx/5) // add/get regexes from cache
				require.NoError(t, rtp.ValidatePattern("4", patern))
			}(i)
		}

		wg.Wait()
	})
}

func TestValidateEnumeration(t *testing.T) {
	enumVals := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I"}
	require.NoError(t, rtp.ValidateEnumeration("A", enumVals...))
	require.Error(t, rtp.ValidateEnumeration("Z", enumVals...))
}

func TestValidateLength(t *testing.T) {
	length := 5
	require.NoError(t, rtp.ValidateLength("abcde", length))
	require.Error(t, rtp.ValidateLength("abcdef", length))
	require.NoError(t, rtp.ValidateLength("abµde", length))
	require.Error(t, rtp.ValidateLength("abµdef", length))
}

func TestValidateMinLength(t *testing.T) {
	minLength := 5
	require.NoError(t, rtp.ValidateMinLength("abcde", minLength))
	require.Error(t, rtp.ValidateMinLength("abcd", minLength))
	require.NoError(t, rtp.ValidateMinLength("abµde", minLength))
	require.Error(t, rtp.ValidateMinLength("abµd", minLength))
}

func TestValidateMaxLength(t *testing.T) {
	maxLength := 5
	require.NoError(t, rtp.ValidateMaxLength("abcde", maxLength))
	require.Error(t, rtp.ValidateMaxLength("abcdef", maxLength))
	require.NoError(t, rtp.ValidateMaxLength("abµde", maxLength))
	require.Error(t, rtp.ValidateMaxLength("abµdef", maxLength))
}

func TestValidateMinInclusive(t *testing.T) {
	minVal := 3
	require.Error(t, rtp.ValidateMinInclusive(3, minVal))
	require.NoError(t, rtp.ValidateMinInclusive(5, minVal))
	require.Error(t, rtp.ValidateMinInclusive(2, minVal))
}

func TestValidateMaxInclusive(t *testing.T) {
	maxVal := 3
	require.Error(t, rtp.ValidateMaxInclusive(3, maxVal))
	require.NoError(t, rtp.ValidateMaxInclusive(2, maxVal))
	require.Error(t, rtp.ValidateMaxInclusive(5, maxVal))
}

func TestValidateMinExclusive(t *testing.T) {
	minVal := 3
	require.NoError(t, rtp.ValidateMinExclusive(3, minVal))
	require.NoError(t, rtp.ValidateMinExclusive(5, minVal))
	require.Error(t, rtp.ValidateMinExclusive(2, minVal))
}

func TestValidateMaxExclusive(t *testing.T) {
	maxVal := 3
	require.NoError(t, rtp.ValidateMaxExclusive(3, maxVal))
	require.NoError(t, rtp.ValidateMaxExclusive(2, maxVal))
	require.Error(t, rtp.ValidateMaxExclusive(5, maxVal))
}

func TestValidateFractionDigits(t *testing.T) {
	maxVal := 5
	require.NoError(t, rtp.ValidateFractionDigits("3.1415", maxVal))
	require.NoError(t, rtp.ValidateFractionDigits("3.14159", maxVal))
	require.Error(t, rtp.ValidateFractionDigits("3.141592", maxVal))
}

func TestValidateTotalDigits(t *testing.T) {
	maxVal := 5
	require.NoError(t, rtp.ValidateTotalDigits("3.141", maxVal))
	require.NoError(t, rtp.ValidateTotalDigits("3.1415", maxVal))
	require.Error(t, rtp.ValidateTotalDigits("3.14159", maxVal))
}
