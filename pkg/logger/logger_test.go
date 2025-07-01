package logger

import (
	"bytes"
	"testing"

	errorcodes "github.com/BhaveshKaushal/base-lib/pkg/errors"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// ... existing test code ...

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
