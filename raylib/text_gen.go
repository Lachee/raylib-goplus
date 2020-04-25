package raylib

/*
//Generated 2020-04-25T14:27:29+10:00
#include "raylib.h"
#include <stdlib.h>
#include "go.h"
*/
import "C"
import "unsafe"

//GetFontDefault : Get the default Font
func GetFontDefault() *Font {
	res := C.GetFontDefault()
	retval := newFontFromPointer(unsafe.Pointer(&res))
	RegisterUnloadable(retval)
	return retval
}

//LoadFont : Load font from file into GPU memory (VRAM)
func LoadFont(fileName string) *Font {
	cfileName := C.CString(fileName)
	defer C.free(unsafe.Pointer(cfileName))
	res := C.LoadFont(cfileName)
	retval := newFontFromPointer(unsafe.Pointer(&res))
	RegisterUnloadable(retval)
	return retval
}

//LoadFontEx : Load font from file with extended parameters
func LoadFontEx(fileName string, fontSize int, charsCount int) (*Font, []int) {

	cCharsCount := C.int(int32(charsCount))
	cFontChars := C.malloc(C.size_t(cCharsCount) * C.size_t(unsafe.Sizeof(int32(0))))
	cfileName := C.CString(fileName)

	defer C.free(unsafe.Pointer(cfileName))
	defer C.free(unsafe.Pointer(cFontChars))

	//Read the font resource
	res := C.LoadFontEx(cfileName, C.int(int32(fontSize)), (*C.int)(unsafe.Pointer(&cFontChars)), cCharsCount)

	//Convert the font chars back to a go array
	tempFontChars := (*[1 << 29]C.int)(cFontChars)
	goFontChars := make([]int, charsCount)
	for i, value := range tempFontChars {
		goFontChars[i] = int(int32(value))
	}

	//Create the font object
	retval := newFontFromPointer(unsafe.Pointer(&res))
	RegisterUnloadable(retval)

	//Return the result
	return retval, goFontChars
}

//LoadFontFromImage : Load font from Image (XNA style)
func LoadFontFromImage(image *Image, key Color, firstChar int) *Font {
	ckey := *key.cptr()
	cimage := *image.cptr()
	res := C.LoadFontFromImage(cimage, ckey, C.int(int32(firstChar)))
	retval := newFontFromPointer(unsafe.Pointer(&res))
	RegisterUnloadable(retval)
	return retval
}

// LoadFontData : Load font data and copy into a new Go Slice. Original is then freed.
func LoadFontData(fileName string, fontSize, charsCount int, fontType FontType) []CharInfo {
	cfileName := C.CString(fileName)
	defer C.free(unsafe.Pointer(cfileName))

	res := C.LoadFontData(cfileName, C.int(fontSize), nil, C.int(charsCount), C.int(fontType))
	defer C.free(unsafe.Pointer(res))

	//Get the slice
	tmpslice := (*[1 << 24]*C.CharInfo)(unsafe.Pointer(res))[:charsCount:charsCount]

	//Convert to a Color array
	goslice := make([]CharInfo, charsCount)
	for i, s := range tmpslice {
		goslice[i] = newCharInfoFromPointer(unsafe.Pointer(s))
	}

	return goslice
}

//Unload : Unload Font from GPU memory (VRAM)
func (font *Font) Unload() {
	cfont := *font.cptr()
	C.UnloadFont(cfont)
	UnregisterUnloadable(font)
}

//UnloadFont : Unload Font from GPU memory (VRAM)
//Recommended to use font.Unload() instead
func UnloadFont(font *Font) {
	font.Unload()
}

//DrawFPS : Shows current FPS
func DrawFPS(posX int, posY int) {
	C.DrawFPS(C.int(int32(posX)), C.int(int32(posY)))
}

//DrawText : Draw text (using default font)
func DrawText(text string, posX int, posY int, fontSize int, color Color) {
	ccolor := *color.cptr()
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	C.DrawText(ctext, C.int(int32(posX)), C.int(int32(posY)), C.int(int32(fontSize)), ccolor)
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

//DrawTextCodepoint : Draw one character (codepoint)
func DrawTextCodepoint(font Font, codepoint int, position Vector2, scale float32, tint Color) {
	ctint := *tint.cptr()
	cposition := *position.cptr()
	cfont := *font.cptr()
	C.DrawTextCodepoint(cfont, C.int(int32(codepoint)), cposition, C.float(scale), ctint)
}

//MeasureText : Measure string width for default font
func MeasureText(text string, fontSize int) int {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	res := C.MeasureText(ctext, C.int(int32(fontSize)))
	return int(int32(res))
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
func GetGlyphIndex(font Font, codepoint int) int {
	cfont := *font.cptr()
	res := C.GetGlyphIndex(cfont, C.int(int32(codepoint)))
	return int(int32(res))
}
