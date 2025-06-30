package errors

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestErrorCodeConstants tests that all error codes are properly defined
func TestErrorCodeConstants(t *testing.T) {
	tests := []struct {
		name string
		code Code
	}{
		// General Errors
		{"ErrCodeUnknown", ErrCodeUnknown},
		{"ErrCodeInternal", ErrCodeInternal},
		{"ErrCodeConfiguration", ErrCodeConfiguration},
		{"ErrCodeInitialization", ErrCodeInitialization},

		// Authentication/Authorization Errors
		{"ErrCodeAuth", ErrCodeAuth},
		{"ErrCodeUnauthorized", ErrCodeUnauthorized},
		{"ErrCodeTokenInvalid", ErrCodeTokenInvalid},
		{"ErrCodeTokenExpired", ErrCodeTokenExpired},
		{"ErrCodePermission", ErrCodePermission},

		// Database Errors
		{"ErrCodeDatabase", ErrCodeDatabase},
		{"ErrCodeDBConnection", ErrCodeDBConnection},
		{"ErrCodeDBQuery", ErrCodeDBQuery},
		{"ErrCodeDBDuplicate", ErrCodeDBDuplicate},
		{"ErrCodeDBNotFound", ErrCodeDBNotFound},
		{"ErrCodeDBValidation", ErrCodeDBValidation},

		// HTTP/Network Errors
		{"ErrCodeHTTP", ErrCodeHTTP},
		{"ErrCodeHTTPRequest", ErrCodeHTTPRequest},
		{"ErrCodeHTTPResponse", ErrCodeHTTPResponse},
		{"ErrCodeNetwork", ErrCodeNetwork},
		{"ErrCodeTimeout", ErrCodeTimeout},

		// Validation Errors
		{"ErrCodeValidation", ErrCodeValidation},
		{"ErrCodeInvalidInput", ErrCodeInvalidInput},
		{"ErrCodeInvalidFormat", ErrCodeInvalidFormat},
		{"ErrCodeMissingField", ErrCodeMissingField},
		{"ErrCodeInvalidState", ErrCodeInvalidState},

		// External Service Errors
		{"ErrCodeExternal", ErrCodeExternal},
		{"ErrCodeAPIError", ErrCodeAPIError},
		{"ErrCodeThirdParty", ErrCodeThirdParty},
		{"ErrCodeIntegration", ErrCodeIntegration},

		// Business Logic Errors
		{"ErrCodeBusiness", ErrCodeBusiness},
		{"ErrCodeWorkflow", ErrCodeWorkflow},
		{"ErrCodeOperation", ErrCodeOperation},
		{"ErrCodeLimit", ErrCodeLimit},

		// Resource Errors
		{"ErrCodeResource", ErrCodeResource},
		{"ErrCodeNotFound", ErrCodeNotFound},
		{"ErrCodeConflict", ErrCodeConflict},
		{"ErrCodeLocked", ErrCodeLocked},
		{"ErrCodeExhausted", ErrCodeExhausted},

		// Configuration Errors
		{"ErrCodeConfig", ErrCodeConfig},
		{"ErrCodeConfigMissing", ErrCodeConfigMissing},
		{"ErrCodeConfigInvalid", ErrCodeConfigInvalid},
		{"ErrCodeConfigType", ErrCodeConfigType},
		{"ErrCodeConfigFile", ErrCodeConfigFile},
		{"ErrCodeConfigEnvironment", ErrCodeConfigEnvironment},
		{"ErrCodeConfigOverride", ErrCodeConfigOverride},
		{"ErrCodeConfigDependency", ErrCodeConfigDependency},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Verify code is not empty
			assert.NotEmpty(t, string(tt.code))
			
			// Verify code is valid
			assert.True(t, IsValidCode(tt.code))
			
			// Verify code has description
			description := GetCodeDescription(tt.code)
			assert.NotEmpty(t, description)
			assert.NotEqual(t, "Unknown error code - please contact support", description)
		})
	}
}

