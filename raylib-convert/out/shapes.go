package raylib

//DrawPixel : Draw a pixel
func DrawPixel(posX int, posY int, color Color) {
	ccolor := *color.cptr()
	C.DrawPixel(C.int(posX), C.int(posY), ccolor)
}

//DrawPixelV : Draw a pixel (Vector version)
func DrawPixelV(position Vector2, color Color) {
	ccolor := *color.cptr()
	cposition := *position.cptr()
	C.DrawPixelV(cposition, ccolor)
}

//DrawLine : Draw a line
func DrawLine(startPosX int, startPosY int, endPosX int, endPosY int, color Color) {
	ccolor := *color.cptr()
	C.DrawLine(C.int(startPosX), C.int(startPosY), C.int(endPosX), C.int(endPosY), ccolor)
}

//DrawLineV : Draw a line (Vector version)
func DrawLineV(startPos Vector2, endPos Vector2, color Color) {
	ccolor := *color.cptr()
	cendPos := *endPos.cptr()
	cstartPos := *startPos.cptr()
	C.DrawLineV(cstartPos, cendPos, ccolor)
}

//DrawLineEx : Draw a line defining thickness
func DrawLineEx(startPos Vector2, endPos Vector2, thick float32, color Color) {
	ccolor := *color.cptr()
	cendPos := *endPos.cptr()
	cstartPos := *startPos.cptr()
	C.DrawLineEx(cstartPos, cendPos, C.float(thick), ccolor)
}

//DrawLineBezier : Draw a line using cubic-bezier curves in-out
func DrawLineBezier(startPos Vector2, endPos Vector2, thick float32, color Color) {
	ccolor := *color.cptr()
	cendPos := *endPos.cptr()
	cstartPos := *startPos.cptr()
	C.DrawLineBezier(cstartPos, cendPos, C.float(thick), ccolor)
}

//DrawLineStrip : Draw lines sequence
func DrawLineStrip(points Vector2, numPoints int, color Color) Vector2 {
	ccolor := *color.cptr()
	cpoints := points.cptr()
	C.DrawLineStrip(&cpoints, C.int(numPoints), ccolor)
	return newVector2FromPointer(unsafe.Pointer(&cpoints))
}

//DrawCircle : Draw a color-filled circle
func DrawCircle(centerX int, centerY int, radius float32, color Color) {
	ccolor := *color.cptr()
	C.DrawCircle(C.int(centerX), C.int(centerY), C.float(radius), ccolor)
}

//DrawCircleSector : Draw a piece of a circle
func DrawCircleSector(center Vector2, radius float32, startAngle int, endAngle int, segments int, color Color) {
	ccolor := *color.cptr()
	ccenter := *center.cptr()
	C.DrawCircleSector(ccenter, C.float(radius), C.int(startAngle), C.int(endAngle), C.int(segments), ccolor)
}

//DrawCircleSectorLines : Draw circle sector outline
func DrawCircleSectorLines(center Vector2, radius float32, startAngle int, endAngle int, segments int, color Color) {
	ccolor := *color.cptr()
	ccenter := *center.cptr()
	C.DrawCircleSectorLines(ccenter, C.float(radius), C.int(startAngle), C.int(endAngle), C.int(segments), ccolor)
}

//DrawCircleGradient : Draw a gradient-filled circle
func DrawCircleGradient(centerX int, centerY int, radius float32, color1 Color, color2 Color) {
	ccolor2 := *color2.cptr()
	ccolor1 := *color1.cptr()
	C.DrawCircleGradient(C.int(centerX), C.int(centerY), C.float(radius), ccolor1, ccolor2)
}

//DrawCircleV : Draw a color-filled circle (Vector version)
func DrawCircleV(center Vector2, radius float32, color Color) {
	ccolor := *color.cptr()
	ccenter := *center.cptr()
	C.DrawCircleV(ccenter, C.float(radius), ccolor)
}

//DrawCircleLines : Draw circle outline
func DrawCircleLines(centerX int, centerY int, radius float32, color Color) {
	ccolor := *color.cptr()
	C.DrawCircleLines(C.int(centerX), C.int(centerY), C.float(radius), ccolor)
}

//DrawRing : Draw ring
func DrawRing(center Vector2, innerRadius float32, outerRadius float32, startAngle int, endAngle int, segments int, color Color) {
	ccolor := *color.cptr()
	ccenter := *center.cptr()
	C.DrawRing(ccenter, C.float(innerRadius), C.float(outerRadius), C.int(startAngle), C.int(endAngle), C.int(segments), ccolor)
}

