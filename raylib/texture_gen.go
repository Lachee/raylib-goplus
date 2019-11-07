package raylib

/*
#include "raylib.h"
#include <stdlib.h>
#include "go.h"
*/
import "C"
import "unsafe"

//LoadImage : Load image from file into CPU memory (RAM)
func LoadImage(fileName string) *Image {
	cfileName := C.CString(fileName)
	defer C.free(unsafe.Pointer(cfileName))
	res := C.LoadImage(cfileName)
	retval := newImageFromPointer(unsafe.Pointer(&res))
	addUnloadable(retval)
	return retval
}

// LoadImageEx Load image data from Color array data (RGBA - 32bit)
func LoadImageEx(pixels []Color, width, height int32) *Image {
	cpixels := pixels[0].cptr()
	cwidth := (C.int)(width)
	cheight := (C.int)(height)
	ret := C.LoadImageEx(cpixels, cwidth, cheight) //We call the C function as raylib uses stb_image to load images.
	v := newImageFromPointer(unsafe.Pointer(&ret))
	addUnloadable(v)
	return v
}

//LoadImagePro loads raw data wtih parameters
func LoadImagePro(pixels []byte, width, height int32, format PixelFormat) *Image {
	data := unsafe.Pointer(&pixels[0])
	res := C.LoadImagePro(data, C.int(width), C.int(height), C.int(format))
	v := newImageFromPointer(unsafe.Pointer(&res))
	addUnloadable(v)
	return v
}

//LoadImageRaw : Load image from RAW file data
func LoadImageRaw(fileName string, width int, height int, format int, headerSize int) *Image {
	cfileName := C.CString(fileName)
	defer C.free(unsafe.Pointer(cfileName))
	res := C.LoadImageRaw(cfileName, C.int(int32(width)), C.int(int32(height)), C.int(int32(format)), C.int(int32(headerSize)))
	retval := newImageFromPointer(unsafe.Pointer(&res))
	addUnloadable(retval)
	return retval
}

//Export : Export image data to file
func (image *Image) Export(fileName string) {
	cfileName := C.CString(fileName)
	defer C.free(unsafe.Pointer(cfileName))
	cimage := *image.cptr()
	C.ExportImage(cimage, cfileName)
}

//ExportImage : Export image data to file
//Recommended to use image.Export(fileName) instead
func ExportImage(image *Image, fileName string) {
	image.Export(fileName)
}

//ExportAsCode : Export image as code file defining an array of bytes
func (image *Image) ExportAsCode(fileName string) {
	cfileName := C.CString(fileName)
	defer C.free(unsafe.Pointer(cfileName))
	cimage := *image.cptr()
	C.ExportImageAsCode(cimage, cfileName)
}

//ExportImageAsCode : Export image as code file defining an array of bytes
//Recommended to use image.ExportAsCode(fileName) instead
func ExportImageAsCode(image *Image, fileName string) {
	image.ExportAsCode(fileName)
}

//LoadTexture : Load texture from file into GPU memory (VRAM)
func LoadTexture(fileName string) *Texture2D {
	cfileName := C.CString(fileName)
	defer C.free(unsafe.Pointer(cfileName))
	res := C.LoadTexture(cfileName)
	retval := newTexture2DFromPointer(unsafe.Pointer(&res))
	addUnloadable(retval)
	return retval
}

//LoadTextureFromImage : Load texture from image data
func LoadTextureFromImage(image *Image) *Texture2D {
	cimage := *image.cptr()
	res := C.LoadTextureFromImage(cimage)
	retval := newTexture2DFromPointer(unsafe.Pointer(&res))
	addUnloadable(retval)
	return retval
}

//LoadTextureCubemap : Load cubemap from image, multiple image cubemap layouts supported
func LoadTextureCubemap(image *Image, layoutType CubemapLayoutType) *TextureCubemap {
	cimage := *image.cptr()
	res := C.LoadTextureCubemap(cimage, C.int(int32(layoutType)))
	retval := newTextureCubemapFromPointer(unsafe.Pointer(&res))
	addUnloadable(retval)
	return retval
}

