//SetCameraAltControl : Set camera alt key to combine with mouse movement (free camera)
func SetCameraAltControl(altKey Key) {
	C.SetCameraAltControl(C.int(altKey))
}