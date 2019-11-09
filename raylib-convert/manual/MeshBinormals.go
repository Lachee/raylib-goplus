//ComputeBinormals : Compute mesh binormals
func (mesh *Mesh) ComputeBinormals() {
	cmesh := mesh.cptr()
	C.MeshBinormals(cmesh)
}

//MeshBinormals : Compute mesh binormals
//Recommended to use mesh.ComputeBinormals() instead
func MeshBinormals(mesh *Mesh) {
	mesh.ComputeBinormals()
}