package raylib

//LoadImage : Load image from file into CPU memory (RAM)
func LoadImage(fileName string) (Image, string) {
	cfileName := C.CString(fileName)
	defer C.free(unsafe.Pointer(cfileName))
	res := C.LoadImage(cfileName)
	return newImageFromPointer(unsafe.Pointer(&res)), C.GoString(cfileName)
}

// LoadImageEx Load image data from Color array data (RGBA - 32bit)
func LoadImageEx(pixels []Color, width, height int32) *Image {
	cpixels := pixels[0].cptr()
	cwidth := (C.int)(width)
	cheight := (C.int)(height)
	ret := C.LoadImageEx(cpixels, cwidth, cheight) //We call the C function as raylib uses stb_image to load images.
	v := newImageFromPointer(unsafe.Pointer(&ret))
	return v
}

//LoadImagePro loads raw data wtih parameters
func LoadImagePro(pixels []byte, width, height int32, format PixelFormat) *Image {
	data := unsafe.Pointer(&pixels[0])
	res := C.LoadImagePro(data, C.int(width), C.int(height), C.int(format))
	return newImageFromPointer(unsafe.Pointer(&res))
}

//LoadImageRaw : Load image from RAW file data
func LoadImageRaw(fileName string, width int, height int, format int, headerSize int) (Image, string) {
	cfileName := C.CString(fileName)
	defer C.free(unsafe.Pointer(cfileName))
	res := C.LoadImageRaw(cfileName, C.int(width), C.int(height), C.int(format), C.int(headerSize))
	return newImageFromPointer(unsafe.Pointer(&res)), C.GoString(cfileName)
}

//ExportImage : Export image data to file
func ExportImage(image Image, fileName string) string {
	cfileName := C.CString(fileName)
	defer C.free(unsafe.Pointer(cfileName))
	cimage := *image.cptr()
	C.ExportImage(cimage, cfileName)
	return C.GoString(cfileName)
}

//ExportImageAsCode : Export image as code file defining an array of bytes
func ExportImageAsCode(image Image, fileName string) string {
	cfileName := C.CString(fileName)
	defer C.free(unsafe.Pointer(cfileName))
	cimage := *image.cptr()
	C.ExportImageAsCode(cimage, cfileName)
	return C.GoString(cfileName)
}

//LoadTexture : Load texture from file into GPU memory (VRAM)
func LoadTexture(fileName string) (Texture2D, string) {
	cfileName := C.CString(fileName)
	defer C.free(unsafe.Pointer(cfileName))
	res := C.LoadTexture(cfileName)
	return newTexture2DFromPointer(unsafe.Pointer(&res)), C.GoString(cfileName)
}

//LoadTextureFromImage : Load texture from image data
func LoadTextureFromImage(image Image) Texture2D {
	cimage := *image.cptr()
	res := C.LoadTextureFromImage(cimage)
	return newTexture2DFromPointer(unsafe.Pointer(&res))
}

//LoadTextureCubemap : Load cubemap from image, multiple image cubemap layouts supported
func LoadTextureCubemap(image Image, layoutType int) TextureCubemap {
	cimage := *image.cptr()
	res := C.LoadTextureCubemap(cimage, C.int(layoutType))
	return newTextureCubemapFromPointer(unsafe.Pointer(&res))
}

//LoadRenderTexture : Load texture for rendering (framebuffer)
func LoadRenderTexture(width int, height int) RenderTexture2D {
	res := C.LoadRenderTexture(C.int(width), C.int(height))
	return newRenderTexture2DFromPointer(unsafe.Pointer(&res))
}

//UnloadImage : Unload image from CPU memory (RAM)
func UnloadImage(image Image) {
	cimage := *image.cptr()
	C.UnloadImage(cimage)
}

//UnloadTexture : Unload texture from GPU memory (VRAM)
func UnloadTexture(texture Texture2D) {
	ctexture := *texture.cptr()
	C.UnloadTexture(ctexture)
}

