package raylib

/*
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
#define SUPPORT_TRACELOG
#include "raylib.h"
#include <stdlib.h>
#include "go.h"
*/
import "C"
import (
	"runtime"
)

//#cgo debug CFLAGS: -O0
//

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

var screenWidth = 0
var screenHeight = 0

func init() {
	//We have to lock the OS Thread as raylib is sensitive to that stuff.
	runtime.LockOSThread()

	//Setup the unloadables so they finalize
	runtime.SetFinalizer(&unloadables, finalizeUnloadables)

	//Set the default trace callback. We are setting a default callback because the Load errors are fetched from the logs.
	SetTraceLogCallback(nil)

	//Reset the scale, this seems to be broken on mac?
	SetMouseScale(1, 1)
}

func ReportGlErrors() {
	C.reportGlErrors()
}
