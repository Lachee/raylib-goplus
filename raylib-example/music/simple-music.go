package main

import (
	"fmt"

	r "github.com/lachee/raylib-goplus/raylib"
)

/*
Music from https://filmmusic.io
"Midnight Tale" by Kevin MacLeod (https://incompetech.com)
License: CC BY (http://creativecommons.org/licenses/by/4.0/)
*/

func main() {

	//Setup Configuration
	r.SetConfigFlags(r.FlagVsyncHint)
	r.InitWindow(800, 400, "Hello world!")
	r.SetTraceLogLevel(r.LogAll)
	r.SetTraceLogCallback(func(logType r.TraceLogType, text string) {
		fmt.Println(logType.ToString(), ": ", text)
	})
	r.TraceLog(r.LogInfo, "Simply, define! ", "Just like", "fmt!")

	r.InitAudioDevice()

	//Load the music
	//Will hard crash if not init audio device.
	music := r.LoadMusicStream("../resources/4710-midnight-tale-by-kevin-macleod.mp3")
	pop := r.LoadSound("pop-A.ogg")

	//TODO: See issue https://github.com/Lachee/raylib-goplus/issues/3
	if !pop.IsValid() {
		r.TraceLog(r.LogError, "Failed to load pop noise.")
	}

	//fmt.Println(music)
	music.PlayStream()

	frame := 0
	for !r.WindowShouldClose() {
		music.UpdateStream()
		frame++

		if r.IsKeyPressed(r.KeySpace) {
			pop.Play()
		}

		duration := float32(music.GetTimePlayed()) / float32(music.GetTimeLength())

		r.GuiProgressBar(r.NewRectangle(100, 100, 200, 25), "Start ", " End", duration, float32(0), float32(1))

		r.BeginDrawing()
		r.ClearBackground(r.RayWhite)
		r.DrawText(fmt.Sprintf("Playing Music %2f%%", float64(duration)), 100, 80, 10, r.Black)
		r.DrawText("Music from https://filmmusic.io", 100, 150, 10, r.Black)
		r.DrawText("\"Midnight Tale\" by Kevin MacLeod (https://incompetech.com)", 100, 160, 10, r.Black)
		r.DrawText("License: CC BY (http://creativecommons.org/licenses/by/4.0/)", 100, 170, 10, r.Black)
		r.DrawFPS(3, 3)
		r.EndDrawing()
	}

	//UNload the Music
	music.Unload()

	//This unload is experimental. In theory, all objects with Unload() will
	// have been generated to register themselves to a global slice. This func
	// will go through the slice, calling Unload() on everything that is
	// still there.
	r.UnloadAll()

	//Close the window.
	r.CloseWindow()
}
