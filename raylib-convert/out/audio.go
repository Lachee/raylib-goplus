package raylib

//InitAudioDevice : Initialize audio device and context
func InitAudioDevice() {
	C.InitAudioDevice()
}

//CloseAudioDevice : Close the audio device and context
func CloseAudioDevice() {
	C.CloseAudioDevice()
}

//IsAudioDeviceReady : Check if audio device has been initialized successfully
func IsAudioDeviceReady() bool {
	res := C.IsAudioDeviceReady()
	return bool(res)
}

//SetMasterVolume : Set master volume (listener)
func SetMasterVolume(volume float32) {
	C.SetMasterVolume(C.float(volume))
}

//LoadWave : Load wave data from file
func LoadWave(fileName string) (Wave, string) {
	cfileName := C.CString(fileName)
	defer C.free(unsafe.Pointer(cfileName))
	res := C.LoadWave(cfileName)
	return newWaveFromPointer(unsafe.Pointer(&res)), C.GoString(cfileName)
}

//LoadSound : Load sound from file
func LoadSound(fileName string) (Sound, string) {
	cfileName := C.CString(fileName)
	defer C.free(unsafe.Pointer(cfileName))
	res := C.LoadSound(cfileName)
	return newSoundFromPointer(unsafe.Pointer(&res)), C.GoString(cfileName)
}

//LoadSoundFromWave : Load sound from wave data
func LoadSoundFromWave(wave Wave) Sound {
	cwave := *wave.cptr()
	res := C.LoadSoundFromWave(cwave)
	return newSoundFromPointer(unsafe.Pointer(&res))
}

//UpdateSound : Update sound buffer with new data
func UpdateSound(sound Sound, data unsafe.Pointer, samplesCount int) {
	csound := *sound.cptr()
	C.UpdateSound(csound, data, C.int(samplesCount))
}

//UnloadWave : Unload wave data
func UnloadWave(wave Wave) {
	cwave := *wave.cptr()
	C.UnloadWave(cwave)
}

//UnloadSound : Unload sound
func UnloadSound(sound Sound) {
	csound := *sound.cptr()
	C.UnloadSound(csound)
}

//ExportWave : Export wave data to file
func ExportWave(wave Wave, fileName string) string {
	cfileName := C.CString(fileName)
	defer C.free(unsafe.Pointer(cfileName))
	cwave := *wave.cptr()
	C.ExportWave(cwave, cfileName)
	return C.GoString(cfileName)
}

//ExportWaveAsCode : Export wave sample data to code (.h)
func ExportWaveAsCode(wave Wave, fileName string) string {
	cfileName := C.CString(fileName)
	defer C.free(unsafe.Pointer(cfileName))
	cwave := *wave.cptr()
	C.ExportWaveAsCode(cwave, cfileName)
	return C.GoString(cfileName)
}

//PlaySound : Play a sound
func PlaySound(sound Sound) {
	csound := *sound.cptr()
	C.PlaySound(csound)
}

//StopSound : Stop playing a sound
func StopSound(sound Sound) {
	csound := *sound.cptr()
	C.StopSound(csound)
}

//PauseSound : Pause a sound
func PauseSound(sound Sound) {
	csound := *sound.cptr()
	C.PauseSound(csound)
}

//ResumeSound : Resume a paused sound
func ResumeSound(sound Sound) {
	csound := *sound.cptr()
	C.ResumeSound(csound)
}

//PlaySoundMulti : Play a sound (using multichannel buffer pool)
func PlaySoundMulti(sound Sound) {
	csound := *sound.cptr()
	C.PlaySoundMulti(csound)
}

//StopSoundMulti : Stop any sound playing (using multichannel buffer pool)
func StopSoundMulti() {
	C.StopSoundMulti()
}

//GetSoundsPlaying : Get number of sounds playing in the multichannel
func GetSoundsPlaying() int {
	res := C.GetSoundsPlaying()
	return int(res)
}

//IsSoundPlaying : Check if a sound is currently playing
func IsSoundPlaying(sound Sound) bool {
	csound := *sound.cptr()
	res := C.IsSoundPlaying(csound)
	return bool(res)
}

//SetSoundVolume : Set volume for a sound (1.0 is max level)
func SetSoundVolume(sound Sound, volume float32) {
	csound := *sound.cptr()
	C.SetSoundVolume(csound, C.float(volume))
}

//SetSoundPitch : Set pitch for a sound (1.0 is base level)
func SetSoundPitch(sound Sound, pitch float32) {
	csound := *sound.cptr()
	C.SetSoundPitch(csound, C.float(pitch))
}

//WaveFormat : Convert wave data to desired format
func WaveFormat(wave Wave, sampleRate int, sampleSize int, channels int) Wave {
	cwave := wave.cptr()
	C.WaveFormat(&cwave, C.int(sampleRate), C.int(sampleSize), C.int(channels))
	return newWaveFromPointer(unsafe.Pointer(&cwave))
}

//WaveCopy : Copy a wave to a new wave
func WaveCopy(wave Wave) Wave {
	cwave := *wave.cptr()
	res := C.WaveCopy(cwave)
	return newWaveFromPointer(unsafe.Pointer(&res))
}

//WaveCrop : Crop a wave to defined samples range
func WaveCrop(wave Wave, initSample int, finalSample int) Wave {
	cwave := wave.cptr()
	C.WaveCrop(&cwave, C.int(initSample), C.int(finalSample))
	return newWaveFromPointer(unsafe.Pointer(&cwave))
}

