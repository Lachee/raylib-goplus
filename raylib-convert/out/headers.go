package raylib

//InitWindow : Initialize window and OpenGL context
func InitWindow(width int, height int, title string) {
	ctitle := C.CString(title)
	defer C.free(unsafe.Pointer(ctitle))
	C.InitWindow(C.int(width), C.int(height), ctitle)

}

//WindowShouldClose : Check if KEY_ESCAPE pressed or Close icon pressed
func WindowShouldClose() bool {
	res := C.WindowShouldClose()
	return bool(&res)
}

//CloseWindow : Close window and unload OpenGL context
func CloseWindow() {
	C.CloseWindow()

}

//IsWindowReady : Check if window has been initialized successfully
func IsWindowReady() bool {
	res := C.IsWindowReady()
	return bool(&res)
}

//IsWindowMinimized : Check if window has been minimized (or lost focus)
func IsWindowMinimized() bool {
	res := C.IsWindowMinimized()
	return bool(&res)
}

//IsWindowResized : Check if window has been resized
func IsWindowResized() bool {
	res := C.IsWindowResized()
	return bool(&res)
}

//IsWindowHidden : Check if window is currently hidden
func IsWindowHidden() bool {
	res := C.IsWindowHidden()
	return bool(&res)
}

//ToggleFullscreen : Toggle fullscreen mode (only PLATFORM_DESKTOP)
func ToggleFullscreen() {
	C.ToggleFullscreen()

}

//UnhideWindow : Show the window
func UnhideWindow() {
	C.UnhideWindow()

}

//HideWindow : Hide the window
func HideWindow() {
	C.HideWindow()

}

//SetWindowIcon : Set icon for window (only PLATFORM_DESKTOP)
func SetWindowIcon(image Image) {
	cimage := *image.cptr()
	C.SetWindowIcon(cimage)

}

//SetWindowTitle : Set title for window (only PLATFORM_DESKTOP)
func SetWindowTitle(title string) {
	ctitle := C.CString(title)
	defer C.free(unsafe.Pointer(ctitle))
	C.SetWindowTitle(ctitle)

}

//SetWindowPosition : Set window position on screen (only PLATFORM_DESKTOP)
func SetWindowPosition(x int, y int) {
	C.SetWindowPosition(C.int(x), C.int(y))

}

//SetWindowMonitor : Set monitor for the current window (fullscreen mode)
func SetWindowMonitor(monitor int) {
	C.SetWindowMonitor(C.int(monitor))

}

//SetWindowMinSize : Set window minimum dimensions (for FLAG_WINDOW_RESIZABLE)
func SetWindowMinSize(width int, height int) {
	C.SetWindowMinSize(C.int(width), C.int(height))

}

//SetWindowSize : Set window dimensions
func SetWindowSize(width int, height int) {
	C.SetWindowSize(C.int(width), C.int(height))

}

//GetWindowHandle : Get native window handle
func GetWindowHandle() {
	C.GetWindowHandle()

}

//GetScreenWidth : Get current screen width
func GetScreenWidth() int {
	res := C.GetScreenWidth()
	return int(&res)
}

//GetScreenHeight : Get current screen height
func GetScreenHeight() int {
	res := C.GetScreenHeight()
	return int(&res)
}

//GetMonitorCount : Get number of connected monitors
func GetMonitorCount() int {
	res := C.GetMonitorCount()
	return int(&res)
}

//GetMonitorWidth : Get primary monitor width
func GetMonitorWidth(monitor int) int {
	res := C.GetMonitorWidth(C.int(monitor))
	return int(&res)
}

//GetMonitorHeight : Get primary monitor height
func GetMonitorHeight(monitor int) int {
	res := C.GetMonitorHeight(C.int(monitor))
	return int(&res)
}

//GetMonitorPhysicalWidth : Get primary monitor physical width in millimetres
func GetMonitorPhysicalWidth(monitor int) int {
	res := C.GetMonitorPhysicalWidth(C.int(monitor))
	return int(&res)
}

//GetMonitorPhysicalHeight : Get primary monitor physical height in millimetres
func GetMonitorPhysicalHeight(monitor int) int {
	res := C.GetMonitorPhysicalHeight(C.int(monitor))
	return int(&res)
}

//GetWindowPosition : Get window position XY on monitor
func GetWindowPosition() Vector2 {
	res := C.GetWindowPosition()
	return newVector2FromPointer(unsafe.Pointer(&res))
}

//GetMonitorName : Get the human-readable, UTF-8 encoded name of the primary monitor
func GetMonitorName(monitor int) string {
	res := C.GetMonitorName(C.int(monitor))
	return C.GoString(&res)
}

//GetClipboardText : Get clipboard text content
func GetClipboardText() string {
	res := C.GetClipboardText()
	return C.GoString(&res)
}

//SetClipboardText : Set clipboard text content
func SetClipboardText(text string) {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	C.SetClipboardText(ctext)

}

//ShowCursor : Shows cursor
func ShowCursor() {
	C.ShowCursor()

}

//HideCursor : Hides cursor
func HideCursor() {
	C.HideCursor()

}

//IsCursorHidden : Check if cursor is not visible
func IsCursorHidden() bool {
	res := C.IsCursorHidden()
	return bool(&res)
}

//EnableCursor : Enables cursor (unlock cursor)
func EnableCursor() {
	C.EnableCursor()

}

//DisableCursor : Disables cursor (lock cursor)
func DisableCursor() {
	C.DisableCursor()

}

//ClearBackground : Set background color (framebuffer clear color)
func ClearBackground(color Color) {
	ccolor := *color.cptr()
	C.ClearBackground(ccolor)

}

//BeginDrawing : Setup canvas (framebuffer) to start drawing
func BeginDrawing() {
	C.BeginDrawing()

}

//EndDrawing : End canvas drawing and swap buffers (double buffering)
func EndDrawing() {
	C.EndDrawing()

}

//BeginMode2D : Initialize 2D mode with custom camera (2D)
func BeginMode2D(camera Camera2D) {
	ccamera := *camera.cptr()
	C.BeginMode2D(ccamera)

}

//EndMode2D : Ends 2D mode with custom camera
func EndMode2D() {
	C.EndMode2D()

}

//BeginMode3D : Initializes 3D mode with custom camera (3D)
func BeginMode3D(camera Camera3D) {
	ccamera := *camera.cptr()
	C.BeginMode3D(ccamera)

}

//EndMode3D : Ends 3D mode and returns to default 2D orthographic mode
func EndMode3D() {
	C.EndMode3D()

}

//BeginTextureMode : Initializes render texture for drawing
func BeginTextureMode(target RenderTexture2D) {
	ctarget := *target.cptr()
	C.BeginTextureMode(ctarget)

}

//EndTextureMode : Ends drawing to render texture
func EndTextureMode() {
	C.EndTextureMode()

}

//BeginScissorMode : Begin scissor mode (define screen area for following drawing)
func BeginScissorMode(x int, y int, width int, height int) {
	C.BeginScissorMode(C.int(x), C.int(y), C.int(width), C.int(height))

}

//EndScissorMode : End scissor mode
func EndScissorMode() {
	C.EndScissorMode()

}

//GetMouseRay : Returns a ray trace from mouse position
func GetMouseRay(mousePosition Vector2, camera Camera) Ray {
	ccamera := *camera.cptr()
	cmousePosition := *mousePosition.cptr()
	res := C.GetMouseRay(cmousePosition, ccamera)
	return newRayFromPointer(unsafe.Pointer(&res))
}

//GetCameraMatrix : Returns camera transform matrix (view matrix)
func GetCameraMatrix(camera Camera) Matrix {
	ccamera := *camera.cptr()
	res := C.GetCameraMatrix(ccamera)
	return newMatrixFromPointer(unsafe.Pointer(&res))
}

//GetCameraMatrix2D : Returns camera 2d transform matrix
func GetCameraMatrix2D(camera Camera2D) Matrix {
	ccamera := *camera.cptr()
	res := C.GetCameraMatrix2D(ccamera)
	return newMatrixFromPointer(unsafe.Pointer(&res))
}

//GetWorldToScreen : Returns the screen space position for a 3d world space position
func GetWorldToScreen(position Vector3, camera Camera) Vector2 {
	ccamera := *camera.cptr()
	cposition := *position.cptr()
	res := C.GetWorldToScreen(cposition, ccamera)
	return newVector2FromPointer(unsafe.Pointer(&res))
}

//GetWorldToScreen2D : Returns the screen space position for a 2d camera world space position
func GetWorldToScreen2D(position Vector2, camera Camera2D) Vector2 {
	ccamera := *camera.cptr()
	cposition := *position.cptr()
	res := C.GetWorldToScreen2D(cposition, ccamera)
	return newVector2FromPointer(unsafe.Pointer(&res))
}

//GetScreenToWorld2D : Returns the world space position for a 2d camera screen space position
func GetScreenToWorld2D(position Vector2, camera Camera2D) Vector2 {
	ccamera := *camera.cptr()
	cposition := *position.cptr()
	res := C.GetScreenToWorld2D(cposition, ccamera)
	return newVector2FromPointer(unsafe.Pointer(&res))
}

//SetTargetFPS : Set target FPS (maximum)
func SetTargetFPS(fps int) {
	C.SetTargetFPS(C.int(fps))

}

//GetFPS : Returns current FPS
func GetFPS() int {
	res := C.GetFPS()
	return int(&res)
}

//GetFrameTime : Returns time in seconds for last frame drawn
func GetFrameTime() float32 {
	res := C.GetFrameTime()
	return float32(&res)
}

//GetTime : Returns elapsed time in seconds since InitWindow()
func GetTime() float64 {
	res := C.GetTime()
	return float64(&res)
}

//ColorToInt : Returns hexadecimal value for a Color
func ColorToInt(color Color) int {
	ccolor := *color.cptr()
	res := C.ColorToInt(ccolor)
	return int(&res)
}

//ColorNormalize : Returns color normalized as float [0..1]
func ColorNormalize(color Color) Vector4 {
	ccolor := *color.cptr()
	res := C.ColorNormalize(ccolor)
	return newVector4FromPointer(unsafe.Pointer(&res))
}

//ColorToHSV : Returns HSV values for a Color
func ColorToHSV(color Color) Vector3 {
	ccolor := *color.cptr()
	res := C.ColorToHSV(ccolor)
	return newVector3FromPointer(unsafe.Pointer(&res))
}

//ColorFromHSV : Returns a Color from HSV values
func ColorFromHSV(hsv Vector3) Color {
	chsv := *hsv.cptr()
	res := C.ColorFromHSV(chsv)
	return newColorFromPointer(unsafe.Pointer(&res))
}

//GetColor : Returns a Color struct from hexadecimal value
func GetColor(hexValue int) Color {
	res := C.GetColor(C.int(hexValue))
	return newColorFromPointer(unsafe.Pointer(&res))
}

//Fade : Color fade-in or fade-out, alpha goes from 0.0f to 1.0f
func Fade(color Color, alpha float32) Color {
	ccolor := *color.cptr()
	res := C.Fade(ccolor, C.float(alpha))
	return newColorFromPointer(unsafe.Pointer(&res))
}

//SetConfigFlags : Setup window configuration flags (view FLAGS)
func SetConfigFlags(flags int) {
	C.SetConfigFlags(C.int(flags))

}

//SetTraceLogLevel : Set the current threshold (minimum) log level
func SetTraceLogLevel(logType int) {
	C.SetTraceLogLevel(C.int(logType))

}

//SetTraceLogExit : Set the exit threshold (minimum) log level
func SetTraceLogExit(logType int) {
	C.SetTraceLogExit(C.int(logType))

}

//SetTraceLogCallback : Set a trace log callback to enable custom logging
func SetTraceLogCallback(callback TraceLogCallback) {
	ccallback := *callback.cptr()
	C.SetTraceLogCallback(ccallback)

}

//TakeScreenshot : Takes a screenshot of current screen (saved a .png)
func TakeScreenshot(fileName string) {
	cfileName := C.CString(fileName)
	defer C.free(unsafe.Pointer(cfileName))
	C.TakeScreenshot(cfileName)

}

//GetRandomValue : Returns a random value between min and max (both included)
func GetRandomValue(min int, max int) int {
	res := C.GetRandomValue(C.int(min), C.int(max))
	return int(&res)
}

//FileExists : Check if file exists
func FileExists(fileName string) bool {
	cfileName := C.CString(fileName)
	defer C.free(unsafe.Pointer(cfileName))
	res := C.FileExists(cfileName)
	return bool(&res)
}

//IsFileExtension : Check file extension
func IsFileExtension(fileName string, ext string) bool {
	cext := C.CString(ext)
	defer C.free(unsafe.Pointer(cext))
	cfileName := C.CString(fileName)
	defer C.free(unsafe.Pointer(cfileName))
	res := C.IsFileExtension(cfileName, cext)
	return bool(&res)
}

//DirectoryExists : Check if a directory path exists
func DirectoryExists(dirPath string) bool {
	cdirPath := C.CString(dirPath)
	defer C.free(unsafe.Pointer(cdirPath))
	res := C.DirectoryExists(cdirPath)
	return bool(&res)
}

//GetExtension : Get pointer to extension for a filename string
func GetExtension(fileName string) string {
	cfileName := C.CString(fileName)
	defer C.free(unsafe.Pointer(cfileName))
	res := C.GetExtension(cfileName)
	return C.GoString(&res)
}

//GetFileName : Get pointer to filename for a path string
func GetFileName(filePath string) string {
	cfilePath := C.CString(filePath)
	defer C.free(unsafe.Pointer(cfilePath))
	res := C.GetFileName(cfilePath)
	return C.GoString(&res)
}

//GetFileNameWithoutExt : Get filename string without extension (uses static string)
func GetFileNameWithoutExt(filePath string) string {
	cfilePath := C.CString(filePath)
	defer C.free(unsafe.Pointer(cfilePath))
	res := C.GetFileNameWithoutExt(cfilePath)
	return C.GoString(&res)
}

//GetDirectoryPath : Get full path for a given fileName with path (uses static string)
func GetDirectoryPath(filePath string) string {
	cfilePath := C.CString(filePath)
	defer C.free(unsafe.Pointer(cfilePath))
	res := C.GetDirectoryPath(cfilePath)
	return C.GoString(&res)
}

//GetPrevDirectoryPath : Get previous directory path for a given path (uses static string)
func GetPrevDirectoryPath(dirPath string) string {
	cdirPath := C.CString(dirPath)
	defer C.free(unsafe.Pointer(cdirPath))
	res := C.GetPrevDirectoryPath(cdirPath)
	return C.GoString(&res)
}

//GetWorkingDirectory : Get current working directory (uses static string)
func GetWorkingDirectory() string {
	res := C.GetWorkingDirectory()
	return C.GoString(&res)
}

//ClearDirectoryFiles : Clear directory files paths buffers (free memory)
func ClearDirectoryFiles() {
	C.ClearDirectoryFiles()

}