//LoadRenderTexture : Load texture for rendering (framebuffer)
func LoadRenderTexture(width int, height int) *RenderTexture2D {
	res := C.LoadRenderTexture(C.int(int32(width)), C.int(int32(height)))
	retval := newRenderTexture2DFromPointer(unsafe.Pointer(&res))
	addUnloadable(retval)
	return retval
}

//Unload : Unload image from CPU memory (RAM)
func (image *Image) Unload() {
	cimage := *image.cptr()
	C.UnloadImage(cimage)
	removeUnloadable(image)
}

//UnloadImage : Unload image from CPU memory (RAM)
//Recommended to use image.Unload() instead
func UnloadImage(image *Image) {
	image.Unload()
}

//Unload : Unload texture from GPU memory (VRAM)
func (texture *Texture2D) Unload() {
	ctexture := *texture.cptr()
	C.UnloadTexture(ctexture)
	removeUnloadable(texture)
}

//UnloadTexture : Unload texture from GPU memory (VRAM)
//Recommended to use texture.Unload() instead
func UnloadTexture(texture *Texture2D) {
	texture.Unload()
}

//Unload : Unload render texture from GPU memory (VRAM)
func (target *RenderTexture2D) Unload() {
	ctarget := *target.cptr()
	C.UnloadRenderTexture(ctarget)
	removeUnloadable(target)
}

//UnloadRenderTexture : Unload render texture from GPU memory (VRAM)
//Recommended to use target.Unload() instead
func UnloadRenderTexture(target *RenderTexture2D) {
	target.Unload()
}

//GetAlphaBorder : Get image alpha border rectangle
func (image *Image) GetAlphaBorder(threshold float32) Rectangle {
	cimage := *image.cptr()
	res := C.GetImageAlphaBorder(cimage, C.float(threshold))
	return newRectangleFromPointer(unsafe.Pointer(&res))
}

//GetImageAlphaBorder : Get image alpha border rectangle
//Recommended to use image.GetAlphaBorder(threshold) instead
func GetImageAlphaBorder(image *Image, threshold float32) Rectangle {
	return image.GetAlphaBorder(threshold)
}

//GetPixelDataSize : Get pixel data size in bytes (image or texture)
func GetPixelDataSize(width int, height int, format PixelFormat) int {
	res := C.GetPixelDataSize(C.int(int32(width)), C.int(int32(height)), C.int(format))
	return int(int32(res))
}

//GetTextureData : Get pixel data from GPU texture and return an Image
func (texture *Texture2D) GetTextureData() *Image {
	ctexture := *texture.cptr()
	res := C.GetTextureData(ctexture)
	retval := newImageFromPointer(unsafe.Pointer(&res))
	addUnloadable(retval)
	return retval
}

//GetTextureData : Get pixel data from GPU texture and return an Image
//Recommended to use texture.GetTextureData() instead
func GetTextureData(texture *Texture2D) *Image {
	return texture.GetTextureData()
}

//GetScreenData : Get pixel data from screen buffer and return an Image (screenshot)
func GetScreenData() *Image {
	res := C.GetScreenData()
	retval := newImageFromPointer(unsafe.Pointer(&res))
	addUnloadable(retval)
	return retval
}

//UpdateTexture : Update GPU texture with new data
func (texture *Texture2D) UpdateTexture(pixels unsafe.Pointer) {
	ctexture := *texture.cptr()
	C.UpdateTexture(ctexture, pixels)
}

//UpdateTexture : Update GPU texture with new data
//Recommended to use texture.UpdateTexture(pixels) instead
func UpdateTexture(texture *Texture2D, pixels unsafe.Pointer) {
	texture.UpdateTexture(pixels)
}

