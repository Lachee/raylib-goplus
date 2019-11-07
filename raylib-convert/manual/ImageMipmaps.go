//CreateMipmaps : Generate all mipmap levels for a provided image
func (image *Image) CreateMipmaps() {
	cimage := image.cptr()
	C.ImageMipmaps(cimage)
}

//ImageMipmaps : Generate all mipmap levels for a provided image
//Recommended to use image.CreateMipmaps() instead
func ImageMipmaps(image *Image) {
	image.CreateMipmaps()
}