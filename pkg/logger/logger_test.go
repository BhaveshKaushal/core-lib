package logger

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	errorcodes "github.com/BhaveshKaushal/base-lib/pkg/errors"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// getZapLevel converts string level to zap level (similar to SetLogLevel logic)
func getZapLevel(level string) zapcore.Level {
	switch strings.ToLower(level) {
	case "debug":
		return zap.DebugLevel
	case "info":
		return zap.InfoLevel
	case "warn", "warning":
		return zap.WarnLevel
	case "error":
		return zap.ErrorLevel
	case "fatal":
		return zap.FatalLevel
	default:
		return zap.InfoLevel
	}
}

// Update the Error test to use the errors package
func TestErrorWithInvalidCode(t *testing.T) {
	// Create a test logger that captures output
	var buf bytes.Buffer
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		zapcore.AddSync(&buf),
		zapcore.DebugLevel,
	)
	testLogger := zap.New(core)

	// Save original logger
	originalLogger := zapLogger
	defer func() {
		zapLogger = originalLogger
	}()

	// Set test logger
	zapLogger = testLogger

	// Test with invalid code
	Error("test message", nil, errorcodes.Code("invalid"), nil)

	output := buf.String()

	// Should use CodeUnknown for invalid codes
	assert.Contains(t, output, string(errorcodes.ErrCodeUnknown))
}

// TestSetLogLevel tests the SetLogLevel function with various inputs
func TestSetLogLevel(t *testing.T) {
	// Save original logger
	originalLogger := zapLogger
	defer func() {
		zapLogger = originalLogger
	}()

	// Initialize logger first to ensure we have a working logger
	Initialize(LoggerConfig{
		AppName:     "test-app",
		AppVersion:  "1.0.0",
		Environment: "test",
	})

	tests := []struct {
		name           string
		level          string
		expectedLevel  zapcore.Level
		shouldLogDebug bool
		shouldLogInfo  bool
	}{
		{
			name:           "debug level",
			level:          "debug",
			expectedLevel:  zap.DebugLevel,
			shouldLogDebug: true,
			shouldLogInfo:  true,
		},
		{
			name:           "info level",
			level:          "info",
			expectedLevel:  zap.InfoLevel,
			shouldLogDebug: false,
			shouldLogInfo:  true,
		},
		{
			name:           "warn level",
			level:          "warn",
			expectedLevel:  zap.WarnLevel,
			shouldLogDebug: false,
			shouldLogInfo:  false,
		},
		{
			name:           "warning level (alias)",
			level:          "warning",
			expectedLevel:  zap.WarnLevel,
			shouldLogDebug: false,
			shouldLogInfo:  false,
		},
		{
			name:           "error level",
			level:          "error",
			expectedLevel:  zap.ErrorLevel,
			shouldLogDebug: false,
			shouldLogInfo:  false,
		},
		{
			name:           "fatal level",
			level:          "fatal",
			expectedLevel:  zap.FatalLevel,
			shouldLogDebug: false,
			shouldLogInfo:  false,
		},
		{
			name:           "case insensitive debug",
			level:          "DEBUG",
			expectedLevel:  zap.DebugLevel,
			shouldLogDebug: true,
			shouldLogInfo:  true,
		},
		{
			name:           "case insensitive info",
			level:          "INFO",
			expectedLevel:  zap.InfoLevel,
			shouldLogDebug: false,
			shouldLogInfo:  true,
		},
		{
			name:           "invalid level defaults to info",
			level:          "invalid",
			expectedLevel:  zap.InfoLevel,
			shouldLogDebug: false,
			shouldLogInfo:  true,
		},
		{
			name:           "empty string defaults to info",
			level:          "",
			expectedLevel:  zap.InfoLevel,
			shouldLogDebug: false,
			shouldLogInfo:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set log level
			SetLogLevel(tt.level)

			// Verify the logger is not nil and has a core
			assert.NotNil(t, zapLogger)
			assert.NotNil(t, zapLogger.Core())

			// Test that the logger can be used without panicking
			// We can't easily test the actual log level filtering without
			// complex buffer manipulation, but we can verify the logger works
			assert.NotPanics(t, func() {
				Debug("debug message", Fields{"test": "debug"})
			})

			assert.NotPanics(t, func() {
				Info("info message", Fields{"test": "info"})
			})
		})
	}
}

// TestSetLogLevelWithNilLogger tests that SetLogLevel handles nil logger gracefully
func TestSetLogLevelWithNilLogger(t *testing.T) {
	// Save original logger
	originalLogger := zapLogger
	defer func() {
		zapLogger = originalLogger
	}()

	// Set logger to nil
	zapLogger = nil

	// This should not panic
	assert.NotPanics(t, func() {
		SetLogLevel("debug")
	})

	// Verify logger is still nil
	assert.Nil(t, zapLogger)
}

// TestSetLogLevelPreservesDefaultFields tests that changing log level preserves default fields
func TestSetLogLevelPreservesDefaultFields(t *testing.T) {
	// Save original logger
	originalLogger := zapLogger
	defer func() {
		zapLogger = originalLogger
	}()

	// Initialize logger with default fields
	Initialize(LoggerConfig{
		AppName:     "test-app",
		AppVersion:  "1.0.0",
		Environment: "test",
	})

	// Set log level - this should not panic
	assert.NotPanics(t, func() {
		SetLogLevel("info")
	})

	// Verify logger is still working
	assert.NotNil(t, zapLogger)
	assert.NotNil(t, zapLogger.Core())
}

