package utils_test

import (
	"testing"

	"github.com/NikhilMJagtap/bunny-cli/utils"
)

func TestValidateIP(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
		hasError bool
	}{
		{"Valid IPv4", "192.168.1.1", "192.168.1.1", false},
		{"Valid IPv6", "2001:0db8:85a3:0000:0000:8a2e:0370:7334", "2001:0db8:85a3:0000:0000:8a2e:0370:7334", false},
		{"Valid IPv4 with wildcard", "192.168.1.*", "192.168.1.*", false},
		{"Invalid IP", "256.0.0.1", "", true},
		{"Invalid IP with wildcard", "192.168.*.300", "", true},
		{"Not an IP", "example.com", "", true},
		{"Valid IPv6 with wildcard", "2001:0db8:85a3:*:0000:8a2e:0370:7334", "", true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := utils.ValidateIP(tc.input)

			if tc.hasError {
				if err == nil {
					t.Errorf("Expected an error, but got none")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if result != tc.expected {
					t.Errorf("Expected %s, but got %s", tc.expected, result)
				}
			}
		})
	}
}

func TestAddRemoveValidator(t *testing.T) {
	testCases := []struct {
		name     string
		args     []string
		hasError bool
	}{
		{"Valid add", []string{"add", "arg1", "arg2"}, false},
		{"Valid remove", []string{"remove", "arg1", "arg2"}, false},
		{"Invalid action", []string{"update", "arg1", "arg2"}, true},
		{"Empty args", []string{}, true},
		{"Only action", []string{"add"}, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := utils.AddRemoveValidator(nil, tc.args)

			if tc.hasError {
				if err == nil {
					t.Errorf("Expected an error, but got none")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
			}
		})
	}
}
