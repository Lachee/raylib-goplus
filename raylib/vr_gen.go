package raylib

/*
//Generated 2019-11-12T18:07:27+11:00
#include "raylib.h"
#include <stdlib.h>
#include "go.h"
*/
import "C"

//InitVrSimulator : Init VR simulator for selected device parameters
func InitVrSimulator() {
	C.InitVrSimulator()
}

//CloseVrSimulator : Close VR simulator for current device
func CloseVrSimulator() {
	C.CloseVrSimulator()
}

//UpdateVrTracking : Update VR tracking (position and orientation) and camera
func UpdateVrTracking(camera *Camera) {
	ccamera := camera.cptr()
	C.UpdateVrTracking(ccamera)
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
