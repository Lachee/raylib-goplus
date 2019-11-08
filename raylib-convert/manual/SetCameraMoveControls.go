//SetCameraMoveControls : Set camera move controls (1st person and 3rd person cameras)
func SetCameraMoveControls(frontKey Key, backKey Key, rightKey Key, leftKey Key, upKey Key, downKey Key) {
	C.SetCameraMoveControls(C.int(frontKey), C.int(backKey), C.int(rightKey), C.int(leftKey), C.int(upKey), C.int(downKey))
}
