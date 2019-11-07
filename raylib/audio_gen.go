package raylib

/*
#include "raylib.h"
#include <stdlib.h>
#include "go.h"
*/
import "C"
import "unsafe"

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
func LoadWave(fileName string) *Wave {
	cfileName := C.CString(fileName)
	defer C.free(unsafe.Pointer(cfileName))
	res := C.LoadWave(cfileName)
	retval := newWaveFromPointer(unsafe.Pointer(&res))
	addUnloadable(retval)
	return retval
}

//LoadSound : Load sound from file
func LoadSound(fileName string) *Sound {
	cfileName := C.CString(fileName)
	defer C.free(unsafe.Pointer(cfileName))
	res := C.LoadSound(cfileName)
	retval := newSoundFromPointer(unsafe.Pointer(&res))
	addUnloadable(retval)
	return retval
}

//LoadSoundFromWave : Load sound from wave data
func LoadSoundFromWave(wave *Wave) *Sound {
	cwave := *wave.cptr()
	res := C.LoadSoundFromWave(cwave)
	retval := newSoundFromPointer(unsafe.Pointer(&res))
	addUnloadable(retval)
	return retval
}

//Update : Update sound buffer with new data
func (sound *Sound) Update(data unsafe.Pointer, samplesCount int) {
	UpdateSound(sound, data, samplesCount)
}

//UpdateSound : Update sound buffer with new data
func UpdateSound(sound *Sound, data unsafe.Pointer, samplesCount int) {
	csound := *sound.cptr()
	C.UpdateSound(csound, data, C.int(int32(samplesCount)))
}

//Unload : Unload wave data
func (wave *Wave) Unload() {
	UnloadWave(wave)
}

//UnloadWave : Unload wave data
func UnloadWave(wave *Wave) {
	cwave := *wave.cptr()
	C.UnloadWave(cwave)
	removeUnloadable(wave)
}

//Unload : Unload sound
func (sound *Sound) Unload() {
	UnloadSound(sound)
}

//UnloadSound : Unload sound
func UnloadSound(sound *Sound) {
	csound := *sound.cptr()
	C.UnloadSound(csound)
	removeUnloadable(sound)
}

//Export : Export wave data to file
func (wave *Wave) Export(fileName string) {
	ExportWave(wave, fileName)
}

//ExportWave : Export wave data to file
func ExportWave(wave *Wave, fileName string) {
	cfileName := C.CString(fileName)
	defer C.free(unsafe.Pointer(cfileName))
	cwave := *wave.cptr()
	C.ExportWave(cwave, cfileName)
}

//ExportAsCode : Export wave sample data to code (.h)
func (wave *Wave) ExportAsCode(fileName string) {
	ExportWaveAsCode(wave, fileName)
}

//ExportWaveAsCode : Export wave sample data to code (.h)
func ExportWaveAsCode(wave *Wave, fileName string) {
	cfileName := C.CString(fileName)
	defer C.free(unsafe.Pointer(cfileName))
	cwave := *wave.cptr()
	C.ExportWaveAsCode(cwave, cfileName)
}

//Play : Play a sound
func (sound *Sound) Play() {
	PlaySound(sound)
}

//PlaySound : Play a sound
func PlaySound(sound *Sound) {
	csound := *sound.cptr()
	C.PlaySound(csound)
}

//Stop : Stop playing a sound
func (sound *Sound) Stop() {
	StopSound(sound)
}

//StopSound : Stop playing a sound
func StopSound(sound *Sound) {
	csound := *sound.cptr()
	C.StopSound(csound)
}

//Pause : Pause a sound
func (sound *Sound) Pause() {
	PauseSound(sound)
}

//PauseSound : Pause a sound
func PauseSound(sound *Sound) {
	csound := *sound.cptr()
	C.PauseSound(csound)
}

//Resume : Resume a paused sound
func (sound *Sound) Resume() {
	ResumeSound(sound)
}

//ResumeSound : Resume a paused sound
func ResumeSound(sound *Sound) {
	csound := *sound.cptr()
	C.ResumeSound(csound)
}

