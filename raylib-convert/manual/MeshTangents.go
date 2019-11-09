//ComputeTangents : Compute mesh tangents
func (mesh *Mesh) ComputeTangents() {
	cmesh := mesh.cptr()
	C.MeshTangents(cmesh)
}

//MeshTangents : Compute mesh tangents
//Recommended to use mesh.ComputeTangents() instead
func MeshTangents(mesh *Mesh) {
	mesh.ComputeTangents()
}