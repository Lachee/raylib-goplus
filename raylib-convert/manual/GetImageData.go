// GetImageData : Get pixel data from image as a Color slice
// Recommended to use image.GetPixels instead.
func GetImageData(image *Image) []Color {
	return image.GetPixels()
}

//GetPixels returns the pixel data from an image as a colour slice.
func (image *Image) GetPixels() []Color {

	//https://github.com/gen2brain/raylib-go/blob/48689803b8b82802b2899cc62afcb473d7ba8c00/raylib/textures.go#L131
	cimg := image.cptr()
	ret := C.GetImageData(*cimg)
	defer C.free(unsafe.Pointer(ret))
	
	tmpslice := (*[1 << 24]Color)(unsafe.Pointer(ret))[0:image.Width*image.Height]	
	goslice := make([]Color, image.Width*image.Height)
	
	for i, _ := range tmpslice {
		goslice[i] = Color{tmpslice[i].R, tmpslice[i].G, tmpslice[i].B, tmpslice[i].A }
	}
	
	return goslice
	
/*
	//Prepare the pointer and get the data
	cimg := image.cptr()
	res := C.GetImageData(*cimg)
	//defer C.free(unsafe.Pointer(res))

	//Calculate the samples
	samples := image.Width*image.Height
	
	//Get the slice
	tmpslice := (*[1 << 24]C.Color)(unsafe.Pointer(res))[:samples:samples]

	//Convert to a Color array
	goslice := make([]Color, samples)
	for i, s := range tmpslice {
		goslice[i] = newColorFromPointer(unsafe.Pointer(&s))
	}

	return goslice
	*/
}