//Copy : Create an image duplicate (useful for transformations)
func (image *Image) Copy() *Image {
	cimage := *image.cptr()
	res := C.ImageCopy(cimage)
	retval := newImageFromPointer(unsafe.Pointer(&res))
	addUnloadable(retval)
	return retval
}

//ImageCopy : Create an image duplicate (useful for transformations)
//Recommended to use image.Copy() instead
func ImageCopy(image *Image) *Image {
	return image.Copy()
}

//FromImage : Create an image from another image piece
func (image *Image) FromImage(rec Rectangle) *Image {
	crec := *rec.cptr()
	cimage := *image.cptr()
	res := C.ImageFromImage(cimage, crec)
	v := newImageFromPointer(unsafe.Pointer(&res))
	addUnloadable(v)
	return v
}

//ImageFromImage : Create an image from another image piece
//Recommended to use image.(rec) instead
func ImageFromImage(image *Image, rec Rectangle) *Image {
	return image.FromImage(rec)
}

//ToPOT : Convert image to POT (power-of-two)
func (image *Image) ToPOT(fillColor Color) {
	cfillColor := *fillColor.cptr()
	cimage := image.cptr()
	C.ImageToPOT(cimage, cfillColor)
}

//ImageToPOT : Convert image to POT (power-of-two)
//Recommended to use image.ToPOT(fillColor) instead
func ImageToPOT(image *Image, fillColor Color) {
	image.ToPOT(fillColor)
}

//Format : Convert image data to desired format
func (image *Image) SetFormat(newFormat PixelFormat) {
	cimage := image.cptr()
	C.ImageFormat(cimage, C.int(newFormat))
}

//ImageFormat : Convert image data to desired format
//Recommended to use image.SetFormat(newFormat) instead
func ImageFormat(image *Image, newFormat PixelFormat) {
	image.SetFormat(newFormat)
}

//AlphaMask : Apply alpha mask to image
func (image *Image) AlphaMask(alphaMask *Image) {
	calphaMask := *alphaMask.cptr()
	cimage := image.cptr()
	C.ImageAlphaMask(cimage, calphaMask)
}

//ImageAlphaMask : Apply alpha mask to image
//Recommended to use image.AlphaMask(alphaMask) instead
func ImageAlphaMask(image *Image, alphaMask *Image) {
	image.AlphaMask(alphaMask)
}

//AlphaClear : Clear alpha channel to desired color
func (image *Image) AlphaClear(color Color, threshold float32) {
	ccolor := *color.cptr()
	cimage := image.cptr()
	C.ImageAlphaClear(cimage, ccolor, C.float(threshold))
}

//ImageAlphaClear : Clear alpha channel to desired color
//Recommended to use image.AlphaClear(color, threshold) instead
func ImageAlphaClear(image *Image, color Color, threshold float32) {
	image.AlphaClear(color, threshold)
}

//AlphaCrop : Crop image depending on alpha value
func (image *Image) AlphaCrop(threshold float32) {
	cimage := image.cptr()
	C.ImageAlphaCrop(cimage, C.float(threshold))
}

//ImageAlphaCrop : Crop image depending on alpha value
//Recommended to use image.AlphaCrop(threshold) instead
func ImageAlphaCrop(image *Image, threshold float32) {
	image.AlphaCrop(threshold)
}

//AlphaPremultiply : Premultiply alpha channel
func (image *Image) AlphaPremultiply() {
	cimage := image.cptr()
	C.ImageAlphaPremultiply(cimage)
}

//ImageAlphaPremultiply : Premultiply alpha channel
//Recommended to use image.AlphaPremultiply() instead
func ImageAlphaPremultiply(image *Image) {
	image.AlphaPremultiply()
}

//Crop : Crop an image to a defined rectangle
func (image *Image) Crop(crop Rectangle) {
	ccrop := *crop.cptr()
	cimage := image.cptr()
	C.ImageCrop(cimage, ccrop)
}

