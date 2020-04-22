package raylib

/*
#include "raylib.h"
#include <stdlib.h>
#ifndef GO_TRACE
#define GO_TRACE

#define MAX_TRACELOG_BUFFER_SIZE   128  // Max length of one trace-log message. From raylib/utils.c

void Go_TraceLogWrapper(int logType, const char *text) {
  TraceLog(logType, text);
}

//This function is the callback that raylib will execute. As you can see, it prepares the data for the onTraceCallback which is a gofunc.
void Go_CustomCallbackHook(int logType, const char *text, va_list args) {
  char buffer[MAX_TRACELOG_BUFFER_SIZE] = { 0 };
  vsprintf (buffer, text, args);
  onTraceCallback(logType, buffer);
}

void Go_EnableCustomCallback() {
  SetTraceLogCallback(&Go_CustomCallbackHook);
}

void Go_DisableCustomCallback() {
  SetTraceLogCallback(NULL);
}


#endif
*/
import "C"
import (
	"fmt"
	"unsafe"
)

var logLevelType TraceLogType = LogInfo
var logLevelExit TraceLogType = LogError

type traceHistory struct {
	Type    TraceLogType
	Message string
}

var logHistory []traceHistory = nil

func beginTraceCapture() {

}

//traceCallback is called when there is a message to log
var traceCallback func(logType TraceLogType, text string)

//tracePanicCheck will panic if the logType is set to exit.
func tracePanicCheck(logType TraceLogType, message string) {
	//If the log is greater than the exit log, then we need to exit too.
	// (this is normally handled on the C function, but since we overridden it, we need to do it manually.)
	if logType >= logLevelExit && logLevelExit != LogNone {
		panic(message)
		//os.Exit(1)
	}
}

//SetTraceLogLevel : Set the current threshold (minimum) log level
func SetTraceLogLevel(logType TraceLogType) {
	logLevelType = logType
	C.SetTraceLogLevel(C.int(logType))
}

//SetTraceLogExit : Set the exit threshold (minimum) log level
// Set to NONE to disable this. (none is above all log levels).
func SetTraceLogExit(logType TraceLogType) {
	logLevelExit = logType
	C.SetTraceLogExit(C.int(logType))
}

//SetTraceLogCallback : Set the callback for custom logging.
func SetTraceLogCallback(callback func(logType TraceLogType, text string)) {
	if callback == nil {
		//Set the default callback. We require a default callback as the LoadXXXX errors use the lastLogMessage.
		traceCallback = defaultTraceCallback
	} else {
		//Set the callback
		traceCallback = callback
	}

	//Update if we should enable or not
	if traceCallback != nil {
		C.Go_EnableCustomCallback()
	} else {
		C.Go_DisableCustomCallback()
	}
}

//defaultTraceCallback : The default callback for raylib
func defaultTraceCallback(logType TraceLogType, text string) {
	fmt.Println("[" + logType.ToUniformedString() + "] " + text)
}

//TraceLog creates a new log with a particular type. If a custom callback for logs
// is set, then it will directly invoke it, avoiding the raylib function.
func TraceLog(logType TraceLogType, a ...interface{}) {
	line := fmt.Sprint(a...)
	if traceCallback != nil {

		//Make sure we are the correct level. Since we are doing this ourselves,
		// we need to keep track of this
		if logType < logLevelType {
			return
		}

		//We have a custom tracer, so lets just use that and avoid the C end entirely.
		//Note that since we are avoiding the C calls, we have to tracePanicCheck again.
		traceCallback(logType, line)
		tracePanicCheck(logType, line)
	} else {

		//We dont have a custom tracer yet, so we will just execute our wrapper over it.
		// We have a wrapper because cgo doesn't support variable length arguments.
		cline := C.CString(line)
		defer C.free(unsafe.Pointer(cline))
		C.Go_TraceLogWrapper(C.int(logType), cline)
	}
}

//Trace creates a trace log
func Trace(a ...interface{}) {
	TraceLog(LogTrace, a...)
}

//TraceInfo creates a info log
func TraceInfo(a ...interface{}) {
	TraceLog(LogInfo, a...)
}

//TraceWarning creates a warning log
func TraceWarning(a ...interface{}) {
	TraceLog(LogWarning, a...)
}

//TraceError creates an error log
func TraceError(a ...interface{}) {
	TraceLog(LogError, a...)
}

//TraceFatal creates an error log
func TraceFatal(a ...interface{}) {
	TraceLog(LogFatal, a...)
}

type TraceLogType int32

const (
	LogAll TraceLogType = iota
	LogTrace
	LogDebug
	LogInfo
	LogWarning
	LogError
	LogFatal
	LogNone
)

//ToString converts the TraceLogType to a string. This can be useful for custom loggers.
//seealso: ToUniformedString for a uniformed formatting.
func (v TraceLogType) ToString() string {
	switch v {
	default:
		return fmt.Sprint(v)
	case LogAll:
		return "All"
	case LogTrace:
		return "Trace"
	case LogDebug:
		return "Debug"
	case LogInfo:
		return "Info"
	case LogWarning:
		return "Warning"
	case LogError:
		return "Error"
	case LogFatal:
		return "Fatal"
	case LogNone:
		return "None"
	}
}

//ToUniformedString converts the trace log type to a string of fixed length (5) and all captials.
func (v TraceLogType) ToUniformedString() string {
	switch v {
	default:
		return fmt.Sprint(v)
	case LogAll:
		return "ALL  "
	case LogTrace:
		return "TRACE"
	case LogDebug:
		return "DEBUG"
	case LogInfo:
		return "INFO "
	case LogWarning:
		return "WARN "
	case LogError:
		return "ERROR"
	case LogFatal:
		return "FATAL"
	case LogNone:
		return "NONE"
	}
}
