package raylib

//#include "raylib.h"
//#include <stdlib.h>
import "C"

//DrawFPS shows the current FPS
func DrawFPS(x, y int) {
	C.DrawFPS(C.int(x), C.int(y))
}
