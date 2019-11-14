//included with gui enable because raygui.h is weird
func newGuiTextBoxStateFromPointer(ptr unsafe.Pointer) GuiTextBoxState {
	return *(*GuiTextBoxState)(ptr)
}

func (w *GuiTextBoxState) cptr() *C.GuiTextBoxState {
	return (*C.GuiTextBoxState)(unsafe.Pointer(w))
}

//GuiEnable : Enable gui controls (global state)
func GuiEnable() {
	C.GuiEnable()
	guiEnabled = true
}
