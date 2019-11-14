//GuiDisable : Disable gui controls (global state)
func GuiDisable() {
	C.GuiDisable()
	guiEnabled = false
}