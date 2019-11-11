package rgif

import (
	"fmt"
	"image"
	"image/draw"
	"image/gif"
	"os"

	r "github.com/lachee/raylib-goplus/raylib"
)

//GifImage represents a gif texture
type GifImage struct {
	//TileSheet is the backend Texture2D tilesheet that contains the frames of the gif
	TileSheet *r.Texture2D
	//Width is the width of a single frame
	Width int
	//Height is the height of a single frame
	Height int
	//Frames is the number of frames available
	Frames int
	//Timing is the delay (in 100ths of seconds) a frame has
	Timing        []int
	currentFrame  int
	lastFrameTime float32
}

//LoadGifFromFile loads a new gif
func LoadGifFromFile(fileName string) (*GifImage, error) {

	//Read the GIF file
	file, err := os.Open(fileName)
	defer file.Close()
	if err != nil {
		return nil, err
	}

	//Defer any panics
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("Error while decoding: %s", r)
		}
	}()

	//Decode teh gif
	gif, err := gif.DecodeAll(file)
	if err != nil {
		return nil, err
	}

	//Prepare the tilesheet and overpaint image.
	imgWidth, imgHeight := getGifDimensions(gif)
	frames := len(gif.Image)

	bounds := image.Rect(0, 0, imgWidth, imgHeight)
	tilesheet := image.NewRGBA(image.Rect(0, 0, imgWidth*frames, imgHeight))

	//Destination
	dest := bounds
	prev := dest
	previousUndisposedIndex := 0

	//Iterate over every frame of the gif
	for i, srcImg := range gif.Image {

		//Update the destinations
		dest = bounds.Add(image.Point{imgWidth * i, 0})
		prev = dest

		switch gif.Disposal[i] {
		default: //full image replacement
			draw.Draw(tilesheet, dest, srcImg, image.ZP, draw.Over) //Copy new frame

		case 0:
			fallthrough
		case 1: //Do not dispose
			draw.Draw(tilesheet, dest, tilesheet, prev.Min, draw.Over) //Copy previous frame on the tilesheet
			draw.Draw(tilesheet, dest, srcImg, image.ZP, draw.Over)    //Copy new frame
			previousUndisposedIndex = i

		case 2: //Restore to Background
			//TODO: Get this working. Golang Decoder doesn't have transparency information so I cannot mask.
			//bg := gif.Image[i].Palette[255-gif.BackgroundIndex]
			//draw.Draw(tilesheet, dest, &image.Uniform{bg}, image.ZP, draw.Over)
			draw.Draw(tilesheet, dest, srcImg, image.ZP, draw.Over) //Copy new frame

		case 3: //Restore to Previous
			//Copy the last previously undisposed frame
			draw.Draw(tilesheet, dest, gif.Image[previousUndisposedIndex], image.ZP, draw.Over) //Copy previous frame original
			draw.Draw(tilesheet, dest, srcImg, image.ZP, draw.Over)                             //Copy new frame

		}

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

//Unload removes the current underlying texture from memory
func (gif *GifImage) Unload() {
	gif.TileSheet.Unload()
}

//Step performs a time step.
func (gif *GifImage) Step(timeSinceLastStep float32) {

	gif.lastFrameTime += (timeSinceLastStep * 100)
	diff := gif.lastFrameTime - float32(gif.Timing[gif.currentFrame])

	if diff >= 0 {
		gif.NextFrame()
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

//CurrentRectangle gets the rectangle crop for the current frame
func (gif *GifImage) CurrentRectangle() r.Rectangle {
	return gif.GetRectangle(gif.currentFrame)
}

//CurrentTiming gets the current timing for the current frame
func (gif *GifImage) CurrentTiming() int { return gif.Timing[gif.currentFrame] }

//GetRectangle gets a rectangle crop for a specified frame
func (gif *GifImage) GetRectangle(frame int) r.Rectangle {
	return r.NewRectangle(float32(gif.Width*frame), 0, float32(gif.Width), float32(gif.Height))
}

//DrawGifFrame draws a single frame of a gif
func DrawGifFrame(gif *GifImage, x int, y int, frame int, tint r.Color) {
	r.DrawTextureRec(*gif.TileSheet, gif.GetRectangle(frame), r.NewVector2(float32(x), float32(y)), tint)
}

//DrawGif draws the current frame of a gif
func DrawGif(gif *GifImage, x int, y int, tint r.Color) {
	DrawGifFrame(gif, x, y, gif.currentFrame, tint)
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