//ImageCrop : Crop an image to a defined rectangle
//Recommended to use image.Crop(crop) instead
func ImageCrop(image *Image, crop Rectangle) {
	image.Crop(crop)
}

//Resize : Resize image (Bicubic scaling algorithm)
func (image *Image) Resize(newWidth int, newHeight int) {
	cimage := image.cptr()
	C.ImageResize(cimage, C.int(int32(newWidth)), C.int(int32(newHeight)))
}

//ImageResize : Resize image (Bicubic scaling algorithm)
//Recommended to use image.Resize(newWidth, newHeight) instead
func ImageResize(image *Image, newWidth int, newHeight int) {
	image.Resize(newWidth, newHeight)
}

//ResizeNN : Resize image (Nearest-Neighbor scaling algorithm)
func (image *Image) ResizeNN(newWidth int, newHeight int) {
	cimage := image.cptr()
	C.ImageResizeNN(cimage, C.int(int32(newWidth)), C.int(int32(newHeight)))
}

//ImageResizeNN : Resize image (Nearest-Neighbor scaling algorithm)
//Recommended to use image.ResizeNN(newWidth, newHeight) instead
func ImageResizeNN(image *Image, newWidth int, newHeight int) {
	image.ResizeNN(newWidth, newHeight)
}

//ResizeCanvas : Resize canvas and fill with color
func (image *Image) ResizeCanvas(newWidth int, newHeight int, offsetX int, offsetY int, color Color) {
	ccolor := *color.cptr()
	cimage := image.cptr()
	C.ImageResizeCanvas(cimage, C.int(int32(newWidth)), C.int(int32(newHeight)), C.int(int32(offsetX)), C.int(int32(offsetY)), ccolor)
}

//ImageResizeCanvas : Resize canvas and fill with color
//Recommended to use image.ResizeCanvas(newWidth, newHeight, offsetX, offsetY, color) instead
func ImageResizeCanvas(image *Image, newWidth int, newHeight int, offsetX int, offsetY int, color Color) {
	image.ResizeCanvas(newWidth, newHeight, offsetX, offsetY, color)
}

//CreateMipmaps : Generate all mipmap levels for a provided image
func (image *Image) CreateMipmaps() {
	cimage := image.cptr()
	C.ImageMipmaps(cimage)
}

//ImageMipmaps : Generate all mipmap levels for a provided image
//Recommended to use image.CreateMipmaps() instead
func ImageMipmaps(image *Image) {
	image.CreateMipmaps()
}

//Dither : Dither image data to 16bpp or lower (Floyd-Steinberg dithering)
func (image *Image) Dither(rBpp int, gBpp int, bBpp int, aBpp int) {
	cimage := image.cptr()
	C.ImageDither(cimage, C.int(int32(rBpp)), C.int(int32(gBpp)), C.int(int32(bBpp)), C.int(int32(aBpp)))
}

//ImageDither : Dither image data to 16bpp or lower (Floyd-Steinberg dithering)
//Recommended to use image.Dither(rBpp, gBpp, bBpp, aBpp) instead
func ImageDither(image *Image, rBpp int, gBpp int, bBpp int, aBpp int) {
	image.Dither(rBpp, gBpp, bBpp, aBpp)
}

//ImageText : Create an image from text (default font)
func ImageText(text string, fontSize int, color Color) *Image {
	ccolor := *color.cptr()
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	res := C.ImageText(ctext, C.int(int32(fontSize)), ccolor)
	retval := newImageFromPointer(unsafe.Pointer(&res))
	addUnloadable(retval)
	return retval
}

//ImageTextEx : Create an image from text (custom sprite font)
func (font *Font) ImageTextEx(text string, fontSize float32, spacing float32, tint Color) *Image {
	ctint := *tint.cptr()
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	cfont := *font.cptr()
	res := C.ImageTextEx(cfont, ctext, C.float(fontSize), C.float(spacing), ctint)
	retval := newImageFromPointer(unsafe.Pointer(&res))
	addUnloadable(retval)
	return retval
}

