package raylib

/*
#include "raylib.h"
#include <stdlib.h>
*/
import "C"
import "unsafe"

//Wave defines audio wave data
type Wave struct {
	SampleCount uint32
	SampleRate  uint32
	SampleSize  uint32
	Channels    uint32
	data        unsafe.Pointer
}

func newWaveFromPointer(ptr unsafe.Pointer) *Wave {
	return (*Wave)(ptr)
}

func (w *Wave) cptr() *C.Wave {
	return (*C.Wave)(unsafe.Pointer(w))
}

//Sound source type
type Sound struct {
	SampleCount uint32
	Stream      AudioStream
}

func newSoundFromPointer(ptr unsafe.Pointer) *Sound {
	return (*Sound)(ptr)
}

func (s *Sound) cptr() *C.Sound {
	return (*C.Sound)(unsafe.Pointer(s))
}

//AudioStream can be used to create custom audio streams.
//Note that Buffer is an unsafe.Pointer and is a C stream.
type AudioStream struct {
	SampleRate uint32
	SampleSize uint32
	Channels   uint32
	Buffer     unsafe.Pointer
}

func newAudioStreamFromPointer(ptr unsafe.Pointer) *AudioStream {
	return (*AudioStream)(ptr)
}

func (as *AudioStream) cptr() *C.AudioStream {
	return (*C.AudioStream)(unsafe.Pointer(as))
}

//Music stream type. Anything longer than ~10 seconds should be streamed.
type Music struct {
	CtxType     int32
	CtxData     unsafe.Pointer
	SampleCount uint32
	LoopCount   uint32
	Stream      AudioStream
}

func newMusicFromPointer(ptr unsafe.Pointer) *Music {
	return (*Music)(ptr)
}

func (music *Music) cptr() *C.Music {
	return (*C.Music)(unsafe.Pointer(music))
}
