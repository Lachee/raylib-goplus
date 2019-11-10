package main

import r "github.com/lachee/raylib-goplus/raylib"

func main() {
	screenWidth := 800
	screenHeight := 450

	r.InitWindow(screenWidth, screenHeight, "Raylib Go Plus - SrcRec DistRec")
	r.SetTargetFPS(60)

	texture := r.LoadTexture("../resources/gopher_sheet.png")
	frameCount := 24
	frameWidth := float32(texture.Width / int32(frameCount))
	frameHeight := float32(texture.Height)

	sourceRec := r.NewRectangle(0, 0, frameWidth, frameHeight)
	destRec := r.NewRectangle(float32(screenWidth)/2, float32(screenHeight)/2, frameWidth*2, frameHeight*2)
	origin := r.NewVector2(frameWidth, frameHeight)

	frame := 0
	for !r.WindowShouldClose() {
		frame++

		if frame%5 == 0 {
			sourceRec.X += frameWidth
		}

		r.BeginDrawing()
		r.ClearBackground(r.RayWhite)

		r.DrawTexturePro(*texture, sourceRec, destRec, origin, float32(frame), r.White)
		r.DrawLine(int(destRec.X), 0, int(destRec.X), screenHeight, r.Gray)
		r.DrawLine(0, int(destRec.Y), screenWidth, int(destRec.Y), r.Gray)
		r.DrawText("Party Gopher by Egon Elbre", screenWidth-200, screenHeight-20, 10, r.Gray)
		r.EndDrawing()
	}

	texture.Unload()
	r.UnloadAll()
	r.CloseWindow()
}
