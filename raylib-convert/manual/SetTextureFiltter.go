//SetTextureFilter : Set texture scaling filter mode
func (texture *Texture2D) SetFilter(filterMode TextureFilterMode) {
	ctexture := *texture.cptr()
	C.SetTextureFilter(ctexture, C.int(int32(filterMode)))
}

//SetTextureFilter : Set texture scaling filter mode
//Recommended to use texture.SetTextureFilter(filterMode) instead
func SetTextureFilter(texture *Texture2D, filterMode TextureFilterMode) {
	texture.SetFilter(filterMode)
}