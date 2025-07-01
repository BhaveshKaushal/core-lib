package errors

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewErr(t *testing.T) {
	tests := []struct {
		name        string
		code        Code
		err         error
		msg         string
		app         string
		expectedErr *Err
	}{
		{
			name: "create new error with all fields",
			code: ErrCodeDatabase,
			err:  errors.New("connection failed"),
			msg:  "Database connection error",
			app:  "testapp",
			expectedErr: &Err{
				code:    ErrCodeDatabase,
				message: "Database connection error",
				er:      errors.New("connection failed"),
				app:     "testapp",
			},
		},
		{
			name: "create new error with empty app",
			code: ErrCodeValidation,
			err:  errors.New("invalid input"),
			msg:  "Validation failed",
			app:  "",
			expectedErr: &Err{
				code:    ErrCodeValidation,
				message: "Validation failed",
				er:      errors.New("invalid input"),
				app:     "",
			},
		},
		{
			name: "create new error with nil underlying error",
			code: ErrCodeInternal,
			err:  nil,
			msg:  "Internal error",
			app:  "myapp",
			expectedErr: &Err{
				code:    ErrCodeInternal,
				message: "Internal error",
				er:      nil,
				app:     "myapp",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NewErr(tt.code, tt.err, tt.msg, tt.app)
			
			assert.NotNil(t, result)
			assert.Equal(t, tt.expectedErr.code, result.code)
			assert.Equal(t, tt.expectedErr.message, result.message)
			assert.Equal(t, tt.expectedErr.app, result.app)
			
			if tt.err != nil {
				assert.Equal(t, tt.err.Error(), result.er.Error())
			} else {
				assert.Nil(t, result.er)
			}
		})
	}
}

func TestNewErrDefault(t *testing.T) {
	tests := []struct {
		name string
		code Code
		msg  string
		app  string
	}{
		{
			name: "create default error",
			code: ErrCodeConfigMissing,
			msg:  "Configuration missing",
			app:  "config-service",
		},
		{
			name: "create default error with empty message",
			code: ErrCodeUnauthorized,
			msg:  "",
			app:  "auth-service",
		},
		{
			name: "create default error with empty app",
			code: ErrCodeTimeout,
			msg:  "Request timeout",
			app:  "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NewErrDefault(tt.code, tt.msg, tt.app)
			
			assert.NotNil(t, result)
			assert.Equal(t, tt.code, result.Code())
			assert.Equal(t, tt.msg, result.Message())
			assert.Equal(t, tt.msg, result.Error())
			assert.NotNil(t, result.Er())
		})
	}
}

func TestErr_Code(t *testing.T) {
	tests := []struct {
		name         string
		err          *Err
		expectedCode Code
	}{
		{
			name: "get code from error",
			err: &Err{
				code:    ErrCodeDatabase,
				message: "DB error",
				er:      errors.New("connection failed"),
				app:     "testapp",
			},
			expectedCode: ErrCodeDatabase,
		},
		{
			name: "get code from error with different code",
			err: &Err{
				code:    ErrCodeValidation,
				message: "Validation error",
				er:      errors.New("invalid data"),
				app:     "validator",
			},
			expectedCode: ErrCodeValidation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.err.Code()
			assert.Equal(t, tt.expectedCode, result)
		})
	}
}

func TestErr_Message(t *testing.T) {
	tests := []struct {
		name            string
		err             *Err
		expectedMessage string
	}{
		{
			name: "get message from error",
			err: &Err{
				code:    ErrCodeDatabase,
				message: "Database connection failed",
				er:      errors.New("connection timeout"),
				app:     "testapp",
			},
			expectedMessage: "Database connection failed",
		},
		{
			name: "get empty message from error",
			err: &Err{
				code:    ErrCodeValidation,
				message: "",
				er:      errors.New("validation failed"),
				app:     "validator",
			},
			expectedMessage: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.err.Message()
			assert.Equal(t, tt.expectedMessage, result)
		})
	}
}

func TestErr_Er(t *testing.T) {
	tests := []struct {
		name        string
		err         *Err
		expectedErr error
	}{
		{
			name: "get underlying error",
			err: &Err{
				code:    ErrCodeDatabase,
				message: "DB error",
				er:      errors.New("connection failed"),
				app:     "testapp",
			},
			expectedErr: errors.New("connection failed"),
		},
		{
			name: "get nil underlying error",
			err: &Err{
				code:    ErrCodeValidation,
				message: "Validation error",
				er:      nil,
				app:     "validator",
			},
			expectedErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.err.Er()
			
			if tt.expectedErr != nil {
				require.NotNil(t, result)
				assert.Equal(t, tt.expectedErr.Error(), result.Error())
			} else {
				assert.Nil(t, result)
			}
		})
	}
}

