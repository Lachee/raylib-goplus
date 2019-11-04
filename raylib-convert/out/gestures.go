package raylib

//SetGesturesEnabled : Enable a set of gestures using flags
func SetGesturesEnabled(gestureFlags int) {
	C.SetGesturesEnabled(C.int(gestureFlags))
}

//IsGestureDetected : Check if a gesture have been detected
func IsGestureDetected(gesture int) bool {
	res := C.IsGestureDetected(C.int(gesture))
	return bool(res)
}

//GetGestureDetected : Get latest detected gesture
func GetGestureDetected() int {
	res := C.GetGestureDetected()
	return int(res)
}

//GetTouchPointsCount : Get touch points count
func GetTouchPointsCount() int {
	res := C.GetTouchPointsCount()
	return int(res)
}

//GetGestureHoldDuration : Get gesture hold time in milliseconds
func GetGestureHoldDuration() float32 {
	res := C.GetGestureHoldDuration()
	return float32(res)
}

//GetGestureDragVector : Get gesture drag vector
func GetGestureDragVector() Vector2 {
	res := C.GetGestureDragVector()
	return newVector2FromPointer(unsafe.Pointer(&res))
}

//GetGestureDragAngle : Get gesture drag angle
func GetGestureDragAngle() float32 {
	res := C.GetGestureDragAngle()
	return float32(res)
}

//GetGesturePinchVector : Get gesture pinch delta
func GetGesturePinchVector() Vector2 {
	res := C.GetGesturePinchVector()
	return newVector2FromPointer(unsafe.Pointer(&res))
}

//GetGesturePinchAngle : Get gesture pinch angle
func GetGesturePinchAngle() float32 {
	res := C.GetGesturePinchAngle()
	return float32(res)
}
