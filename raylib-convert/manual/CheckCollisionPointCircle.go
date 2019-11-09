//CheckCollisionPointCircle : Check if point is inside circle
func CheckCollisionPointCircle(point Vector2, center Vector2, radius float32) bool {
	return CheckCollisionCircles(point, 0, center, radius)
}