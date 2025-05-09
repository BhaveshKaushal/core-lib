package logger

import (
	"strings"
	"testing"
)

func TestGetCodeDescription(t *testing.T) {
	tests := []struct {
		name string
		code Code
		want string
	}{
		{
			name: "Valid general error code",
			code: CodeUnknown,
			want: "Unknown or unexpected error",
		},
		{
			name: "Valid database error code",
			code: ErrCodeDBNotFound,
			want: "Record not found",
		},
		{
			name: "Valid config error code",
			code: ErrCodeConfigMissing,
			want: "Missing configuration parameter",
		},
		{
			name: "Invalid error code",
			code: Code("9999"),
			want: "Unknown error code",
		},
		{
			name: "Empty error code",
			code: Code(""),
			want: "Unknown error code",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetCodeDescription(tt.code)
			if got != tt.want {
				t.Errorf("GetCodeDescription(%s) = %v, want %v", tt.code, got, tt.want)
			}
		})
	}
}

func TestIsValidCode(t *testing.T) {
	tests := []struct {
		name string
		code Code
		want bool
	}{
		{
			name: "Valid error code",
			code: ErrCodeConfigInvalid,
			want: true,
		},
		{
			name: "Invalid error code",
			code: Code("9999"),
			want: false,
		},
		{
			name: "Empty error code",
			code: Code(""),
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsValidCode(tt.code)
			if got != tt.want {
				t.Errorf("IsValidCode(%s) = %v, want %v", tt.code, got, tt.want)
			}
		})
	}
}

func TestErrorCodeCategorization(t *testing.T) {
	// Test that codes follow the expected categorization pattern
	categories := map[string]string{
		"10": "General",
		"11": "Authentication",
		"12": "Database",
		"13": "HTTP/Network",
		"14": "Validation",
		"15": "External Service",
		"16": "Business Logic",
		"17": "Resource",
		"18": "Configuration",
	}

	for code, desc := range CodeDetails {
		prefix := string(code)[0:2]
		category, exists := categories[prefix]
		if !exists {
			t.Errorf("Code %s doesn't match any known category prefix", code)
			continue
		}

		// Verify the description matches the category in some way
		// This is a loose check but helps catch miscategorized errors
		if !strings.Contains(strings.ToLower(desc), strings.ToLower(category)) &&
			!strings.Contains(strings.ToLower(category), "general") {
			t.Logf("Warning: Code %s with description '%s' might not match its category '%s'",
				code, desc, category)
		}
	}
}
