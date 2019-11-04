//GuiListViewEx : List View with extended parameters
func GuiListViewEx(bounds Rectangle, text []string, count int, focus int, scrollIndex int, active int) (int, int, int) {
	cscrollIndex := C.int(scrollIndex)
	cfocus := C.int(focus)
	cbounds := *bounds.cptr()
	
	//Copies the string into an array in C memory
	cargs := C.makeCharArray(C.int(len(text)))
	defer C.freeCharArray(cargs, C.int(len(text)))
	for i, s := range text {
        C.setArrayString(cargs, C.CString(s), C.int(i))
	}

	//Create the thingy
	res := C.GuiListViewEx(cbounds, cargs, C.int(count), &cfocus, &cscrollIndex, C.int(active))
	return int(res), int(cfocus), int(cscrollIndex)
}
