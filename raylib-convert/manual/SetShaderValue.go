//SetValueFloat32 : Set shader uniform value
func (shader *Shader) SetValueFloat32(uniformLoc int, value float32, uniformType ShaderUniformDataType) {
	shader.SetValueFloat32V(uniformLoc, []float32{ value }, uniformType)
}

//SetShaderValueFloat32 : Set shader uniform value
//Recommended to use shader.SetValueFloat32(uniformLoc, value, uniformType) instead
func SetShaderValueFloat32(shader *Shader, uniformLoc int, value float32, uniformType ShaderUniformDataType) {
	shader.SetValueFloat32(uniformLoc, value, uniformType)
}

//SetValueInt32 : Set shader uniform value
func (shader *Shader) SetValueInt32(uniformLoc int, value int32, uniformType ShaderUniformDataType) {
	shader.SetValueInt32V(uniformLoc, []int32{ value }, uniformType)
}

//SetShaderValueInt32 : Set shader uniform value
//Recommended to use shader.SetValueInt32(uniformLoc, value, uniformType) instead
func SetShaderValueInt32(shader *Shader, uniformLoc int, value int32, uniformType ShaderUniformDataType) {
	shader.SetValueInt32(uniformLoc, value, uniformType)
}