//UnloadRenderTexture : Unload render texture from GPU memory (VRAM)
func UnloadRenderTexture(target RenderTexture2D) {
	ctarget := *target.cptr()
	C.UnloadRenderTexture(ctarget)
}

//GetImageAlphaBorder : Get image alpha border rectangle
func GetImageAlphaBorder(image Image, threshold float32) Rectangle {
	cimage := *image.cptr()
	res := C.GetImageAlphaBorder(cimage, C.float(threshold))
	return newRectangleFromPointer(unsafe.Pointer(&res))
}

//GetPixelDataSize : Get pixel data size in bytes (image or texture)
func GetPixelDataSize(width int, height int, format int) int {
	res := C.GetPixelDataSize(C.int(width), C.int(height), C.int(format))
	return int(res)
}

//GetTextureData : Get pixel data from GPU texture and return an Image
func GetTextureData(texture Texture2D) Image {
	ctexture := *texture.cptr()
	res := C.GetTextureData(ctexture)
	return newImageFromPointer(unsafe.Pointer(&res))
}

//GetScreenData : Get pixel data from screen buffer and return an Image (screenshot)
func GetScreenData() Image {
	res := C.GetScreenData()
	return newImageFromPointer(unsafe.Pointer(&res))
}

//UpdateTexture : Update GPU texture with new data
func UpdateTexture(texture Texture2D, pixels unsafe.Pointer) {
	ctexture := *texture.cptr()
	C.UpdateTexture(ctexture, pixels)
}

//ImageCopy : Create an image duplicate (useful for transformations)
func ImageCopy(image Image) Image {
	cimage := *image.cptr()
	res := C.ImageCopy(cimage)
	return newImageFromPointer(unsafe.Pointer(&res))
}

//ImageFromImage : Create an image from another image piece
func ImageFromImage(image Image, rec Rectangle) Image {
	crec := *rec.cptr()
	cimage := *image.cptr()
	res := C.ImageFromImage(cimage, crec)
	return newImageFromPointer(unsafe.Pointer(&res))
}

//ImageToPOT : Convert image to POT (power-of-two)
func ImageToPOT(image Image, fillColor Color) Image {
	cfillColor := *fillColor.cptr()
	cimage := image.cptr()
	C.ImageToPOT(&cimage, cfillColor)
	return newImageFromPointer(unsafe.Pointer(&cimage))
}

//ImageFormat : Convert image data to desired format
func ImageFormat(image Image, newFormat int) Image {
	cimage := image.cptr()
	C.ImageFormat(&cimage, C.int(newFormat))
	return newImageFromPointer(unsafe.Pointer(&cimage))
}

//ImageAlphaMask : Apply alpha mask to image
func ImageAlphaMask(image Image, alphaMask Image) Image {
	calphaMask := *alphaMask.cptr()
	cimage := image.cptr()
	C.ImageAlphaMask(&cimage, calphaMask)
	return newImageFromPointer(unsafe.Pointer(&cimage))
}

//ImageAlphaClear : Clear alpha channel to desired color
func ImageAlphaClear(image Image, color Color, threshold float32) Image {
	ccolor := *color.cptr()
	cimage := image.cptr()
	C.ImageAlphaClear(&cimage, ccolor, C.float(threshold))
	return newImageFromPointer(unsafe.Pointer(&cimage))
}

//ImageAlphaCrop : Crop image depending on alpha value
func ImageAlphaCrop(image Image, threshold float32) Image {
	cimage := image.cptr()
	C.ImageAlphaCrop(&cimage, C.float(threshold))
	return newImageFromPointer(unsafe.Pointer(&cimage))
}

//ImageAlphaPremultiply : Premultiply alpha channel
func ImageAlphaPremultiply(image Image) Image {
	cimage := image.cptr()
	C.ImageAlphaPremultiply(&cimage)
	return newImageFromPointer(unsafe.Pointer(&cimage))
}

