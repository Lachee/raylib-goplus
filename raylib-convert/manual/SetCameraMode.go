//SetMode : Set camera mode (multiple camera modes available)
func (camera *Camera) SetMode(mode CameraMode) {
	ccamera := *camera.cptr()
	C.SetCameraMode(ccamera, C.int(mode))
}

//SetCameraMode : Set camera mode (multiple camera modes available)
//Recommended to use camera.SetMode(mode) instead
func SetCameraMode(camera *Camera, mode CameraMode) {
	camera.SetMode(mode)
}