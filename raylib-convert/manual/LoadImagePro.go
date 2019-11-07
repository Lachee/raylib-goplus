//LoadImagePro loads raw data wtih parameters
func LoadImagePro(pixels []byte, width, height int32, format PixelFormat) *Image {
	data := unsafe.Pointer(&pixels[0])
	res := C.LoadImagePro(data, C.int(width), C.int(height), C.int(format))
	v := newImageFromPointer(unsafe.Pointer(&res))
	addUnloadable(v)
	return v
}