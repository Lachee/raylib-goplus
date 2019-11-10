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