//FromImage : Create an image from another image piece
func (image *Image) FromImage(rec Rectangle) (*Image) {
	crec := *rec.cptr()
	cimage := *image.cptr()
	res := C.ImageFromImage(cimage, crec)
	v := newImageFromPointer(unsafe.Pointer(&res)) 
	addUnloadable(v)
	return v
}

//ImageFromImage : Create an image from another image piece
//Recommended to use image.(rec) instead
func ImageFromImage(image *Image, rec Rectangle) *Image {
	return image.FromImage(rec) 
}