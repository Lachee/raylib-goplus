//GuiTextInputBox : Text Input Box control, ask for text
func GuiTextInputBox(bounds Rectangle, title string, message string, buttons string, text string) (int, string) {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	cbuttons := C.CString(buttons)
	defer C.free(unsafe.Pointer(cbuttons))
	cmessage := C.CString(message)
	defer C.free(unsafe.Pointer(cmessage))
	ctitle := C.CString(title)
	defer C.free(unsafe.Pointer(ctitle))
	cbounds := *bounds.cptr()
	res := C.GuiTextInputBox(cbounds, ctitle, cmessage, cbuttons, ctext)
	return int(res), C.GoString(text)
}