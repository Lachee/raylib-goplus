package rgif

import (
	"fmt"
	"image/gif"
	"os"

	r "github.com/lachee/raylib-goplus/raylib"
)

type FrameDisposal int

const (
	FrameDisposalNone FrameDisposal = iota
	FrameDisposalDontDispose
	FrameDisposalReplace
	FrameDisposalRestorePrevious
)

//GifImage represents a gif texture
type GifImage struct {

	//currently loaded textures
	textures     []*r.Texture2D
	textureCount int

	//cache of images
	images []*r.Image

	//Width is the width of a single frame
	Width int
	//Height is the height of a single frame
	Height int
	//Frames is the number of frames available
	Frames int
	//Timing is the delay (in 100ths of seconds) a frame has
	Timing []int
	//Disposal is the disposal for each frame
	Disposal []FrameDisposal

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

	disposals := make([]FrameDisposal, frames)
	images := make([]*r.Image, frames)
	textures := make([]*r.Texture2D, frames)
	for i, img := range gif.Image {
		disposals[i] = FrameDisposal(gif.Disposal[i])
		images[i] = r.LoadImageFromGo(img)
	}

	//Load the first initial texture
	textures[0] = r.LoadTextureFromImage(images[0])

	return &GifImage{
		textures:     textures,
		textureCount: 0,
		images:       images,
		Width:        imgWidth,
		Height:       imgHeight,
		Frames:       frames,
		Timing:       gif.Delay,
		Disposal:     disposals,
	}, nil
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

	gif.rollingDisposal()
}

func (gif *GifImage) rollingDisposal() {
	if gif.currentFrame == 0 {
		gif.unloadCurrentFrames(0)
		gif.textures[0] = r.UnregisteredLoadTextureFromImage(gif.images[gif.currentFrame])
		gif.textureCount = 0
	} else {

		switch gif.Disposal[gif.currentFrame] {
		default:
			fallthrough

		//Replace each frame with the new one
		case FrameDisposalReplace:
			gif.unloadCurrentFrames(0)
			gif.textures[0] = r.UnregisteredLoadTextureFromImage(gif.images[gif.currentFrame])
			gif.textureCount = 0

		case FrameDisposalNone:
			fallthrough
		case FrameDisposalDontDispose:
			if gif.textureCount < gif.currentFrame {
				gif.textureCount++
				gif.textures[gif.textureCount] = r.UnregisteredLoadTextureFromImage(gif.images[gif.currentFrame])
			}
		case FrameDisposalRestorePrevious:
			gif.unloadCurrentFrames(1)
			gif.textureCount++
			gif.textures[gif.textureCount] = r.UnregisteredLoadTextureFromImage(gif.images[gif.currentFrame])
		}
	}
}

//unloadCurrentFrames unloads all the current texture frames
func (gif *GifImage) unloadCurrentFrames(offset int) {
	for i := offset; i < gif.textureCount+1; i++ {
		r.UnregisteredUnloadTexture(gif.textures[i])
		gif.textures[i] = nil
	}
	gif.textureCount = offset - 1
}

//Reset clears the last frame time and resets the current frame to zero
func (gif *GifImage) Reset() {
	gif.currentFrame = 0
	gif.lastFrameTime = 0
	gif.unloadCurrentFrames(0)
}

//Unload unloads all the textures and images, making this gif unusable.
func (gif *GifImage) Unload() {
	for _, t := range gif.textures {
		if t != nil {
			r.UnregisteredUnloadTexture(t)
		}
	}

	for _, i := range gif.images {
		if i != nil {
			i.Unload()
		}
	}

	gif.textures = nil
	gif.images = nil
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

//DrawGif draws a single frame of a gif
func DrawGif(gif *GifImage, x int, y int, tint r.Color) {
	for i := 0; i <= gif.textureCount; i++ {
		r.DrawTexture(*gif.textures[i], x, y, tint)
	}
}

/*
//DrawGifFrame draws a single frame of a gif
func DrawGifFrame(gif *GifImage, x int, y int, frame int, tint r.Color) {
	r.DrawTextureRec(*gif.TileSheet, gif.GetRectangle(frame), r.NewVector2(float32(x), float32(y)), tint)
}

//DrawGifFrameV draws a single frame of a gif at a vector position
func DrawGifFrameV(gif *GifImage, position r.Vector2, frame int, tint r.Color) {
	r.DrawTextureRec(*gif.TileSheet, gif.GetRectangle(frame), position, tint)
}

//DrawGifFrameEx draws a frame of a gif at a vector position, rotation and scale
func DrawGifFrameEx(gif *GifImage, position r.Vector2, rotation float32, scale float32, frame int, tint r.Color) {
	w := float32(gif.Width) * scale
	h := float32(gif.Height) * scale
	rect := gif.GetRectangle(frame)
	origin := r.NewVector2(w/2.0, h/2.0)
	r.DrawTexturePro(*gif.TileSheet, rect, r.NewRectangle(position.X, position.Y, w, h), origin, rotation, tint)
}

//DrawGif draws the current frame of a gif
func DrawGif(gif *GifImage, x int, y int, tint r.Color) {
	DrawGifFrame(gif, x, y, gif.currentFrame, tint)
}

//DrawGifV draws the current frame of a gif at a vector position
func DrawGifV(gif *GifImage, position r.Vector2, tint r.Color) {
	DrawGifFrameV(gif, position, gif.currentFrame, tint)
}

//DrawGifEx draws the current frame of a gif at a vector position, rotation and scale
func DrawGifEx(gif *GifImage, position r.Vector2, rotation float32, scale float32, tint r.Color) {
	DrawGifFrameEx(gif, position, rotation, scale, gif.currentFrame, tint)
}
*/

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
