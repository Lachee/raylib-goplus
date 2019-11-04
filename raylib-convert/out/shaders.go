package raylib

//LoadShader : Load shader from files and bind default locations
func LoadShader(vsFileName string, fsFileName string) (Shader, string, string) {
	cfsFileName := C.CString(fsFileName)
	defer C.free(unsafe.Pointer(cfsFileName))
	cvsFileName := C.CString(vsFileName)
	defer C.free(unsafe.Pointer(cvsFileName))
	res := C.LoadShader(cvsFileName, cfsFileName)
	return newShaderFromPointer(unsafe.Pointer(&res)), C.GoString(cvsFileName), C.GoString(cfsFileName)
}

//LoadShaderCode : Load shader from code strings and bind default locations
func LoadShaderCode(vsCode string, fsCode string) (Shader, string, string) {
	cfsCode := C.CString(fsCode)
	defer C.free(unsafe.Pointer(cfsCode))
	cvsCode := C.CString(vsCode)
	defer C.free(unsafe.Pointer(cvsCode))
	res := C.LoadShaderCode(cvsCode, cfsCode)
	return newShaderFromPointer(unsafe.Pointer(&res)), C.GoString(cvsCode), C.GoString(cfsCode)
}

//UnloadShader : Unload shader from GPU memory (VRAM)
func UnloadShader(shader Shader) {
	cshader := *shader.cptr()
	C.UnloadShader(cshader)
}

//GetShaderDefault : Get default shader
func GetShaderDefault() Shader {
	res := C.GetShaderDefault()
	return newShaderFromPointer(unsafe.Pointer(&res))
}

//GetTextureDefault : Get default texture
func GetTextureDefault() Texture2D {
	res := C.GetTextureDefault()
	return newTexture2DFromPointer(unsafe.Pointer(&res))
}

//GetShaderLocation : Get shader uniform location
func GetShaderLocation(shader Shader, uniformName string) (int, string) {
	cuniformName := C.CString(uniformName)
	defer C.free(unsafe.Pointer(cuniformName))
	cshader := *shader.cptr()
	res := C.GetShaderLocation(cshader, cuniformName)
	return int(res), C.GoString(cuniformName)
}

//SetShaderValue : Set shader uniform value
func SetShaderValue(shader Shader, uniformLoc int, value unsafe.Pointer, uniformType int) {
	cshader := *shader.cptr()
	C.SetShaderValue(cshader, C.int(uniformLoc), value, C.int(uniformType))
}

//SetShaderValueV : Set shader uniform value vector
func SetShaderValueV(shader Shader, uniformLoc int, value unsafe.Pointer, uniformType int, count int) {
	cshader := *shader.cptr()
	C.SetShaderValueV(cshader, C.int(uniformLoc), value, C.int(uniformType), C.int(count))
}

//SetShaderValueMatrix : Set shader uniform value (matrix 4x4)
func SetShaderValueMatrix(shader Shader, uniformLoc int, mat Matrix) {
	cmat := *mat.cptr()
	cshader := *shader.cptr()
	C.SetShaderValueMatrix(cshader, C.int(uniformLoc), cmat)
}

//SetShaderValueTexture : Set shader uniform value for texture
func SetShaderValueTexture(shader Shader, uniformLoc int, texture Texture2D) {
	ctexture := *texture.cptr()
	cshader := *shader.cptr()
	C.SetShaderValueTexture(cshader, C.int(uniformLoc), ctexture)
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
func GenTextureCubemap(shader Shader, skyHDR Texture2D, size int) Texture2D {
	cskyHDR := *skyHDR.cptr()
	cshader := *shader.cptr()
	res := C.GenTextureCubemap(cshader, cskyHDR, C.int(size))
	return newTexture2DFromPointer(unsafe.Pointer(&res))
}

//GenTextureIrradiance : Generate irradiance texture using cubemap data
func GenTextureIrradiance(shader Shader, cubemap Texture2D, size int) Texture2D {
	ccubemap := *cubemap.cptr()
	cshader := *shader.cptr()
	res := C.GenTextureIrradiance(cshader, ccubemap, C.int(size))
	return newTexture2DFromPointer(unsafe.Pointer(&res))
}

//GenTexturePrefilter : Generate prefilter texture using cubemap data
func GenTexturePrefilter(shader Shader, cubemap Texture2D, size int) Texture2D {
	ccubemap := *cubemap.cptr()
	cshader := *shader.cptr()
	res := C.GenTexturePrefilter(cshader, ccubemap, C.int(size))
	return newTexture2DFromPointer(unsafe.Pointer(&res))
}

//GenTextureBRDF : Generate BRDF texture
func GenTextureBRDF(shader Shader, size int) Texture2D {
	cshader := *shader.cptr()
	res := C.GenTextureBRDF(cshader, C.int(size))
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
func BeginBlendMode(mode int) {
	C.BeginBlendMode(C.int(mode))
}

//EndBlendMode : End blending mode (reset to default: alpha blending)
func EndBlendMode() {
	C.EndBlendMode()
}

//InitVrSimulator : Init VR simulator for selected device parameters
func InitVrSimulator() {
	C.InitVrSimulator()
}

//CloseVrSimulator : Close VR simulator for current device
func CloseVrSimulator() {
	C.CloseVrSimulator()
}

//UpdateVrTracking : Update VR tracking (position and orientation) and camera
func UpdateVrTracking(camera Camera) Camera {
	ccamera := camera.cptr()
	C.UpdateVrTracking(&ccamera)
	return newCameraFromPointer(unsafe.Pointer(&ccamera))
}

//SetVrConfiguration : Set stereo rendering configuration parameters
func SetVrConfiguration(info VrDeviceInfo, distortion Shader) {
	cdistortion := *distortion.cptr()
	cinfo := *info.cptr()
	C.SetVrConfiguration(cinfo, cdistortion)
}

//IsVrSimulatorReady : Detect if VR simulator is ready
func IsVrSimulatorReady() bool {
	res := C.IsVrSimulatorReady()
	return bool(res)
}

//ToggleVrMode : Enable/Disable VR experience
func ToggleVrMode() {
	C.ToggleVrMode()
}

//BeginVrDrawing : Begin VR simulator stereo rendering
func BeginVrDrawing() {
	C.BeginVrDrawing()
}

//EndVrDrawing : End VR simulator stereo rendering
func EndVrDrawing() {
	C.EndVrDrawing()
}
