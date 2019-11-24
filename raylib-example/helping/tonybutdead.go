package main

import (
	"github.com/lachee/raylib-goplus/raylib"
)

func main() {
	raylib.InitWindow(800, 600, "Hello World!")

	var camera raylib.Camera
	camera.Position = raylib.Vector3{X: 3, Y: 3, Z: 3}
	camera.Target = raylib.Vector3{X: 0, Y: 1.5, Z: 0}
	camera.Up = raylib.Vector3{X: 0, Y: 1, Z: 0}
	camera.FOVY = 45
	camera.Type = raylib.CameraTypePerspective

	raylib.SetTargetFPS(60)

	for !raylib.WindowShouldClose() {
		raylib.BeginDrawing()
		raylib.ClearBackground(raylib.White)

		raylib.BeginMode3D(camera)

		raylib.DrawGrid(10, 1)

		raylib.EndMode3D()

		raylib.EndDrawing()
	}

	raylib.CloseWindow()
}
