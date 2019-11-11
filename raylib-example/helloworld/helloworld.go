package main

import r "github.com/lachee/raylib-goplus/raylib"

func main() {

	//Prepare the window
	r.InitWindow(800, 450, "Raylib Go Plus - Hellow World")
	r.SetTargetFPS(60)

	//Draw every frame
	for !r.WindowShouldClose() {

		//Begin the drawing
		r.BeginDrawing()
		r.ClearBackground(r.RayWhite)

		//Draw Hello. DrawText(message string, x int, y int, size int, color r.Color)
		r.DrawText("Hello World!", 10, 10, 25, r.GopherBlue)

		//End the drawing
		r.EndDrawing()
	}

	//Close the window
	r.CloseWindow()
}
