package raylib

/*
#include "raylib.h"
#include <stdlib.h>
#include "go.h"
*/
import "C"

type Key uint32

const (
	// alphanumeric KeyS
	KeyApostrophe Key = 39
	KeyComma          = 44
	KeyMinus          = 45
	KeyPeriod         = 46
	KeySlash          = 47
	KeyZero           = 48
	KeyOne            = 49
	KeyTwo            = 50
	KeyThree          = 51
	KeyFour           = 52
	KeyFive           = 53
	KeySix            = 54
	KeySeven          = 55
	KeyEight          = 56
	KeyNine           = 57
	KeySemicolon      = 59
	KeyEqual          = 61
	KeyA              = 65
	KeyB              = 66
	KeyC              = 67
	KeyD              = 68
	KeyE              = 69
	KeyF              = 70
	KeyG              = 71
	KeyH              = 72
	KeyI              = 73
	KeyJ              = 74
	KeyK              = 75
	KeyL              = 76
	KeyM              = 77
	KeyN              = 78
	KeyO              = 79
	KeyP              = 80
	KeyQ              = 81
	KeyR              = 82
	KeyS              = 83
	KeyT              = 84
	KeyU              = 85
	KeyV              = 86
	keyw              = 87
	KeyX              = 88
	KeyY              = 89
	KeyZ              = 90

	// function KeyS
	KeySpace        = 32
	KeyEscape       = 256
	KeyEnter        = 257
	KeyTab          = 258
	KeyBackspace    = 259
	KeyInsert       = 260
	KeyDelete       = 261
	KeyRight        = 262
	KeyLeft         = 263
	KeyDown         = 264
	KeyUp           = 265
	KeyPageUp       = 266
	KeyPageDown     = 267
	KeyHome         = 268
	KeyEnd          = 269
	KeyCapsLock     = 280
	KeyScrollLock   = 281
	KeyNumLock      = 282
	KeyPrintScreen  = 283
	KeyPause        = 284
	KeyF1           = 290
	KeyF2           = 291
	KeyF3           = 292
	KeyF4           = 293
	KeyF5           = 294
	KeyF6           = 295
	KeyF7           = 296
	KeyF8           = 297
	KeyF9           = 298
	KeyF10          = 299
	KeyF11          = 300
	KeyF12          = 301
	KeyLeftShift    = 340
	KeyLeftControl  = 341
	KeyLeftAlt      = 342
	KeyLeftSuper    = 343
	KeyRightShift   = 344
	KeyRightControl = 345
	KeyRightAlt     = 346
	KeyRightSuper   = 347
	KeyKbMenu       = 348
	KeyLeftBracket  = 91
	KeyBackslash    = 92
	KeyRightBracket = 93
	KeyGrave        = 96

	// KeyPad KeyS
	KeyKp0        = 320
	KeyKp1        = 321
	KeyKp2        = 322
	KeyKp3        = 323
	KeyKp4        = 324
	KeyKp5        = 325
	KeyKp6        = 326
	KeyKp7        = 327
	KeyKp8        = 328
	KeyKp9        = 329
	KeyKpDecimal  = 330
	KeyKpDivide   = 331
	KeyKpMultiply = 332
	KeyKpSubtract = 333
	KeyKpAdd      = 334
	KeyKpEnter    = 335
	KeyKpEqual    = 336
)

//IsKeyPressed : Detect if a key has been pressed once
func IsKeyPressed(key Key) bool {
	res := C.IsKeyPressed(C.int(key))
	return bool(res)
}

//IsKeyDown : Detect if a key is being pressed
func IsKeyDown(key Key) bool {
	res := C.IsKeyDown(C.int(key))
	return bool(res)
}

//IsKeyReleased : Detect if a key has been released once
func IsKeyReleased(key Key) bool {
	res := C.IsKeyReleased(C.int(key))
	return bool(res)
}

//IsKeyUp : Detect if a key is NOT being pressed
func IsKeyUp(key Key) bool {
	res := C.IsKeyUp(C.int(key))
	return bool(res)
}

//GetKeyPressed : Get latest key pressed
func GetKeyPressed() Key {
	res := C.GetKeyPressed()
	return Key(res)
}

//SetExitKey : Set a custom key to exit program (default is ESC)
func SetExitKey(key Key) {
	C.SetExitKey(C.int(key))
}
