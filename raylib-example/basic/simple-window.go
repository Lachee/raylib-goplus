package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"

	r "github.com/lachee/raylib-goplus/raylib"
)

type ball struct {
	x       int
	y       int
	vx      int
	vy      int
	texture *r.Texture2D
}

func (b *ball) move(maxX int, maxY int) {

	if b.x+int((*b.texture).Width) >= maxX || b.x <= 0 {
		b.vx *= -1
	}
	if b.y+int((*b.texture).Height) >= maxY || b.y <= 0 {
		b.vy *= -1
	}

	b.x += b.vx
	b.y += b.vy
}

const (
	maxSamples          = 22050
	maxSamplesPerUpdate = 4096
)

func main() {

	r.SetConfigFlags(r.FlagVsyncHint)
	r.InitWindow(800, 400, "Hello world!")

	r.SetTraceLogCallback(func(logType r.TraceLogType, text string) {
		fmt.Println(logType.ToString(), ": ", text)
	})

	var frame int
	var offset int
	var completion float32

	windowVisible := true

	stream := r.InitAudioStream(22050, 32, 1)

	//// Fill audio stream with some samples (sine wave)
	data := make([]float32, maxSamples)

	for i := 0; i < maxSamples; i++ {
		data[i] = float32(math.Sin(float64(((2 * r.PI * float32(i)) / 2) * r.Deg2Rad)))
	}

	rgp := r.LoadTexture("../../logo/raylib_goplus_48x48.png")
	rlb := r.LoadTexture("../../logo/raylib_32x32.png")
	balls := [...]*ball{
		&ball{x: 10, y: 10, vx: -rand.Intn(10), vy: rand.Intn(5), texture: rgp},
		&ball{x: 10, y: 10, vx: -rand.Intn(10), vy: rand.Intn(5), texture: rgp},
		&ball{x: 10, y: 10, vx: -rand.Intn(10), vy: rand.Intn(5), texture: rgp},
		&ball{x: 10, y: 10, vx: -rand.Intn(10), vy: rand.Intn(5), texture: rgp},

		&ball{x: 700, y: 300, vx: -rand.Intn(10), vy: rand.Intn(5), texture: rlb},
		&ball{x: 700, y: 300, vx: -rand.Intn(10), vy: rand.Intn(5), texture: rlb},
		&ball{x: 700, y: 300, vx: -rand.Intn(10), vy: rand.Intn(5), texture: rlb},
		&ball{x: 700, y: 300, vx: -rand.Intn(10), vy: rand.Intn(5), texture: rlb},
		&ball{x: 700, y: 300, vx: -rand.Intn(10), vy: rand.Intn(5), texture: rlb},
		&ball{x: 700, y: 300, vx: -rand.Intn(10), vy: rand.Intn(5), texture: rlb},
		&ball{x: 700, y: 300, vx: -rand.Intn(10), vy: rand.Intn(5), texture: rlb},
		&ball{x: 700, y: 300, vx: -rand.Intn(10), vy: rand.Intn(5), texture: rlb},
		&ball{x: 700, y: 300, vx: -rand.Intn(10), vy: rand.Intn(5), texture: rlb},
		&ball{x: 700, y: 300, vx: -rand.Intn(10), vy: rand.Intn(5), texture: rlb},
	}

	//r.GuiLoadStyle("cyber/cyber.rgs")

	for !r.WindowShouldClose() {

		if r.IsFileDropped() {
			files := r.GetDroppedFiles()
			fmt.Println("Dropped Files: ", files)
			r.ClearDroppedFiles()
		}

		frame++
		offset = (offset + 5) % 400
		completion = 360 * (float32(offset) / 400.0)

		hsv := r.NewColorFromHSV(r.NewVector3(completion, 1, 1))
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

		if r.GuiImageButton(r.NewRectangle(400, 30, 48, 48), "Raylib Logo", *rlb) {
			r.OpenURL("https://www.raylib.com/")
		}

		for _, b := range balls {
			b.move(800, 400)
			r.DrawTexture(*b.texture, b.x, b.y, r.White)
		}

		if r.IsGamepadAvailable(r.GamepadPlayer1) {
			r.DrawText("Gamepad Connected", 10, 50, 15, r.GopherBlue)
		}

		if r.IsKeyReleased(r.KeyF1) {
			r.UnloadAll()
		}

		r.DrawFPS(3, 3)
		r.EndDrawing()
	}

	r.UnloadAll()
	r.CloseWindow()
}
