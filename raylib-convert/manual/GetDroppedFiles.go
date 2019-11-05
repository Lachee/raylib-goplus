//GetDroppedFiles : Get dropped files names (memory should be freed)
func GetDroppedFiles() []string {
	ccount := C.int(0)
	res := C.GetDroppedFiles(&ccount)
	count := int(ccount)

	tmpslice := (*[1 << 24]*C.char)(unsafe.Pointer(res))[:count:count]
	gostrings := make([]string, count)
	for i, s := range tmpslice {
		gostrings[i] = C.GoString(s)
	}

	return gostrings
}