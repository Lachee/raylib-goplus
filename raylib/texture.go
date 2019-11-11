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

func newTexture2DFromPointer(ptr unsafe.Pointer) *Texture2D {
	return (*Texture2D)(ptr)
}

//LoadTextureFromGoImage loads image data from image.Image. Uses NewImageFromGoImage.
func LoadTextureFromGoImage(image image.Image) *Texture2D {
	img := NewImageFromGoImage(image)
	defer img.Unload()
	return LoadTextureFromImage(img)
}

//TextureCubemap type, actuall the same as a Texture2D
type TextureCubemap Texture2D
type CubemapLayoutType int32

const (
	//CubmapAutoDetect automatically detect layout type
	CubmapAutoDetect CubemapLayoutType = iota
	//CubemapLineVertical Layout is defined by a vertical line with faces
	CubemapLineVertical
	//CubemapLineHorizontal Layout is defined by an horizontal line with faces
	CubemapLineHorizontal
	//CubemapCrossThreeByFour Layout is defined by a 3x4 cross with cubemap faces
	CubemapCrossThreeByFour
	//CubemapCrossFourByThree Layout is defined by a 4x3 cross with cubemap faces
	CubemapCrossFourByThree
	//CubemapPanorama Layout is defined by a panorama image (equirectangular map)
	CubemapPanorama
)

func (t *TextureCubemap) cptr() *C.TextureCubemap {
	return (*C.TextureCubemap)(unsafe.Pointer(t))
}

func newTextureCubemapFromPointer(ptr unsafe.Pointer) *TextureCubemap {
	return (*TextureCubemap)(ptr)
}

//Unload : Unload texture from GPU memory (VRAM)
func (texture *TextureCubemap) Unload() {
	ctexture := C.Texture2D(*texture.cptr())
	C.UnloadTexture(ctexture)
	removeUnloadable(texture)
}

//RenderTexture2D is a texture used for rendering
type RenderTexture2D struct {
	Id           uint32
	Texture      Texture2D
	Depth        Texture2D
	DepthTexture bool
}

func newRenderTexture2DFromPointer(ptr unsafe.Pointer) *RenderTexture2D {
	return (*RenderTexture2D)(ptr)
}

func (t *RenderTexture2D) cptr() *C.RenderTexture2D {
	return (*C.RenderTexture2D)(unsafe.Pointer(t))
}

type NPatchType int32

const (
	NPT9Patch NPatchType = iota
	NPT3PatchVertical
	NPT3PatchHorizontal
)

//NPatchInfo is the N-Patch layout information
type NPatchInfo struct {
	SourceRectangle Rectangle
	Left            int32
	Top             int32
	Right           int32
	Bottom          int32
	Type            NPatchType
}

func newNPatchInfoFromPointer(ptr unsafe.Pointer) NPatchInfo {
	return *(*NPatchInfo)(ptr)
}

func (t *NPatchInfo) cptr() *C.NPatchInfo {
	return (*C.NPatchInfo)(unsafe.Pointer(t))
}

type TextureWrapMode int32

const (
	WrapRepeat TextureWrapMode = iota
	WrapClamp
	WrapMirrorRepeat
	WrapMirrorClamp
)

type TextureFilterMode int32

const (
	FilterPoint TextureFilterMode = iota
	FilterBilinear
	FilterTrilinear
	FilterAnisotropic4x
	FilterAnisotropic8x
	FilterAnisotropic16x
)
