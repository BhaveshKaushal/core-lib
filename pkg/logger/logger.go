package logger

import (
	"context"
	"os"

	"github.com/sirupsen/logrus"
)

var (
	log           = logrus.New()
	defaultFields = logrus.Fields{
		"app_name":    "unknown", // Default application name
		"app_version": "unknown", // Default version
		"environment": "unknown", // Default environment
	}
)

// LoggerConfig holds the configuration for the logger
type LoggerConfig struct {
	AppName     string
	AppVersion  string
	Environment string
}

// Initialize sets up the logger with application information
func Initialize(config LoggerConfig) {
	// Update default fields
	defaultFields["app_name"] = config.AppName
	if config.AppVersion != "" {
		defaultFields["app_version"] = config.AppVersion
	}
	if config.Environment != "" {
		defaultFields["environment"] = config.Environment
	}

	// Set default configuration
	log.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02T15:04:05.000Z07:00",
	})
	log.SetOutput(os.Stdout)
	log.SetLevel(logrus.InfoLevel)
}

// addDefaultFields adds the default fields to the provided fields
func addDefaultFields(fields Fields) Fields {
	if fields == nil {
		fields = Fields{}
	}
	// Add default fields to the provided fields
	for k, v := range defaultFields {
		if _, exists := fields[k]; !exists {
			fields[k] = v
		}
	}
	return fields
}

// SetLogLevel sets the logging level
func SetLogLevel(level string) {
	switch level {
	case "debug":
		log.SetLevel(logrus.DebugLevel)
	case "info":
		log.SetLevel(logrus.InfoLevel)
	case "warn":
		log.SetLevel(logrus.WarnLevel)
	case "error":
		log.SetLevel(logrus.ErrorLevel)
	default:
		log.SetLevel(logrus.InfoLevel)
	}
}

// Fields type for structured logging
type Fields map[string]interface{}

// Debug logs a message at debug level
func Debug(msg string, fields Fields) {
	fields = addDefaultFields(fields)
	log.WithFields(logrus.Fields(fields)).Debug(msg)
}

// Info logs a message at info level
func Info(msg string, fields Fields) {
	fields = addDefaultFields(fields)
	log.WithFields(logrus.Fields(fields)).Info(msg)
}

// Warn logs a message at warn level
func Warn(msg string, err error, fields Fields) {
	fields = addDefaultFields(fields)
	if err != nil {
		fields["error"] = err.Error()
	}
	log.WithFields(logrus.Fields(fields)).Warn(msg)
}

// Error logs a message at error level with error code
func Error(msg string, err error, code Code, fields Fields) {
	if fields == nil {
		fields = Fields{}
	}

	// Validate error code
	if !IsValidCode(code) {
		code = CodeUnknown
	}

	fields["code"] = code
	fields["code_description"] = GetCodeDescription(code)
	if err != nil {
		fields["error"] = err.Error()
	}

	log.WithFields(logrus.Fields(fields)).Error(msg)
}

// Fatal logs a message at fatal level with error code and then exits
func Fatal(msg string, err error, code Code, fields Fields) {
	if fields == nil {
		fields = Fields{}
	}

	// Validate error code
	if !IsValidCode(code) {
		code = CodeUnknown
	}

	fields["code"] = code
	fields["code_description"] = GetCodeDescription(code)
	if err != nil {
		fields["error"] = err.Error()
	}

	log.WithFields(logrus.Fields(fields)).Fatal(msg)
}

// WithContext adds context fields to the logger
func WithContext(ctx context.Context) *logrus.Entry {
	return log.WithContext(ctx)
}

// SetFormatter sets the log formatter (JSON or Text)
func SetFormatter(format string) {
	switch format {
	case "json":
		log.SetFormatter(&logrus.JSONFormatter{})
	case "text":
		log.SetFormatter(&logrus.TextFormatter{
			FullTimestamp: true,
		})
	default:
		log.SetFormatter(&logrus.JSONFormatter{})
	}
}

/*import(
	log "github.com/Sirupsen/logrus"
)

func LogWarning(msessage string, err error){
   str := fmt.Sprintf("Message: %s, Error: %+v",message, err)
   log.
}*/
