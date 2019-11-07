
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


### License
This project is still a work in progress, but the license will be `zlib/libpng` to keep it inline with Raylib license.
