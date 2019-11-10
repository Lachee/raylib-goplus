package raylib

/*
//Generated 2019-11-10T19:06:35+11:00
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
	csound := *sound.cptr()
	C.UpdateSound(csound, data, C.int(int32(samplesCount)))
}

//UpdateSound : Update sound buffer with new data
//Recommended to use sound.Update(data, samplesCount) instead
func UpdateSound(sound *Sound, data unsafe.Pointer, samplesCount int) {
	sound.Update(data, samplesCount)
}

//Unload : Unload wave data
func (wave *Wave) Unload() {
	cwave := *wave.cptr()
	C.UnloadWave(cwave)
	removeUnloadable(wave)
}

//UnloadWave : Unload wave data
//Recommended to use wave.Unload() instead
func UnloadWave(wave *Wave) {
	wave.Unload()
}

//Unload : Unload sound
func (sound *Sound) Unload() {
	csound := *sound.cptr()
	C.UnloadSound(csound)
	removeUnloadable(sound)
}

//UnloadSound : Unload sound
//Recommended to use sound.Unload() instead
func UnloadSound(sound *Sound) {
	sound.Unload()
}

//Export : Export wave data to file
func (wave *Wave) Export(fileName string) {
	cfileName := C.CString(fileName)
	defer C.free(unsafe.Pointer(cfileName))
	cwave := *wave.cptr()
	C.ExportWave(cwave, cfileName)
}

//ExportWave : Export wave data to file
//Recommended to use wave.Export(fileName) instead
func ExportWave(wave *Wave, fileName string) {
	wave.Export(fileName)
}

//ExportAsCode : Export wave sample data to code (.h)
func (wave *Wave) ExportAsCode(fileName string) {
	cfileName := C.CString(fileName)
	defer C.free(unsafe.Pointer(cfileName))
	cwave := *wave.cptr()
	C.ExportWaveAsCode(cwave, cfileName)
}

//ExportWaveAsCode : Export wave sample data to code (.h)
//Recommended to use wave.ExportAsCode(fileName) instead
func ExportWaveAsCode(wave *Wave, fileName string) {
	wave.ExportAsCode(fileName)
}

//Play : Play a sound
func (sound *Sound) Play() {
	csound := *sound.cptr()
	C.PlaySound(csound)
}

//PlaySound : Play a sound
//Recommended to use sound.Play() instead
func PlaySound(sound *Sound) {
	sound.Play()
}

//Stop : Stop playing a sound
func (sound *Sound) Stop() {
	csound := *sound.cptr()
	C.StopSound(csound)
}

//StopSound : Stop playing a sound
//Recommended to use sound.Stop() instead
func StopSound(sound *Sound) {
	sound.Stop()
}

//Pause : Pause a sound
func (sound *Sound) Pause() {
	csound := *sound.cptr()
	C.PauseSound(csound)
}

//PauseSound : Pause a sound
//Recommended to use sound.Pause() instead
func PauseSound(sound *Sound) {
	sound.Pause()
}

//Resume : Resume a paused sound
func (sound *Sound) Resume() {
	csound := *sound.cptr()
	C.ResumeSound(csound)
}

//ResumeSound : Resume a paused sound
//Recommended to use sound.Resume() instead
func ResumeSound(sound *Sound) {
	sound.Resume()
}

//PlayMulti : Play a sound (using multichannel buffer pool)
func (sound *Sound) PlayMulti() {
	csound := *sound.cptr()
	C.PlaySoundMulti(csound)
}

//PlaySoundMulti : Play a sound (using multichannel buffer pool)
//Recommended to use sound.PlayMulti() instead
func PlaySoundMulti(sound *Sound) {
	sound.PlayMulti()
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
	csound := *sound.cptr()
	res := C.IsSoundPlaying(csound)
	return bool(res)
}

//IsSoundPlaying : Check if a sound is currently playing
//Recommended to use sound.IsPlaying() instead
func IsSoundPlaying(sound *Sound) bool {
	return sound.IsPlaying()
}

//SetVolume : Set volume for a sound (1.0 is max level)
func (sound *Sound) SetVolume(volume float32) {
	csound := *sound.cptr()
	C.SetSoundVolume(csound, C.float(volume))
}

//SetSoundVolume : Set volume for a sound (1.0 is max level)
//Recommended to use sound.SetVolume(volume) instead
func SetSoundVolume(sound *Sound, volume float32) {
	sound.SetVolume(volume)
}

//SetPitch : Set pitch for a sound (1.0 is base level)
func (sound *Sound) SetPitch(pitch float32) {
	csound := *sound.cptr()
	C.SetSoundPitch(csound, C.float(pitch))
}

//SetSoundPitch : Set pitch for a sound (1.0 is base level)
//Recommended to use sound.SetPitch(pitch) instead
func SetSoundPitch(sound *Sound, pitch float32) {
	sound.SetPitch(pitch)
}

//Format : Convert wave data to desired format
func (wave *Wave) Format(sampleRate int, sampleSize int, channels int) {
	cwave := wave.cptr()
	C.WaveFormat(cwave, C.int(int32(sampleRate)), C.int(int32(sampleSize)), C.int(int32(channels)))
}

//WaveFormat : Convert wave data to desired format
//Recommended to use wave.Format(sampleRate, sampleSize, channels) instead
func WaveFormat(wave *Wave, sampleRate int, sampleSize int, channels int) {
	wave.Format(sampleRate, sampleSize, channels)
}

//Copy : Copy a wave to a new wave
func (wave *Wave) Copy() *Wave {
	cwave := *wave.cptr()
	res := C.WaveCopy(cwave)
	retval := newWaveFromPointer(unsafe.Pointer(&res))
	addUnloadable(retval)
	return retval
}

//WaveCopy : Copy a wave to a new wave
//Recommended to use wave.Copy() instead
func WaveCopy(wave *Wave) *Wave {
	return wave.Copy()
}

//Crop : Crop a wave to defined samples range
func (wave *Wave) Crop(initSample int, finalSample int) {
	cwave := wave.cptr()
	C.WaveCrop(cwave, C.int(int32(initSample)), C.int(int32(finalSample)))
}

//WaveCrop : Crop a wave to defined samples range
//Recommended to use wave.Crop(initSample, finalSample) instead
func WaveCrop(wave *Wave, initSample int, finalSample int) {
	wave.Crop(initSample, finalSample)
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
	cmusic := *music.cptr()
	C.PlayMusicStream(cmusic)
}

//PlayMusicStream : Start music playing
//Recommended to use music.PlayStream() instead
func PlayMusicStream(music *Music) {
	music.PlayStream()
}

//UpdateStream : Updates buffers for music streaming
func (music *Music) UpdateStream() {
	cmusic := *music.cptr()
	C.UpdateMusicStream(cmusic)
}

//UpdateMusicStream : Updates buffers for music streaming
//Recommended to use music.UpdateStream() instead
func UpdateMusicStream(music *Music) {
	music.UpdateStream()
}

//StopStream : Stop music playing
func (music *Music) StopStream() {
	cmusic := *music.cptr()
	C.StopMusicStream(cmusic)
}

//StopMusicStream : Stop music playing
//Recommended to use music.StopStream() instead
func StopMusicStream(music *Music) {
	music.StopStream()
}

//PauseStream : Pause music playing
func (music *Music) PauseStream() {
	cmusic := *music.cptr()
	C.PauseMusicStream(cmusic)
}

//PauseMusicStream : Pause music playing
//Recommended to use music.PauseStream() instead
func PauseMusicStream(music *Music) {
	music.PauseStream()
}

//ResumeStream : Resume playing paused music
func (music *Music) ResumeStream() {
	cmusic := *music.cptr()
	C.ResumeMusicStream(cmusic)
}

//ResumeMusicStream : Resume playing paused music
//Recommended to use music.ResumeStream() instead
func ResumeMusicStream(music *Music) {
	music.ResumeStream()
}

//IsPlaying : Check if music is playing
func (music *Music) IsPlaying() bool {
	cmusic := *music.cptr()
	res := C.IsMusicPlaying(cmusic)
	return bool(res)
}

//IsMusicPlaying : Check if music is playing
//Recommended to use music.IsPlaying() instead
func IsMusicPlaying(music *Music) bool {
	return music.IsPlaying()
}

//SetVolume : Set volume for music (1.0 is max level)
func (music *Music) SetVolume(volume float32) {
	cmusic := *music.cptr()
	C.SetMusicVolume(cmusic, C.float(volume))
}

//SetMusicVolume : Set volume for music (1.0 is max level)
//Recommended to use music.SetVolume(volume) instead
func SetMusicVolume(music *Music, volume float32) {
	music.SetVolume(volume)
}

//SetPitch : Set pitch for a music (1.0 is base level)
func (music *Music) SetPitch(pitch float32) {
	cmusic := *music.cptr()
	C.SetMusicPitch(cmusic, C.float(pitch))
}

//SetMusicPitch : Set pitch for a music (1.0 is base level)
//Recommended to use music.SetPitch(pitch) instead
func SetMusicPitch(music *Music, pitch float32) {
	music.SetPitch(pitch)
}

//SetLoopCount : Set music loop count (loop repeats)
func (music *Music) SetLoopCount(count int) {
	cmusic := *music.cptr()
	C.SetMusicLoopCount(cmusic, C.int(int32(count)))
}

//SetMusicLoopCount : Set music loop count (loop repeats)
//Recommended to use music.SetLoopCount(count) instead
func SetMusicLoopCount(music *Music, count int) {
	music.SetLoopCount(count)
}

//GetTimeLength : Get music time length (in seconds)
func (music *Music) GetTimeLength() float32 {
	cmusic := *music.cptr()
	res := C.GetMusicTimeLength(cmusic)
	return float32(res)
}

//GetMusicTimeLength : Get music time length (in seconds)
//Recommended to use music.GetTimeLength() instead
func GetMusicTimeLength(music *Music) float32 {
	return music.GetTimeLength()
}

//GetTimePlayed : Get current music time played (in seconds)
func (music *Music) GetTimePlayed() float32 {
	cmusic := *music.cptr()
	res := C.GetMusicTimePlayed(cmusic)
	return float32(res)
}

//GetMusicTimePlayed : Get current music time played (in seconds)
//Recommended to use music.GetTimePlayed() instead
func GetMusicTimePlayed(music *Music) float32 {
	return music.GetTimePlayed()
}

//InitAudioStream : Init audio stream (to stream raw audio pcm data)
func InitAudioStream(sampleRate uint32, sampleSize uint32, channels uint32) *AudioStream {
	res := C.InitAudioStream(C.uint(sampleRate), C.uint(sampleSize), C.uint(channels))
	retval := newAudioStreamFromPointer(unsafe.Pointer(&res))
	addUnloadable(retval)
	return retval
}

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

//IsProcessed : Check if any audio stream buffers requires refill
func (stream *AudioStream) IsProcessed() bool {
	cstream := *stream.cptr()
	res := C.IsAudioStreamProcessed(cstream)
	return bool(res)
}

//IsAudioStreamProcessed : Check if any audio stream buffers requires refill
//Recommended to use stream.IsProcessed() instead
func IsAudioStreamProcessed(stream *AudioStream) bool {
	return stream.IsProcessed()
}

//Play : Play audio stream
func (stream *AudioStream) Play() {
	cstream := *stream.cptr()
	C.PlayAudioStream(cstream)
}

//PlayAudioStream : Play audio stream
//Recommended to use stream.Play() instead
func PlayAudioStream(stream *AudioStream) {
	stream.Play()
}

//Pause : Pause audio stream
func (stream *AudioStream) Pause() {
	cstream := *stream.cptr()
	C.PauseAudioStream(cstream)
}

//PauseAudioStream : Pause audio stream
//Recommended to use stream.Pause() instead
func PauseAudioStream(stream *AudioStream) {
	stream.Pause()
}

//Resume : Resume audio stream
func (stream *AudioStream) Resume() {
	cstream := *stream.cptr()
	C.ResumeAudioStream(cstream)
}

//ResumeAudioStream : Resume audio stream
//Recommended to use stream.Resume() instead
func ResumeAudioStream(stream *AudioStream) {
	stream.Resume()
}

//IsPlaying : Check if audio stream is playing
func (stream *AudioStream) IsPlaying() bool {
	cstream := *stream.cptr()
	res := C.IsAudioStreamPlaying(cstream)
	return bool(res)
}

//IsAudioStreamPlaying : Check if audio stream is playing
//Recommended to use stream.IsPlaying() instead
func IsAudioStreamPlaying(stream *AudioStream) bool {
	return stream.IsPlaying()
}

//Stop : Stop audio stream
func (stream *AudioStream) Stop() {
	cstream := *stream.cptr()
	C.StopAudioStream(cstream)
}

//StopAudioStream : Stop audio stream
//Recommended to use stream.Stop() instead
func StopAudioStream(stream *AudioStream) {
	stream.Stop()
}

//SetVolume : Set volume for audio stream (1.0 is max level)
func (stream *AudioStream) SetVolume(volume float32) {
	cstream := *stream.cptr()
	C.SetAudioStreamVolume(cstream, C.float(volume))
}

//SetAudioStreamVolume : Set volume for audio stream (1.0 is max level)
//Recommended to use stream.SetVolume(volume) instead
func SetAudioStreamVolume(stream *AudioStream, volume float32) {
	stream.SetVolume(volume)
}

//SetPitch : Set pitch for audio stream (1.0 is base level)
func (stream *AudioStream) SetPitch(pitch float32) {
	cstream := *stream.cptr()
	C.SetAudioStreamPitch(cstream, C.float(pitch))
}

//SetAudioStreamPitch : Set pitch for audio stream (1.0 is base level)
//Recommended to use stream.SetPitch(pitch) instead
func SetAudioStreamPitch(stream *AudioStream, pitch float32) {
	stream.SetPitch(pitch)
}
