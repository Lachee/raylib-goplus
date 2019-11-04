package raylib
//InitWindow :  Initialize window and OpenGL context
func InitWindow(width int, height int, title string) () {
ctitle := C.CString(title)
defer C.free(unsafe.Pointer(&ctitle))
C.InitWindow(C.int(width), C.int(height), ctitle)
}
//WindowShouldClose :  Check if KEY_ESCAPE pressed or Close icon pressed
func WindowShouldClose() () {
C.WindowShouldClose()
}
//CloseWindow :  Close window and unload OpenGL context
func CloseWindow() () {
C.CloseWindow()
}
//IsWindowReady :  Check if window has been initialized successfully
func IsWindowReady() () {
C.IsWindowReady()
}
//IsWindowMinimized :  Check if window has been minimized (or lost focus)
func IsWindowMinimized() () {
C.IsWindowMinimized()
}
//IsWindowResized :  Check if window has been resized
func IsWindowResized() () {
C.IsWindowResized()
}
//IsWindowHidden :  Check if window is currently hidden
func IsWindowHidden() () {
C.IsWindowHidden()
}
//ToggleFullscreen :  Toggle fullscreen mode (only PLATFORM_DESKTOP)
func ToggleFullscreen() () {
C.ToggleFullscreen()
}
//UnhideWindow :  Show the window
func UnhideWindow() () {
C.UnhideWindow()
}
//HideWindow :  Hide the window
func HideWindow() () {
C.HideWindow()
}
//SetWindowIcon :  Set icon for window (only PLATFORM_DESKTOP)
func SetWindowIcon(image Image) () {
cimage := *image.cptr()
C.SetWindowIcon(cimage)
}
//SetWindowTitle :  Set title for window (only PLATFORM_DESKTOP)
func SetWindowTitle(title string) () {
ctitle := C.CString(title)
defer C.free(unsafe.Pointer(&ctitle))
C.SetWindowTitle(ctitle)
}
//SetWindowPosition :  Set window position on screen (only PLATFORM_DESKTOP)
func SetWindowPosition(x int, y int) () {
C.SetWindowPosition(C.int(x), C.int(y))
}
//SetWindowMonitor :  Set monitor for the current window (fullscreen mode)
func SetWindowMonitor(monitor int) () {
C.SetWindowMonitor(C.int(monitor))
}
//SetWindowMinSize :  Set window minimum dimensions (for FLAG_WINDOW_RESIZABLE)
func SetWindowMinSize(width int, height int) () {
C.SetWindowMinSize(C.int(width), C.int(height))
}
//SetWindowSize :  Set window dimensions
func SetWindowSize(width int, height int) () {
C.SetWindowSize(C.int(width), C.int(height))
}
//GetScreenWidth :  Get current screen width
func GetScreenWidth() () {
C.GetScreenWidth()
}
//GetScreenHeight :  Get current screen height
func GetScreenHeight() () {
C.GetScreenHeight()
}
//GetMonitorCount :  Get number of connected monitors
func GetMonitorCount() () {
C.GetMonitorCount()
}
//GetMonitorWidth :  Get primary monitor width
func GetMonitorWidth(monitor int) () {
C.GetMonitorWidth(C.int(monitor))
}
//GetMonitorHeight :  Get primary monitor height
func GetMonitorHeight(monitor int) () {
C.GetMonitorHeight(C.int(monitor))
}
//GetMonitorPhysicalWidth :  Get primary monitor physical width in millimetres
func GetMonitorPhysicalWidth(monitor int) () {
C.GetMonitorPhysicalWidth(C.int(monitor))
}
//GetMonitorPhysicalHeight :  Get primary monitor physical height in millimetres
func GetMonitorPhysicalHeight(monitor int) () {
C.GetMonitorPhysicalHeight(C.int(monitor))
}
//GetWindowPosition :  Get window position XY on monitor
func GetWindowPosition() () {
C.GetWindowPosition()
}
//GetMonitorName :  Get the human-readable, UTF-8 encoded name of the primary monitor
func GetMonitorName(monitor int) () {
C.GetMonitorName(C.int(monitor))
}
//GetClipboardText :  Get clipboard text content
func GetClipboardText() () {
C.GetClipboardText()
}
//SetClipboardText :  Set clipboard text content
func SetClipboardText(text string) () {
ctext := C.CString(text)
defer C.free(unsafe.Pointer(&ctext))
C.SetClipboardText(ctext)
}
//ShowCursor :  Shows cursor
func ShowCursor() () {
C.ShowCursor()
}
//HideCursor :  Hides cursor
func HideCursor() () {
C.HideCursor()
}
//IsCursorHidden :  Check if cursor is not visible
func IsCursorHidden() () {
C.IsCursorHidden()
}
//EnableCursor :  Enables cursor (unlock cursor)
func EnableCursor() () {
C.EnableCursor()
}
//DisableCursor :  Disables cursor (lock cursor)
func DisableCursor() () {
C.DisableCursor()
}
//ClearBackground :  Set background color (framebuffer clear color)
func ClearBackground(color Color) () {
ccolor := *color.cptr()
C.ClearBackground(ccolor)
}
//BeginDrawing :  Setup canvas (framebuffer) to start drawing
func BeginDrawing() () {
C.BeginDrawing()
}
//EndDrawing :  End canvas drawing and swap buffers (double buffering)
func EndDrawing() () {
C.EndDrawing()
}
//BeginMode2D :  Initialize 2D mode with custom camera (2D)
func BeginMode2D(camera Camera2D) () {
ccamera := *camera.cptr()
C.BeginMode2D(ccamera)
}
//EndMode2D :  Ends 2D mode with custom camera
func EndMode2D() () {
C.EndMode2D()
}
//BeginMode3D :  Initializes 3D mode with custom camera (3D)
func BeginMode3D(camera Camera3D) () {
ccamera := *camera.cptr()
C.BeginMode3D(ccamera)
}
//EndMode3D :  Ends 3D mode and returns to default 2D orthographic mode
func EndMode3D() () {
C.EndMode3D()
}
//BeginTextureMode :  Initializes render texture for drawing
func BeginTextureMode(target RenderTexture2D) () {
ctarget := *target.cptr()
C.BeginTextureMode(ctarget)
}
//EndTextureMode :  Ends drawing to render texture
func EndTextureMode() () {
C.EndTextureMode()
}
//BeginScissorMode :  Begin scissor mode (define screen area for following drawing)
func BeginScissorMode(x int, y int, width int, height int) () {
C.BeginScissorMode(C.int(x), C.int(y), C.int(width), C.int(height))
}
//EndScissorMode :  End scissor mode
func EndScissorMode() () {
C.EndScissorMode()
}
//GetMouseRay :  Returns a ray trace from mouse position
func GetMouseRay(mousePosition Vector2, camera Camera) () {
ccamera := *camera.cptr()
cmousePosition := *mousePosition.cptr()
C.GetMouseRay(cmousePosition, ccamera)
}
//GetCameraMatrix :  Returns camera transform matrix (view matrix)
func GetCameraMatrix(camera Camera) () {
ccamera := *camera.cptr()
C.GetCameraMatrix(ccamera)
}
//GetCameraMatrix2D :  Returns camera 2d transform matrix
func GetCameraMatrix2D(camera Camera2D) () {
ccamera := *camera.cptr()
C.GetCameraMatrix2D(ccamera)
}
//GetWorldToScreen :  Returns the screen space position for a 3d world space position
func GetWorldToScreen(position Vector3, camera Camera) () {
ccamera := *camera.cptr()
cposition := *position.cptr()
C.GetWorldToScreen(cposition, ccamera)
}
//GetWorldToScreen2D :  Returns the screen space position for a 2d camera world space position
func GetWorldToScreen2D(position Vector2, camera Camera2D) () {
ccamera := *camera.cptr()
cposition := *position.cptr()
C.GetWorldToScreen2D(cposition, ccamera)
}
//GetScreenToWorld2D :  Returns the world space position for a 2d camera screen space position
func GetScreenToWorld2D(position Vector2, camera Camera2D) () {
ccamera := *camera.cptr()
cposition := *position.cptr()
C.GetScreenToWorld2D(cposition, ccamera)
}
//SetTargetFPS :  Set target FPS (maximum)
func SetTargetFPS(fps int) () {
C.SetTargetFPS(C.int(fps))
}
//GetFPS :  Returns current FPS
func GetFPS() () {
C.GetFPS()
}
//GetFrameTime :  Returns time in seconds for last frame drawn
func GetFrameTime() () {
C.GetFrameTime()
}
//GetTime :  Returns elapsed time in seconds since InitWindow()
func GetTime() () {
C.GetTime()
}
//ColorToInt :  Returns hexadecimal value for a Color
func ColorToInt(color Color) () {
ccolor := *color.cptr()
C.ColorToInt(ccolor)
}
//ColorNormalize :  Returns color normalized as float [0..1]
func ColorNormalize(color Color) () {
ccolor := *color.cptr()
C.ColorNormalize(ccolor)
}
//ColorToHSV :  Returns HSV values for a Color
func ColorToHSV(color Color) () {
ccolor := *color.cptr()
C.ColorToHSV(ccolor)
}
//ColorFromHSV :  Returns a Color from HSV values
func ColorFromHSV(hsv Vector3) () {
chsv := *hsv.cptr()
C.ColorFromHSV(chsv)
}
//GetColor :  Returns a Color struct from hexadecimal value
func GetColor(hexValue int) () {
C.GetColor(C.int(hexValue))
}
//Fade :  Color fade-in or fade-out, alpha goes from 0.0f to 1.0f
func Fade(color Color, alpha float32) () {
ccolor := *color.cptr()
C.Fade(ccolor, C.float(alpha))
}
//SetConfigFlags :  Setup window configuration flags (view FLAGS)
func SetConfigFlags(flags int) () {
C.SetConfigFlags(C.int(flags))
}
//SetTraceLogLevel :  Set the current threshold (minimum) log level
func SetTraceLogLevel(logType int) () {
C.SetTraceLogLevel(C.int(logType))
}
//SetTraceLogExit :  Set the exit threshold (minimum) log level
func SetTraceLogExit(logType int) () {
C.SetTraceLogExit(C.int(logType))
}
//SetTraceLogCallback :  Set a trace log callback to enable custom logging
func SetTraceLogCallback(callback TraceLogCallback) () {
ccallback := *callback.cptr()
C.SetTraceLogCallback(ccallback)
}
//TakeScreenshot :  Takes a screenshot of current screen (saved a .png)
func TakeScreenshot(fileName string) () {
cfileName := C.CString(fileName)
defer C.free(unsafe.Pointer(&cfileName))
C.TakeScreenshot(cfileName)
}
//GetRandomValue :  Returns a random value between min and max (both included)
func GetRandomValue(min int, max int) () {
C.GetRandomValue(C.int(min), C.int(max))
}
//FileExists :  Check if file exists
func FileExists(fileName string) () {
cfileName := C.CString(fileName)
defer C.free(unsafe.Pointer(&cfileName))
C.FileExists(cfileName)
}
//IsFileExtension :  Check file extension
func IsFileExtension(fileName string, ext string) () {
cext := C.CString(ext)
defer C.free(unsafe.Pointer(&cext))
cfileName := C.CString(fileName)
defer C.free(unsafe.Pointer(&cfileName))
C.IsFileExtension(cfileName, cext)
}
//DirectoryExists :  Check if a directory path exists
func DirectoryExists(dirPath string) () {
cdirPath := C.CString(dirPath)
defer C.free(unsafe.Pointer(&cdirPath))
C.DirectoryExists(cdirPath)
}
//GetExtension :  Get pointer to extension for a filename string
func GetExtension(fileName string) () {
cfileName := C.CString(fileName)
defer C.free(unsafe.Pointer(&cfileName))
C.GetExtension(cfileName)
}
//GetFileName :  Get pointer to filename for a path string
func GetFileName(filePath string) () {
cfilePath := C.CString(filePath)
defer C.free(unsafe.Pointer(&cfilePath))
C.GetFileName(cfilePath)
}
//GetFileNameWithoutExt :  Get filename string without extension (uses static string)
func GetFileNameWithoutExt(filePath string) () {
cfilePath := C.CString(filePath)
defer C.free(unsafe.Pointer(&cfilePath))
C.GetFileNameWithoutExt(cfilePath)
}
//GetDirectoryPath :  Get full path for a given fileName with path (uses static string)
func GetDirectoryPath(filePath string) () {
cfilePath := C.CString(filePath)
defer C.free(unsafe.Pointer(&cfilePath))
C.GetDirectoryPath(cfilePath)
}
//GetPrevDirectoryPath :  Get previous directory path for a given path (uses static string)
func GetPrevDirectoryPath(dirPath string) () {
cdirPath := C.CString(dirPath)
defer C.free(unsafe.Pointer(&cdirPath))
C.GetPrevDirectoryPath(cdirPath)
}
//GetWorkingDirectory :  Get current working directory (uses static string)
func GetWorkingDirectory() () {
C.GetWorkingDirectory()
}
//ClearDirectoryFiles :  Clear directory files paths buffers (free memory)
func ClearDirectoryFiles() () {
C.ClearDirectoryFiles()
}
//ChangeDirectory :  Change working directory, returns true if success
func ChangeDirectory(dir string) () {
cdir := C.CString(dir)
defer C.free(unsafe.Pointer(&cdir))
C.ChangeDirectory(cdir)
}
//IsFileDropped :  Check if a file has been dropped into window
func IsFileDropped() () {
C.IsFileDropped()
}
//ClearDroppedFiles :  Clear dropped files paths buffer (free memory)
func ClearDroppedFiles() () {
C.ClearDroppedFiles()
}
//GetFileModTime :  Get file modification time (last write time)
func GetFileModTime(fileName string) () {
cfileName := C.CString(fileName)
defer C.free(unsafe.Pointer(&cfileName))
C.GetFileModTime(cfileName)
}
//StorageSaveValue :  Save integer value to storage file (to defined position)
func StorageSaveValue(position int, value int) () {
C.StorageSaveValue(C.int(position), C.int(value))
}
//StorageLoadValue :  Load integer value from storage file (from defined position)
func StorageLoadValue(position int) () {
C.StorageLoadValue(C.int(position))
}
//OpenURL :  Open URL with default system browser (if available)
func OpenURL(url string) () {
curl := C.CString(url)
defer C.free(unsafe.Pointer(&curl))
C.OpenURL(curl)
}
//IsKeyPressed :  Detect if a key has been pressed once
func IsKeyPressed(key int) () {
C.IsKeyPressed(C.int(key))
}
//IsKeyDown :  Detect if a key is being pressed
func IsKeyDown(key int) () {
C.IsKeyDown(C.int(key))
}
//IsKeyReleased :  Detect if a key has been released once
func IsKeyReleased(key int) () {
C.IsKeyReleased(C.int(key))
}
//IsKeyUp :  Detect if a key is NOT being pressed
func IsKeyUp(key int) () {
C.IsKeyUp(C.int(key))
}
//GetKeyPressed :  Get latest key pressed
func GetKeyPressed() () {
C.GetKeyPressed()
}
//SetExitKey :  Set a custom key to exit program (default is ESC)
func SetExitKey(key int) () {
C.SetExitKey(C.int(key))
}
//IsGamepadAvailable :  Detect if a gamepad is available
func IsGamepadAvailable(gamepad int) () {
C.IsGamepadAvailable(C.int(gamepad))
}
//IsGamepadName :  Check gamepad name (if available)
func IsGamepadName(gamepad int, name string) () {
cname := C.CString(name)
defer C.free(unsafe.Pointer(&cname))
C.IsGamepadName(C.int(gamepad), cname)
}
//GetGamepadName :  Return gamepad internal name id
func GetGamepadName(gamepad int) () {
C.GetGamepadName(C.int(gamepad))
}
//IsGamepadButtonPressed :  Detect if a gamepad button has been pressed once
func IsGamepadButtonPressed(gamepad int, button int) () {
C.IsGamepadButtonPressed(C.int(gamepad), C.int(button))
}
//IsGamepadButtonDown :  Detect if a gamepad button is being pressed
func IsGamepadButtonDown(gamepad int, button int) () {
C.IsGamepadButtonDown(C.int(gamepad), C.int(button))
}
//IsGamepadButtonReleased :  Detect if a gamepad button has been released once
func IsGamepadButtonReleased(gamepad int, button int) () {
C.IsGamepadButtonReleased(C.int(gamepad), C.int(button))
}
//IsGamepadButtonUp :  Detect if a gamepad button is NOT being pressed
func IsGamepadButtonUp(gamepad int, button int) () {
C.IsGamepadButtonUp(C.int(gamepad), C.int(button))
}
//GetGamepadButtonPressed :  Get the last gamepad button pressed
func GetGamepadButtonPressed() () {
C.GetGamepadButtonPressed()
}
//GetGamepadAxisCount :  Return gamepad axis count for a gamepad
func GetGamepadAxisCount(gamepad int) () {
C.GetGamepadAxisCount(C.int(gamepad))
}
//GetGamepadAxisMovement :  Return axis movement value for a gamepad axis
func GetGamepadAxisMovement(gamepad int, axis int) () {
C.GetGamepadAxisMovement(C.int(gamepad), C.int(axis))
}
//IsMouseButtonPressed :  Detect if a mouse button has been pressed once
func IsMouseButtonPressed(button int) () {
C.IsMouseButtonPressed(C.int(button))
}
//IsMouseButtonDown :  Detect if a mouse button is being pressed
func IsMouseButtonDown(button int) () {
C.IsMouseButtonDown(C.int(button))
}
//IsMouseButtonReleased :  Detect if a mouse button has been released once
func IsMouseButtonReleased(button int) () {
C.IsMouseButtonReleased(C.int(button))
}
//IsMouseButtonUp :  Detect if a mouse button is NOT being pressed
func IsMouseButtonUp(button int) () {
C.IsMouseButtonUp(C.int(button))
}
//GetMouseX :  Returns mouse position X
func GetMouseX() () {
C.GetMouseX()
}
//GetMouseY :  Returns mouse position Y
func GetMouseY() () {
C.GetMouseY()
}
//GetMousePosition :  Returns mouse position XY
func GetMousePosition() () {
C.GetMousePosition()
}
//SetMousePosition :  Set mouse position XY
func SetMousePosition(x int, y int) () {
C.SetMousePosition(C.int(x), C.int(y))
}
//SetMouseOffset :  Set mouse offset
func SetMouseOffset(offsetX int, offsetY int) () {
C.SetMouseOffset(C.int(offsetX), C.int(offsetY))
}
//SetMouseScale :  Set mouse scaling
func SetMouseScale(scaleX float32, scaleY float32) () {
C.SetMouseScale(C.float(scaleX), C.float(scaleY))
}
//GetMouseWheelMove :  Returns mouse wheel movement Y
func GetMouseWheelMove() () {
C.GetMouseWheelMove()
}
//GetTouchX :  Returns touch position X for touch point 0 (relative to screen size)
func GetTouchX() () {
C.GetTouchX()
}
//GetTouchY :  Returns touch position Y for touch point 0 (relative to screen size)
func GetTouchY() () {
C.GetTouchY()
}
//GetTouchPosition :  Returns touch position XY for a touch point index (relative to screen size)
func GetTouchPosition(index int) () {
C.GetTouchPosition(C.int(index))
}
//SetGesturesEnabled :  Enable a set of gestures using flags
func SetGesturesEnabled(gestureFlags int) () {
C.SetGesturesEnabled(C.int(gestureFlags))
}
//IsGestureDetected :  Check if a gesture have been detected
func IsGestureDetected(gesture int) () {
C.IsGestureDetected(C.int(gesture))
}
//GetGestureDetected :  Get latest detected gesture
func GetGestureDetected() () {
C.GetGestureDetected()
}
//GetTouchPointsCount :  Get touch points count
func GetTouchPointsCount() () {
C.GetTouchPointsCount()
}
//GetGestureHoldDuration :  Get gesture hold time in milliseconds
func GetGestureHoldDuration() () {
C.GetGestureHoldDuration()
}
//GetGestureDragVector :  Get gesture drag vector
func GetGestureDragVector() () {
C.GetGestureDragVector()
}
//GetGestureDragAngle :  Get gesture drag angle
func GetGestureDragAngle() () {
C.GetGestureDragAngle()
}
//GetGesturePinchVector :  Get gesture pinch delta
func GetGesturePinchVector() () {
C.GetGesturePinchVector()
}
//GetGesturePinchAngle :  Get gesture pinch angle
func GetGesturePinchAngle() () {
C.GetGesturePinchAngle()
}
//SetCameraMode :  Set camera mode (multiple camera modes available)
func SetCameraMode(camera Camera, mode int) () {
ccamera := *camera.cptr()
C.SetCameraMode(ccamera, C.int(mode))
}
//SetCameraPanControl :  Set camera pan key to combine with mouse movement (free camera)
func SetCameraPanControl(panKey int) () {
C.SetCameraPanControl(C.int(panKey))
}
//SetCameraAltControl :  Set camera alt key to combine with mouse movement (free camera)
func SetCameraAltControl(altKey int) () {
C.SetCameraAltControl(C.int(altKey))
}
//SetCameraSmoothZoomControl :  Set camera smooth zoom key to combine with mouse (free camera)
func SetCameraSmoothZoomControl(szKey int) () {
C.SetCameraSmoothZoomControl(C.int(szKey))
}
//SetCameraMoveControls :  Set camera move controls (1st person and 3rd person cameras)
func SetCameraMoveControls(frontKey int, backKey int, rightKey int, leftKey int, upKey int, downKey int) () {
C.SetCameraMoveControls(C.int(frontKey), C.int(backKey), C.int(rightKey), C.int(leftKey), C.int(upKey), C.int(downKey))
}
//DrawPixel :  Draw a pixel
func DrawPixel(posX int, posY int, color Color) () {
ccolor := *color.cptr()
C.DrawPixel(C.int(posX), C.int(posY), ccolor)
}
//DrawPixelV :  Draw a pixel (Vector version)
func DrawPixelV(position Vector2, color Color) () {
ccolor := *color.cptr()
cposition := *position.cptr()
C.DrawPixelV(cposition, ccolor)
}
//DrawLine :  Draw a line
func DrawLine(startPosX int, startPosY int, endPosX int, endPosY int, color Color) () {
ccolor := *color.cptr()
C.DrawLine(C.int(startPosX), C.int(startPosY), C.int(endPosX), C.int(endPosY), ccolor)
}
//DrawLineV :  Draw a line (Vector version)
func DrawLineV(startPos Vector2, endPos Vector2, color Color) () {
ccolor := *color.cptr()
cendPos := *endPos.cptr()
cstartPos := *startPos.cptr()
C.DrawLineV(cstartPos, cendPos, ccolor)
}
//DrawLineEx :  Draw a line defining thickness
func DrawLineEx(startPos Vector2, endPos Vector2, thick float32, color Color) () {
ccolor := *color.cptr()
cendPos := *endPos.cptr()
cstartPos := *startPos.cptr()
C.DrawLineEx(cstartPos, cendPos, C.float(thick), ccolor)
}
//DrawLineBezier :  Draw a line using cubic-bezier curves in-out
func DrawLineBezier(startPos Vector2, endPos Vector2, thick float32, color Color) () {
ccolor := *color.cptr()
cendPos := *endPos.cptr()
cstartPos := *startPos.cptr()
C.DrawLineBezier(cstartPos, cendPos, C.float(thick), ccolor)
}
//DrawCircle :  Draw a color-filled circle
func DrawCircle(centerX int, centerY int, radius float32, color Color) () {
ccolor := *color.cptr()
C.DrawCircle(C.int(centerX), C.int(centerY), C.float(radius), ccolor)
}
//DrawCircleSector :  Draw a piece of a circle
func DrawCircleSector(center Vector2, radius float32, startAngle int, endAngle int, segments int, color Color) () {
ccolor := *color.cptr()
ccenter := *center.cptr()
C.DrawCircleSector(ccenter, C.float(radius), C.int(startAngle), C.int(endAngle), C.int(segments), ccolor)
}
//DrawCircleSectorLines :  Draw circle sector outline
func DrawCircleSectorLines(center Vector2, radius float32, startAngle int, endAngle int, segments int, color Color) () {
ccolor := *color.cptr()
ccenter := *center.cptr()
C.DrawCircleSectorLines(ccenter, C.float(radius), C.int(startAngle), C.int(endAngle), C.int(segments), ccolor)
}
//DrawCircleGradient :  Draw a gradient-filled circle
func DrawCircleGradient(centerX int, centerY int, radius float32, color1 Color, color2 Color) () {
ccolor2 := *color2.cptr()
ccolor1 := *color1.cptr()
C.DrawCircleGradient(C.int(centerX), C.int(centerY), C.float(radius), ccolor1, ccolor2)
}
//DrawCircleV :  Draw a color-filled circle (Vector version)
func DrawCircleV(center Vector2, radius float32, color Color) () {
ccolor := *color.cptr()
ccenter := *center.cptr()
C.DrawCircleV(ccenter, C.float(radius), ccolor)
}
//DrawCircleLines :  Draw circle outline
func DrawCircleLines(centerX int, centerY int, radius float32, color Color) () {
ccolor := *color.cptr()
C.DrawCircleLines(C.int(centerX), C.int(centerY), C.float(radius), ccolor)
}
//DrawRing :  Draw ring
func DrawRing(center Vector2, innerRadius float32, outerRadius float32, startAngle int, endAngle int, segments int, color Color) () {
ccolor := *color.cptr()
ccenter := *center.cptr()
C.DrawRing(ccenter, C.float(innerRadius), C.float(outerRadius), C.int(startAngle), C.int(endAngle), C.int(segments), ccolor)
}
//DrawRingLines :  Draw ring outline
func DrawRingLines(center Vector2, innerRadius float32, outerRadius float32, startAngle int, endAngle int, segments int, color Color) () {
ccolor := *color.cptr()
ccenter := *center.cptr()
C.DrawRingLines(ccenter, C.float(innerRadius), C.float(outerRadius), C.int(startAngle), C.int(endAngle), C.int(segments), ccolor)
}
//DrawRectangle :  Draw a color-filled rectangle
func DrawRectangle(posX int, posY int, width int, height int, color Color) () {
ccolor := *color.cptr()
C.DrawRectangle(C.int(posX), C.int(posY), C.int(width), C.int(height), ccolor)
}
//DrawRectangleV :  Draw a color-filled rectangle (Vector version)
func DrawRectangleV(position Vector2, size Vector2, color Color) () {
ccolor := *color.cptr()
csize := *size.cptr()
cposition := *position.cptr()
C.DrawRectangleV(cposition, csize, ccolor)
}
//DrawRectangleRec :  Draw a color-filled rectangle
func DrawRectangleRec(rec Rectangle, color Color) () {
ccolor := *color.cptr()
crec := *rec.cptr()
C.DrawRectangleRec(crec, ccolor)
}
//DrawRectanglePro :  Draw a color-filled rectangle with pro parameters
func DrawRectanglePro(rec Rectangle, origin Vector2, rotation float32, color Color) () {
ccolor := *color.cptr()
corigin := *origin.cptr()
crec := *rec.cptr()
C.DrawRectanglePro(crec, corigin, C.float(rotation), ccolor)
}
//DrawRectangleGradientV :  Draw a vertical-gradient-filled rectangle
func DrawRectangleGradientV(posX int, posY int, width int, height int, color1 Color, color2 Color) () {
ccolor2 := *color2.cptr()
ccolor1 := *color1.cptr()
C.DrawRectangleGradientV(C.int(posX), C.int(posY), C.int(width), C.int(height), ccolor1, ccolor2)
}
//DrawRectangleGradientH :  Draw a horizontal-gradient-filled rectangle
func DrawRectangleGradientH(posX int, posY int, width int, height int, color1 Color, color2 Color) () {
ccolor2 := *color2.cptr()
ccolor1 := *color1.cptr()
C.DrawRectangleGradientH(C.int(posX), C.int(posY), C.int(width), C.int(height), ccolor1, ccolor2)
}
//DrawRectangleGradientEx :  Draw a gradient-filled rectangle with custom vertex colors
func DrawRectangleGradientEx(rec Rectangle, col1 Color, col2 Color, col3 Color, col4 Color) () {
ccol4 := *col4.cptr()
ccol3 := *col3.cptr()
ccol2 := *col2.cptr()
ccol1 := *col1.cptr()
crec := *rec.cptr()
C.DrawRectangleGradientEx(crec, ccol1, ccol2, ccol3, ccol4)
}
//DrawRectangleLines :  Draw rectangle outline
func DrawRectangleLines(posX int, posY int, width int, height int, color Color) () {
ccolor := *color.cptr()
C.DrawRectangleLines(C.int(posX), C.int(posY), C.int(width), C.int(height), ccolor)
}
//DrawRectangleLinesEx :  Draw rectangle outline with extended parameters
func DrawRectangleLinesEx(rec Rectangle, lineThick int, color Color) () {
ccolor := *color.cptr()
crec := *rec.cptr()
C.DrawRectangleLinesEx(crec, C.int(lineThick), ccolor)
}
//DrawRectangleRounded :  Draw rectangle with rounded edges
func DrawRectangleRounded(rec Rectangle, roundness float32, segments int, color Color) () {
ccolor := *color.cptr()
crec := *rec.cptr()
C.DrawRectangleRounded(crec, C.float(roundness), C.int(segments), ccolor)
}
//DrawRectangleRoundedLines :  Draw rectangle with rounded edges outline
func DrawRectangleRoundedLines(rec Rectangle, roundness float32, segments int, lineThick int, color Color) () {
ccolor := *color.cptr()
crec := *rec.cptr()
C.DrawRectangleRoundedLines(crec, C.float(roundness), C.int(segments), C.int(lineThick), ccolor)
}
//DrawTriangle :  Draw a color-filled triangle (vertex in counter-clockwise order!)
func DrawTriangle(v1 Vector2, v2 Vector2, v3 Vector2, color Color) () {
ccolor := *color.cptr()
cv3 := *v3.cptr()
cv2 := *v2.cptr()
cv1 := *v1.cptr()
C.DrawTriangle(cv1, cv2, cv3, ccolor)
}
//DrawTriangleLines :  Draw triangle outline (vertex in counter-clockwise order!)
func DrawTriangleLines(v1 Vector2, v2 Vector2, v3 Vector2, color Color) () {
ccolor := *color.cptr()
cv3 := *v3.cptr()
cv2 := *v2.cptr()
cv1 := *v1.cptr()
C.DrawTriangleLines(cv1, cv2, cv3, ccolor)
}
//DrawPoly :  Draw a regular polygon (Vector version)
func DrawPoly(center Vector2, sides int, radius float32, rotation float32, color Color) () {
ccolor := *color.cptr()
ccenter := *center.cptr()
C.DrawPoly(ccenter, C.int(sides), C.float(radius), C.float(rotation), ccolor)
}
//SetShapesTexture :  Define default texture used to draw shapes
func SetShapesTexture(texture Texture2D, source Rectangle) () {
csource := *source.cptr()
ctexture := *texture.cptr()
C.SetShapesTexture(ctexture, csource)
}
//CheckCollisionRecs :  Check collision between two rectangles
func CheckCollisionRecs(rec1 Rectangle, rec2 Rectangle) () {
crec2 := *rec2.cptr()
crec1 := *rec1.cptr()
C.CheckCollisionRecs(crec1, crec2)
}
//CheckCollisionCircles :  Check collision between two circles
func CheckCollisionCircles(center1 Vector2, radius1 float32, center2 Vector2, radius2 float32) () {
ccenter2 := *center2.cptr()
ccenter1 := *center1.cptr()
C.CheckCollisionCircles(ccenter1, C.float(radius1), ccenter2, C.float(radius2))
}
//CheckCollisionCircleRec :  Check collision between circle and rectangle
func CheckCollisionCircleRec(center Vector2, radius float32, rec Rectangle) () {
crec := *rec.cptr()
ccenter := *center.cptr()
C.CheckCollisionCircleRec(ccenter, C.float(radius), crec)
}
//GetCollisionRec :  Get collision rectangle for two rectangles collision
func GetCollisionRec(rec1 Rectangle, rec2 Rectangle) () {
crec2 := *rec2.cptr()
crec1 := *rec1.cptr()
C.GetCollisionRec(crec1, crec2)
}
//CheckCollisionPointRec :  Check if point is inside rectangle
func CheckCollisionPointRec(point Vector2, rec Rectangle) () {
crec := *rec.cptr()
cpoint := *point.cptr()
C.CheckCollisionPointRec(cpoint, crec)
}
//CheckCollisionPointCircle :  Check if point is inside circle
func CheckCollisionPointCircle(point Vector2, center Vector2, radius float32) () {
ccenter := *center.cptr()
cpoint := *point.cptr()
C.CheckCollisionPointCircle(cpoint, ccenter, C.float(radius))
}
//CheckCollisionPointTriangle :  Check if point is inside a triangle
func CheckCollisionPointTriangle(point Vector2, p1 Vector2, p2 Vector2, p3 Vector2) () {
cp3 := *p3.cptr()
cp2 := *p2.cptr()
cp1 := *p1.cptr()
cpoint := *point.cptr()
C.CheckCollisionPointTriangle(cpoint, cp1, cp2, cp3)
}
//LoadImage :  Load image from file into CPU memory (RAM)
func LoadImage(fileName string) () {
cfileName := C.CString(fileName)
defer C.free(unsafe.Pointer(&cfileName))
C.LoadImage(cfileName)
}
//LoadImageRaw :  Load image from RAW file data
func LoadImageRaw(fileName string, width int, height int, format int, headerSize int) () {
cfileName := C.CString(fileName)
defer C.free(unsafe.Pointer(&cfileName))
C.LoadImageRaw(cfileName, C.int(width), C.int(height), C.int(format), C.int(headerSize))
}
//ExportImage :  Export image data to file
func ExportImage(image Image, fileName string) () {
cfileName := C.CString(fileName)
defer C.free(unsafe.Pointer(&cfileName))
cimage := *image.cptr()
C.ExportImage(cimage, cfileName)
}
//ExportImageAsCode :  Export image as code file defining an array of bytes
func ExportImageAsCode(image Image, fileName string) () {
cfileName := C.CString(fileName)
defer C.free(unsafe.Pointer(&cfileName))
cimage := *image.cptr()
C.ExportImageAsCode(cimage, cfileName)
}
//LoadTexture :  Load texture from file into GPU memory (VRAM)
func LoadTexture(fileName string) () {
cfileName := C.CString(fileName)
defer C.free(unsafe.Pointer(&cfileName))
C.LoadTexture(cfileName)
}
//LoadTextureFromImage :  Load texture from image data
func LoadTextureFromImage(image Image) () {
cimage := *image.cptr()
C.LoadTextureFromImage(cimage)
}
//LoadTextureCubemap :  Load cubemap from image, multiple image cubemap layouts supported
func LoadTextureCubemap(image Image, layoutType int) () {
cimage := *image.cptr()
C.LoadTextureCubemap(cimage, C.int(layoutType))
}
//LoadRenderTexture :  Load texture for rendering (framebuffer)
func LoadRenderTexture(width int, height int) () {
C.LoadRenderTexture(C.int(width), C.int(height))
}
//UnloadImage :  Unload image from CPU memory (RAM)
func UnloadImage(image Image) () {
cimage := *image.cptr()
C.UnloadImage(cimage)
}
//UnloadTexture :  Unload texture from GPU memory (VRAM)
func UnloadTexture(texture Texture2D) () {
ctexture := *texture.cptr()
C.UnloadTexture(ctexture)
}
//UnloadRenderTexture :  Unload render texture from GPU memory (VRAM)
func UnloadRenderTexture(target RenderTexture2D) () {
ctarget := *target.cptr()
C.UnloadRenderTexture(ctarget)
}
//GetImageAlphaBorder :  Get image alpha border rectangle
func GetImageAlphaBorder(image Image, threshold float32) () {
cimage := *image.cptr()
C.GetImageAlphaBorder(cimage, C.float(threshold))
}
//GetPixelDataSize :  Get pixel data size in bytes (image or texture)
func GetPixelDataSize(width int, height int, format int) () {
C.GetPixelDataSize(C.int(width), C.int(height), C.int(format))
}
//GetTextureData :  Get pixel data from GPU texture and return an Image
func GetTextureData(texture Texture2D) () {
ctexture := *texture.cptr()
C.GetTextureData(ctexture)
}
//GetScreenData :  Get pixel data from screen buffer and return an Image (screenshot)
func GetScreenData() () {
C.GetScreenData()
}
//ImageCopy :  Create an image duplicate (useful for transformations)
func ImageCopy(image Image) () {
cimage := *image.cptr()
C.ImageCopy(cimage)
}
//ImageFromImage :  Create an image from another image piece
func ImageFromImage(image Image, rec Rectangle) () {
crec := *rec.cptr()
cimage := *image.cptr()
C.ImageFromImage(cimage, crec)
}
//ImageText :  Create an image from text (default font)
func ImageText(text string, fontSize int, color Color) () {
ccolor := *color.cptr()
ctext := C.CString(text)
defer C.free(unsafe.Pointer(&ctext))
C.ImageText(ctext, C.int(fontSize), ccolor)
}
//ImageTextEx :  Create an image from text (custom sprite font)
func ImageTextEx(font Font, text string, fontSize float32, spacing float32, tint Color) () {
ctint := *tint.cptr()
ctext := C.CString(text)
defer C.free(unsafe.Pointer(&ctext))
cfont := *font.cptr()
C.ImageTextEx(cfont, ctext, C.float(fontSize), C.float(spacing), ctint)
}
//GenImageColor :  Generate image: plain color
func GenImageColor(width int, height int, color Color) () {
ccolor := *color.cptr()
C.GenImageColor(C.int(width), C.int(height), ccolor)
}
//GenImageGradientV :  Generate image: vertical gradient
func GenImageGradientV(width int, height int, top Color, bottom Color) () {
cbottom := *bottom.cptr()
ctop := *top.cptr()
C.GenImageGradientV(C.int(width), C.int(height), ctop, cbottom)
}
//GenImageGradientH :  Generate image: horizontal gradient
func GenImageGradientH(width int, height int, left Color, right Color) () {
cright := *right.cptr()
cleft := *left.cptr()
C.GenImageGradientH(C.int(width), C.int(height), cleft, cright)
}
//GenImageGradientRadial :  Generate image: radial gradient
func GenImageGradientRadial(width int, height int, density float32, inner Color, outer Color) () {
couter := *outer.cptr()
cinner := *inner.cptr()
C.GenImageGradientRadial(C.int(width), C.int(height), C.float(density), cinner, couter)
}
//GenImageChecked :  Generate image: checked
func GenImageChecked(width int, height int, checksX int, checksY int, col1 Color, col2 Color) () {
ccol2 := *col2.cptr()
ccol1 := *col1.cptr()
C.GenImageChecked(C.int(width), C.int(height), C.int(checksX), C.int(checksY), ccol1, ccol2)
}
//GenImageWhiteNoise :  Generate image: white noise
func GenImageWhiteNoise(width int, height int, factor float32) () {
C.GenImageWhiteNoise(C.int(width), C.int(height), C.float(factor))
}
//GenImagePerlinNoise :  Generate image: perlin noise
func GenImagePerlinNoise(width int, height int, offsetX int, offsetY int, scale float32) () {
C.GenImagePerlinNoise(C.int(width), C.int(height), C.int(offsetX), C.int(offsetY), C.float(scale))
}
//GenImageCellular :  Generate image: cellular algorithm. Bigger tileSize means bigger cells
func GenImageCellular(width int, height int, tileSize int) () {
C.GenImageCellular(C.int(width), C.int(height), C.int(tileSize))
}
//SetTextureFilter :  Set texture scaling filter mode
func SetTextureFilter(texture Texture2D, filterMode int) () {
ctexture := *texture.cptr()
C.SetTextureFilter(ctexture, C.int(filterMode))
}
//SetTextureWrap :  Set texture wrapping mode
func SetTextureWrap(texture Texture2D, wrapMode int) () {
ctexture := *texture.cptr()
C.SetTextureWrap(ctexture, C.int(wrapMode))
}
//DrawTexture :  Draw a Texture2D
func DrawTexture(texture Texture2D, posX int, posY int, tint Color) () {
ctint := *tint.cptr()
ctexture := *texture.cptr()
C.DrawTexture(ctexture, C.int(posX), C.int(posY), ctint)
}
//DrawTextureV :  Draw a Texture2D with position defined as Vector2
func DrawTextureV(texture Texture2D, position Vector2, tint Color) () {
ctint := *tint.cptr()
cposition := *position.cptr()
ctexture := *texture.cptr()
C.DrawTextureV(ctexture, cposition, ctint)
}
//DrawTextureEx :  Draw a Texture2D with extended parameters
func DrawTextureEx(texture Texture2D, position Vector2, rotation float32, scale float32, tint Color) () {
ctint := *tint.cptr()
cposition := *position.cptr()
ctexture := *texture.cptr()
C.DrawTextureEx(ctexture, cposition, C.float(rotation), C.float(scale), ctint)
}
//DrawTextureRec :  Draw a part of a texture defined by a rectangle
func DrawTextureRec(texture Texture2D, sourceRec Rectangle, position Vector2, tint Color) () {
ctint := *tint.cptr()
cposition := *position.cptr()
csourceRec := *sourceRec.cptr()
ctexture := *texture.cptr()
C.DrawTextureRec(ctexture, csourceRec, cposition, ctint)
}
//DrawTextureQuad :  Draw texture quad with tiling and offset parameters
func DrawTextureQuad(texture Texture2D, tiling Vector2, offset Vector2, quad Rectangle, tint Color) () {
ctint := *tint.cptr()
cquad := *quad.cptr()
coffset := *offset.cptr()
ctiling := *tiling.cptr()
ctexture := *texture.cptr()
C.DrawTextureQuad(ctexture, ctiling, coffset, cquad, ctint)
}
//DrawTexturePro :  Draw a part of a texture defined by a rectangle with 'pro' parameters
func DrawTexturePro(texture Texture2D, sourceRec Rectangle, destRec Rectangle, origin Vector2, rotation float32, tint Color) () {
ctint := *tint.cptr()
corigin := *origin.cptr()
cdestRec := *destRec.cptr()
csourceRec := *sourceRec.cptr()
ctexture := *texture.cptr()
C.DrawTexturePro(ctexture, csourceRec, cdestRec, corigin, C.float(rotation), ctint)
}
//DrawTextureNPatch :  Draws a texture (or part of it) that stretches or shrinks nicely
func DrawTextureNPatch(texture Texture2D, nPatchInfo NPatchInfo, destRec Rectangle, origin Vector2, rotation float32, tint Color) () {
ctint := *tint.cptr()
corigin := *origin.cptr()
cdestRec := *destRec.cptr()
cnPatchInfo := *nPatchInfo.cptr()
ctexture := *texture.cptr()
C.DrawTextureNPatch(ctexture, cnPatchInfo, cdestRec, corigin, C.float(rotation), ctint)
}
//GetFontDefault :  Get the default Font
func GetFontDefault() () {
C.GetFontDefault()
}
//LoadFont :  Load font from file into GPU memory (VRAM)
func LoadFont(fileName string) () {
cfileName := C.CString(fileName)
defer C.free(unsafe.Pointer(&cfileName))
C.LoadFont(cfileName)
}
//LoadFontFromImage :  Load font from Image (XNA style)
func LoadFontFromImage(image Image, key Color, firstChar int) () {
ckey := *key.cptr()
cimage := *image.cptr()
C.LoadFontFromImage(cimage, ckey, C.int(firstChar))
}
//UnloadFont :  Unload Font from GPU memory (VRAM)
func UnloadFont(font Font) () {
cfont := *font.cptr()
C.UnloadFont(cfont)
}
//DrawFPS :  Shows current FPS
func DrawFPS(posX int, posY int) () {
C.DrawFPS(C.int(posX), C.int(posY))
}
//DrawText :  Draw text (using default font)
func DrawText(text string, posX int, posY int, fontSize int, color Color) () {
ccolor := *color.cptr()
ctext := C.CString(text)
defer C.free(unsafe.Pointer(&ctext))
C.DrawText(ctext, C.int(posX), C.int(posY), C.int(fontSize), ccolor)
}
//DrawTextEx :  Draw text using font and additional parameters
func DrawTextEx(font Font, text string, position Vector2, fontSize float32, spacing float32, tint Color) () {
ctint := *tint.cptr()
cposition := *position.cptr()
ctext := C.CString(text)
defer C.free(unsafe.Pointer(&ctext))
cfont := *font.cptr()
C.DrawTextEx(cfont, ctext, cposition, C.float(fontSize), C.float(spacing), ctint)
}
//DrawTextRec :  Draw text using font inside rectangle limits
func DrawTextRec(font Font, text string, rec Rectangle, fontSize float32, spacing float32, wordWrap bool, tint Color) () {
ctint := *tint.cptr()
crec := *rec.cptr()
ctext := C.CString(text)
defer C.free(unsafe.Pointer(&ctext))
cfont := *font.cptr()
C.DrawTextRec(cfont, ctext, crec, C.float(fontSize), C.float(spacing), C.bool(wordWrap), ctint)
}
//MeasureText :  Measure string width for default font
func MeasureText(text string, fontSize int) () {
ctext := C.CString(text)
defer C.free(unsafe.Pointer(&ctext))
C.MeasureText(ctext, C.int(fontSize))
}
//MeasureTextEx :  Measure string size for Font
func MeasureTextEx(font Font, text string, fontSize float32, spacing float32) () {
ctext := C.CString(text)
defer C.free(unsafe.Pointer(&ctext))
cfont := *font.cptr()
C.MeasureTextEx(cfont, ctext, C.float(fontSize), C.float(spacing))
}
//GetGlyphIndex :  Get index position for a unicode character on font
func GetGlyphIndex(font Font, character int) () {
cfont := *font.cptr()
C.GetGlyphIndex(cfont, C.int(character))
}
//TextIsEqual :  Check if two text string are equal
func TextIsEqual(text1 string, text2 string) () {
ctext2 := C.CString(text2)
defer C.free(unsafe.Pointer(&ctext2))
ctext1 := C.CString(text1)
defer C.free(unsafe.Pointer(&ctext1))
C.TextIsEqual(ctext1, ctext2)
}
//TextSubtext :  Get a piece of a text string
func TextSubtext(text string, position int, length int) () {
ctext := C.CString(text)
defer C.free(unsafe.Pointer(&ctext))
C.TextSubtext(ctext, C.int(position), C.int(length))
}
//TextReplace :  Replace text string (memory must be freed!)
func TextReplace(text string, replace string, by string) () {
cby := C.CString(by)
defer C.free(unsafe.Pointer(&cby))
creplace := C.CString(replace)
defer C.free(unsafe.Pointer(&creplace))
ctext := C.CString(text)
defer C.free(unsafe.Pointer(&ctext))
C.TextReplace(ctext, creplace, cby)
}
//TextInsert :  Insert text in a position (memory must be freed!)
func TextInsert(text string, insert string, position int) () {
cinsert := C.CString(insert)
defer C.free(unsafe.Pointer(&cinsert))
ctext := C.CString(text)
defer C.free(unsafe.Pointer(&ctext))
C.TextInsert(ctext, cinsert, C.int(position))
}

