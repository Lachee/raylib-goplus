//SetCameraSmoothZoomControl : Set camera smooth zoom key to combine with mouse (free camera)
func SetCameraSmoothZoomControl(szKey Key) {
	C.SetCameraSmoothZoomControl(C.int(szKey))
}
