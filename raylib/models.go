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
	Vertices *[1 << 28]Vector3

	// Vertex texture coordinates (UV - 2 components per vertex) (shader-location = 1)
	Texcoords *[1 << 28]Vector2

	// Vertex second texture coordinates (useful for lightmaps) (shader-location = 5)
	Texcoords2 *[1 << 28]Vector2

	// Vertex normals (XYZ - 3 components per vertex) (shader-location = 2)
	Normals *[1 << 28]Vector3

	// Vertex tangents (XYZ - 3 components per vertex) (shader-location = 4)
	Tangents *[1 << 28]Vector3

	// Vertex colors (RGBA - 4 components per vertex) (shader-location = 3)
	Colors *[1 << 28]Color

	// Vertex indices (in case vertex data comes indexed)
	Indices *[1 << 2]uint16

	AnimVertices *[1 << 28]Vector3
	AnimNormals  *[1 << 28]Vector3
	BoneIds      *[1 << 2]int32
	BoneWeights  *[1 << 2]float32

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
	Transform Matrix
	MeshCount int32
	Meshes    *[1 << 28]Mesh

	MaterialCount int32
	Materials     *[1 << 8]Material
	MeshMaterial  *int32

	BoneCount int32
	Bones     *[1 << 8]BoneInfo
	BindPos   *[1 << 8]Transform
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
