//DrawTextRecEx : Draw text using font inside rectangle limits with support for text selection
func DrawTextRecEx(font Font, text string, rec Rectangle, fontSize float32, spacing float32, wordWrap bool, tint Color, selectStart int, selectLength int, selectText Color, selectBack Color) {
	cselectBack := *selectBack.cptr()
	cselectText := *selectText.cptr()
	ctint := *tint.cptr()
	crec := *rec.cptr()
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	cfont := *font.cptr()
	C.DrawTextRecEx(cfont, ctext, crec, C.float(fontSize), C.float(spacing), C.bool(wordWrap), ctint, C.int(selectStart), C.int(selectLength), cselectText, cselectBack)
}