package raylib

//#include "raylib.h"
//#import <stdlib.h>
import "C"
import (
	"image"
	"unsafe"
)

// Image type, bpp always RGBA (32bit)
// NOTE: Data stored in CPU memory (RAM)
// Source: https://github.com/gen2brain/raylib-go/blob/master/raylib/raylib.go
type Image struct {
	// Image raw data
	data unsafe.Pointer
	// Image base width
	Width int32
	// Image base height
	Height int32
	// Mipmap levels, 1 by default
	Mipmaps int32
	// Data format (PixelFormat)
	Format PixelFormat
}

// NewImage Creates a new iamge
func NewImage(data []byte, width, height, mipmaps int32, format PixelFormat) *Image {
	d := unsafe.Pointer(&data[0])
	return &Image{d, width, height, mipmaps, format}
}

//NewImageFromGoImage Creates a new image from a Go Image
func NewImageFromGoImage(img image.Image) *Image {
	size := img.Bounds().Size()
	pixels := make([]Color, size.X*size.Y)

	for y := 0; y < size.Y; y++ {
		for x := 0; x < size.X; x++ {
			color := img.At(x, y)
			r, g, b, a := color.RGBA()
			pixels[x+y*size.X] = NewColor(uint8(r), uint8(g), uint8(b), uint8(a))
		}
	}

	return LoadImageEx(pixels, int32(size.X), int32(size.Y))
}

func newImageFromPointer(ptr unsafe.Pointer) *Image {
	return (*Image)(ptr)
}

// LoadImage Load an image into CPU memory (RAM)
func LoadImage(fileName string) *Image {
	cfileName := C.CString(fileName)
	defer C.free(unsafe.Pointer(cfileName))
	ret := C.LoadImage(cfileName)
	v := newImageFromPointer(unsafe.Pointer(&ret))
	return v
}

// LoadImageEx Load image data from Color array data (RGBA - 32bit)
func LoadImageEx(pixels []Color, width, height int32) *Image {
	cpixels := pixels[0].cptr()
	cwidth := (C.int)(width)
	cheight := (C.int)(height)
	ret := C.LoadImageEx(cpixels, cwidth, cheight)
	v := newImageFromPointer(unsafe.Pointer(&ret))
	return v
}

//Texture2D stores GPU based textures.
type Texture2D struct {
	Id      uint32
	Width   int32
	Height  int32
	Mipmaps int32
	Format  int32
}

func newTextureFromPointer(ptr unsafe.Pointer) Texture2D {
	return *(*Texture2D)(ptr)
}

//LoadTexture from a file into GPU memory
func LoadTexture(filename string) Texture2D {
	cs := C.CString(filename)
	defer C.free(unsafe.Pointer(cs))
	ct2d := C.LoadTexture(cs)
	return newTextureFromPointer(unsafe.Pointer(&ct2d))
}

//LoadTextureFromImage loads image data
func LoadTextureFromImage(image Image) Texture2D {
	i := *image.cptr()
	ct2d := C.LoadTextureFromImage(i)
	return newTextureFromPointer(unsafe.Pointer(&ct2d))
}

//LoadTextureFromGoImage loads image data from image.Image. Uses NewImageFromGoImage.
func LoadTextureFromGoImage(image image.Image) Texture2D {
	return LoadTextureFromImage(*NewImageFromGoImage(image))
}

// PixelFormat - Texture format
type PixelFormat int32

// Texture formats
// NOTE: Support depends on OpenGL version and platform
const (
	// 8 bit per pixel (no alpha)
	UncompressedGrayscale PixelFormat = iota + 1
	// 8*2 bpp (2 channels)
	UncompressedGrayAlpha
	// 16 bpp
	UncompressedR5g6b5
	// 24 bpp
	UncompressedR8g8b8
	// 16 bpp (1 bit alpha)
	UncompressedR5g5b5a1
	// 16 bpp (4 bit alpha)
	UncompressedR4g4b4a4
	// 32 bpp
	UncompressedR8g8b8a8
	// 32 bpp (1 channel - float)
	UncompressedR32
	// 32*3 bpp (3 channels - float)
	UncompressedR32g32b32
	// 32*4 bpp (4 channels - float)
	UncompressedR32g32b32a32
	// 4 bpp (no alpha)
	CompressedDxt1Rgb
	// 4 bpp (1 bit alpha)
	CompressedDxt1Rgba
	// 8 bpp
	CompressedDxt3Rgba
	// 8 bpp
	CompressedDxt5Rgba
	// 4 bpp
	CompressedEtc1Rgb
	// 4 bpp
	CompressedEtc2Rgb
	// 8 bpp
	CompressedEtc2EacRgba
	// 4 bpp
	CompressedPvrtRgb
	// 4 bpp
	CompressedPvrtRgba
	// 8 bpp
	CompressedAstc4x4Rgba
	// 2 bpp
	CompressedAstc8x8Rgba
)
