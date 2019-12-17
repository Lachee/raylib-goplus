package raylib

//#include "raylib.h"
import "C"
import "unsafe"

const (
	MaxMeshVertices         = 1 << 28
	MaxMeshIndices          = 1 << 28
	MaxMeshTexCoords        = 1 << 28
	MaxMeshAnimatedVertices = 1 << 28
	MaxMeshBones            = 1 << 28
)

//Mesh is vertex data that is stored in CPU and GPU memory.
// Note that the values are pointing to C memory
type Mesh struct {
	// Number of vertices stored in arrays
	VertexCount int32

	// Number of triangles stored (indexed or not)
	TriangleCount int32

	// Vertex position (XYZ - 3 components per vertex) (shader-location = 0)
	Vertices *[MaxMeshVertices]Vector3

	// Vertex texture coordinates (UV - 2 components per vertex) (shader-location = 1)
	Texcoords *[MaxMeshTexCoords]Vector2

	// Vertex second texture coordinates (useful for lightmaps) (shader-location = 5)
	Texcoords2 *[MaxMeshTexCoords]Vector2

	// Vertex normals (XYZ - 3 components per vertex) (shader-location = 2)
	Normals *[MaxMeshVertices]Vector3

	// Vertex tangents (XYZ - 3 components per vertex) (shader-location = 4)
	Tangents *[MaxMeshVertices]Vector3

	// Vertex colors (RGBA - 4 components per vertex) (shader-location = 3)
	Colors *[MaxMeshVertices]Color

	// Vertex indices (in case vertex data comes indexed)
	Indices *[MaxMeshIndices]uint16

	AnimVertices *[MaxMeshAnimatedVertices]Vector3
	AnimNormals  *[MaxMeshAnimatedVertices]Vector3
	BoneIds      *[MaxMeshBones]int32
	BoneWeights  *[MaxMeshBones]float32

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

const (
	MaxModelMeshes    = 1 << 28
	MaxModelMaterials = 1 << 28
	MaxModelBones     = 1 << 28
	MaxModelBinds     = 1 << 28
)

type Model struct {
	Transform Matrix
	MeshCount int32
	Meshes    *[MaxModelMeshes]Mesh

	MaterialCount int32
	Materials     *[MaxModelMaterials]Material
	MeshMaterial  *int32

	BoneCount int32
	Bones     *[MaxModelBones]BoneInfo
	BindPos   *[MaxModelBinds]Transform
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
