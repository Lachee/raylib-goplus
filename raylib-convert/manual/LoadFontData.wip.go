//LoadFontData is currently not supported
// LoadFontData : Load font data for further use
// NOTE: Requires TTF font and can generate SDF data
// It is REQUIRED to free the slice with UnloadFontData
//func LoadFontData(fileName string, fontSize, charsCount, fontType FontType) []CharInfo {
//	cfileName := C.CString(fileName)
//	cfontChars := C.int(fontChars)
//	defer C.free(unsafe.Pointer(cfilename))
//
//	result := C.LoadFontData(cfileName, C.int(fontSize), nil, C.int(charsCount), C.int(fontType))
//	slice := (*[1 << 30]CharInfo)(unsafe.Pointer(result))[:length:length]
//
//	return slice, int(cfontChars)
//}