//UpdateVrTracking : Update VR tracking (position and orientation) and camera
func UpdateVrTracking(camera *Camera) {
	ccamera := camera.cptr()
	C.UpdateVrTracking(ccamera)
}