//ImageCrop : Crop an image to a defined rectangle
func ImageCrop(image Image, crop Rectangle) Image {
	ccrop := *crop.cptr()
	cimage := image.cptr()
	C.ImageCrop(&cimage, ccrop)
	return newImageFromPointer(unsafe.Pointer(&cimage))
}

//ImageResize : Resize image (Bicubic scaling algorithm)
func ImageResize(image Image, newWidth int, newHeight int) Image {
	cimage := image.cptr()
	C.ImageResize(&cimage, C.int(newWidth), C.int(newHeight))
	return newImageFromPointer(unsafe.Pointer(&cimage))
}

//ImageResizeNN : Resize image (Nearest-Neighbor scaling algorithm)
func ImageResizeNN(image Image, newWidth int, newHeight int) Image {
	cimage := image.cptr()
	C.ImageResizeNN(&cimage, C.int(newWidth), C.int(newHeight))
	return newImageFromPointer(unsafe.Pointer(&cimage))
}

//ImageResizeCanvas : Resize canvas and fill with color
func ImageResizeCanvas(image Image, newWidth int, newHeight int, offsetX int, offsetY int, color Color) Image {
	ccolor := *color.cptr()
	cimage := image.cptr()
	C.ImageResizeCanvas(&cimage, C.int(newWidth), C.int(newHeight), C.int(offsetX), C.int(offsetY), ccolor)
	return newImageFromPointer(unsafe.Pointer(&cimage))
}

//ImageMipmaps : Generate all mipmap levels for a provided image
func ImageMipmaps(image Image) Image {
	cimage := image.cptr()
	C.ImageMipmaps(&cimage)
	return newImageFromPointer(unsafe.Pointer(&cimage))
}

//ImageDither : Dither image data to 16bpp or lower (Floyd-Steinberg dithering)
func ImageDither(image Image, rBpp int, gBpp int, bBpp int, aBpp int) Image {
	cimage := image.cptr()
	C.ImageDither(&cimage, C.int(rBpp), C.int(gBpp), C.int(bBpp), C.int(aBpp))
	return newImageFromPointer(unsafe.Pointer(&cimage))
}

//ImageText : Create an image from text (default font)
func ImageText(text string, fontSize int, color Color) (Image, string) {
	ccolor := *color.cptr()
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	res := C.ImageText(ctext, C.int(fontSize), ccolor)
	return newImageFromPointer(unsafe.Pointer(&res)), C.GoString(ctext)
}

//ImageTextEx : Create an image from text (custom sprite font)
func ImageTextEx(font Font, text string, fontSize float32, spacing float32, tint Color) (Image, string) {
	ctint := *tint.cptr()
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	cfont := *font.cptr()
	res := C.ImageTextEx(cfont, ctext, C.float(fontSize), C.float(spacing), ctint)
	return newImageFromPointer(unsafe.Pointer(&res)), C.GoString(ctext)
}

//ImageDraw : Draw a source image within a destination image (tint applied to source)
func ImageDraw(dst Image, src Image, srcRec Rectangle, dstRec Rectangle, tint Color) Image {
	ctint := *tint.cptr()
	cdstRec := *dstRec.cptr()
	csrcRec := *srcRec.cptr()
	csrc := *src.cptr()
	cdst := dst.cptr()
	C.ImageDraw(&cdst, csrc, csrcRec, cdstRec, ctint)
	return newImageFromPointer(unsafe.Pointer(&cdst))
}

//ImageDrawRectangle : Draw rectangle within an image
func ImageDrawRectangle(dst Image, rec Rectangle, color Color) Image {
	ccolor := *color.cptr()
	crec := *rec.cptr()
	cdst := dst.cptr()
	C.ImageDrawRectangle(&cdst, crec, ccolor)
	return newImageFromPointer(unsafe.Pointer(&cdst))
}