//ImageTextEx : Create an image from text (custom sprite font)
//Recommended to use font.ImageTextEx(text, fontSize, spacing, tint) instead
func ImageTextEx(font *Font, text string, fontSize float32, spacing float32, tint Color) *Image {
	return font.ImageTextEx(text, fontSize, spacing, tint)
}

//Draw : Draw a source image within a destination image (tint applied to source)
func (dst *Image) Draw(src *Image, srcRec Rectangle, dstRec Rectangle, tint Color) {
	ctint := *tint.cptr()
	cdstRec := *dstRec.cptr()
	csrcRec := *srcRec.cptr()
	csrc := *src.cptr()
	cdst := dst.cptr()
	C.ImageDraw(cdst, csrc, csrcRec, cdstRec, ctint)
}

//ImageDraw : Draw a source image within a destination image (tint applied to source)
//Recommended to use dst.Draw(src, srcRec, dstRec, tint) instead
func ImageDraw(dst *Image, src *Image, srcRec Rectangle, dstRec Rectangle, tint Color) {
	dst.Draw(src, srcRec, dstRec, tint)
}

//DrawRectangle : Draw rectangle within an image
func (dst *Image) DrawRectangle(rec Rectangle, color Color) {
	ccolor := *color.cptr()
	crec := *rec.cptr()
	cdst := dst.cptr()
	C.ImageDrawRectangle(cdst, crec, ccolor)
}

//ImageDrawRectangle : Draw rectangle within an image
//Recommended to use dst.DrawRectangle(rec, color) instead
func ImageDrawRectangle(dst *Image, rec Rectangle, color Color) {
	dst.DrawRectangle(rec, color)
}

//DrawRectangleLines : Draw rectangle lines within an image
func (dst *Image) DrawRectangleLines(rec Rectangle, thick int, color Color) {
	ccolor := *color.cptr()
	crec := *rec.cptr()
	cdst := dst.cptr()
	C.ImageDrawRectangleLines(cdst, crec, C.int(int32(thick)), ccolor)
}

//ImageDrawRectangleLines : Draw rectangle lines within an image
//Recommended to use dst.DrawRectangleLines(rec, thick, color) instead
func ImageDrawRectangleLines(dst *Image, rec Rectangle, thick int, color Color) {
	dst.DrawRectangleLines(rec, thick, color)
}

//DrawText : Draw text (default font) within an image (destination)
func (dst *Image) DrawText(position Vector2, text string, fontSize int, color Color) {
	ccolor := *color.cptr()
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	cposition := *position.cptr()
	cdst := dst.cptr()
	C.ImageDrawText(cdst, cposition, ctext, C.int(int32(fontSize)), ccolor)
}

//ImageDrawText : Draw text (default font) within an image (destination)
//Recommended to use dst.DrawText(position, text, fontSize, color) instead
func ImageDrawText(dst *Image, position Vector2, text string, fontSize int, color Color) {
	dst.DrawText(position, text, fontSize, color)
}

//DrawTextEx : Draw text (custom sprite font) within an image (destination)
func (dst *Image) DrawTextEx(position Vector2, font *Font, text string, fontSize float32, spacing float32, color Color) {
	ccolor := *color.cptr()
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	cfont := *font.cptr()
	cposition := *position.cptr()
	cdst := dst.cptr()
	C.ImageDrawTextEx(cdst, cposition, cfont, ctext, C.float(fontSize), C.float(spacing), ccolor)
}

//ImageDrawTextEx : Draw text (custom sprite font) within an image (destination)
//Recommended to use dst.DrawTextEx(position, font, text, fontSize, spacing, color) instead
func ImageDrawTextEx(dst *Image, position Vector2, font *Font, text string, fontSize float32, spacing float32, color Color) {
	dst.DrawTextEx(position, font, text, fontSize, spacing, color)
}

//FlipVertical : Flip image vertically
func (image *Image) FlipVertical() {
	cimage := image.cptr()
	C.ImageFlipVertical(cimage)
}

