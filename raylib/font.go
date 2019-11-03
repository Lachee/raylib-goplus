package raylib

//#include "raylib.h"
import "C"
import "unsafe"

type CharInfo struct {
	Value    uint32
	OffsetX  uint32
	OffsetY  uint32
	AdvanceX uint32
	Image    Image
}

type Font struct {
	BaseSize  int32
	CharCount int32
	Texture   Texture2D
	Chars     *CharInfo
	Recs      *Rectangle
}

func (f *Font) cptr() *C.Font {
	return (*C.Font)(unsafe.Pointer(f))
}

func newFontFromPointer(ptr unsafe.Pointer) Font {
	return *(*Font)(ptr)
}
