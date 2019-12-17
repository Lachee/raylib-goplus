package raylib

/*
//Generated 2019-12-17T17:11:45+11:00
#include "raylib.h"
#include <stdlib.h>
#include "go.h"
*/
import "C"
import (
	"math"
	"unsafe"
)

//DrawPixel : Draw a pixel
func DrawPixel(posX int, posY int, color Color) {
	ccolor := *color.cptr()
	C.DrawPixel(C.int(int32(posX)), C.int(int32(posY)), ccolor)
}

//DrawPixelV : Draw a pixel (Vector version)
func DrawPixelV(position Vector2, color Color) {
	ccolor := *color.cptr()
	cposition := *position.cptr()
	C.DrawPixelV(cposition, ccolor)
}

//DrawLine : Draw a line
func DrawLine(startPosX int, startPosY int, endPosX int, endPosY int, color Color) {
	ccolor := *color.cptr()
	C.DrawLine(C.int(int32(startPosX)), C.int(int32(startPosY)), C.int(int32(endPosX)), C.int(int32(endPosY)), ccolor)
}

//DrawLineV : Draw a line (Vector version)
func DrawLineV(startPos Vector2, endPos Vector2, color Color) {
	ccolor := *color.cptr()
	cendPos := *endPos.cptr()
	cstartPos := *startPos.cptr()
	C.DrawLineV(cstartPos, cendPos, ccolor)
}

//DrawLineEx : Draw a line defining thickness
func DrawLineEx(startPos Vector2, endPos Vector2, thick float32, color Color) {
	ccolor := *color.cptr()
	cendPos := *endPos.cptr()
	cstartPos := *startPos.cptr()
	C.DrawLineEx(cstartPos, cendPos, C.float(thick), ccolor)
}

//DrawLineBezier : Draw a line using cubic-bezier curves in-out
func DrawLineBezier(startPos Vector2, endPos Vector2, thick float32, color Color) {
	ccolor := *color.cptr()
	cendPos := *endPos.cptr()
	cstartPos := *startPos.cptr()
	C.DrawLineBezier(cstartPos, cendPos, C.float(thick), ccolor)
}

//DrawLineStrip : Draw lines sequence
func DrawLineStrip(points Vector2, numPoints int, color Color) Vector2 {
	ccolor := *color.cptr()
	cpoints := points.cptr()
	C.DrawLineStrip(cpoints, C.int(int32(numPoints)), ccolor)
	return newVector2FromPointer(unsafe.Pointer(cpoints))
}

//DrawCircle : Draw a color-filled circle
func DrawCircle(centerX int, centerY int, radius float32, color Color) {
	ccolor := *color.cptr()
	C.DrawCircle(C.int(int32(centerX)), C.int(int32(centerY)), C.float(radius), ccolor)
}

//DrawCircleSector : Draw a piece of a circle
func DrawCircleSector(center Vector2, radius float32, startAngle int, endAngle int, segments int, color Color) {
	ccolor := *color.cptr()
	ccenter := *center.cptr()
	C.DrawCircleSector(ccenter, C.float(radius), C.int(int32(startAngle)), C.int(int32(endAngle)), C.int(int32(segments)), ccolor)
}

//DrawCircleSectorLines : Draw circle sector outline
func DrawCircleSectorLines(center Vector2, radius float32, startAngle int, endAngle int, segments int, color Color) {
	ccolor := *color.cptr()
	ccenter := *center.cptr()
	C.DrawCircleSectorLines(ccenter, C.float(radius), C.int(int32(startAngle)), C.int(int32(endAngle)), C.int(int32(segments)), ccolor)
}

//DrawCircleGradient : Draw a gradient-filled circle
func DrawCircleGradient(centerX int, centerY int, radius float32, color1 Color, color2 Color) {
	ccolor2 := *color2.cptr()
	ccolor1 := *color1.cptr()
	C.DrawCircleGradient(C.int(int32(centerX)), C.int(int32(centerY)), C.float(radius), ccolor1, ccolor2)
}

