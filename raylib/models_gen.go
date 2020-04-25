package raylib

/*
//Generated 2020-04-25T14:27:29+10:00
#include "raylib.h"
#include <stdlib.h>
#include "go.h"
*/
import "C"
import "unsafe"

//LoadModel : Load model from files (meshes and materials)
func LoadModel(fileName string) *Model {
	cfileName := C.CString(fileName)
	defer C.free(unsafe.Pointer(cfileName))
	res := C.LoadModel(cfileName)
	retval := newModelFromPointer(unsafe.Pointer(&res))
	RegisterUnloadable(retval)
	return retval
}

//LoadModelFromMesh : Load model from generated mesh (default material)
func LoadModelFromMesh(mesh *Mesh) *Model {
	cmesh := *mesh.cptr()
	res := C.LoadModelFromMesh(cmesh)
	retval := newModelFromPointer(unsafe.Pointer(&res))
	RegisterUnloadable(retval)
	return retval
}

//Unload : Unload model from memory (RAM and/or VRAM)
func (model *Model) Unload() {
	cmodel := *model.cptr()
	C.UnloadModel(cmodel)
	UnregisterUnloadable(model)
}

//UnloadModel : Unload model from memory (RAM and/or VRAM)
//Recommended to use model.Unload() instead
func UnloadModel(model *Model) {
	model.Unload()
}

//Export : Export mesh data to file
func (mesh *Mesh) Export(fileName string) {
	cfileName := C.CString(fileName)
	defer C.free(unsafe.Pointer(cfileName))
	cmesh := *mesh.cptr()
	C.ExportMesh(cmesh, cfileName)
}

//ExportMesh : Export mesh data to file
//Recommended to use mesh.Export(fileName) instead
func ExportMesh(mesh *Mesh, fileName string) {
	mesh.Export(fileName)
}

//Unload : Unload mesh from memory (RAM and/or VRAM)
func (mesh *Mesh) Unload() {
	cmesh := *mesh.cptr()
	C.UnloadMesh(cmesh)
	UnregisterUnloadable(mesh)
}

//UnloadMesh : Unload mesh from memory (RAM and/or VRAM)
//Recommended to use mesh.Unload() instead
func UnloadMesh(mesh *Mesh) {
	mesh.Unload()
}

//LoadMaterialDefault : Load default material (Supports: DIFFUSE, SPECULAR, NORMAL maps)
func LoadMaterialDefault() *Material {
	res := C.LoadMaterialDefault()
	retval := newMaterialFromPointer(unsafe.Pointer(&res))
	RegisterUnloadable(retval)
	return retval
}

//Unload : Unload material from GPU memory (VRAM)
func (material *Material) Unload() {
	cmaterial := *material.cptr()
	C.UnloadMaterial(cmaterial)
	UnregisterUnloadable(material)
}

//UnloadMaterial : Unload material from GPU memory (VRAM)
//Recommended to use material.Unload() instead
func UnloadMaterial(material *Material) {
	material.Unload()
}

//SetTexture : Set texture for a material map type (MAP_DIFFUSE, MAP_SPECULAR...)
func (material *Material) SetTexture(mapType MaterialMapType, texture Texture2D) {
	ctexture := *texture.cptr()
	cmaterial := material.cptr()
	C.SetMaterialTexture(cmaterial, C.int(int32(mapType)), ctexture)
	material.Maps[int(mapType)].Texture = texture
}

//SetMaterialTexture : Set texture for a material map type (MAP_DIFFUSE, MAP_SPECULAR...)
//Recommended to use material.SetTexture(mapType, texture) instead
func SetMaterialTexture(material *Material, mapType MaterialMapType, texture Texture2D) {
	material.SetTexture(mapType, texture)
}

//SetMeshMaterial : Set material for a mesh
func (model *Model) SetMeshMaterial(meshId int, materialId int) {
	cmodel := model.cptr()
	C.SetModelMeshMaterial(cmodel, C.int(int32(meshId)), C.int(int32(materialId)))
}

//SetModelMeshMaterial : Set material for a mesh
//Recommended to use model.SetMeshMaterial(meshId, materialId) instead
func SetModelMeshMaterial(model *Model, meshId int, materialId int) {
	model.SetMeshMaterial(meshId, materialId)
}

//LoadModelAnimations : Load model animations from file
func LoadModelAnimations(fileName string) ([]ModelAnimation, int32) {
	cfileName := C.CString(fileName)
	ccount := C.int(0)
	defer C.free(unsafe.Pointer(cfileName))

	res := C.LoadModelAnimations(cfileName, &ccount)

	samples := int32(ccount)
	tmpslice := (*[1 << 24]*C.ModelAnimation)(unsafe.Pointer(res))[:samples:samples]

	goslice := make([]ModelAnimation, samples)
	for i, s := range tmpslice {
		goslice[i] = *newModelAnimationFromPointer(unsafe.Pointer(s))
	}

	return goslice, samples
}

