//GuiSetStyle : Set one style property
func GuiSetStyle(control GuiControl, property GuiProperty, value int) {
	C.GuiSetStyle(C.int(control), C.int(property), C.int(value))
}