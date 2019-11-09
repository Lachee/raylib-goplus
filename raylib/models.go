package raylib

//#include "raylib.h"
import "C"
import "unsafe"

//Mesh is vertex data that is stored in CPU and GPU memory.
// Note that the values are pointing to C memory
type Mesh struct {
	// Number of vertices stored in arrays
	VertexCount int32
	// Number of triangles stored (indexed or not)
	TriangleCount int32
	// Vertex position (XYZ - 3 components per vertex) (shader-location = 0)
	Vertices *[]Vector3
	// Vertex texture coordinates (UV - 2 components per vertex) (shader-location = 1)
	Texcoords *[]Vector2
	// Vertex second texture coordinates (useful for lightmaps) (shader-location = 5)
	Texcoords2 *[]Vector2
	// Vertex normals (XYZ - 3 components per vertex) (shader-location = 2)
	Normals *[]Vector4
	// Vertex tangents (XYZ - 3 components per vertex) (shader-location = 4)
	Tangents *[]float32
	// Vertex colors (RGBA - 4 components per vertex) (shader-location = 3)
	Colors *[]Color

	AnimVertices *[]Vector3
	AnimNormals  *[]Vector4
	BoneIds      *[]int32
	BoneWeights  *[]float32

	// Vertex indices (in case vertex data comes indexed)
	Indices *[]uint16

	// OpenGL Vertex Array Object id
	VaoID uint32

	// OpenGL Vertex Buffer Objects id (7 types of vertex data)
	VboID unsafe.Pointer
}

func newMeshFromPointer(ptr unsafe.Pointer) *Mesh {
	return (*Mesh)(ptr)
}

func (s *Mesh) cptr() *C.Mesh {
	return (*C.Mesh)(unsafe.Pointer(s))
}

type BoneInfo struct {
	Name   [32]byte
	Parent int32
}

type Model struct {
	Transform     Matrix
	MeshCount     int32
	Meshes        *[]Mesh
	MaterialCount int32
	Materials     *[]Material
	BoneCount     int32
	Bones         *[]BoneInfo
	BindPos       *[]Transform
}

type ModelAnimation struct {
	BoneCount  int32
	Bones      *[]BoneInfo
	FrameCount int32
	FramePoses *[](*[]Transform)
}

func newModelFromPointer(ptr unsafe.Pointer) *Model {
	return (*Model)(ptr)
}

func (s *Model) cptr() *C.Model {
	return (*C.Model)(unsafe.Pointer(s))
}

func newModelAnimationFromPointer(ptr unsafe.Pointer) *ModelAnimation {
	return (*ModelAnimation)(ptr)
}

func (s *ModelAnimation) cptr() *C.ModelAnimation {
	return (*C.ModelAnimation)(unsafe.Pointer(s))
}
