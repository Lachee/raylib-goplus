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

func (v *Vector2) cptr() *C.Vector2 {
	return (*C.Vector2)(unsafe.Pointer(v))
}
func (v *Vector3) cptr() *C.Vector3 {
	return (*C.Vector3)(unsafe.Pointer(v))
}
func (v *Vector4) cptr() *C.Vector4 {
	return (*C.Vector4)(unsafe.Pointer(v))
}
func (m *Matrix) cptr() *C.Matrix {
	return (*C.Matrix)(unsafe.Pointer(m))
}
func (q *Quaternion) cptr() *C.Quaternion {
	return (*C.Quaternion)(unsafe.Pointer(q))
}
func (color *Color) cptr() *C.Color {
	return (*C.Color)(unsafe.Pointer(color))
}
func (r *Rectangle) cptr() *C.Rectangle {
	return (*C.Rectangle)(unsafe.Pointer(r))
}
