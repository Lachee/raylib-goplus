package raylib

import (
	"math"
	"unsafe"
)

var (
	// LightGray ...
	LightGray = NewColor(200, 200, 200, 255)
	//Gray ...
	Gray = NewColor(130, 130, 130, 255)
	//DarkGray ...
	DarkGray = NewColor(80, 80, 80, 255)
	//Yellow ...
	Yellow = NewColor(253, 249, 0, 255) // Yellow
	//Gold ...
	Gold = NewColor(255, 203, 0, 255) // Gold
	//Orange ...
	Orange = NewColor(255, 161, 0, 255) // Orange
	//Pink ...
	Pink = NewColor(255, 109, 194, 255) // Pink
	//Red ...
	Red = NewColor(230, 41, 55, 255) // Red
	//Maroon ...
	Maroon = NewColor(190, 33, 55, 255) // Maroon
	//Green ...
	Green = NewColor(0, 228, 48, 255) // Green
	//Lime ...
	Lime = NewColor(0, 158, 47, 255) // Lime
	//DarkGreen ...
	DarkGreen = NewColor(0, 117, 44, 255) // Dark Green
	//SkyBlue ...
	SkyBlue = NewColor(102, 191, 255, 255) // Sky Blue
	//Blue ...
	Blue = NewColor(0, 121, 241, 255) // Blue
	//DarkBlue ...
	DarkBlue = NewColor(0, 82, 172, 255) // Dark Blue
	//Purple ...
	Purple = NewColor(200, 122, 255, 255) // Purple
	//Violet ...
	Violet = NewColor(135, 60, 190, 255) // Violet
	//DarkPurple ...
	DarkPurple = NewColor(112, 31, 126, 255) // Dark Purple
	//Beige ...
	Beige = NewColor(211, 176, 131, 255) // Beige
	//Brown ...
	Brown = NewColor(127, 106, 79, 255) // Brown
	//DarkBrown ...
	DarkBrown = NewColor(76, 63, 47, 255) // Dark Brown
	//White ...
	White = NewColor(255, 255, 255, 255) // White
	//Black ...
	Black = NewColor(0, 0, 0, 255) // Black
	//Blank (Transparent)
	Blank = NewColor(0, 0, 0, 0) // Blank (Transparent)
	//Magenta ...
	Magenta = NewColor(255, 0, 255, 255) // Magenta
	//RayWhite - Off White
	RayWhite = NewColor(245, 245, 245, 255) // My own White (raylib logo)
	//Aqua
	Aqua = NewColor(0, 162, 156, 255)
	//Gopher Blue
	GopherBlue = NewColor(1, 173, 216, 255)
	//Transparent
	Transparent = NewColor(0, 0, 0, 0)
)

// Color type, RGBA (32bit)
type Color struct {
	R uint8
	G uint8
	B uint8
	A uint8
}

func newColorFromPointer(ptr unsafe.Pointer) Color { return *(*Color)(ptr) }

//NewColor creates a new colour
func NewColor(r, g, b, a uint8) Color { return Color{R: r, G: g, B: b, A: a} }

//NewColorFromHSV turns a HSV to a colour
func NewColorFromHSV(hsv Vector3) Color {

	h, s, v := float64(hsv.X), float64(hsv.Y), float64(hsv.Z)
	var k, t, r, g, b float64

	//RED
	k = math.Mod((5 + h/60), 6)
	t = 4 - k
	if t < k {
		k = t
	}
	k = Clamp(k, 0, 1)
	r = (v - v*s*k) * 255

	//Green
	k = math.Mod((3 + h/60), 6)
	t = 4 - k
	if t < k {
		k = t
	}
	k = Clamp(k, 0, 1)
	g = (v - v*s*k) * 255

	//Blue
	k = math.Mod((1 + h/60), 6)
	t = 4 - k
	if t < k {
		k = t
	}
	k = Clamp(k, 0, 1)
	b = (v - v*s*k) * 255

	return Color{
		R: uint8(r),
		G: uint8(g),
		B: uint8(b),
		A: 255,
	}
}

//NewColorFromNormalized creates a colour from the normalized value [0..1]
func NewColorFromNormalized(normalized Vector4) Color {
	return Color{
		R: uint8(255 * Clamp32(normalized.X, 0, 1)),
		G: uint8(255 * Clamp32(normalized.Y, 0, 1)),
		B: uint8(255 * Clamp32(normalized.Z, 0, 1)),
		A: uint8(255 * Clamp32(normalized.W, 0, 1)),
	}
}

//NewColorInt creates a colour from a hexadecimal int
func NewColorInt(hexValue int) Color {
	return Color{
		R: uint8((hexValue >> 24) & 0xFF),
		G: uint8((hexValue >> 16) & 0xFF),
		B: uint8((hexValue >> 8) & 0xFF),
		A: uint8(hexValue & 0xFF),
	}
}

//ToInt converts the colour into a hexadecimal int
func (c Color) ToInt() int {
	return (int(c.R) << 24) | (int(c.G) << 16) | (int(c.B) << 8) | int(c.A)
}

//Normalize returns the normalized colour as floats [0..1]
func (c Color) Normalize() Vector4 {
	return NewVector4(float32(c.R)/255.0, float32(c.G)/255.0, float32(c.B)/255.0, float32(c.A)/255.0)
}

func (c Color) ToHSV() Vector3 {
	r, g, b := float64(c.R)/255.0, float64(c.G)/255.0, float64(c.B)/255.0
	hsv := &Vector3{X: 0, Y: 0, Z: 0}

	min := math.Min(math.Min(r, g), b)
	max := math.Max(math.Max(r, g), b)

	delta := max - min
	hsv.Z = float32(max)
	if delta < 0.00001 {
		return *hsv
	}

	if max > 0 {
		hsv.Y = float32(delta / max)
	} else {
		hsv.X = NaN32()
		return *hsv
	}

	if r >= max {
		hsv.X = float32((g - b) / delta)
	} else {
		if g >= max {
			hsv.X = 2 + float32((b - r/delta))
		} else {
			hsv.X = 4 + float32((r - g/delta))
		}
	}

	//Convert to degrees
	hsv.X *= 60
	if hsv.X < 0 {
		hsv.X += 360
	}

	return *hsv
}

//Fade a colour
func (c Color) Fade(alpha float32) Color {
	return Color{R: c.R, B: c.B, G: c.G, A: uint8(255 * Clamp32(alpha, 0, 1))}
}

//Lerp a color towards another color
func (c Color) Lerp(target Color, amount float32) Color {
	return NewColorFromNormalized(c.Normalize().Lerp(target.Normalize(), amount))
}

//LerpHSV will lerp the color towards another, using their calculated HSV to do so.
// Note it is more efficient to store the calculated HSV of each and lerp between them
func (c Color) LerpHSV(target Color, amount float32) Color {
	hsv1 := c.ToHSV()
	hsv2 := target.ToHSV()
	lerp := hsv1.Lerp(hsv2, amount)
	return NewColorFromHSV(lerp)
}