//ImageFlipVertical : Flip image vertically
//Recommended to use image.FlipVertical() instead
func ImageFlipVertical(image *Image) {
	image.FlipVertical()
}

//FlipHorizontal : Flip image horizontally
func (image *Image) FlipHorizontal() {
	cimage := image.cptr()
	C.ImageFlipHorizontal(cimage)
}

//ImageFlipHorizontal : Flip image horizontally
//Recommended to use image.FlipHorizontal() instead
func ImageFlipHorizontal(image *Image) {
	image.FlipHorizontal()
}

//RotateCW : Rotate image clockwise 90deg
func (image *Image) RotateCW() {
	cimage := image.cptr()
	C.ImageRotateCW(cimage)
}

//ImageRotateCW : Rotate image clockwise 90deg
//Recommended to use image.RotateCW() instead
func ImageRotateCW(image *Image) {
	image.RotateCW()
}

//RotateCCW : Rotate image counter-clockwise 90deg
func (image *Image) RotateCCW() {
	cimage := image.cptr()
	C.ImageRotateCCW(cimage)
}

//ImageRotateCCW : Rotate image counter-clockwise 90deg
//Recommended to use image.RotateCCW() instead
func ImageRotateCCW(image *Image) {
	image.RotateCCW()
}

//ColorTint : Modify image color: tint
func (image *Image) ColorTint(color Color) {
	ccolor := *color.cptr()
	cimage := image.cptr()
	C.ImageColorTint(cimage, ccolor)
}

//ImageColorTint : Modify image color: tint
//Recommended to use image.ColorTint(color) instead
func ImageColorTint(image *Image, color Color) {
	image.ColorTint(color)
}

//ColorInvert : Modify image color: invert
func (image *Image) ColorInvert() {
	cimage := image.cptr()
	C.ImageColorInvert(cimage)
}

//ImageColorInvert : Modify image color: invert
//Recommended to use image.ColorInvert() instead
func ImageColorInvert(image *Image) {
	image.ColorInvert()
}

//ColorGrayscale : Modify image color: grayscale
func (image *Image) ColorGrayscale() {
	cimage := image.cptr()
	C.ImageColorGrayscale(cimage)
}

//ImageColorGrayscale : Modify image color: grayscale
//Recommended to use image.ColorGrayscale() instead
func ImageColorGrayscale(image *Image) {
	image.ColorGrayscale()
}

//ColorContrast : Modify image color: contrast (-100 to 100)
func (image *Image) ColorContrast(contrast float32) {
	cimage := image.cptr()
	C.ImageColorContrast(cimage, C.float(contrast))
}

//ImageColorContrast : Modify image color: contrast (-100 to 100)
//Recommended to use image.ColorContrast(contrast) instead
func ImageColorContrast(image *Image, contrast float32) {
	image.ColorContrast(contrast)
}

//ColorBrightness : Modify image color: brightness (-255 to 255)
func (image *Image) ColorBrightness(brightness int) {
	cimage := image.cptr()
	C.ImageColorBrightness(cimage, C.int(int32(brightness)))
}

//ImageColorBrightness : Modify image color: brightness (-255 to 255)
//Recommended to use image.ColorBrightness(brightness) instead
func ImageColorBrightness(image *Image, brightness int) {
	image.ColorBrightness(brightness)
}

//ColorReplace : Modify image color: replace color
func (image *Image) ColorReplace(color Color, replace Color) {
	creplace := *replace.cptr()
	ccolor := *color.cptr()
	cimage := image.cptr()
	C.ImageColorReplace(cimage, ccolor, creplace)
}

//ImageColorReplace : Modify image color: replace color
//Recommended to use image.ColorReplace(color, replace) instead
func ImageColorReplace(image *Image, color Color, replace Color) {
	image.ColorReplace(color, replace)
}

