//GetGestureDetected : Get latest detected gesture
func GetGestureDetected() GestureType {
	res := C.GetGestureDetected()
	return GestureType(res)
}