package errors

// contains common error across all packages

var (
	MissingAppName = NewErrDefault(1000, "Missing app name","errors")
)
