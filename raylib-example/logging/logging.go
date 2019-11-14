package main

import (
	"fmt"

	r "github.com/lachee/raylib-goplus/raylib"
)

func main() {

	/*
		Set up the target log level. In this case we want ALL, but there is numerous to choose.
		All Options:
			LogAll TraceLogType = iota
			LogTrace
			LogDebug
			LogInfo
			LogWarning
			LogError
			LogFatal
			LogNone
	*/
	r.SetTraceLogLevel(r.LogAll)

	//We don't want our logs to throw a panic
	r.SetTraceLogExit(r.LogNone)

	//Set up a custom callback to the log function.
	//Raylib will automatically filter our logs for us and prepend them appropriately.
	//
	//NOTE:
	//	It is important that we set this callback up BEFORE the init,
	//	otherwise we will miss the init logs.
	r.SetTraceLogCallback(func(logType r.TraceLogType, text string) {

		//For simplicity, we will just log it using fmt. But you can easily add a file
		// logger or some other tracking.
		//Note that r.TraceLogType has a function ToString that will convert its number
		// but in this case, we are using ToUniformedString that converts it to a 5 letter code.
		fmt.Println("[" + logType.ToUniformedString() + "] " + text)
	})

	//Prepare the window. This should be done after the SetTraceLogCallback so we dont
	// miss anything.
	r.InitWindow(800, 450, "Raylib Go Plus - Hellow World")
	r.SetTargetFPS(60)

	//Draw every frame
	for !r.WindowShouldClose() {

		//Begin the drawing
		r.BeginDrawing()
		r.ClearBackground(r.RayWhite)

		//Draw a new GUI button that we can press to show the logs
		if r.GuiButton(r.NewRectangle(10, 10, 200, 25), "I make a log!") {

			//These logs work just like fmt.Println, except they have an additional type.
			// Note that since we are running a custom callback, it won't actually go through raylib.
			//   if we didn't have a custom callback, then it will call the C.TraceLog function.
			r.TraceLog(r.LogTrace, "This is a trace!")
			r.TraceLog(r.LogDebug, "This is a debug!")
			r.TraceLog(r.LogInfo, "This is a info!")
			r.TraceLog(r.LogWarning, "This is a warning!")
			r.TraceLog(r.LogError, "This is a error!")
			r.TraceLog(r.LogFatal, "This is a fatal!")
		}

		//End the drawing
		r.EndDrawing()
	}

	//Close the window
	r.CloseWindow()
}
