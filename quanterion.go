package ray

import "math"

//Quaternion A represntation of rotations that does not suffer from gimbal lock
type Quaternion struct {
	X float32
	Y float32
	Z float32
	W float32
}

func newQuaternionIdentity() Quaternion { return Quaternion{X: 0, Y: 0, Z: 0, W: 1} }
func newQuaternionVector3ToVector3(from, too Vector3) Quaternion {
	cos2theta := from.DotProduct(too)
	cross := from.CrossProduct(too)
	return Quaternion{X: cross.X, Y: cross.Y, Z: cross.Z, W: 1 + cos2theta}.Normalize()
}
func newQuaternionFromMatrix(matrix Matrix) Quaternion {
	return nil
}
func newQuaternionFromEuler(euler Vector3) Quaternion {
	return nil
}

//Invert a quaternions components
func (q Quaternion) Invert() Quaternion {
	length := q.SqrLength()
	if length != 0 {
		i := 1 / length
		return Quaternion{
			X: q.X * -i,
			Y: q.Y * -i,
			Z: q.Z * -i,
			W: q.W * i,
		}
	}
	return q
}

//Decompose the vector into a slice of floats
func (q Quaternion) Decompose() []float32 { return []float32{q.X, q.Y, q.Z, q.W} }

//Length of the vector
func (q Quaternion) Length() float32 {
	return float32(math.Sqrt(float64(q.X*q.X) + float64(q.Y*q.Y) + float64(q.Z*q.Z) + float64(q.W*q.W)))
}

//SqrLength is the squared length of the vector
func (q Quaternion) SqrLength() float32 {
	return float32(float64(q.X*q.X) + float64(q.Y*q.Y) + float64(q.Z*q.Z) + float64(q.W*q.W))
}

//Scale the vector (v * scale)
func (q Quaternion) Scale(scale float32) Quaternion {
	return Quaternion{X: q.X * scale, Y: q.Y * scale, Z: q.Z * scale, W: q.W * scale}
}

//Normalize a vector
func (q Quaternion) Normalize() Quaternion {
	length := q.Length()
	if length == 0 {
		length = 1
	}

	ilength := 1 / length
	return q.Scale(ilength)
}

//Multiply two Quaternion together, doing queraternion mathmatics
func (q Quaternion) Multiply(q2 Quaternion) Quaternion {
	return Quaternion{
		X: q.X*q2.W + q.W*q2.X + q.Y*q2.Z - q.Z*q2.Y,
		Y: q.Y*q2.W + q.W*q2.Y + q.Z*q2.X - q.X*q2.Z,
		Z: q.Z*q2.W + q.W*q2.Z + q.X*q2.Y - q.Y*q2.X,
		W: q.W*q2.W - q.X*q2.X - q.Y*q2.Y - q.Z*q2.Z,
	}
}

//Lerp a vector towards another vector
func (q Quaternion) Lerp(target Quaternion, amount float32) Quaternion {
	return Quaternion{
		X: q.X + amount*(target.X-q.X),
		Y: q.Y + amount*(target.Y-q.Y),
		Z: q.Z + amount*(target.Z-q.Z),
		W: q.W + amount*(target.W-q.W),
	}
}

//Nlerp slerp-optimized interpolation between two quaternions
func (q Quaternion) Nlerp(target Quaternion, amount float32) Quaternion {
	return q.Lerp(target, amount).Normalize()
}

//Slerp Spherically Lerped
func (q Quaternion) Slerp(q2 Quaternion, amount float32) Quaternion {
	cosHalfTheta := q.X*q2.X + q.Y*q2.Y + q.Z*q2.Z + q.W*q2.W
	if math.Abs((float64(cosHalfTheta))) >= 1 {
		return q
	}

	if cosHalfTheta > 0.95 {
		return q.Nlerp(q2, amount)
	}

	halfTheta := float32(math.Acos(float64(cosHalfTheta)))
	sinHalfTheta := float32(math.Sqrt(float64(1 - cosHalfTheta*cosHalfTheta)))

	if math.Abs(float64(sinHalfTheta)) < 0.001 {
		return Quaternion{
			X: q.X*0.5 + q.X*0.5,
			Y: q.Y*0.5 + q.Y*0.5,
			Z: q.Z*0.5 + q.Z*0.5,
			W: q.W*0.5 + q.W*0.5,
		}
	}

	ratioA := float32(math.Sin(float64((1-amount)*halfTheta)) / float64(sinHalfTheta))
	ratioB := float32(math.Sin(float64(amount*halfTheta)) / float64(sinHalfTheta))

	return Quaternion{
		X: q.X*ratioA + q.X*ratioB,
		Y: q.Y*ratioA + q.Y*ratioB,
		Z: q.Z*ratioA + q.Z*ratioB,
		W: q.W*ratioA + q.W*ratioB,
	}
}

//https://github.com/raysan5/raylib/blob/master/src/raymath.h
func (q Quaternion) ToMatrix() Matrix {
	return nil
}

func (q Quaternion) ToAxisAngle(outAxis *Vector3, outAngle *float32) {

}

func (q Quaternion) ToEuler() Vector3 {
	return nil
}

func (q Quaternion) Transform(Matrix m) Quaternion {
	return nil
}
