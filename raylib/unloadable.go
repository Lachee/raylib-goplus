package raylib

//Unloadable is any object that has a Unload function and needs to be freed
// when it has finished being used.
type Unloadable interface {
	Unload()
}

var unloadingAll bool = false
var unloadables []Unloadable = make([]Unloadable, 100)

//TODO: Fix this
func finalizeUnloadables(unlds *[]Unloadable) {
	TraceLog(LogInfo, "[UNLOAD] Finalizing Unloadables")
	UnloadAll()
}

//RegisterUnloadable registers an unloadable to the slice
// This is called on Load functions
func RegisterUnloadable(unloadable Unloadable) {
	TraceLog(LogTrace, "[UNLOAD] New unloadable created")
	unloadables = append(unloadables, unloadable)
}

//UnregisterUnloadable unregisters an unloadable to the slice
// This is called on Unload functions
// This does not remove from the slice if unloadingAll is true (as that will clear post)
func UnregisterUnloadable(unloadable Unloadable) {
	if !unloadingAll {
		for i, u := range unloadables {
			if u == unloadable {
				unloadables[i] = unloadables[len(unloadables)-1]
				unloadables[len(unloadables)-1] = nil
				unloadables = unloadables[:len(unloadables)-1]
				TraceLog(LogTrace, "[UNLOAD] Removed Unloadable")
				break
			}
		}
	}
}

//UnloadAll clears all unloadables that have been recorded.
// NOTE: Not everything maybe included in this list and it is experimental feature.
// 			 Please unload these objects when you are not using them anyways.
func UnloadAll() {
	TraceLog(LogInfo, "[UNLOAD] Unloading all unloadables: ", len(unloadables))

	//Count the tally
	tally := 0

	//Unload everything
	unloadingAll = true
	for _, ul := range unloadables {
		if ul != nil {
			ul.Unload()
			tally++
		}
	}
	unloadingAll = false

	//Clear the unloadables
	unloadables = unloadables[:0]
	TraceLog(LogInfo, "[UNLOAD] Unloaded ", tally)
}
