//IsGestureDetected : Check if a gesture have been detected
func IsGestureDetected(gesture GestureType) bool {
	res := C.IsGestureDetected(C.int(gesture))
	return bool(res)
}