//UpdateAnimation : Update model animation pose
func (model *Model) UpdateAnimation(anim *ModelAnimation, frame int) {
	canim := *anim.cptr()
	cmodel := *model.cptr()
	C.UpdateModelAnimation(cmodel, canim, C.int(int32(frame)))
}

//UpdateModelAnimation : Update model animation pose
//Recommended to use model.UpdateAnimation(anim, frame) instead
func UpdateModelAnimation(model *Model, anim *ModelAnimation, frame int) {
	model.UpdateAnimation(anim, frame)
}

//Unload : Unload animation data
func (anim *ModelAnimation) Unload() {
	canim := *anim.cptr()
	C.UnloadModelAnimation(canim)
	UnregisterUnloadable(anim)
}

//UnloadModelAnimation : Unload animation data
//Recommended to use anim.Unload() instead
func UnloadModelAnimation(anim *ModelAnimation) {
	anim.Unload()
}

//IsAnimationValid : Check model animation skeleton match
func (model *Model) IsAnimationValid(anim *ModelAnimation) bool {
	canim := *anim.cptr()
	cmodel := *model.cptr()
	res := C.IsModelAnimationValid(cmodel, canim)
	return bool(res)
}

//IsModelAnimationValid : Check model animation skeleton match
//Recommended to use model.IsAnimationValid(anim) instead
func IsModelAnimationValid(model *Model, anim *ModelAnimation) bool {
	return model.IsAnimationValid(anim)
}

//GenMeshPoly : Generate polygonal mesh
func GenMeshPoly(sides int, radius float32) *Mesh {
	res := C.GenMeshPoly(C.int(int32(sides)), C.float(radius))
	retval := newMeshFromPointer(unsafe.Pointer(&res))
	RegisterUnloadable(retval)
	return retval
}

//GenMeshPlane : Generate plane mesh (with subdivisions)
func GenMeshPlane(width float32, length float32, resX int, resZ int) *Mesh {
	res := C.GenMeshPlane(C.float(width), C.float(length), C.int(int32(resX)), C.int(int32(resZ)))
	retval := newMeshFromPointer(unsafe.Pointer(&res))
	RegisterUnloadable(retval)
	return retval
}

//GenMeshCube : Generate cuboid mesh
func GenMeshCube(width float32, height float32, length float32) *Mesh {
	res := C.GenMeshCube(C.float(width), C.float(height), C.float(length))
	retval := newMeshFromPointer(unsafe.Pointer(&res))
	RegisterUnloadable(retval)
	return retval
}

//GenMeshSphere : Generate sphere mesh (standard sphere)
func GenMeshSphere(radius float32, rings int, slices int) *Mesh {
	res := C.GenMeshSphere(C.float(radius), C.int(int32(rings)), C.int(int32(slices)))
	retval := newMeshFromPointer(unsafe.Pointer(&res))
	RegisterUnloadable(retval)
	return retval
}

//GenMeshHemiSphere : Generate half-sphere mesh (no bottom cap)
func GenMeshHemiSphere(radius float32, rings int, slices int) *Mesh {
	res := C.GenMeshHemiSphere(C.float(radius), C.int(int32(rings)), C.int(int32(slices)))
	retval := newMeshFromPointer(unsafe.Pointer(&res))
	RegisterUnloadable(retval)
	return retval
}

//GenMeshCylinder : Generate cylinder mesh
func GenMeshCylinder(radius float32, height float32, slices int) *Mesh {
	res := C.GenMeshCylinder(C.float(radius), C.float(height), C.int(int32(slices)))
	retval := newMeshFromPointer(unsafe.Pointer(&res))
	RegisterUnloadable(retval)
	return retval
}

//GenMeshTorus : Generate torus mesh
func GenMeshTorus(radius float32, size float32, radSeg int, sides int) *Mesh {
	res := C.GenMeshTorus(C.float(radius), C.float(size), C.int(int32(radSeg)), C.int(int32(sides)))
	retval := newMeshFromPointer(unsafe.Pointer(&res))
	RegisterUnloadable(retval)
	return retval
}

//GenMeshKnot : Generate trefoil knot mesh
func GenMeshKnot(radius float32, size float32, radSeg int, sides int) *Mesh {
	res := C.GenMeshKnot(C.float(radius), C.float(size), C.int(int32(radSeg)), C.int(int32(sides)))
	retval := newMeshFromPointer(unsafe.Pointer(&res))
	RegisterUnloadable(retval)
	return retval
}

