//Update : Update audio stream buffers with data
func (stream *AudioStream) Update(data []float32, samplesCount int) {
	cstream := *stream.cptr()
	C.UpdateAudioStream(cstream, unsafe.Pointer(&data[0]), C.int(int32(samplesCount)))
}

//UpdateSound : Update audio stream buffers with data
//Recommended to use stream.Update(data, samplesCount) instead
func UpdateAudioStream(stream *AudioStream, data []float32, samplesCount int) {
	stream.Update(data, samplesCount)
}