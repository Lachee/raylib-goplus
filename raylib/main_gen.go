package raylib

/*
//Generated 2019-11-12T14:15:54+11:00
#include "raylib.h"
#include <stdlib.h>
#include "go.h"
*/
import "C"
import (
	"os/exec"
	"runtime"
	"unsafe"
)

//InitWindow : Initialize window and OpenGL context
func InitWindow(width int, height int, title string) {
	ctitle := C.CString(title)
	defer C.free(unsafe.Pointer(ctitle))
	C.InitWindow(C.int(int32(width)), C.int(int32(height)), ctitle)
}

//WindowShouldClose : Check if KEY_ESCAPE pressed or Close icon pressed
func WindowShouldClose() bool {
	res := C.WindowShouldClose()
	return bool(res)
}

//CloseWindow : Close window and unload OpenGL context
func CloseWindow() {
	C.CloseWindow()
}

//IsWindowReady : Check if window has been initialized successfully
func IsWindowReady() bool {
	res := C.IsWindowReady()
	return bool(res)
}

//IsWindowMinimized : Check if window has been minimized (or lost focus)
func IsWindowMinimized() bool {
	res := C.IsWindowMinimized()
	return bool(res)
}

//IsWindowResized : Check if window has been resized
func IsWindowResized() bool {
	res := C.IsWindowResized()
	return bool(res)
}

//IsWindowHidden : Check if window is currently hidden
func IsWindowHidden() bool {
	res := C.IsWindowHidden()
	return bool(res)
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
	C.SetWindowPosition(C.int(int32(x)), C.int(int32(y)))
}

//SetWindowMonitor : Set monitor for the current window (fullscreen mode)
func SetWindowMonitor(monitor int) {
	C.SetWindowMonitor(C.int(int32(monitor)))
}

//SetWindowMinSize : Set window minimum dimensions (for FLAG_WINDOW_RESIZABLE)
func SetWindowMinSize(width int, height int) {
	C.SetWindowMinSize(C.int(int32(width)), C.int(int32(height)))
}

//SetWindowSize : Set window dimensions
func SetWindowSize(width int, height int) {
	C.SetWindowSize(C.int(int32(width)), C.int(int32(height)))
}

//GetWindowHandle : Get native window handle
func GetWindowHandle() {
	C.GetWindowHandle()
}

//GetScreenWidth : Get current screen width
func GetScreenWidth() int {
	res := C.GetScreenWidth()
	return int(int32(res))
}

//GetScreenHeight : Get current screen height
func GetScreenHeight() int {
	res := C.GetScreenHeight()
	return int(int32(res))
}

//GetMonitorCount : Get number of connected monitors
func GetMonitorCount() int {
	res := C.GetMonitorCount()
	return int(int32(res))
}

//GetMonitorWidth : Get primary monitor width
func GetMonitorWidth(monitor int) int {
	res := C.GetMonitorWidth(C.int(int32(monitor)))
	return int(int32(res))
}

//GetMonitorHeight : Get primary monitor height
func GetMonitorHeight(monitor int) int {
	res := C.GetMonitorHeight(C.int(int32(monitor)))
	return int(int32(res))
}

//GetMonitorPhysicalWidth : Get primary monitor physical width in millimetres
func GetMonitorPhysicalWidth(monitor int) int {
	res := C.GetMonitorPhysicalWidth(C.int(int32(monitor)))
	return int(int32(res))
}

//GetMonitorPhysicalHeight : Get primary monitor physical height in millimetres
func GetMonitorPhysicalHeight(monitor int) int {
	res := C.GetMonitorPhysicalHeight(C.int(int32(monitor)))
	return int(int32(res))
}

//GetWindowPosition : Get window position XY on monitor
func GetWindowPosition() Vector2 {
	res := C.GetWindowPosition()
	return newVector2FromPointer(unsafe.Pointer(&res))
}

