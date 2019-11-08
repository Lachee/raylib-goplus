package raylib

//#include "raylib.h"
//#include <stdlib.h>
import "C"
import "unsafe"

const MaxShaderLocations = 32

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
