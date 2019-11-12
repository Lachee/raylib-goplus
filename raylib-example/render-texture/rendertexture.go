package main

import (
	"fmt"

	r "github.com/lachee/raylib-goplus/raylib"
)

func main() {
	screenWidth := 800
	screenHeight := 450

	r.SetConfigFlags(r.FlagMsaa4xHint) // Enable Multi Sampling Anti Aliasing 4x (if available)
	r.SetTraceLogLevel(r.LogAll)
	r.SetTraceLogCallback(func(logType r.TraceLogType, text string) {
		fmt.Println(logType.ToString(), ": ", text)
	})

	r.InitWindow(screenWidth, screenHeight, "raylib [shaders] example - custom uniform variable")
	defer r.UnloadAll()

	camera := r.Camera{}
	camera.Position = r.NewVector3(3.0, 3.0, 3.0)
	camera.Target = r.NewVector3(0.0, 1.5, 0.0)
	camera.Up = r.NewVector3(0.0, 1.0, 0.0)
	camera.FOVY = 45.0

	dwarf := r.LoadModel("../resources/dwarf.obj")             // Load OBJ model
	texture := r.LoadTexture("../resources/dwarf_diffuse.png") // Load model texture

	//dwarf.Materials.Maps[r.MapDiffuse].Texture = texture // Set dwarf model diffuse texture
	mat := dwarf.Materials[0]
	mat.SetTexture(r.MapAlbedo, texture)

	position := r.NewVector3(0.0, 0.0, 0.0) // Set model position

	shader := r.LoadShader("", "../resources/glsl330/swirl.fs") // Load postpro shader

	// Get variable (uniform) location on the shader to connect with the program
	// NOTE: If uniform variable could not be found in the shader, function returns -1
	//swirlCenterLoc := r.GetShaderLocation(shader, "center")
	swirlCenterLoc := shader.GetLocation("center")
	swirlCenter := r.NewVector2(float32(screenWidth)/2, float32(screenHeight)/2)

	// Create a RenderTexture2D to be used for render to texture
	target := r.LoadRenderTexture(screenWidth-10, screenHeight-10)

	// Setup orbital camera
	//r.SetCameraMode(camera, r.CameraOrbital) // Set an orbital camera mode
	camera.SetMode(r.CameraOrbital)

	r.SetTargetFPS(60)

	for !r.WindowShouldClose() {
		// Update
		//----------------------------------------------------------------------------------

		camera.Update()

		//Begin drawing like normal
		r.BeginDrawing()
		r.ClearBackground(r.RayWhite)

		//Start a texture drawing. Anything rendered will now be on the target
		r.BeginTextureMode(target)                          // Enable drawing to texture
		r.ClearBackground(r.RayWhite)                       // Clear texture background. 3D wont work if you forget this
		r.DrawText("Drawing Before 3D", 200, 10, 30, r.Red) //Draw some text to show the 3D models draw order

		//Start the 3D camera
		r.BeginMode3D(camera)
		r.DrawModel(*dwarf, position, 2.0, r.White) // Draw 3d model with texture
		r.DrawGrid(10, 1.0)                         // Draw a grid
		r.EndMode3D()

		//Stop the 3D camera, and draw text again above.
		r.DrawText("Drawing After 3D", 200, screenHeight-40, 30, r.Red)
		r.EndTextureMode() // End drawing to texture (now we have a texture available for next passes)

		//Button press to export the generated image.
		if r.IsKeyReleased(r.KeyS) {
			td := r.GetTextureData(target.Texture)
			td.Export("C:/myimage.png")
			td.Unload()
		}

		//Update the swirl position, converting Screen Space to Shader Space (by flipping the Y)
		swirlCenter = r.GetMousePosition()
		swirlCenter.Y = float32(screenHeight) - swirlCenter.Y

		//Set teh shader `center` value (its a Vector2).
		//  Note that  because we are setting 1 value, we use SetValueFloat32,
		//  and NOT the SetValueFloat32V.
		shader.SetValueFloat32(swirlCenterLoc, swirlCenter.Decompose(), r.UniformVec2)
		r.BeginShaderMode(*shader)

		// NOTE: Render texture must be y-flipped due to default OpenGL coordinates (left-bottom)
		r.DrawTextureRec(target.Texture, r.NewRectangle(0, 0, float32(target.Texture.Width), float32(-target.Texture.Height)), r.NewVector2(0, 0), r.White)

		r.EndShaderMode()

		r.DrawText("(c) Dwarf 3D model by David Moreno", screenWidth-200, screenHeight-20, 10, r.Gray)
		r.DrawFPS(10, 10)

		r.EndDrawing()
	}

	r.CloseWindow()
}
