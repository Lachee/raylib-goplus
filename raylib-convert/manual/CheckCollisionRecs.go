//CheckCollisionRecs : Check collision between two rectangles
// Alias of rec1.Overlaps(rect2) instead.
func CheckCollisionRecs(r Rectangle, rect Rectangle) bool {
	return (r.X < (rect.X+rect.Width) && (r.X+r.Width) > rect.X) && (r.Y < (rect.Y+rect.Height) && (r.Y+r.Height) > rect.Y)
}

//Overlaps checks if a rectangle overlaps another.
func (r Rectangle) Overlaps(rect Rectangle) bool {
	return CheckCollisionRecs(r, rect)
}