//GetMonitorName : Get the human-readable, UTF-8 encoded name of the primary monitor
func GetMonitorName(monitor int) string {
	res := C.GetMonitorName(C.int(int32(monitor)))
	return C.GoString(res)
}

//GetClipboardText : Get clipboard text content
func GetClipboardText() string {
	res := C.GetClipboardText()
	return C.GoString(res)
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
	return bool(res)
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
func BeginMode3D(camera Camera) {
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
	C.BeginScissorMode(C.int(int32(x)), C.int(int32(y)), C.int(int32(width)), C.int(int32(height)))
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
	C.SetTargetFPS(C.int(int32(fps)))
}

//GetFPS : Returns current FPS
func GetFPS() int {
	res := C.GetFPS()
	return int(int32(res))
}

//GetFrameTime : Returns time in seconds for last frame drawn
func GetFrameTime() float32 {
	res := C.GetFrameTime()
	return float32(res)
}

//GetTime : Returns elapsed time in seconds since InitWindow()
func GetTime() float64 {
	res := C.GetTime()
	return float64(res)
}

//ColorToInt : Returns hexadecimal value for a Color
func ColorToInt(color Color) int {
	ccolor := *color.cptr()
	res := C.ColorToInt(ccolor)
	return int(int32(res))
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
	res := C.GetColor(C.int(int32(hexValue)))
	return newColorFromPointer(unsafe.Pointer(&res))
}

//Fade : Color fade-in or fade-out, alpha goes from 0.0f to 1.0f
func Fade(color Color, alpha float32) Color {
	ccolor := *color.cptr()
	res := C.Fade(ccolor, C.float(alpha))
	return newColorFromPointer(unsafe.Pointer(&res))
}

//SetConfigFlags : Setup window configuration flags (view FLAGS)
func SetConfigFlags(flags uint32) {
	C.SetConfigFlags(C.uint(flags))
}

//SetTraceLogLevel is in trace.go

//SetTraceLogExit is in trace.go

//SetTraceLogCallback is in trace.go
//TakeScreenshot : Takes a screenshot of current screen (saved a .png)
func TakeScreenshot(fileName string) {
	cfileName := C.CString(fileName)
	defer C.free(unsafe.Pointer(cfileName))
	C.TakeScreenshot(cfileName)
}

//GetRandomValue : Returns a random value between min and max (both included)
func GetRandomValue(min int, max int) int {
	res := C.GetRandomValue(C.int(int32(min)), C.int(int32(max)))
	return int(int32(res))
}

//IsFileDropped : Check if a file has been dropped into window
func IsFileDropped() bool {
	res := C.IsFileDropped()
	return bool(res)
}

//GetDroppedFiles : Get dropped files names (memory should be freed)
func GetDroppedFiles() []string {
	ccount := C.int(0)
	res := C.GetDroppedFiles(&ccount)
	count := int(ccount)

	tmpslice := (*[1 << 24]*C.char)(unsafe.Pointer(res))[:count:count]
	gostrings := make([]string, count)
	for i, s := range tmpslice {
		gostrings[i] = C.GoString(s)
	}

	return gostrings
}

//ClearDroppedFiles : Clear dropped files paths buffer (free memory)
func ClearDroppedFiles() {
	C.ClearDroppedFiles()
}

//StorageSaveValue : Save integer value to storage file (to defined position)
func StorageSaveValue(position int, value int) {
	C.StorageSaveValue(C.int(int32(position)), C.int(int32(value)))
}

//StorageLoadValue : Load integer value from storage file (from defined position)
func StorageLoadValue(position int) int {
	res := C.StorageLoadValue(C.int(int32(position)))
	return int(int32(res))
}

//OpenURL opens a URL in the system browser.
func OpenURL(url string) error {
	switch runtime.GOOS {
	case "linux":
		return exec.Command("xdg-open", url).Start()
	case "windows":
		return exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		return exec.Command("open", url).Start()
	default:
		cstring := C.CString(url)
		defer C.free(unsafe.Pointer(cstring))
		C.OpenURL(cstring)
		return nil
	}
}
