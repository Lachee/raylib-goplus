package raylib

/*
//Generated 2019-11-11T18:03:30+11:00
#include "raylib.h"
#include <stdlib.h>
#include "go.h"
*/
import "C"
import "unsafe"

/*
//The following function has been ommitted because it is only available in OpenGL 1.1 and otherwise useless
// It was not worth the effort to write functionality as raylib-convert is unable to parse this function (pointer return)

//LoadText load chars array from text file
func LoadText(fileName string) string { return "" }
*/
//LoadShader : Load shader from files and bind default locations
func LoadShader(vsFileName string, fsFileName string) *Shader {
	cfsFileName := C.CString(fsFileName)
	defer C.free(unsafe.Pointer(cfsFileName))
	cvsFileName := C.CString(vsFileName)
	defer C.free(unsafe.Pointer(cvsFileName))
	res := C.LoadShader(cvsFileName, cfsFileName)
	retval := newShaderFromPointer(unsafe.Pointer(&res))
	addUnloadable(retval)
	return retval
}

//LoadShaderCode : Load shader from code strings and bind default locations
func LoadShaderCode(vsCode string, fsCode string) *Shader {
	cfsCode := C.CString(fsCode)
	defer C.free(unsafe.Pointer(cfsCode))
	cvsCode := C.CString(vsCode)
	defer C.free(unsafe.Pointer(cvsCode))
	res := C.LoadShaderCode(cvsCode, cfsCode)
	retval := newShaderFromPointer(unsafe.Pointer(&res))
	addUnloadable(retval)
	return retval
}

//Unload : Unload shader from GPU memory (VRAM)
func (shader *Shader) Unload() {
	cshader := *shader.cptr()
	C.UnloadShader(cshader)
	removeUnloadable(shader)
}

//UnloadShader : Unload shader from GPU memory (VRAM)
//Recommended to use shader.Unload() instead
func UnloadShader(shader *Shader) {
	shader.Unload()
}

//GetShaderDefault : Get default shader
func GetShaderDefault() *Shader {
	res := C.GetShaderDefault()
	retval := newShaderFromPointer(unsafe.Pointer(&res))
	addUnloadable(retval)
	return retval
}

//GetTextureDefault : Get default texture
func GetTextureDefault() *Texture2D {
	res := C.GetTextureDefault()
	retval := newTexture2DFromPointer(unsafe.Pointer(&res))
	addUnloadable(retval)
	return retval
}

//GetLocation : Get shader uniform location
func (shader *Shader) GetLocation(uniformName string) int {
	cuniformName := C.CString(uniformName)
	defer C.free(unsafe.Pointer(cuniformName))
	cshader := *shader.cptr()
	res := C.GetShaderLocation(cshader, cuniformName)
	return int(int32(res))
}

//GetShaderLocation : Get shader uniform location
//Recommended to use shader.GetLocation(uniformName) instead
func GetShaderLocation(shader *Shader, uniformName string) int {
	return shader.GetLocation(uniformName)
}

//SetValueFloat32 : Set shader uniform value
func (shader *Shader) SetValueFloat32(uniformLoc int, value float32, uniformType ShaderUniformDataType) {
	shader.SetValueFloat32V(uniformLoc, []float32{value}, uniformType)
}

//SetShaderValueFloat32 : Set shader uniform value
//Recommended to use shader.SetValueFloat32(uniformLoc, value, uniformType) instead
func SetShaderValueFloat32(shader *Shader, uniformLoc int, value float32, uniformType ShaderUniformDataType) {
	shader.SetValueFloat32(uniformLoc, value, uniformType)
}

//SetValueInt32 : Set shader uniform value
func (shader *Shader) SetValueInt32(uniformLoc int, value int32, uniformType ShaderUniformDataType) {
	shader.SetValueInt32V(uniformLoc, []int32{value}, uniformType)
}

//SetShaderValueInt32 : Set shader uniform value
//Recommended to use shader.SetValueInt32(uniformLoc, value, uniformType) instead
func SetShaderValueInt32(shader *Shader, uniformLoc int, value int32, uniformType ShaderUniformDataType) {
	shader.SetValueInt32(uniformLoc, value, uniformType)
}

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

