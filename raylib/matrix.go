package raylib

import "math"

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

//NewMatrixIdentity creates a identity
func NewMatrixIdentity() Matrix {
	return Matrix{
		M0: 1, M1: 0, M2: 0, M3: 0,
		M4: 0, M5: 1, M6: 0, M7: 0,
		M8: 0, M9: 0, M10: 1, M11: 0,
		M12: 0, M13: 0, M14: 0, M15: 1,
	}
}

//NewMatrixTranslate creates a blank translation matrix
func NewMatrixTranslate(x, y, z float32) Matrix {
	return Matrix{
		M0: 1, M1: 0, M2: 0, M3: x,
		M4: 0, M5: 1, M6: 0, M7: y,
		M8: 0, M9: 0, M10: 1, M11: z,
		M12: 0, M13: 0, M14: 0, M15: 1,
	}
}

//NewMatrixRotate creates a rotation matrix based of the acis and radians
func NewMatrixRotate(axis Vector3, radians float32) Matrix {
	x := axis.X
	y := axis.Y
	z := axis.Z
	length := axis.Length()
	if length != 1 && length != 0 {
		length = 1 / length
		x *= length
		y *= length
		z *= length
	}

	sinres := float32(math.Sin(float64(radians)))
	cosres := float32(math.Cos(float64(radians)))
	t := 1 - cosres
	return Matrix{
		M0: x*x*t + cosres,
		M1: y*x*t + z*sinres,
		M2: z*x*t - y*sinres,
		M3: 0,

		M4: x*y*t - z*sinres,
		M5: y*y*t + cosres,
		M6: z*y*t + x*sinres,
		M7: 0,

		M8:  x*z*t + y*sinres,
		M9:  y*z*t - x*sinres,
		M10: z*z*t + cosres,
		M11: 0,

		M12: 0,
		M13: 0,
		M14: 0,
		M15: 1,
	}
}

//ToQuaternion turns the rotation matrix into a Quaternion. Alias of newQuaternionFromMatrix
func (m Matrix) ToQuaternion() Quaternion { return NewQuaternionFromMatrix(m) }

//Trace of the matrix (sum of values along diagonal)
func (m Matrix) Trace() float32 {
	return m.M0 + m.M5 + m.M10 + m.M15
}

//Detrimant of the matrix
func (m Matrix) Detrimant() float32 {
	// Cache the matrix values (speed optimization)
	a00 := m.M0
	a01 := m.M1
	a02 := m.M2
	a03 := m.M3
	a10 := m.M4
	a11 := m.M5
	a12 := m.M6
	a13 := m.M7
	a20 := m.M8
	a21 := m.M9
	a22 := m.M10
	a23 := m.M11
	a30 := m.M12
	a31 := m.M13
	a32 := m.M14
	a33 := m.M15

	return a30*a21*a12*a03 - a20*a31*a12*a03 - a30*a11*a22*a03 + a10*a31*a22*a03 +
		a20*a11*a32*a03 - a10*a21*a32*a03 - a30*a21*a02*a13 + a20*a31*a02*a13 +
		a30*a01*a22*a13 - a00*a31*a22*a13 - a20*a01*a32*a13 + a00*a21*a32*a13 +
		a30*a11*a02*a23 - a10*a31*a02*a23 - a30*a01*a12*a23 + a00*a31*a12*a23 +
		a10*a01*a32*a23 - a00*a11*a32*a23 - a20*a11*a02*a33 + a10*a21*a02*a33 +
		a20*a01*a12*a33 - a00*a21*a12*a33 - a10*a01*a22*a33 + a00*a11*a22*a33
}

//Transpose the matrix
func (m Matrix) Transpose() Matrix {
	return Matrix{
		M0:  m.M0,
		M1:  m.M4,
		M2:  m.M8,
		M3:  m.M12,
		M4:  m.M1,
		M5:  m.M5,
		M6:  m.M9,
		M7:  m.M13,
		M8:  m.M2,
		M9:  m.M6,
		M10: m.M10,
		M11: m.M14,
		M12: m.M3,
		M13: m.M7,
		M14: m.M11,
		M15: m.M15,
	}
}

