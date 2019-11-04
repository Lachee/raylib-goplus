// LoadImageEx Load image data from Color array data (RGBA - 32bit)
func LoadImageEx(pixels []Color, width, height int32) *Image {
	cpixels := pixels[0].cptr()
	cwidth := (C.int)(width)
	cheight := (C.int)(height)
	ret := C.LoadImageEx(cpixels, cwidth, cheight) //We call the C function as raylib uses stb_image to load images.
	v := newImageFromPointer(unsafe.Pointer(&ret))
	return v
}