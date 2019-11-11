//LoadTextureCubemap : Load cubemap from image, multiple image cubemap layouts supported
func LoadTextureCubemap(image *Image, layoutType CubemapLayoutType) *TextureCubemap {
	cimage := *image.cptr()
	res := C.LoadTextureCubemap(cimage, C.int(int32(layoutType)))
	retval := newTextureCubemapFromPointer(unsafe.Pointer(&res))
	RegisterUnloadable(retval)
	return retval
}