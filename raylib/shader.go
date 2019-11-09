package raylib

//#include "raylib.h"
//#include <stdlib.h>
import "C"
import "unsafe"

const MaxShaderLocations = 32
const MaxMaterialMaps = 12

type Shader struct {
	Id   uint32
	Locs [MaxShaderLocations]int32
}

func newShaderFromPointer(ptr unsafe.Pointer) *Shader {
	return (*Shader)(ptr)
}

func (s *Shader) cptr() *C.Shader {
	return (*C.Shader)(unsafe.Pointer(s))
}

type ShaderUniformDataType int32

const (
	UniformFloat ShaderUniformDataType = iota
	UniformVec2
	UniformVec3
	UniformVec4
	UniformInt
	UniformIVec2
	UniformIVec3
	UniformIVec4
	UniformSampler2D
)

// Shader location point type
type ShaderLocationIndex int32

//https://dencode.com/en/string/camel-case this is Unsucesfully

const (
	LocVertexPosition ShaderLocationIndex = iota
	LocVertexTexcoord01
	LocVertexTexcoord02
	LocVertexNormal
	LocVertexTangent
	LocVertexColor
	LocMatrixMvp
	LocMatrixModel
	LocMatrixView
	LocMatrixProjection
	LocVectorView
	LocColorDiffuse
	LocColorSpecular
	LocColorAmbient
	LocMapAlbedoLocMapDiffuse
	LocMapMetalnessLocMapSpecular
	LocMapNormal
	LocMapRoughness
	LocMapOcclusion
	LocMapEmission
	LocMapHeight
	LocMapCubemap
	LocMapIrradiance
	LocMapPrefilter
	LocMapBrdf
)

// BlendMode type
type BlendMode int32

const (
	BlendAlpha BlendMode = iota
	BlendAdditive
	BlendMultiplied
)

type Material struct {
	Shader Shader
	Maps   [MaxMaterialMaps]MaterialMap

	//Padding. Apparently required according to https://github.com/gen2brain/raylib-go/blob/02424e2e10eab68b875539f0532a3d51516c4c95/raylib/raylib.go
	_ [4]byte

	Params *[]float32
}

func newMaterialFromPointer(ptr unsafe.Pointer) *Material {
	return (*Material)(ptr)
}

func (s *Material) cptr() *C.Material {
	return (*C.Material)(unsafe.Pointer(s))
}

type MaterialMap struct {
	Texture Texture2D
	Color   Color
	Value   float32
}

func newMaterialMapFromPointer(ptr unsafe.Pointer) *MaterialMap {
	return (*MaterialMap)(ptr)
}

func (s *MaterialMap) cptr() *C.MaterialMap {
	return (*C.MaterialMap)(unsafe.Pointer(s))
}

type MaterialMapType int32

const (
	MapAlbedo MaterialMapType = iota
	MapMetalness
	MapNormal
	MapRoughness
	MapOcclusion
	MapEmission
	MapHeight
	MapCubemap
	MapIrradiance
	MapPrefilter
	MapBRDF
)
