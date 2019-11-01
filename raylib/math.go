package raylib

import "unsafe"

const (
	//PI 3.1415
	PI float32 = 3.14159265358979323846

	//Deg2Rad used to convert degrees into radians
	Deg2Rad float32 = PI / 180

	//Rad2Deg used to convert radians into degrees
	Rad2Deg float32 = 180 / PI
)

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

//Clamp32 a value between min and max
func Clamp32(v float32, min float32, max float32) float32 {
	if v < min {
		return min
	}
	if v > max {
		return max
	}
	return v
}

// Float32frombits returns the floating point number corresponding
// to the IEEE 754 binary representation b.
func Float32frombits(b uint32) float32 { return *(*float32)(unsafe.Pointer(&b)) }

//NaN32 returns an IEEE 754 ``not-a-number'' value. Source: https://github.com/chewxy/math32/blob/9a000fcb79dff2019bd78fc28bd676198ff3a616/bits.go
func NaN32() float32 { return Float32frombits(0x7FE00000) }

// Inf returns positive infinity if sign >= 0, negative infinity if sign < 0.
func Inf(sign int) float32 {
	if sign >= 0 {
		return Float32frombits(0x7F800000)
	}

	return Float32frombits(0xFF800000)
}
