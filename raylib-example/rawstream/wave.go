package main

import (
	"math"

	r "github.com/lachee/raylib-goplus/raylib"
)

const (
	MaxSamples          = 22050
	MaxSamplesPerUpdate = 4096
)

func main() {
	screenWidth := 800
	screenHeight := 450

	r.InitWindow(screenWidth, screenHeight, "Raylib Go Plus - Raw Stream")
	r.InitAudioDevice()
	defer r.CloseAudioDevice()
	defer r.CloseWindow()
	defer r.UnloadAll()

	stream := r.InitAudioStream(22050, 32, 1)
	defer stream.Unload()

	frequency := float32(1)
	oldFrequency := float32(0)

	samplesLeft := MaxSamples
	data := [MaxSamples]float32{}

	stream.Play()

	r.SetTargetFPS(30)
	for !r.WindowShouldClose() {

		if frequency != oldFrequency {
			for i := 0; i < MaxSamples; i++ {
				data[i] = float32(math.Sin(float64(frequency * ((2 * r.PI * float32(i)) / 2) * r.Deg2Rad)))
			}
			oldFrequency = frequency
		}

		if r.IsAudioStreamProcessed(stream) {

			samples := samplesLeft
			if samplesLeft >= MaxSamplesPerUpdate {
				samples = MaxSamplesPerUpdate
			}

			//r.UpdateAudioStream(stream, data, samplesCount)
			stream.Update(data[MaxSamples-samplesLeft:], samples)

			samplesLeft -= samples
			if samplesLeft <= 0 {
				samplesLeft = MaxSamples
			}

		}

		r.BeginDrawing()

		r.ClearBackground(r.RayWhite)
		r.DrawText("Playing a Sine Wave", 240, 140, 20, r.LightGray)
		frequency = r.GuiSlider(r.NewRectangle(150, 25, 200, 25), "0", "100", frequency, 1, 1.1)

		// NOTE: Draw a part of the sine wave (only screen width)
		for i := 0; i < MaxSamples; i++ {
			position := r.NewVector2(float32(i), 250+50*data[i])
			r.DrawPixelV(position, r.GopherBlue)
		}

		r.EndDrawing()

	}
}
