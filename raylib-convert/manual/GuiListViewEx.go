//GuiListViewEx :  List View with extended parameters
func GuiListViewEx(bounds Rectangle, text string, count int, focus int, scrollIndex int, active int) (int, string) {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	cbounds := *bounds.cptr()
	res := C.GuiListViewEx(cbounds, &ctext, C.int(count), C.int(focus), C.int(scrollIndex), C.int(active))
	return int(res), C.GoString(ctext)
}