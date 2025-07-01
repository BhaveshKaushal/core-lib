package logger

import (
	"context"
	"fmt"
	"strings"

	"github.com/BhaveshKaushal/base-lib/pkg/errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Global variables for logger state management
var (
	// zapLogger is the global zap logger instance used throughout the application
	zapLogger *zap.Logger

	// defaultFields contains application-level metadata that gets added to every log entry
	defaultFields = map[string]interface{}{
		"app_name":    "unknown", // Application name identifier
		"app_version": "unknown", // Application version for tracking deployments
		"environment": "local",   // Environment (dev, staging, prod) for filtering logs
	}
)

// LoggerConfig holds the configuration parameters for initializing the logger
// These values will be included in every log entry as default fields
type LoggerConfig struct {
	AppName     string // Name of the application
	AppVersion  string // Version of the application (e.g., "1.0.0")
	Environment string // Environment where the app is running (e.g., "production")
}

// createZapConfig creates a standardized zap configuration with common settings
func createZapConfig(level zapcore.Level) zap.Config {
	zapConfig := zap.NewProductionConfig()
	zapConfig.OutputPaths = []string{"stdout"}                      // Log to standard output
	zapConfig.ErrorOutputPaths = []string{"stderr"}                 // Error logs to standard error
	zapConfig.EncoderConfig.TimeKey = "timestamp"                   // Use "timestamp" as the time field key
	zapConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder // ISO8601 time format
	zapConfig.Level = zap.NewAtomicLevelAt(level)                   // Set the specified level
	return zapConfig
}

// Initialize sets up the zap logger with application-specific configuration
// This should be called once at application startup before any logging occurs
func Initialize(config LoggerConfig) {
	// Update default fields with provided configuration
	// These fields will be automatically added to every log entry
	defaultFields["app_name"] = config.AppName
	if config.AppVersion != "" {
		defaultFields["app_version"] = config.AppVersion
	}
	if config.Environment != "" {
		defaultFields["environment"] = config.Environment
	}

	// Create production-ready zap configuration
	zapConfig := createZapConfig(zap.InfoLevel) // Default to Info level

	// Build the logger instance
	var err error
	zapLogger, err = zapConfig.Build()
	if err != nil {
		// Logger initialization failure is critical - panic to prevent silent failures
		panic("Failed to initialize zap logger: " + err.Error())
	}

	// Add default fields to every log entry by creating a new logger with these fields
	fields := make([]zap.Field, 0, len(defaultFields))
	for k, v := range defaultFields {
		fields = append(fields, zap.Any(k, v))
	}
	zapLogger = zapLogger.With(fields...)
}

// addDefaultFields merges application default fields with user-provided fields
// User-provided fields take precedence over defaults if there are conflicts
func addDefaultFields(fields Fields) Fields {
	if fields == nil {
		fields = Fields{}
	}

	// Add default fields only if they don't already exist in user fields
	// This allows users to override default values if needed
	for k, v := range defaultFields {
		if _, exists := fields[k]; !exists {
			fields[k] = v
		}
	}
	return fields
}

// fieldsToZapFields converts our Fields map to zap's native field format
// This enables type-safe and efficient field handling in zap
func fieldsToZapFields(fields Fields) []zap.Field {
	if fields == nil {
		return nil
	}

	// Pre-allocate slice for better performance
	zapFields := make([]zap.Field, 0, len(fields))
	for k, v := range fields {
		// zap.Any automatically determines the best field type for the value
		zapFields = append(zapFields, zap.Any(k, v))
	}
	return zapFields
}

// SetLogLevel dynamically changes the logging level at runtime
// Valid levels: debug, info, warn/warning, error, fatal
func SetLogLevel(level string) {
	if zapLogger == nil {
		fmt.Println("Logger is not initialized")
		return // Gracefully handle uninitialized logger
	}

	// Convert string level to zap's level type
	var zapLevel zapcore.Level
	switch strings.ToLower(level) {
	case "debug":
		zapLevel = zap.DebugLevel
	case "info":
		zapLevel = zap.InfoLevel
	case "warn", "warning":
		zapLevel = zap.WarnLevel
	case "error":
		zapLevel = zap.ErrorLevel
	case "fatal":
		zapLevel = zap.FatalLevel
	default:
		zapLevel = zap.InfoLevel // Default to info for invalid levels
	}

	// Create new logger with updated level using common configuration
	zapConfig := createZapConfig(zapLevel)
	newLogger, err := zapConfig.Build()
	if err != nil {
		return // Keep existing logger if new one fails to build
	}

	// Add default fields to the new logger to preserve them
	fields := make([]zap.Field, 0, len(defaultFields))
	for k, v := range defaultFields {
		fields = append(fields, zap.Any(k, v))
	}
	newLogger = newLogger.With(fields...)

	// Replace current logger after ensuring old one is synced
	zapLogger.Sync()
	zapLogger = newLogger
}

// Fields type alias for structured logging key-value pairs
// Use this to provide context and metadata with your log entries
type Fields map[string]interface{}

// Debug logs a message at debug level with structured fields
// Debug logs are typically used for detailed diagnostic information
// Only logged when log level is set to debug
func Debug(msg string, fields Fields) {
	if zapLogger == nil {
		fmt.Println("Logger is not initialized")
		return // Gracefully handle uninitialized logger
	}
	fields = addDefaultFields(fields)
	zapLogger.Debug(msg, fieldsToZapFields(fields)...)
}

// Info logs a message at info level with structured fields
// Info logs are for general application flow and important events
func Info(msg string, fields Fields) {
	if zapLogger == nil {
		return // Gracefully handle uninitialized logger
	}
	fields = addDefaultFields(fields)
	zapLogger.Info(msg, fieldsToZapFields(fields)...)
}

// Warn logs a message at warning level with structured fields
// Warning logs indicate potential issues that don't prevent operation
func Warn(msg string, fields Fields) {
	if zapLogger == nil {
		return // Gracefully handle uninitialized logger
	}
	fields = addDefaultFields(fields)
	zapLogger.Warn(msg, fieldsToZapFields(fields)...)
}

// Error logs a message at error level with error code and structured fields
// Error logs indicate serious problems that need attention
// Automatically includes error code and description for standardized error tracking
func Error(msg string, err error, code errors.Code, fields Fields) {
	if zapLogger == nil {
		return // Gracefully handle uninitialized logger
	}

	if fields == nil {
		fields = Fields{}
	}

	// Validate and set error code - use unknown if invalid
	if !errors.IsValidCode(code) {
		code = errors.ErrCodeUnknown
	}

	// Add standardized error information to fields
	fields["code"] = code
	fields["code_description"] = errors.GetCodeDescription(code)
	if err != nil {
		fields["error"] = err.Error()
	}

	fields = addDefaultFields(fields)
	zapLogger.Error(msg, fieldsToZapFields(fields)...)
}

// Fatal logs a message at fatal level with error code and then exits the application
// Use sparingly - only for unrecoverable errors that require application termination
// Automatically includes error code and description for standardized error tracking
func Fatal(msg string, err error, code errors.Code, fields Fields) {
	if zapLogger == nil {
		return // Gracefully handle uninitialized logger
	}

	if fields == nil {
		fields = Fields{}
	}

	// Validate and set error code - use unknown if invalid
	if !errors.IsValidCode(code) {
		code = errors.ErrCodeUnknown
	}

	// Add standardized error information to fields
	fields["code"] = code
	fields["code_description"] = errors.GetCodeDescription(code)
	if err != nil {
		fields["error"] = err.Error()
	}

	fields = addDefaultFields(fields)
	// Fatal will log the message and then call os.Exit(1)
	zapLogger.Fatal(msg, fieldsToZapFields(fields)...)
}

// WithContext returns a logger that can be used with context
// Note: Zap doesn't have built-in context support like logrus
// You can extend this to extract values from context if needed
func WithContext(ctx context.Context) *zap.Logger {
	if zapLogger == nil {
		return nil
	}
	// Future enhancement: extract correlation IDs, user IDs, etc. from context
	// and add them as fields to the returned logger
	return zapLogger
}

// SetFormatter changes the log output format at runtime
// Supports "json" (structured, machine-readable) and "text"/"console" (human-readable)
func SetFormatter(format string) {
	if zapLogger == nil {
		fmt.Println("Logger is not initialized")
		return // Gracefully handle uninitialized logger
	}

	// Choose appropriate zap configuration based on desired format
	var zapConfig zap.Config
	switch strings.ToLower(format) {
	case "json":
		// Production config uses JSON encoding - better for log aggregation
		zapConfig = zap.NewProductionConfig()
	case "text", "console":
		// Development config uses console encoding - better for human reading
		zapConfig = zap.NewDevelopmentConfig()
	default:
		// Default to JSON for unknown`` formats
		zapConfig = zap.NewProductionConfig()
	}

	// Build new logger with updated format
	newLogger, err := zapConfig.Build()
	if err != nil {
		zapLogger.Error("Failed to build new logger", zap.Error(err))
		return // Keep existing logger if new one fails to build
	}

	// Replace current logger after ensuring old one is synced
	zapLogger.Sync()
	zapLogger = newLogger
}

// Sync flushes any buffered log entries to the output
// Should be called before application shutdown to ensure all logs are written
func Sync() {
	if zapLogger != nil {
		zapLogger.Sync()
	}
}

// GetLogger returns the underlying zap logger for advanced usage
// Use this when you need zap-specific functionality not exposed by this wrapper
func GetLogger() *zap.Logger {
	return zapLogger
}

// init automatically initializes the logger with default configuration
// This ensures the logger is always ready to use, even if Initialize() isn't called
func init() {
	// Initialize with default configuration
	// Applications should call Initialize() with proper config to override these defaults
	Initialize(LoggerConfig{
		AppName:     "unknown",
		AppVersion:  "unknown",
		Environment: "unknown",
	})
}
