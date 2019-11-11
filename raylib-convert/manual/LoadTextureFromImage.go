//LoadTextureFromImage : Load texture from image data
func LoadTextureFromImage(image *Image) *Texture2D {
	cimage := *image.cptr()
	res := C.LoadTextureFromImage(cimage)
	retval := newTexture2DFromPointer(unsafe.Pointer(&res))
	RegisterUnloadable(retval)
	return retval
}

//UnregisteredLoadTextureFromImage : Load texture from image data without registering it to the unloadables.
func UnregisteredLoadTextureFromImage(image *Image) *Texture2D {
	cimage := *image.cptr()
	res := C.LoadTextureFromImage(cimage)
	retval := newTexture2DFromPointer(unsafe.Pointer(&res))
	return retval
}

//UnregisteredUnloadTexture : Unloads a texture without unregistering it from the unloadables.
func UnregisteredUnloadTexture(texture *Texture2D) {
	ctexture := *texture.cptr()
	C.UnloadTexture(ctexture)
}