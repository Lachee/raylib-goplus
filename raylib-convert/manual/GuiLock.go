//GuiLock : Lock gui controls (global state)
func GuiLock() {
	C.GuiLock()
	guiLocked = true
}