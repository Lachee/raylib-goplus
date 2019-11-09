//LoadModelAnimations : Load model animations from file
func LoadModelAnimations(fileName string) ([]ModelAnimation, int32) {
	cfileName := C.CString(fileName)
	ccount := C.int(0)
	defer C.free(unsafe.Pointer(cfileName))
	
	res := C.LoadModelAnimations(cfileName, &ccount)
	
	samples := int32(ccount)
	tmpslice := (*[1 << 24]*C.ModelAnimation)(unsafe.Pointer(res))[:samples:samples]
	
	goslice := make([]ModelAnimation, samples)
	for i, s := range tmpslice {
		goslice[i] = *newModelAnimationFromPointer(unsafe.Pointer(s))
	}

	return goslice, samples
}