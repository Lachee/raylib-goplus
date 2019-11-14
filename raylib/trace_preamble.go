package raylib

import "C"

//export onTraceCallback
func onTraceCallback(logType TraceLogType, text *C.char) {
	str := C.GoString(text)

	//Execlute our callback and panic if nessary.
	traceCallback(logType, str)
	tracePanicCheck(logType, str)
}
