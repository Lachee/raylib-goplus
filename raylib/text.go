package raylib

/*
#include "raylib.h"
#include <stdlib.h>
*/
import "C"
import "unsafe"

//DrawFPS :  Shows current FPS
func DrawFPS(x, y int) {
	C.DrawFPS(C.int(x), C.int(y))
}

//DrawText :  Draw text (using default font)
func DrawText(label string, x int, y int, w int, color Color) {
	ctext := C.CString(label)
	clr := *color.cptr()
	cx := C.int(x)
	cy := C.int(y)
	cw := C.int(w)

	defer C.free(unsafe.Pointer(ctext))
	C.DrawText(ctext, cx, cy, cw, clr)
}
