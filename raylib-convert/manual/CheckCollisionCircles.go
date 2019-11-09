//CheckCollisionCircles : Check collision between two circles
func CheckCollisionCircles(center1 Vector2, radius1 float32, center2 Vector2, radius2 float32) bool {
	distance := center1.Distance(center2)
	return distance <= radius1 + radius2
}
