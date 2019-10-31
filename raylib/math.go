package raylib

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
