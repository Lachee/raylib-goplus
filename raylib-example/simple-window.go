package main

import r "github.com/lachee/raylib-goplus/raylib"

func main() {
	RaylibTestWorld()
	//r.TestWorld()

	/*
		color := r.Color{R: 255, G: 0, B: 255, A: 255}
		color2 := r.Color{R: 255, G: 255, B: 255, A: 255}
		r.InitWindow(800, 400, "Hello World")

		for !r.WindowShouldClose() {
			r.BeginDrawing()
			r.ClearBackground(color2)
			r.DrawText("Drawing A Label!", 10, 10, 20, color)

			//r.Button(r.Rectangle{X: 30, Y: 10, Width: 200, Height: 25}, "Testing")
			r.EndDrawing()
		}

		r.CloseWindow()
	*/
	//r.HelloWorld()

}

func RaylibTestWorld() {

	r.SetConfigFlags(r.FlagVsyncHint)
	r.InitWindow(800, 400, "Hello world!")

	color := r.Color{R: 255, G: 0, B: 255, A: 255}
	color2 := r.Color{R: 255, G: 255, B: 255, A: 255}
	var frame int = 0

	for !r.WindowShouldClose() {
		frame++

		r.BeginDrawing()
		r.ClearBackground(color2)
		r.DrawText("How are you today? ", 10, int32(frame%200), 20, color)
		r.GuiButton(r.Rectangle{X: 10, Y: 30, Width: 200, Height: 20}, "Im a Button!")
		r.EndDrawing()
	}

	r.CloseWindow()
}
