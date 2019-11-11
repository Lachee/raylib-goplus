package raylib

/*
//Generated 2019-11-11T20:13:51+11:00
#include "raylib.h"
#include <stdlib.h>
#include "go.h"
*/
import "C"
import "unsafe"

//IsGamepadAvailable : Detect if a gamepad is available
func IsGamepadAvailable(gamepad GamepadNumber) bool {
	res := C.IsGamepadAvailable(C.int(int32(gamepad)))
	return bool(res)
}

//IsGamepadName : Check gamepad name (if available)
func IsGamepadName(gamepad GamepadNumber, name string) bool {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	res := C.IsGamepadName(C.int(int32(gamepad)), cname)
	return bool(res)
}

//GetGamepadName : Return gamepad internal name id
func GetGamepadName(gamepad GamepadNumber) string {
	res := C.GetGamepadName(C.int(int32(gamepad)))
	return C.GoString(res)
}

//IsGamepadButtonPressed : Detect if a gamepad button has been pressed once
func IsGamepadButtonPressed(gamepad GamepadNumber, button GamepadButton) bool {
	res := C.IsGamepadButtonPressed(C.int(int32(gamepad)), C.int(int32(button)))
	return bool(res)
}

//IsGamepadButtonDown : Detect if a gamepad button is being pressed
func IsGamepadButtonDown(gamepad GamepadNumber, button GamepadButton) bool {
	res := C.IsGamepadButtonDown(C.int(int32(gamepad)), C.int(int32(button)))
	return bool(res)
}

//IsGamepadButtonReleased : Detect if a gamepad button has been released once
func IsGamepadButtonReleased(gamepad GamepadNumber, button GamepadButton) bool {
	res := C.IsGamepadButtonReleased(C.int(int32(gamepad)), C.int(int32(button)))
	return bool(res)
}

//IsGamepadButtonUp : Detect if a gamepad button is NOT being pressed
func IsGamepadButtonUp(gamepad GamepadNumber, button GamepadButton) bool {
	res := C.IsGamepadButtonUp(C.int(int32(gamepad)), C.int(int32(button)))
	return bool(res)
}

//GetGamepadButtonPressed : Get the last gamepad button pressed
func GetGamepadButtonPressed() int {
	res := C.GetGamepadButtonPressed()
	return int(int32(res))
}

//GetGamepadAxisCount : Return gamepad axis count for a gamepad
func GetGamepadAxisCount(gamepad GamepadNumber) int {
	res := C.GetGamepadAxisCount(C.int(int32(gamepad)))
	return int(int32(res))
}

//GetGamepadAxisMovement : Return axis movement value for a gamepad axis
func GetGamepadAxisMovement(gamepad GamepadNumber, axis GamepadAxis) float32 {
	res := C.GetGamepadAxisMovement(C.int(int32(gamepad)), C.int(int32(axis)))
	return float32(res)
}

//IsMouseButtonPressed : Detect if a mouse button has been pressed once
func IsMouseButtonPressed(button MouseButton) bool {
	res := C.IsMouseButtonPressed(C.int(int32(button)))
	return bool(res)
}

//IsMouseButtonDown : Detect if a mouse button is being pressed
func IsMouseButtonDown(button MouseButton) bool {
	res := C.IsMouseButtonDown(C.int(int32(button)))
	return bool(res)
}

//IsMouseButtonReleased : Detect if a mouse button has been released once
func IsMouseButtonReleased(button MouseButton) bool {
	res := C.IsMouseButtonReleased(C.int(int32(button)))
	return bool(res)
}

//IsMouseButtonUp : Detect if a mouse button is NOT being pressed
func IsMouseButtonUp(button MouseButton) bool {
	res := C.IsMouseButtonUp(C.int(int32(button)))
	return bool(res)
}

//GetMouseX : Returns mouse position X
func GetMouseX() int {
	res := C.GetMouseX()
	return int(int32(res))
}

//GetMouseY : Returns mouse position Y
func GetMouseY() int {
	res := C.GetMouseY()
	return int(int32(res))
}

//GetMousePosition : Returns mouse position XY
func GetMousePosition() Vector2 {
	res := C.GetMousePosition()
	return newVector2FromPointer(unsafe.Pointer(&res))
}

//SetMousePosition : Set mouse position XY
func SetMousePosition(x int, y int) {
	C.SetMousePosition(C.int(int32(x)), C.int(int32(y)))
}

//SetMouseOffset : Set mouse offset
func SetMouseOffset(offsetX int, offsetY int) {
	C.SetMouseOffset(C.int(int32(offsetX)), C.int(int32(offsetY)))
}

//SetMouseScale : Set mouse scaling
func SetMouseScale(scaleX float32, scaleY float32) {
	C.SetMouseScale(C.float(scaleX), C.float(scaleY))
}

//GetMouseWheelMove : Returns mouse wheel movement Y
func GetMouseWheelMove() int {
	res := C.GetMouseWheelMove()
	return int(int32(res))
}

//GetTouchX : Returns touch position X for touch point 0 (relative to screen size)
func GetTouchX() int {
	res := C.GetTouchX()
	return int(int32(res))
}

//GetTouchY : Returns touch position Y for touch point 0 (relative to screen size)
func GetTouchY() int {
	res := C.GetTouchY()
	return int(int32(res))
}

//GetTouchPosition : Returns touch position XY for a touch point index (relative to screen size)
func GetTouchPosition(index int) Vector2 {
	res := C.GetTouchPosition(C.int(int32(index)))
	return newVector2FromPointer(unsafe.Pointer(&res))
}
