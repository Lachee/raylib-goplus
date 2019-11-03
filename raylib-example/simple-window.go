package main

import (
	"strconv"

	r "github.com/lachee/raylib-goplus/raylib"
)

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

	//color := r.Color{R: 255, G: 0, B: 255, A: 255}
	//color2 := r.Color{R: 255, G: 255, B: 255, A: 255}

	var frame int
	var offset int
	var completion float32

	windowVisible := true

	texture := r.LoadTexture("raylib_32x32.png")

	for !r.WindowShouldClose() {
		frame++
		offset = (offset + 5) % 400
		completion = 360 * (float32(offset) / 400.0)

		hsv := r.NewColorHSV(r.NewVector3(completion, 1, 1))
		text := "I LOVE RAINBOW TEXT " + strconv.FormatFloat(float64(completion), 'f', 6, 64)
		r.BeginDrawing()
		r.ClearBackground(r.RayWhite)
		r.DrawText(text, 10, offset, 20, hsv)

		if r.GuiButton(r.Rectangle{X: 10, Y: 30, Width: 200, Height: 20}, "Toggle Window") {
			windowVisible = !windowVisible
		}

		if windowVisible && r.GuiWindowBox(r.Rectangle{X: 10, Y: 50, Width: 200, Height: 200}, "Im a Window!") {
			windowVisible = false
		}

		r.GuiImageButton(r.NewRectangle(400, 30, 32, 32), "Raylib Logo", texture)

		r.DrawFPS(3, 3)
		r.EndDrawing()
	}

	r.CloseWindow()
}
