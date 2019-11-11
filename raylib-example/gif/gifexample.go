package main

import (
	"fmt"

	r "github.com/lachee/raylib-goplus/raylib"
	"github.com/lachee/raylib-goplus/raylib-gif"
)

func main() {
	screenWidth := 800
	screenHeight := 450

	defer r.UnloadAll()

	r.SetTraceLogLevel(r.LogAll)
	r.SetTraceLogCallback(func(logType r.TraceLogType, text string) {
		fmt.Println(logType.ToString(), ": ", text)
	})

	//http://www.theimage.com/animation/pages/disposal2.html
	r.InitWindow(screenWidth, screenHeight, "Raylib Go Plus - GIF Example")
	r.SetTargetFPS(60)
	frame := 0

	var texture *rgif.GifImage

	for !r.WindowShouldClose() {
		frame++

		//======= This is the input part
		if r.IsKeyPressed(r.KeyA) || texture == nil {
			if texture != nil {
				texture.Unload()
			}
			texture, _ = rgif.LoadGifFromFile("../resources/gopher.gif")
		}
		if r.IsKeyPressed(r.KeyS) {
			if texture != nil {
				texture.Unload()
			}
			texture, _ = rgif.LoadGifFromFile("../resources/testgif_04.gif")
		}
		if r.IsKeyPressed(r.KeyD) {
			if texture != nil {
				texture.Unload()
			}
			texture, _ = rgif.LoadGifFromFile("../resources/testgif_05.gif")
		}
		if r.IsKeyPressed(r.KeyF) {
			if texture != nil {
				texture.Unload()
			}
			texture, _ = rgif.LoadGifFromFile("../resources/laugh.gif")
		}
		if r.IsKeyPressed(r.KeyG) {
			if texture != nil {
				texture.Unload()
			}
			texture, _ = rgif.LoadGifFromFile("../resources/important.gif")
		}

		r.BeginDrawing()
		r.ClearBackground(r.RayWhite)

		//pos := r.NewVector2(10, float32(screenHeight-50))
		//scale := float32(0.25)

		if texture != nil {
			//Step the animation and then draw it
			//===== THIS IS THE IMPORTANT PART =======
			texture.Step(r.GetFrameTime())
			rgif.DrawGif(texture, 100, 100, r.White) //This draws it normally, and a nice flat gif

			//===== This is the debugging part =====
			//Draw Debug about the current frame
			r.DrawRectangleLinesEx(r.NewRectangle(100, 100, float32(texture.Width), float32(texture.Height)), 1, r.LightGray)
			r.DrawText(fmt.Sprintf("#%d", texture.CurrentFrame()), 100, 75, 20, r.LightGray)
			r.DrawText(fmt.Sprintf(".0%ds", texture.CurrentTiming()), 100+texture.Width-50, 75, 20, r.LightGray)

		}

		r.DrawText("Party Gopher by Egon Elbre", screenWidth-200, screenHeight-20, 10, r.Gray)
		r.DrawText("Press A, S, F, or G to load different GIFs", 10, 10, 20, r.Gray)
		r.DrawFPS(screenWidth-30, 10)
		r.EndDrawing()
	}

	r.CloseWindow()
}
