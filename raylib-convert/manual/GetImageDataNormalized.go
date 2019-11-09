// GetImageDataNormalized : Get pixel data from image as a Color slice
// Recommended to use image.GetPixelsNormalized instead.
func GetImageDataNormalized(image *Image) []Vector4 {
	return image.GetPixelsNormalized()
}

//GetPixelsNormalized returns the pixel data from an image as a colour slice.
func (image *Image) GetPixelsNormalized() []Vector4 {
	//Prepare the pointer and get the data
	cimg := image.cptr()
	res := C.GetImageDataNormalized(*cimg)
	defer C.free(unsafe.Pointer(res))

	//Calculate the samples
	samples := image.Width*image.Height
	
	//Get the slice
	tmpslice := (*[1 << 24]*C.Vector4)(unsafe.Pointer(res))[:samples:samples]

	//Convert to a Color array
	goslice := make([]Vector4, samples)
	for i, s := range tmpslice {
		goslice[i] = newVector4FromPointer(unsafe.Pointer(s))
	}

	return goslice
}