//PlayMulti : Play a sound (using multichannel buffer pool)
func (sound *Sound) PlayMulti() {
	PlaySoundMulti(sound)
}

//PlaySoundMulti : Play a sound (using multichannel buffer pool)
func PlaySoundMulti(sound *Sound) {
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
	return int(int32(res))
}

//IsPlaying : Check if a sound is currently playing
func (sound *Sound) IsPlaying() bool {
	return IsSoundPlaying(sound)
}

//IsSoundPlaying : Check if a sound is currently playing
func IsSoundPlaying(sound *Sound) bool {
	csound := *sound.cptr()
	res := C.IsSoundPlaying(csound)
	return bool(res)
}

//SetVolume : Set volume for a sound (1.0 is max level)
func (sound *Sound) SetVolume(volume float32) {
	SetSoundVolume(sound, volume)
}

//SetSoundVolume : Set volume for a sound (1.0 is max level)
func SetSoundVolume(sound *Sound, volume float32) {
	csound := *sound.cptr()
	C.SetSoundVolume(csound, C.float(volume))
}

//SetPitch : Set pitch for a sound (1.0 is base level)
func (sound *Sound) SetPitch(pitch float32) {
	SetSoundPitch(sound, pitch)
}

//SetSoundPitch : Set pitch for a sound (1.0 is base level)
func SetSoundPitch(sound *Sound, pitch float32) {
	csound := *sound.cptr()
	C.SetSoundPitch(csound, C.float(pitch))
}

//Format : Convert wave data to desired format
func (wave *Wave) Format(sampleRate int, sampleSize int, channels int) *Wave {
	return WaveFormat(wave, sampleRate, sampleSize, channels)
}

//WaveFormat : Convert wave data to desired format
func WaveFormat(wave *Wave, sampleRate int, sampleSize int, channels int) *Wave {
	cwave := wave.cptr()
	C.WaveFormat(cwave, C.int(int32(sampleRate)), C.int(int32(sampleSize)), C.int(int32(channels)))
	return newWaveFromPointer(unsafe.Pointer(cwave))
}

//Copy : Copy a wave to a new wave
func (wave *Wave) Copy() *Wave {
	return WaveCopy(wave)
}

//WaveCopy : Copy a wave to a new wave
func WaveCopy(wave *Wave) *Wave {
	cwave := *wave.cptr()
	res := C.WaveCopy(cwave)
	return newWaveFromPointer(unsafe.Pointer(&res))
}

//Crop : Crop a wave to defined samples range
func (wave *Wave) Crop(initSample int, finalSample int) *Wave {
	return WaveCrop(wave, initSample, finalSample)
}

//WaveCrop : Crop a wave to defined samples range
func WaveCrop(wave *Wave, initSample int, finalSample int) *Wave {
	cwave := wave.cptr()
	C.WaveCrop(cwave, C.int(int32(initSample)), C.int(int32(finalSample)))
	return newWaveFromPointer(unsafe.Pointer(cwave))
}

//GetWaveData : Get samples data from wave as a floats array
func GetWaveData(wave Wave) []float32 {

	samples := wave.SampleCount * wave.Channels
	cwave := *wave.cptr()

	res := C.GetWaveData(cwave)
	tmpslice := (*[1 << 24]*C.float)(unsafe.Pointer(res))[:samples:samples]
	defer C.free(unsafe.Pointer(res))

	gostrings := make([]float32, samples)
	for i, s := range tmpslice {
		gostrings[i] = float32(*s)
	}

	return gostrings
}

//LoadMusicStream : Load music stream from file
func LoadMusicStream(fileName string) *Music {
	cfileName := C.CString(fileName)
	defer C.free(unsafe.Pointer(cfileName))
	res := C.LoadMusicStream(cfileName)
	retval := newMusicFromPointer(unsafe.Pointer(&res))
	addUnloadable(retval)
	return retval
}

//UnloadStream : Unload music stream
func (music *Music) Unload() {
	UnloadMusicStream(music)
}