//ChangeDirectory : Change working directory, returns true if success
func ChangeDirectory(dir string) bool {
	cdir := C.CString(dir)
	defer C.free(unsafe.Pointer(cdir))
	res := C.ChangeDirectory(cdir)
	return bool(&res)
}

//IsFileDropped : Check if a file has been dropped into window
func IsFileDropped() bool {
	res := C.IsFileDropped()
	return bool(&res)
}

//ClearDroppedFiles : Clear dropped files paths buffer (free memory)
func ClearDroppedFiles() {
	C.ClearDroppedFiles()

}

//GetFileModTime : Get file modification time (last write time)
func GetFileModTime(fileName string) long {
	cfileName := C.CString(fileName)
	defer C.free(unsafe.Pointer(cfileName))
	res := C.GetFileModTime(cfileName)
	return newlongFromPointer(unsafe.Pointer(&res))
}

//CompressData : Compress data (DEFLATE algorythm)
func CompressData(data string, dataLength int, compDataLength int) (string, int) {
	ccompDataLength := C.int(compDataLength)
	cdata := C.CString(data)
	defer C.free(unsafe.Pointer(cdata))
	res := C.CompressData(cdata, C.int(dataLength), &ccompDataLength)
	return C.GoString(&res), int(ccompDataLength)
}

//DecompressData : Decompress data (DEFLATE algorythm)
func DecompressData(compData string, compDataLength int, dataLength int) (string, int) {
	cdataLength := C.int(dataLength)
	ccompData := C.CString(compData)
	defer C.free(unsafe.Pointer(ccompData))
	res := C.DecompressData(ccompData, C.int(compDataLength), &cdataLength)
	return C.GoString(&res), int(cdataLength)
}

//StorageSaveValue : Save integer value to storage file (to defined position)
func StorageSaveValue(position int, value int) {
	C.StorageSaveValue(C.int(position), C.int(value))

}

//StorageLoadValue : Load integer value from storage file (from defined position)
func StorageLoadValue(position int) int {
	res := C.StorageLoadValue(C.int(position))
	return int(&res)
}

//OpenURL : Open URL with default system browser (if available)
func OpenURL(url string) {
	curl := C.CString(url)
	defer C.free(unsafe.Pointer(curl))
	C.OpenURL(curl)

}

//IsKeyPressed : Detect if a key has been pressed once
func IsKeyPressed(key int) bool {
	res := C.IsKeyPressed(C.int(key))
	return bool(&res)
}

//IsKeyDown : Detect if a key is being pressed
func IsKeyDown(key int) bool {
	res := C.IsKeyDown(C.int(key))
	return bool(&res)
}

//IsKeyReleased : Detect if a key has been released once
func IsKeyReleased(key int) bool {
	res := C.IsKeyReleased(C.int(key))
	return bool(&res)
}

//IsKeyUp : Detect if a key is NOT being pressed
func IsKeyUp(key int) bool {
	res := C.IsKeyUp(C.int(key))
	return bool(&res)
}

//GetKeyPressed : Get latest key pressed
func GetKeyPressed() int {
	res := C.GetKeyPressed()
	return int(&res)
}

//SetExitKey : Set a custom key to exit program (default is ESC)
func SetExitKey(key int) {
	C.SetExitKey(C.int(key))

}

//IsGamepadAvailable : Detect if a gamepad is available
func IsGamepadAvailable(gamepad int) bool {
	res := C.IsGamepadAvailable(C.int(gamepad))
	return bool(&res)
}

//IsGamepadName : Check gamepad name (if available)
func IsGamepadName(gamepad int, name string) bool {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	res := C.IsGamepadName(C.int(gamepad), cname)
	return bool(&res)
}

//GetGamepadName : Return gamepad internal name id
func GetGamepadName(gamepad int) string {
	res := C.GetGamepadName(C.int(gamepad))
	return C.GoString(&res)
}

//IsGamepadButtonPressed : Detect if a gamepad button has been pressed once
func IsGamepadButtonPressed(gamepad int, button int) bool {
	res := C.IsGamepadButtonPressed(C.int(gamepad), C.int(button))
	return bool(&res)
}

//IsGamepadButtonDown : Detect if a gamepad button is being pressed
func IsGamepadButtonDown(gamepad int, button int) bool {
	res := C.IsGamepadButtonDown(C.int(gamepad), C.int(button))
	return bool(&res)
}

//IsGamepadButtonReleased : Detect if a gamepad button has been released once
func IsGamepadButtonReleased(gamepad int, button int) bool {
	res := C.IsGamepadButtonReleased(C.int(gamepad), C.int(button))
	return bool(&res)
}

//IsGamepadButtonUp : Detect if a gamepad button is NOT being pressed
func IsGamepadButtonUp(gamepad int, button int) bool {
	res := C.IsGamepadButtonUp(C.int(gamepad), C.int(button))
	return bool(&res)
}

//GetGamepadButtonPressed : Get the last gamepad button pressed
func GetGamepadButtonPressed() int {
	res := C.GetGamepadButtonPressed()
	return int(&res)
}

//GetGamepadAxisCount : Return gamepad axis count for a gamepad
func GetGamepadAxisCount(gamepad int) int {
	res := C.GetGamepadAxisCount(C.int(gamepad))
	return int(&res)
}

//GetGamepadAxisMovement : Return axis movement value for a gamepad axis
func GetGamepadAxisMovement(gamepad int, axis int) float32 {
	res := C.GetGamepadAxisMovement(C.int(gamepad), C.int(axis))
	return float32(&res)
}

//IsMouseButtonPressed : Detect if a mouse button has been pressed once
func IsMouseButtonPressed(button int) bool {
	res := C.IsMouseButtonPressed(C.int(button))
	return bool(&res)
}

//IsMouseButtonDown : Detect if a mouse button is being pressed
func IsMouseButtonDown(button int) bool {
	res := C.IsMouseButtonDown(C.int(button))
	return bool(&res)
}

//IsMouseButtonReleased : Detect if a mouse button has been released once
func IsMouseButtonReleased(button int) bool {
	res := C.IsMouseButtonReleased(C.int(button))
	return bool(&res)
}

//IsMouseButtonUp : Detect if a mouse button is NOT being pressed
func IsMouseButtonUp(button int) bool {
	res := C.IsMouseButtonUp(C.int(button))
	return bool(&res)
}

//GetMouseX : Returns mouse position X
func GetMouseX() int {
	res := C.GetMouseX()
	return int(&res)
}

//GetMouseY : Returns mouse position Y
func GetMouseY() int {
	res := C.GetMouseY()
	return int(&res)
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
	return int(&res)
}

//GetTouchX : Returns touch position X for touch point 0 (relative to screen size)
func GetTouchX() int {
	res := C.GetTouchX()
	return int(&res)
}

//GetTouchY : Returns touch position Y for touch point 0 (relative to screen size)
func GetTouchY() int {
	res := C.GetTouchY()
	return int(&res)
}

//GetTouchPosition : Returns touch position XY for a touch point index (relative to screen size)
func GetTouchPosition(index int) Vector2 {
	res := C.GetTouchPosition(C.int(index))
	return newVector2FromPointer(unsafe.Pointer(&res))
}

//SetGesturesEnabled : Enable a set of gestures using flags
func SetGesturesEnabled(gestureFlags int) {
	C.SetGesturesEnabled(C.int(gestureFlags))

}

//IsGestureDetected : Check if a gesture have been detected
func IsGestureDetected(gesture int) bool {
	res := C.IsGestureDetected(C.int(gesture))
	return bool(&res)
}

//GetGestureDetected : Get latest detected gesture
func GetGestureDetected() int {
	res := C.GetGestureDetected()
	return int(&res)
}

//GetTouchPointsCount : Get touch points count
func GetTouchPointsCount() int {
	res := C.GetTouchPointsCount()
	return int(&res)
}

//GetGestureHoldDuration : Get gesture hold time in milliseconds
func GetGestureHoldDuration() float32 {
	res := C.GetGestureHoldDuration()
	return float32(&res)
}

//GetGestureDragVector : Get gesture drag vector
func GetGestureDragVector() Vector2 {
	res := C.GetGestureDragVector()
	return newVector2FromPointer(unsafe.Pointer(&res))
}

//GetGestureDragAngle : Get gesture drag angle
func GetGestureDragAngle() float32 {
	res := C.GetGestureDragAngle()
	return float32(&res)
}

//GetGesturePinchVector : Get gesture pinch delta
func GetGesturePinchVector() Vector2 {
	res := C.GetGesturePinchVector()
	return newVector2FromPointer(unsafe.Pointer(&res))
}

//GetGesturePinchAngle : Get gesture pinch angle
func GetGesturePinchAngle() float32 {
	res := C.GetGesturePinchAngle()
	return float32(&res)
}

//SetCameraMode : Set camera mode (multiple camera modes available)
func SetCameraMode(camera Camera, mode int) {
	ccamera := *camera.cptr()
	C.SetCameraMode(ccamera, C.int(mode))

}

//UpdateCamera : Update camera position for selected mode
func UpdateCamera(camera Camera) Camera {
	ccamera := camera.cptr()
	C.UpdateCamera(&ccamera)
	return newCameraFromPointer(unsafe.Pointer(ccamera))
}

//SetCameraPanControl : Set camera pan key to combine with mouse movement (free camera)
func SetCameraPanControl(panKey int) {
	C.SetCameraPanControl(C.int(panKey))

}

//SetCameraAltControl : Set camera alt key to combine with mouse movement (free camera)
func SetCameraAltControl(altKey int) {
	C.SetCameraAltControl(C.int(altKey))

}

//SetCameraSmoothZoomControl : Set camera smooth zoom key to combine with mouse (free camera)
func SetCameraSmoothZoomControl(szKey int) {
	C.SetCameraSmoothZoomControl(C.int(szKey))

}

//SetCameraMoveControls : Set camera move controls (1st person and 3rd person cameras)
func SetCameraMoveControls(frontKey int, backKey int, rightKey int, leftKey int, upKey int, downKey int) {
	C.SetCameraMoveControls(C.int(frontKey), C.int(backKey), C.int(rightKey), C.int(leftKey), C.int(upKey), C.int(downKey))

}

//DrawPixel : Draw a pixel
func DrawPixel(posX int, posY int, color Color) {
	ccolor := *color.cptr()
	C.DrawPixel(C.int(posX), C.int(posY), ccolor)

}

//DrawPixelV : Draw a pixel (Vector version)
func DrawPixelV(position Vector2, color Color) {
	ccolor := *color.cptr()
	cposition := *position.cptr()
	C.DrawPixelV(cposition, ccolor)

}

//DrawLine : Draw a line
func DrawLine(startPosX int, startPosY int, endPosX int, endPosY int, color Color) {
	ccolor := *color.cptr()
	C.DrawLine(C.int(startPosX), C.int(startPosY), C.int(endPosX), C.int(endPosY), ccolor)

}

//DrawLineV : Draw a line (Vector version)
func DrawLineV(startPos Vector2, endPos Vector2, color Color) {
	ccolor := *color.cptr()
	cendPos := *endPos.cptr()
	cstartPos := *startPos.cptr()
	C.DrawLineV(cstartPos, cendPos, ccolor)

}

//DrawLineEx : Draw a line defining thickness
func DrawLineEx(startPos Vector2, endPos Vector2, thick float32, color Color) {
	ccolor := *color.cptr()
	cendPos := *endPos.cptr()
	cstartPos := *startPos.cptr()
	C.DrawLineEx(cstartPos, cendPos, C.float(thick), ccolor)

}

//DrawLineBezier : Draw a line using cubic-bezier curves in-out
func DrawLineBezier(startPos Vector2, endPos Vector2, thick float32, color Color) {
	ccolor := *color.cptr()
	cendPos := *endPos.cptr()
	cstartPos := *startPos.cptr()
	C.DrawLineBezier(cstartPos, cendPos, C.float(thick), ccolor)

}

//DrawLineStrip : Draw lines sequence
func DrawLineStrip(points Vector2, numPoints int, color Color) Vector2 {
	ccolor := *color.cptr()
	cpoints := points.cptr()
	C.DrawLineStrip(&cpoints, C.int(numPoints), ccolor)
	return newVector2FromPointer(unsafe.Pointer(cpoints))
}

//DrawCircle : Draw a color-filled circle
func DrawCircle(centerX int, centerY int, radius float32, color Color) {
	ccolor := *color.cptr()
	C.DrawCircle(C.int(centerX), C.int(centerY), C.float(radius), ccolor)

}

//DrawCircleSector : Draw a piece of a circle
func DrawCircleSector(center Vector2, radius float32, startAngle int, endAngle int, segments int, color Color) {
	ccolor := *color.cptr()
	ccenter := *center.cptr()
	C.DrawCircleSector(ccenter, C.float(radius), C.int(startAngle), C.int(endAngle), C.int(segments), ccolor)

}

//DrawCircleSectorLines : Draw circle sector outline
func DrawCircleSectorLines(center Vector2, radius float32, startAngle int, endAngle int, segments int, color Color) {
	ccolor := *color.cptr()
	ccenter := *center.cptr()
	C.DrawCircleSectorLines(ccenter, C.float(radius), C.int(startAngle), C.int(endAngle), C.int(segments), ccolor)

}

//DrawCircleGradient : Draw a gradient-filled circle
func DrawCircleGradient(centerX int, centerY int, radius float32, color1 Color, color2 Color) {
	ccolor2 := *color2.cptr()
	ccolor1 := *color1.cptr()
	C.DrawCircleGradient(C.int(centerX), C.int(centerY), C.float(radius), ccolor1, ccolor2)

}

//DrawCircleV : Draw a color-filled circle (Vector version)
func DrawCircleV(center Vector2, radius float32, color Color) {
	ccolor := *color.cptr()
	ccenter := *center.cptr()
	C.DrawCircleV(ccenter, C.float(radius), ccolor)

}

//DrawCircleLines : Draw circle outline
func DrawCircleLines(centerX int, centerY int, radius float32, color Color) {
	ccolor := *color.cptr()
	C.DrawCircleLines(C.int(centerX), C.int(centerY), C.float(radius), ccolor)

}

//DrawRing : Draw ring
func DrawRing(center Vector2, innerRadius float32, outerRadius float32, startAngle int, endAngle int, segments int, color Color) {
	ccolor := *color.cptr()
	ccenter := *center.cptr()
	C.DrawRing(ccenter, C.float(innerRadius), C.float(outerRadius), C.int(startAngle), C.int(endAngle), C.int(segments), ccolor)

}

//DrawRingLines : Draw ring outline
func DrawRingLines(center Vector2, innerRadius float32, outerRadius float32, startAngle int, endAngle int, segments int, color Color) {
	ccolor := *color.cptr()
	ccenter := *center.cptr()
	C.DrawRingLines(ccenter, C.float(innerRadius), C.float(outerRadius), C.int(startAngle), C.int(endAngle), C.int(segments), ccolor)

}

