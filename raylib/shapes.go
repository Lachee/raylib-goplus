package raylib

//DrawRectangleGradientV : Draw a vertical-gradient-filled rectangle
func DrawRectangleGradientVRec(rect Rectangle, color1 Color, color2 Color) {
	DrawRectangleGradientEx(rect, color1, color2, color2, color1)
}

//DrawRectangleGradientH : Draw a horizontal-gradient-filled rectangle
func DrawRectangleGradientHRec(rect Rectangle, color1 Color, color2 Color) {
	DrawRectangleGradientEx(rect, color1, color1, color2, color2)
}
