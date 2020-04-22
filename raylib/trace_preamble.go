package raylib

import "C"

//This function is called by Raylib when the traceCallback has been set to
// a value other than nil. It needs to be declared first, so it is in a preamble.

//export onTraceCallback
func onTraceCallback(logType TraceLogType, text *C.char) {
	logLastMessage = C.GoString(text)
	logLastType = logType

	//Execlute our callback and panic if nessary.
	traceCallback(logLastType, logLastMessage)
	tracePanicCheck(logLastType, logLastMessage)
}