//DrawRingLines : Draw ring outline
func DrawRingLines(center Vector2, innerRadius float32, outerRadius float32, startAngle int, endAngle int, segments int, color Color) {
	ccolor := *color.cptr()
	ccenter := *center.cptr()
	C.DrawRingLines(ccenter, C.float(innerRadius), C.float(outerRadius), C.int(startAngle), C.int(endAngle), C.int(segments), ccolor)
}

//DrawRectangle : Draw a color-filled rectangle
func DrawRectangle(posX int, posY int, width int, height int, color Color) {
	ccolor := *color.cptr()
	C.DrawRectangle(C.int(posX), C.int(posY), C.int(width), C.int(height), ccolor)
}

//DrawRectangleV : Draw a color-filled rectangle (Vector version)
func DrawRectangleV(position Vector2, size Vector2, color Color) {
	ccolor := *color.cptr()
	csize := *size.cptr()
	cposition := *position.cptr()
	C.DrawRectangleV(cposition, csize, ccolor)
}

//DrawRectangleRec : Draw a color-filled rectangle
func DrawRectangleRec(rec Rectangle, color Color) {
	ccolor := *color.cptr()
	crec := *rec.cptr()
	C.DrawRectangleRec(crec, ccolor)
}

//DrawRectanglePro : Draw a color-filled rectangle with pro parameters
func DrawRectanglePro(rec Rectangle, origin Vector2, rotation float32, color Color) {
	ccolor := *color.cptr()
	corigin := *origin.cptr()
	crec := *rec.cptr()
	C.DrawRectanglePro(crec, corigin, C.float(rotation), ccolor)
}

//DrawRectangleGradientV : Draw a vertical-gradient-filled rectangle
func DrawRectangleGradientV(posX int, posY int, width int, height int, color1 Color, color2 Color) {
	ccolor2 := *color2.cptr()
	ccolor1 := *color1.cptr()
	C.DrawRectangleGradientV(C.int(posX), C.int(posY), C.int(width), C.int(height), ccolor1, ccolor2)
}

//DrawRectangleGradientH : Draw a horizontal-gradient-filled rectangle
func DrawRectangleGradientH(posX int, posY int, width int, height int, color1 Color, color2 Color) {
	ccolor2 := *color2.cptr()
	ccolor1 := *color1.cptr()
	C.DrawRectangleGradientH(C.int(posX), C.int(posY), C.int(width), C.int(height), ccolor1, ccolor2)
}

//DrawRectangleGradientEx : Draw a gradient-filled rectangle with custom vertex colors
func DrawRectangleGradientEx(rec Rectangle, col1 Color, col2 Color, col3 Color, col4 Color) {
	ccol4 := *col4.cptr()
	ccol3 := *col3.cptr()
	ccol2 := *col2.cptr()
	ccol1 := *col1.cptr()
	crec := *rec.cptr()
	C.DrawRectangleGradientEx(crec, ccol1, ccol2, ccol3, ccol4)
}

//DrawRectangleLines : Draw rectangle outline
func DrawRectangleLines(posX int, posY int, width int, height int, color Color) {
	ccolor := *color.cptr()
	C.DrawRectangleLines(C.int(posX), C.int(posY), C.int(width), C.int(height), ccolor)
}

//DrawRectangleLinesEx : Draw rectangle outline with extended parameters
func DrawRectangleLinesEx(rec Rectangle, lineThick int, color Color) {
	ccolor := *color.cptr()
	crec := *rec.cptr()
	C.DrawRectangleLinesEx(crec, C.int(lineThick), ccolor)
}

//DrawRectangleRounded : Draw rectangle with rounded edges
func DrawRectangleRounded(rec Rectangle, roundness float32, segments int, color Color) {
	ccolor := *color.cptr()
	crec := *rec.cptr()
	C.DrawRectangleRounded(crec, C.float(roundness), C.int(segments), ccolor)
}

//DrawRectangleRoundedLines : Draw rectangle with rounded edges outline
func DrawRectangleRoundedLines(rec Rectangle, roundness float32, segments int, lineThick int, color Color) {
	ccolor := *color.cptr()
	crec := *rec.cptr()
	C.DrawRectangleRoundedLines(crec, C.float(roundness), C.int(segments), C.int(lineThick), ccolor)
}

//DrawTriangle : Draw a color-filled triangle (vertex in counter-clockwise order!)
func DrawTriangle(v1 Vector2, v2 Vector2, v3 Vector2, color Color) {
	ccolor := *color.cptr()
	cv3 := *v3.cptr()
	cv2 := *v2.cptr()
	cv1 := *v1.cptr()
	C.DrawTriangle(cv1, cv2, cv3, ccolor)
}