//DrawRectangle : Draw a color-filled rectangle
func DrawRectangle(posX int, posY int, width int, height int, color Color) {
	ccolor := *color.cptr()
	C.DrawRectangle(C.int(posX), C.int(posY), C.int(width), C.int(height), ccolor)

}

//DrawRectangleV : Draw a color-filled rectangle (Vector version)
func DrawRectangleV(position Vector2, size Vector2, color Color) {
	ccolor := *color.cptr()
	csize := *size.cptr()
	cposition := *position.cptr()
	C.DrawRectangleV(cposition, csize, ccolor)

}

//DrawRectangleRec : Draw a color-filled rectangle
func DrawRectangleRec(rec Rectangle, color Color) {
	ccolor := *color.cptr()
	crec := *rec.cptr()
	C.DrawRectangleRec(crec, ccolor)

}

//DrawRectanglePro : Draw a color-filled rectangle with pro parameters
func DrawRectanglePro(rec Rectangle, origin Vector2, rotation float32, color Color) {
	ccolor := *color.cptr()
	corigin := *origin.cptr()
	crec := *rec.cptr()
	C.DrawRectanglePro(crec, corigin, C.float(rotation), ccolor)

}

//DrawRectangleGradientV : Draw a vertical-gradient-filled rectangle
func DrawRectangleGradientV(posX int, posY int, width int, height int, color1 Color, color2 Color) {
	ccolor2 := *color2.cptr()
	ccolor1 := *color1.cptr()
	C.DrawRectangleGradientV(C.int(posX), C.int(posY), C.int(width), C.int(height), ccolor1, ccolor2)

}

//DrawRectangleGradientH : Draw a horizontal-gradient-filled rectangle
func DrawRectangleGradientH(posX int, posY int, width int, height int, color1 Color, color2 Color) {
	ccolor2 := *color2.cptr()
	ccolor1 := *color1.cptr()
	C.DrawRectangleGradientH(C.int(posX), C.int(posY), C.int(width), C.int(height), ccolor1, ccolor2)

}

//DrawRectangleGradientEx : Draw a gradient-filled rectangle with custom vertex colors
func DrawRectangleGradientEx(rec Rectangle, col1 Color, col2 Color, col3 Color, col4 Color) {
	ccol4 := *col4.cptr()
	ccol3 := *col3.cptr()
	ccol2 := *col2.cptr()
	ccol1 := *col1.cptr()
	crec := *rec.cptr()
	C.DrawRectangleGradientEx(crec, ccol1, ccol2, ccol3, ccol4)

}

//DrawRectangleLines : Draw rectangle outline
func DrawRectangleLines(posX int, posY int, width int, height int, color Color) {
	ccolor := *color.cptr()
	C.DrawRectangleLines(C.int(posX), C.int(posY), C.int(width), C.int(height), ccolor)

}

//DrawRectangleLinesEx : Draw rectangle outline with extended parameters
func DrawRectangleLinesEx(rec Rectangle, lineThick int, color Color) {
	ccolor := *color.cptr()
	crec := *rec.cptr()
	C.DrawRectangleLinesEx(crec, C.int(lineThick), ccolor)

}

//DrawRectangleRounded : Draw rectangle with rounded edges
func DrawRectangleRounded(rec Rectangle, roundness float32, segments int, color Color) {
	ccolor := *color.cptr()
	crec := *rec.cptr()
	C.DrawRectangleRounded(crec, C.float(roundness), C.int(segments), ccolor)

}

//DrawRectangleRoundedLines : Draw rectangle with rounded edges outline
func DrawRectangleRoundedLines(rec Rectangle, roundness float32, segments int, lineThick int, color Color) {
	ccolor := *color.cptr()
	crec := *rec.cptr()
	C.DrawRectangleRoundedLines(crec, C.float(roundness), C.int(segments), C.int(lineThick), ccolor)

}

//DrawTriangle : Draw a color-filled triangle (vertex in counter-clockwise order!)
func DrawTriangle(v1 Vector2, v2 Vector2, v3 Vector2, color Color) {
	ccolor := *color.cptr()
	cv3 := *v3.cptr()
	cv2 := *v2.cptr()
	cv1 := *v1.cptr()
	C.DrawTriangle(cv1, cv2, cv3, ccolor)

}

//DrawTriangleLines : Draw triangle outline (vertex in counter-clockwise order!)
func DrawTriangleLines(v1 Vector2, v2 Vector2, v3 Vector2, color Color) {
	ccolor := *color.cptr()
	cv3 := *v3.cptr()
	cv2 := *v2.cptr()
	cv1 := *v1.cptr()
	C.DrawTriangleLines(cv1, cv2, cv3, ccolor)

}

//DrawTriangleFan : Draw a triangle fan defined by points (first vertex is the center)
func DrawTriangleFan(points Vector2, numPoints int, color Color) Vector2 {
	ccolor := *color.cptr()
	cpoints := points.cptr()
	C.DrawTriangleFan(&cpoints, C.int(numPoints), ccolor)
	return newVector2FromPointer(unsafe.Pointer(cpoints))
}

//DrawTriangleStrip : Draw a triangle strip defined by points
func DrawTriangleStrip(points Vector2, pointsCount int, color Color) Vector2 {
	ccolor := *color.cptr()
	cpoints := points.cptr()
	C.DrawTriangleStrip(&cpoints, C.int(pointsCount), ccolor)
	return newVector2FromPointer(unsafe.Pointer(cpoints))
}

//DrawPoly : Draw a regular polygon (Vector version)
func DrawPoly(center Vector2, sides int, radius float32, rotation float32, color Color) {
	ccolor := *color.cptr()
	ccenter := *center.cptr()
	C.DrawPoly(ccenter, C.int(sides), C.float(radius), C.float(rotation), ccolor)

}

//SetShapesTexture : Define default texture used to draw shapes
func SetShapesTexture(texture Texture2D, source Rectangle) {
	csource := *source.cptr()
	ctexture := *texture.cptr()
	C.SetShapesTexture(ctexture, csource)

}

//CheckCollisionRecs : Check collision between two rectangles
func CheckCollisionRecs(rec1 Rectangle, rec2 Rectangle) bool {
	crec2 := *rec2.cptr()
	crec1 := *rec1.cptr()
	res := C.CheckCollisionRecs(crec1, crec2)
	return bool(&res)
}

//CheckCollisionCircles : Check collision between two circles
func CheckCollisionCircles(center1 Vector2, radius1 float32, center2 Vector2, radius2 float32) bool {
	ccenter2 := *center2.cptr()
	ccenter1 := *center1.cptr()
	res := C.CheckCollisionCircles(ccenter1, C.float(radius1), ccenter2, C.float(radius2))
	return bool(&res)
}

//CheckCollisionCircleRec : Check collision between circle and rectangle
func CheckCollisionCircleRec(center Vector2, radius float32, rec Rectangle) bool {
	crec := *rec.cptr()
	ccenter := *center.cptr()
	res := C.CheckCollisionCircleRec(ccenter, C.float(radius), crec)
	return bool(&res)
}

//GetCollisionRec : Get collision rectangle for two rectangles collision
func GetCollisionRec(rec1 Rectangle, rec2 Rectangle) Rectangle {
	crec2 := *rec2.cptr()
	crec1 := *rec1.cptr()
	res := C.GetCollisionRec(crec1, crec2)
	return newRectangleFromPointer(unsafe.Pointer(&res))
}

//CheckCollisionPointRec : Check if point is inside rectangle
func CheckCollisionPointRec(point Vector2, rec Rectangle) bool {
	crec := *rec.cptr()
	cpoint := *point.cptr()
	res := C.CheckCollisionPointRec(cpoint, crec)
	return bool(&res)
}

//CheckCollisionPointCircle : Check if point is inside circle
func CheckCollisionPointCircle(point Vector2, center Vector2, radius float32) bool {
	ccenter := *center.cptr()
	cpoint := *point.cptr()
	res := C.CheckCollisionPointCircle(cpoint, ccenter, C.float(radius))
	return bool(&res)
}

//CheckCollisionPointTriangle : Check if point is inside a triangle
func CheckCollisionPointTriangle(point Vector2, p1 Vector2, p2 Vector2, p3 Vector2) bool {
	cp3 := *p3.cptr()
	cp2 := *p2.cptr()
	cp1 := *p1.cptr()
	cpoint := *point.cptr()
	res := C.CheckCollisionPointTriangle(cpoint, cp1, cp2, cp3)
	return bool(&res)
}

//LoadImage : Load image from file into CPU memory (RAM)
func LoadImage(fileName string) Image {
	cfileName := C.CString(fileName)
	defer C.free(unsafe.Pointer(cfileName))
	res := C.LoadImage(cfileName)
	return newImageFromPointer(unsafe.Pointer(&res))
}

// LoadImageEx Load image data from Color array data (RGBA - 32bit)
func LoadImageEx(pixels []Color, width, height int32) *Image {
	cpixels := pixels[0].cptr()
	cwidth := (C.int)(width)
	cheight := (C.int)(height)
	ret := C.LoadImageEx(cpixels, cwidth, cheight) //We call the C function as raylib uses stb_image to load images.
	v := newImageFromPointer(unsafe.Pointer(&ret))
	return v
}

//LoadImagePro loads raw data wtih parameters
func LoadImagePro(pixels []byte, width, height int32, format PixelFormat) *Image {
	data := unsafe.Pointer(&pixels[0])
	res := C.LoadImagePro(data, C.int(width), C.int(height), C.int(format))
	return newImageFromPointer(unsafe.Pointer(&res))
}

//LoadImageRaw : Load image from RAW file data
func LoadImageRaw(fileName string, width int, height int, format int, headerSize int) Image {
	cfileName := C.CString(fileName)
	defer C.free(unsafe.Pointer(cfileName))
	res := C.LoadImageRaw(cfileName, C.int(width), C.int(height), C.int(format), C.int(headerSize))
	return newImageFromPointer(unsafe.Pointer(&res))
}

//ExportImage : Export image data to file
func ExportImage(image Image, fileName string) {
	cfileName := C.CString(fileName)
	defer C.free(unsafe.Pointer(cfileName))
	cimage := *image.cptr()
	C.ExportImage(cimage, cfileName)

}

//ExportImageAsCode : Export image as code file defining an array of bytes
func ExportImageAsCode(image Image, fileName string) {
	cfileName := C.CString(fileName)
	defer C.free(unsafe.Pointer(cfileName))
	cimage := *image.cptr()
	C.ExportImageAsCode(cimage, cfileName)

}

//LoadTexture : Load texture from file into GPU memory (VRAM)
func LoadTexture(fileName string) Texture2D {
	cfileName := C.CString(fileName)
	defer C.free(unsafe.Pointer(cfileName))
	res := C.LoadTexture(cfileName)
	return newTexture2DFromPointer(unsafe.Pointer(&res))
}

//LoadTextureFromImage : Load texture from image data
func LoadTextureFromImage(image Image) Texture2D {
	cimage := *image.cptr()
	res := C.LoadTextureFromImage(cimage)
	return newTexture2DFromPointer(unsafe.Pointer(&res))
}

//LoadTextureCubemap : Load cubemap from image, multiple image cubemap layouts supported
func LoadTextureCubemap(image Image, layoutType int) TextureCubemap {
	cimage := *image.cptr()
	res := C.LoadTextureCubemap(cimage, C.int(layoutType))
	return newTextureCubemapFromPointer(unsafe.Pointer(&res))
}

//LoadRenderTexture : Load texture for rendering (framebuffer)
func LoadRenderTexture(width int, height int) RenderTexture2D {
	res := C.LoadRenderTexture(C.int(width), C.int(height))
	return newRenderTexture2DFromPointer(unsafe.Pointer(&res))
}

//UnloadImage : Unload image from CPU memory (RAM)
func UnloadImage(image Image) {
	cimage := *image.cptr()
	C.UnloadImage(cimage)

}

//UnloadTexture : Unload texture from GPU memory (VRAM)
func UnloadTexture(texture Texture2D) {
	ctexture := *texture.cptr()
	C.UnloadTexture(ctexture)

}

//UnloadRenderTexture : Unload render texture from GPU memory (VRAM)
func UnloadRenderTexture(target RenderTexture2D) {
	ctarget := *target.cptr()
	C.UnloadRenderTexture(ctarget)

}

//GetImageAlphaBorder : Get image alpha border rectangle
func GetImageAlphaBorder(image Image, threshold float32) Rectangle {
	cimage := *image.cptr()
	res := C.GetImageAlphaBorder(cimage, C.float(threshold))
	return newRectangleFromPointer(unsafe.Pointer(&res))
}

//GetPixelDataSize : Get pixel data size in bytes (image or texture)
func GetPixelDataSize(width int, height int, format int) int {
	res := C.GetPixelDataSize(C.int(width), C.int(height), C.int(format))
	return int(&res)
}

//GetTextureData : Get pixel data from GPU texture and return an Image
func GetTextureData(texture Texture2D) Image {
	ctexture := *texture.cptr()
	res := C.GetTextureData(ctexture)
	return newImageFromPointer(unsafe.Pointer(&res))
}

//GetScreenData : Get pixel data from screen buffer and return an Image (screenshot)
func GetScreenData() Image {
	res := C.GetScreenData()
	return newImageFromPointer(unsafe.Pointer(&res))
}

//UpdateTexture : Update GPU texture with new data
func UpdateTexture(texture Texture2D, pixels unsafe.Pointer) {
	ctexture := *texture.cptr()
	C.UpdateTexture(ctexture, pixels)

}

//ImageCopy : Create an image duplicate (useful for transformations)
func ImageCopy(image Image) Image {
	cimage := *image.cptr()
	res := C.ImageCopy(cimage)
	return newImageFromPointer(unsafe.Pointer(&res))
}

//ImageFromImage : Create an image from another image piece
func ImageFromImage(image Image, rec Rectangle) Image {
	crec := *rec.cptr()
	cimage := *image.cptr()
	res := C.ImageFromImage(cimage, crec)
	return newImageFromPointer(unsafe.Pointer(&res))
}

//ImageToPOT : Convert image to POT (power-of-two)
func ImageToPOT(image Image, fillColor Color) Image {
	cfillColor := *fillColor.cptr()
	cimage := image.cptr()
	C.ImageToPOT(&cimage, cfillColor)
	return newImageFromPointer(unsafe.Pointer(cimage))
}

//ImageFormat : Convert image data to desired format
func ImageFormat(image Image, newFormat int) Image {
	cimage := image.cptr()
	C.ImageFormat(&cimage, C.int(newFormat))
	return newImageFromPointer(unsafe.Pointer(cimage))
}

//ImageAlphaMask : Apply alpha mask to image
func ImageAlphaMask(image Image, alphaMask Image) Image {
	calphaMask := *alphaMask.cptr()
	cimage := image.cptr()
	C.ImageAlphaMask(&cimage, calphaMask)
	return newImageFromPointer(unsafe.Pointer(cimage))
}