//ImageDrawRectangleLines : Draw rectangle lines within an image
func ImageDrawRectangleLines(dst Image, rec Rectangle, thick int, color Color) Image {
	ccolor := *color.cptr()
	crec := *rec.cptr()
	cdst := dst.cptr()
	C.ImageDrawRectangleLines(&cdst, crec, C.int(thick), ccolor)
	return newImageFromPointer(unsafe.Pointer(&cdst))
}

//ImageDrawText : Draw text (default font) within an image (destination)
func ImageDrawText(dst Image, position Vector2, text string, fontSize int, color Color) (Image, string) {
	ccolor := *color.cptr()
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	cposition := *position.cptr()
	cdst := dst.cptr()
	C.ImageDrawText(&cdst, cposition, ctext, C.int(fontSize), ccolor)
	return newImageFromPointer(unsafe.Pointer(&cdst)), C.GoString(ctext)
}

//ImageDrawTextEx : Draw text (custom sprite font) within an image (destination)
func ImageDrawTextEx(dst Image, position Vector2, font Font, text string, fontSize float32, spacing float32, color Color) (Image, string) {
	ccolor := *color.cptr()
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	cfont := *font.cptr()
	cposition := *position.cptr()
	cdst := dst.cptr()
	C.ImageDrawTextEx(&cdst, cposition, cfont, ctext, C.float(fontSize), C.float(spacing), ccolor)
	return newImageFromPointer(unsafe.Pointer(&cdst)), C.GoString(ctext)
}

//ImageFlipVertical : Flip image vertically
func ImageFlipVertical(image Image) Image {
	cimage := image.cptr()
	C.ImageFlipVertical(&cimage)
	return newImageFromPointer(unsafe.Pointer(&cimage))
}

//ImageFlipHorizontal : Flip image horizontally
func ImageFlipHorizontal(image Image) Image {
	cimage := image.cptr()
	C.ImageFlipHorizontal(&cimage)
	return newImageFromPointer(unsafe.Pointer(&cimage))
}

//ImageRotateCW : Rotate image clockwise 90deg
func ImageRotateCW(image Image) Image {
	cimage := image.cptr()
	C.ImageRotateCW(&cimage)
	return newImageFromPointer(unsafe.Pointer(&cimage))
}

//ImageRotateCCW : Rotate image counter-clockwise 90deg
func ImageRotateCCW(image Image) Image {
	cimage := image.cptr()
	C.ImageRotateCCW(&cimage)
	return newImageFromPointer(unsafe.Pointer(&cimage))
}

//ImageColorTint : Modify image color: tint
func ImageColorTint(image Image, color Color) Image {
	ccolor := *color.cptr()
	cimage := image.cptr()
	C.ImageColorTint(&cimage, ccolor)
	return newImageFromPointer(unsafe.Pointer(&cimage))
}

//ImageColorInvert : Modify image color: invert
func ImageColorInvert(image Image) Image {
	cimage := image.cptr()
	C.ImageColorInvert(&cimage)
	return newImageFromPointer(unsafe.Pointer(&cimage))
}

//ImageColorGrayscale : Modify image color: grayscale
func ImageColorGrayscale(image Image) Image {
	cimage := image.cptr()
	C.ImageColorGrayscale(&cimage)
	return newImageFromPointer(unsafe.Pointer(&cimage))
}

//ImageColorContrast : Modify image color: contrast (-100 to 100)
func ImageColorContrast(image Image, contrast float32) Image {
	cimage := image.cptr()
	C.ImageColorContrast(&cimage, C.float(contrast))
	return newImageFromPointer(unsafe.Pointer(&cimage))
}

//ImageColorBrightness : Modify image color: brightness (-255 to 255)
func ImageColorBrightness(image Image, brightness int) Image {
	cimage := image.cptr()
	C.ImageColorBrightness(&cimage, C.int(brightness))
	return newImageFromPointer(unsafe.Pointer(&cimage))
}

//ImageColorReplace : Modify image color: replace color
func ImageColorReplace(image Image, color Color, replace Color) Image {
	creplace := *replace.cptr()
	ccolor := *color.cptr()
	cimage := image.cptr()
	C.ImageColorReplace(&cimage, ccolor, creplace)
	return newImageFromPointer(unsafe.Pointer(&cimage))
}

