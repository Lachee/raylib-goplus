package raylib

//#include "raylib.h"
//#include <stdlib.h>
import "C"
import "unsafe"

type CharInfo struct {
	Value    uint32
	OffsetX  uint32
	OffsetY  uint32
	AdvanceX uint32
	Image    Image
}

func newCharInfoFromPointer(ptr unsafe.Pointer) CharInfo {
	return *(*CharInfo)(ptr)
}

type Font struct {
	BaseSize  int32
	CharCount int32
	Texture   Texture2D
	Chars     []CharInfo
	Recs      *Rectangle
}

//FontType defines generation method of the font
type FontType int32

const (
	//FontTypeDefault Default font generation, anti-aliased
	FontTypeDefault FontType = iota
	//FontTypeBitmap Bitmap font generation, no anti-aliasing
	FontTypeBitmap
	//FontTypeSDF SDF font generation, requires external shader
	FontTypeSDF
)

func (f *Font) cptr() *C.Font {
	return (*C.Font)(unsafe.Pointer(f))
}

func newFontFromPointer(ptr unsafe.Pointer) *Font {
	return (*Font)(ptr)
}
