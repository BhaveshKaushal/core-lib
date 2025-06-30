package errors

import "strings"

// Code type represents standardized error codes for consistent error categorization
// Error codes follow a hierarchical numbering scheme for easy identification and filtering
type Code string

// Standardized error codes organized by category for consistent error handling
// Each category uses a specific number range to avoid conflicts and enable filtering
const (
	// General Errors (1000-1099) - System-level and unclassified errors
	ErrCodeUnknown        Code = "1000" // Unknown or unexpected error - fallback for unhandled cases
	ErrCodeInternal       Code = "1001" // Internal server error - unexpected system failures
	ErrCodeConfiguration  Code = "1002" // Configuration error - invalid or missing configuration
	ErrCodeInitialization Code = "1003" // Initialization error - startup and setup failures

	// Authentication/Authorization Errors (1100-1199) - Security and access control
	ErrCodeAuth         Code = "1100" // General authentication error - catch-all for auth issues
	ErrCodeUnauthorized Code = "1101" // Unauthorized access - missing or invalid credentials
	ErrCodeTokenInvalid Code = "1102" // Invalid token - malformed or corrupted tokens
	ErrCodeTokenExpired Code = "1103" // Expired token - valid but time-expired tokens
	ErrCodePermission   Code = "1104" // Permission denied - insufficient privileges

	// Database Errors (1200-1299) - Data persistence and retrieval issues
	ErrCodeDatabase     Code = "1200" // General database error - unspecified DB issues
	ErrCodeDBConnection Code = "1201" // Database connection error - connectivity problems
	ErrCodeDBQuery      Code = "1202" // Database query error - SQL syntax or execution issues
	ErrCodeDBDuplicate  Code = "1203" // Duplicate entry - unique constraint violations
	ErrCodeDBNotFound   Code = "1204" // Record not found - query returned no results
	ErrCodeDBValidation Code = "1205" // Database validation error - constraint violations

	// HTTP/Network Errors (1300-1399) - Communication and protocol issues
	ErrCodeHTTP         Code = "1300" // General HTTP error - unspecified HTTP issues
	ErrCodeHTTPRequest  Code = "1301" // Invalid HTTP request - malformed requests
	ErrCodeHTTPResponse Code = "1302" // Invalid HTTP response - unexpected response format
	ErrCodeNetwork      Code = "1303" // Network error - connectivity and routing issues
	ErrCodeTimeout      Code = "1304" // Request timeout - operations exceeding time limits

	// Validation Errors (1400-1499) - Input validation and data format issues
	ErrCodeValidation    Code = "1400" // General validation error - unspecified validation failures
	ErrCodeInvalidInput  Code = "1401" // Invalid input - malformed or incorrect input data
	ErrCodeInvalidFormat Code = "1402" // Invalid format - wrong data format or structure
	ErrCodeMissingField  Code = "1403" // Missing required field - incomplete data submissions
	ErrCodeInvalidState  Code = "1404" // Invalid state - operations in wrong system state

	// External Service Errors (1500-1599) - Third-party integration issues
	ErrCodeExternal    Code = "1500" // General external service error - unspecified external issues
	ErrCodeAPIError    Code = "1501" // External API error - third-party API failures
	ErrCodeThirdParty  Code = "1502" // Third-party service error - external service unavailable
	ErrCodeIntegration Code = "1503" // Integration error - integration setup or configuration issues

	// Business Logic Errors (1600-1699) - Application-specific logic violations
	ErrCodeBusiness  Code = "1600" // General business logic error - unspecified business rule violations
	ErrCodeWorkflow  Code = "1601" // Workflow error - process or state machine violations
	ErrCodeOperation Code = "1602" // Operation error - invalid operations or sequences
	ErrCodeLimit     Code = "1603" // Limit exceeded - rate limits, quotas, or capacity exceeded

	// Resource Errors (1700-1799) - Resource management and availability issues
	ErrCodeResource  Code = "1700" // General resource error - unspecified resource issues
	ErrCodeNotFound  Code = "1701" // Resource not found - requested resource doesn't exist
	ErrCodeConflict  Code = "1702" // Resource conflict - concurrent modification conflicts
	ErrCodeLocked    Code = "1703" // Resource locked - resource temporarily unavailable
	ErrCodeExhausted Code = "1704" // Resource exhausted - insufficient resources available

	// Configuration Errors (1800-1899) - Configuration and environment issues
	ErrCodeConfig            Code = "1800" // General configuration error - unspecified config issues
	ErrCodeConfigMissing     Code = "1801" // Missing configuration - required config parameters absent
	ErrCodeConfigInvalid     Code = "1802" // Invalid configuration value - config values out of range or format
	ErrCodeConfigType        Code = "1803" // Configuration type mismatch - wrong data type for config
	ErrCodeConfigFile        Code = "1804" // Configuration file error - file reading or parsing issues
	ErrCodeConfigEnvironment Code = "1805" // Environment configuration error - environment variable issues
	ErrCodeConfigOverride    Code = "1806" // Configuration override error - conflicting config sources
	ErrCodeConfigDependency  Code = "1807" // Configuration dependency error - missing dependent configs
)

