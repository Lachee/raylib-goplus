package raylibgif

import (
	"fmt"
	"image"
	"image/draw"
	"image/gif"
	"os"

	r "github.com/lachee/raylib-goplus/raylib"
)

type GifImage struct {
	TileSheet     *r.Texture2D
	Width         int
	Height        int
	Frames        int
	Timing        []int
	currentFrame  int
	lastFrameTime float32
}

func LoadGifFromFile(fileName string) (*GifImage, error) {

	file, err := os.Open(fileName)
	defer file.Close()
	if err != nil {
		return nil, err
	}

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("Error while decoding: %s", r)
		}
	}()

	gif, err := gif.DecodeAll(file)
	if err != nil {
		return nil, err
	}

	imgWidth, imgHeight := getGifDimensions(gif)
	frames := len(gif.Image)

	overpaintImage := image.NewRGBA(image.Rect(0, 0, imgWidth, imgHeight))
	tilesheet := image.NewRGBA(image.Rect(0, 0, imgWidth*frames, imgHeight))

	draw.Draw(overpaintImage, overpaintImage.Bounds(), gif.Image[0], image.ZP, draw.Src)

	for i, srcImg := range gif.Image {
		draw.Draw(overpaintImage, overpaintImage.Bounds(), srcImg, image.ZP, draw.Over)

		dst := overpaintImage.Bounds().Add(image.Point{imgWidth * i, 0})
		draw.Draw(tilesheet, dst, srcImg, image.ZP, draw.Over)
	}

	rayimg := r.LoadTextureFromGoImage(tilesheet)
	return &GifImage{
		TileSheet: rayimg,
		Width:     imgWidth,
		Height:    imgHeight,
		Frames:    frames,
		Timing:    gif.Delay,
	}, nil
}

func (gif *GifImage) Unload() {
	gif.TileSheet.Unload()
}

//Step performs a time step.
func (gif *GifImage) Step(timeSinceLastStep float32) {

	gif.lastFrameTime += (timeSinceLastStep * 100)
	diff := gif.lastFrameTime - float32(gif.Timing[gif.currentFrame])

	if diff >= 0 {
		gif.NextFrame()
		gif.lastFrameTime = 0
	}
}

//NextFrame increments the frame counter and resets the timing buffer
func (gif *GifImage) NextFrame() {
	gif.lastFrameTime -= float32(gif.Timing[gif.currentFrame])
	gif.currentFrame = (gif.currentFrame + 1) % gif.Frames
	if gif.lastFrameTime < 0 {
		gif.lastFrameTime = 0
	}
}

//Reset clears the last frame time and resets the current frame to zero
func (gif *GifImage) Reset() {
	gif.currentFrame = 0
	gif.lastFrameTime = 0
}

//CurrentFrame returns the current frame index
func (gif *GifImage) CurrentFrame() int { return gif.currentFrame }

func DrawGifFrame(gif *GifImage, x int, y int, frame int, tint r.Color) {
	r.DrawTextureRec(*gif.TileSheet, r.NewRectangle(float32(gif.Width*frame), 0, float32(gif.Width), float32(gif.Height)), r.NewVector2(float32(x), float32(y)), tint)
}

func DrawGif(gif *GifImage, x int, y int, tint r.Color) {
	r.DrawTextureRec(*gif.TileSheet, r.NewRectangle(float32(gif.Width*gif.CurrentFrame()), 0, float32(gif.Width), float32(gif.Height)), r.NewVector2(float32(x), float32(y)), tint)
}

func getGifDimensions(gif *gif.GIF) (x, y int) {
	var lowestX int
	var lowestY int
	var highestX int
	var highestY int

	for _, img := range gif.Image {
		if img.Rect.Min.X < lowestX {
			lowestX = img.Rect.Min.X
		}
		if img.Rect.Min.Y < lowestY {
			lowestY = img.Rect.Min.Y
		}
		if img.Rect.Max.X > highestX {
			highestX = img.Rect.Max.X
		}
		if img.Rect.Max.Y > highestY {
			highestY = img.Rect.Max.Y
		}
	}

	return highestX - lowestX, highestY - lowestY
}
