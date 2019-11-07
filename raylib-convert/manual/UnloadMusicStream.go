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