//ImageAlphaClear : Clear alpha channel to desired color
func ImageAlphaClear(image Image, color Color, threshold float32) Image {
	ccolor := *color.cptr()
	cimage := image.cptr()
	C.ImageAlphaClear(&cimage, ccolor, C.float(threshold))
	return newImageFromPointer(unsafe.Pointer(cimage))
}

//ImageAlphaCrop : Crop image depending on alpha value
func ImageAlphaCrop(image Image, threshold float32) Image {
	cimage := image.cptr()
	C.ImageAlphaCrop(&cimage, C.float(threshold))
	return newImageFromPointer(unsafe.Pointer(cimage))
}

//ImageAlphaPremultiply : Premultiply alpha channel
func ImageAlphaPremultiply(image Image) Image {
	cimage := image.cptr()
	C.ImageAlphaPremultiply(&cimage)
	return newImageFromPointer(unsafe.Pointer(cimage))
}

//ImageCrop : Crop an image to a defined rectangle
func ImageCrop(image Image, crop Rectangle) Image {
	ccrop := *crop.cptr()
	cimage := image.cptr()
	C.ImageCrop(&cimage, ccrop)
	return newImageFromPointer(unsafe.Pointer(cimage))
}

//ImageResize : Resize image (Bicubic scaling algorithm)
func ImageResize(image Image, newWidth int, newHeight int) Image {
	cimage := image.cptr()
	C.ImageResize(&cimage, C.int(newWidth), C.int(newHeight))
	return newImageFromPointer(unsafe.Pointer(cimage))
}

//ImageResizeNN : Resize image (Nearest-Neighbor scaling algorithm)
func ImageResizeNN(image Image, newWidth int, newHeight int) Image {
	cimage := image.cptr()
	C.ImageResizeNN(&cimage, C.int(newWidth), C.int(newHeight))
	return newImageFromPointer(unsafe.Pointer(cimage))
}

//ImageResizeCanvas : Resize canvas and fill with color
func ImageResizeCanvas(image Image, newWidth int, newHeight int, offsetX int, offsetY int, color Color) Image {
	ccolor := *color.cptr()
	cimage := image.cptr()
	C.ImageResizeCanvas(&cimage, C.int(newWidth), C.int(newHeight), C.int(offsetX), C.int(offsetY), ccolor)
	return newImageFromPointer(unsafe.Pointer(cimage))
}

//ImageMipmaps : Generate all mipmap levels for a provided image
func ImageMipmaps(image Image) Image {
	cimage := image.cptr()
	C.ImageMipmaps(&cimage)
	return newImageFromPointer(unsafe.Pointer(cimage))
}

//ImageDither : Dither image data to 16bpp or lower (Floyd-Steinberg dithering)
func ImageDither(image Image, rBpp int, gBpp int, bBpp int, aBpp int) Image {
	cimage := image.cptr()
	C.ImageDither(&cimage, C.int(rBpp), C.int(gBpp), C.int(bBpp), C.int(aBpp))
	return newImageFromPointer(unsafe.Pointer(cimage))
}

//ImageText : Create an image from text (default font)
func ImageText(text string, fontSize int, color Color) Image {
	ccolor := *color.cptr()
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	res := C.ImageText(ctext, C.int(fontSize), ccolor)
	return newImageFromPointer(unsafe.Pointer(&res))
}

//ImageTextEx : Create an image from text (custom sprite font)
func ImageTextEx(font Font, text string, fontSize float32, spacing float32, tint Color) Image {
	ctint := *tint.cptr()
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	cfont := *font.cptr()
	res := C.ImageTextEx(cfont, ctext, C.float(fontSize), C.float(spacing), ctint)
	return newImageFromPointer(unsafe.Pointer(&res))
}

//ImageDraw : Draw a source image within a destination image (tint applied to source)
func ImageDraw(dst Image, src Image, srcRec Rectangle, dstRec Rectangle, tint Color) Image {
	ctint := *tint.cptr()
	cdstRec := *dstRec.cptr()
	csrcRec := *srcRec.cptr()
	csrc := *src.cptr()
	cdst := dst.cptr()
	C.ImageDraw(&cdst, csrc, csrcRec, cdstRec, ctint)
	return newImageFromPointer(unsafe.Pointer(cdst))
}

//ImageDrawRectangle : Draw rectangle within an image
func ImageDrawRectangle(dst Image, rec Rectangle, color Color) Image {
	ccolor := *color.cptr()
	crec := *rec.cptr()
	cdst := dst.cptr()
	C.ImageDrawRectangle(&cdst, crec, ccolor)
	return newImageFromPointer(unsafe.Pointer(cdst))
}

//ImageDrawRectangleLines : Draw rectangle lines within an image
func ImageDrawRectangleLines(dst Image, rec Rectangle, thick int, color Color) Image {
	ccolor := *color.cptr()
	crec := *rec.cptr()
	cdst := dst.cptr()
	C.ImageDrawRectangleLines(&cdst, crec, C.int(thick), ccolor)
	return newImageFromPointer(unsafe.Pointer(cdst))
}

//ImageDrawText : Draw text (default font) within an image (destination)
func ImageDrawText(dst Image, position Vector2, text string, fontSize int, color Color) Image {
	ccolor := *color.cptr()
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	cposition := *position.cptr()
	cdst := dst.cptr()
	C.ImageDrawText(&cdst, cposition, ctext, C.int(fontSize), ccolor)
	return newImageFromPointer(unsafe.Pointer(cdst))
}

//ImageDrawTextEx : Draw text (custom sprite font) within an image (destination)
func ImageDrawTextEx(dst Image, position Vector2, font Font, text string, fontSize float32, spacing float32, color Color) Image {
	ccolor := *color.cptr()
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	cfont := *font.cptr()
	cposition := *position.cptr()
	cdst := dst.cptr()
	C.ImageDrawTextEx(&cdst, cposition, cfont, ctext, C.float(fontSize), C.float(spacing), ccolor)
	return newImageFromPointer(unsafe.Pointer(cdst))
}

//ImageFlipVertical : Flip image vertically
func ImageFlipVertical(image Image) Image {
	cimage := image.cptr()
	C.ImageFlipVertical(&cimage)
	return newImageFromPointer(unsafe.Pointer(cimage))
}

//ImageFlipHorizontal : Flip image horizontally
func ImageFlipHorizontal(image Image) Image {
	cimage := image.cptr()
	C.ImageFlipHorizontal(&cimage)
	return newImageFromPointer(unsafe.Pointer(cimage))
}

//ImageRotateCW : Rotate image clockwise 90deg
func ImageRotateCW(image Image) Image {
	cimage := image.cptr()
	C.ImageRotateCW(&cimage)
	return newImageFromPointer(unsafe.Pointer(cimage))
}

//ImageRotateCCW : Rotate image counter-clockwise 90deg
func ImageRotateCCW(image Image) Image {
	cimage := image.cptr()
	C.ImageRotateCCW(&cimage)
	return newImageFromPointer(unsafe.Pointer(cimage))
}

//ImageColorTint : Modify image color: tint
func ImageColorTint(image Image, color Color) Image {
	ccolor := *color.cptr()
	cimage := image.cptr()
	C.ImageColorTint(&cimage, ccolor)
	return newImageFromPointer(unsafe.Pointer(cimage))
}

//ImageColorInvert : Modify image color: invert
func ImageColorInvert(image Image) Image {
	cimage := image.cptr()
	C.ImageColorInvert(&cimage)
	return newImageFromPointer(unsafe.Pointer(cimage))
}

//ImageColorGrayscale : Modify image color: grayscale
func ImageColorGrayscale(image Image) Image {
	cimage := image.cptr()
	C.ImageColorGrayscale(&cimage)
	return newImageFromPointer(unsafe.Pointer(cimage))
}

//ImageColorContrast : Modify image color: contrast (-100 to 100)
func ImageColorContrast(image Image, contrast float32) Image {
	cimage := image.cptr()
	C.ImageColorContrast(&cimage, C.float(contrast))
	return newImageFromPointer(unsafe.Pointer(cimage))
}

//ImageColorBrightness : Modify image color: brightness (-255 to 255)
func ImageColorBrightness(image Image, brightness int) Image {
	cimage := image.cptr()
	C.ImageColorBrightness(&cimage, C.int(brightness))
	return newImageFromPointer(unsafe.Pointer(cimage))
}

//ImageColorReplace : Modify image color: replace color
func ImageColorReplace(image Image, color Color, replace Color) Image {
	creplace := *replace.cptr()
	ccolor := *color.cptr()
	cimage := image.cptr()
	C.ImageColorReplace(&cimage, ccolor, creplace)
	return newImageFromPointer(unsafe.Pointer(cimage))
}

//GenImageColor : Generate image: plain color
func GenImageColor(width int, height int, color Color) Image {
	ccolor := *color.cptr()
	res := C.GenImageColor(C.int(width), C.int(height), ccolor)
	return newImageFromPointer(unsafe.Pointer(&res))
}

//GenImageGradientV : Generate image: vertical gradient
func GenImageGradientV(width int, height int, top Color, bottom Color) Image {
	cbottom := *bottom.cptr()
	ctop := *top.cptr()
	res := C.GenImageGradientV(C.int(width), C.int(height), ctop, cbottom)
	return newImageFromPointer(unsafe.Pointer(&res))
}

//GenImageGradientH : Generate image: horizontal gradient
func GenImageGradientH(width int, height int, left Color, right Color) Image {
	cright := *right.cptr()
	cleft := *left.cptr()
	res := C.GenImageGradientH(C.int(width), C.int(height), cleft, cright)
	return newImageFromPointer(unsafe.Pointer(&res))
}

//GenImageGradientRadial : Generate image: radial gradient
func GenImageGradientRadial(width int, height int, density float32, inner Color, outer Color) Image {
	couter := *outer.cptr()
	cinner := *inner.cptr()
	res := C.GenImageGradientRadial(C.int(width), C.int(height), C.float(density), cinner, couter)
	return newImageFromPointer(unsafe.Pointer(&res))
}

//GenImageChecked : Generate image: checked
func GenImageChecked(width int, height int, checksX int, checksY int, col1 Color, col2 Color) Image {
	ccol2 := *col2.cptr()
	ccol1 := *col1.cptr()
	res := C.GenImageChecked(C.int(width), C.int(height), C.int(checksX), C.int(checksY), ccol1, ccol2)
	return newImageFromPointer(unsafe.Pointer(&res))
}

//GenImageWhiteNoise : Generate image: white noise
func GenImageWhiteNoise(width int, height int, factor float32) Image {
	res := C.GenImageWhiteNoise(C.int(width), C.int(height), C.float(factor))
	return newImageFromPointer(unsafe.Pointer(&res))
}

//GenImagePerlinNoise : Generate image: perlin noise
func GenImagePerlinNoise(width int, height int, offsetX int, offsetY int, scale float32) Image {
	res := C.GenImagePerlinNoise(C.int(width), C.int(height), C.int(offsetX), C.int(offsetY), C.float(scale))
	return newImageFromPointer(unsafe.Pointer(&res))
}

//GenImageCellular : Generate image: cellular algorithm. Bigger tileSize means bigger cells
func GenImageCellular(width int, height int, tileSize int) Image {
	res := C.GenImageCellular(C.int(width), C.int(height), C.int(tileSize))
	return newImageFromPointer(unsafe.Pointer(&res))
}

//GenTextureMipmaps : Generate GPU mipmaps for a texture
func GenTextureMipmaps(texture Texture2D) Texture2D {
	ctexture := texture.cptr()
	C.GenTextureMipmaps(&ctexture)
	return newTexture2DFromPointer(unsafe.Pointer(ctexture))
}

//SetTextureFilter : Set texture scaling filter mode
func SetTextureFilter(texture Texture2D, filterMode int) {
	ctexture := *texture.cptr()
	C.SetTextureFilter(ctexture, C.int(filterMode))

}

//SetTextureWrap : Set texture wrapping mode
func SetTextureWrap(texture Texture2D, wrapMode int) {
	ctexture := *texture.cptr()
	C.SetTextureWrap(ctexture, C.int(wrapMode))

}

//DrawTexture : Draw a Texture2D
func DrawTexture(texture Texture2D, posX int, posY int, tint Color) {
	ctint := *tint.cptr()
	ctexture := *texture.cptr()
	C.DrawTexture(ctexture, C.int(posX), C.int(posY), ctint)

}

//DrawTextureV : Draw a Texture2D with position defined as Vector2
func DrawTextureV(texture Texture2D, position Vector2, tint Color) {
	ctint := *tint.cptr()
	cposition := *position.cptr()
	ctexture := *texture.cptr()
	C.DrawTextureV(ctexture, cposition, ctint)

}

//DrawTextureEx : Draw a Texture2D with extended parameters
func DrawTextureEx(texture Texture2D, position Vector2, rotation float32, scale float32, tint Color) {
	ctint := *tint.cptr()
	cposition := *position.cptr()
	ctexture := *texture.cptr()
	C.DrawTextureEx(ctexture, cposition, C.float(rotation), C.float(scale), ctint)

}

//DrawTextureRec : Draw a part of a texture defined by a rectangle
func DrawTextureRec(texture Texture2D, sourceRec Rectangle, position Vector2, tint Color) {
	ctint := *tint.cptr()
	cposition := *position.cptr()
	csourceRec := *sourceRec.cptr()
	ctexture := *texture.cptr()
	C.DrawTextureRec(ctexture, csourceRec, cposition, ctint)

}

//DrawTextureQuad : Draw texture quad with tiling and offset parameters
func DrawTextureQuad(texture Texture2D, tiling Vector2, offset Vector2, quad Rectangle, tint Color) {
	ctint := *tint.cptr()
	cquad := *quad.cptr()
	coffset := *offset.cptr()
	ctiling := *tiling.cptr()
	ctexture := *texture.cptr()
	C.DrawTextureQuad(ctexture, ctiling, coffset, cquad, ctint)

}

//DrawTexturePro : Draw a part of a texture defined by a rectangle with 'pro' parameters
func DrawTexturePro(texture Texture2D, sourceRec Rectangle, destRec Rectangle, origin Vector2, rotation float32, tint Color) {
	ctint := *tint.cptr()
	corigin := *origin.cptr()
	cdestRec := *destRec.cptr()
	csourceRec := *sourceRec.cptr()
	ctexture := *texture.cptr()
	C.DrawTexturePro(ctexture, csourceRec, cdestRec, corigin, C.float(rotation), ctint)

}

//DrawTextureNPatch : Draws a texture (or part of it) that stretches or shrinks nicely
func DrawTextureNPatch(texture Texture2D, nPatchInfo NPatchInfo, destRec Rectangle, origin Vector2, rotation float32, tint Color) {
	ctint := *tint.cptr()
	corigin := *origin.cptr()
	cdestRec := *destRec.cptr()
	cnPatchInfo := *nPatchInfo.cptr()
	ctexture := *texture.cptr()
	C.DrawTextureNPatch(ctexture, cnPatchInfo, cdestRec, corigin, C.float(rotation), ctint)

}

