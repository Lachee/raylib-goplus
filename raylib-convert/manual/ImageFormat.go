//Format : Convert image data to desired format
func (image *Image) SetFormat(newFormat PixelFormat) {
	cimage := image.cptr()
	C.ImageFormat(cimage, C.int(newFormat))
}

//ImageFormat : Convert image data to desired format
//Recommended to use image.SetFormat(newFormat) instead
func ImageFormat(image *Image, newFormat PixelFormat) {
	image.SetFormat(newFormat)
}