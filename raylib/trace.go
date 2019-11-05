package raylib

/*
#include "raylib.h"
#include <stdlib.h>
#ifndef GO_TRACE
#define GO_TRACE

#define MAX_TRACELOG_BUFFER_SIZE   128  // Max length of one trace-log message. From raylib/utils.c

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

var traceCallback func(logType TraceLogType, text string)

//SetTraceLogLevel : Set the current threshold (minimum) log level
func SetTraceLogLevel(logType TraceLogType) {
	C.SetTraceLogLevel(C.int(logType))
}

//SetTraceLogExit : Set the exit threshold (minimum) log level
func SetTraceLogExit(logType TraceLogType) {
	C.SetTraceLogExit(C.int(logType))
}

//SetTraceLogCallback : Set the callback for custom logging.
func SetTraceLogCallback(callback func(logType TraceLogType, text string)) {
	traceCallback = callback
	if traceCallback != nil {
		C.Go_EnableCustomCallback()
	} else {
		C.Go_DisableCustomCallback()
	}
}