//DrawTriangleLines : Draw triangle outline (vertex in counter-clockwise order!)
func DrawTriangleLines(v1 Vector2, v2 Vector2, v3 Vector2, color Color) {
	ccolor := *color.cptr()
	cv3 := *v3.cptr()
	cv2 := *v2.cptr()
	cv1 := *v1.cptr()
	C.DrawTriangleLines(cv1, cv2, cv3, ccolor)
}

//DrawTriangleFan : Draw a triangle fan defined by points (first vertex is the center)
func DrawTriangleFan(points Vector2, numPoints int, color Color) Vector2 {
	ccolor := *color.cptr()
	cpoints := points.cptr()
	C.DrawTriangleFan(&cpoints, C.int(numPoints), ccolor)
	return newVector2FromPointer(unsafe.Pointer(&cpoints))
}

//DrawTriangleStrip : Draw a triangle strip defined by points
func DrawTriangleStrip(points Vector2, pointsCount int, color Color) Vector2 {
	ccolor := *color.cptr()
	cpoints := points.cptr()
	C.DrawTriangleStrip(&cpoints, C.int(pointsCount), ccolor)
	return newVector2FromPointer(unsafe.Pointer(&cpoints))
}

//DrawPoly : Draw a regular polygon (Vector version)
func DrawPoly(center Vector2, sides int, radius float32, rotation float32, color Color) {
	ccolor := *color.cptr()
	ccenter := *center.cptr()
	C.DrawPoly(ccenter, C.int(sides), C.float(radius), C.float(rotation), ccolor)
}

//SetShapesTexture : Define default texture used to draw shapes
func SetShapesTexture(texture Texture2D, source Rectangle) {
	csource := *source.cptr()
	ctexture := *texture.cptr()
	C.SetShapesTexture(ctexture, csource)
}

//CheckCollisionRecs : Check collision between two rectangles
func CheckCollisionRecs(rec1 Rectangle, rec2 Rectangle) bool {
	crec2 := *rec2.cptr()
	crec1 := *rec1.cptr()
	res := C.CheckCollisionRecs(crec1, crec2)
	return bool(res)
}

//CheckCollisionCircles : Check collision between two circles
func CheckCollisionCircles(center1 Vector2, radius1 float32, center2 Vector2, radius2 float32) bool {
	ccenter2 := *center2.cptr()
	ccenter1 := *center1.cptr()
	res := C.CheckCollisionCircles(ccenter1, C.float(radius1), ccenter2, C.float(radius2))
	return bool(res)
}

//CheckCollisionCircleRec : Check collision between circle and rectangle
func CheckCollisionCircleRec(center Vector2, radius float32, rec Rectangle) bool {
	crec := *rec.cptr()
	ccenter := *center.cptr()
	res := C.CheckCollisionCircleRec(ccenter, C.float(radius), crec)
	return bool(res)
}

//GetCollisionRec : Get collision rectangle for two rectangles collision
func GetCollisionRec(rec1 Rectangle, rec2 Rectangle) Rectangle {
	crec2 := *rec2.cptr()
	crec1 := *rec1.cptr()
	res := C.GetCollisionRec(crec1, crec2)
	return newRectangleFromPointer(unsafe.Pointer(&res))
}

//CheckCollisionPointRec : Check if point is inside rectangle
func CheckCollisionPointRec(point Vector2, rec Rectangle) bool {
	crec := *rec.cptr()
	cpoint := *point.cptr()
	res := C.CheckCollisionPointRec(cpoint, crec)
	return bool(res)
}

//CheckCollisionPointCircle : Check if point is inside circle
func CheckCollisionPointCircle(point Vector2, center Vector2, radius float32) bool {
	ccenter := *center.cptr()
	cpoint := *point.cptr()
	res := C.CheckCollisionPointCircle(cpoint, ccenter, C.float(radius))
	return bool(res)
}

//CheckCollisionPointTriangle : Check if point is inside a triangle
func CheckCollisionPointTriangle(point Vector2, p1 Vector2, p2 Vector2, p3 Vector2) bool {
	cp3 := *p3.cptr()
	cp2 := *p2.cptr()
	cp1 := *p1.cptr()
	cpoint := *point.cptr()
	res := C.CheckCollisionPointTriangle(cpoint, cp1, cp2, cp3)
	return bool(res)
}