//GetFontDefault : Get the default Font
func GetFontDefault() Font {
	res := C.GetFontDefault()
	return newFontFromPointer(unsafe.Pointer(&res))
}

//LoadFont : Load font from file into GPU memory (VRAM)
func LoadFont(fileName string) Font {
	cfileName := C.CString(fileName)
	defer C.free(unsafe.Pointer(cfileName))
	res := C.LoadFont(cfileName)
	return newFontFromPointer(unsafe.Pointer(&res))
}

//LoadFontEx : Load font from file with extended parameters
func LoadFontEx(fileName string, fontSize int, fontChars int, charsCount int) (Font, int) {
	cfontChars := C.int(fontChars)
	cfileName := C.CString(fileName)
	defer C.free(unsafe.Pointer(cfileName))
	res := C.LoadFontEx(cfileName, C.int(fontSize), &cfontChars, C.int(charsCount))
	return newFontFromPointer(unsafe.Pointer(&res)), int(cfontChars)
}

//LoadFontFromImage : Load font from Image (XNA style)
func LoadFontFromImage(image Image, key Color, firstChar int) Font {
	ckey := *key.cptr()
	cimage := *image.cptr()
	res := C.LoadFontFromImage(cimage, ckey, C.int(firstChar))
	return newFontFromPointer(unsafe.Pointer(&res))
}

//UnloadFont : Unload Font from GPU memory (VRAM)
func UnloadFont(font Font) {
	cfont := *font.cptr()
	C.UnloadFont(cfont)

}

//DrawFPS : Shows current FPS
func DrawFPS(posX int, posY int) {
	C.DrawFPS(C.int(posX), C.int(posY))

}

//DrawText : Draw text (using default font)
func DrawText(text string, posX int, posY int, fontSize int, color Color) {
	ccolor := *color.cptr()
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	C.DrawText(ctext, C.int(posX), C.int(posY), C.int(fontSize), ccolor)

}

//DrawTextEx : Draw text using font and additional parameters
func DrawTextEx(font Font, text string, position Vector2, fontSize float32, spacing float32, tint Color) {
	ctint := *tint.cptr()
	cposition := *position.cptr()
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	cfont := *font.cptr()
	C.DrawTextEx(cfont, ctext, cposition, C.float(fontSize), C.float(spacing), ctint)

}

//DrawTextRec : Draw text using font inside rectangle limits
func DrawTextRec(font Font, text string, rec Rectangle, fontSize float32, spacing float32, wordWrap bool, tint Color) {
	ctint := *tint.cptr()
	crec := *rec.cptr()
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	cfont := *font.cptr()
	C.DrawTextRec(cfont, ctext, crec, C.float(fontSize), C.float(spacing), C.bool(wordWrap), ctint)

}

//MeasureText : Measure string width for default font
func MeasureText(text string, fontSize int) int {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	res := C.MeasureText(ctext, C.int(fontSize))
	return int(&res)
}

//MeasureTextEx : Measure string size for Font
func MeasureTextEx(font Font, text string, fontSize float32, spacing float32) Vector2 {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	cfont := *font.cptr()
	res := C.MeasureTextEx(cfont, ctext, C.float(fontSize), C.float(spacing))
	return newVector2FromPointer(unsafe.Pointer(&res))
}

//GetGlyphIndex : Get index position for a unicode character on font
func GetGlyphIndex(font Font, character int) int {
	cfont := *font.cptr()
	res := C.GetGlyphIndex(cfont, C.int(character))
	return int(&res)
}

//TextIsEqual : Check if two text string are equal
func TextIsEqual(text1 string, text2 string) bool {
	ctext2 := C.CString(text2)
	defer C.free(unsafe.Pointer(ctext2))
	ctext1 := C.CString(text1)
	defer C.free(unsafe.Pointer(ctext1))
	res := C.TextIsEqual(ctext1, ctext2)
	return bool(&res)
}

//TextLength : Get text length, checks for '\0' ending
func TextLength(text string) int {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	res := C.TextLength(ctext)
	return int(&res)
}

//TextSubtext : Get a piece of a text string
func TextSubtext(text string, position int, length int) string {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	res := C.TextSubtext(ctext, C.int(position), C.int(length))
	return C.GoString(&res)
}

//TextReplace : Replace text string (memory must be freed!)
func TextReplace(text string, replace string, by string) string {
	cby := C.CString(by)
	defer C.free(unsafe.Pointer(cby))
	creplace := C.CString(replace)
	defer C.free(unsafe.Pointer(creplace))
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	res := C.TextReplace(ctext, creplace, cby)
	return C.GoString(&res)
}

//TextInsert : Insert text in a position (memory must be freed!)
func TextInsert(text string, insert string, position int) string {
	cinsert := C.CString(insert)
	defer C.free(unsafe.Pointer(cinsert))
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	res := C.TextInsert(ctext, cinsert, C.int(position))
	return C.GoString(&res)
}

//TextJoin :  Join text strings with delimiter
func TextJoin(char string, count int, delimiter string) {
	fmt.Println("test")
}

//TextAppend : Append text at specific position and move cursor!
func TextAppend(text string, append string, position int) int {
	cposition := C.int(position)
	cappend := C.CString(append)
	defer C.free(unsafe.Pointer(cappend))
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	C.TextAppend(ctext, cappend, &cposition)
	return int(cposition)
}

//TextFindIndex : Find first text occurrence within a string
func TextFindIndex(text string, find string) int {
	cfind := C.CString(find)
	defer C.free(unsafe.Pointer(cfind))
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	res := C.TextFindIndex(ctext, cfind)
	return int(&res)
}

//TextToUpper : Get upper case version of provided string
func TextToUpper(text string) string {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	res := C.TextToUpper(ctext)
	return C.GoString(&res)
}

//TextToLower : Get lower case version of provided string
func TextToLower(text string) string {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	res := C.TextToLower(ctext)
	return C.GoString(&res)
}

//TextToPascal : Get Pascal case notation version of provided string
func TextToPascal(text string) string {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	res := C.TextToPascal(ctext)
	return C.GoString(&res)
}

//TextToInteger : Get integer value from text (negative values not supported)
func TextToInteger(text string) int {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	res := C.TextToInteger(ctext)
	return int(&res)
}

//TextToUtf8 : Encode text codepoint into utf8 text (memory must be freed!)
func TextToUtf8(codepoints int, length int) (string, int) {
	ccodepoints := C.int(codepoints)
	res := C.TextToUtf8(&ccodepoints, C.int(length))
	return C.GoString(&res), int(ccodepoints)
}

//GetCodepointsCount : Get total number of characters (codepoints) in a UTF8 encoded string
func GetCodepointsCount(text string) int {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	res := C.GetCodepointsCount(ctext)
	return int(&res)
}

//GetNextCodepoint : Returns next codepoint in a UTF8 encoded string; 0x3f('?') is returned on failure
func GetNextCodepoint(text string, bytesProcessed int) (int, int) {
	cbytesProcessed := C.int(bytesProcessed)
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	res := C.GetNextCodepoint(ctext, &cbytesProcessed)
	return int(&res), int(cbytesProcessed)
}

//CodepointToUtf8 : Encode codepoint into utf8 text (char array length returned as parameter)
func CodepointToUtf8(codepoint int, byteLength int) (string, int) {
	cbyteLength := C.int(byteLength)
	res := C.CodepointToUtf8(C.int(codepoint), &cbyteLength)
	return C.GoString(&res), int(cbyteLength)
}

//DrawLine3D : Draw a line in 3D world space
func DrawLine3D(startPos Vector3, endPos Vector3, color Color) {
	ccolor := *color.cptr()
	cendPos := *endPos.cptr()
	cstartPos := *startPos.cptr()
	C.DrawLine3D(cstartPos, cendPos, ccolor)

}

//DrawCircle3D : Draw a circle in 3D world space
func DrawCircle3D(center Vector3, radius float32, rotationAxis Vector3, rotationAngle float32, color Color) {
	ccolor := *color.cptr()
	crotationAxis := *rotationAxis.cptr()
	ccenter := *center.cptr()
	C.DrawCircle3D(ccenter, C.float(radius), crotationAxis, C.float(rotationAngle), ccolor)

}

//DrawCube : Draw cube
func DrawCube(position Vector3, width float32, height float32, length float32, color Color) {
	ccolor := *color.cptr()
	cposition := *position.cptr()
	C.DrawCube(cposition, C.float(width), C.float(height), C.float(length), ccolor)

}

//DrawCubeV : Draw cube (Vector version)
func DrawCubeV(position Vector3, size Vector3, color Color) {
	ccolor := *color.cptr()
	csize := *size.cptr()
	cposition := *position.cptr()
	C.DrawCubeV(cposition, csize, ccolor)

}

//DrawCubeWires : Draw cube wires
func DrawCubeWires(position Vector3, width float32, height float32, length float32, color Color) {
	ccolor := *color.cptr()
	cposition := *position.cptr()
	C.DrawCubeWires(cposition, C.float(width), C.float(height), C.float(length), ccolor)

}

//DrawCubeWiresV : Draw cube wires (Vector version)
func DrawCubeWiresV(position Vector3, size Vector3, color Color) {
	ccolor := *color.cptr()
	csize := *size.cptr()
	cposition := *position.cptr()
	C.DrawCubeWiresV(cposition, csize, ccolor)

}

//DrawCubeTexture : Draw cube textured
func DrawCubeTexture(texture Texture2D, position Vector3, width float32, height float32, length float32, color Color) {
	ccolor := *color.cptr()
	cposition := *position.cptr()
	ctexture := *texture.cptr()
	C.DrawCubeTexture(ctexture, cposition, C.float(width), C.float(height), C.float(length), ccolor)

}

//DrawSphere : Draw sphere
func DrawSphere(centerPos Vector3, radius float32, color Color) {
	ccolor := *color.cptr()
	ccenterPos := *centerPos.cptr()
	C.DrawSphere(ccenterPos, C.float(radius), ccolor)

}

//DrawSphereEx : Draw sphere with extended parameters
func DrawSphereEx(centerPos Vector3, radius float32, rings int, slices int, color Color) {
	ccolor := *color.cptr()
	ccenterPos := *centerPos.cptr()
	C.DrawSphereEx(ccenterPos, C.float(radius), C.int(rings), C.int(slices), ccolor)

}

//DrawSphereWires : Draw sphere wires
func DrawSphereWires(centerPos Vector3, radius float32, rings int, slices int, color Color) {
	ccolor := *color.cptr()
	ccenterPos := *centerPos.cptr()
	C.DrawSphereWires(ccenterPos, C.float(radius), C.int(rings), C.int(slices), ccolor)

}

//DrawCylinder : Draw a cylinder/cone
func DrawCylinder(position Vector3, radiusTop float32, radiusBottom float32, height float32, slices int, color Color) {
	ccolor := *color.cptr()
	cposition := *position.cptr()
	C.DrawCylinder(cposition, C.float(radiusTop), C.float(radiusBottom), C.float(height), C.int(slices), ccolor)

}

//DrawCylinderWires : Draw a cylinder/cone wires
func DrawCylinderWires(position Vector3, radiusTop float32, radiusBottom float32, height float32, slices int, color Color) {
	ccolor := *color.cptr()
	cposition := *position.cptr()
	C.DrawCylinderWires(cposition, C.float(radiusTop), C.float(radiusBottom), C.float(height), C.int(slices), ccolor)

}

//DrawPlane : Draw a plane XZ
func DrawPlane(centerPos Vector3, size Vector2, color Color) {
	ccolor := *color.cptr()
	csize := *size.cptr()
	ccenterPos := *centerPos.cptr()
	C.DrawPlane(ccenterPos, csize, ccolor)

}

//DrawRay : Draw a ray line
func DrawRay(ray Ray, color Color) {
	ccolor := *color.cptr()
	cray := *ray.cptr()
	C.DrawRay(cray, ccolor)

}

//DrawGrid : Draw a grid (centered at (0, 0, 0))
func DrawGrid(slices int, spacing float32) {
	C.DrawGrid(C.int(slices), C.float(spacing))

}

//DrawGizmo : Draw simple gizmo
func DrawGizmo(position Vector3) {
	cposition := *position.cptr()
	C.DrawGizmo(cposition)

}

//LoadModel : Load model from files (meshes and materials)
func LoadModel(fileName string) Model {
	cfileName := C.CString(fileName)
	defer C.free(unsafe.Pointer(cfileName))
	res := C.LoadModel(cfileName)
	return newModelFromPointer(unsafe.Pointer(&res))
}

//LoadModelFromMesh : Load model from generated mesh (default material)
func LoadModelFromMesh(mesh Mesh) Model {
	cmesh := *mesh.cptr()
	res := C.LoadModelFromMesh(cmesh)
	return newModelFromPointer(unsafe.Pointer(&res))
}

//UnloadModel : Unload model from memory (RAM and/or VRAM)
func UnloadModel(model Model) {
	cmodel := *model.cptr()
	C.UnloadModel(cmodel)

}

//ExportMesh : Export mesh data to file
func ExportMesh(mesh Mesh, fileName string) {
	cfileName := C.CString(fileName)
	defer C.free(unsafe.Pointer(cfileName))
	cmesh := *mesh.cptr()
	C.ExportMesh(cmesh, cfileName)

}

//UnloadMesh : Unload mesh from memory (RAM and/or VRAM)
func UnloadMesh(mesh Mesh) {
	cmesh := *mesh.cptr()
	C.UnloadMesh(cmesh)

}

//LoadMaterialDefault : Load default material (Supports: DIFFUSE, SPECULAR, NORMAL maps)
func LoadMaterialDefault() Material {
	res := C.LoadMaterialDefault()
	return newMaterialFromPointer(unsafe.Pointer(&res))
}

//UnloadMaterial : Unload material from GPU memory (VRAM)
func UnloadMaterial(material Material) {
	cmaterial := *material.cptr()
	C.UnloadMaterial(cmaterial)

}

//SetMaterialTexture : Set texture for a material map type (MAP_DIFFUSE, MAP_SPECULAR...)
func SetMaterialTexture(material Material, mapType int, texture Texture2D) Material {
	ctexture := *texture.cptr()
	cmaterial := material.cptr()
	C.SetMaterialTexture(&cmaterial, C.int(mapType), ctexture)
	return newMaterialFromPointer(unsafe.Pointer(cmaterial))
}

//SetModelMeshMaterial : Set material for a mesh
func SetModelMeshMaterial(model Model, meshId int, materialId int) Model {
	cmodel := model.cptr()
	C.SetModelMeshMaterial(&cmodel, C.int(meshId), C.int(materialId))
	return newModelFromPointer(unsafe.Pointer(cmodel))
}

//UpdateModelAnimation : Update model animation pose
func UpdateModelAnimation(model Model, anim ModelAnimation, frame int) {
	canim := *anim.cptr()
	cmodel := *model.cptr()
	C.UpdateModelAnimation(cmodel, canim, C.int(frame))

}