func TestErr_Error(t *testing.T) {
	tests := []struct {
		name           string
		err            *Err
		expectedString string
	}{
		{
			name: "get error string",
			err: &Err{
				code:    ErrCodeDatabase,
				message: "DB error",
				er:      errors.New("connection failed"),
				app:     "testapp",
			},
			expectedString: "connection failed",
		},
		{
			name: "get error string with different message",
			err: &Err{
				code:    ErrCodeValidation,
				message: "Validation error",
				er:      errors.New("invalid input format"),
				app:     "validator",
			},
			expectedString: "invalid input format",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.err.Error()
			assert.Equal(t, tt.expectedString, result)
		})
	}
}

func TestErr_Cause(t *testing.T) {
	tests := []struct {
		name          string
		err           *Err
		expectedCause error
	}{
		{
			name: "get cause of simple error",
			err: &Err{
				code:    ErrCodeDatabase,
				message: "DB error",
				er:      errors.New("connection failed"),
				app:     "testapp",
			},
			expectedCause: errors.New("connection failed"),
		},
		{
			name: "get cause of wrapped error",
			err: &Err{
				code:    ErrCodeExternal,
				message: "External service error",
				er:      errors.Wrap(errors.New("network timeout"), "API call failed"),
				app:     "api-client",
			},
			expectedCause: errors.New("network timeout"),
		},
		{
			name: "get cause of nil error",
			err: &Err{
				code:    ErrCodeValidation,
				message: "Validation error",
				er:      nil,
				app:     "validator",
			},
			expectedCause: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.err.Cause()
			
			if tt.expectedCause != nil {
				require.NotNil(t, result)
				assert.Equal(t, tt.expectedCause.Error(), result.Error())
			} else {
				assert.Nil(t, result)
			}
		})
	}
}

func TestErr_Wrap(t *testing.T) {
	tests := []struct {
		name           string
		err            *Err
		wrapMsg        string
		expectedResult string
	}{
		{
			name: "wrap error with message",
			err: &Err{
				code:    ErrCodeDatabase,
				message: "DB error",
				er:      errors.New("connection failed"),
				app:     "testapp",
			},
			wrapMsg:        "failed to save user",
			expectedResult: "failed to save user: connection failed",
		},
		{
			name: "wrap error with empty message",
			err: &Err{
				code:    ErrCodeValidation,
				message: "Validation error",
				er:      errors.New("invalid input"),
				app:     "validator",
			},
			wrapMsg:        "",
			expectedResult: ": invalid input",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.err.Wrap(tt.wrapMsg)
			
			assert.NotNil(t, result)
			assert.Equal(t, tt.expectedResult, result.Error())
		})
	}
}

func TestErr_InterfaceCompliance(t *testing.T) {
	t.Run("implements Error interface", func(t *testing.T) {
		err := NewErrDefault(ErrCodeDatabase, "test error", "testapp")
		
		// Test that it implements the Error interface
		var _ Error = err
		
		// Test that it implements the standard error interface
		var _ error = err
		
		// Test all interface methods
		assert.NotEmpty(t, err.Code())
		assert.NotEmpty(t, err.Message())
		assert.NotNil(t, err.Er())
		assert.NotNil(t, err.Cause())
		assert.NotEmpty(t, err.Error())
		
		wrapped := err.Wrap("wrapper message")
		assert.NotNil(t, wrapped)
	})
}

func TestErr_NilUnderlyingError(t *testing.T) {
	t.Run("handle nil underlying error gracefully", func(t *testing.T) {
		err := NewErr(ErrCodeInternal, nil, "test message", "testapp")
		
		assert.Equal(t, ErrCodeInternal, err.Code())
		assert.Equal(t, "test message", err.Message())
		assert.Nil(t, err.Er())
		assert.Nil(t, err.Cause())
		
		// Error() method should handle nil underlying error
		assert.NotPanics(t, func() {
			_ = err.Error()
		})
		
		// Wrap should handle nil underlying error
		assert.NotPanics(t, func() {
			_ = err.Wrap("wrap message")
		})
	})
}

func TestErr_ChainedErrors(t *testing.T) {
	t.Run("handle chained errors correctly", func(t *testing.T) {
		// Create a chain of errors
		originalErr := errors.New("original error")
		wrappedErr := errors.Wrap(originalErr, "first wrap")
		doubleWrappedErr := errors.Wrap(wrappedErr, "second wrap")
		
		err := NewErr(ErrCodeExternal, doubleWrappedErr, "Custom error message", "testapp")
		
		// Check that Cause() returns the original error
		cause := err.Cause()
		assert.NotNil(t, cause)
		assert.Equal(t, "original error", cause.Error())
		
		// Check that Error() returns the full wrapped message
		assert.Contains(t, err.Error(), "second wrap")
		assert.Contains(t, err.Error(), "first wrap")
		assert.Contains(t, err.Error(), "original error")
	})
} 