//Invert the matrix
func (m Matrix) Invert() Matrix {
	a00 := m.M0
	a01 := m.M1
	a02 := m.M2
	a03 := m.M3
	a10 := m.M4
	a11 := m.M5
	a12 := m.M6
	a13 := m.M7
	a20 := m.M8
	a21 := m.M9
	a22 := m.M10
	a23 := m.M11
	a30 := m.M12
	a31 := m.M13
	a32 := m.M14
	a33 := m.M15

	b00 := a00*a11 - a01*a10
	b01 := a00*a12 - a02*a10
	b02 := a00*a13 - a03*a10
	b03 := a01*a12 - a02*a11
	b04 := a01*a13 - a03*a11
	b05 := a02*a13 - a03*a12
	b06 := a20*a31 - a21*a30
	b07 := a20*a32 - a22*a30
	b08 := a20*a33 - a23*a30
	b09 := a21*a32 - a22*a31
	b10 := a21*a33 - a23*a31
	b11 := a22*a33 - a23*a32

	// Calculate the invert determinant (inlined to avoid double-caching)
	invDet := 1 / (b00*b11 - b01*b10 + b02*b09 + b03*b08 - b04*b07 + b05*b06)
	return Matrix{
		M0:  (a11*b11 - a12*b10 + a13*b09) * invDet,
		M1:  (-a01*b11 + a02*b10 - a03*b09) * invDet,
		M2:  (a31*b05 - a32*b04 + a33*b03) * invDet,
		M3:  (-a21*b05 + a22*b04 - a23*b03) * invDet,
		M4:  (-a10*b11 + a12*b08 - a13*b07) * invDet,
		M5:  (a00*b11 - a02*b08 + a03*b07) * invDet,
		M6:  (-a30*b05 + a32*b02 - a33*b01) * invDet,
		M7:  (a20*b05 - a22*b02 + a23*b01) * invDet,
		M8:  (a10*b10 - a11*b08 + a13*b06) * invDet,
		M9:  (-a00*b10 + a01*b08 - a03*b06) * invDet,
		M10: (a30*b04 - a31*b02 + a33*b00) * invDet,
		M11: (-a20*b04 + a21*b02 - a23*b00) * invDet,
		M12: (-a10*b09 + a11*b07 - a12*b06) * invDet,
		M13: (a00*b09 - a01*b07 + a02*b06) * invDet,
		M14: (-a30*b03 + a31*b01 - a32*b00) * invDet,
		M15: (a20*b03 - a21*b01 + a22*b00) * invDet,
	}
}

//Normalize calcuates the normal of the matrix
func (m Matrix) Normalize() Matrix {
	det := m.Detrimant()
	return Matrix{
		M0:  m.M0 / det,
		M1:  m.M1 / det,
		M2:  m.M2 / det,
		M3:  m.M3 / det,
		M4:  m.M4 / det,
		M5:  m.M5 / det,
		M6:  m.M6 / det,
		M7:  m.M7 / det,
		M8:  m.M8 / det,
		M9:  m.M9 / det,
		M10: m.M10 / det,
		M11: m.M11 / det,
		M12: m.M12 / det,
		M13: m.M13 / det,
		M14: m.M14 / det,
		M15: m.M15 / det,
	}
}

//Add two matrices
func (m Matrix) Add(right Matrix) Matrix {
	return Matrix{
		M0:  m.M0 + right.M0,
		M1:  m.M1 + right.M1,
		M2:  m.M2 + right.M2,
		M3:  m.M3 + right.M3,
		M4:  m.M4 + right.M4,
		M5:  m.M5 + right.M5,
		M6:  m.M6 + right.M6,
		M7:  m.M7 + right.M7,
		M8:  m.M8 + right.M8,
		M9:  m.M9 + right.M9,
		M10: m.M10 + right.M10,
		M11: m.M11 + right.M11,
		M12: m.M12 + right.M12,
		M13: m.M13 + right.M13,
		M14: m.M14 + right.M14,
		M15: m.M15 + right.M15,
	}
}