//GenImageColor : Generate image: plain color
func GenImageColor(width int, height int, color Color) *Image {
	ccolor := *color.cptr()
	res := C.GenImageColor(C.int(int32(width)), C.int(int32(height)), ccolor)
	retval := newImageFromPointer(unsafe.Pointer(&res))
	addUnloadable(retval)
	return retval
}

//GenImageGradientV : Generate image: vertical gradient
func GenImageGradientV(width int, height int, top Color, bottom Color) *Image {
	cbottom := *bottom.cptr()
	ctop := *top.cptr()
	res := C.GenImageGradientV(C.int(int32(width)), C.int(int32(height)), ctop, cbottom)
	retval := newImageFromPointer(unsafe.Pointer(&res))
	addUnloadable(retval)
	return retval
}

//GenImageGradientH : Generate image: horizontal gradient
func GenImageGradientH(width int, height int, left Color, right Color) *Image {
	cright := *right.cptr()
	cleft := *left.cptr()
	res := C.GenImageGradientH(C.int(int32(width)), C.int(int32(height)), cleft, cright)
	retval := newImageFromPointer(unsafe.Pointer(&res))
	addUnloadable(retval)
	return retval
}

//GenImageGradientRadial : Generate image: radial gradient
func GenImageGradientRadial(width int, height int, density float32, inner Color, outer Color) *Image {
	couter := *outer.cptr()
	cinner := *inner.cptr()
	res := C.GenImageGradientRadial(C.int(int32(width)), C.int(int32(height)), C.float(density), cinner, couter)
	retval := newImageFromPointer(unsafe.Pointer(&res))
	addUnloadable(retval)
	return retval
}

//GenImageChecked : Generate image: checked
func GenImageChecked(width int, height int, checksX int, checksY int, col1 Color, col2 Color) *Image {
	ccol2 := *col2.cptr()
	ccol1 := *col1.cptr()
	res := C.GenImageChecked(C.int(int32(width)), C.int(int32(height)), C.int(int32(checksX)), C.int(int32(checksY)), ccol1, ccol2)
	retval := newImageFromPointer(unsafe.Pointer(&res))
	addUnloadable(retval)
	return retval
}

//GenImageWhiteNoise : Generate image: white noise
func GenImageWhiteNoise(width int, height int, factor float32) *Image {
	res := C.GenImageWhiteNoise(C.int(int32(width)), C.int(int32(height)), C.float(factor))
	retval := newImageFromPointer(unsafe.Pointer(&res))
	addUnloadable(retval)
	return retval
}

//GenImagePerlinNoise : Generate image: perlin noise
func GenImagePerlinNoise(width int, height int, offsetX int, offsetY int, scale float32) *Image {
	res := C.GenImagePerlinNoise(C.int(int32(width)), C.int(int32(height)), C.int(int32(offsetX)), C.int(int32(offsetY)), C.float(scale))
	retval := newImageFromPointer(unsafe.Pointer(&res))
	addUnloadable(retval)
	return retval
}

//GenImageCellular : Generate image: cellular algorithm. Bigger tileSize means bigger cells
func GenImageCellular(width int, height int, tileSize int) *Image {
	res := C.GenImageCellular(C.int(int32(width)), C.int(int32(height)), C.int(int32(tileSize)))
	retval := newImageFromPointer(unsafe.Pointer(&res))
	addUnloadable(retval)
	return retval
}

//GenTextureMipmaps : Generate GPU mipmaps for a texture
func (texture *Texture2D) GenTextureMipmaps() {
	ctexture := texture.cptr()
	C.GenTextureMipmaps(ctexture)
}

//GenTextureMipmaps : Generate GPU mipmaps for a texture
//Recommended to use texture.GenTextureMipmaps() instead
func GenTextureMipmaps(texture *Texture2D) {
	texture.GenTextureMipmaps()
}

//SetTextureFilter : Set texture scaling filter mode
func (texture *Texture2D) SetTextureFilter(filterMode int) {
	ctexture := *texture.cptr()
	C.SetTextureFilter(ctexture, C.int(int32(filterMode)))
}