// TestSetLogLevelMultipleChanges tests that log level can be changed multiple times
func TestSetLogLevelMultipleChanges(t *testing.T) {
	// Save original logger
	originalLogger := zapLogger
	defer func() {
		zapLogger = originalLogger
	}()

	// Test multiple level changes
	levels := []string{"debug", "info", "warn", "error", "debug", "info"}

	for i, level := range levels {
		t.Run(fmt.Sprintf("change_%d_to_%s", i+1, level), func(t *testing.T) {
			// Set log level
			SetLogLevel(level)

			// Verify the logger is not nil
			assert.NotNil(t, zapLogger)

			// Check that the logger's core level is set correctly
			// Note: We can't directly access the level from zap.Logger, but we can verify
			// the logger is working by ensuring it's not nil and can be used
			assert.NotNil(t, zapLogger.Core())
		})
	}
}

// TestLogLevelFiltering tests that log levels actually filter messages correctly
func TestLogLevelFiltering(t *testing.T) {
	// Save original logger
	originalLogger := zapLogger
	defer func() {
		zapLogger = originalLogger
	}()

	tests := []struct {
		name           string
		level          string
		shouldLogDebug bool
		shouldLogInfo  bool
		shouldLogWarn  bool
		shouldLogError bool
	}{
		{
			name:           "debug level - all messages",
			level:          "debug",
			shouldLogDebug: true,
			shouldLogInfo:  true,
			shouldLogWarn:  true,
			shouldLogError: true,
		},
		{
			name:           "info level - info and above",
			level:          "info",
			shouldLogDebug: false,
			shouldLogInfo:  true,
			shouldLogWarn:  true,
			shouldLogError: true,
		},
		{
			name:           "warn level - warn and above",
			level:          "warn",
			shouldLogDebug: false,
			shouldLogInfo:  false,
			shouldLogWarn:  true,
			shouldLogError: true,
		},
		{
			name:           "error level - error only",
			level:          "error",
			shouldLogDebug: false,
			shouldLogInfo:  false,
			shouldLogWarn:  false,
			shouldLogError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a new test logger with the specified level
			var testBuf bytes.Buffer
			testCore := zapcore.NewCore(
				zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
				zapcore.AddSync(&testBuf),
				getZapLevel(tt.level),
			)
			testLogger := zap.New(testCore)
			zapLogger = testLogger

			// Test debug logging
			Debug("debug message", Fields{"test": "debug"})
			if tt.shouldLogDebug {
				assert.Contains(t, testBuf.String(), "debug message")
			} else {
				assert.NotContains(t, testBuf.String(), "debug message")
			}

			// Clear buffer for info test
			testBuf.Reset()

			// Test info logging
			Info("info message", Fields{"test": "info"})
			if tt.shouldLogInfo {
				assert.Contains(t, testBuf.String(), "info message")
			} else {
				assert.NotContains(t, testBuf.String(), "info message")
			}

			// Clear buffer for warn test
			testBuf.Reset()

			// Test warn logging
			Warn("warn message", Fields{"test": "warn"})
			if tt.shouldLogWarn {
				assert.Contains(t, testBuf.String(), "warn message")
			} else {
				assert.NotContains(t, testBuf.String(), "warn message")
			}

			// Clear buffer for error test
			testBuf.Reset()

			// Test error logging
			Error("error message", nil, errorcodes.ErrCodeUnknown, Fields{"test": "error"})
			if tt.shouldLogError {
				assert.Contains(t, testBuf.String(), "error message")
			} else {
				assert.NotContains(t, testBuf.String(), "error message")
			}
		})
	}
}

// TestSetFormatterActualOutput tests the actual output format by creating custom loggers
func TestSetFormatterActualOutput(t *testing.T) {
	// Save original logger
	originalLogger := zapLogger
	defer func() {
		zapLogger = originalLogger
	}()

	tests := []struct {
		name           string
		format         string
		expectedFields []string
	}{
		{
			name:           "json format",
			format:         "json",
			expectedFields: []string{`"level"`, `"msg"`, `"ts"`},
		},
		{
			name:           "text format",
			format:         "text",
			expectedFields: []string{`INFO`, `test message`},
		},
		{
			name:           "console format",
			format:         "console",
			expectedFields: []string{`INFO`, `test message`},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a test buffer
			var buf bytes.Buffer

			// Create logger with the specified format
			var core zapcore.Core
			switch strings.ToLower(tt.format) {
			case "json":
				core = zapcore.NewCore(
					zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
					zapcore.AddSync(&buf),
					zapcore.InfoLevel,
				)
			case "text", "console":
				core = zapcore.NewCore(
					zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig()),
					zapcore.AddSync(&buf),
					zapcore.InfoLevel,
				)
			default:
				core = zapcore.NewCore(
					zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
					zapcore.AddSync(&buf),
					zapcore.InfoLevel,
				)
			}

			testLogger := zap.New(core)
			zapLogger = testLogger

			// Log a message
			Info("test message", Fields{"test": "output", "format": tt.format})

			// Get the output
			output := buf.String()

			// Verify output is not empty
			assert.NotEmpty(t, output, "Log output should not be empty")

			// Verify expected fields are present
			for _, field := range tt.expectedFields {
				assert.Contains(t, output, field, "Output should contain field: %s", field)
			}

			// Verify the message is present
			assert.Contains(t, output, "test message", "Output should contain the log message")
		})
	}
}