//Subtract two matrices
func (m Matrix) Subtract(right Matrix) Matrix {
	return Matrix{
		M0:  m.M0 - right.M0,
		M1:  m.M1 - right.M1,
		M2:  m.M2 - right.M2,
		M3:  m.M3 - right.M3,
		M4:  m.M4 - right.M4,
		M5:  m.M5 - right.M5,
		M6:  m.M6 - right.M6,
		M7:  m.M7 - right.M7,
		M8:  m.M8 - right.M8,
		M9:  m.M9 - right.M9,
		M10: m.M10 - right.M10,
		M11: m.M11 - right.M11,
		M12: m.M12 - right.M12,
		M13: m.M13 - right.M13,
		M14: m.M14 - right.M14,
		M15: m.M15 - right.M15,
	}
}

//NewMatrixRotateXYZ new xyz-rotation matrix (in radians)
func NewMatrixRotateXYZ(radians Vector3) Matrix {
	cosz := float32(math.Cos(float64(-radians.Z)))
	sinz := float32(math.Sin(float64(-radians.Z)))
	cosy := float32(math.Cos(float64(-radians.Y)))
	siny := float32(math.Sin(float64(-radians.Y)))
	cosx := float32(math.Cos(float64(-radians.X)))
	sinx := float32(math.Sin(float64(-radians.X)))
	result := NewMatrixIdentity()
	result.M0 = cosz * cosy
	result.M4 = (cosz * siny * sinx) - (sinz * cosx)
	result.M8 = (cosz * siny * cosx) + (sinz * sinx)
	result.M1 = sinz * cosy
	result.M5 = (sinz * siny * sinx) + (cosz * cosx)
	result.M9 = (sinz * siny * cosx) - (cosz * sinx)
	result.M2 = -siny
	result.M6 = cosy * sinx
	result.M10 = cosy * cosx
	return result

}

//NewMatrixRotateX creates a new matrix that is rotated
func NewMatrixRotateX(radians float32) Matrix {
	result := NewMatrixIdentity()
	cosres := float32(math.Cos(float64(radians)))
	sinres := float32(math.Sin(float64(radians)))
	result.M5 = cosres
	result.M6 = -sinres
	result.M9 = sinres
	result.M10 = cosres
	return result
}

//NewMatrixRotateY creates a new matrix that is rotated
func NewMatrixRotateY(radians float32) Matrix {
	result := NewMatrixIdentity()
	cosres := float32(math.Cos(float64(radians)))
	sinres := float32(math.Sin(float64(radians)))
	result.M0 = cosres
	result.M2 = sinres
	result.M8 = -sinres
	result.M10 = cosres
	return result

}

//NewMatrixRotateZ creates a new matrix that is rotated
func NewMatrixRotateZ(radians float32) Matrix {
	result := NewMatrixIdentity()
	cosres := float32(math.Cos(float64(radians)))
	sinres := float32(math.Sin(float64(radians)))
	result.M0 = cosres
	result.M1 = -sinres
	result.M4 = sinres
	result.M5 = cosres
	return result

}

//NewMatrixScale creates a new scalling matrix
func NewMatrixScale(scale Vector3) Matrix {
	return Matrix{
		M0: scale.X, M1: 0, M2: 0, M3: 0,
		M4: 0, M5: scale.Y, M6: 0, M7: 0,
		M8: 0, M9: 0, M10: scale.Z, M11: 0,
		M12: 0, M13: 0, M14: 0, M15: 1,
	}
}