//UnloadModelAnimation : Unload animation data
func UnloadModelAnimation(anim ModelAnimation) {
	canim := *anim.cptr()
	C.UnloadModelAnimation(canim)

}

//IsModelAnimationValid : Check model animation skeleton match
func IsModelAnimationValid(model Model, anim ModelAnimation) bool {
	canim := *anim.cptr()
	cmodel := *model.cptr()
	res := C.IsModelAnimationValid(cmodel, canim)
	return bool(&res)
}

//GenMeshPoly : Generate polygonal mesh
func GenMeshPoly(sides int, radius float32) Mesh {
	res := C.GenMeshPoly(C.int(sides), C.float(radius))
	return newMeshFromPointer(unsafe.Pointer(&res))
}

//GenMeshPlane : Generate plane mesh (with subdivisions)
func GenMeshPlane(width float32, length float32, resX int, resZ int) Mesh {
	res := C.GenMeshPlane(C.float(width), C.float(length), C.int(resX), C.int(resZ))
	return newMeshFromPointer(unsafe.Pointer(&res))
}

//GenMeshCube : Generate cuboid mesh
func GenMeshCube(width float32, height float32, length float32) Mesh {
	res := C.GenMeshCube(C.float(width), C.float(height), C.float(length))
	return newMeshFromPointer(unsafe.Pointer(&res))
}

//GenMeshSphere : Generate sphere mesh (standard sphere)
func GenMeshSphere(radius float32, rings int, slices int) Mesh {
	res := C.GenMeshSphere(C.float(radius), C.int(rings), C.int(slices))
	return newMeshFromPointer(unsafe.Pointer(&res))
}

//GenMeshHemiSphere : Generate half-sphere mesh (no bottom cap)
func GenMeshHemiSphere(radius float32, rings int, slices int) Mesh {
	res := C.GenMeshHemiSphere(C.float(radius), C.int(rings), C.int(slices))
	return newMeshFromPointer(unsafe.Pointer(&res))
}

//GenMeshCylinder : Generate cylinder mesh
func GenMeshCylinder(radius float32, height float32, slices int) Mesh {
	res := C.GenMeshCylinder(C.float(radius), C.float(height), C.int(slices))
	return newMeshFromPointer(unsafe.Pointer(&res))
}

//GenMeshTorus : Generate torus mesh
func GenMeshTorus(radius float32, size float32, radSeg int, sides int) Mesh {
	res := C.GenMeshTorus(C.float(radius), C.float(size), C.int(radSeg), C.int(sides))
	return newMeshFromPointer(unsafe.Pointer(&res))
}

//GenMeshKnot : Generate trefoil knot mesh
func GenMeshKnot(radius float32, size float32, radSeg int, sides int) Mesh {
	res := C.GenMeshKnot(C.float(radius), C.float(size), C.int(radSeg), C.int(sides))
	return newMeshFromPointer(unsafe.Pointer(&res))
}

//GenMeshHeightmap : Generate heightmap mesh from image data
func GenMeshHeightmap(heightmap Image, size Vector3) Mesh {
	csize := *size.cptr()
	cheightmap := *heightmap.cptr()
	res := C.GenMeshHeightmap(cheightmap, csize)
	return newMeshFromPointer(unsafe.Pointer(&res))
}

//GenMeshCubicmap : Generate cubes-based map mesh from image data
func GenMeshCubicmap(cubicmap Image, cubeSize Vector3) Mesh {
	ccubeSize := *cubeSize.cptr()
	ccubicmap := *cubicmap.cptr()
	res := C.GenMeshCubicmap(ccubicmap, ccubeSize)
	return newMeshFromPointer(unsafe.Pointer(&res))
}

//MeshBoundingBox : Compute mesh bounding box limits
func MeshBoundingBox(mesh Mesh) BoundingBox {
	cmesh := *mesh.cptr()
	res := C.MeshBoundingBox(cmesh)
	return newBoundingBoxFromPointer(unsafe.Pointer(&res))
}

//MeshTangents : Compute mesh tangents
func MeshTangents(mesh Mesh) Mesh {
	cmesh := mesh.cptr()
	C.MeshTangents(&cmesh)
	return newMeshFromPointer(unsafe.Pointer(cmesh))
}

//MeshBinormals : Compute mesh binormals
func MeshBinormals(mesh Mesh) Mesh {
	cmesh := mesh.cptr()
	C.MeshBinormals(&cmesh)
	return newMeshFromPointer(unsafe.Pointer(cmesh))
}

//DrawModel : Draw a model (with texture if set)
func DrawModel(model Model, position Vector3, scale float32, tint Color) {
	ctint := *tint.cptr()
	cposition := *position.cptr()
	cmodel := *model.cptr()
	C.DrawModel(cmodel, cposition, C.float(scale), ctint)

}

//DrawModelEx : Draw a model with extended parameters
func DrawModelEx(model Model, position Vector3, rotationAxis Vector3, rotationAngle float32, scale Vector3, tint Color) {
	ctint := *tint.cptr()
	cscale := *scale.cptr()
	crotationAxis := *rotationAxis.cptr()
	cposition := *position.cptr()
	cmodel := *model.cptr()
	C.DrawModelEx(cmodel, cposition, crotationAxis, C.float(rotationAngle), cscale, ctint)

}

//DrawModelWires : Draw a model wires (with texture if set)
func DrawModelWires(model Model, position Vector3, scale float32, tint Color) {
	ctint := *tint.cptr()
	cposition := *position.cptr()
	cmodel := *model.cptr()
	C.DrawModelWires(cmodel, cposition, C.float(scale), ctint)

}

//DrawModelWiresEx : Draw a model wires (with texture if set) with extended parameters
func DrawModelWiresEx(model Model, position Vector3, rotationAxis Vector3, rotationAngle float32, scale Vector3, tint Color) {
	ctint := *tint.cptr()
	cscale := *scale.cptr()
	crotationAxis := *rotationAxis.cptr()
	cposition := *position.cptr()
	cmodel := *model.cptr()
	C.DrawModelWiresEx(cmodel, cposition, crotationAxis, C.float(rotationAngle), cscale, ctint)

}

//DrawBoundingBox : Draw bounding box (wires)
func DrawBoundingBox(box BoundingBox, color Color) {
	ccolor := *color.cptr()
	cbox := *box.cptr()
	C.DrawBoundingBox(cbox, ccolor)

}

//DrawBillboard : Draw a billboard texture
func DrawBillboard(camera Camera, texture Texture2D, center Vector3, size float32, tint Color) {
	ctint := *tint.cptr()
	ccenter := *center.cptr()
	ctexture := *texture.cptr()
	ccamera := *camera.cptr()
	C.DrawBillboard(ccamera, ctexture, ccenter, C.float(size), ctint)

}

//DrawBillboardRec : Draw a billboard texture defined by sourceRec
func DrawBillboardRec(camera Camera, texture Texture2D, sourceRec Rectangle, center Vector3, size float32, tint Color) {
	ctint := *tint.cptr()
	ccenter := *center.cptr()
	csourceRec := *sourceRec.cptr()
	ctexture := *texture.cptr()
	ccamera := *camera.cptr()
	C.DrawBillboardRec(ccamera, ctexture, csourceRec, ccenter, C.float(size), ctint)

}

//CheckCollisionSpheres : Detect collision between two spheres
func CheckCollisionSpheres(centerA Vector3, radiusA float32, centerB Vector3, radiusB float32) bool {
	ccenterB := *centerB.cptr()
	ccenterA := *centerA.cptr()
	res := C.CheckCollisionSpheres(ccenterA, C.float(radiusA), ccenterB, C.float(radiusB))
	return bool(&res)
}

//CheckCollisionBoxes : Detect collision between two bounding boxes
func CheckCollisionBoxes(box1 BoundingBox, box2 BoundingBox) bool {
	cbox2 := *box2.cptr()
	cbox1 := *box1.cptr()
	res := C.CheckCollisionBoxes(cbox1, cbox2)
	return bool(&res)
}

//CheckCollisionBoxSphere : Detect collision between box and sphere
func CheckCollisionBoxSphere(box BoundingBox, center Vector3, radius float32) bool {
	ccenter := *center.cptr()
	cbox := *box.cptr()
	res := C.CheckCollisionBoxSphere(cbox, ccenter, C.float(radius))
	return bool(&res)
}

//CheckCollisionRaySphere : Detect collision between ray and sphere
func CheckCollisionRaySphere(ray Ray, center Vector3, radius float32) bool {
	ccenter := *center.cptr()
	cray := *ray.cptr()
	res := C.CheckCollisionRaySphere(cray, ccenter, C.float(radius))
	return bool(&res)
}

//CheckCollisionRaySphereEx : Detect collision between ray and sphere, returns collision point
func CheckCollisionRaySphereEx(ray Ray, center Vector3, radius float32, collisionPoint Vector3) (bool, Vector3) {
	ccollisionPoint := collisionPoint.cptr()
	ccenter := *center.cptr()
	cray := *ray.cptr()
	res := C.CheckCollisionRaySphereEx(cray, ccenter, C.float(radius), &ccollisionPoint)
	return bool(&res), newVector3FromPointer(unsafe.Pointer(ccollisionPoint))
}

//CheckCollisionRayBox : Detect collision between ray and box
func CheckCollisionRayBox(ray Ray, box BoundingBox) bool {
	cbox := *box.cptr()
	cray := *ray.cptr()
	res := C.CheckCollisionRayBox(cray, cbox)
	return bool(&res)
}

//GetCollisionRayModel : Get collision info between ray and model
func GetCollisionRayModel(ray Ray, model Model) RayHitInfo {
	cmodel := *model.cptr()
	cray := *ray.cptr()
	res := C.GetCollisionRayModel(cray, cmodel)
	return newRayHitInfoFromPointer(unsafe.Pointer(&res))
}

//GetCollisionRayTriangle : Get collision info between ray and triangle
func GetCollisionRayTriangle(ray Ray, p1 Vector3, p2 Vector3, p3 Vector3) RayHitInfo {
	cp3 := *p3.cptr()
	cp2 := *p2.cptr()
	cp1 := *p1.cptr()
	cray := *ray.cptr()
	res := C.GetCollisionRayTriangle(cray, cp1, cp2, cp3)
	return newRayHitInfoFromPointer(unsafe.Pointer(&res))
}

//GetCollisionRayGround : Get collision info between ray and ground plane (Y-normal plane)
func GetCollisionRayGround(ray Ray, groundHeight float32) RayHitInfo {
	cray := *ray.cptr()
	res := C.GetCollisionRayGround(cray, C.float(groundHeight))
	return newRayHitInfoFromPointer(unsafe.Pointer(&res))
}

//LoadText : Load chars array from text file
func LoadText(fileName string) string {
	cfileName := C.CString(fileName)
	defer C.free(unsafe.Pointer(cfileName))
	res := C.LoadText(cfileName)
	return C.GoString(&res)
}

//LoadShader : Load shader from files and bind default locations
func LoadShader(vsFileName string, fsFileName string) Shader {
	cfsFileName := C.CString(fsFileName)
	defer C.free(unsafe.Pointer(cfsFileName))
	cvsFileName := C.CString(vsFileName)
	defer C.free(unsafe.Pointer(cvsFileName))
	res := C.LoadShader(cvsFileName, cfsFileName)
	return newShaderFromPointer(unsafe.Pointer(&res))
}

//LoadShaderCode : Load shader from code strings and bind default locations
func LoadShaderCode(vsCode string, fsCode string) Shader {
	cfsCode := C.CString(fsCode)
	defer C.free(unsafe.Pointer(cfsCode))
	cvsCode := C.CString(vsCode)
	defer C.free(unsafe.Pointer(cvsCode))
	res := C.LoadShaderCode(cvsCode, cfsCode)
	return newShaderFromPointer(unsafe.Pointer(&res))
}

//UnloadShader : Unload shader from GPU memory (VRAM)
func UnloadShader(shader Shader) {
	cshader := *shader.cptr()
	C.UnloadShader(cshader)

}

//GetShaderDefault : Get default shader
func GetShaderDefault() Shader {
	res := C.GetShaderDefault()
	return newShaderFromPointer(unsafe.Pointer(&res))
}

//GetTextureDefault : Get default texture
func GetTextureDefault() Texture2D {
	res := C.GetTextureDefault()
	return newTexture2DFromPointer(unsafe.Pointer(&res))
}

//GetShaderLocation : Get shader uniform location
func GetShaderLocation(shader Shader, uniformName string) int {
	cuniformName := C.CString(uniformName)
	defer C.free(unsafe.Pointer(cuniformName))
	cshader := *shader.cptr()
	res := C.GetShaderLocation(cshader, cuniformName)
	return int(&res)
}

//SetShaderValue : Set shader uniform value
func SetShaderValue(shader Shader, uniformLoc int, value unsafe.Pointer, uniformType int) {
	cshader := *shader.cptr()
	C.SetShaderValue(cshader, C.int(uniformLoc), value, C.int(uniformType))

}

//SetShaderValueV : Set shader uniform value vector
func SetShaderValueV(shader Shader, uniformLoc int, value unsafe.Pointer, uniformType int, count int) {
	cshader := *shader.cptr()
	C.SetShaderValueV(cshader, C.int(uniformLoc), value, C.int(uniformType), C.int(count))

}

//SetShaderValueMatrix : Set shader uniform value (matrix 4x4)
func SetShaderValueMatrix(shader Shader, uniformLoc int, mat Matrix) {
	cmat := *mat.cptr()
	cshader := *shader.cptr()
	C.SetShaderValueMatrix(cshader, C.int(uniformLoc), cmat)

}

//SetShaderValueTexture : Set shader uniform value for texture
func SetShaderValueTexture(shader Shader, uniformLoc int, texture Texture2D) {
	ctexture := *texture.cptr()
	cshader := *shader.cptr()
	C.SetShaderValueTexture(cshader, C.int(uniformLoc), ctexture)

}

//SetMatrixProjection : Set a custom projection matrix (replaces internal projection matrix)
func SetMatrixProjection(proj Matrix) {
	cproj := *proj.cptr()
	C.SetMatrixProjection(cproj)

}

//SetMatrixModelview : Set a custom modelview matrix (replaces internal modelview matrix)
func SetMatrixModelview(view Matrix) {
	cview := *view.cptr()
	C.SetMatrixModelview(cview)

}

//GetMatrixModelview : Get internal modelview matrix
func GetMatrixModelview() Matrix {
	res := C.GetMatrixModelview()
	return newMatrixFromPointer(unsafe.Pointer(&res))
}

//GetMatrixProjection : Get internal projection matrix
func GetMatrixProjection() Matrix {
	res := C.GetMatrixProjection()
	return newMatrixFromPointer(unsafe.Pointer(&res))
}

//GenTextureCubemap : Generate cubemap texture from HDR texture
func GenTextureCubemap(shader Shader, skyHDR Texture2D, size int) Texture2D {
	cskyHDR := *skyHDR.cptr()
	cshader := *shader.cptr()
	res := C.GenTextureCubemap(cshader, cskyHDR, C.int(size))
	return newTexture2DFromPointer(unsafe.Pointer(&res))
}

