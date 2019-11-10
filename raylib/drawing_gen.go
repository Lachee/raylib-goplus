package raylib

/*
//Generated 2019-11-10T17:59:53+11:00
#include "raylib.h"
#include <stdlib.h>
#include "go.h"
*/
import "C"

//DrawLine3D : Draw a line in 3D world space
func DrawLine3D(startPos Vector3, endPos Vector3, color Color) {
	ccolor := *color.cptr()
	cendPos := *endPos.cptr()
	cstartPos := *startPos.cptr()
	C.DrawLine3D(cstartPos, cendPos, ccolor)
}

//DrawCircle3D : Draw a circle in 3D world space
func DrawCircle3D(center Vector3, radius float32, rotationAxis Vector3, rotationAngle float32, color Color) {
	ccolor := *color.cptr()
	crotationAxis := *rotationAxis.cptr()
	ccenter := *center.cptr()
	C.DrawCircle3D(ccenter, C.float(radius), crotationAxis, C.float(rotationAngle), ccolor)
}

//DrawCube : Draw cube
func DrawCube(position Vector3, width float32, height float32, length float32, color Color) {
	ccolor := *color.cptr()
	cposition := *position.cptr()
	C.DrawCube(cposition, C.float(width), C.float(height), C.float(length), ccolor)
}

//DrawCubeV : Draw cube (Vector version)
func DrawCubeV(position Vector3, size Vector3, color Color) {
	ccolor := *color.cptr()
	csize := *size.cptr()
	cposition := *position.cptr()
	C.DrawCubeV(cposition, csize, ccolor)
}

//DrawCubeWires : Draw cube wires
func DrawCubeWires(position Vector3, width float32, height float32, length float32, color Color) {
	ccolor := *color.cptr()
	cposition := *position.cptr()
	C.DrawCubeWires(cposition, C.float(width), C.float(height), C.float(length), ccolor)
}

//DrawCubeWiresV : Draw cube wires (Vector version)
func DrawCubeWiresV(position Vector3, size Vector3, color Color) {
	ccolor := *color.cptr()
	csize := *size.cptr()
	cposition := *position.cptr()
	C.DrawCubeWiresV(cposition, csize, ccolor)
}

//DrawCubeTexture : Draw cube textured
func DrawCubeTexture(texture Texture2D, position Vector3, width float32, height float32, length float32, color Color) {
	ccolor := *color.cptr()
	cposition := *position.cptr()
	ctexture := *texture.cptr()
	C.DrawCubeTexture(ctexture, cposition, C.float(width), C.float(height), C.float(length), ccolor)
}

//DrawSphere : Draw sphere
func DrawSphere(centerPos Vector3, radius float32, color Color) {
	ccolor := *color.cptr()
	ccenterPos := *centerPos.cptr()
	C.DrawSphere(ccenterPos, C.float(radius), ccolor)
}

//DrawSphereEx : Draw sphere with extended parameters
func DrawSphereEx(centerPos Vector3, radius float32, rings int, slices int, color Color) {
	ccolor := *color.cptr()
	ccenterPos := *centerPos.cptr()
	C.DrawSphereEx(ccenterPos, C.float(radius), C.int(int32(rings)), C.int(int32(slices)), ccolor)
}

//DrawSphereWires : Draw sphere wires
func DrawSphereWires(centerPos Vector3, radius float32, rings int, slices int, color Color) {
	ccolor := *color.cptr()
	ccenterPos := *centerPos.cptr()
	C.DrawSphereWires(ccenterPos, C.float(radius), C.int(int32(rings)), C.int(int32(slices)), ccolor)
}

//DrawCylinder : Draw a cylinder/cone
func DrawCylinder(position Vector3, radiusTop float32, radiusBottom float32, height float32, slices int, color Color) {
	ccolor := *color.cptr()
	cposition := *position.cptr()
	C.DrawCylinder(cposition, C.float(radiusTop), C.float(radiusBottom), C.float(height), C.int(int32(slices)), ccolor)
}

//DrawCylinderWires : Draw a cylinder/cone wires
func DrawCylinderWires(position Vector3, radiusTop float32, radiusBottom float32, height float32, slices int, color Color) {
	ccolor := *color.cptr()
	cposition := *position.cptr()
	C.DrawCylinderWires(cposition, C.float(radiusTop), C.float(radiusBottom), C.float(height), C.int(int32(slices)), ccolor)
}

//DrawPlane : Draw a plane XZ
func DrawPlane(centerPos Vector3, size Vector2, color Color) {
	ccolor := *color.cptr()
	csize := *size.cptr()
	ccenterPos := *centerPos.cptr()
	C.DrawPlane(ccenterPos, csize, ccolor)
}

//DrawRay : Draw a ray line
func DrawRay(ray Ray, color Color) {
	ccolor := *color.cptr()
	cray := *ray.cptr()
	C.DrawRay(cray, ccolor)
}

//DrawGrid : Draw a grid (centered at (0, 0, 0))
func DrawGrid(slices int, spacing float32) {
	C.DrawGrid(C.int(int32(slices)), C.float(spacing))
}

//DrawGizmo : Draw simple gizmo
func DrawGizmo(position Vector3) {
	cposition := *position.cptr()
	C.DrawGizmo(cposition)
}