//GenMeshHeightmap : Generate heightmap mesh from image data
func (heightamap *Image) GenMeshHeightmap(size Vector3) *Mesh {
	csize := *size.cptr()
	cheightamap := *heightamap.cptr()
	res := C.GenMeshHeightmap(cheightamap, csize)
	retval := newMeshFromPointer(unsafe.Pointer(&res))
	RegisterUnloadable(retval)
	return retval
}

//GenMeshHeightmap : Generate heightmap mesh from image data
//Recommended to use heightamap.GenMeshHeightmap(size) instead
func GenMeshHeightmap(heightamap *Image, size Vector3) *Mesh {
	return heightamap.GenMeshHeightmap(size)
}

//GenMeshCubicmap : Generate cubes-based map mesh from image data
func (cubicamap *Image) GenMeshCubicmap(cubeSize Vector3) *Mesh {
	ccubeSize := *cubeSize.cptr()
	ccubicamap := *cubicamap.cptr()
	res := C.GenMeshCubicmap(ccubicamap, ccubeSize)
	retval := newMeshFromPointer(unsafe.Pointer(&res))
	RegisterUnloadable(retval)
	return retval
}

//GenMeshCubicmap : Generate cubes-based map mesh from image data
//Recommended to use cubicamap.GenMeshCubicmap(cubeSize) instead
func GenMeshCubicmap(cubicamap *Image, cubeSize Vector3) *Mesh {
	return cubicamap.GenMeshCubicmap(cubeSize)
}

//BoundingBox : Compute mesh bounding box limits
func (mesh *Mesh) BoundingBox() BoundingBox {
	cmesh := *mesh.cptr()
	res := C.MeshBoundingBox(cmesh)
	return newBoundingBoxFromPointer(unsafe.Pointer(&res))
}

//MeshBoundingBox : Compute mesh bounding box limits
//Recommended to use mesh.BoundingBox() instead
func MeshBoundingBox(mesh *Mesh) BoundingBox {
	return mesh.BoundingBox()
}

//ComputeTangents : Compute mesh tangents
func (mesh *Mesh) ComputeTangents() {
	cmesh := mesh.cptr()
	C.MeshTangents(cmesh)
}

//MeshTangents : Compute mesh tangents
//Recommended to use mesh.ComputeTangents() instead
func MeshTangents(mesh *Mesh) {
	mesh.ComputeTangents()
}

//ComputeBinormals : Compute mesh binormals
func (mesh *Mesh) ComputeBinormals() {
	cmesh := mesh.cptr()
	C.MeshBinormals(cmesh)
}

//MeshBinormals : Compute mesh binormals
//Recommended to use mesh.ComputeBinormals() instead
func MeshBinormals(mesh *Mesh) {
	mesh.ComputeBinormals()
}

//DrawModel : Draw a model (with texture if set)
func DrawModel(model Model, position Vector3, scale float32, tint Color) {
	ctint := *tint.cptr()
	cposition := *position.cptr()
	cmodel := *model.cptr()
	C.DrawModel(cmodel, cposition, C.float(scale), ctint)
}

//DrawModelEx : Draw a model with extended parameters
func DrawModelEx(model Model, position Vector3, rotationAxis Vector3, rotationAngle float32, scale Vector3, tint Color) {
	ctint := *tint.cptr()
	cscale := *scale.cptr()
	crotationAxis := *rotationAxis.cptr()
	cposition := *position.cptr()
	cmodel := *model.cptr()
	C.DrawModelEx(cmodel, cposition, crotationAxis, C.float(rotationAngle), cscale, ctint)
}

//DrawModelWires : Draw a model wires (with texture if set)
func DrawModelWires(model Model, position Vector3, scale float32, tint Color) {
	ctint := *tint.cptr()
	cposition := *position.cptr()
	cmodel := *model.cptr()
	C.DrawModelWires(cmodel, cposition, C.float(scale), ctint)
}

//DrawModelWiresEx : Draw a model wires (with texture if set) with extended parameters
func DrawModelWiresEx(model Model, position Vector3, rotationAxis Vector3, rotationAngle float32, scale Vector3, tint Color) {
	ctint := *tint.cptr()
	cscale := *scale.cptr()
	crotationAxis := *rotationAxis.cptr()
	cposition := *position.cptr()
	cmodel := *model.cptr()
	C.DrawModelWiresEx(cmodel, cposition, crotationAxis, C.float(rotationAngle), cscale, ctint)
}

