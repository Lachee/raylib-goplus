package ray

import "math"

const (
	PI      float32 = 3.14159265358979323846
	Deg2Rad float32 = PI / 180
	Rad2Deg float32 = 180 / PI
)

//Vector a interface for any vector
type Vector interface {
	Decompose() []float32
}

//Vector2 a 2 dimensional vector
type Vector2 struct {
	X float32
	Y float32
}

//Decompose the vector into a slice of floats
func (v Vector2) Decompose() []float32 { return []float32{v.X, v.Y} }

//Add two vectors (v1 + v2)
func (v Vector2) Add(v2 Vector2) Vector2 {
	return Vector2{X: v.X + v2.X, Y: v.Y + v2.Y}
}

//Subtract two vectors (v1 - v2)
func (v Vector2) Subtract(v2 Vector2) Vector2 {
	return Vector2{X: v.X - v2.X, Y: v.Y - v2.Y}
}

//Length of the vector
func (v Vector2) Length() float32 {
	return float32(math.Sqrt(float64(v.X*v.X) + float64(v.Y*v.Y)))
}

//SqrLength is the squared length of the vector
func (v Vector2) SqrLength() float32 {
	return float32(float64(v.X*v.X) + float64(v.Y*v.Y))
}

//DotProduct of the vector
func (v Vector2) DotProduct(v2 Vector2) float32 {
	return v.X*v2.X + v.Y*v2.Y
}

//Angle the vector creates with another vector
func (v Vector2) Angle(v2 Vector2) float32 {
	result := float32(math.Atan2(float64(v2.Y-v.Y), float64(v2.X-v.X))) * Rad2Deg
	if result < 0 {
		result += 360
	}
	return result
}

//Scale the vector (v * scale)
func (v Vector2) Scale(scale float32) Vector2 {
	return Vector2{X: v.X * scale, Y: v.Y * scale}
}

//Multiply a vector by another vector
func (v Vector2) Multiply(v2 Vector2) Vector2 {
	return Vector2{X: v.X * v2.Y, Y: v.Y * v2.Y}
}

//Negate or Inverts a vector
func (v Vector2) Negate() Vector2 {
	return Vector2{X: -v.X, Y: -v.Y}
}

//Divide  a vector by a value (v / d)
func (v Vector2) Divide(d float32) Vector2 {
	return Vector2{X: v.X / d, Y: v.Y / d}
}

//DivideV a vector by another vecotr (v / v2)
func (v Vector2) DivideV(v2 Vector2) Vector2 {
	return Vector2{X: v.X / v2.Y, Y: v.Y / v2.Y}
}

//Normalize a vector
func (v Vector2) Normalize() Vector2 {
	return v.Divide(v.Length())
}

//Lerp a vector towards another vector
func (v Vector2) Lerp(target Vector2, amount float32) Vector2 {
	return Vector2{
		X: v.X + amount*(target.X-v.X),
		Y: v.Y + amount*(target.Y-v.Y),
	}
}

//Distance between two vectors
func (v Vector2) Distance(v2 Vector2) float32 {
	d := v2.Subtract(v)
	return d.Length()
}

//Reflect a vector. The mirror normal can be invisioned as a mirror perpendicular to the surface that is hit.
func (v Vector2) Reflect(mirrorNormal Vector2) Vector2 {
	return v.Add(v.Scale(-2 * v.DotProduct(mirrorNormal)))
}

//Vector3 a 3 dimensional vector
type Vector3 struct {
	X float32
	Y float32
	Z float32
}

//Decompose the vector into a slice of floats
func (v Vector3) Decompose() []float32 { return []float32{v.X, v.Y, v.Z} }

//Add two vectors (v1 + v2)
func (v Vector3) Add(v2 Vector3) Vector3 {
	return Vector3{X: v.X + v2.X, Y: v.Y + v2.Y, Z: v.Z + v2.Z}
}

//Subtract two vectors (v1 - v2)
func (v Vector3) Subtract(v2 Vector3) Vector3 {
	return Vector3{X: v.X - v2.X, Y: v.Y - v2.Y, Z: v.Z - v2.Z}
}

