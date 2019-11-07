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
	TraceLog(LogInfo, "Finalizing Unloadables")
	UnloadAll()
}

//addUnloadable registers an unloadable to the slice
// This is called on Load functions
func addUnloadable(unloadable Unloadable) {
	TraceLog(LogTrace, "New unloadable created")
	unloadables = append(unloadables, unloadable)
}

//removeUnloadable unregisters an unloadable to the slice
// This is called on Unload functions
// This does not remove from the slice if unloadingAll is true (as that will clear post)
func removeUnloadable(unloadable Unloadable) {
	if !unloadingAll {
		TraceLog(LogTrace, "Remocing unloadable")
		for i, u := range unloadables {
			if u == unloadable {
				copy(unloadables[i:], unloadables[i+1:])
				break
			}
		}
	}
}

//UnloadAll clears all unloadables that have been recorded.
// NOTE: Not everything maybe included in this list and it is experimental feature.
// 			 Please unload these objects when you are not using them anyways.
func UnloadAll() {
	TraceLog(LogInfo, "Unloading all unloadables")

	//Unload everything
	unloadingAll = true
	for _, ul := range unloadables {
		if ul != nil {
			ul.Unload()
		}
	}
	unloadingAll = false

	//Clear the unloadables
	unloadables = unloadables[:0]
}
