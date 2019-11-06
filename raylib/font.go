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

func LoadFontData(fileName string, fontSize, charsCount int, fontType FontType) []CharInfo {
	cfileName := C.CString(fileName)
	defer C.free(unsafe.Pointer(cfileName))

	result := C.LoadFontData(cfileName, C.int(fontSize), nil, C.int(charsCount), C.int(fontType))
	defer C.free(unsafe.Pointer(result))

	//Load into a mega slice
	resultArray := (*[1 << 28]C.CharInfo)(unsafe.Pointer(result))[:charsCount:charsCount]

	//Copy the values over
	slice := make([]CharInfo, charsCount)
	for i := 0; i < charsCount; i++ {
		slice[i] = newCharInfoFromPointer(unsafe.Pointer(&resultArray[i]))
	}

	//Return the slice
	return slice
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

func newFontFromPointer(ptr unsafe.Pointer) Font {
	return *(*Font)(ptr)
}
