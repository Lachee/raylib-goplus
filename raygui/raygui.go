package main

import (
	rl "github.com/lachee/raylib-goplus/raylib"
)

/*
#define RAYGUI_IMPLEMENTATION
#cgo CFLAGS: -I../raylib -DPLATFORM_DESKTOP
#include "../raylib/raylib.h"

#include "raygui.h"
#include <stdlib.h>
*/
import "C"

func main() {
	color := rl.Color{R: 255, G: 0, B: 255, A: 255}
	color2 := rl.Color{R: 255, G: 255, B: 255, A: 255}
	rl.InitWindow(800, 400, "Hello World")

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(color2)
		rl.DrawText("Drawing A Label!", 10, 10, 20, color)
		rl.EndDrawing()
	}

	rl.CloseWindow()
}