//SetTextureFilter : Set texture scaling filter mode
//Recommended to use texture.SetTextureFilter(filterMode) instead
func SetTextureFilter(texture *Texture2D, filterMode int) {
	texture.SetTextureFilter(filterMode)
}

//SetWrap : Set texture wrapping mode
func (texture *Texture2D) SetWrap(wrapMode TextureWrapMode) {
	ctexture := *texture.cptr()
	C.SetTextureWrap(ctexture, C.int(int32(wrapMode)))
}

//SetTextureWrap : Set texture wrapping mode
//Recommended to use texture.SetWrap(wrapMode) instead
func SetTextureWrap(texture *Texture2D, wrapMode TextureWrapMode) {
	texture.SetWrap(wrapMode)
}

//DrawTexture : Draw a Texture2D
func DrawTexture(texture Texture2D, posX int, posY int, tint Color) {
	ctint := *tint.cptr()
	ctexture := *texture.cptr()
	C.DrawTexture(ctexture, C.int(int32(posX)), C.int(int32(posY)), ctint)
}

//DrawTextureV : Draw a Texture2D with position defined as Vector2
func DrawTextureV(texture Texture2D, position Vector2, tint Color) {
	ctint := *tint.cptr()
	cposition := *position.cptr()
	ctexture := *texture.cptr()
	C.DrawTextureV(ctexture, cposition, ctint)
}

//DrawTextureEx : Draw a Texture2D with extended parameters
func DrawTextureEx(texture Texture2D, position Vector2, rotation float32, scale float32, tint Color) {
	ctint := *tint.cptr()
	cposition := *position.cptr()
	ctexture := *texture.cptr()
	C.DrawTextureEx(ctexture, cposition, C.float(rotation), C.float(scale), ctint)
}

//DrawTextureRec : Draw a part of a texture defined by a rectangle
func DrawTextureRec(texture Texture2D, sourceRec Rectangle, position Vector2, tint Color) {
	ctint := *tint.cptr()
	cposition := *position.cptr()
	csourceRec := *sourceRec.cptr()
	ctexture := *texture.cptr()
	C.DrawTextureRec(ctexture, csourceRec, cposition, ctint)
}

//DrawTextureQuad : Draw texture quad with tiling and offset parameters
func DrawTextureQuad(texture Texture2D, tiling Vector2, offset Vector2, quad Rectangle, tint Color) {
	ctint := *tint.cptr()
	cquad := *quad.cptr()
	coffset := *offset.cptr()
	ctiling := *tiling.cptr()
	ctexture := *texture.cptr()
	C.DrawTextureQuad(ctexture, ctiling, coffset, cquad, ctint)
}

//DrawTexturePro : Draw a part of a texture defined by a rectangle with 'pro' parameters
func DrawTexturePro(texture Texture2D, sourceRec Rectangle, destRec Rectangle, origin Vector2, rotation float32, tint Color) {
	ctint := *tint.cptr()
	corigin := *origin.cptr()
	cdestRec := *destRec.cptr()
	csourceRec := *sourceRec.cptr()
	ctexture := *texture.cptr()
	C.DrawTexturePro(ctexture, csourceRec, cdestRec, corigin, C.float(rotation), ctint)
}

//DrawTextureNPatch : Draws a texture (or part of it) that stretches or shrinks nicely
func DrawTextureNPatch(texture Texture2D, nPatchInfo NPatchInfo, destRec Rectangle, origin Vector2, rotation float32, tint Color) {
	ctint := *tint.cptr()
	corigin := *origin.cptr()
	cdestRec := *destRec.cptr()
	cnPatchInfo := *nPatchInfo.cptr()
	ctexture := *texture.cptr()
	C.DrawTextureNPatch(ctexture, cnPatchInfo, cdestRec, corigin, C.float(rotation), ctint)
}
