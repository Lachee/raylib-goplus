//ExtractPalette : Extract color palette from image to maximum size
func (image *Image) ExtractPalette(maxPaletteSize int) ([]Color, int) {
	cextractCount := C.int(0)
	cimage := *image.cptr()
	res := C.ImageExtractPalette(cimage, C.int(int32(maxPaletteSize)), &cextractCount)
	defer C.free(unsafe.Pointer(res))
	
	//Calculate the samples
	samples := maxPaletteSize
	
	//Get the slice
	tmpslice := (*[1 << 24]Color)(unsafe.Pointer(res))[0:maxPaletteSize]	

	//Convert to a Color array
	goslice := make([]Color, samples)
	for i, _ := range tmpslice {
		//goslice[i] = newColorFromPointer(unsafe.Pointer(s))
		goslice[i] = Color{tmpslice[i].R, tmpslice[i].G, tmpslice[i].B, tmpslice[i].A }
	}

	
	return goslice, int(int32(cextractCount))
}

//ImageExtractPalette : Extract color palette from image to maximum size (memory should be freed)
//Recommended to use image.ExtractPalette(maxPaletteSize) instead
func ImageExtractPalette(image *Image, maxPaletteSize int) ([]Color, int) {
	return image.ExtractPalette(maxPaletteSize)
}