// TestGetCodeDescription tests the GetCodeDescription function
func TestGetCodeDescription(t *testing.T) {
	tests := []struct {
		name     string
		code     Code
		expected string
	}{
		{
			name:     "valid code",
			code:     ErrCodeInternal,
			expected: "Internal server error - please contact support",
		},
		{
			name:     "another valid code",
			code:     ErrCodeConfigFile,
			expected: "Configuration file could not be read or parsed",
		},
		{
			name:     "valid general error code",
			code:     ErrCodeUnknown,
			expected: "Unknown or unexpected error occurred",
		},
		{
			name:     "valid database error code", 
			code:     ErrCodeDBNotFound,
			expected: "Requested record not found in database",
		},
		{
			name:     "valid config error code",
			code:     ErrCodeConfigMissing,
			expected: "Required configuration parameter is missing",
		},
		{
			name:     "invalid code",
			code:     Code("9999"),
			expected: "Unknown error code - please contact support",
		},
		{
			name:     "empty code",
			code:     Code(""),
			expected: "Unknown error code - please contact support",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetCodeDescription(tt.code)
			assert.Equal(t, tt.expected, result)
		})
	}
}

// TestIsValidCode tests the IsValidCode function
func TestIsValidCode(t *testing.T) {
	tests := []struct {
		name     string
		code     Code
		expected bool
	}{
		{
			name:     "valid code",
			code:     ErrCodeInternal,
			expected: true,
		},
		{
			name:     "another valid code",
			code:     ErrCodeConfigFile,
			expected: true,
		},
		{
			name:     "invalid code",
			code:     Code("invalid"),
			expected: false,
		},
		{
			name:     "empty code",
			code:     Code(""),
			expected: false,
		},
		{
			name:     "numeric string not in range",
			code:     Code("9999"),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsValidCode(tt.code)
			assert.Equal(t, tt.expected, result)
		})
	}
}

// TestGetAllCodes tests the GetAllCodes function
func TestGetAllCodes(t *testing.T) {
	codes := GetAllCodes()
	
	// Should return all defined codes
	assert.Greater(t, len(codes), 0)
	
	// Should contain known codes
	assert.Contains(t, codes, ErrCodeInternal)
	assert.Contains(t, codes, ErrCodeConfigFile)
	
	// All returned codes should be valid
	for _, code := range codes {
		assert.True(t, IsValidCode(code))
	}
}

// TestGetCodesByCategory tests the GetCodesByCategory function
func TestGetCodesByCategory(t *testing.T) {
	tests := []struct {
		name           string
		category       string
		expectedCodes  []Code
		shouldContain  []Code
	}{
		{
			name:          "general category",
			category:      "general",
			expectedCodes: []Code{ErrCodeUnknown, ErrCodeInternal, ErrCodeConfiguration, ErrCodeInitialization},
		},
		{
			name:          "auth category",
			category:      "auth",
			shouldContain: []Code{ErrCodeAuth, ErrCodeUnauthorized, ErrCodeTokenInvalid},
		},
		{
			name:          "authentication category",
			category:      "authentication",
			shouldContain: []Code{ErrCodeAuth, ErrCodeUnauthorized, ErrCodeTokenInvalid},
		},
		{
			name:          "database category",
			category:      "database",
			shouldContain: []Code{ErrCodeDatabase, ErrCodeDBConnection, ErrCodeDBQuery},
		},
		{
			name:          "db category",
			category:      "db",
			shouldContain: []Code{ErrCodeDatabase, ErrCodeDBConnection, ErrCodeDBQuery},
		},
		{
			name:          "http category",
			category:      "http",
			shouldContain: []Code{ErrCodeHTTP, ErrCodeHTTPRequest, ErrCodeNetwork},
		},
		{
			name:          "network category",
			category:      "network",
			shouldContain: []Code{ErrCodeHTTP, ErrCodeHTTPRequest, ErrCodeNetwork},
		},
		{
			name:          "validation category",
			category:      "validation",
			shouldContain: []Code{ErrCodeValidation, ErrCodeInvalidInput, ErrCodeMissingField},
		},
		{
			name:          "external category",
			category:      "external",
			shouldContain: []Code{ErrCodeExternal, ErrCodeAPIError, ErrCodeThirdParty},
		},
		{
			name:          "business category",
			category:      "business",
			shouldContain: []Code{ErrCodeBusiness, ErrCodeWorkflow, ErrCodeOperation},
		},
		{
			name:          "resource category",
			category:      "resource",
			shouldContain: []Code{ErrCodeResource, ErrCodeNotFound, ErrCodeConflict},
		},
		{
			name:          "config category",
			category:      "config",
			shouldContain: []Code{ErrCodeConfig, ErrCodeConfigMissing, ErrCodeConfigFile},
		},
		{
			name:          "configuration category",
			category:      "configuration",
			shouldContain: []Code{ErrCodeConfig, ErrCodeConfigMissing, ErrCodeConfigFile},
		},
		{
			name:     "invalid category",
			category: "invalid",
			shouldContain: []Code{}, // Should return empty slice
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			codes := GetCodesByCategory(tt.category)
			
			if len(tt.expectedCodes) > 0 {
				// Exact match test
				assert.ElementsMatch(t, tt.expectedCodes, codes)
			} else if len(tt.shouldContain) > 0 {
				// Contains test
				for _, expectedCode := range tt.shouldContain {
					assert.Contains(t, codes, expectedCode)
				}
			} else {
				// Should be empty for invalid categories
				assert.Empty(t, codes)
			}
			
			// All returned codes should be valid
			for _, code := range codes {
				assert.True(t, IsValidCode(code))
			}
		})
	}
}

