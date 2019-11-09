package raylib

/*
Handles the translation between go structures and their C equivelence.
All moved to a monolithic file instead of their respective classes to limit the
  usage of import "C" and import "unsafe"
*/

//#include "raylib.h"
//#include <stdlib.h>
import "C"
import "unsafe"

func newVector2FromPointer(ptr unsafe.Pointer) Vector2 { return *(*Vector2)(ptr) }
func (v *Vector2) cptr() *C.Vector2 {
	return (*C.Vector2)(unsafe.Pointer(v))
}
func newVector3FromPointer(ptr unsafe.Pointer) Vector3 { return *(*Vector3)(ptr) }
func (v *Vector3) cptr() *C.Vector3 {
	return (*C.Vector3)(unsafe.Pointer(v))
}
func newVector4FromPointer(ptr unsafe.Pointer) Vector4 { return *(*Vector4)(ptr) }
func (v *Vector4) cptr() *C.Vector4 {
	return (*C.Vector4)(unsafe.Pointer(v))
}

func newTransformFromPointer(ptr unsafe.Pointer) Transform { return *(*Transform)(ptr) }
func (v *Transform) cptr() *C.Transform {
	return (*C.Transform)(unsafe.Pointer(v))
}

func (q *Quaternion) cptr() *C.Quaternion {
	return (*C.Quaternion)(unsafe.Pointer(q))
}
func (color *Color) cptr() *C.Color {
	return (*C.Color)(unsafe.Pointer(color))
}
func newRectangleFromPointer(ptr unsafe.Pointer) Rectangle { return *(*Rectangle)(ptr) }
func (r *Rectangle) cptr() *C.Rectangle {
	return (*C.Rectangle)(unsafe.Pointer(r))
}
