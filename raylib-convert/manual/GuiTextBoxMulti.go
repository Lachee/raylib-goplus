//GuiTextBox : Text Box control, updates input text
func GuiTextBoxMulti(bounds Rectangle, text string, maxCharacters int, editMode bool) (bool, string) {

	//Allocate a new chunk of memory to put the characters in.
	// Then copy all the characters across. If there is not enough characters, fill with 0.
	data := make([]C.char, maxCharacters)
	for i := 0; i < maxCharacters; i++ {
		if len(text) <= i {
			data[i] = C.char(0)
		} else {
			data[i] = C.char(text[i])
		}
	}

	//Prepare an unsafe pointer to that chunk. C will move it arround
	ctext := (*C.char)(unsafe.Pointer(&data[0]))
	cbounds := *bounds.cptr()

	//Create the textbox
	res := C.GuiTextBoxMulti(cbounds, ctext, C.int(int32(maxCharacters)), C.bool(editMode))

	//Return the result.
	return bool(res), C.GoString(ctext)
}
