//GetLocation : Get shader uniform location
func (shader Shader) GetLocation(uniformName string) int32 {
	cuniformName := C.CString(uniformName)
	defer C.free(unsafe.Pointer(cuniformName))
	cshader := *shader.cptr()
	res := C.GetShaderLocation(cshader, cuniformName)
	return int32(res)
}