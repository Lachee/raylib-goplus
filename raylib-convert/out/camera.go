package raylib

//SetCameraMode : Set camera mode (multiple camera modes available)
func SetCameraMode(camera Camera, mode int) {
	ccamera := *camera.cptr()
	C.SetCameraMode(ccamera, C.int(mode))
}

//UpdateCamera : Update camera position for selected mode
func UpdateCamera(camera Camera) Camera {
	ccamera := camera.cptr()
	C.UpdateCamera(&ccamera)
	return newCameraFromPointer(unsafe.Pointer(&ccamera))
}

//SetCameraPanControl : Set camera pan key to combine with mouse movement (free camera)
func SetCameraPanControl(panKey int) {
	C.SetCameraPanControl(C.int(panKey))
}

//SetCameraAltControl : Set camera alt key to combine with mouse movement (free camera)
func SetCameraAltControl(altKey int) {
	C.SetCameraAltControl(C.int(altKey))
}

//SetCameraSmoothZoomControl : Set camera smooth zoom key to combine with mouse (free camera)
func SetCameraSmoothZoomControl(szKey int) {
	C.SetCameraSmoothZoomControl(C.int(szKey))
}

//SetCameraMoveControls : Set camera move controls (1st person and 3rd person cameras)
func SetCameraMoveControls(frontKey int, backKey int, rightKey int, leftKey int, upKey int, downKey int) {
	C.SetCameraMoveControls(C.int(frontKey), C.int(backKey), C.int(rightKey), C.int(leftKey), C.int(upKey), C.int(downKey))
}
