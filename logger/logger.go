package logger

var (
    PrintInfo   func(format string, v ...interface{})
    PrintWarning func(format string, v ...interface{})
    PrintError   func(format string, v ...interface{})
    PrintDebug   func(format string, v ...interface{})
)

// Setup initializes the logging functions
func Setup(infoFunc, warningFunc, errorFunc, debugFunc func(format string, v ...interface{})) {
    PrintInfo = infoFunc
    PrintWarning = warningFunc
    PrintError = errorFunc
    PrintDebug = debugFunc
}