//DrawCircleV : Draw a color-filled circle (Vector version)
func DrawCircleV(center Vector2, radius float32, color Color) {
	ccolor := *color.cptr()
	ccenter := *center.cptr()
	C.DrawCircleV(ccenter, C.float(radius), ccolor)
}

//DrawCircleLines : Draw circle outline
func DrawCircleLines(centerX int, centerY int, radius float32, color Color) {
	ccolor := *color.cptr()
	C.DrawCircleLines(C.int(int32(centerX)), C.int(int32(centerY)), C.float(radius), ccolor)
}

//DrawRing : Draw ring
func DrawRing(center Vector2, innerRadius float32, outerRadius float32, startAngle int, endAngle int, segments int, color Color) {
	ccolor := *color.cptr()
	ccenter := *center.cptr()
	C.DrawRing(ccenter, C.float(innerRadius), C.float(outerRadius), C.int(int32(startAngle)), C.int(int32(endAngle)), C.int(int32(segments)), ccolor)
}

//DrawRingLines : Draw ring outline
func DrawRingLines(center Vector2, innerRadius float32, outerRadius float32, startAngle int, endAngle int, segments int, color Color) {
	ccolor := *color.cptr()
	ccenter := *center.cptr()
	C.DrawRingLines(ccenter, C.float(innerRadius), C.float(outerRadius), C.int(int32(startAngle)), C.int(int32(endAngle)), C.int(int32(segments)), ccolor)
}

//DrawRectangle : Draw a color-filled rectangle
func DrawRectangle(posX int, posY int, width int, height int, color Color) {
	ccolor := *color.cptr()
	C.DrawRectangle(C.int(int32(posX)), C.int(int32(posY)), C.int(int32(width)), C.int(int32(height)), ccolor)
}

//DrawRectangleV : Draw a color-filled rectangle (Vector version)
func DrawRectangleV(position Vector2, size Vector2, color Color) {
	ccolor := *color.cptr()
	csize := *size.cptr()
	cposition := *position.cptr()
	C.DrawRectangleV(cposition, csize, ccolor)
}

//DrawRectangleRec : Draw a color-filled rectangle
func DrawRectangleRec(rec Rectangle, color Color) {
	ccolor := *color.cptr()
	crec := *rec.cptr()
	C.DrawRectangleRec(crec, ccolor)
}

//DrawRectanglePro : Draw a color-filled rectangle with pro parameters
func DrawRectanglePro(rec Rectangle, origin Vector2, rotation float32, color Color) {
	ccolor := *color.cptr()
	corigin := *origin.cptr()
	crec := *rec.cptr()
	C.DrawRectanglePro(crec, corigin, C.float(rotation), ccolor)
}

//DrawRectangleGradientV : Draw a vertical-gradient-filled rectangle
func DrawRectangleGradientV(posX int, posY int, width int, height int, color1 Color, color2 Color) {
	ccolor2 := *color2.cptr()
	ccolor1 := *color1.cptr()
	C.DrawRectangleGradientV(C.int(int32(posX)), C.int(int32(posY)), C.int(int32(width)), C.int(int32(height)), ccolor1, ccolor2)
}

//DrawRectangleGradientH : Draw a horizontal-gradient-filled rectangle
func DrawRectangleGradientH(posX int, posY int, width int, height int, color1 Color, color2 Color) {
	ccolor2 := *color2.cptr()
	ccolor1 := *color1.cptr()
	C.DrawRectangleGradientH(C.int(int32(posX)), C.int(int32(posY)), C.int(int32(width)), C.int(int32(height)), ccolor1, ccolor2)
}

//DrawRectangleGradientEx : Draw a gradient-filled rectangle with custom vertex colors
func DrawRectangleGradientEx(rec Rectangle, col1 Color, col2 Color, col3 Color, col4 Color) {
	ccol4 := *col4.cptr()
	ccol3 := *col3.cptr()
	ccol2 := *col2.cptr()
	ccol1 := *col1.cptr()
	crec := *rec.cptr()
	C.DrawRectangleGradientEx(crec, ccol1, ccol2, ccol3, ccol4)
}

