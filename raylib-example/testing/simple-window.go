package main

import (
	"fmt"

	r "github.com/lachee/raylib-goplus/raylib"
)

func main() {

	//Prepare the window
	r.InitWindow(800, 450, "Raylib Go Plus - Hellow World")
	r.SetTargetFPS(60)

	img := r.LoadImage("../resources/raylib_goplus_256x256.png")
	tex := r.LoadTextureFromImage(img)

	//Draw every frame
	for !r.WindowShouldClose() {

		//Begin the drawing
		r.BeginDrawing()
		r.ClearBackground(r.RayWhite)

		//Draw Hello. DrawText(message string, x int, y int, size int, color r.Color)
		r.DrawText("Hello World!", 10, 10, 25, r.GopherBlue)
		r.DrawTexture(tex, 10, 10, r.White)

		colors := img.GetPixels()
		fmt.Println(colors)
		//End the drawing
		r.EndDrawing()
	}

	//Close the window
	r.CloseWindow()
}
