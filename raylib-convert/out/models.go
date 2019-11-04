package raylib

//LoadModel : Load model from files (meshes and materials)
func LoadModel(fileName string) (Model, string) {
	cfileName := C.CString(fileName)
	defer C.free(unsafe.Pointer(cfileName))
	res := C.LoadModel(cfileName)
	return newModelFromPointer(unsafe.Pointer(&res)), C.GoString(cfileName)
}

//LoadModelFromMesh : Load model from generated mesh (default material)
func LoadModelFromMesh(mesh Mesh) Model {
	cmesh := *mesh.cptr()
	res := C.LoadModelFromMesh(cmesh)
	return newModelFromPointer(unsafe.Pointer(&res))
}

//UnloadModel : Unload model from memory (RAM and/or VRAM)
func UnloadModel(model Model) {
	cmodel := *model.cptr()
	C.UnloadModel(cmodel)
}

//ExportMesh : Export mesh data to file
func ExportMesh(mesh Mesh, fileName string) string {
	cfileName := C.CString(fileName)
	defer C.free(unsafe.Pointer(cfileName))
	cmesh := *mesh.cptr()
	C.ExportMesh(cmesh, cfileName)
	return C.GoString(cfileName)
}

//UnloadMesh : Unload mesh from memory (RAM and/or VRAM)
func UnloadMesh(mesh Mesh) {
	cmesh := *mesh.cptr()
	C.UnloadMesh(cmesh)
}

//LoadMaterialDefault : Load default material (Supports: DIFFUSE, SPECULAR, NORMAL maps)
func LoadMaterialDefault() Material {
	res := C.LoadMaterialDefault()
	return newMaterialFromPointer(unsafe.Pointer(&res))
}

//UnloadMaterial : Unload material from GPU memory (VRAM)
func UnloadMaterial(material Material) {
	cmaterial := *material.cptr()
	C.UnloadMaterial(cmaterial)
}

//SetMaterialTexture : Set texture for a material map type (MAP_DIFFUSE, MAP_SPECULAR...)
func SetMaterialTexture(material Material, mapType int, texture Texture2D) Material {
	ctexture := *texture.cptr()
	cmaterial := material.cptr()
	C.SetMaterialTexture(&cmaterial, C.int(mapType), ctexture)
	return newMaterialFromPointer(unsafe.Pointer(&cmaterial))
}

//SetModelMeshMaterial : Set material for a mesh
func SetModelMeshMaterial(model Model, meshId int, materialId int) Model {
	cmodel := model.cptr()
	C.SetModelMeshMaterial(&cmodel, C.int(meshId), C.int(materialId))
	return newModelFromPointer(unsafe.Pointer(&cmodel))
}

//UpdateModelAnimation : Update model animation pose
func UpdateModelAnimation(model Model, anim ModelAnimation, frame int) {
	canim := *anim.cptr()
	cmodel := *model.cptr()
	C.UpdateModelAnimation(cmodel, canim, C.int(frame))
}

//UnloadModelAnimation : Unload animation data
func UnloadModelAnimation(anim ModelAnimation) {
	canim := *anim.cptr()
	C.UnloadModelAnimation(canim)
}

//IsModelAnimationValid : Check model animation skeleton match
func IsModelAnimationValid(model Model, anim ModelAnimation) bool {
	canim := *anim.cptr()
	cmodel := *model.cptr()
	res := C.IsModelAnimationValid(cmodel, canim)
	return bool(res)
}

//GenMeshPoly : Generate polygonal mesh
func GenMeshPoly(sides int, radius float32) Mesh {
	res := C.GenMeshPoly(C.int(sides), C.float(radius))
	return newMeshFromPointer(unsafe.Pointer(&res))
}

//GenMeshPlane : Generate plane mesh (with subdivisions)
func GenMeshPlane(width float32, length float32, resX int, resZ int) Mesh {
	res := C.GenMeshPlane(C.float(width), C.float(length), C.int(resX), C.int(resZ))
	return newMeshFromPointer(unsafe.Pointer(&res))
}

//GenMeshCube : Generate cuboid mesh
func GenMeshCube(width float32, height float32, length float32) Mesh {
	res := C.GenMeshCube(C.float(width), C.float(height), C.float(length))
	return newMeshFromPointer(unsafe.Pointer(&res))
}

//GenMeshSphere : Generate sphere mesh (standard sphere)
func GenMeshSphere(radius float32, rings int, slices int) Mesh {
	res := C.GenMeshSphere(C.float(radius), C.int(rings), C.int(slices))
	return newMeshFromPointer(unsafe.Pointer(&res))
}

//GenMeshHemiSphere : Generate half-sphere mesh (no bottom cap)
func GenMeshHemiSphere(radius float32, rings int, slices int) Mesh {
	res := C.GenMeshHemiSphere(C.float(radius), C.int(rings), C.int(slices))
	return newMeshFromPointer(unsafe.Pointer(&res))
}

//GenMeshCylinder : Generate cylinder mesh
func GenMeshCylinder(radius float32, height float32, slices int) Mesh {
	res := C.GenMeshCylinder(C.float(radius), C.float(height), C.int(slices))
	return newMeshFromPointer(unsafe.Pointer(&res))
}

//GenMeshTorus : Generate torus mesh
func GenMeshTorus(radius float32, size float32, radSeg int, sides int) Mesh {
	res := C.GenMeshTorus(C.float(radius), C.float(size), C.int(radSeg), C.int(sides))
	return newMeshFromPointer(unsafe.Pointer(&res))
}

//GenMeshKnot : Generate trefoil knot mesh
func GenMeshKnot(radius float32, size float32, radSeg int, sides int) Mesh {
	res := C.GenMeshKnot(C.float(radius), C.float(size), C.int(radSeg), C.int(sides))
	return newMeshFromPointer(unsafe.Pointer(&res))
}

//GenMeshHeightmap : Generate heightmap mesh from image data
func GenMeshHeightmap(heightmap Image, size Vector3) Mesh {
	csize := *size.cptr()
	cheightmap := *heightmap.cptr()
	res := C.GenMeshHeightmap(cheightmap, csize)
	return newMeshFromPointer(unsafe.Pointer(&res))
}

//GenMeshCubicmap : Generate cubes-based map mesh from image data
func GenMeshCubicmap(cubicmap Image, cubeSize Vector3) Mesh {
	ccubeSize := *cubeSize.cptr()
	ccubicmap := *cubicmap.cptr()
	res := C.GenMeshCubicmap(ccubicmap, ccubeSize)
	return newMeshFromPointer(unsafe.Pointer(&res))
}

//MeshBoundingBox : Compute mesh bounding box limits
func MeshBoundingBox(mesh Mesh) BoundingBox {
	cmesh := *mesh.cptr()
	res := C.MeshBoundingBox(cmesh)
	return newBoundingBoxFromPointer(unsafe.Pointer(&res))
}

//MeshTangents : Compute mesh tangents
func MeshTangents(mesh Mesh) Mesh {
	cmesh := mesh.cptr()
	C.MeshTangents(&cmesh)
	return newMeshFromPointer(unsafe.Pointer(&cmesh))
}

//MeshBinormals : Compute mesh binormals
func MeshBinormals(mesh Mesh) Mesh {
	cmesh := mesh.cptr()
	C.MeshBinormals(&cmesh)
	return newMeshFromPointer(unsafe.Pointer(&cmesh))
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
	res := C.CheckCollisionRaySphereEx(cray, ccenter, C.float(radius), &ccollisionPoint)
	return bool(res), newVector3FromPointer(unsafe.Pointer(&ccollisionPoint))
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
