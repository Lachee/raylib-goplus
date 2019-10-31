package raylib

/*
#include "raylib.h"
#include <stdlib.h>
*/
import "C"
import "unsafe"

type Rectangle struct {
	X      float32
	Y      float32
	Width  float32
	Height float32
}

// cptr returns C pointer
func (r *Rectangle) cptr() *C.Rectangle {
	return (*C.Rectangle)(unsafe.Pointer(r))
}

func Button(rect Rectangle, label string) {
	//ctext := C.CString(label)
	//defer C.free(unsafe.Pointer(ctext))
	//C.GuiButton(*rect.cptr(), ctext)
}
