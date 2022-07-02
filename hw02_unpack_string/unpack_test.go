package hw02unpackstring

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnpack(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{input: "a4bc2d5e", expected: "aaaabccddddde"},
		{input: "abccd", expected: "abccd"},
		{input: "abccd0", expected: "abcc"},
		{input: "", expected: ""},
		{input: "aaa0b", expected: "aab"},
		// uncomment if task with asterisk completed
		// {input: `qwe\4\5`, expected: `qwe45`},
		// {input: `qwe\45`, expected: `qwe44444`},
		// {input: `qwe\\5`, expected: `qwe\\\\\`},
		// {input: `qwe\\\3`, expected: `qwe\3`},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.input, func(t *testing.T) {
			result, err := Unpack(tc.input)
			require.NoError(t, err)
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestUnpackInvalidString(t *testing.T) {
	invalidStrings := []struct {
		input       string
		expectedErr error
	}{
		{input: "3abc", expectedErr: ErrFirstCharIsDigit},
		{input: "45", expectedErr: ErrFirstCharIsDigit},
		{input: "459887276", expectedErr: ErrFirstCharIsDigit},
		{input: "aaa10b", expectedErr: ErrTwoDigits},
		{input: "aa01ab", expectedErr: ErrTwoDigits},
	}
	for _, tc := range invalidStrings {
		_, err := Unpack(tc.input)

		require.Truef(t, errors.Is(err, tc.expectedErr), "actual error %q expected %q", err, tc.expectedErr)
	}
}
