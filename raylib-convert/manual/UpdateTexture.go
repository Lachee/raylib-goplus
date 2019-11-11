//UpdateTexture : Update GPU texture with new data
func (texture *Texture2D) UpdateTexture(pixels []Color) {
	ctexture := *texture.cptr()
	cpixels := pixels[0].cptr()
	C.UpdateTexture(ctexture, unsafe.Pointer(cpixels))
}

//UpdateTexture : Update GPU texture with new data
//Recommended to use texture.UpdateTexture(pixels) instead
func UpdateTexture(texture *Texture2D, pixels []Color) {
	texture.UpdateTexture(pixels)
}