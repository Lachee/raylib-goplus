//GetPixelDataSize : Get pixel data size in bytes (image or texture)
func GetPixelDataSize(width int, height int, format PixelFormat) int {
	res := C.GetPixelDataSize(C.int(int32(width)), C.int(int32(height)), C.int(format))
	return int(int32(res))
}