//UnloadMusicStream : Unload music stream
func UnloadMusicStream(music *Music) {
	cmusic := *music.cptr()
	C.UnloadMusicStream(cmusic)
	removeUnloadable(music)
}

//PlayStream : Start music playing
func (music *Music) PlayStream() {
	PlayMusicStream(music)
}

//PlayMusicStream : Start music playing
func PlayMusicStream(music *Music) {
	cmusic := *music.cptr()
	C.PlayMusicStream(cmusic)
}

//UpdateStream : Updates buffers for music streaming
func (music *Music) UpdateStream() {
	UpdateMusicStream(music)
}

//UpdateMusicStream : Updates buffers for music streaming
func UpdateMusicStream(music *Music) {
	cmusic := *music.cptr()
	C.UpdateMusicStream(cmusic)
}

//StopStream : Stop music playing
func (music *Music) StopStream() {
	StopMusicStream(music)
}

//StopMusicStream : Stop music playing
func StopMusicStream(music *Music) {
	cmusic := *music.cptr()
	C.StopMusicStream(cmusic)
}

//PauseStream : Pause music playing
func (music *Music) PauseStream() {
	PauseMusicStream(music)
}

//PauseMusicStream : Pause music playing
func PauseMusicStream(music *Music) {
	cmusic := *music.cptr()
	C.PauseMusicStream(cmusic)
}

//ResumeStream : Resume playing paused music
func (music *Music) ResumeStream() {
	ResumeMusicStream(music)
}

//ResumeMusicStream : Resume playing paused music
func ResumeMusicStream(music *Music) {
	cmusic := *music.cptr()
	C.ResumeMusicStream(cmusic)
}

//IsPlaying : Check if music is playing
func (music *Music) IsPlaying() bool {
	return IsMusicPlaying(music)
}

//IsMusicPlaying : Check if music is playing
func IsMusicPlaying(music *Music) bool {
	cmusic := *music.cptr()
	res := C.IsMusicPlaying(cmusic)
	return bool(res)
}

//SetVolume : Set volume for music (1.0 is max level)
func (music *Music) SetVolume(volume float32) {
	SetMusicVolume(music, volume)
}

//SetMusicVolume : Set volume for music (1.0 is max level)
func SetMusicVolume(music *Music, volume float32) {
	cmusic := *music.cptr()
	C.SetMusicVolume(cmusic, C.float(volume))
}

//SetPitch : Set pitch for a music (1.0 is base level)
func (music *Music) SetPitch(pitch float32) {
	SetMusicPitch(music, pitch)
}

//SetMusicPitch : Set pitch for a music (1.0 is base level)
func SetMusicPitch(music *Music, pitch float32) {
	cmusic := *music.cptr()
	C.SetMusicPitch(cmusic, C.float(pitch))
}

//SetLoopCount : Set music loop count (loop repeats)
func (music *Music) SetLoopCount(count int) {
	SetMusicLoopCount(music, count)
}

//SetMusicLoopCount : Set music loop count (loop repeats)
func SetMusicLoopCount(music *Music, count int) {
	cmusic := *music.cptr()
	C.SetMusicLoopCount(cmusic, C.int(int32(count)))
}

//GetTimeLength : Get music time length (in seconds)
func (music *Music) GetTimeLength() float32 {
	return GetMusicTimeLength(music)
}

//GetMusicTimeLength : Get music time length (in seconds)
func GetMusicTimeLength(music *Music) float32 {
	cmusic := *music.cptr()
	res := C.GetMusicTimeLength(cmusic)
	return float32(res)
}

//GetTimePlayed : Get current music time played (in seconds)
func (music *Music) GetTimePlayed() float32 {
	return GetMusicTimePlayed(music)
}

//GetMusicTimePlayed : Get current music time played (in seconds)
func GetMusicTimePlayed(music *Music) float32 {
	cmusic := *music.cptr()
	res := C.GetMusicTimePlayed(cmusic)
	return float32(res)
}

