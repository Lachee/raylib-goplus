//LoadFontEx : Load font from file with extended parameters
func LoadFontEx(fileName string, fontSize int, charsCount int) (*Font, []int) {

	cCharsCount 	:= C.int(int32(charsCount))
	cFontChars 		:= C.malloc(C.size_t(cCharsCount) * C.size_t(unsafe.Sizeof(int32(0))))
	cfileName 		:= C.CString(fileName)
	
	defer C.free(unsafe.Pointer(cfileName))
	defer C.free(unsafe.Pointer(cFontChars))
	
	//Read the font resource
	res := C.LoadFontEx(cfileName, C.int(int32(fontSize)), (*C.int)(unsafe.Pointer(&cFontChars)), cCharsCount)
	
	//Convert the font chars back to a go array
	tempFontChars 	:= (*[1 << 29]C.int)(cFontChars)
	goFontChars 	:= make([]int, charsCount)
	for i, value := range tempFontChars {
		goFontChars[i] = int(int32(value))
	}
	
	//Create the font object
	retval := newFontFromPointer(unsafe.Pointer(&res))
	RegisterUnloadable(retval)
	
	//Return the result
	return retval, goFontChars
}