//DrawRectangleLines : Draw rectangle outline
func DrawRectangleLines(posX int, posY int, width int, height int, color Color) {
	ccolor := *color.cptr()
	C.DrawRectangleLines(C.int(int32(posX)), C.int(int32(posY)), C.int(int32(width)), C.int(int32(height)), ccolor)
}

//DrawRectangleLinesEx : Draw rectangle outline with extended parameters
func DrawRectangleLinesEx(rec Rectangle, lineThick int, color Color) {
	ccolor := *color.cptr()
	crec := *rec.cptr()
	C.DrawRectangleLinesEx(crec, C.int(int32(lineThick)), ccolor)
}

//DrawRectangleRounded : Draw rectangle with rounded edges
func DrawRectangleRounded(rec Rectangle, roundness float32, segments int, color Color) {
	ccolor := *color.cptr()
	crec := *rec.cptr()
	C.DrawRectangleRounded(crec, C.float(roundness), C.int(int32(segments)), ccolor)
}

//DrawRectangleRoundedLines : Draw rectangle with rounded edges outline
func DrawRectangleRoundedLines(rec Rectangle, roundness float32, segments int, lineThick int, color Color) {
	ccolor := *color.cptr()
	crec := *rec.cptr()
	C.DrawRectangleRoundedLines(crec, C.float(roundness), C.int(int32(segments)), C.int(int32(lineThick)), ccolor)
}

//DrawTriangle : Draw a color-filled triangle (vertex in counter-clockwise order!)
func DrawTriangle(v1 Vector2, v2 Vector2, v3 Vector2, color Color) {
	ccolor := *color.cptr()
	cv3 := *v3.cptr()
	cv2 := *v2.cptr()
	cv1 := *v1.cptr()
	C.DrawTriangle(cv1, cv2, cv3, ccolor)
}

//DrawTriangleLines : Draw triangle outline (vertex in counter-clockwise order!)
func DrawTriangleLines(v1 Vector2, v2 Vector2, v3 Vector2, color Color) {
	ccolor := *color.cptr()
	cv3 := *v3.cptr()
	cv2 := *v2.cptr()
	cv1 := *v1.cptr()
	C.DrawTriangleLines(cv1, cv2, cv3, ccolor)
}

//DrawTriangleFan : Draw a triangle fan defined by points (first vertex is the center)
func DrawTriangleFan(points Vector2, numPoints int, color Color) Vector2 {
	ccolor := *color.cptr()
	cpoints := points.cptr()
	C.DrawTriangleFan(cpoints, C.int(int32(numPoints)), ccolor)
	return newVector2FromPointer(unsafe.Pointer(cpoints))
}

//DrawTriangleStrip : Draw a triangle strip defined by points
func DrawTriangleStrip(points Vector2, pointsCount int, color Color) Vector2 {
	ccolor := *color.cptr()
	cpoints := points.cptr()
	C.DrawTriangleStrip(cpoints, C.int(int32(pointsCount)), ccolor)
	return newVector2FromPointer(unsafe.Pointer(cpoints))
}

//DrawPoly : Draw a regular polygon (Vector version)
func DrawPoly(center Vector2, sides int, radius float32, rotation float32, color Color) {
	ccolor := *color.cptr()
	ccenter := *center.cptr()
	C.DrawPoly(ccenter, C.int(int32(sides)), C.float(radius), C.float(rotation), ccolor)
}

//SetShapesTexture : Define default texture used to draw shapes
func SetShapesTexture(texture Texture2D, source Rectangle) {
	csource := *source.cptr()
	ctexture := *texture.cptr()
	C.SetShapesTexture(ctexture, csource)
}

//CheckCollisionRecs : Check collision between two rectangles
// Alias of rec1.Overlaps(rect2) instead.
func CheckCollisionRecs(r Rectangle, rect Rectangle) bool {
	return (r.X < (rect.X+rect.Width) && (r.X+r.Width) > rect.X) && (r.Y < (rect.Y+rect.Height) && (r.Y+r.Height) > rect.Y)
}

//Overlaps checks if a rectangle overlaps another.
func (r Rectangle) Overlaps(rect Rectangle) bool {
	return CheckCollisionRecs(r, rect)
}

