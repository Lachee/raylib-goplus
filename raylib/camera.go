package raylib

/*
#include "raylib.h"
#include <stdlib.h>
*/
import "C"
import "unsafe"

type CameraType int32

const (
	CameraTypePerspective CameraType = iota
	CameraTypeOrthographic
)

type CameraMode int32

const (
	CameraCustom CameraMode = iota
	CameraFree
	CameraOrbital
	CameraFirstPerson
	CameraThirdPerson
)

//Camera3D defines a camera in 3D space
type Camera3D struct {
	Position Vector3
	Target   Vector3
	Up       Vector3
	FOVY     float32
	Type     CameraType
}

//Camera is a fallback, defaults to Camera3D
type Camera Camera3D

func newCamera3DFromPointer(ptr unsafe.Pointer) *Camera3D {
	return (*Camera3D)(ptr)
}

func newCameraFromPointer(ptr unsafe.Pointer) *Camera {
	return (*Camera)(ptr)
}

func (w *Camera3D) cptr() *C.Camera3D {
	return (*C.Camera3D)(unsafe.Pointer(w))
}

func (w *Camera) cptr() *C.Camera {
	return (*C.Camera)(unsafe.Pointer(w))
}

type Camera2D struct {
	Offset   Vector2
	Target   Vector2
	Rotation float32
	Zoom     float32
}

func newCamera2DFromPointer(ptr unsafe.Pointer) *Camera2D {
	return (*Camera2D)(ptr)
}

func (w *Camera2D) cptr() *C.Camera2D {
	return (*C.Camera2D)(unsafe.Pointer(w))
}