//TextJoin :  Join text strings with delimiter
func TextJoin(char string, count int, delimiter string) {
	fmt.Println("test")
}

//TextFindIndex :  Find first text occurrence within a string
func TextFindIndex(text string, find string) () {
cfind := C.CString(find)
defer C.free(unsafe.Pointer(&cfind))
ctext := C.CString(text)
defer C.free(unsafe.Pointer(&ctext))
C.TextFindIndex(ctext, cfind)
}
//TextToUpper :  Get upper case version of provided string
func TextToUpper(text string) () {
ctext := C.CString(text)
defer C.free(unsafe.Pointer(&ctext))
C.TextToUpper(ctext)
}
//TextToLower :  Get lower case version of provided string
func TextToLower(text string) () {
ctext := C.CString(text)
defer C.free(unsafe.Pointer(&ctext))
C.TextToLower(ctext)
}
//TextToPascal :  Get Pascal case notation version of provided string
func TextToPascal(text string) () {
ctext := C.CString(text)
defer C.free(unsafe.Pointer(&ctext))
C.TextToPascal(ctext)
}
//TextToInteger :  Get integer value from text (negative values not supported)
func TextToInteger(text string) () {
ctext := C.CString(text)
defer C.free(unsafe.Pointer(&ctext))
C.TextToInteger(ctext)
}
//GetCodepointsCount :  Get total number of characters (codepoints) in a UTF8 encoded string
func GetCodepointsCount(text string) () {
ctext := C.CString(text)
defer C.free(unsafe.Pointer(&ctext))
C.GetCodepointsCount(ctext)
}
//DrawLine3D :  Draw a line in 3D world space
func DrawLine3D(startPos Vector3, endPos Vector3, color Color) () {
ccolor := *color.cptr()
cendPos := *endPos.cptr()
cstartPos := *startPos.cptr()
C.DrawLine3D(cstartPos, cendPos, ccolor)
}
//DrawCircle3D :  Draw a circle in 3D world space
func DrawCircle3D(center Vector3, radius float32, rotationAxis Vector3, rotationAngle float32, color Color) () {
ccolor := *color.cptr()
crotationAxis := *rotationAxis.cptr()
ccenter := *center.cptr()
C.DrawCircle3D(ccenter, C.float(radius), crotationAxis, C.float(rotationAngle), ccolor)
}
//DrawCube :  Draw cube
func DrawCube(position Vector3, width float32, height float32, length float32, color Color) () {
ccolor := *color.cptr()
cposition := *position.cptr()
C.DrawCube(cposition, C.float(width), C.float(height), C.float(length), ccolor)
}
//DrawCubeV :  Draw cube (Vector version)
func DrawCubeV(position Vector3, size Vector3, color Color) () {
ccolor := *color.cptr()
csize := *size.cptr()
cposition := *position.cptr()
C.DrawCubeV(cposition, csize, ccolor)
}
//DrawCubeWires :  Draw cube wires
func DrawCubeWires(position Vector3, width float32, height float32, length float32, color Color) () {
ccolor := *color.cptr()
cposition := *position.cptr()
C.DrawCubeWires(cposition, C.float(width), C.float(height), C.float(length), ccolor)
}
//DrawCubeWiresV :  Draw cube wires (Vector version)
func DrawCubeWiresV(position Vector3, size Vector3, color Color) () {
ccolor := *color.cptr()
csize := *size.cptr()
cposition := *position.cptr()
C.DrawCubeWiresV(cposition, csize, ccolor)
}
//DrawCubeTexture :  Draw cube textured
func DrawCubeTexture(texture Texture2D, position Vector3, width float32, height float32, length float32, color Color) () {
ccolor := *color.cptr()
cposition := *position.cptr()
ctexture := *texture.cptr()
C.DrawCubeTexture(ctexture, cposition, C.float(width), C.float(height), C.float(length), ccolor)
}
//DrawSphere :  Draw sphere
func DrawSphere(centerPos Vector3, radius float32, color Color) () {
ccolor := *color.cptr()
ccenterPos := *centerPos.cptr()
C.DrawSphere(ccenterPos, C.float(radius), ccolor)
}
//DrawSphereEx :  Draw sphere with extended parameters
func DrawSphereEx(centerPos Vector3, radius float32, rings int, slices int, color Color) () {
ccolor := *color.cptr()
ccenterPos := *centerPos.cptr()
C.DrawSphereEx(ccenterPos, C.float(radius), C.int(rings), C.int(slices), ccolor)
}
//DrawSphereWires :  Draw sphere wires
func DrawSphereWires(centerPos Vector3, radius float32, rings int, slices int, color Color) () {
ccolor := *color.cptr()
ccenterPos := *centerPos.cptr()
C.DrawSphereWires(ccenterPos, C.float(radius), C.int(rings), C.int(slices), ccolor)
}
//DrawCylinder :  Draw a cylinder/cone
func DrawCylinder(position Vector3, radiusTop float32, radiusBottom float32, height float32, slices int, color Color) () {
ccolor := *color.cptr()
cposition := *position.cptr()
C.DrawCylinder(cposition, C.float(radiusTop), C.float(radiusBottom), C.float(height), C.int(slices), ccolor)
}
//DrawCylinderWires :  Draw a cylinder/cone wires
func DrawCylinderWires(position Vector3, radiusTop float32, radiusBottom float32, height float32, slices int, color Color) () {
ccolor := *color.cptr()
cposition := *position.cptr()
C.DrawCylinderWires(cposition, C.float(radiusTop), C.float(radiusBottom), C.float(height), C.int(slices), ccolor)
}
//DrawPlane :  Draw a plane XZ
func DrawPlane(centerPos Vector3, size Vector2, color Color) () {
ccolor := *color.cptr()
csize := *size.cptr()
ccenterPos := *centerPos.cptr()
C.DrawPlane(ccenterPos, csize, ccolor)
}
//DrawRay :  Draw a ray line
func DrawRay(ray Ray, color Color) () {
ccolor := *color.cptr()
cray := *ray.cptr()
C.DrawRay(cray, ccolor)
}
//DrawGrid :  Draw a grid (centered at (0, 0, 0))
func DrawGrid(slices int, spacing float32) () {
C.DrawGrid(C.int(slices), C.float(spacing))
}
//DrawGizmo :  Draw simple gizmo
func DrawGizmo(position Vector3) () {
cposition := *position.cptr()
C.DrawGizmo(cposition)
}
//LoadModel :  Load model from files (meshes and materials)
func LoadModel(fileName string) () {
cfileName := C.CString(fileName)
defer C.free(unsafe.Pointer(&cfileName))
C.LoadModel(cfileName)
}
//LoadModelFromMesh :  Load model from generated mesh (default material)
func LoadModelFromMesh(mesh Mesh) () {
cmesh := *mesh.cptr()
C.LoadModelFromMesh(cmesh)
}
//UnloadModel :  Unload model from memory (RAM and/or VRAM)
func UnloadModel(model Model) () {
cmodel := *model.cptr()
C.UnloadModel(cmodel)
}
//ExportMesh :  Export mesh data to file
func ExportMesh(mesh Mesh, fileName string) () {
cfileName := C.CString(fileName)
defer C.free(unsafe.Pointer(&cfileName))
cmesh := *mesh.cptr()
C.ExportMesh(cmesh, cfileName)
}
//UnloadMesh :  Unload mesh from memory (RAM and/or VRAM)
func UnloadMesh(mesh Mesh) () {
cmesh := *mesh.cptr()
C.UnloadMesh(cmesh)
}
//LoadMaterialDefault :  Load default material (Supports: DIFFUSE, SPECULAR, NORMAL maps)
func LoadMaterialDefault() () {
C.LoadMaterialDefault()
}
//UnloadMaterial :  Unload material from GPU memory (VRAM)
func UnloadMaterial(material Material) () {
cmaterial := *material.cptr()
C.UnloadMaterial(cmaterial)
}
//UpdateModelAnimation :  Update model animation pose
func UpdateModelAnimation(model Model, anim ModelAnimation, frame int) () {
canim := *anim.cptr()
cmodel := *model.cptr()
C.UpdateModelAnimation(cmodel, canim, C.int(frame))
}
//UnloadModelAnimation :  Unload animation data
func UnloadModelAnimation(anim ModelAnimation) () {
canim := *anim.cptr()
C.UnloadModelAnimation(canim)
}
//IsModelAnimationValid :  Check model animation skeleton match
func IsModelAnimationValid(model Model, anim ModelAnimation) () {
canim := *anim.cptr()
cmodel := *model.cptr()
C.IsModelAnimationValid(cmodel, canim)
}
//GenMeshPoly :  Generate polygonal mesh
func GenMeshPoly(sides int, radius float32) () {
C.GenMeshPoly(C.int(sides), C.float(radius))
}
//GenMeshPlane :  Generate plane mesh (with subdivisions)
func GenMeshPlane(width float32, length float32, resX int, resZ int) () {
C.GenMeshPlane(C.float(width), C.float(length), C.int(resX), C.int(resZ))
}
//GenMeshCube :  Generate cuboid mesh
func GenMeshCube(width float32, height float32, length float32) () {
C.GenMeshCube(C.float(width), C.float(height), C.float(length))
}
//GenMeshSphere :  Generate sphere mesh (standard sphere)
func GenMeshSphere(radius float32, rings int, slices int) () {
C.GenMeshSphere(C.float(radius), C.int(rings), C.int(slices))
}
//GenMeshHemiSphere :  Generate half-sphere mesh (no bottom cap)
func GenMeshHemiSphere(radius float32, rings int, slices int) () {
C.GenMeshHemiSphere(C.float(radius), C.int(rings), C.int(slices))
}
//GenMeshCylinder :  Generate cylinder mesh
func GenMeshCylinder(radius float32, height float32, slices int) () {
C.GenMeshCylinder(C.float(radius), C.float(height), C.int(slices))
}
//GenMeshTorus :  Generate torus mesh
func GenMeshTorus(radius float32, size float32, radSeg int, sides int) () {
C.GenMeshTorus(C.float(radius), C.float(size), C.int(radSeg), C.int(sides))
}
//GenMeshKnot :  Generate trefoil knot mesh
func GenMeshKnot(radius float32, size float32, radSeg int, sides int) () {
C.GenMeshKnot(C.float(radius), C.float(size), C.int(radSeg), C.int(sides))
}
//GenMeshHeightmap :  Generate heightmap mesh from image data
func GenMeshHeightmap(heightmap Image, size Vector3) () {
csize := *size.cptr()
cheightmap := *heightmap.cptr()
C.GenMeshHeightmap(cheightmap, csize)
}
//GenMeshCubicmap :  Generate cubes-based map mesh from image data
func GenMeshCubicmap(cubicmap Image, cubeSize Vector3) () {
ccubeSize := *cubeSize.cptr()
ccubicmap := *cubicmap.cptr()
C.GenMeshCubicmap(ccubicmap, ccubeSize)
}
//MeshBoundingBox :  Compute mesh bounding box limits
func MeshBoundingBox(mesh Mesh) () {
cmesh := *mesh.cptr()
C.MeshBoundingBox(cmesh)
}
//DrawModel :  Draw a model (with texture if set)
func DrawModel(model Model, position Vector3, scale float32, tint Color) () {
ctint := *tint.cptr()
cposition := *position.cptr()
cmodel := *model.cptr()
C.DrawModel(cmodel, cposition, C.float(scale), ctint)
}
//DrawModelEx :  Draw a model with extended parameters
func DrawModelEx(model Model, position Vector3, rotationAxis Vector3, rotationAngle float32, scale Vector3, tint Color) () {
ctint := *tint.cptr()
cscale := *scale.cptr()
crotationAxis := *rotationAxis.cptr()
cposition := *position.cptr()
cmodel := *model.cptr()
C.DrawModelEx(cmodel, cposition, crotationAxis, C.float(rotationAngle), cscale, ctint)
}
//DrawModelWires :  Draw a model wires (with texture if set)
func DrawModelWires(model Model, position Vector3, scale float32, tint Color) () {
ctint := *tint.cptr()
cposition := *position.cptr()
cmodel := *model.cptr()
C.DrawModelWires(cmodel, cposition, C.float(scale), ctint)
}
//DrawModelWiresEx :  Draw a model wires (with texture if set) with extended parameters
func DrawModelWiresEx(model Model, position Vector3, rotationAxis Vector3, rotationAngle float32, scale Vector3, tint Color) () {
ctint := *tint.cptr()
cscale := *scale.cptr()
crotationAxis := *rotationAxis.cptr()
cposition := *position.cptr()
cmodel := *model.cptr()
C.DrawModelWiresEx(cmodel, cposition, crotationAxis, C.float(rotationAngle), cscale, ctint)
}
//DrawBoundingBox :  Draw bounding box (wires)
func DrawBoundingBox(box BoundingBox, color Color) () {
ccolor := *color.cptr()
cbox := *box.cptr()
C.DrawBoundingBox(cbox, ccolor)
}
//DrawBillboard :  Draw a billboard texture
func DrawBillboard(camera Camera, texture Texture2D, center Vector3, size float32, tint Color) () {
ctint := *tint.cptr()
ccenter := *center.cptr()
ctexture := *texture.cptr()
ccamera := *camera.cptr()
C.DrawBillboard(ccamera, ctexture, ccenter, C.float(size), ctint)
}
//DrawBillboardRec :  Draw a billboard texture defined by sourceRec
func DrawBillboardRec(camera Camera, texture Texture2D, sourceRec Rectangle, center Vector3, size float32, tint Color) () {
ctint := *tint.cptr()
ccenter := *center.cptr()
csourceRec := *sourceRec.cptr()
ctexture := *texture.cptr()
ccamera := *camera.cptr()
C.DrawBillboardRec(ccamera, ctexture, csourceRec, ccenter, C.float(size), ctint)
}
//CheckCollisionSpheres :  Detect collision between two spheres
func CheckCollisionSpheres(centerA Vector3, radiusA float32, centerB Vector3, radiusB float32) () {
ccenterB := *centerB.cptr()
ccenterA := *centerA.cptr()
C.CheckCollisionSpheres(ccenterA, C.float(radiusA), ccenterB, C.float(radiusB))
}
//CheckCollisionBoxes :  Detect collision between two bounding boxes
func CheckCollisionBoxes(box1 BoundingBox, box2 BoundingBox) () {
cbox2 := *box2.cptr()
cbox1 := *box1.cptr()
C.CheckCollisionBoxes(cbox1, cbox2)
}
//CheckCollisionBoxSphere :  Detect collision between box and sphere
func CheckCollisionBoxSphere(box BoundingBox, center Vector3, radius float32) () {
ccenter := *center.cptr()
cbox := *box.cptr()
C.CheckCollisionBoxSphere(cbox, ccenter, C.float(radius))
}
//CheckCollisionRaySphere :  Detect collision between ray and sphere
func CheckCollisionRaySphere(ray Ray, center Vector3, radius float32) () {
ccenter := *center.cptr()
cray := *ray.cptr()
C.CheckCollisionRaySphere(cray, ccenter, C.float(radius))
}
//CheckCollisionRayBox :  Detect collision between ray and box
func CheckCollisionRayBox(ray Ray, box BoundingBox) () {
cbox := *box.cptr()
cray := *ray.cptr()
C.CheckCollisionRayBox(cray, cbox)
}
//GetCollisionRayModel :  Get collision info between ray and model
func GetCollisionRayModel(ray Ray, model Model) () {
cmodel := *model.cptr()
cray := *ray.cptr()
C.GetCollisionRayModel(cray, cmodel)
}
//GetCollisionRayTriangle :  Get collision info between ray and triangle
func GetCollisionRayTriangle(ray Ray, p1 Vector3, p2 Vector3, p3 Vector3) () {
cp3 := *p3.cptr()
cp2 := *p2.cptr()
cp1 := *p1.cptr()
cray := *ray.cptr()
C.GetCollisionRayTriangle(cray, cp1, cp2, cp3)
}
//GetCollisionRayGround :  Get collision info between ray and ground plane (Y-normal plane)
func GetCollisionRayGround(ray Ray, groundHeight float32) () {
cray := *ray.cptr()
C.GetCollisionRayGround(cray, C.float(groundHeight))
}
//LoadText :  Load chars array from text file
func LoadText(fileName string) () {
cfileName := C.CString(fileName)
defer C.free(unsafe.Pointer(&cfileName))
C.LoadText(cfileName)
}
//LoadShader :  Load shader from files and bind default locations
func LoadShader(vsFileName string, fsFileName string) () {
cfsFileName := C.CString(fsFileName)
defer C.free(unsafe.Pointer(&cfsFileName))
cvsFileName := C.CString(vsFileName)
defer C.free(unsafe.Pointer(&cvsFileName))
C.LoadShader(cvsFileName, cfsFileName)
}
//LoadShaderCode :  Load shader from code strings and bind default locations
func LoadShaderCode(vsCode string, fsCode string) () {
cfsCode := C.CString(fsCode)
defer C.free(unsafe.Pointer(&cfsCode))
cvsCode := C.CString(vsCode)
defer C.free(unsafe.Pointer(&cvsCode))
C.LoadShaderCode(cvsCode, cfsCode)
}
//UnloadShader :  Unload shader from GPU memory (VRAM)
func UnloadShader(shader Shader) () {
cshader := *shader.cptr()
C.UnloadShader(cshader)
}
//GetShaderDefault :  Get default shader
func GetShaderDefault() () {
C.GetShaderDefault()
}
//GetTextureDefault :  Get default texture
func GetTextureDefault() () {
C.GetTextureDefault()
}
//GetShaderLocation :  Get shader uniform location
func GetShaderLocation(shader Shader, uniformName string) () {
cuniformName := C.CString(uniformName)
defer C.free(unsafe.Pointer(&cuniformName))
cshader := *shader.cptr()
C.GetShaderLocation(cshader, cuniformName)
}
//SetShaderValueMatrix :  Set shader uniform value (matrix 4x4)
func SetShaderValueMatrix(shader Shader, uniformLoc int, mat Matrix) () {
cmat := *mat.cptr()
cshader := *shader.cptr()
C.SetShaderValueMatrix(cshader, C.int(uniformLoc), cmat)
}
//SetShaderValueTexture :  Set shader uniform value for texture
func SetShaderValueTexture(shader Shader, uniformLoc int, texture Texture2D) () {
ctexture := *texture.cptr()
cshader := *shader.cptr()
C.SetShaderValueTexture(cshader, C.int(uniformLoc), ctexture)
}
//SetMatrixProjection :  Set a custom projection matrix (replaces internal projection matrix)
func SetMatrixProjection(proj Matrix) () {
cproj := *proj.cptr()
C.SetMatrixProjection(cproj)
}
//SetMatrixModelview :  Set a custom modelview matrix (replaces internal modelview matrix)
func SetMatrixModelview(view Matrix) () {
cview := *view.cptr()
C.SetMatrixModelview(cview)
}
//GetMatrixModelview :  Get internal modelview matrix
func GetMatrixModelview() () {
C.GetMatrixModelview()
}
//GetMatrixProjection :  Get internal projection matrix
func GetMatrixProjection() () {
C.GetMatrixProjection()
}
//GenTextureCubemap :  Generate cubemap texture from HDR texture
func GenTextureCubemap(shader Shader, skyHDR Texture2D, size int) () {
cskyHDR := *skyHDR.cptr()
cshader := *shader.cptr()
C.GenTextureCubemap(cshader, cskyHDR, C.int(size))
}
//GenTextureIrradiance :  Generate irradiance texture using cubemap data
func GenTextureIrradiance(shader Shader, cubemap Texture2D, size int) () {
ccubemap := *cubemap.cptr()
cshader := *shader.cptr()
C.GenTextureIrradiance(cshader, ccubemap, C.int(size))
}
//GenTexturePrefilter :  Generate prefilter texture using cubemap data
func GenTexturePrefilter(shader Shader, cubemap Texture2D, size int) () {
ccubemap := *cubemap.cptr()
cshader := *shader.cptr()
C.GenTexturePrefilter(cshader, ccubemap, C.int(size))
}
//GenTextureBRDF :  Generate BRDF texture
func GenTextureBRDF(shader Shader, size int) () {
cshader := *shader.cptr()
C.GenTextureBRDF(cshader, C.int(size))
}
//BeginShaderMode :  Begin custom shader drawing
func BeginShaderMode(shader Shader) () {
cshader := *shader.cptr()
C.BeginShaderMode(cshader)
}
//EndShaderMode :  End custom shader drawing (use default shader)
func EndShaderMode() () {
C.EndShaderMode()
}
//BeginBlendMode :  Begin blending mode (alpha, additive, multiplied)
func BeginBlendMode(mode int) () {
C.BeginBlendMode(C.int(mode))
}
//EndBlendMode :  End blending mode (reset to default: alpha blending)
func EndBlendMode() () {
C.EndBlendMode()
}
//InitVrSimulator :  Init VR simulator for selected device parameters
func InitVrSimulator() () {
C.InitVrSimulator()
}
//CloseVrSimulator :  Close VR simulator for current device
func CloseVrSimulator() () {
C.CloseVrSimulator()
}
//SetVrConfiguration :  Set stereo rendering configuration parameters
func SetVrConfiguration(info VrDeviceInfo, distortion Shader) () {
cdistortion := *distortion.cptr()
cinfo := *info.cptr()
C.SetVrConfiguration(cinfo, cdistortion)
}
//IsVrSimulatorReady :  Detect if VR simulator is ready
func IsVrSimulatorReady() () {
C.IsVrSimulatorReady()
}
//ToggleVrMode :  Enable/Disable VR experience
func ToggleVrMode() () {
C.ToggleVrMode()
}
//BeginVrDrawing :  Begin VR simulator stereo rendering
func BeginVrDrawing() () {
C.BeginVrDrawing()
}
//EndVrDrawing :  End VR simulator stereo rendering
func EndVrDrawing() () {
C.EndVrDrawing()
}
//InitAudioDevice :  Initialize audio device and context
func InitAudioDevice() () {
C.InitAudioDevice()
}
//CloseAudioDevice :  Close the audio device and context
func CloseAudioDevice() () {
C.CloseAudioDevice()
}
//IsAudioDeviceReady :  Check if audio device has been initialized successfully
func IsAudioDeviceReady() () {
C.IsAudioDeviceReady()
}
//SetMasterVolume :  Set master volume (listener)
func SetMasterVolume(volume float32) () {
C.SetMasterVolume(C.float(volume))
}
//LoadWave :  Load wave data from file
func LoadWave(fileName string) () {
cfileName := C.CString(fileName)
defer C.free(unsafe.Pointer(&cfileName))
C.LoadWave(cfileName)
}
//LoadSound :  Load sound from file
func LoadSound(fileName string) () {
cfileName := C.CString(fileName)
defer C.free(unsafe.Pointer(&cfileName))
C.LoadSound(cfileName)
}
//LoadSoundFromWave :  Load sound from wave data
func LoadSoundFromWave(wave Wave) () {
cwave := *wave.cptr()
C.LoadSoundFromWave(cwave)
}
//UnloadWave :  Unload wave data
func UnloadWave(wave Wave) () {
cwave := *wave.cptr()
C.UnloadWave(cwave)
}
//UnloadSound :  Unload sound
func UnloadSound(sound Sound) () {
csound := *sound.cptr()
C.UnloadSound(csound)
}
//ExportWave :  Export wave data to file
func ExportWave(wave Wave, fileName string) () {
cfileName := C.CString(fileName)
defer C.free(unsafe.Pointer(&cfileName))
cwave := *wave.cptr()
C.ExportWave(cwave, cfileName)
}
//ExportWaveAsCode :  Export wave sample data to code (.h)
func ExportWaveAsCode(wave Wave, fileName string) () {
cfileName := C.CString(fileName)
defer C.free(unsafe.Pointer(&cfileName))
cwave := *wave.cptr()
C.ExportWaveAsCode(cwave, cfileName)
}
//PlaySound :  Play a sound
func PlaySound(sound Sound) () {
csound := *sound.cptr()
C.PlaySound(csound)
}
//StopSound :  Stop playing a sound
func StopSound(sound Sound) () {
csound := *sound.cptr()
C.StopSound(csound)
}
//PauseSound :  Pause a sound
func PauseSound(sound Sound) () {
csound := *sound.cptr()
C.PauseSound(csound)
}
//ResumeSound :  Resume a paused sound
func ResumeSound(sound Sound) () {
csound := *sound.cptr()
C.ResumeSound(csound)
}
//PlaySoundMulti :  Play a sound (using multichannel buffer pool)
func PlaySoundMulti(sound Sound) () {
csound := *sound.cptr()
C.PlaySoundMulti(csound)
}
//StopSoundMulti :  Stop any sound playing (using multichannel buffer pool)
func StopSoundMulti() () {
C.StopSoundMulti()
}
//GetSoundsPlaying :  Get number of sounds playing in the multichannel
func GetSoundsPlaying() () {
C.GetSoundsPlaying()
}
//IsSoundPlaying :  Check if a sound is currently playing
func IsSoundPlaying(sound Sound) () {
csound := *sound.cptr()
C.IsSoundPlaying(csound)
}
//SetSoundVolume :  Set volume for a sound (1.0 is max level)
func SetSoundVolume(sound Sound, volume float32) () {
csound := *sound.cptr()
C.SetSoundVolume(csound, C.float(volume))
}
//SetSoundPitch :  Set pitch for a sound (1.0 is base level)
func SetSoundPitch(sound Sound, pitch float32) () {
csound := *sound.cptr()
C.SetSoundPitch(csound, C.float(pitch))
}
//WaveCopy :  Copy a wave to a new wave
func WaveCopy(wave Wave) () {
cwave := *wave.cptr()
C.WaveCopy(cwave)
}
//LoadMusicStream :  Load music stream from file
func LoadMusicStream(fileName string) () {
cfileName := C.CString(fileName)
defer C.free(unsafe.Pointer(&cfileName))
C.LoadMusicStream(cfileName)
}
//UnloadMusicStream :  Unload music stream
func UnloadMusicStream(music Music) () {
cmusic := *music.cptr()
C.UnloadMusicStream(cmusic)
}
//PlayMusicStream :  Start music playing
func PlayMusicStream(music Music) () {
cmusic := *music.cptr()
C.PlayMusicStream(cmusic)
}
//UpdateMusicStream :  Updates buffers for music streaming
func UpdateMusicStream(music Music) () {
cmusic := *music.cptr()
C.UpdateMusicStream(cmusic)
}
//StopMusicStream :  Stop music playing
func StopMusicStream(music Music) () {
cmusic := *music.cptr()
C.StopMusicStream(cmusic)
}
//PauseMusicStream :  Pause music playing
func PauseMusicStream(music Music) () {
cmusic := *music.cptr()
C.PauseMusicStream(cmusic)
}
//ResumeMusicStream :  Resume playing paused music
func ResumeMusicStream(music Music) () {
cmusic := *music.cptr()
C.ResumeMusicStream(cmusic)
}
//IsMusicPlaying :  Check if music is playing
func IsMusicPlaying(music Music) () {
cmusic := *music.cptr()
C.IsMusicPlaying(cmusic)
}
//SetMusicVolume :  Set volume for music (1.0 is max level)
func SetMusicVolume(music Music, volume float32) () {
cmusic := *music.cptr()
C.SetMusicVolume(cmusic, C.float(volume))
}
//SetMusicPitch :  Set pitch for a music (1.0 is base level)
func SetMusicPitch(music Music, pitch float32) () {
cmusic := *music.cptr()
C.SetMusicPitch(cmusic, C.float(pitch))
}
//SetMusicLoopCount :  Set music loop count (loop repeats)
func SetMusicLoopCount(music Music, count int) () {
cmusic := *music.cptr()
C.SetMusicLoopCount(cmusic, C.int(count))
}
//GetMusicTimeLength :  Get music time length (in seconds)
func GetMusicTimeLength(music Music) () {
cmusic := *music.cptr()
C.GetMusicTimeLength(cmusic)
}
//GetMusicTimePlayed :  Get current music time played (in seconds)
func GetMusicTimePlayed(music Music) () {
cmusic := *music.cptr()
C.GetMusicTimePlayed(cmusic)
}
//InitAudioStream :  Init audio stream (to stream raw audio pcm data)
func InitAudioStream(sampleRate int, sampleSize int, channels int) () {
C.InitAudioStream(C.int(sampleRate), C.int(sampleSize), C.int(channels))
}
//CloseAudioStream :  Close audio stream and free memory
func CloseAudioStream(stream AudioStream) () {
cstream := *stream.cptr()
C.CloseAudioStream(cstream)
}
//IsAudioStreamProcessed :  Check if any audio stream buffers requires refill
func IsAudioStreamProcessed(stream AudioStream) () {
cstream := *stream.cptr()
C.IsAudioStreamProcessed(cstream)
}
//PlayAudioStream :  Play audio stream
func PlayAudioStream(stream AudioStream) () {
cstream := *stream.cptr()
C.PlayAudioStream(cstream)
}
//PauseAudioStream :  Pause audio stream
func PauseAudioStream(stream AudioStream) () {
cstream := *stream.cptr()
C.PauseAudioStream(cstream)
}
//ResumeAudioStream :  Resume audio stream
func ResumeAudioStream(stream AudioStream) () {
cstream := *stream.cptr()
C.ResumeAudioStream(cstream)
}
//IsAudioStreamPlaying :  Check if audio stream is playing
func IsAudioStreamPlaying(stream AudioStream) () {
cstream := *stream.cptr()
C.IsAudioStreamPlaying(cstream)
}
//StopAudioStream :  Stop audio stream
func StopAudioStream(stream AudioStream) () {
cstream := *stream.cptr()
C.StopAudioStream(cstream)
}
//SetAudioStreamVolume :  Set volume for audio stream (1.0 is max level)
func SetAudioStreamVolume(stream AudioStream, volume float32) () {
cstream := *stream.cptr()
C.SetAudioStreamVolume(cstream, C.float(volume))
}
//SetAudioStreamPitch :  Set pitch for audio stream (1.0 is base level)
func SetAudioStreamPitch(stream AudioStream, pitch float32) () {
cstream := *stream.cptr()
C.SetAudioStreamPitch(cstream, C.float(pitch))
}