package raylib

//IsKeyPressed : Detect if a key has been pressed once
func IsKeyPressed(key int) bool {
	res := C.IsKeyPressed(C.int(key))
	return bool(res)
}

//IsKeyDown : Detect if a key is being pressed
func IsKeyDown(key int) bool {
	res := C.IsKeyDown(C.int(key))
	return bool(res)
}

//IsKeyReleased : Detect if a key has been released once
func IsKeyReleased(key int) bool {
	res := C.IsKeyReleased(C.int(key))
	return bool(res)
}

//IsKeyUp : Detect if a key is NOT being pressed
func IsKeyUp(key int) bool {
	res := C.IsKeyUp(C.int(key))
	return bool(res)
}

//GetKeyPressed : Get latest key pressed
func GetKeyPressed() int {
	res := C.GetKeyPressed()
	return int(res)
}

//SetExitKey : Set a custom key to exit program (default is ESC)
func SetExitKey(key int) {
	C.SetExitKey(C.int(key))
}

//IsGamepadAvailable : Detect if a gamepad is available
func IsGamepadAvailable(gamepad int) bool {
	res := C.IsGamepadAvailable(C.int(gamepad))
	return bool(res)
}

//IsGamepadName : Check gamepad name (if available)
func IsGamepadName(gamepad int, name string) (bool, string) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	res := C.IsGamepadName(C.int(gamepad), cname)
	return bool(res), C.GoString(cname)
}

//GetGamepadName : Return gamepad internal name id
func GetGamepadName(gamepad int) string {
	res := C.GetGamepadName(C.int(gamepad))
	return C.GoString(res)
}

//IsGamepadButtonPressed : Detect if a gamepad button has been pressed once
func IsGamepadButtonPressed(gamepad int, button int) bool {
	res := C.IsGamepadButtonPressed(C.int(gamepad), C.int(button))
	return bool(res)
}

//IsGamepadButtonDown : Detect if a gamepad button is being pressed
func IsGamepadButtonDown(gamepad int, button int) bool {
	res := C.IsGamepadButtonDown(C.int(gamepad), C.int(button))
	return bool(res)
}

//IsGamepadButtonReleased : Detect if a gamepad button has been released once
func IsGamepadButtonReleased(gamepad int, button int) bool {
	res := C.IsGamepadButtonReleased(C.int(gamepad), C.int(button))
	return bool(res)
}

//IsGamepadButtonUp : Detect if a gamepad button is NOT being pressed
func IsGamepadButtonUp(gamepad int, button int) bool {
	res := C.IsGamepadButtonUp(C.int(gamepad), C.int(button))
	return bool(res)
}

//GetGamepadButtonPressed : Get the last gamepad button pressed
func GetGamepadButtonPressed() int {
	res := C.GetGamepadButtonPressed()
	return int(res)
}

//GetGamepadAxisCount : Return gamepad axis count for a gamepad
func GetGamepadAxisCount(gamepad int) int {
	res := C.GetGamepadAxisCount(C.int(gamepad))
	return int(res)
}

//GetGamepadAxisMovement : Return axis movement value for a gamepad axis
func GetGamepadAxisMovement(gamepad int, axis int) float32 {
	res := C.GetGamepadAxisMovement(C.int(gamepad), C.int(axis))
	return float32(res)
}

//IsMouseButtonPressed : Detect if a mouse button has been pressed once
func IsMouseButtonPressed(button int) bool {
	res := C.IsMouseButtonPressed(C.int(button))
	return bool(res)
}

//IsMouseButtonDown : Detect if a mouse button is being pressed
func IsMouseButtonDown(button int) bool {
	res := C.IsMouseButtonDown(C.int(button))
	return bool(res)
}

//IsMouseButtonReleased : Detect if a mouse button has been released once
func IsMouseButtonReleased(button int) bool {
	res := C.IsMouseButtonReleased(C.int(button))
	return bool(res)
}

//IsMouseButtonUp : Detect if a mouse button is NOT being pressed
func IsMouseButtonUp(button int) bool {
	res := C.IsMouseButtonUp(C.int(button))
	return bool(res)
}

//GetMouseX : Returns mouse position X
func GetMouseX() int {
	res := C.GetMouseX()
	return int(res)
}

//GetMouseY : Returns mouse position Y
func GetMouseY() int {
	res := C.GetMouseY()
	return int(res)
}

//GetMousePosition : Returns mouse position XY
func GetMousePosition() Vector2 {
	res := C.GetMousePosition()
	return newVector2FromPointer(unsafe.Pointer(&res))
}

//SetMousePosition : Set mouse position XY
func SetMousePosition(x int, y int) {
	C.SetMousePosition(C.int(x), C.int(y))
}

//SetMouseOffset : Set mouse offset
func SetMouseOffset(offsetX int, offsetY int) {
	C.SetMouseOffset(C.int(offsetX), C.int(offsetY))
}

//SetMouseScale : Set mouse scaling
func SetMouseScale(scaleX float32, scaleY float32) {
	C.SetMouseScale(C.float(scaleX), C.float(scaleY))
}

//GetMouseWheelMove : Returns mouse wheel movement Y
func GetMouseWheelMove() int {
	res := C.GetMouseWheelMove()
	return int(res)
}

//GetTouchX : Returns touch position X for touch point 0 (relative to screen size)
func GetTouchX() int {
	res := C.GetTouchX()
	return int(res)
}

//GetTouchY : Returns touch position Y for touch point 0 (relative to screen size)
func GetTouchY() int {
	res := C.GetTouchY()
	return int(res)
}

//GetTouchPosition : Returns touch position XY for a touch point index (relative to screen size)
func GetTouchPosition(index int) Vector2 {
	res := C.GetTouchPosition(C.int(index))
	return newVector2FromPointer(unsafe.Pointer(&res))
}