// CodeDetails maps error codes to human-readable descriptions
// Used for error reporting, logging, and user-facing error messages
var CodeDetails = map[Code]string{
	// General Errors - System-level error descriptions
	ErrCodeUnknown:        "Unknown or unexpected error occurred",
	ErrCodeInternal:       "Internal server error - please contact support",
	ErrCodeConfiguration:  "System configuration error detected",
	ErrCodeInitialization: "Application initialization failed",

	// Authentication/Authorization Errors - Security-related descriptions
	ErrCodeAuth:         "Authentication failed - please check credentials",
	ErrCodeUnauthorized: "Access denied - authentication required",
	ErrCodeTokenInvalid: "Invalid authentication token provided",
	ErrCodeTokenExpired: "Authentication token has expired",
	ErrCodePermission:   "Insufficient permissions for this operation",

	// Database Errors - Data persistence descriptions
	ErrCodeDatabase:     "Database operation failed",
	ErrCodeDBConnection: "Unable to connect to database",
	ErrCodeDBQuery:      "Database query execution failed",
	ErrCodeDBDuplicate:  "Duplicate entry - record already exists",
	ErrCodeDBNotFound:   "Requested record not found in database",
	ErrCodeDBValidation: "Database validation constraint violated",

	// HTTP/Network Errors - Communication issue descriptions
	ErrCodeHTTP:         "HTTP request processing failed",
	ErrCodeHTTPRequest:  "Malformed or invalid HTTP request",
	ErrCodeHTTPResponse: "Invalid or unexpected HTTP response",
	ErrCodeNetwork:      "Network connectivity issue detected",
	ErrCodeTimeout:      "Operation timed out - please try again",

	// Validation Errors - Input validation descriptions
	ErrCodeValidation:    "Input validation failed",
	ErrCodeInvalidInput:  "Invalid input data provided",
	ErrCodeInvalidFormat: "Data format is incorrect or unsupported",
	ErrCodeMissingField:  "Required field is missing or empty",
	ErrCodeInvalidState:  "Operation not allowed in current state",

	// External Service Errors - Third-party integration descriptions
	ErrCodeExternal:    "External service operation failed",
	ErrCodeAPIError:    "Third-party API returned an error",
	ErrCodeThirdParty:  "Third-party service is unavailable",
	ErrCodeIntegration: "Service integration configuration error",

	// Business Logic Errors - Application logic descriptions
	ErrCodeBusiness:  "Business rule validation failed",
	ErrCodeWorkflow:  "Workflow process violation detected",
	ErrCodeOperation: "Invalid operation or sequence attempted",
	ErrCodeLimit:     "Rate limit or quota exceeded",

	// Resource Errors - Resource management descriptions
	ErrCodeResource:  "Resource operation failed",
	ErrCodeNotFound:  "Requested resource could not be found",
	ErrCodeConflict:  "Resource conflict - concurrent modification detected",
	ErrCodeLocked:    "Resource is temporarily locked or unavailable",
	ErrCodeExhausted: "Insufficient resources available",

	// Configuration Errors - Configuration issue descriptions
	ErrCodeConfig:            "Configuration error detected",
	ErrCodeConfigMissing:     "Required configuration parameter is missing",
	ErrCodeConfigInvalid:     "Configuration value is invalid or out of range",
	ErrCodeConfigType:        "Configuration parameter has wrong data type",
	ErrCodeConfigFile:        "Configuration file could not be read or parsed",
	ErrCodeConfigEnvironment: "Environment configuration variable error",
	ErrCodeConfigOverride:    "Conflicting configuration sources detected",
	ErrCodeConfigDependency:  "Missing configuration dependency",
}

// GetCodeDescription retrieves the human-readable description for an error code
// Returns a default message for unknown codes to prevent panics
func GetCodeDescription(code Code) string {
	if desc, ok := CodeDetails[code]; ok {
		return desc
	}
	return "Unknown error code - please contact support"
}

// IsValidCode validates whether an error code exists in our defined set
// Used for error code validation before logging or processing
func IsValidCode(code Code) bool {
	_, ok := CodeDetails[code]
	return ok
}

// GetAllCodes returns all defined error codes for documentation or testing
// Useful for generating error code documentation or comprehensive testing
func GetAllCodes() []Code {
	codes := make([]Code, 0, len(CodeDetails))
	for code := range CodeDetails {
		codes = append(codes, code)
	}
	return codes
}

// GetCodesByCategory returns error codes for a specific category based on number range
// Useful for filtering and categorizing errors in monitoring systems
func GetCodesByCategory(category string) []Code {
	var codes []Code
	
	switch strings.ToLower(category) {
	case "general":
		codes = []Code{ErrCodeUnknown, ErrCodeInternal, ErrCodeConfiguration, ErrCodeInitialization}
	case "auth", "authentication":
		codes = []Code{ErrCodeAuth, ErrCodeUnauthorized, ErrCodeTokenInvalid, ErrCodeTokenExpired, ErrCodePermission}
	case "database", "db":
		codes = []Code{ErrCodeDatabase, ErrCodeDBConnection, ErrCodeDBQuery, ErrCodeDBDuplicate, ErrCodeDBNotFound, ErrCodeDBValidation}
	case "http", "network":
		codes = []Code{ErrCodeHTTP, ErrCodeHTTPRequest, ErrCodeHTTPResponse, ErrCodeNetwork, ErrCodeTimeout}
	case "validation":
		codes = []Code{ErrCodeValidation, ErrCodeInvalidInput, ErrCodeInvalidFormat, ErrCodeMissingField, ErrCodeInvalidState}
	case "external":
		codes = []Code{ErrCodeExternal, ErrCodeAPIError, ErrCodeThirdParty, ErrCodeIntegration}
	case "business":
		codes = []Code{ErrCodeBusiness, ErrCodeWorkflow, ErrCodeOperation, ErrCodeLimit}
	case "resource":
		codes = []Code{ErrCodeResource, ErrCodeNotFound, ErrCodeConflict, ErrCodeLocked, ErrCodeExhausted}
	case "config", "configuration":
		codes = []Code{ErrCodeConfig, ErrCodeConfigMissing, ErrCodeConfigInvalid, ErrCodeConfigType, 
			ErrCodeConfigFile, ErrCodeConfigEnvironment, ErrCodeConfigOverride, ErrCodeConfigDependency}
	}
	
	return codes
} 