//Length of the vector
func (v Vector3) Length() float32 {
	return float32(math.Sqrt(float64(v.X*v.X) + float64(v.Y*v.Y) + float64(v.Z*v.Z)))
}

//SqrLength is the squared length of the vector
func (v Vector3) SqrLength() float32 {
	return float32(float64(v.X*v.X) + float64(v.Y*v.Y) + float64(v.Z*v.Z))
}

//DotProduct of the vector
func (v Vector3) DotProduct(v2 Vector3) float32 {
	return v.X*v2.X + v.Y*v2.Y + v.Z*v2.Z
}

//CrossProduct of the vector
func (v Vector3) CrossProduct(v2 Vector3) Vector3 {
	return Vector3{
		X: v.Y*v2.Z - v.Z*v2.Y,
		Y: v.Z*v2.X - v.X*v2.Z,
		Z: v.X*v2.Y - v.Y*v2.X,
	}
}

//Perpendicular to this vector
func (v Vector3) Perpendicular() Vector3 {
	cardinalAxis := &Vector3{X: 1, Y: 0, Z: 0}
	min := math.Abs(float64(v.X))

	if math.Abs(float64(v.Y)) < min {
		min = math.Abs(float64(v.Y))
		cardinalAxis.X = 0
		cardinalAxis.Y = 1
		cardinalAxis.Z = 0
	}

	if math.Abs(float64(v.Z)) < min {
		cardinalAxis.X = 0
		cardinalAxis.Y = 0
		cardinalAxis.Z = 1
	}

	return v.CrossProduct(*cardinalAxis)
}

//Distance between two vectors
func (v Vector3) Distance(v2 Vector3) float32 {
	d := v2.Subtract(v)
	return d.Length()
}

//Angle the vector creates with another vector
func (v Vector3) Angle(v2 Vector3) float32 {
	result := float32(math.Acos(Clamp(float64(v.DotProduct(v2)), -1, 1))) * Rad2Deg
	if result < 0 {
		result += 360
	}
	return result
}

//Scale the vector (v * scale)
func (v Vector3) Scale(scale float32) Vector3 {
	return Vector3{X: v.X * scale, Y: v.Y * scale, Z: v.Z * scale}
}

//Multiply a vector by another vector
func (v Vector3) Multiply(v2 Vector3) Vector3 {
	return Vector3{X: v.X * v2.Y, Y: v.Y * v2.Y, Z: v.Z * v2.Z}
}

//Negate or Inverts a vector
func (v Vector3) Negate() Vector3 {
	return Vector3{X: -v.X, Y: -v.Y, Z: -v.Z}
}

//Divide  a vector by a value (v / d)
func (v Vector3) Divide(d float32) Vector3 {
	return Vector3{X: v.X / d, Y: v.Y / d, Z: v.Z / d}
}

//DivideV a vector by another vecotr (v / v2)
func (v Vector3) DivideV(v2 Vector3) Vector3 {
	return Vector3{X: v.X / v2.Y, Y: v.Y / v2.Y, Z: v.Z / v2.Z}
}

//Normalize a vector
func (v Vector3) Normalize() Vector3 {
	length := v.Length()
	if length == 0 {
		length = 1
	}

	ilength := 1 / length
	return v.Scale(ilength)
}

//Lerp a vector towards another vector
func (v Vector3) Lerp(target Vector3, amount float32) Vector3 {
	return Vector3{
		X: v.X + amount*(target.X-v.X),
		Y: v.Y + amount*(target.Y-v.Y),
		Z: v.Z + amount*(target.Z-v.Z),
	}
}

//Reflect a vector. The mirror normal can be invisioned as a mirror perpendicular to the surface that is hit.
func (v Vector3) Reflect(mirrorNormal Vector3) Vector3 {
	return v.Subtract(mirrorNormal.Scale(2)).Scale(v.DotProduct(mirrorNormal))
}

