package main

import r "github.com/lachee/raylib-goplus/raylib"

func main() {

	//This is the same as HelloWorld, but using recommended Go paradigms

	//Prepare the window and when we are all done, close it
	r.InitWindow(800, 450, "Raylib Go Plus - Hellow World") //Creates the window
	defer r.CloseWindow()                                   //Closes the window at the end of the application
	defer r.UnloadAll()                                     //Unloads any textures or sounds we may have missed. Experimental.

	//Set FPS
	r.SetTargetFPS(60)

	//Draw every frame
	for !r.WindowShouldClose() {
		drawFrame()
	}
}

//drawFrame renders a single frame
func drawFrame() {
	r.BeginDrawing()     //Begin drawing the frame
	defer r.EndDrawing() //Defer ending the drawing

	//Clear the background
	r.ClearBackground(r.RayWhite)

	//Draw Hello. DrawText(message string, x int, y int, size int, color r.Color)
	r.DrawText("Hello World!", 10, 10, 25, r.GopherBlue)
}
