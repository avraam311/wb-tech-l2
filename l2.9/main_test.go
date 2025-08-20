package main

import (
	"testing"
)

func TestUnpackString(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		input    string
		expected string
	}{
		{"a4bc2d5e", "aaaabccddddde"},
		{"abcd", "abcd"},
		{"", ""},
		{"qwe\\4\\5", "qwe45"},
		{"qwe\\45", "qwe44444"},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			result := unpackString(tc.input)
			if result != tc.expected {
				t.Errorf("Test failed with input %s: expected '%s', got '%s'", tc.input, tc.expected, result)
			}
		})
	}
}

func TestErrorHandling(t *testing.T) {
	t.Parallel()
	errTests := []struct {
		input    string
		expected string
	}{
		{"45", errInvalidString.Error()},
	}

	for _, et := range errTests {
		t.Run(et.input, func(t *testing.T) {
			result := unpackString(et.input)
			if result != et.expected {
				t.Errorf("Expected error message for input %s: expected '%s', got '%s'", et.input, et.expected, result)
			}
		})
	}
}
