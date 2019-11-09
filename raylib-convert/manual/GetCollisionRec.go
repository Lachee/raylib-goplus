//GetOverlapRec : Get collision rectangle for two rectangles collision
// Alias of GetCollisionRec
func (rec1 Rectangle) GetOverlapRec(rec2 Rectangle) Rectangle {
	return GetCollisionRec(rec1, rec2)
}

//GetCollisionRec : Get collision rectangle for two rectangles collision
func GetCollisionRec(rec1 Rectangle, rec2 Rectangle) Rectangle {
	retRec := Rectangle{X:0, Y:0, Width:0, Height:0}
	
	if CheckCollisionRecs(rec1, rec2) {
		dxx := float32(math.Abs(float64(rec1.X - rec2.X)))
		dyy := float32(math.Abs(float64(rec1.Y - rec2.Y)))
		
		if rec1.X <= rec2.X {
			if rec1.Y <= rec2.Y {
				retRec.X = rec2.X;
				retRec.Y = rec2.Y;
				retRec.Width = rec1.Width - dxx;
				retRec.Height = rec1.Height - dyy;
			} else {
				retRec.X = rec2.X;
				retRec.Y = rec1.Y;
				retRec.Width = rec1.Width - dxx;
				retRec.Height = rec2.Height - dyy;
			}
		} else {
			if rec1.Y <= rec2.Y {
				retRec.X = rec1.X;
				retRec.Y = rec2.Y;
				retRec.Width = rec2.Width - dxx;
				retRec.Height = rec1.Height - dyy;
			} else {
				retRec.X = rec1.X;
				retRec.Y = rec1.Y;
				retRec.Width = rec2.Width - dxx;
				retRec.Height = rec2.Height - dyy;
			}
		}
	}
	
	if rec1.Width > rec2.Width {
		if retRec.Width >= rec2.Width {
			retRec.Width = rec2.Width
		} 
	} else {
		if retRec.Width >= rec1.Width {
			retRec.Width = rec1.Width
		}
	}
	
	if rec1.Height > rec2.Height {
		if retRec.Height >= rec2.Height {
			retRec.Height = rec2.Height
		} 
	} else {
		if retRec.Height >= rec1.Height {
			retRec.Height = rec1.Height
		}
	}
	
	return retRec
}