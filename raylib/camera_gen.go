package raylib

/*
//Generated 2019-11-08T15:51:06+11:00
#include "raylib.h"
#include <stdlib.h>
#include "go.h"
*/
import "C"

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

//Update : Update camera position for selected mode
func (camera *Camera) Update() {
	ccamera := camera.cptr()
	C.UpdateCamera(ccamera)
}

//UpdateCamera : Update camera position for selected mode
//Recommended to use camera.Update() instead
func UpdateCamera(camera *Camera) {
	camera.Update()
}

//SetCameraPanControl : Set camera pan key to combine with mouse movement (free camera)
func SetCameraPanControl(panKey Key) {
	C.SetCameraPanControl(C.int(panKey))
}

//SetCameraAltControl : Set camera alt key to combine with mouse movement (free camera)
func SetCameraAltControl(altKey Key) {
	C.SetCameraAltControl(C.int(altKey))
}

//SetCameraSmoothZoomControl : Set camera smooth zoom key to combine with mouse (free camera)
func SetCameraSmoothZoomControl(szKey Key) {
	C.SetCameraSmoothZoomControl(C.int(szKey))
}

//SetCameraMoveControls : Set camera move controls (1st person and 3rd person cameras)
func SetCameraMoveControls(frontKey Key, backKey Key, rightKey Key, leftKey Key, upKey Key, downKey Key) {
	C.SetCameraMoveControls(C.int(frontKey), C.int(backKey), C.int(rightKey), C.int(leftKey), C.int(upKey), C.int(downKey))
}
