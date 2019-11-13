package raylib

//#include "raylib.h"
import "C"
import "unsafe"

/*
Rectangle Structure
author: Lachee
source: https://github.com/raysan5/raylib/blob/master/src/raylib.h
*/

type Rectangle struct {
	X      float32
	Y      float32
	Width  float32
	Height float32
}

//NewRectangle creates a new rect
func NewRectangle(x, y, width, height float32) Rectangle {
	return Rectangle{X: x, Y: y, Width: width, Height: height}
}

//NewRectangleFromVector4 creates a rectangle out of a vector4
func NewRectangleFromVector4(vector Vector4) Rectangle {
	return NewRectangle(vector.X, vector.Y, vector.Z, vector.W)
}

//NewRectangleFromPositionSize creates a rectangle out of a position and size
func NewRectangleFromPositionSize(position, size Vector2) Rectangle {
	return NewRectangle(position.X, position.Y, size.X, size.Y)
}

//ToVector4 creates a Vector4 out of the rectangle components
func (r Rectangle) ToVector4() Vector4 {
	return NewVector4(r.X, r.Y, r.Width, r.Height)
}

//Position gets the position of the rectangle. Alias of MinPosition().
func (r Rectangle) Position() Vector2 {
	return NewVector2(r.X, r.Y)
}

//SetPosition sets the position of the rectangle
func (r Rectangle) SetPosition(v Vector2) Rectangle {
	return NewRectangle(v.X, v.Y, r.Width, r.Height)
}

//Size gets the size of the rectangle
func (r Rectangle) Size() Vector2 {
	return NewVector2(r.Width, r.Height)
}

//SetSize sets the size of the rectangle
func (r Rectangle) SetSize(v Vector2) Rectangle {
	return NewRectangle(r.X, r.Y, v.X, v.Y)
}

//MinPosition gets the smallest position the rectangle can be. Alias of Position().
func (r Rectangle) MinPosition() Vector2 {
	return NewVector2(r.X, r.Y)
}

//Center gets the center position of the rectangle
func (r Rectangle) Center() Vector2 {
	return NewVector2(r.X+r.Width/2, r.Y+r.Height/2)
}

//MaxPosition gets the maximum position within the bounds
func (r Rectangle) MaxPosition() Vector2 {
	return NewVector2(r.X+r.Width, r.Y+r.Height)
}

//Move shifts the rectangle by a vector
func (r Rectangle) Move(v Vector2) Rectangle {
	return NewRectangle(r.X+v.X, r.Y+v.Y, r.Width, r.Height)
}

//Grow adds the vector to the width and height of the rectangle.
func (r Rectangle) Grow(v Vector2) Rectangle {
	return NewRectangle(r.X, r.Y, v.X+r.Width, v.Y+r.Height)
}

//Scale the width and height of the rectangle
func (r Rectangle) Scale(scale float32) Rectangle {
	return NewRectangle(r.X, r.Y, r.Width*scale, r.Height*scale)
}

//Lerp a rectangle to a target rectangle
func (r Rectangle) Lerp(target Rectangle, amount float32) Rectangle {
	return Rectangle{
		X:      r.X + amount*(target.X-r.X),
		Y:      r.Y + amount*(target.Y-r.Y),
		Width:  r.Width + amount*(target.Width-r.Width),
		Height: r.Height + amount*(target.Height-r.Height),
	}
}

//LerpPosition a rectangle to a target position
func (r Rectangle) LerpPosition(pos Vector2, amount float32) Rectangle {
	return Rectangle{
		X:      r.X + amount*(pos.X-r.X),
		Y:      r.Y + amount*(pos.Y-r.Y),
		Width:  r.Width,
		Height: r.Height,
	}
}

type BoundingBox struct {
	Min Vector3
	Max Vector3
}

// NewBoundingBox - Returns new BoundingBox
func NewBoundingBox(min, max Vector3) BoundingBox {
	return BoundingBox{min, max}
}

//Size gets the size of the bounding box
func (bb BoundingBox) Size() Vector3 {
	return bb.Max.Subtract(bb.Min)
}

//Center gets the center position of the bounding box
func (bb BoundingBox) Center() Vector3 {
	return bb.Min.Add(bb.Max.Subtract(bb.Min).Divide(2))
}

func newBoundingBoxFromPointer(ptr unsafe.Pointer) BoundingBox {
	return *(*BoundingBox)(ptr)
}
func (r *BoundingBox) cptr() *C.BoundingBox {
	return (*C.BoundingBox)(unsafe.Pointer(r))
}
