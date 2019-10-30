package ray

import "math"

const (
	PI      float32 = 3.14159265358979323846
	Deg2Rad float32 = PI / 180
	Rad2Deg float32 = 180 / PI
)

type Vector interface {
	Decompose() []float32
}

type Vector2 struct {
	X float32
	Y float32
}

func (v1 Vector2) Decompose() []float32 { return []float32{v1.X, v1.Y} }

func (v1 Vector2) Add(v2 Vector2) Vector2 {
	return Vector2{X: v1.X + v2.X, Y: v1.Y + v2.Y}
}
func (v1 Vector2) Subtract(v2 Vector2) Vector2 {
	return Vector2{X: v1.X - v2.X, Y: v1.Y - v2.Y}
}
func (v Vector2) Length() float32 {
	return float32(math.Sqrt(float64(v.X*v.X) + float64(v.Y*v.Y)))
}
func (v Vector2) DotProduct(v2 Vector2) float32 {
	return v.X*v2.X + v.Y*v2.Y
}
func (v Vector2) Angle(v2 Vector2) float32 {
	angle := float32(math.Atan2(float64(v2.Y-v.Y), float64(v2.X-v.X)) * (180.0 / float64(PI)))

	if angle < 0 {
		angle += 360.0
	}

	return angle
}
func (v Vector2) Scale(scale float32) Vector2 {
	return Vector2{X: v.X * scale, Y: v.Y * scale}
}
func (v Vector2) Multiply(v2 Vector2) Vector2 {
	return Vector2{X: v.X * v2.Y, Y: v.Y * v2.Y}
}
func (v Vector2) Negate() Vector2 {
	return Vector2{X: -v.X, Y: -v.Y}
}
func (v Vector2) Divide(d float32) Vector2 {
	return Vector2{X: v.X / d, Y: v.Y / d}
}
func (v Vector2) DivideV(v2 Vector2) Vector2 {
	return Vector2{X: v.X / v2.Y, Y: v.Y / v2.Y}
}
func (v Vector2) Normalize() Vector2 {
	return v.Divide(v.Length())
}
func (v Vector2) Lerp(target Vector2, amount float32) Vector2 {
	return Vector2{
		X: v.X + amount*(target.X-v.X),
		Y: v.Y + amount*(target.Y-v.Y),
	}
}

type Vector3 struct {
	X float32
	Y float32
	Z float32
}

type Vector4 struct {
	X float32
	Y float32
	Z float32
	W float32
}

type Quaternion Vector4

type Matrix struct {
	m0  float32
	m1  float32
	m2  float32
	m3  float32
	m4  float32
	m5  float32
	m6  float32
	m7  float32
	m8  float32
	m9  float32
	m10 float32
	m11 float32
	m12 float32
	m13 float32
	m14 float32
	m15 float32
}

func main() {
	v1.Add(v1)
}
