//Unload : Close audio stream and free memory
func (stream *AudioStream) Unload() {
	cstream := *stream.cptr()
	C.CloseAudioStream(cstream)
}

//CloseAudioStream : Close audio stream and free memory
//Recommended to use stream.Unload() instead
func CloseAudioStream(stream *AudioStream) {
	stream.Unload()
}