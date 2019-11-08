package raylib

type GestureType int32

const (
	GestureNone GestureType = 1 << iota
	GestureTap
	GestureDoubleTap
	GestureHold
	GestureDrag
	GestureSwipeRight
	GestureSwipeLeft
	GestureSwipeUp
	GestureSwipeDown
	GesturePinchIn
	GesturePinchOut
)