//RotateByQuaternion rotates the vector
func (v Vector3) RotateByQuaternion(q Quaternion) Vector3 {
	return Vector3{
		X: v.X*(q.X*q.X+q.W*q.W-q.Y*q.Y-q.Z*q.Z) + v.Y*(2*q.X*q.Y-2*q.W*q.Z) + v.Z*(2*q.X*q.Z+2*q.W*q.Y),
		Y: v.X*(2*q.W*q.Z+2*q.X*q.Y) + v.Y*(q.W*q.W-q.X*q.X+q.Y*q.Y-q.Z*q.Z) + v.Z*(-2*q.W*q.X+2*q.Y*q.Z),
		Z: v.X*(-2*q.W*q.Y+2*q.X*q.Z) + v.Y*(2*q.W*q.X+2*q.Y*q.Z) + v.Z*(q.W*q.W-q.X*q.X-q.Y*q.Y+q.Z*q.Z),
	}
}

//OrthoNormalize makes two vectors normalized and orthogonal to each other
func (v *Vector3) OrthoNormalize(v2 *Vector3) {
	var tmp Vector3
	tmp = v.Normalize()
	v.X = tmp.X
	v.Y = tmp.Y
	v.Z = tmp.Z

	tmp = v.CrossProduct(*v2).Normalize().CrossProduct(*v)
	v2.X = tmp.X
	v2.Y = tmp.Y
	v2.Z = tmp.Z
}

//Transform a vector by a given matrix
func (v Vector3) Transform(m Matrix) Vector3 {
	return Vector3{
		X: m.M0*v.X + m.M4*v.Y + m.M8*v.Z + m.M12,
		Y: m.M1*v.X + m.M5*v.Y + m.M9*v.Z + m.M13,
		Z: m.M2*v.X + m.M6*v.Y + m.M10*v.Z + m.M14,
	}
}

//Min value for each pair of components
func (v Vector3) Min(v2 Vector3) Vector3 {
	return Vector3{
		X: float32(math.Min(float64(v.X), float64(v2.X))),
		Y: float32(math.Min(float64(v.Y), float64(v2.Z))),
		Z: float32(math.Min(float64(v.Y), float64(v2.Z))),
	}
}

//Max value for each pair of components
func (v Vector3) Max(v2 Vector3) Vector3 {
	return Vector3{
		X: float32(math.Max(float64(v.X), float64(v2.X))),
		Y: float32(math.Max(float64(v.Y), float64(v2.Z))),
		Z: float32(math.Max(float64(v.Y), float64(v2.Z))),
	}
}

//Barycenter computers the coordinates (u, v, w) for the vector with respect to triangle (a, b, c). Assumes vector is on plane with triangle
func (v Vector3) Barycenter(a, b, c Vector3) Vector3 {
	v0 := b.Subtract(a)
	v1 := c.Subtract(a)
	v2 := v.Subtract(a)

	d00 := v0.DotProduct(v0)
	d01 := v0.DotProduct(v1)
	d11 := v1.DotProduct(v1)
	d20 := v2.DotProduct(v0)
	d21 := v2.DotProduct(v1)

	denom := d00*d11 - d01*d01
	y := (d11*d20 - d01*d21) / denom
	z := (d00*d21 - d01*d20) / denom
	x := 1 - (z + y)
	return Vector3{X: x, Y: y, Z: z}
}

//Vector4 A 4 dimensional vector
type Vector4 struct {
	X float32
	Y float32
	Z float32
	W float32
}

//Quaternion are used to represent rotation without gimbal lock
type Quaternion struct {
	X float32
	Y float32
	Z float32
	W float32
}

type Matrix struct {
	M0  float32
	M1  float32
	M2  float32
	M3  float32
	M4  float32
	M5  float32
	M6  float32
	M7  float32
	M8  float32
	M9  float32
	M10 float32
	M11 float32
	M12 float32
	M13 float32
	M14 float32
	M15 float32
}

//Clamp a value between min and max
func Clamp(v float64, min float64, max float64) float64 {
	if v < min {
		return min
	}
	if v > max {
		return max
	}
	return v
}
