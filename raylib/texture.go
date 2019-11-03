package raylib

//#include "raylib.h"
//#include <stdlib.h>
import "C"
import (
	"image"
	"unsafe"
)

//Texture2D stores GPU based textures.
type Texture2D struct {
	Id      uint32
	Width   int32
	Height  int32
	Mipmaps int32
	Format  int32
}

func (t *Texture2D) cptr() *C.Texture2D {
	return (*C.Texture2D)(unsafe.Pointer(t))
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
func LoadTextureFromImage(image *Image) Texture2D {
	i := *image.cptr()
	ct2d := C.LoadTextureFromImage(i)
	return newTextureFromPointer(unsafe.Pointer(&ct2d))
}

//LoadTextureFromGoImage loads image data from image.Image. Uses NewImageFromGoImage.
func LoadTextureFromGoImage(image image.Image) Texture2D {
	return LoadTextureFromImage(NewImageFromGoImage(image))
}

/*
RLAPI TextureCubemap LoadTextureCubemap(Image image, int layoutType);                                    // Load cubemap from image, multiple image cubemap layouts supported
RLAPI RenderTexture2D LoadRenderTexture(int width, int height);                                          // Load texture for rendering (framebuffer)
RLAPI TextureCubemap LoadTextureCubemap(Image image, int layoutType);                                    // Load cubemap from image, multiple image cubemap layouts supported
RLAPI RenderTexture2D LoadRenderTexture(int width, int height);                                          // Load texture for rendering (framebuffer)
RLAPI Image GetTextureData(Texture2D texture);
RLAPI void UpdateTexture(Texture2D texture, const void *pixels);

// Texture2D configuration functions
RLAPI void GenTextureMipmaps(Texture2D *texture);                                                        // Generate GPU mipmaps for a texture
RLAPI void SetTextureFilter(Texture2D texture, int filterMode);                                          // Set texture scaling filter mode
RLAPI void SetTextureWrap(Texture2D texture, int wrapMode);                                              // Set texture wrapping mode

// Texture2D drawing functions
RLAPI void DrawTexture(Texture2D texture, int posX, int posY, Color tint);                               // Draw a Texture2D
RLAPI void DrawTextureV(Texture2D texture, Vector2 position, Color tint);                                // Draw a Texture2D with position defined as Vector2
RLAPI void DrawTextureEx(Texture2D texture, Vector2 position, float rotation, float scale, Color tint);  // Draw a Texture2D with extended parameters
RLAPI void DrawTextureRec(Texture2D texture, Rectangle sourceRec, Vector2 position, Color tint);         // Draw a part of a texture defined by a rectangle
RLAPI void DrawTextureQuad(Texture2D texture, Vector2 tiling, Vector2 offset, Rectangle quad, Color tint);  // Draw texture quad with tiling and offset parameters
RLAPI void DrawTexturePro(Texture2D texture, Rectangle sourceRec, Rectangle destRec, Vector2 origin, float rotation, Color tint);       // Draw a part of a texture defined by a rectangle with 'pro' parameters
RLAPI void DrawTextureNPatch(Texture2D texture, NPatchInfo nPatchInfo, Rectangle destRec, Vector2 origin, float rotation, Color tint);  // Draws a texture (or part of it) that stretches or shrinks nicely
*/
