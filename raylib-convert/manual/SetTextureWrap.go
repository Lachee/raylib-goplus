//SetWrap : Set texture wrapping mode
func (texture *Texture2D) SetWrap(wrapMode TextureWrapMode) {
	ctexture := *texture.cptr()
	C.SetTextureWrap(ctexture, C.int(int32(wrapMode)))
}

//SetTextureWrap : Set texture wrapping mode
//Recommended to use texture.SetWrap(wrapMode) instead
func SetTextureWrap(texture *Texture2D, wrapMode TextureWrapMode) {
	texture.SetWrap(wrapMode)
}