//DrawBoundingBox : Draw bounding box (wires)
func DrawBoundingBox(box BoundingBox, color Color) {
	ccolor := *color.cptr()
	cbox := *box.cptr()
	C.DrawBoundingBox(cbox, ccolor)
}

//DrawBillboard : Draw a billboard texture
func DrawBillboard(camera Camera, texture Texture2D, center Vector3, size float32, tint Color) {
	ctint := *tint.cptr()
	ccenter := *center.cptr()
	ctexture := *texture.cptr()
	ccamera := *camera.cptr()
	C.DrawBillboard(ccamera, ctexture, ccenter, C.float(size), ctint)
}

//DrawBillboardRec : Draw a billboard texture defined by sourceRec
func DrawBillboardRec(camera Camera, texture Texture2D, sourceRec Rectangle, center Vector3, size float32, tint Color) {
	ctint := *tint.cptr()
	ccenter := *center.cptr()
	csourceRec := *sourceRec.cptr()
	ctexture := *texture.cptr()
	ccamera := *camera.cptr()
	C.DrawBillboardRec(ccamera, ctexture, csourceRec, ccenter, C.float(size), ctint)
}

//CheckCollisionSpheres : Detect collision between two spheres
func CheckCollisionSpheres(centerA Vector3, radiusA float32, centerB Vector3, radiusB float32) bool {
	ccenterB := *centerB.cptr()
	ccenterA := *centerA.cptr()
	res := C.CheckCollisionSpheres(ccenterA, C.float(radiusA), ccenterB, C.float(radiusB))
	return bool(res)
}

//CheckCollisionBoxes : Detect collision between two bounding boxes
func CheckCollisionBoxes(box1 BoundingBox, box2 BoundingBox) bool {
	cbox2 := *box2.cptr()
	cbox1 := *box1.cptr()
	res := C.CheckCollisionBoxes(cbox1, cbox2)
	return bool(res)
}

//CheckCollisionBoxSphere : Detect collision between box and sphere
func CheckCollisionBoxSphere(box BoundingBox, center Vector3, radius float32) bool {
	ccenter := *center.cptr()
	cbox := *box.cptr()
	res := C.CheckCollisionBoxSphere(cbox, ccenter, C.float(radius))
	return bool(res)
}

//CheckCollisionRaySphere : Detect collision between ray and sphere
func CheckCollisionRaySphere(ray Ray, center Vector3, radius float32) bool {
	ccenter := *center.cptr()
	cray := *ray.cptr()
	res := C.CheckCollisionRaySphere(cray, ccenter, C.float(radius))
	return bool(res)
}

//CheckCollisionRaySphereEx : Detect collision between ray and sphere, returns collision point
func CheckCollisionRaySphereEx(ray Ray, center Vector3, radius float32, collisionPoint Vector3) (bool, Vector3) {
	ccollisionPoint := collisionPoint.cptr()
	ccenter := *center.cptr()
	cray := *ray.cptr()
	res := C.CheckCollisionRaySphereEx(cray, ccenter, C.float(radius), ccollisionPoint)
	return bool(res), newVector3FromPointer(unsafe.Pointer(ccollisionPoint))
}

//CheckCollisionRayBox : Detect collision between ray and box
func CheckCollisionRayBox(ray Ray, box BoundingBox) bool {
	cbox := *box.cptr()
	cray := *ray.cptr()
	res := C.CheckCollisionRayBox(cray, cbox)
	return bool(res)
}

//GetCollisionRayModel : Get collision info between ray and model
func GetCollisionRayModel(ray Ray, model Model) RayHitInfo {
	cmodel := *model.cptr()
	cray := *ray.cptr()
	res := C.GetCollisionRayModel(cray, cmodel)
	return newRayHitInfoFromPointer(unsafe.Pointer(&res))
}

//GetCollisionRayTriangle : Get collision info between ray and triangle
func GetCollisionRayTriangle(ray Ray, p1 Vector3, p2 Vector3, p3 Vector3) RayHitInfo {
	cp3 := *p3.cptr()
	cp2 := *p2.cptr()
	cp1 := *p1.cptr()
	cray := *ray.cptr()
	res := C.GetCollisionRayTriangle(cray, cp1, cp2, cp3)
	return newRayHitInfoFromPointer(unsafe.Pointer(&res))
}

//GetCollisionRayGround : Get collision info between ray and ground plane (Y-normal plane)
func GetCollisionRayGround(ray Ray, groundHeight float32) RayHitInfo {
	cray := *ray.cptr()
	res := C.GetCollisionRayGround(cray, C.float(groundHeight))
	return newRayHitInfoFromPointer(unsafe.Pointer(&res))
}
