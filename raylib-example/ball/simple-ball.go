package main

import (
	"fmt"
	"math"
	"math/rand"

	r "github.com/lachee/raylib-goplus/raylib"
)

type ball struct {
	x       int
	y       int
	vx      int
	vy      int
	texture *r.Texture2D
}

func (b *ball) getRect() r.Rectangle {
	return r.Rectangle{X: float32(b.x), Y: float32(b.y), Width: float32(b.texture.Width), Height: float32(b.texture.Height)}
}

func (b *ball) move(maxX int, maxY int, others []*ball) bool {

	collided := false
	if b.x+int((*b.texture).Width) >= maxX || b.x <= 0 {
		b.vx *= -1
		collided = true
	} else if b.y+int((*b.texture).Height) >= maxY || b.y <= 0 {
		b.vy *= -1
		collided = true
	} else {
		for _, b2 := range others {
			if b != b2 {
				if b.getRect().Overlaps(b2.getRect()) {
					collided = true
					if b2.x-b.x > 0 {
						b.vx = -int(math.Abs(float64(b.vx)))
					} else {
						b.vx = int(math.Abs(float64(b.vx)))
					}

					if b2.y-b.y > 0 {
						b.vy = -int(math.Abs(float64(b.vy)))
					} else {
						b.vy = int(math.Abs(float64(b.vy)))
					}
					break
				}
			}
		}
	}

	b.x += b.vx
	b.y += b.vy
	return collided
}

func main() {

	r.SetConfigFlags(r.FlagVsyncHint)
	r.InitWindow(800, 400, "Ball Example")

	r.SetTraceLogCallback(func(logType r.TraceLogType, text string) {
		fmt.Println(logType.ToString(), ": ", text)
	})

	r.InitAudioDevice()

	var bonkX int = 0
	var bonkY int = 0
	var bonkFade float32 = 1

	rgp := r.LoadTexture("../../logo/raylib_goplus_48x48.png")
	rlb := r.LoadTexture("../../logo/raylib_32x32.png")
	balls := []*ball{
		&ball{x: 10, y: 10, vx: rand.Intn(10) + 1, vy: rand.Intn(5) + 10, texture: rgp},
		&ball{x: 700, y: 300, vx: rand.Intn(10) + 1, vy: rand.Intn(5) + 1, texture: rlb},
		&ball{x: 700, y: 300, vx: rand.Intn(10) + 1, vy: rand.Intn(5) + 1, texture: rlb},
		&ball{x: 700, y: 300, vx: rand.Intn(10) + 1, vy: rand.Intn(5) + 1, texture: rlb},
	}

	for !r.WindowShouldClose() {
		r.BeginDrawing()
		r.ClearBackground(r.RayWhite)

		//Iterate over allt he balls, drawing them and detecting collisions
		for _, b := range balls {
			if b.move(800, 400, balls) {

				if bonkFade < 0.1 {
					bonkX = b.x
					bonkY = b.y
					bonkFade = 1
				}

			}
			r.DrawTexture(*b.texture, b.x, b.y, r.White)
		}

		//Draw when we bonk
		r.DrawText("Bonk!", bonkX, bonkY, 20, r.Aqua.Fade(bonkFade))

		//Decrement the bonk alpha
		bonkFade = bonkFade * 0.95

		if r.IsKeyPressed(r.KeyLeft) {
			balls = append(balls, &ball{x: 10, y: 10, vx: -rand.Intn(10) + 1, vy: rand.Intn(5) + 1, texture: rgp})
		}

		if r.IsKeyPressed(r.KeyRight) {
			balls = append(balls, &ball{x: 700, y: 300, vx: -rand.Intn(10) + 1, vy: rand.Intn(5) + 1, texture: rlb})
		}

		r.DrawText("Press <- or -> to make more!", 160, 300, 30, r.GopherBlue)

		if r.IsKeyReleased(r.KeyF1) {
			r.UnloadAll()
		}

		r.DrawFPS(3, 3)
		r.EndDrawing()
	}

	r.UnloadAll()
	r.CloseWindow()
}
