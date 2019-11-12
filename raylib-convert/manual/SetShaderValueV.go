//SetValueFloat32V : Sets a vector (array) of uniform values
func (shader *Shader) SetValueFloat32V(uniformLoc int, values []float32, uniformType ShaderUniformDataType) {
	cshader := *shader.cptr()
	cvalue := (*C.float)(unsafe.Pointer((*reflect.SliceHeader)(unsafe.Pointer(&values)).Data))
	clen := C.int(int32(len(values)))
	C.SetShaderValueV(cshader, C.int(int32(uniformLoc)), unsafe.Pointer(cvalue), C.int(int32(uniformType)), clen)
}

//SetShaderValueFloat32V : Sets a float vector (array) of uniform values
//Recommended to use shader.SetValueFloat32V(uniformLoc, value, uniformType) instead
func SetShaderValueFloat32V(shader *Shader, uniformLoc int, values []float32, uniformType ShaderUniformDataType) {
	shader.SetValueFloat32V(uniformLoc, values, uniformType)
}

//SetValueInt32V : Sets a integer vector (array) of uniform values
func (shader *Shader) SetValueInt32V(uniformLoc int, values []int32, uniformType ShaderUniformDataType) {
	cshader := *shader.cptr()
	cvalue := (*C.int)(unsafe.Pointer((*reflect.SliceHeader)(unsafe.Pointer(&values)).Data))
	clen := C.int(int32(len(values)))
	C.SetShaderValueV(cshader, C.int(int32(uniformLoc)), unsafe.Pointer(cvalue), C.int(int32(uniformType)), clen)
}

//SetShaderValueInt32V : Sets a vector (array) of uniform values
//Recommended to use shader.SetValueInt32V(uniformLoc, value, uniformType) instead
func SetShaderValueInt32V(shader *Shader, uniformLoc int, values []int32, uniformType ShaderUniformDataType) {
	shader.SetValueInt32V(uniformLoc, values, uniformType)
}