//CheckCollisionCircles : Check collision between two circles
func CheckCollisionCircles(center1 Vector2, radius1 float32, center2 Vector2, radius2 float32) bool {
	distance := center1.Distance(center2)
	return distance <= radius1+radius2
}

//CheckCollisionCircleRec : Check collision between circle and rectangle
func CheckCollisionCircleRec(center Vector2, radius float32, rec Rectangle) bool {
	recCenter := rec.Center()
	dx := float32(math.Abs(float64(center.X - recCenter.X)))
	dy := float32(math.Abs(float64(center.Y - recCenter.Y)))

	if dx > (rec.Width/2+radius) || dy > (rec.Height/2+radius) {
		return false
	}

	if dx <= rec.Width/2 || dy <= rec.Height/2 {
		return true
	}

	ddx := dx - rec.Width/2
	ddy := dy - rec.Height/2
	cornerDistanceSq := ddx*ddx + ddy*ddy
	return cornerDistanceSq <= radius*radius
}

//GetOverlapRec : Get collision rectangle for two rectangles collision
// Alias of GetCollisionRec
func (rec1 Rectangle) GetOverlapRec(rec2 Rectangle) Rectangle {
	return GetCollisionRec(rec1, rec2)
}

//GetCollisionRec : Get collision rectangle for two rectangles collision
func GetCollisionRec(rec1 Rectangle, rec2 Rectangle) Rectangle {
	retRec := Rectangle{X: 0, Y: 0, Width: 0, Height: 0}

	if CheckCollisionRecs(rec1, rec2) {
		dxx := float32(math.Abs(float64(rec1.X - rec2.X)))
		dyy := float32(math.Abs(float64(rec1.Y - rec2.Y)))

		if rec1.X <= rec2.X {
			if rec1.Y <= rec2.Y {
				retRec.X = rec2.X
				retRec.Y = rec2.Y
				retRec.Width = rec1.Width - dxx
				retRec.Height = rec1.Height - dyy
			} else {
				retRec.X = rec2.X
				retRec.Y = rec1.Y
				retRec.Width = rec1.Width - dxx
				retRec.Height = rec2.Height - dyy
			}
		} else {
			if rec1.Y <= rec2.Y {
				retRec.X = rec1.X
				retRec.Y = rec2.Y
				retRec.Width = rec2.Width - dxx
				retRec.Height = rec1.Height - dyy
			} else {
				retRec.X = rec1.X
				retRec.Y = rec1.Y
				retRec.Width = rec2.Width - dxx
				retRec.Height = rec2.Height - dyy
			}
		}
	}

	if rec1.Width > rec2.Width {
		if retRec.Width >= rec2.Width {
			retRec.Width = rec2.Width
		}
	} else {
		if retRec.Width >= rec1.Width {
			retRec.Width = rec1.Width
		}
	}

	if rec1.Height > rec2.Height {
		if retRec.Height >= rec2.Height {
			retRec.Height = rec2.Height
		}
	} else {
		if retRec.Height >= rec1.Height {
			retRec.Height = rec1.Height
		}
	}

	return retRec
}

//CheckCollisionPointRec : Check if point is inside rectangle
func CheckCollisionPointRec(point Vector2, r Rectangle) bool {
	return point.X >= r.X && point.X <= (r.X+r.Width) && point.Y >= r.Y && point.Y <= (r.Y+r.Height)
}

//Contains checks if the rectangle contains a point
func (r Rectangle) Contains(point Vector2) bool {
	return CheckCollisionPointRec(point, r)
}

//CheckCollisionPointCircle : Check if point is inside circle
func CheckCollisionPointCircle(point Vector2, center Vector2, radius float32) bool {
	return CheckCollisionCircles(point, 0, center, radius)
}

//CheckCollisionPointTriangle : Check if point is inside a triangle
func CheckCollisionPointTriangle(point Vector2, p1 Vector2, p2 Vector2, p3 Vector2) bool {
	cp3 := *p3.cptr()
	cp2 := *p2.cptr()
	cp1 := *p1.cptr()
	cpoint := *point.cptr()
	res := C.CheckCollisionPointTriangle(cpoint, cp1, cp2, cp3)
	return bool(res)
}
