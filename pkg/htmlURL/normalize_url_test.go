package htmlURL

import (
	"errors"
	"strings"
	"testing"
)

func TestNormalizeURL(t *testing.T) {
	tests := []struct {
		name     string
		inputURL string
		expected string
		err      bool
	}{
		{
			name:     "remove scheme",
			inputURL: "https://blog.boot.dev/path",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "remove scheme",
			inputURL: "https://blog.boot.dev/path/",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "remove scheme",
			inputURL: "http://blog.boot.dev/path",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "remove scheme",
			inputURL: "http://blog.boot.dev/path/",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "empty string",
			inputURL: "",
			expected: "",
			err:      true,
		},
		{
			name:     "garbage string",
			inputURL: "asdasda",
			expected: "",
			err:      true,
		},
		{
			name:     "Non-http/non-https scheme",
			inputURL: "ftp://blog.boot.dev/path",
			err:      true,
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := normalizeURL(tc.inputURL)
			if err != nil && !tc.err {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			}
			if err == nil && tc.err {
				t.Errorf("Test %v - '%s' FAIL: expected error but got none", i, tc.name)
				return
			}
			if actual != tc.expected {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}

func TestNormalizeURLInvalidInput(t *testing.T) {
	tests := []struct {
		name     string
		inputURL string
		expected error
	}{
		{
			name:     "Empty string",
			inputURL: "",
			expected: errors.New("invalid URL"),
		},
		{
			name:     "Invalid scheme",
			inputURL: "ftp://example.com/path/to/resource",
			expected: errors.New("URL scheme must be http or https"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := normalizeURL(tt.inputURL)
			if result != "" || (err == nil && !strings.Contains(err.Error(), tt.expected.Error())) {
				t.Errorf("normalizeURL(%q) error = %v, wantErr %v", tt.inputURL, err, tt.expected)
			}
		})
	}
}
