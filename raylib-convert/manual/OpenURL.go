//OpenURL opens a URL in the system browser.
func OpenURL(url string) error {
	switch runtime.GOOS {
	case "linux":
		return exec.Command("xdg-open", url).Start()
	case "windows":
		return exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		return exec.Command("open", url).Start()
	default:
		cstring := C.CString(url)
		defer C.free(unsafe.Pointer(cstring))
		C.OpenURL(cstring)
		return nil
	}	
}