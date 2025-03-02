# Logger Package

A structured logging framework built on top of [logrus](https://github.com/sirupsen/logrus) providing easy-to-use logging capabilities with support for multiple log levels, structured logging, and context awareness.

## Features

- Multiple log levels (Debug, Info, Warn, Error, Fatal)
- Structured logging with fields support
- JSON and Text formatting options
- Context-aware logging
- Configurable log levels
- Error handling integration

## Installation

First, ensure you have Go installed and your Go workspace set up. Then install the required dependency:

```bash
go get github.com/sirupsen/logrus
```

## Usage

### Basic Logging

```go
import "your-module/pkg/logger"

// Simple info logging
logger.Info("Server started", nil)

// Info with fields
logger.Info("User action", logger.Fields{
    "userId": "123",
    "action": "login",
    "timestamp": time.Now(),
})

// Error logging
err := someFunction()
if err != nil {
    logger.Error("Operation failed", err, logger.Fields{
        "operation": "someFunction",
        "details": "additional context",
    })
}
```

### Log Levels

The framework supports multiple log levels:

- Debug: Detailed debugging information
- Info: General operational information
- Warn: Warning messages for potentially harmful situations
- Error: Error messages for serious problems
- Fatal: Critical errors that result in program termination

```go
// Setting log level
logger.SetLogLevel("debug") // Options: "debug", "info", "warn", "error"
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

### Error Handling

```go
err := db.Query("SELECT * FROM users")
if err != nil {
    logger.Error("Database query failed", err, logger.Fields{
        "query": "SELECT * FROM users",
        "component": "UserService",
    })
}
```

### Context-Aware Logging

```go
ctx := context.Background()
logEntry := logger.WithContext(ctx)
// Use logEntry for subsequent logging
```

### Formatting Options

The logger supports both JSON and text formatting:

```go
// Set JSON formatting (default)
logger.SetFormatter("json")

// Set text formatting
logger.SetFormatter("text")
```

## Real-World Examples

### HTTP Server Logging

```go
func handleUserRegistration(w http.ResponseWriter, r *http.Request) {
    // Start timing the request
    start := time.Now()
    
    // Log incoming request
    logger.Info("Received registration request", logger.Fields{
        "method": r.Method,
        "path": r.URL.Path,
        "remoteAddr": r.RemoteAddr,
        "userAgent": r.UserAgent(),
    })

    // Process registration...
    user, err := processRegistration(r)
    if err != nil {
        logger.Error("Registration failed", err, logger.Fields{
            "email": user.Email,
            "duration": time.Since(start).String(),
        })
        http.Error(w, "Registration failed", http.StatusBadRequest)
        return
    }

    // Log success
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
        logger.Error("Failed to begin transaction", err, nil)
        return err
    }

    // Insert user
    result, err := tx.ExecContext(ctx, 
        "INSERT INTO users (email, name, role) VALUES (?, ?, ?)",
        user.Email, user.Name, user.Role,
    )
    if err != nil {
        logger.Error("Failed to insert user", err, logger.Fields{
            "email": user.Email,
            "error_code": "DB_INSERT_ERROR",
        })
        tx.Rollback()
        return err
    }

    // Get inserted ID
    id, err := result.LastInsertId()
    if err != nil {
        logger.Warn("Could not get last insert ID", err, nil)
    } else {
        user.ID = id
    }

    if err := tx.Commit(); err != nil {
        logger.Error("Failed to commit transaction", err, logger.Fields{
            "userId": user.ID,
        })
        return err
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

    // Add trace ID to context
    ctx = context.WithValue(ctx, "traceId", uuid.New().String())
    logEntry := logger.WithContext(ctx)

    // Process job stages
    for _, stage := range job.Stages {
        stageStart := time.Now()
        
        logger.Debug("Processing job stage", logger.Fields{
            "jobId": job.ID,
            "stage": stage.Name,
            "attempt": stage.Attempts,
        })

        err := stage.Execute(ctx)
        if err != nil {
            logger.Error("Job stage failed", err, logger.Fields{
                "jobId": job.ID,
                "stage": stage.Name,
                "duration": time.Since(stageStart).String(),
                "attempt": stage.Attempts,
            })
            
            if stage.Attempts >= 3 {
                logger.Fatal("Job failed after max retries", err, logger.Fields{
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

## Default Configuration

The logger is configured with these defaults:
- JSON formatter
- Output to stdout
- Info level logging

## Best Practices

1. Use appropriate log levels:
   - Debug: For detailed debugging information
   - Info: For general operational information
   - Warn: For potentially harmful situations
   - Error: For serious problems
   - Fatal: For critical errors requiring immediate attention

2. Always include relevant context in structured fields

3. Use error logging with proper error objects

4. Include relevant metadata in your logs (timestamps, request IDs, etc.)

## Initialization

Before using the logger, initialize it with your application information:

```go
logger.Initialize(logger.LoggerConfig{
    AppName:     "your-app-name",
    AppVersion:  "1.0.0",
    Environment: "production", // or "development", "staging", etc.
})
```

This will add default fields to all log entries:
- app_name: Your application name
- app_version: Your application version
- environment: Your deployment environment

These fields will be automatically included in all log entries for better tracking and filtering. 