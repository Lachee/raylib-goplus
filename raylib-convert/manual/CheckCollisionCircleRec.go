//CheckCollisionCircleRec : Check collision between circle and rectangle
func CheckCollisionCircleRec(center Vector2, radius float32, rec Rectangle) bool {
	recCenter := rec.Center()
	dx := float32(math.Abs(float64(center.X - recCenter.X)))
	dy := float32(math.Abs(float64(center.Y - recCenter.Y)))
	
	if dx > (rec.Width/2 + radius) || dy > (rec.Height/2+radius) {
		return false
	}
	
	if dx <= rec.Width/2 || dy <= rec.Height/2 {
		return true
	}
	
	ddx := dx - rec.Width/2
	ddy := dy - rec.Height/2
	cornerDistanceSq := ddx*ddx + ddy*ddy
	return cornerDistanceSq <= radius * radius
}