//GenTextureIrradiance : Generate irradiance texture using cubemap data
func GenTextureIrradiance(shader Shader, cubemap Texture2D, size int) Texture2D {
	ccubemap := *cubemap.cptr()
	cshader := *shader.cptr()
	res := C.GenTextureIrradiance(cshader, ccubemap, C.int(size))
	return newTexture2DFromPointer(unsafe.Pointer(&res))
}

//GenTexturePrefilter : Generate prefilter texture using cubemap data
func GenTexturePrefilter(shader Shader, cubemap Texture2D, size int) Texture2D {
	ccubemap := *cubemap.cptr()
	cshader := *shader.cptr()
	res := C.GenTexturePrefilter(cshader, ccubemap, C.int(size))
	return newTexture2DFromPointer(unsafe.Pointer(&res))
}

//GenTextureBRDF : Generate BRDF texture
func GenTextureBRDF(shader Shader, size int) Texture2D {
	cshader := *shader.cptr()
	res := C.GenTextureBRDF(cshader, C.int(size))
	return newTexture2DFromPointer(unsafe.Pointer(&res))
}

//BeginShaderMode : Begin custom shader drawing
func BeginShaderMode(shader Shader) {
	cshader := *shader.cptr()
	C.BeginShaderMode(cshader)

}

//EndShaderMode : End custom shader drawing (use default shader)
func EndShaderMode() {
	C.EndShaderMode()

}

//BeginBlendMode : Begin blending mode (alpha, additive, multiplied)
func BeginBlendMode(mode int) {
	C.BeginBlendMode(C.int(mode))

}

//EndBlendMode : End blending mode (reset to default: alpha blending)
func EndBlendMode() {
	C.EndBlendMode()

}

//InitVrSimulator : Init VR simulator for selected device parameters
func InitVrSimulator() {
	C.InitVrSimulator()

}

//CloseVrSimulator : Close VR simulator for current device
func CloseVrSimulator() {
	C.CloseVrSimulator()

}

//UpdateVrTracking : Update VR tracking (position and orientation) and camera
func UpdateVrTracking(camera Camera) Camera {
	ccamera := camera.cptr()
	C.UpdateVrTracking(&ccamera)
	return newCameraFromPointer(unsafe.Pointer(ccamera))
}

//SetVrConfiguration : Set stereo rendering configuration parameters
func SetVrConfiguration(info VrDeviceInfo, distortion Shader) {
	cdistortion := *distortion.cptr()
	cinfo := *info.cptr()
	C.SetVrConfiguration(cinfo, cdistortion)

}

//IsVrSimulatorReady : Detect if VR simulator is ready
func IsVrSimulatorReady() bool {
	res := C.IsVrSimulatorReady()
	return bool(&res)
}

//ToggleVrMode : Enable/Disable VR experience
func ToggleVrMode() {
	C.ToggleVrMode()

}

//BeginVrDrawing : Begin VR simulator stereo rendering
func BeginVrDrawing() {
	C.BeginVrDrawing()

}

//EndVrDrawing : End VR simulator stereo rendering
func EndVrDrawing() {
	C.EndVrDrawing()

}

//InitAudioDevice : Initialize audio device and context
func InitAudioDevice() {
	C.InitAudioDevice()

}

//CloseAudioDevice : Close the audio device and context
func CloseAudioDevice() {
	C.CloseAudioDevice()

}

//IsAudioDeviceReady : Check if audio device has been initialized successfully
func IsAudioDeviceReady() bool {
	res := C.IsAudioDeviceReady()
	return bool(&res)
}

//SetMasterVolume : Set master volume (listener)
func SetMasterVolume(volume float32) {
	C.SetMasterVolume(C.float(volume))

}

//LoadWave : Load wave data from file
func LoadWave(fileName string) Wave {
	cfileName := C.CString(fileName)
	defer C.free(unsafe.Pointer(cfileName))
	res := C.LoadWave(cfileName)
	return newWaveFromPointer(unsafe.Pointer(&res))
}

//LoadSound : Load sound from file
func LoadSound(fileName string) Sound {
	cfileName := C.CString(fileName)
	defer C.free(unsafe.Pointer(cfileName))
	res := C.LoadSound(cfileName)
	return newSoundFromPointer(unsafe.Pointer(&res))
}

//LoadSoundFromWave : Load sound from wave data
func LoadSoundFromWave(wave Wave) Sound {
	cwave := *wave.cptr()
	res := C.LoadSoundFromWave(cwave)
	return newSoundFromPointer(unsafe.Pointer(&res))
}

//UpdateSound : Update sound buffer with new data
func UpdateSound(sound Sound, data unsafe.Pointer, samplesCount int) {
	csound := *sound.cptr()
	C.UpdateSound(csound, data, C.int(samplesCount))

}

//UnloadWave : Unload wave data
func UnloadWave(wave Wave) {
	cwave := *wave.cptr()
	C.UnloadWave(cwave)

}

//UnloadSound : Unload sound
func UnloadSound(sound Sound) {
	csound := *sound.cptr()
	C.UnloadSound(csound)

}

//ExportWave : Export wave data to file
func ExportWave(wave Wave, fileName string) {
	cfileName := C.CString(fileName)
	defer C.free(unsafe.Pointer(cfileName))
	cwave := *wave.cptr()
	C.ExportWave(cwave, cfileName)

}

//ExportWaveAsCode : Export wave sample data to code (.h)
func ExportWaveAsCode(wave Wave, fileName string) {
	cfileName := C.CString(fileName)
	defer C.free(unsafe.Pointer(cfileName))
	cwave := *wave.cptr()
	C.ExportWaveAsCode(cwave, cfileName)

}

//PlaySound : Play a sound
func PlaySound(sound Sound) {
	csound := *sound.cptr()
	C.PlaySound(csound)

}

//StopSound : Stop playing a sound
func StopSound(sound Sound) {
	csound := *sound.cptr()
	C.StopSound(csound)

}

//PauseSound : Pause a sound
func PauseSound(sound Sound) {
	csound := *sound.cptr()
	C.PauseSound(csound)

}

//ResumeSound : Resume a paused sound
func ResumeSound(sound Sound) {
	csound := *sound.cptr()
	C.ResumeSound(csound)

}

//PlaySoundMulti : Play a sound (using multichannel buffer pool)
func PlaySoundMulti(sound Sound) {
	csound := *sound.cptr()
	C.PlaySoundMulti(csound)

}

//StopSoundMulti : Stop any sound playing (using multichannel buffer pool)
func StopSoundMulti() {
	C.StopSoundMulti()

}

//GetSoundsPlaying : Get number of sounds playing in the multichannel
func GetSoundsPlaying() int {
	res := C.GetSoundsPlaying()
	return int(&res)
}

//IsSoundPlaying : Check if a sound is currently playing
func IsSoundPlaying(sound Sound) bool {
	csound := *sound.cptr()
	res := C.IsSoundPlaying(csound)
	return bool(&res)
}

//SetSoundVolume : Set volume for a sound (1.0 is max level)
func SetSoundVolume(sound Sound, volume float32) {
	csound := *sound.cptr()
	C.SetSoundVolume(csound, C.float(volume))

}

//SetSoundPitch : Set pitch for a sound (1.0 is base level)
func SetSoundPitch(sound Sound, pitch float32) {
	csound := *sound.cptr()
	C.SetSoundPitch(csound, C.float(pitch))

}

//WaveFormat : Convert wave data to desired format
func WaveFormat(wave Wave, sampleRate int, sampleSize int, channels int) Wave {
	cwave := wave.cptr()
	C.WaveFormat(&cwave, C.int(sampleRate), C.int(sampleSize), C.int(channels))
	return newWaveFromPointer(unsafe.Pointer(cwave))
}

//WaveCopy : Copy a wave to a new wave
func WaveCopy(wave Wave) Wave {
	cwave := *wave.cptr()
	res := C.WaveCopy(cwave)
	return newWaveFromPointer(unsafe.Pointer(&res))
}

//WaveCrop : Crop a wave to defined samples range
func WaveCrop(wave Wave, initSample int, finalSample int) Wave {
	cwave := wave.cptr()
	C.WaveCrop(&cwave, C.int(initSample), C.int(finalSample))
	return newWaveFromPointer(unsafe.Pointer(cwave))
}

//LoadMusicStream : Load music stream from file
func LoadMusicStream(fileName string) Music {
	cfileName := C.CString(fileName)
	defer C.free(unsafe.Pointer(cfileName))
	res := C.LoadMusicStream(cfileName)
	return newMusicFromPointer(unsafe.Pointer(&res))
}

//UnloadMusicStream : Unload music stream
func UnloadMusicStream(music Music) {
	cmusic := *music.cptr()
	C.UnloadMusicStream(cmusic)

}

//PlayMusicStream : Start music playing
func PlayMusicStream(music Music) {
	cmusic := *music.cptr()
	C.PlayMusicStream(cmusic)

}

//UpdateMusicStream : Updates buffers for music streaming
func UpdateMusicStream(music Music) {
	cmusic := *music.cptr()
	C.UpdateMusicStream(cmusic)

}

//StopMusicStream : Stop music playing
func StopMusicStream(music Music) {
	cmusic := *music.cptr()
	C.StopMusicStream(cmusic)

}

//PauseMusicStream : Pause music playing
func PauseMusicStream(music Music) {
	cmusic := *music.cptr()
	C.PauseMusicStream(cmusic)

}

//ResumeMusicStream : Resume playing paused music
func ResumeMusicStream(music Music) {
	cmusic := *music.cptr()
	C.ResumeMusicStream(cmusic)

}

//IsMusicPlaying : Check if music is playing
func IsMusicPlaying(music Music) bool {
	cmusic := *music.cptr()
	res := C.IsMusicPlaying(cmusic)
	return bool(&res)
}

//SetMusicVolume : Set volume for music (1.0 is max level)
func SetMusicVolume(music Music, volume float32) {
	cmusic := *music.cptr()
	C.SetMusicVolume(cmusic, C.float(volume))

}

//SetMusicPitch : Set pitch for a music (1.0 is base level)
func SetMusicPitch(music Music, pitch float32) {
	cmusic := *music.cptr()
	C.SetMusicPitch(cmusic, C.float(pitch))

}

//SetMusicLoopCount : Set music loop count (loop repeats)
func SetMusicLoopCount(music Music, count int) {
	cmusic := *music.cptr()
	C.SetMusicLoopCount(cmusic, C.int(count))

}

//GetMusicTimeLength : Get music time length (in seconds)
func GetMusicTimeLength(music Music) float32 {
	cmusic := *music.cptr()
	res := C.GetMusicTimeLength(cmusic)
	return float32(&res)
}

//GetMusicTimePlayed : Get current music time played (in seconds)
func GetMusicTimePlayed(music Music) float32 {
	cmusic := *music.cptr()
	res := C.GetMusicTimePlayed(cmusic)
	return float32(&res)
}

//InitAudioStream : Init audio stream (to stream raw audio pcm data)
func InitAudioStream(sampleRate int, sampleSize int, channels int) AudioStream {
	res := C.InitAudioStream(C.int(sampleRate), C.int(sampleSize), C.int(channels))
	return newAudioStreamFromPointer(unsafe.Pointer(&res))
}

//UpdateAudioStream : Update audio stream buffers with data
func UpdateAudioStream(stream AudioStream, data unsafe.Pointer, samplesCount int) {
	cstream := *stream.cptr()
	C.UpdateAudioStream(cstream, data, C.int(samplesCount))

}

//CloseAudioStream : Close audio stream and free memory
func CloseAudioStream(stream AudioStream) {
	cstream := *stream.cptr()
	C.CloseAudioStream(cstream)

}

//IsAudioStreamProcessed : Check if any audio stream buffers requires refill
func IsAudioStreamProcessed(stream AudioStream) bool {
	cstream := *stream.cptr()
	res := C.IsAudioStreamProcessed(cstream)
	return bool(&res)
}

//PlayAudioStream : Play audio stream
func PlayAudioStream(stream AudioStream) {
	cstream := *stream.cptr()
	C.PlayAudioStream(cstream)

}

//PauseAudioStream : Pause audio stream
func PauseAudioStream(stream AudioStream) {
	cstream := *stream.cptr()
	C.PauseAudioStream(cstream)

}

//ResumeAudioStream : Resume audio stream
func ResumeAudioStream(stream AudioStream) {
	cstream := *stream.cptr()
	C.ResumeAudioStream(cstream)

}

//IsAudioStreamPlaying : Check if audio stream is playing
func IsAudioStreamPlaying(stream AudioStream) bool {
	cstream := *stream.cptr()
	res := C.IsAudioStreamPlaying(cstream)
	return bool(&res)
}

//StopAudioStream : Stop audio stream
func StopAudioStream(stream AudioStream) {
	cstream := *stream.cptr()
	C.StopAudioStream(cstream)

}

//SetAudioStreamVolume : Set volume for audio stream (1.0 is max level)
func SetAudioStreamVolume(stream AudioStream, volume float32) {
	cstream := *stream.cptr()
	C.SetAudioStreamVolume(cstream, C.float(volume))

}

//SetAudioStreamPitch : Set pitch for audio stream (1.0 is base level)
func SetAudioStreamPitch(stream AudioStream, pitch float32) {
	cstream := *stream.cptr()
	C.SetAudioStreamPitch(cstream, C.float(pitch))

}

//GuiEnable : Enable gui controls (global state)
func GuiEnable() {
	C.GuiEnable()

}

//GuiDisable : Disable gui controls (global state)
func GuiDisable() {
	C.GuiDisable()

}

//GuiLock : Lock gui controls (global state)
func GuiLock() {
	C.GuiLock()

}

//GuiUnlock : Unlock gui controls (global state)
func GuiUnlock() {
	C.GuiUnlock()

}

//GuiFade : Set gui controls alpha (global state), alpha goes from 0.0f to 1.0f
func GuiFade(alpha float32) {
	C.GuiFade(C.float(alpha))

}

//GuiSetState : Set gui state (global state)
func GuiSetState(state int) {
	C.GuiSetState(C.int(state))

}

//GuiGetState : Get gui state (global state)
func GuiGetState() int {
	res := C.GuiGetState()
	return int(&res)
}

//GuiSetFont : Set gui custom font (global state)
func GuiSetFont(font Font) {
	cfont := *font.cptr()
	C.GuiSetFont(cfont)

}

//GuiGetFont : Get gui custom font (global state)
func GuiGetFont() Font {
	res := C.GuiGetFont()
	return newFontFromPointer(unsafe.Pointer(&res))
}

//GuiSetStyle : Set one style property
func GuiSetStyle(control int, property int, value int) {
	C.GuiSetStyle(C.int(control), C.int(property), C.int(value))

}

//GuiGetStyle : Get one style property
func GuiGetStyle(control int, property int) int {
	res := C.GuiGetStyle(C.int(control), C.int(property))
	return int(&res)
}

