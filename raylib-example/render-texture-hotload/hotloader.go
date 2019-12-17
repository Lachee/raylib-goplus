package main

//ATTENTION: This example is a work in progress and there is no garuantee that it will work.

import (
	"errors"
	"fmt"
	"path/filepath"
	"time"
	"unsafe"

	r "github.com/lachee/raylib-goplus/raylib"
	"github.com/radovskyb/watcher"
)

//HotLoadable can be dynamically reloaded
type HotLoadable interface {
	HotLoad(absoluteFilePath string) error
}

//HotLoader manages hot loading of files
type HotLoader struct {
	watcher   *watcher.Watcher
	loadables map[string]HotLoadable
}

//NewHotLoader creates a new hot loader
func NewHotLoader(pollrate time.Duration) *HotLoader {
	loader := &HotLoader{
		watcher:   watcher.New(),
		loadables: make(map[string]HotLoadable),
	}

	loader.watcher.SetMaxEvents(1)
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

//Load adds teh hot loadable to the map at the specifed file path and then calls HotLoad.
// Returns the absolute path
func (hl *HotLoader) Load(filePath string, hotLoadable HotLoadable) (string, error) {
	if hotLoadable == nil {
		return "", errors.New("hot loadable cannot be nil")
	}

	if filePath == "" {
		return "", errors.New("filepath cannot be empty")
	}

	fp, err := filepath.Abs(filePath)
	if err != nil {
		return fp, err
	}

	if hl.loadables[fp] != nil {
		return fp, errors.New("filepath already exists")
	}

	hl.loadables[fp] = hotLoadable
	hl.watcher.Add(fp)
	return fp, hotLoadable.HotLoad(fp)
}

//Update the hot loader and reload any sources.
func (hl *HotLoader) Update() error {
	select {
	//We have an event, so lets handle the update
	case event := <-hl.watcher.Event:
		r.Trace("Hot Loading", event.Path)

		//Check for shaders. If we have a shader then reload it.
		if ld := hl.loadables[event.Path]; ld != nil {
			r.Trace("Hot Loading ", event.Path)
			if err := ld.HotLoad(event.Path); err != nil {
				r.TraceError("Hot Loading Exception", err)
			}
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

//HotShader is a hot loadable shader
type HotShader struct {
	Shader     *r.Shader
	FsFileName string
	VsFileName string
}

//HotLoad reloads the shader
func (hs *HotShader) HotLoad(afp string) error {
	if hs.Shader == nil {
		return errors.New("Shader not yet initialized")
	}

	r.Trace("Reloading Shader ", hs.Shader)
	hs.Shader.Unload()
	*hs.Shader = *r.LoadShader(hs.VsFileName, hs.FsFileName)
	if hs.Shader == nil {
		return errors.New("Failed to load the shader")
	}

	return nil
}

//LoadShader loads a new shader
func (hl *HotLoader) LoadShader(vsFileName string, fsFileName string) *HotShader {
	hs := &HotShader{
		VsFileName: vsFileName,
		FsFileName: fsFileName,
	}

	//Prepare the VS link
	if vsFileName != "" {
		hl.Load(vsFileName, hs)
	}

	//Prepare the FS link
	if fsFileName != "" {
		hl.Load(fsFileName, hs)
	}

	//Load the actual shader
	shader := r.LoadShader(vsFileName, fsFileName)
	hs.Shader = shader

	return hs
}

//HotTexture is a texture that can be hotloaded. The texture is updated with the new pixels of the image.
type HotTexture r.Texture2D

//HotLoad reloads the texture
func (ht HotTexture) HotLoad(afp string) error {
	if ht.Id == 0 {
		return errors.New("hot texture isn't loaded")
	}

	//Load the image
	img := r.LoadImage(afp)
	defer img.Unload()

	//Get the image pixels and apply it to the texture
	pixels := img.GetPixels()
	r.UpdateTexture((*r.Texture2D)(unsafe.Pointer(&ht)), pixels)
	return nil
}

//LoadTexture2D hot loads a new texture
func (hl *HotLoader) LoadTexture2D(fileName string) HotTexture {
	t := r.LoadTexture(fileName)
	ht := *(*HotTexture)(unsafe.Pointer(&t))
	hl.Load(fileName, ht)
	return ht
}

//ToTexture2D converts the HotTexture to a Texture2D
func (ht HotTexture) ToTexture2D() r.Texture2D {
	return *(*r.Texture2D)(unsafe.Pointer(&ht))
}