// TestErrorCodeRanges tests that error codes follow the expected numbering scheme
func TestErrorCodeRanges(t *testing.T) {
	tests := []struct {
		name      string
		codes     []Code
		minRange  int
		maxRange  int
	}{
		{
			name:     "General Errors (1000-1099)",
			codes:    []Code{ErrCodeUnknown, ErrCodeInternal, ErrCodeConfiguration, ErrCodeInitialization},
			minRange: 1000,
			maxRange: 1099,
		},
		{
			name:     "Auth Errors (1100-1199)",
			codes:    []Code{ErrCodeAuth, ErrCodeUnauthorized, ErrCodeTokenInvalid, ErrCodeTokenExpired, ErrCodePermission},
			minRange: 1100,
			maxRange: 1199,
		},
		{
			name:     "Database Errors (1200-1299)",
			codes:    []Code{ErrCodeDatabase, ErrCodeDBConnection, ErrCodeDBQuery, ErrCodeDBDuplicate, ErrCodeDBNotFound, ErrCodeDBValidation},
			minRange: 1200,
			maxRange: 1299,
		},
		{
			name:     "Configuration Errors (1800-1899)",
			codes:    GetCodesByCategory("config"),
			minRange: 1800,
			maxRange: 1899,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, code := range tt.codes {
				// Convert code to integer for range checking
				var codeNum int
				_, err := fmt.Sscanf(string(code), "%d", &codeNum)
				require.NoError(t, err, "Code should be numeric: %s", code)
				
				assert.GreaterOrEqual(t, codeNum, tt.minRange, "Code %s should be >= %d", code, tt.minRange)
				assert.LessOrEqual(t, codeNum, tt.maxRange, "Code %s should be <= %d", code, tt.maxRange)
			}
		})
	}
}

// TestCodeDetailsCompleteness tests that all defined codes have descriptions
func TestCodeDetailsCompleteness(t *testing.T) {
	// Get all codes using reflection or by testing known codes
	allCodes := GetAllCodes()
	
	for _, code := range allCodes {
		t.Run("code_"+string(code), func(t *testing.T) {
			description := GetCodeDescription(code)
			assert.NotEmpty(t, description)
			assert.NotEqual(t, "Unknown error code - please contact support", description)
		})
	}
	
	// Verify CodeDetails map has entries
	assert.Greater(t, len(CodeDetails), 0)
} 