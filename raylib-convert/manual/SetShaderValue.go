//SetValueFloat32 : Set shader uniform value
func (shader *Shader) SetValueFloat32(uniformLoc int, value []float32, uniformType ShaderUniformDataType) {
	cshader := *shader.cptr()
	cvalue := (*C.float)(unsafe.Pointer((*reflect.SliceHeader)(unsafe.Pointer(&value)).Data))
	clen := C.int(1)
	C.SetShaderValueV(cshader, C.int(int32(uniformLoc)), unsafe.Pointer(cvalue), C.int(int32(uniformType)), clen)
}

//SetShaderValueFloat32 : Set shader uniform value
//Recommended to use shader.SetValueFloat32(uniformLoc, value, uniformType) instead
func SetShaderValueFloat32(shader *Shader, uniformLoc int, value []float32, uniformType ShaderUniformDataType) {
	shader.SetValueFloat32(uniformLoc, value, uniformType)
}

//SetValueInt32 : Set shader uniform value
func (shader *Shader) SetValueInt32(uniformLoc int, value []int32, uniformType ShaderUniformDataType) {
	cshader := *shader.cptr()
	cvalue := (*C.int)(unsafe.Pointer((*reflect.SliceHeader)(unsafe.Pointer(&value)).Data))
	clen := C.int(1)
	C.SetShaderValueV(cshader, C.int(int32(uniformLoc)), unsafe.Pointer(cvalue), C.int(int32(uniformType)), clen)
}

//SetShaderValueInt32 : Set shader uniform value
//Recommended to use shader.SetValueInt32(uniformLoc, value, uniformType) instead
func SetShaderValueInt32(shader *Shader, uniformLoc int, value []int32, uniformType ShaderUniformDataType) {
	shader.SetValueInt32(uniformLoc, value, uniformType)
}