package main

import (
	"errors"
	"fmt"
	"path/filepath"
	"time"

	r "github.com/lachee/raylib-goplus/raylib"
	"github.com/radovskyb/watcher"
)

type lShader struct {
	shader     *r.Shader
	fsFileName string
	vsFileName string
}

func (ls *lShader) hotload() {
	if ls.shader != nil {
		r.TraceLog(r.LogTrace, "Reloading Shader", ls.shader)
		ls.shader.Unload()
		*ls.shader = *r.LoadShader(ls.vsFileName, ls.fsFileName)
	}
}

//HotLoader manages hot loading of files
type HotLoader struct {
	watcher  *watcher.Watcher
	shaders  map[string]*lShader
	textures map[string]*r.Texture2D
}

//NewHotLoader creates a new hot loader
func NewHotLoader(pollrate time.Duration) *HotLoader {
	loader := &HotLoader{
		watcher:  watcher.New(),
		shaders:  make(map[string]*lShader),
		textures: make(map[string]*r.Texture2D),
	}

	loader.watcher.FilterOps(watcher.Write)

	go func() {
		if err := loader.watcher.Start(pollrate); err != nil {
			fmt.Println("Failed", err)
		}
	}()

	r.RegisterUnloadable(loader)
	return loader
}

//Unload unloads the hot loader
func (hl *HotLoader) Unload() {
	r.UnregisterUnloadable(hl)
	hl.watcher.Close()
}

//LoadShader loads a new shader to the hot loader
func (hl *HotLoader) LoadShader(vsFileName string, fsFileName string) *r.Shader {
	shader := r.LoadShader(vsFileName, fsFileName)
	lshader := &lShader{
		shader:     shader,
		vsFileName: vsFileName,
		fsFileName: fsFileName,
	}

	hl.watcher.Add(fsFileName)

	if absolute, err := filepath.Abs(vsFileName); err == nil {
		hl.watcher.Add(absolute)
		hl.shaders[absolute] = lshader
	}

	if absolute, err := filepath.Abs(fsFileName); err == nil {
		hl.watcher.Add(absolute)
		hl.shaders[absolute] = lshader
	}

	return shader
}

//LoadTexture loads a texture
func (hl *HotLoader) LoadTexture(fileName string) *r.Texture2D {
	texture := r.LoadTexture(fileName)
	if absolute, err := filepath.Abs(fileName); err == nil {
		hl.watcher.Add(absolute)
		hl.textures[absolute] = &texture
	}
	return &texture
}

//Update the hot loader and reload any sources.
func (hl *HotLoader) Update() error {
	select {
	//We have an event, so lets handle the update
	case event := <-hl.watcher.Event:
		r.TraceLog(r.LogTrace, "Hot Loading", event.Path)

		//Check for shaders. If we have a shader then reload it.
		if shader := hl.shaders[event.Path]; shader != nil {
			r.TraceLog(r.LogInfo, "Hot Loading Shader", event.Path)
			shader.hotload()
		}

		//Check for textures. If we have a texture then reload it.
		if texture := hl.textures[event.Path]; texture != nil {
			r.TraceLog(r.LogInfo, "Hot Loading Texture", event.Path)
			updateTextureFromFile(*texture, event.Path)
		}

	//Check if the watcher errors
	case err := <-hl.watcher.Error:
		return err

	//Check if the watch has closed
	case <-hl.watcher.Closed:
		return errors.New("watcher has closed")

	//None of the channels returned, so continue on our merry way
	default:
	}

	return nil
}

//LoadImageFromGo Creates a new image from a Go Image
func updateTextureFromFile(texture r.Texture2D, fileName string) {
	img := r.LoadImage(fileName)
	defer img.Unload()
	pixels := img.GetPixelsNormalized()
	fmt.Println(pixels)
	//texture.UpdateTexture(pixels)
}