//Multiply two matrix together. Note that order matters.
func (m Matrix) Multiply(right Matrix) Matrix {
	return Matrix{
		M0:  m.M0*right.M0 + m.M1*right.M4 + m.M2*right.M8 + m.M3*right.M12,
		M1:  m.M0*right.M1 + m.M1*right.M5 + m.M2*right.M9 + m.M3*right.M13,
		M2:  m.M0*right.M2 + m.M1*right.M6 + m.M2*right.M10 + m.M3*right.M14,
		M3:  m.M0*right.M3 + m.M1*right.M7 + m.M2*right.M11 + m.M3*right.M15,
		M4:  m.M4*right.M0 + m.M5*right.M4 + m.M6*right.M8 + m.M7*right.M12,
		M5:  m.M4*right.M1 + m.M5*right.M5 + m.M6*right.M9 + m.M7*right.M13,
		M6:  m.M4*right.M2 + m.M5*right.M6 + m.M6*right.M10 + m.M7*right.M14,
		M7:  m.M4*right.M3 + m.M5*right.M7 + m.M6*right.M11 + m.M7*right.M15,
		M8:  m.M8*right.M0 + m.M9*right.M4 + m.M10*right.M8 + m.M11*right.M12,
		M9:  m.M8*right.M1 + m.M9*right.M5 + m.M10*right.M9 + m.M11*right.M13,
		M10: m.M8*right.M2 + m.M9*right.M6 + m.M10*right.M10 + m.M11*right.M14,
		M11: m.M8*right.M3 + m.M9*right.M7 + m.M10*right.M11 + m.M11*right.M15,
		M12: m.M12*right.M0 + m.M13*right.M4 + m.M14*right.M8 + m.M15*right.M12,
		M13: m.M12*right.M1 + m.M13*right.M5 + m.M14*right.M9 + m.M15*right.M13,
		M14: m.M12*right.M2 + m.M13*right.M6 + m.M14*right.M10 + m.M15*right.M14,
		M15: m.M12*right.M3 + m.M13*right.M7 + m.M14*right.M11 + m.M15*right.M15,
	}
}

//NewMatrixFrustum creates a new perspective projection matrix
func NewMatrixFrustum(left, right, bottom, top, near, far float64) Matrix {
	rl := (right - left)
	tb := (top - bottom)
	fn := (far - near)

	return Matrix{
		M0:  float32((near * 2) / rl),
		M1:  0,
		M2:  0,
		M3:  0,
		M4:  0,
		M5:  float32((near * 2) / tb),
		M6:  0,
		M7:  0,
		M8:  float32((right + left) / rl),
		M9:  float32((top + bottom) / tb),
		M10: float32(-(far + near) / fn),
		M11: -1,
		M12: 0,
		M13: 0,
		M14: float32(-(far * near * 2) / fn),
		M15: 0,
	}
}

//NewMatrixPerspective creates a perspective projection matrix. Angles should be provided in radians
func NewMatrixPerspective(fovy, aspect, near, far float64) Matrix {
	top := near * math.Tan(fovy*0.5)
	right := top * aspect
	return NewMatrixFrustum(-right, right, -top, top, near, far)
}

//NewMatrixOrtho creates a orthographic projection
func NewMatrixOrtho(left, right, bottom, top, near, far float64) Matrix {
	rl := (right - left)
	tb := (top - bottom)
	fn := (far - near)
	return Matrix{
		M0:  float32(2 / rl),
		M1:  0,
		M2:  0,
		M3:  0,
		M4:  0,
		M5:  float32(2 / tb),
		M6:  0,
		M7:  0,
		M8:  0,
		M9:  0,
		M10: float32(-2 / fn),
		M11: 0,
		M12: float32(-(left + right) / rl),
		M13: float32(-(top + bottom) / tb),
		M14: float32(-(far + near) / fn),
		M15: 1,
	}
}

//NewMatrixLookAt creates a matrix to look at a target
func NewMatrixLookAt(eye, target, up Vector3) Matrix {
	z := eye.Subtract(target).Normalize()
	x := up.CrossProduct(z).Normalize()
	y := z.CrossProduct(x).Normalize()
	return Matrix{
		M0:  x.X,
		M1:  x.Y,
		M2:  x.Z,
		M3:  0,
		M4:  y.X,
		M5:  y.Y,
		M6:  y.Z,
		M7:  0,
		M8:  z.X,
		M9:  z.Y,
		M10: z.Z,
		M11: 0,
		M12: eye.X,
		M13: eye.Y,
		M14: eye.Z,
		M15: 1,
	}
}

//Decompose turns a matrix into an slice of floats
func (m Matrix) Decompose() []float32 {
	return []float32{
		m.M0, m.M1, m.M2, m.M3, m.M4, m.M5, m.M6, m.M7, m.M8, m.M9,
		m.M10, m.M11, m.M12, m.M13, m.M14, m.M15,
	}
}
