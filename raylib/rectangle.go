package raylib

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
	return Rectange{X: x, Y: y, Width: width, Height: height}
}

//NewRectangleFromVector4 creates a rectangle out of a vector4
func NewRectangleFromVector4(vector Vector4) Rectangle {
	return NewRectangle(vector.X, vector.Y, vector.Z, vector.W)
}

//NewRectangleFromPositionSize creates a rectangle out of a position and size
func NewRectangleFromPositionSize(position, size Vector2) {
	return NewRectangle(position.X, position.Y, size.X, size.Y)
}

//ToVector4 creates a Vector4 out of the rectangle components
func (r Rectangle) ToVector4() Vector4 {
	return NewVector4(r.X, r.Y, r.Width, r.Height)
}

//Position gets the position of the rectangle
func (r Rectangle) Position() Vector2 {
	return NewVector2(r.X, r.Y)
}

//Size gets the size of the rectangle
func (r Rectangle) Size() Vector2 {
	return NewVector2(r.Width, r.Height)
}

//MinPosition gets the smallest position the rectangle can be. Alias of Position().
func (r Rectangle) MinPosition() Vector2 {
	return NewVector2(r.X, r.Y)
}

//MaxPosition gets the maximum position within the bounds
func (r Rectangle) MaxPosition() Vector2 {
	return NewVector2(r.X+r.Width, r.Y+r.Height)
}

//Contains checks if the rectangle contains a point
func (r Rectangle) Contains(point Vector2) bool {
	return point.X >= r.X && point.X <= (r.X+r.Width) && point.Y >= r.Y && point.Y <= (r.Y+r.Height)
}

//Overlaps checks if a rectangle overlaps another.
func (r Rectangle) Overlaps(rect Rectangle) bool {
  return (r.X < (rect.X + rect.Width) && (r.X + r.Width) > rect.X) &&  (r.Y < (rect.Y + rect.Height) && (r.Y + r.Height) > rect.Y))
}