//LoadMusicStream : Load music stream from file
func LoadMusicStream(fileName string) (Music, string) {
	cfileName := C.CString(fileName)
	defer C.free(unsafe.Pointer(cfileName))
	res := C.LoadMusicStream(cfileName)
	return newMusicFromPointer(unsafe.Pointer(&res)), C.GoString(cfileName)
}

//UnloadMusicStream : Unload music stream
func UnloadMusicStream(music Music) {
	cmusic := *music.cptr()
	C.UnloadMusicStream(cmusic)
}

//PlayMusicStream : Start music playing
func PlayMusicStream(music Music) {
	cmusic := *music.cptr()
	C.PlayMusicStream(cmusic)
}

//UpdateMusicStream : Updates buffers for music streaming
func UpdateMusicStream(music Music) {
	cmusic := *music.cptr()
	C.UpdateMusicStream(cmusic)
}

//StopMusicStream : Stop music playing
func StopMusicStream(music Music) {
	cmusic := *music.cptr()
	C.StopMusicStream(cmusic)
}

//PauseMusicStream : Pause music playing
func PauseMusicStream(music Music) {
	cmusic := *music.cptr()
	C.PauseMusicStream(cmusic)
}

//ResumeMusicStream : Resume playing paused music
func ResumeMusicStream(music Music) {
	cmusic := *music.cptr()
	C.ResumeMusicStream(cmusic)
}

//IsMusicPlaying : Check if music is playing
func IsMusicPlaying(music Music) bool {
	cmusic := *music.cptr()
	res := C.IsMusicPlaying(cmusic)
	return bool(res)
}

//SetMusicVolume : Set volume for music (1.0 is max level)
func SetMusicVolume(music Music, volume float32) {
	cmusic := *music.cptr()
	C.SetMusicVolume(cmusic, C.float(volume))
}

//SetMusicPitch : Set pitch for a music (1.0 is base level)
func SetMusicPitch(music Music, pitch float32) {
	cmusic := *music.cptr()
	C.SetMusicPitch(cmusic, C.float(pitch))
}

//SetMusicLoopCount : Set music loop count (loop repeats)
func SetMusicLoopCount(music Music, count int) {
	cmusic := *music.cptr()
	C.SetMusicLoopCount(cmusic, C.int(count))
}

//GetMusicTimeLength : Get music time length (in seconds)
func GetMusicTimeLength(music Music) float32 {
	cmusic := *music.cptr()
	res := C.GetMusicTimeLength(cmusic)
	return float32(res)
}

//GetMusicTimePlayed : Get current music time played (in seconds)
func GetMusicTimePlayed(music Music) float32 {
	cmusic := *music.cptr()
	res := C.GetMusicTimePlayed(cmusic)
	return float32(res)
}

//InitAudioStream : Init audio stream (to stream raw audio pcm data)
func InitAudioStream(sampleRate int, sampleSize int, channels int) AudioStream {
	res := C.InitAudioStream(C.int(sampleRate), C.int(sampleSize), C.int(channels))
	return newAudioStreamFromPointer(unsafe.Pointer(&res))
}

//UpdateAudioStream : Update audio stream buffers with data
func UpdateAudioStream(stream AudioStream, data unsafe.Pointer, samplesCount int) {
	cstream := *stream.cptr()
	C.UpdateAudioStream(cstream, data, C.int(samplesCount))
}

//CloseAudioStream : Close audio stream and free memory
func CloseAudioStream(stream AudioStream) {
	cstream := *stream.cptr()
	C.CloseAudioStream(cstream)
}

//IsAudioStreamProcessed : Check if any audio stream buffers requires refill
func IsAudioStreamProcessed(stream AudioStream) bool {
	cstream := *stream.cptr()
	res := C.IsAudioStreamProcessed(cstream)
	return bool(res)
}

//PlayAudioStream : Play audio stream
func PlayAudioStream(stream AudioStream) {
	cstream := *stream.cptr()
	C.PlayAudioStream(cstream)
}

//PauseAudioStream : Pause audio stream
func PauseAudioStream(stream AudioStream) {
	cstream := *stream.cptr()
	C.PauseAudioStream(cstream)
}

//ResumeAudioStream : Resume audio stream
func ResumeAudioStream(stream AudioStream) {
	cstream := *stream.cptr()
	C.ResumeAudioStream(cstream)
}

//IsAudioStreamPlaying : Check if audio stream is playing
func IsAudioStreamPlaying(stream AudioStream) bool {
	cstream := *stream.cptr()
	res := C.IsAudioStreamPlaying(cstream)
	return bool(res)
}

//StopAudioStream : Stop audio stream
func StopAudioStream(stream AudioStream) {
	cstream := *stream.cptr()
	C.StopAudioStream(cstream)
}

//SetAudioStreamVolume : Set volume for audio stream (1.0 is max level)
func SetAudioStreamVolume(stream AudioStream, volume float32) {
	cstream := *stream.cptr()
	C.SetAudioStreamVolume(cstream, C.float(volume))
}

//SetAudioStreamPitch : Set pitch for audio stream (1.0 is base level)
func SetAudioStreamPitch(stream AudioStream, pitch float32) {
	cstream := *stream.cptr()
	C.SetAudioStreamPitch(cstream, C.float(pitch))
}
