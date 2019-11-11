package main

import (
	r "github.com/lachee/raylib-goplus/raylib"
	gif "github.com/lachee/raylib-goplus/raylib-gif"
)

func main() {
	screenWidth := 800
	screenHeight := 450

	r.InitWindow(screenWidth, screenHeight, "Raylib Go Plus - GIF Example")
	r.SetTargetFPS(60)

	var texture *gif.GifImage
	frame := 0
	gifFrame := 0

	for !r.WindowShouldClose() {
		frame++

		if frame%5 == 0 {
			gifFrame++
		}

		if r.IsKeyPressed(r.KeyA) || texture == nil {
			if texture != nil {
				texture.Unload()
			}
			texture, _ = gif.LoadGifFromFile("../resources/gopher.gif")
		}
		if r.IsKeyPressed(r.KeyS) {
			if texture != nil {
				texture.Unload()
			}
			texture, _ = gif.LoadGifFromFile("../resources/laugh.gif")
		}
		if r.IsKeyPressed(r.KeyD) {
			if texture != nil {
				texture.Unload()
			}
			texture, _ = gif.LoadGifFromFile("../resources/important.gif")
		}

		//tint := r.ColorFromHSV(r.NewVector3(float32(frame%360), 1, 1))

		r.BeginDrawing()
		r.ClearBackground(r.RayWhite)

		if texture != nil {
			texture.Step(r.GetFrameTime())
			gif.DrawGif(texture, 100, 100, r.White)
		}

		r.DrawText("Party Gopher by Egon Elbre", screenWidth-200, screenHeight-20, 10, r.Gray)
		r.DrawText("Press A, S, or D to load different GIFs", 10, screenHeight-20, 20, r.Gray)
		r.DrawFPS(10, 10)
		r.EndDrawing()
	}

	texture.Unload()
	r.UnloadAll()
	r.CloseWindow()
}
