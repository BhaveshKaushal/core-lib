package logger

// Code type for standardized error codes
type Code string

// Error codes for different categories
const (
    // General Errors (1000-1099)
    CodeUnknown        Code = "1000" // Unknown or unexpected error
    CodeInternal       Code = "1001" // Internal server error
    CodeConfiguration  Code = "1002" // Configuration error
    CodeInitialization Code = "1003" // Initialization error

    // Authentication/Authorization Errors (1100-1199)
    ErrCodeAuth           Code = "1100" // General authentication error
    ErrCodeUnauthorized   Code = "1101" // Unauthorized access
    ErrCodeTokenInvalid   Code = "1102" // Invalid token
    ErrCodeTokenExpired   Code = "1103" // Expired token
    ErrCodePermission     Code = "1104" // Permission denied

    // Database Errors (1200-1299)
    ErrCodeDatabase      Code = "1200" // General database error
    ErrCodeDBConnection  Code = "1201" // Database connection error
    ErrCodeDBQuery       Code = "1202" // Database query error
    ErrCodeDBDuplicate   Code = "1203" // Duplicate entry
    ErrCodeDBNotFound    Code = "1204" // Record not found
    ErrCodeDBValidation  Code = "1205" // Database validation error

    // HTTP/Network Errors (1300-1399)
    ErrCodeHTTP         Code = "1300" // General HTTP error
    ErrCodeHTTPRequest  Code = "1301" // Invalid HTTP request
    ErrCodeHTTPResponse Code = "1302" // Invalid HTTP response
    ErrCodeNetwork      Code = "1303" // Network error
    ErrCodeTimeout      Code = "1304" // Request timeout

    // Validation Errors (1400-1499)
    ErrCodeValidation    Code = "1400" // General validation error
    ErrCodeInvalidInput  Code = "1401" // Invalid input
    ErrCodeInvalidFormat Code = "1402" // Invalid format
    ErrCodeMissingField  Code = "1403" // Missing required field
    ErrCodeInvalidState  Code = "1404" // Invalid state

    // External Service Errors (1500-1599)
    ErrCodeExternal    Code = "1500" // General external service error
    ErrCodeAPIError    Code = "1501" // External API error
    ErrCodeThirdParty  Code = "1502" // Third-party service error
    ErrCodeIntegration Code = "1503" // Integration error

    // Business Logic Errors (1600-1699)
    ErrCodeBusiness  Code = "1600" // General business logic error
    ErrCodeWorkflow  Code = "1601" // Workflow error
    ErrCodeOperation Code = "1602" // Operation error
    ErrCodeLimit     Code = "1603" // Limit exceeded

    // Resource Errors (1700-1799)
    ErrCodeResource  Code = "1700" // General resource error
    ErrCodeNotFound  Code = "1701" // Resource not found
    ErrCodeConflict  Code = "1702" // Resource conflict
    ErrCodeLocked    Code = "1703" // Resource locked
    ErrCodeExhausted Code = "1704" // Resource exhausted
)

// CodeDetails contains the description for each error code
var CodeDetails = map[Code]string{
    // General Errors
    CodeUnknown:        "Unknown or unexpected error",
    CodeInternal:       "Internal server error",
    CodeConfiguration:  "Configuration error",
    CodeInitialization: "Initialization error",

    // Authentication/Authorization Errors
    ErrCodeAuth:         "General authentication error",
    ErrCodeUnauthorized: "Unauthorized access",
    ErrCodeTokenInvalid: "Invalid token",
    ErrCodeTokenExpired: "Expired token",
    ErrCodePermission:   "Permission denied",

    // Database Errors
    ErrCodeDatabase:     "General database error",
    ErrCodeDBConnection: "Database connection error",
    ErrCodeDBQuery:      "Database query error",
    ErrCodeDBDuplicate:  "Duplicate entry",
    ErrCodeDBNotFound:   "Record not found",
    ErrCodeDBValidation: "Database validation error",

    // HTTP/Network Errors
    ErrCodeHTTP:         "General HTTP error",
    ErrCodeHTTPRequest:  "Invalid HTTP request",
    ErrCodeHTTPResponse: "Invalid HTTP response",
    ErrCodeNetwork:      "Network error",
    ErrCodeTimeout:      "Request timeout",

    // Validation Errors
    ErrCodeValidation:    "General validation error",
    ErrCodeInvalidInput:  "Invalid input",
    ErrCodeInvalidFormat: "Invalid format",
    ErrCodeMissingField:  "Missing required field",
    ErrCodeInvalidState:  "Invalid state",

    // External Service Errors
    ErrCodeExternal:    "General external service error",
    ErrCodeAPIError:    "External API error",
    ErrCodeThirdParty:  "Third-party service error",
    ErrCodeIntegration: "Integration error",

    // Business Logic Errors
    ErrCodeBusiness:  "General business logic error",
    ErrCodeWorkflow:  "Workflow error",
    ErrCodeOperation: "Operation error",
    ErrCodeLimit:     "Limit exceeded",

    // Resource Errors
    ErrCodeResource:  "General resource error",
    ErrCodeNotFound:  "Resource not found",
    ErrCodeConflict:  "Resource conflict",
    ErrCodeLocked:    "Resource locked",
    ErrCodeExhausted: "Resource exhausted",
}

// GetCodeDescription returns the description for a given error code
func GetCodeDescription(code Code) string {
    if desc, ok := CodeDetails[code]; ok {
        return desc
    }
    return "Unknown error code"
}

// IsValidCode checks if an error code is valid
func IsValidCode(code Code) bool {
    _, ok := CodeDetails[code]
    return ok
} 