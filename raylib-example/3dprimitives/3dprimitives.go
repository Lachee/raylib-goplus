package main

import r "github.com/lachee/raylib-goplus/raylib"
import "math"

var orbitSpeed = float64(r.Deg2Rad)
var zoomSpeed = 1.0
var camera r.Camera

func main() {
	screenWidth := 800
	screenHeight := 450

	r.InitWindow(screenWidth, screenHeight, "raylib [models] example - geometric shapes")

	camera = r.NewCamera(r.NewVector3(0.0, 10.0, 10.0), r.NewVector3(0.0, 0.0, 0.0), r.NewVector3(0.0, 1.0, 0.0), 45, r.CameraTypePerspective)
	r.SetTargetFPS(60)

	selfOrbit := true
	orbit := 0.0
	distance := 10.0

	for !r.WindowShouldClose() {

		//Modify the orbits
		odif, ddif := GetOrbitInput()
		orbit += odif
		distance += ddif

		if selfOrbit {
			orbit += orbitSpeed
		}

		//Loop the orbit
		if orbit > 2*math.Pi {
			orbit = -2 * math.Pi
		} else if orbit < -2*math.Pi {
			orbit = 2 * math.Pi
		}

		//Apply the orbit modification
		camera.Position.X = float32(math.Sin(orbit) * distance)
		camera.Position.Z = float32(math.Cos(orbit) * distance)

		r.BeginDrawing()

		//Draw the static shapes
		drawShapes()

		//Draw a UI to update the orbits
		r.GuiLabel(r.NewRectangle(10, 25, 100, 20), "distance")
		distance = float64(r.GuiSlider(r.NewRectangle(100, 25, 150, 20), "0%", "100%", float32(distance), 2, 100))

		r.GuiLabel(r.NewRectangle(10, 50, 100, 20), "orbit")
		orbit = float64(r.GuiSlider(r.NewRectangle(100, 50, 150, 20), "-360", "360", float32(orbit)*r.Rad2Deg, -360, 360) * r.Deg2Rad)

		r.GuiLabel(r.NewRectangle(10, 75, 100, 20), "speed")
		orbitSpeed = float64(r.GuiSlider(r.NewRectangle(100, 75, 150, 20), "-100%", "100%", float32(orbitSpeed)*r.Rad2Deg, -10, 10) * r.Deg2Rad)

		//Enable / Disable the auto fly mode
		selfOrbit = r.GuiCheckBox(r.NewRectangle(100, 100, 20, 20), "Auto Orbit", selfOrbit)

		r.EndDrawing()
	}

	r.CloseWindow()
}

func drawShapes() {

	//Shape Source: https://github.com/gen2brain/raylib-go/blob/master/examples/models/geometric_shapes/main.go
	r.ClearBackground(r.RayWhite)

	r.BeginMode3D(camera)

	r.DrawCube(r.NewVector3(-4.0, 0.0, 2.0), 2.0, 5.0, 2.0, r.Red)
	r.DrawCubeWires(r.NewVector3(-4.0, 0.0, 2.0), 2.0, 5.0, 2.0, r.Gold)
	r.DrawCubeWires(r.NewVector3(-4.0, 0.0, -2.0), 3.0, 6.0, 2.0, r.Maroon)

	r.DrawSphere(r.NewVector3(-1.0, 0.0, -2.0), 1.0, r.Green)
	r.DrawSphereWires(r.NewVector3(1.0, 0.0, 2.0), 2.0, 16, 16, r.Lime)

	r.DrawCylinder(r.NewVector3(4.0, 0.0, -2.0), 1.0, 2.0, 3.0, 4, r.GopherBlue)
	r.DrawCylinderWires(r.NewVector3(4.0, 0.0, -2.0), 1.0, 2.0, 3.0, 4, r.DarkBlue)
	r.DrawCylinderWires(r.NewVector3(4.5, -1.0, 2.0), 1.0, 1.0, 2.0, 6, r.Brown)

	r.DrawCylinder(r.NewVector3(1.0, 0.0, -4.0), 0.0, 1.5, 3.0, 8, r.Gold)
	r.DrawCylinderWires(r.NewVector3(1.0, 0.0, -4.0), 0.0, 1.5, 3.0, 8, r.Pink)

	r.DrawGrid(10, 1.0) // Draw a grid

	r.EndMode3D()

	r.DrawFPS(10, 10)

}

func GetOrbitInput() (float64, float64) {
	orbit := 0.0
	dist := 0.0

	if r.IsKeyDown(r.KeyA) {
		orbit -= orbitSpeed
	}

	if r.IsKeyDown(r.KeyD) {
		orbit += orbitSpeed
	}

	if r.IsKeyDown(r.KeyW) {
		dist -= zoomSpeed
	}

	if r.IsKeyDown(r.KeyS) {
		dist += zoomSpeed
	}

	return orbit, dist
}
