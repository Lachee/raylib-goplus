//CheckCollisionPointRec : Check if point is inside rectangle
func CheckCollisionPointRec(point Vector2, r Rectangle) bool {
	return point.X >= r.X && point.X <= (r.X+r.Width) && point.Y >= r.Y && point.Y <= (r.Y+r.Height)
}

//Contains checks if the rectangle contains a point
func (r Rectangle) Contains(point Vector2) bool {
	return CheckCollisionPointRec(point, r)
}