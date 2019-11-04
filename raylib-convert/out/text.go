package raylib

//GetFontDefault : Get the default Font
func GetFontDefault() Font {
	res := C.GetFontDefault()
	return newFontFromPointer(unsafe.Pointer(&res))
}

//LoadFont : Load font from file into GPU memory (VRAM)
func LoadFont(fileName string) (Font, string) {
	cfileName := C.CString(fileName)
	defer C.free(unsafe.Pointer(cfileName))
	res := C.LoadFont(cfileName)
	return newFontFromPointer(unsafe.Pointer(&res)), C.GoString(cfileName)
}

//LoadFontEx : Load font from file with extended parameters
func LoadFontEx(fileName string, fontSize int, fontChars int, charsCount int) (Font, string, int) {
	cfontChars := C.int(fontChars)
	cfileName := C.CString(fileName)
	defer C.free(unsafe.Pointer(cfileName))
	res := C.LoadFontEx(cfileName, C.int(fontSize), &cfontChars, C.int(charsCount))
	return newFontFromPointer(unsafe.Pointer(&res)), C.GoString(cfileName), int(cfontChars)
}

//LoadFontFromImage : Load font from Image (XNA style)
func LoadFontFromImage(image Image, key Color, firstChar int) Font {
	ckey := *key.cptr()
	cimage := *image.cptr()
	res := C.LoadFontFromImage(cimage, ckey, C.int(firstChar))
	return newFontFromPointer(unsafe.Pointer(&res))
}

//UnloadFont : Unload Font from GPU memory (VRAM)
func UnloadFont(font Font) {
	cfont := *font.cptr()
	C.UnloadFont(cfont)
}

//DrawFPS : Shows current FPS
func DrawFPS(posX int, posY int) {
	C.DrawFPS(C.int(posX), C.int(posY))
}

//DrawText : Draw text (using default font)
func DrawText(text string, posX int, posY int, fontSize int, color Color) string {
	ccolor := *color.cptr()
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	C.DrawText(ctext, C.int(posX), C.int(posY), C.int(fontSize), ccolor)
	return C.GoString(ctext)
}

//DrawTextEx : Draw text using font and additional parameters
func DrawTextEx(font Font, text string, position Vector2, fontSize float32, spacing float32, tint Color) string {
	ctint := *tint.cptr()
	cposition := *position.cptr()
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	cfont := *font.cptr()
	C.DrawTextEx(cfont, ctext, cposition, C.float(fontSize), C.float(spacing), ctint)
	return C.GoString(ctext)
}

//DrawTextRec : Draw text using font inside rectangle limits
func DrawTextRec(font Font, text string, rec Rectangle, fontSize float32, spacing float32, wordWrap bool, tint Color) string {
	ctint := *tint.cptr()
	crec := *rec.cptr()
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	cfont := *font.cptr()
	C.DrawTextRec(cfont, ctext, crec, C.float(fontSize), C.float(spacing), C.bool(wordWrap), ctint)
	return C.GoString(ctext)
}

//MeasureText : Measure string width for default font
func MeasureText(text string, fontSize int) (int, string) {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	res := C.MeasureText(ctext, C.int(fontSize))
	return int(res), C.GoString(ctext)
}

//MeasureTextEx : Measure string size for Font
func MeasureTextEx(font Font, text string, fontSize float32, spacing float32) (Vector2, string) {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	cfont := *font.cptr()
	res := C.MeasureTextEx(cfont, ctext, C.float(fontSize), C.float(spacing))
	return newVector2FromPointer(unsafe.Pointer(&res)), C.GoString(ctext)
}

//GetGlyphIndex : Get index position for a unicode character on font
func GetGlyphIndex(font Font, character int) int {
	cfont := *font.cptr()
	res := C.GetGlyphIndex(cfont, C.int(character))
	return int(res)
}

//TextIsEqual : Check if two text string are equal
func TextIsEqual(text1 string, text2 string) (bool, string, string) {
	ctext2 := C.CString(text2)
	defer C.free(unsafe.Pointer(ctext2))
	ctext1 := C.CString(text1)
	defer C.free(unsafe.Pointer(ctext1))
	res := C.TextIsEqual(ctext1, ctext2)
	return bool(res), C.GoString(ctext1), C.GoString(ctext2)
}

//TextLength : Get text length, checks for '\0' ending
func TextLength(text string) (int, string) {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	res := C.TextLength(ctext)
	return int(res), C.GoString(ctext)
}

//TextSubtext : Get a piece of a text string
func TextSubtext(text string, position int, length int) (string, string) {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	res := C.TextSubtext(ctext, C.int(position), C.int(length))
	return C.GoString(res), C.GoString(ctext)
}

//TextJoin :  Join text strings with delimiter
func TextJoin(char string, count int, delimiter string) {
	fmt.Println("test")
}

//TextAppend : Append text at specific position and move cursor!
func TextAppend(text string, append string, position int) (string, string, int) {
	cposition := C.int(position)
	cappend := C.CString(append)
	defer C.free(unsafe.Pointer(cappend))
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	C.TextAppend(ctext, cappend, &cposition)
	return C.GoString(ctext), C.GoString(cappend), int(cposition)
}

//TextFindIndex : Find first text occurrence within a string
func TextFindIndex(text string, find string) (int, string, string) {
	cfind := C.CString(find)
	defer C.free(unsafe.Pointer(cfind))
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	res := C.TextFindIndex(ctext, cfind)
	return int(res), C.GoString(ctext), C.GoString(cfind)
}

//TextToUpper : Get upper case version of provided string
func TextToUpper(text string) (string, string) {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	res := C.TextToUpper(ctext)
	return C.GoString(res), C.GoString(ctext)
}

//TextToLower : Get lower case version of provided string
func TextToLower(text string) (string, string) {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	res := C.TextToLower(ctext)
	return C.GoString(res), C.GoString(ctext)
}

//TextToPascal : Get Pascal case notation version of provided string
func TextToPascal(text string) (string, string) {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	res := C.TextToPascal(ctext)
	return C.GoString(res), C.GoString(ctext)
}

//TextToInteger : Get integer value from text (negative values not supported)
func TextToInteger(text string) (int, string) {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	res := C.TextToInteger(ctext)
	return int(res), C.GoString(ctext)
}

//GetCodepointsCount : Get total number of characters (codepoints) in a UTF8 encoded string
func GetCodepointsCount(text string) (int, string) {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	res := C.GetCodepointsCount(ctext)
	return int(res), C.GoString(ctext)
}

//GetNextCodepoint : Returns next codepoint in a UTF8 encoded string; 0x3f('?') is returned on failure
func GetNextCodepoint(text string, bytesProcessed int) (int, string, int) {
	cbytesProcessed := C.int(bytesProcessed)
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	res := C.GetNextCodepoint(ctext, &cbytesProcessed)
	return int(res), C.GoString(ctext), int(cbytesProcessed)
}

//CodepointToUtf8 : Encode codepoint into utf8 text (char array length returned as parameter)
func CodepointToUtf8(codepoint int, byteLength int) (string, int) {
	cbyteLength := C.int(byteLength)
	res := C.CodepointToUtf8(C.int(codepoint), &cbyteLength)
	return C.GoString(res), int(cbyteLength)
}
