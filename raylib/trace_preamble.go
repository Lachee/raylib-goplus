package raylib

import "C"

//export onTraceCallback
func onTraceCallback(logType TraceLogType, text *C.char) {
	str := C.GoString(text)
	traceCallback(logType, str)
}
