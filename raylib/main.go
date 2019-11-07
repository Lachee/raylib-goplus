package raylib

/*
#cgo debug CFLAGS: -O0
#cgo CFLAGS: -w

#cgo android LDFLAGS: -llog -landroid -lEGL -lGLESv2 -lOpenSLES -lm
#cgo android CFLAGS: -DPLATFORM_ANDROID -DGRAPHICS_API_OPENGL_ES2 -DMAL_NO_SDL -Iexternal/android/native_app_glue

#cgo darwin LDFLAGS: -framework OpenGL -framework Cocoa -framework IOKit -framework CoreVideo -framework CoreFoundation
#cgo darwin CFLAGS: -x objective-c -Iexternal/glfw/include -D_GLFW_COCOA -D_GLFW_USE_CHDIR -D_GLFW_USE_MENUBAR -D_GLFW_USE_RETINA -Wno-deprecated-declarations -DPLATFORM_DESKTOP -DMAL_NO_COREAUDIO

#cgo darwin,opengl11 CFLAGS: -DGRAPHICS_API_OPENGL_11
#cgo darwin,opengl21 CFLAGS: -DGRAPHICS_API_OPENGL_21
#cgo darwin,!opengl11,!opengl21 CFLAGS: -DGRAPHICS_API_OPENGL_33

#cgo linux,arm LDFLAGS: -L/opt/vc/lib -L/opt/vc/lib64 -lbrcmGLESv2 -lbrcmEGL -lpthread -lrt -lm -lbcm_host -lvcos -lvchiq_arm -ldl
#cgo linux,arm CFLAGS: -DPLATFORM_RPI -DGRAPHICS_API_OPENGL_ES2 -I/opt/vc/include -I/opt/vc/include/interface/vcos -I/opt/vc/include/interface/vmcs_host/linux -I/opt/vc/include/interface/vcos/pthreads

#cgo linux CFLAGS: -Iexternal/glfw/include -DPLATFORM_DESKTOP -Wno-stringop-overflow

#cgo linux,!wayland LDFLAGS: -lGL -lm -pthread -ldl -lrt -lX11
#cgo linux,wayland LDFLAGS: -lGL -lm -pthread -ldl -lrt -lwayland-client -lwayland-cursor -lwayland-egl -lxkbcommon

#cgo linux,!wayland CFLAGS: -D_GLFW_X11
#cgo linux,wayland CFLAGS: -D_GLFW_WAYLAND

#cgo linux,opengl11 CFLAGS: -DGRAPHICS_API_OPENGL_11
#cgo linux,opengl21 CFLAGS: -DGRAPHICS_API_OPENGL_21
#cgo linux,!opengl11,!opengl21 CFLAGS: -DGRAPHICS_API_OPENGL_33

#cgo windows LDFLAGS: -lopengl32 -lgdi32 -lwinmm -lole32
#cgo windows CFLAGS: -D_GLFW_WIN32 -Iexternal -Iexternal/glfw/include -Iexternal/glfw/deps/mingw -DPLATFORM_DESKTOP -DRAYGUI_IMPLEMENTATION

#cgo windows,opengl11 CFLAGS: -DGRAPHICS_API_OPENGL_11
#cgo windows,opengl21 CFLAGS: -DGRAPHICS_API_OPENGL_21
#cgo windows,!opengl11,!opengl21 CFLAGS: -DGRAPHICS_API_OPENGL_33

#include "raylib.h"
#include <stdlib.h>
*/
import "C"
import (
	"runtime"
	"time"
	"unsafe"
)

const (
	// Set to show raylib logo at startup
	FlagShowLogo = 1
	// Set to run program in fullscreen
	FlagFullscreenMode = 2
	// Set to allow resizable window
	FlagWindowResizable = 4
	// Set to disable window decoration (frame and buttons)
	FlagWindowUndecorated = 8
	// Set to allow transparent window
	FlagWindowTransparent = 16
	// Set to try enabling MSAA 4X
	FlagMsaa4xHint = 32
	// Set to try enabling V-Sync on GPU
	FlagVsyncHint = 64
)

func init() {
	//We have to lock the OS Thread as raylib is sensitive to that stuff.
	runtime.LockOSThread()

	//Setup the unloadables so they finalize
	runtime.SetFinalizer(&unloadables, finalizeUnloadables)
	time.Sleep(time.Second)
	unloadables = nil
	time.Sleep(time.Second)
	runtime.GC()
}

// SetConfigFlags - Setup some window configuration flags
func SetConfigFlags(flags int) {
	cflags := (C.uint)(flags)
	C.SetConfigFlags(cflags)
}

//WindowShouldClose Should the window clsoe
func WindowShouldClose() bool {
	return bool(C.WindowShouldClose())
}

// InitWindow - Initialize Window and OpenGL Graphics
func InitWindow(width int, height int, title string) {
	cwidth := (C.int)(width)
	cheight := (C.int)(height)

	ctitle := C.CString(title)
	defer C.free(unsafe.Pointer(ctitle))

	C.InitWindow(cwidth, cheight, ctitle)
}

func BeginDrawing() {
	runtime.LockOSThread()
	C.BeginDrawing()
}

func EndDrawing() {
	C.EndDrawing()
	runtime.UnlockOSThread()
}

func ClearBackground(color Color) {
	clr := *color.cptr()
	C.ClearBackground(clr)
}

func CloseWindow() {
	C.CloseWindow()
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