//SetValueMatrix : Set shader uniform value (matrix 4x4)
func (shader *Shader) SetValueMatrix(uniformLoc int, mat Matrix) {
	cmat := *mat.cptr()
	cshader := *shader.cptr()
	C.SetShaderValueMatrix(cshader, C.int(int32(uniformLoc)), cmat)
}

//SetShaderValueMatrix : Set shader uniform value (matrix 4x4)
//Recommended to use shader.SetValueMatrix(uniformLoc, mat) instead
func SetShaderValueMatrix(shader *Shader, uniformLoc int, mat Matrix) {
	shader.SetValueMatrix(uniformLoc, mat)
}

//SetValueTexture : Set shader uniform value for texture
func (shader *Shader) SetValueTexture(uniformLoc int, texture *Texture2D) {
	ctexture := *texture.cptr()
	cshader := *shader.cptr()
	C.SetShaderValueTexture(cshader, C.int(int32(uniformLoc)), ctexture)
}

//SetShaderValueTexture : Set shader uniform value for texture
//Recommended to use shader.SetValueTexture(uniformLoc, texture) instead
func SetShaderValueTexture(shader *Shader, uniformLoc int, texture *Texture2D) {
	shader.SetValueTexture(uniformLoc, texture)
}

//SetMatrixProjection : Set a custom projection matrix (replaces internal projection matrix)
func SetMatrixProjection(proj Matrix) {
	cproj := *proj.cptr()
	C.SetMatrixProjection(cproj)
}

//SetMatrixModelview : Set a custom modelview matrix (replaces internal modelview matrix)
func SetMatrixModelview(view Matrix) {
	cview := *view.cptr()
	C.SetMatrixModelview(cview)
}

//GetMatrixModelview : Get internal modelview matrix
func GetMatrixModelview() Matrix {
	res := C.GetMatrixModelview()
	return newMatrixFromPointer(unsafe.Pointer(&res))
}

//GetMatrixProjection : Get internal projection matrix
func GetMatrixProjection() Matrix {
	res := C.GetMatrixProjection()
	return newMatrixFromPointer(unsafe.Pointer(&res))
}

//GenTextureCubemap : Generate cubemap texture from HDR texture
func GenTextureCubemap(shader Shader, skyHDR Texture2D, size int) *Texture2D {
	cskyHDR := *skyHDR.cptr()
	cshader := *shader.cptr()
	res := C.GenTextureCubemap(cshader, cskyHDR, C.int(int32(size)))
	return newTexture2DFromPointer(unsafe.Pointer(&res))
}

//GenTextureIrradiance : Generate irradiance texture using cubemap data
func GenTextureIrradiance(shader Shader, cubemap Texture2D, size int) *Texture2D {
	ccubemap := *cubemap.cptr()
	cshader := *shader.cptr()
	res := C.GenTextureIrradiance(cshader, ccubemap, C.int(int32(size)))
	return newTexture2DFromPointer(unsafe.Pointer(&res))
}

//GenTexturePrefilter : Generate prefilter texture using cubemap data
func GenTexturePrefilter(shader Shader, cubemap Texture2D, size int) *Texture2D {
	ccubemap := *cubemap.cptr()
	cshader := *shader.cptr()
	res := C.GenTexturePrefilter(cshader, ccubemap, C.int(int32(size)))
	return newTexture2DFromPointer(unsafe.Pointer(&res))
}

//GenTextureBRDF : Generate BRDF texture
func GenTextureBRDF(shader Shader, size int) *Texture2D {
	cshader := *shader.cptr()
	res := C.GenTextureBRDF(cshader, C.int(int32(size)))
	return newTexture2DFromPointer(unsafe.Pointer(&res))
}

//BeginShaderMode : Begin custom shader drawing
func BeginShaderMode(shader Shader) {
	cshader := *shader.cptr()
	C.BeginShaderMode(cshader)
}

//EndShaderMode : End custom shader drawing (use default shader)
func EndShaderMode() {
	C.EndShaderMode()
}

//BeginBlendMode : Begin blending mode (alpha, additive, multiplied)
func BeginBlendMode(mode BlendMode) {
	C.BeginBlendMode(C.int(int32(mode)))
}

//EndBlendMode : End blending mode (reset to default: alpha blending)
func EndBlendMode() {
	C.EndBlendMode()
}