//GuiTextBoxSetActive : Sets the active textbox
func GuiTextBoxSetActive(bounds Rectangle) {
	cbounds := *bounds.cptr()
	C.GuiTextBoxSetActive(cbounds)

}

//GuiTextBoxGetActive : Get bounds of active textbox
func GuiTextBoxGetActive() Rectangle {
	res := C.GuiTextBoxGetActive()
	return newRectangleFromPointer(unsafe.Pointer(&res))
}

//GuiTextBoxSetCursor : Set cursor position of active textbox
func GuiTextBoxSetCursor(cursor int) {
	C.GuiTextBoxSetCursor(C.int(cursor))

}

//GuiTextBoxGetCursor : Get cursor position of active textbox
func GuiTextBoxGetCursor() int {
	res := C.GuiTextBoxGetCursor()
	return int(&res)
}

//GuiTextBoxSetSelection : Set selection of active textbox
func GuiTextBoxSetSelection(start int, length int) {
	C.GuiTextBoxSetSelection(C.int(start), C.int(length))

}

//GuiTextBoxGetSelection : Get selection of active textbox (x - selection start  y - selection length)
func GuiTextBoxGetSelection() Vector2 {
	res := C.GuiTextBoxGetSelection()
	return newVector2FromPointer(unsafe.Pointer(&res))
}

//GuiTextBoxIsActive : Returns true if a textbox control with specified `bounds` is the active textbox
func GuiTextBoxIsActive(bounds Rectangle) bool {
	cbounds := *bounds.cptr()
	res := C.GuiTextBoxIsActive(cbounds)
	return bool(&res)
}

//GuiTextBoxGetState : Get state for the active textbox
func GuiTextBoxGetState() GuiTextBoxState {
	res := C.GuiTextBoxGetState()
	return newGuiTextBoxStateFromPointer(unsafe.Pointer(&res))
}

//GuiTextBoxSetState : Set state for the active textbox (state must be valid else things will break)
func GuiTextBoxSetState(state GuiTextBoxState) {
	cstate := *state.cptr()
	C.GuiTextBoxSetState(cstate)

}

//GuiTextBoxSelectAll : Select all characters in the active textbox (same as pressing `CTRL` + `A`)
func GuiTextBoxSelectAll(text string) {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	C.GuiTextBoxSelectAll(ctext)

}

//GuiTextBoxCopy : Copy selected text to clipboard from the active textbox (same as pressing `CTRL` + `C`)
func GuiTextBoxCopy(text string) {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	C.GuiTextBoxCopy(ctext)

}

//GuiTextBoxPaste : Paste text from clipboard into the textbox (same as pressing `CTRL` + `V`)
func GuiTextBoxPaste(text string, textSize int) {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	C.GuiTextBoxPaste(ctext, C.int(textSize))

}

//GuiTextBoxCut : Cut selected text in the active textbox and copy it to clipboard (same as pressing `CTRL` + `X`)
func GuiTextBoxCut(text string) {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	C.GuiTextBoxCut(ctext)

}

//GuiTextBoxDelete : Deletes a character or selection before from the active textbox (depending on `before`). Returns bytes deleted.
func GuiTextBoxDelete(text string, length int, before bool) int {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	res := C.GuiTextBoxDelete(ctext, C.int(length), C.bool(before))
	return int(&res)
}

//GuiTextBoxGetByteIndex : Get the byte index for a character starting at position `from` with index `start` until position `to`.
func GuiTextBoxGetByteIndex(text string, start int, from int, to int) int {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	res := C.GuiTextBoxGetByteIndex(ctext, C.int(start), C.int(from), C.int(to))
	return int(&res)
}

//GuiWindowBox : Window Box control, shows a window that can be closed
func GuiWindowBox(bounds Rectangle, title string) bool {
	ctitle := C.CString(title)
	defer C.free(unsafe.Pointer(ctitle))
	cbounds := *bounds.cptr()
	res := C.GuiWindowBox(cbounds, ctitle)
	return bool(&res)
}

//GuiGroupBox : Group Box control with text name
func GuiGroupBox(bounds Rectangle, text string) {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	cbounds := *bounds.cptr()
	C.GuiGroupBox(cbounds, ctext)

}

//GuiLine : Line separator control, could contain text
func GuiLine(bounds Rectangle, text string) {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	cbounds := *bounds.cptr()
	C.GuiLine(cbounds, ctext)

}

//GuiPanel : Panel control, useful to group controls
func GuiPanel(bounds Rectangle) {
	cbounds := *bounds.cptr()
	C.GuiPanel(cbounds)

}

//GuiScrollPanel : Scroll Panel control
func GuiScrollPanel(bounds Rectangle, content Rectangle, scroll Vector2) (Rectangle, Vector2) {
	cscroll := scroll.cptr()
	ccontent := *content.cptr()
	cbounds := *bounds.cptr()
	res := C.GuiScrollPanel(cbounds, ccontent, &cscroll)
	return newRectangleFromPointer(unsafe.Pointer(&res)), newVector2FromPointer(unsafe.Pointer(cscroll))
}

//GuiLabel : Label control, shows text
func GuiLabel(bounds Rectangle, text string) {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	cbounds := *bounds.cptr()
	C.GuiLabel(cbounds, ctext)

}

//GuiButton : Button control, returns true when clicked
func GuiButton(bounds Rectangle, text string) bool {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	cbounds := *bounds.cptr()
	res := C.GuiButton(cbounds, ctext)
	return bool(&res)
}

//GuiLabelButton : Label button control, show true when clicked
func GuiLabelButton(bounds Rectangle, text string) bool {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	cbounds := *bounds.cptr()
	res := C.GuiLabelButton(cbounds, ctext)
	return bool(&res)
}

//GuiImageButton : Image button control, returns true when clicked
func GuiImageButton(bounds Rectangle, text string, texture Texture2D) bool {
	ctexture := *texture.cptr()
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	cbounds := *bounds.cptr()
	res := C.GuiImageButton(cbounds, ctext, ctexture)
	return bool(&res)
}

//GuiImageButtonEx : Image button extended control, returns true when clicked
func GuiImageButtonEx(bounds Rectangle, text string, texture Texture2D, texSource Rectangle) bool {
	ctexSource := *texSource.cptr()
	ctexture := *texture.cptr()
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	cbounds := *bounds.cptr()
	res := C.GuiImageButtonEx(cbounds, ctext, ctexture, ctexSource)
	return bool(&res)
}

//GuiToggle : Toggle Button control, returns true when active
func GuiToggle(bounds Rectangle, text string, active bool) bool {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	cbounds := *bounds.cptr()
	res := C.GuiToggle(cbounds, ctext, C.bool(active))
	return bool(&res)
}

//GuiToggleGroup : Toggle Group control, returns active toggle index
func GuiToggleGroup(bounds Rectangle, text string, active int) int {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	cbounds := *bounds.cptr()
	res := C.GuiToggleGroup(cbounds, ctext, C.int(active))
	return int(&res)
}

//GuiCheckBox : Check Box control, returns true when active
func GuiCheckBox(bounds Rectangle, text string, checked bool) bool {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	cbounds := *bounds.cptr()
	res := C.GuiCheckBox(cbounds, ctext, C.bool(checked))
	return bool(&res)
}

//GuiComboBox : Combo Box control, returns selected item index
func GuiComboBox(bounds Rectangle, text string, active int) int {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	cbounds := *bounds.cptr()
	res := C.GuiComboBox(cbounds, ctext, C.int(active))
	return int(&res)
}

//GuiDropdownBox : Dropdown Box control, returns selected item
func GuiDropdownBox(bounds Rectangle, text string, active int, editMode bool) (bool, int) {
	cactive := C.int(active)
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	cbounds := *bounds.cptr()
	res := C.GuiDropdownBox(cbounds, ctext, &cactive, C.bool(editMode))
	return bool(&res), int(cactive)
}

//GuiSpinner : Spinner control, returns selected value
func GuiSpinner(bounds Rectangle, text string, value int, minValue int, maxValue int, editMode bool) (bool, int) {
	cvalue := C.int(value)
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	cbounds := *bounds.cptr()
	res := C.GuiSpinner(cbounds, ctext, &cvalue, C.int(minValue), C.int(maxValue), C.bool(editMode))
	return bool(&res), int(cvalue)
}

//GuiValueBox : Value Box control, updates input text with numbers
func GuiValueBox(bounds Rectangle, text string, value int, minValue int, maxValue int, editMode bool) (bool, int) {
	cvalue := C.int(value)
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	cbounds := *bounds.cptr()
	res := C.GuiValueBox(cbounds, ctext, &cvalue, C.int(minValue), C.int(maxValue), C.bool(editMode))
	return bool(&res), int(cvalue)
}

//GuiTextBox : Text Box control, updates input text
func GuiTextBox(bounds Rectangle, text string, textSize int, editMode bool) bool {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	cbounds := *bounds.cptr()
	res := C.GuiTextBox(cbounds, ctext, C.int(textSize), C.bool(editMode))
	return bool(&res)
}

//GuiTextBoxMulti : Text Box control with multiple lines
func GuiTextBoxMulti(bounds Rectangle, text string, textSize int, editMode bool) bool {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	cbounds := *bounds.cptr()
	res := C.GuiTextBoxMulti(cbounds, ctext, C.int(textSize), C.bool(editMode))
	return bool(&res)
}

//GuiSlider : Slider control, returns selected value
func GuiSlider(bounds Rectangle, textLeft string, textRight string, value float32, minValue float32, maxValue float32) float32 {
	ctextRight := C.CString(textRight)
	defer C.free(unsafe.Pointer(ctextRight))
	ctextLeft := C.CString(textLeft)
	defer C.free(unsafe.Pointer(ctextLeft))
	cbounds := *bounds.cptr()
	res := C.GuiSlider(cbounds, ctextLeft, ctextRight, C.float(value), C.float(minValue), C.float(maxValue))
	return float32(&res)
}

//GuiSliderBar : Slider Bar control, returns selected value
func GuiSliderBar(bounds Rectangle, textLeft string, textRight string, value float32, minValue float32, maxValue float32) float32 {
	ctextRight := C.CString(textRight)
	defer C.free(unsafe.Pointer(ctextRight))
	ctextLeft := C.CString(textLeft)
	defer C.free(unsafe.Pointer(ctextLeft))
	cbounds := *bounds.cptr()
	res := C.GuiSliderBar(cbounds, ctextLeft, ctextRight, C.float(value), C.float(minValue), C.float(maxValue))
	return float32(&res)
}

//GuiProgressBar : Progress Bar control, shows current progress value
func GuiProgressBar(bounds Rectangle, textLeft string, textRight string, value float32, minValue float32, maxValue float32) float32 {
	ctextRight := C.CString(textRight)
	defer C.free(unsafe.Pointer(ctextRight))
	ctextLeft := C.CString(textLeft)
	defer C.free(unsafe.Pointer(ctextLeft))
	cbounds := *bounds.cptr()
	res := C.GuiProgressBar(cbounds, ctextLeft, ctextRight, C.float(value), C.float(minValue), C.float(maxValue))
	return float32(&res)
}

//GuiStatusBar : Status Bar control, shows info text
func GuiStatusBar(bounds Rectangle, text string) {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	cbounds := *bounds.cptr()
	C.GuiStatusBar(cbounds, ctext)

}

//GuiDummyRec : Dummy control for placeholders
func GuiDummyRec(bounds Rectangle, text string) {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	cbounds := *bounds.cptr()
	C.GuiDummyRec(cbounds, ctext)

}

//GuiScrollBar : Scroll Bar control
func GuiScrollBar(bounds Rectangle, value int, minValue int, maxValue int) int {
	cbounds := *bounds.cptr()
	res := C.GuiScrollBar(cbounds, C.int(value), C.int(minValue), C.int(maxValue))
	return int(&res)
}

//GuiGrid : Grid control
func GuiGrid(bounds Rectangle, spacing float32, subdivs int) Vector2 {
	cbounds := *bounds.cptr()
	res := C.GuiGrid(cbounds, C.float(spacing), C.int(subdivs))
	return newVector2FromPointer(unsafe.Pointer(&res))
}

//GuiListView : List View control, returns selected list item index
func GuiListView(bounds Rectangle, text string, scrollIndex int, active int) (int, int) {
	cscrollIndex := C.int(scrollIndex)
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	cbounds := *bounds.cptr()
	res := C.GuiListView(cbounds, ctext, &cscrollIndex, C.int(active))
	return int(&res), int(cscrollIndex)
}

//GuiListViewEx :  List View with extended parameters
func GuiListViewEx(bounds Rectangle, text string, count int, focus int, scrollIndex int, active int) (int, string) {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	cbounds := *bounds.cptr()
	res := C.GuiListViewEx(cbounds, &ctext, C.int(count), C.int(focus), C.int(scrollIndex), C.int(active))
	return int(res), C.GoString(ctext)
}

//GuiMessageBox : Message Box control, displays a message
func GuiMessageBox(bounds Rectangle, title string, message string, buttons string) int {
	cbuttons := C.CString(buttons)
	defer C.free(unsafe.Pointer(cbuttons))
	cmessage := C.CString(message)
	defer C.free(unsafe.Pointer(cmessage))
	ctitle := C.CString(title)
	defer C.free(unsafe.Pointer(ctitle))
	cbounds := *bounds.cptr()
	res := C.GuiMessageBox(cbounds, ctitle, cmessage, cbuttons)
	return int(&res)
}

//GuiTextInputBox : Text Input Box control, ask for text
func GuiTextInputBox(bounds Rectangle, title string, message string, buttons string, text string) int {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	cbuttons := C.CString(buttons)
	defer C.free(unsafe.Pointer(cbuttons))
	cmessage := C.CString(message)
	defer C.free(unsafe.Pointer(cmessage))
	ctitle := C.CString(title)
	defer C.free(unsafe.Pointer(ctitle))
	cbounds := *bounds.cptr()
	res := C.GuiTextInputBox(cbounds, ctitle, cmessage, cbuttons, ctext)
	return int(&res)
}

//GuiColorPicker : Color Picker control
func GuiColorPicker(bounds Rectangle, color Color) Color {
	ccolor := *color.cptr()
	cbounds := *bounds.cptr()
	res := C.GuiColorPicker(cbounds, ccolor)
	return newColorFromPointer(unsafe.Pointer(&res))
}

//GuiLoadStyle : Load style file (.rgs)
func GuiLoadStyle(fileName string) {
	cfileName := C.CString(fileName)
	defer C.free(unsafe.Pointer(cfileName))
	C.GuiLoadStyle(cfileName)

}

//GuiLoadStyleDefault : Load style default over global style
func GuiLoadStyleDefault() {
	C.GuiLoadStyleDefault()

}

//GuiIconText : Get text with icon id prepended
func GuiIconText(iconId int, text string) string {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	res := C.GuiIconText(C.int(iconId), ctext)
	return C.GoString(&res)
}
