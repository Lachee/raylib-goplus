//GetDroppedFiles : Get dropped files names (memory should be freed)
func GetDroppedFiles() []string {
	ccount := C.int(count)
	res := C.GetDroppedFiles(&ccount)
		
	tmpslice := (*[1 << 24]*C.char)(unsafe.Pointer(res))[:count:count]
	gostrings := make([]string, count)
	for i, s := range tmpslice {
		gostrings[i] = C.GoString(s)
	}
	
	return gostrings
}