//InitAudioStream : Init audio stream (to stream raw audio pcm data)
func InitAudioStream(sampleRate uint32, sampleSize uint32, channels uint32) *AudioStream {
	res := C.InitAudioStream(C.uint(sampleRate), C.uint(sampleSize), C.uint(channels))
	return newAudioStreamFromPointer(unsafe.Pointer(&res))
}

//Update : Update audio stream buffers with data
func (stream *AudioStream) Update(data unsafe.Pointer, samplesCount int) {
	UpdateAudioStream(stream, data, samplesCount)
}

//UpdateAudioStream : Update audio stream buffers with data
func UpdateAudioStream(stream *AudioStream, data unsafe.Pointer, samplesCount int) {
	cstream := *stream.cptr()
	C.UpdateAudioStream(cstream, data, C.int(int32(samplesCount)))
}

//Close : Close audio stream and free memory
func (stream *AudioStream) Close() {
	CloseAudioStream(stream)
}

//CloseAudioStream : Close audio stream and free memory
func CloseAudioStream(stream *AudioStream) {
	cstream := *stream.cptr()
	C.CloseAudioStream(cstream)
}

//IsProcessed : Check if any audio stream buffers requires refill
func (stream *AudioStream) IsProcessed() bool {
	return IsAudioStreamProcessed(stream)
}

//IsAudioStreamProcessed : Check if any audio stream buffers requires refill
func IsAudioStreamProcessed(stream *AudioStream) bool {
	cstream := *stream.cptr()
	res := C.IsAudioStreamProcessed(cstream)
	return bool(res)
}

//Play : Play audio stream
func (stream *AudioStream) Play() {
	PlayAudioStream(stream)
}

//PlayAudioStream : Play audio stream
func PlayAudioStream(stream *AudioStream) {
	cstream := *stream.cptr()
	C.PlayAudioStream(cstream)
}

//Pause : Pause audio stream
func (stream *AudioStream) Pause() {
	PauseAudioStream(stream)
}

//PauseAudioStream : Pause audio stream
func PauseAudioStream(stream *AudioStream) {
	cstream := *stream.cptr()
	C.PauseAudioStream(cstream)
}

//Resume : Resume audio stream
func (stream *AudioStream) Resume() {
	ResumeAudioStream(stream)
}

//ResumeAudioStream : Resume audio stream
func ResumeAudioStream(stream *AudioStream) {
	cstream := *stream.cptr()
	C.ResumeAudioStream(cstream)
}

//IsPlaying : Check if audio stream is playing
func (stream *AudioStream) IsPlaying() bool {
	return IsAudioStreamPlaying(stream)
}

//IsAudioStreamPlaying : Check if audio stream is playing
func IsAudioStreamPlaying(stream *AudioStream) bool {
	cstream := *stream.cptr()
	res := C.IsAudioStreamPlaying(cstream)
	return bool(res)
}

//Stop : Stop audio stream
func (stream *AudioStream) Stop() {
	StopAudioStream(stream)
}

//StopAudioStream : Stop audio stream
func StopAudioStream(stream *AudioStream) {
	cstream := *stream.cptr()
	C.StopAudioStream(cstream)
}

//SetVolume : Set volume for audio stream (1.0 is max level)
func (stream *AudioStream) SetVolume(volume float32) {
	SetAudioStreamVolume(stream, volume)
}

//SetAudioStreamVolume : Set volume for audio stream (1.0 is max level)
func SetAudioStreamVolume(stream *AudioStream, volume float32) {
	cstream := *stream.cptr()
	C.SetAudioStreamVolume(cstream, C.float(volume))
}

//SetPitch : Set pitch for audio stream (1.0 is base level)
func (stream *AudioStream) SetPitch(pitch float32) {
	SetAudioStreamPitch(stream, pitch)
}

//SetAudioStreamPitch : Set pitch for audio stream (1.0 is base level)
func SetAudioStreamPitch(stream *AudioStream, pitch float32) {
	cstream := *stream.cptr()
	C.SetAudioStreamPitch(cstream, C.float(pitch))
}