//GenImageColor : Generate image: plain color
func GenImageColor(width int, height int, color Color) Image {
	ccolor := *color.cptr()
	res := C.GenImageColor(C.int(width), C.int(height), ccolor)
	return newImageFromPointer(unsafe.Pointer(&res))
}

//GenImageGradientV : Generate image: vertical gradient
func GenImageGradientV(width int, height int, top Color, bottom Color) Image {
	cbottom := *bottom.cptr()
	ctop := *top.cptr()
	res := C.GenImageGradientV(C.int(width), C.int(height), ctop, cbottom)
	return newImageFromPointer(unsafe.Pointer(&res))
}

//GenImageGradientH : Generate image: horizontal gradient
func GenImageGradientH(width int, height int, left Color, right Color) Image {
	cright := *right.cptr()
	cleft := *left.cptr()
	res := C.GenImageGradientH(C.int(width), C.int(height), cleft, cright)
	return newImageFromPointer(unsafe.Pointer(&res))
}

//GenImageGradientRadial : Generate image: radial gradient
func GenImageGradientRadial(width int, height int, density float32, inner Color, outer Color) Image {
	couter := *outer.cptr()
	cinner := *inner.cptr()
	res := C.GenImageGradientRadial(C.int(width), C.int(height), C.float(density), cinner, couter)
	return newImageFromPointer(unsafe.Pointer(&res))
}

//GenImageChecked : Generate image: checked
func GenImageChecked(width int, height int, checksX int, checksY int, col1 Color, col2 Color) Image {
	ccol2 := *col2.cptr()
	ccol1 := *col1.cptr()
	res := C.GenImageChecked(C.int(width), C.int(height), C.int(checksX), C.int(checksY), ccol1, ccol2)
	return newImageFromPointer(unsafe.Pointer(&res))
}

//GenImageWhiteNoise : Generate image: white noise
func GenImageWhiteNoise(width int, height int, factor float32) Image {
	res := C.GenImageWhiteNoise(C.int(width), C.int(height), C.float(factor))
	return newImageFromPointer(unsafe.Pointer(&res))
}

//GenImagePerlinNoise : Generate image: perlin noise
func GenImagePerlinNoise(width int, height int, offsetX int, offsetY int, scale float32) Image {
	res := C.GenImagePerlinNoise(C.int(width), C.int(height), C.int(offsetX), C.int(offsetY), C.float(scale))
	return newImageFromPointer(unsafe.Pointer(&res))
}

//GenImageCellular : Generate image: cellular algorithm. Bigger tileSize means bigger cells
func GenImageCellular(width int, height int, tileSize int) Image {
	res := C.GenImageCellular(C.int(width), C.int(height), C.int(tileSize))
	return newImageFromPointer(unsafe.Pointer(&res))
}

//GenTextureMipmaps : Generate GPU mipmaps for a texture
func GenTextureMipmaps(texture Texture2D) Texture2D {
	ctexture := texture.cptr()
	C.GenTextureMipmaps(&ctexture)
	return newTexture2DFromPointer(unsafe.Pointer(&ctexture))
}

//SetTextureFilter : Set texture scaling filter mode
func SetTextureFilter(texture Texture2D, filterMode int) {
	ctexture := *texture.cptr()
	C.SetTextureFilter(ctexture, C.int(filterMode))
}

//SetTextureWrap : Set texture wrapping mode
func SetTextureWrap(texture Texture2D, wrapMode int) {
	ctexture := *texture.cptr()
	C.SetTextureWrap(ctexture, C.int(wrapMode))
}

//DrawTexture : Draw a Texture2D
func DrawTexture(texture Texture2D, posX int, posY int, tint Color) {
	ctint := *tint.cptr()
	ctexture := *texture.cptr()
	C.DrawTexture(ctexture, C.int(posX), C.int(posY), ctint)
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
