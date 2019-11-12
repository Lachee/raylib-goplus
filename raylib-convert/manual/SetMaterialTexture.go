//SetTexture : Set texture for a material map type (MAP_DIFFUSE, MAP_SPECULAR...)
func (material *Material) SetTexture(mapType MaterialMapType, texture Texture2D) {
	ctexture := *texture.cptr()
	cmaterial := material.cptr()
	C.SetMaterialTexture(cmaterial, C.int(int32(mapType)), ctexture)
	material.Maps[int(mapType)].Texture = texture
}

//SetMaterialTexture : Set texture for a material map type (MAP_DIFFUSE, MAP_SPECULAR...)
//Recommended to use material.SetTexture(mapType, texture) instead
func SetMaterialTexture(material *Material, mapType MaterialMapType, texture Texture2D) {
	material.SetTexture(mapType, texture)
}