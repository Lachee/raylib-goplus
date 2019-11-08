//SetValueFloat32V : Set shader uniform value vector
func (shader *Shader) SetValueFloat32V(uniformLoc int, values []float32, uniformType ShaderUniformDataType) {
	cshader := *shader.cptr()
	C.SetShaderValueV(cshader, C.int(int32(uniformLoc)), unsafe.Pointer(&values[0]), C.int(int32(uniformType)), C.int(int32(len(values))))
}

//SetShaderValueFloat32V : Set shader uniform value vector
//Recommended to use shader.SetValueFloat32V(uniformLoc, value, uniformType) instead
func SetShaderValueFloat32V(shader *Shader, uniformLoc int, values []float32, uniformType ShaderUniformDataType) {
	shader.SetValueFloat32V(uniformLoc, values, uniformType)
}

//SetValueInt32V : Set shader uniform value vector
func (shader *Shader) SetValueInt32V(uniformLoc int, values []int32, uniformType ShaderUniformDataType) {
	cshader := *shader.cptr()
	C.SetShaderValueV(cshader, C.int(int32(uniformLoc)), unsafe.Pointer(&values[0]), C.int(int32(uniformType)), C.int(int32(len(values))))
}

//SetShaderValueInt32V : Set shader uniform value vector
//Recommended to use shader.SetValueInt32V(uniformLoc, value, uniformType) instead
func SetShaderValueInt32V(shader *Shader, uniformLoc int, values []int32, uniformType ShaderUniformDataType) {
	shader.SetValueInt32V(uniformLoc, values, uniformType)
}