//GuiUnlock : Unlock gui controls (global state)
func GuiUnlock() {
	C.GuiUnlock()
	guiLocked = false
}