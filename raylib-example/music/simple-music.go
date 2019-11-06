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
	music := r.LoadMusicStream("4710-midnight-tale-by-kevin-macleod.mp3")
	//fmt.Println(music)
	music.PlayStream()

	frame := 0
	for !r.WindowShouldClose() {
		music.UpdateStream()
		frame++

		duration := float32(music.GetTimePlayed()) / float32(music.GetTimeLength())
		//duration := float32(frame%255) / 255

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

	r.CloseWindow()
}
