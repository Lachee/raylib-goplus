package raylib

/*
//Generated 2020-04-25T14:27:29+10:00
#include "raylib.h"
#include <stdlib.h>
#include "go.h"
*/
import "C"
import "unsafe"

//SetGesturesEnabled : Enable a set of gestures using flags
func SetGesturesEnabled(gestureFlags uint32) {
	C.SetGesturesEnabled(C.uint(gestureFlags))
}

//IsGestureDetected : Check if a gesture have been detected
func IsGestureDetected(gesture GestureType) bool {
	res := C.IsGestureDetected(C.int(gesture))
	return bool(res)
}

//GetGestureDetected : Get latest detected gesture
func GetGestureDetected() GestureType {
	res := C.GetGestureDetected()
	return GestureType(res)
}

//GetTouchPointsCount : Get touch points count
func GetTouchPointsCount() int {
	res := C.GetTouchPointsCount()
	return int(int32(res))
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
