
<img align="left" src="https://github.com/Lachee/raylib-goplus/raw/master/logo/raylib_goplus_256x256.png" width=256>

Raylib Go Plus is an improved [Go]([https://golang.org/](https://golang.org/)) bindings for the video game library [Raylib](https://raylib.com/)

This library is a work in progress and still new. We are always looking for improvements and help translating the bindings. Any cross-platform testing is most welcome too.




### Installation
[![GoDoc](https://godoc.org/github.com/Lachee/raylib-goplus/raylib?status.svg)](https://godoc.org/github.com/Lachee/raylib-goplus/raylib)

`go get github.com/lachee/raylib-goplus/raylib`

Its that simple! Once you have fetch the resource, be sure to include it and get going!

```go
package main
import r "github.com/lachee/raylib-goplus/raylib"
func main() {
	r.InitWindow(800, 450, "Raylib Go Plus")
	for !r.WindowShouldClose() {
		r.BeginDrawing()
		r.ClearBackground(r.RayWhite)
		r.DrawText("Woo! Raylib-Go-Plus! Now with ++", 20, 20, 20, r.GopherBlue)
		r.EndDrawing()
	}
	r.CloseWindow()
}
```

### Building / Updating Raylib / Contribution
If you wish to build and update raylib, the project comes with a useful converter.
Make sure you update the source files in `raylib/` and update the `raylib-convert/headings.txt` headings (this tells the converter what to export). Once that is done, simply run `./build.sh`.

I am happy for any contributions. This is a big project with over 477 functions! Its likely I have missed something or something doesn't work right for a particular platform (I only test on Windows).

**NOTE**: All contributions to code within  `_gen.go` files will be rejected!
Those additions must be made within a `raylib-convert/manual/` go file, with the function name as the file name.

### Unloadables
There is an experimental feature in this library called "Unloadables". When possible, Raylib Go Plus will have a record to every object that is loaded via LoadXXXX pattern, and will delete their record when `Unload()` is called on them. This is a useful safety feature to just make sure everything is unloaded.

Because they are tracked, you can call `r.UnloadAll()` at the end of your application to free up all recorded `Unloadables` from the C memory, preventing memory leaks.

Please note that this is an experimental feature, and not all Loadables maybe added to the tracker, nor may they all be added for every Load functions (ie: some GetXXXX may require unloading too). It is recommended **to always free up the object yourself** using `Unload()`

Anything that is loaded using a `LoadXXXX` should be in OOP mode of the converter. The converter will add Unload methods. In cases where they are named Close instead (ie AudioStream), please make the OOP method Unload, and the functional method the original Close.

### License
This project is still a work in progress, but the license will be `zlib/libpng` to keep it inline with Raylib license.
