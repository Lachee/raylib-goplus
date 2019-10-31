package raylib

//Matrix A representation of a 4 x 4 matrix
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

//NewMatrixFromQuaternion creates a new rotation matrix from a quaternion
func NewMatrixFromQuaternion(q Quaternion) Matrix {
	x := q.X
	y := q.Y
	z := q.Z
	w := q.W
	x2 := x + x
	y2 := y + y
	z2 := z + z
	lengthSquared := q.SqrLength()

	xx := x * x2 / lengthSquared
	xy := x * y2 / lengthSquared
	xz := x * z2 / lengthSquared

	yy := y * y2 / lengthSquared
	yz := y * z2 / lengthSquared
	zz := z * z2 / lengthSquared

	wx := w * x2 / lengthSquared
	wy := w * y2 / lengthSquared
	wz := w * z2 / lengthSquared

	return Matrix{
		M0:  1.0 - (yy + zz),
		M1:  xy - wz,
		M2:  xz + wy,
		M3:  0.0,
		M4:  xy + wz,
		M5:  1.0 - (xx + zz),
		M6:  yz - wx,
		M7:  0.0,
		M8:  xz - wy,
		M9:  yz + wx,
		M10: 1.0 - (xx + yy),
		M11: 0.0,
		M12: 0.0,
		M13: 0.0,
		M14: 0.0,
		M15: 1.0,
	}
}

func NewMatrixIdentity() Matrix {}

//ToQuaternion turns the rotation matrix into a Quaternion. Alias of newQuaternionFromMatrix
func (m Matrix) ToQuaternion() Quaternion { return NewQuaternionFromMatrix(m) }

//Trace of the matrix (sum of values along diagonal)
func (m Matrix) Trace() float32 {
	return m.M0 + m.M5 + m.M10 + m.M15
}

func (m Matrix) Detrimant() float32                                       {}
func (m Matrix) Transpose() Matrix                                        {}
func (m Matrix) Invert() Matrix                                           {}
func (m Matrix) Normalize() Matrix                                        {}
func (m Matrix) Add(right Matrix) Matrix                                  {}
func (m Matrix) Subtract(right Matrix) Matrix                             {}
func NewMatrixTranslate(x, y, z float32) Matrix                           {}
func NewMatrixRotate(axis Vector3, angle float32) Matrix                  {}
func NewMatrixRotateXYZ(ang Vector3) Matrix                               {}
func (m Matrix) RotateX(angle float32) Matrix                             {}
func (m Matrix) RotateY(angle float32) Matrix                             {}
func (m Matrix) RotateZ(angle float32) Matrix                             {}
func (m Matrix) Scale(scale Vector3) Matrix                               {}
func (m Matrix) Multiply(right Matrix) Matrix                             {}
func NewMatrixFrustum(left, right, bottom, top, near, far float64) Matrix {}
func NewMatrixPerspective(fovby, aspect, near, far float64) Matrix        {}
func NewMatrixOrtho(left, right, bottom, top, near, far float64) Matrix   {}
func NewMatrixLookAt(eye, target, up Vector3) Matrix                      {}
func (m Matrix) Decompose() []float32                                     {}
