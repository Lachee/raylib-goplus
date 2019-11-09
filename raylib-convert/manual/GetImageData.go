// GetImageData : Get pixel data from image as a Color slice
// Recommended to use image.GetPixels instead.
func GetImageData(image *Image) []Color {
	return image.GetPixels()
}

//GetPixels returns the pixel data from an image as a colour slice.
func (image *Image) GetPixels() []Color {
	//Prepare the pointer and get the data
	cimg := image.cptr()
	res := C.GetImageData(*cimg)
	defer C.free(unsafe.Pointer(res))

	//Calculate the samples
	samples := image.Width*image.Height
	
	//Get the slice
	tmpslice := (*[1 << 24]*C.Color)(unsafe.Pointer(res))[:samples:samples]

	//Convert to a Color array
	goslice := make([]Color, samples)
	for i, s := range tmpslice {
		goslice[i] = newColorFromPointer(unsafe.Pointer(s))
	}

	return goslice
}