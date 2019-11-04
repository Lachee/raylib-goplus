//GuiGetStyle : Get one style property
func GuiGetStyle(control GuiControl, property GuiProperty) int {
	res := C.GuiGetStyle(C.int(control), C.int(property))
	return int(res)
}
