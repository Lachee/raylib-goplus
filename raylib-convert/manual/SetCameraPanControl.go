//SetCameraPanControl : Set camera pan key to combine with mouse movement (free camera)
func SetCameraPanControl(panKey Key) {
	C.SetCameraPanControl(C.int(panKey))
}
