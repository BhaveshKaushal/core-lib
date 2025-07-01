# Logger Package

A structured logging framework built on top of [zap](https://github.com/uber-go/zap) providing high-performance, structured logging capabilities with support for multiple log levels, structured logging, and integrated error handling.

## Features

- **High-performance logging** using Uber's zap library
- Multiple log levels (Debug, Info, Warn, Error, Fatal)
- Structured logging with fields support
- JSON and Text formatting options
- **Integrated error handling** with custom error types
- Runtime log level and format configuration
- Application metadata injection
- Graceful handling of uninitialized logger

## Installation

First, ensure you have Go installed and your Go workspace set up. The logger package uses zap as its underlying logging library:

```bash
go get go.uber.org/zap
```

## Quick Start

The logger is automatically initialized when the package is imported, so you can start using it immediately:

```go
import "github.com/BhaveshKaushal/base-lib/pkg/logger"

// Simple info logging
logger.Info("Server started", nil)

// Info with structured fields
logger.Info("User action", logger.Fields{
    "userId": "123",
    "action": "login",
    "timestamp": time.Now(),
})
```

## Usage

### Basic Logging

```go
import "github.com/BhaveshKaushal/base-lib/pkg/logger"

// Simple info logging
logger.Info("Server started", nil)

// Info with fields
logger.Info("User action", logger.Fields{
    "userId": "123",
    "action": "login",
    "timestamp": time.Now(),
})

// Error logging with custom error
err := errors.NewErr(errors.ErrCodeDatabase, someError, "Database operation failed", "myapp")
if err != nil {
    logger.Error("Operation failed", err, logger.Fields{
        "operation": "someFunction",
        "details": "additional context",
    })
}
```

### Log Levels

The framework supports multiple log levels:

- **Debug**: Detailed debugging information
- **Info**: General operational information  
- **Warn**: Warning messages for potentially harmful situations
- **Error**: Error messages for serious problems (requires custom error object)
- **Fatal**: Critical errors that result in program termination (requires custom error object)

```go
// Setting log level at runtime
logger.SetLogLevel("debug") // Options: "debug", "info", "warn"/"warning", "error", "fatal"
```

### Structured Logging with Fields

```go
logger.Info("Database operation", logger.Fields{
    "operation": "insert",
    "table": "users",
    "rowsAffected": 1,
    "duration": "100ms",
})
```

### Error Handling with Custom Errors

The logger integrates with the custom error system for standardized error tracking:

```go
err := db.Query("SELECT * FROM users")
if err != nil {
    customErr := errors.NewErr(errors.ErrCodeDatabase, err, "Database query failed", "UserService")
    logger.Error("Database query failed", customErr, logger.Fields{
        "query": "SELECT * FROM users",
        "component": "UserService",
    })
}
```

### Context-Aware Logging

```go
ctx := context.Background()
logEntry := logger.WithContext(ctx)
// Use logEntry for subsequent logging (returns underlying zap logger)
```

### Formatting Options

The logger supports both JSON and text formatting:

```go
// Set JSON formatting (default, production-ready)
logger.SetFormatter("json")

// Set text formatting (human-readable, development-friendly)
logger.SetFormatter("text")
// or
logger.SetFormatter("console")
```

## Real-World Examples

### HTTP Server Logging

```go
func handleUserRegistration(w http.ResponseWriter, r *http.Request) {
    start := time.Now()
    
    logger.Info("Received registration request", logger.Fields{
        "method": r.Method,
        "path": r.URL.Path,
        "remoteAddr": r.RemoteAddr,
        "userAgent": r.UserAgent(),
    })

    user, err := processRegistration(r)
    if err != nil {
        customErr := errors.NewErr(errors.ErrCodeValidation, err, "Registration validation failed", "AuthService")
        logger.Error("Registration failed", customErr, logger.Fields{
            "email": user.Email,
            "duration": time.Since(start).String(),
        })
        http.Error(w, "Registration failed", http.StatusBadRequest)
        return
    }

    logger.Info("User registered successfully", logger.Fields{
        "userId": user.ID,
        "email": user.Email,
        "duration": time.Since(start).String(),
    })
}
```

### Database Operations

```go
func (repo *UserRepository) CreateUser(ctx context.Context, user *User) error {
    logger.Debug("Attempting to create user", logger.Fields{
        "email": user.Email,
        "role": user.Role,
    })

    tx, err := repo.db.BeginTx(ctx, nil)
    if err != nil {
        customErr := errors.NewErr(errors.ErrCodeDatabase, err, "Failed to begin transaction", "UserRepo")
        logger.Error("Failed to begin transaction", customErr, nil)
        return customErr
    }

    result, err := tx.ExecContext(ctx, 
        "INSERT INTO users (email, name, role) VALUES (?, ?, ?)",
        user.Email, user.Name, user.Role,
    )
    if err != nil {
        customErr := errors.NewErr(errors.ErrCodeDatabase, err, "Failed to insert user", "UserRepo")
        logger.Error("Failed to insert user", customErr, logger.Fields{
            "email": user.Email,
        })
        tx.Rollback()
        return customErr
    }

    id, err := result.LastInsertId()
    if err != nil {
        logger.Warn("Could not get last insert ID", logger.Fields{
            "error": err.Error(),
        })
    } else {
        user.ID = id
    }

    if err := tx.Commit(); err != nil {
        customErr := errors.NewErr(errors.ErrCodeDatabase, err, "Failed to commit transaction", "UserRepo")
        logger.Error("Failed to commit transaction", customErr, logger.Fields{
            "userId": user.ID,
        })
        return customErr
    }

    logger.Info("User created successfully", logger.Fields{
        "userId": user.ID,
        "email": user.Email,
    })
    return nil
}
```

### Background Job Processing

```go
func processBackgroundJob(ctx context.Context, job Job) {
    logger.Info("Starting background job", logger.Fields{
        "jobId": job.ID,
        "jobType": job.Type,
        "priority": job.Priority,
    })

    for _, stage := range job.Stages {
        stageStart := time.Now()
        
        logger.Debug("Processing job stage", logger.Fields{
            "jobId": job.ID,
            "stage": stage.Name,
            "attempt": stage.Attempts,
        })

        err := stage.Execute(ctx)
        if err != nil {
            customErr := errors.NewErr(errors.ErrCodeProcessing, err, "Job stage execution failed", "JobProcessor")
            logger.Error("Job stage failed", customErr, logger.Fields{
                "jobId": job.ID,
                "stage": stage.Name,
                "duration": time.Since(stageStart).String(),
                "attempt": stage.Attempts,
            })
            
            if stage.Attempts >= 3 {
                fatalErr := errors.NewErr(errors.ErrCodeProcessing, err, "Job failed after max retries", "JobProcessor")
                logger.Fatal("Job failed after max retries", fatalErr, logger.Fields{
                    "jobId": job.ID,
                    "stage": stage.Name,
                    "totalAttempts": stage.Attempts,
                })
                return
            }
            continue
        }

        logger.Info("Job stage completed", logger.Fields{
            "jobId": job.ID,
            "stage": stage.Name,
            "duration": time.Since(stageStart).String(),
        })
    }

    logger.Info("Background job completed", logger.Fields{
        "jobId": job.ID,
        "totalDuration": time.Since(job.StartTime).String(),
    })
}
```

## Configuration

### Default Configuration

The logger is automatically initialized with these defaults:
- **JSON formatter** (production-ready)
- **Info level** logging
- **Stdout** output
- **ISO8601 timestamp** format
- **Default fields**: `app_name: "unknown"`, `app_version: "unknown"`, `environment: "unknown"`

### Custom Initialization

For production applications, initialize the logger with your application information:

```go
logger.Initialize(logger.LoggerConfig{
    AppName:     "my-application",
    AppVersion:  "1.0.0",
    Environment: "production", // or "development", "staging", etc.
})
```

This will add default fields to all log entries:
- `app_name`: Your application name
- `app_version`: Your application version  
- `environment`: Your deployment environment

### Runtime Configuration

You can change log levels and formats at runtime:

```go
// Change log level
logger.SetLogLevel("debug")

// Change format
logger.SetFormatter("text")  // For development
logger.SetFormatter("json")  // For production
```

## Advanced Usage

### Accessing the Underlying Zap Logger

For advanced zap-specific functionality:

```go
zapLogger := logger.GetLogger()
// Use zap-specific features
zapLogger.WithOptions(zap.AddCaller())
```

### Flushing Buffered Logs

Ensure all logs are written before shutdown:

```go
defer logger.Sync()
```

### Context Integration

```go
ctx := context.WithValue(context.Background(), "requestId", "12345")
zapLogger := logger.WithContext(ctx)
// Future enhancement: extract values from context for logging
```

## Best Practices

1. **Use appropriate log levels**:
   - Debug: For detailed debugging information
   - Info: For general operational information
   - Warn: For potentially harmful situations
   - Error: For serious problems (always use custom error objects)
   - Fatal: For critical errors requiring immediate attention

2. **Always include relevant context** in structured fields

3. **Use custom error objects** for error and fatal logging to get standardized error tracking

4. **Include relevant metadata** in your logs (timestamps, request IDs, user IDs, etc.)

5. **Initialize with proper configuration** in production applications

6. **Use JSON format in production** for better log aggregation and parsing

7. **Use text format in development** for better human readability

## Error Integration

The logger automatically extracts and includes comprehensive error information from custom error objects:

```go
customErr := errors.NewErr(errors.ErrCodeDatabase, dbErr, "Database connection failed", "MyApp")
logger.Error("Database connection failed", customErr, logger.Fields{
    "host": "localhost",
    "port": 5432,
})
```

This produces structured output with:
- `code`: The error code (e.g., "1200" for database errors)
- `code_description`: Human-readable description of the error code
- `error_message`: The custom error message
- `error_cause`: The underlying cause error message
- `app`: The application identifier
- All your custom fields

## Performance

The logger is built on zap, which is designed for high-performance logging:
- Zero-allocation JSON encoding
- Structured logging with type safety
- Efficient field handling
- Minimal memory allocations

## Migration from Other Loggers

If migrating from other logging libraries:

1. Replace log level calls with appropriate logger functions
2. Convert log fields to `logger.Fields` type
3. **Create custom error objects** for error and fatal logging using `errors.NewErr()`
4. Initialize with your application configuration
5. Use structured fields instead of string concatenation 