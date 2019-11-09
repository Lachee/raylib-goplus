package raylib

//#include "raylib.h"
import "C"
import "unsafe"

type Ray struct {
	Position  Vector3
	Direction Vector3
}

func NewRay(position, direction Vector3) Ray {
	return Ray{Position: position, Direction: direction}
}

func newRayFromPointer(ptr unsafe.Pointer) Ray { return *(*Ray)(ptr) }
func (ray *Ray) cptr() *C.Ray {
	return (*C.Ray)(unsafe.Pointer(ray))
}

type RayHitInfo struct {
	Hit      bool
	Distance float32
	Position Vector3
	Normal   Vector3
}

func newRayHitInfoFromPointer(ptr unsafe.Pointer) RayHitInfo { return *(*RayHitInfo)(ptr) }
