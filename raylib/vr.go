package raylib

//#include "raylib.h"
import "C"
import "unsafe"

type VrDeviceInfo struct {
	HorizontalResolution       int32
	VerticalResolution         int32
	HorizontalScreenSize       float32
	VerticalScreenSize         float32
	VerticalScreenCenter       float32
	EyeToScreenDistance        float32
	LensSeperationDistance     float32
	InterpupillaryDistance     float32
	LensDistortionValues       [4]float32
	ChromaAberrationCorrection [4]float32
}

func newVrDeviceInfoFromPointer(ptr unsafe.Pointer) *VrDeviceInfo {
	return (*VrDeviceInfo)(ptr)
}
func (vr *VrDeviceInfo) cptr() *C.VrDeviceInfo {
	return (*C.VrDeviceInfo)(unsafe.Pointer(vr))
}
