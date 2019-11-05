//GetWaveData : Get samples data from wave as a floats array
func GetWaveData(wave Wave) []float32 {

	samples := wave.SampleCount * wave.Channels
	cwave := *wave.cptr()

	res := C.GetWaveData(cwave)
	tmpslice := (*[1 << 24]*C.float)(unsafe.Pointer(res))[:samples:samples]
	defer C.free(unsafe.Pointer(res))

	gostrings := make([]float32, samples)
	for i, s := range tmpslice {
		gostrings[i] = float32(s)
	}

	return gostrings
}
