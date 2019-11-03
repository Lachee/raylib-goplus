package raylib

/*
Function Bindings for raygui.
author: lachee
source: https://github.com/raysan5/raygui/blob/master/src/raygui.h
*/

/*
#include "raygui.h"
#include <stdlib.h>
*/
import "C"
import "unsafe"

//GuiState the state of the GUI
type GuiState int32

//GuiControl the ID of a control
type GuiControl int32

//GuiProperty the property of a control
type GuiProperty int32

const (
	GuiStateNormal GuiState = iota
	GuiStateFocused
	GuiStatePressed
	GuiStateDisabled
)

const (
	GuiControlDefault GuiControl = iota
	GuiControlLabel
	GuiControlButton
	GuiControlToggle
	GuiControlSlider
	GuiControlProgressBar
	GuiControlCheckBox
	GuiControlComboBox
	GuiControlDropDownBox
	GuiControlTextBox
	GuiControlValueBox
	GuiControlSpinner
	GuiControlListView
	GuiControlColorPicker
	GuiControlScrollBar
	GuiControlStatusBar
)

const (
	GuiPropertyBorderColorNormal GuiProperty = iota
	GuiPropertyBaseColorNormal
	GuiPropertyTextColorNormal
	GuiPropertyBorderColorFocused
)

//GuiEnable enables GUI controls
func GuiEnable() { C.GuiEnable() }

//GuiDisable disables GUI controls
func GuiDisable() { C.GuiDisable() }

//GuiLock locks the GUI controls
func GuiLock() { C.GuiLock() }

//GuiUnlock unlocks the GUI controls
func GuiUnlock() { C.GuiUnlock() }

//GuiFade sets teh GUI controls alpha [0..1]
func GuiFade(alpha float32) {
	f := C.float(alpha)
	C.GuiFade(f)
}

//GuiSetState sets the GUI state
func GuiSetState(state GuiState) {
	i := C.int(state)
	C.GuiSetState(i)
}

//GuiGetState gets the GUI current state
func GuiGetState() GuiState {
	s := C.GuiGetState()
	return GuiState(s)
}

func GuiSetFont(f Font) {
	cf := *f.cptr()
	C.GuiSetFont(cf)
}
func GuiGetFont() Font {
	cf := C.GuiGetFont()
	return newFontFromPointer(unsafe.Pointer(&cf))
}

//GuiSetStyle sets a style property
func GuiSetStyle(control GuiControl, property GuiProperty, value int) {
	c0 := C.int(control)
	c1 := C.int(property)
	c2 := C.int(value)
	C.GuiSetStyle(c0, c1, c2)
}

//GuiGetStyle gets a style property
func GuiGetStyle(control GuiControl, property GuiProperty) int {
	c0 := C.int(control)
	c1 := C.int(property)
	o := C.GuiGetStyle(c0, c1)
	return int(o)
}

//GuiGetStyleColor gets the color of a style property
func GuiGetStyleColor(control GuiControl, property GuiProperty) Color {
	return NewColorInt(GuiGetStyle(control, property))
}

// Container Seperator Controls

//GuiWindowBox shows a window that can be closed
func GuiWindowBox(bounds Rectangle, title string) bool {
	bptr := *bounds.cptr()
	ctitle := C.CString(title)
	defer C.free(unsafe.Pointer(ctitle))
	o := C.GuiWindowBox(bptr, ctitle)
	return bool(o)
}

//GuiGroupBox is a group box control with a name
func GuiGroupBox(bounds Rectangle, title string) {
	bptr := *bounds.cptr()
	ctitle := C.CString(title)
	defer C.free(unsafe.Pointer(ctitle))
	C.GuiGroupBox(bptr, ctitle)
}

//GuiLine is a line seperator, that optionally could contain text
func GuiLine(bounds Rectangle, title string) {
	bptr := *bounds.cptr()
	ctitle := C.CString(title)
	defer C.free(unsafe.Pointer(ctitle))
	C.GuiLine(bptr, ctitle)
}

//GuiPanel is a panel
func GuiPanel(bounds Rectangle) {
	bptr := *bounds.cptr()
	C.GuiPanel(bptr)
}

//GuiPanel is a panel
func GuiScrollPanel(bounds, content Rectangle, scroll *Vector2) Rectangle {
	bptr := *bounds.cptr()
	cptr := *content.cptr()
	sptr := scroll.cptr()
	r := C.GuiScrollPanel(bptr, cptr, sptr)
	return newRectangleFromPointer(unsafe.Pointer(&r))
}

//GuiLabel shows text
func GuiLabel(bounds Rectangle, text string) {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	C.GuiLabel(*bounds.cptr(), ctext)
}

//GuiButton returns true when clicked
func GuiButton(bounds Rectangle, text string) bool {
	ct := C.CString(text)
	cb := *bounds.cptr()
	defer C.free(unsafe.Pointer(ct))
	o := C.GuiButton(cb, ct)
	return bool(o)
}

//GuiLabelButton control
func GuiLabelButton(bounds Rectangle, text string) bool {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	return bool(C.GuiLabelButton(*bounds.cptr(), ctext))
}

//GuiImageButton adds a button with an image within it
func GuiImageButton(bounds Rectangle, text string, texture Texture2D) bool {
	cbounds := *bounds.cptr()
	ctexture := *texture.cptr()
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	res := C.GuiImageButton(cbounds, ctext, ctexture)
	return bool(res)
}

//GuiImageButtonEx adds a button with an image within it
func GuiImageButtonEx(bounds Rectangle, text string, texture Texture2D, texSource Rectangle) bool {
	cbounds := *bounds.cptr()
	ctexSource := *texSource.cptr()
	ctexture := *texture.cptr()
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	res := C.GuiImageButtonEx(cbounds, ctext, ctexture, ctexSource)
	return bool(res)
}

//GuiToggle returns true when active
func GuiToggle(bounds Rectangle, text string, active bool) bool {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	return bool(C.GuiToggle(*bounds.cptr(), ctext, C.bool(active)))
}

//GuiToggleGroup returns active toggle index
func GuiToggleGroup(bounds Rectangle, text string, active int) int {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	return int(C.GuiToggleGroup(*bounds.cptr(), ctext, C.int(active)))
}

//GuiCheckBox returns true when checked
func GuiCheckBox(bounds Rectangle, text string, checked bool) bool {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	return bool(C.GuiCheckBox(*bounds.cptr(), ctext, C.bool(checked)))
}
func GuiDropdownBox(bounds Rectangle, text string, active int, editMode bool) (bool, int) {
	ctext := C.CString(text)
	cactive := C.int(active)
	defer C.free(unsafe.Pointer(ctext))
	res := C.GuiDropdownBox(*bounds.cptr(), ctext, &cactive, C.bool(editMode))
	return bool(res), int(cactive)
}
func GuiSpinner(bounds Rectangle, text string, value int, min int, max int, editMode bool) (bool, int) {
	cbounds := *bounds.cptr()
	ctext := C.CString(text)
	cvalue := C.int(value)
	defer C.free(unsafe.Pointer(ctext))
	res := C.GuiSpinner(cbounds, ctext, &cvalue, C.int(min), C.int(max), C.bool(editMode))
	return bool(res), int(cvalue)
}
func GuiValueBox(bounds Rectangle, text string, value int, min int, max int, editMode bool) (bool, int) {
	cbounds := *bounds.cptr()
	ctext := C.CString(text)
	cvalue := C.int(value)
	defer C.free(unsafe.Pointer(ctext))
	res := C.GuiValueBox(cbounds, ctext, &cvalue, C.int(min), C.int(max), C.bool(editMode))
	return bool(res), int(cvalue)
}
func GuiTextBox(bounds Rectangle, text string, textSize int, editMode bool) (bool, string) {
	cbounds := *bounds.cptr()
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	res := C.GuiTextBox(cbounds, ctext, C.int(textSize), C.bool(editMode))
	return bool(res), C.GoString(ctext)
}
func GuiTextBoxMulti(bounds Rectangle, text string, textSize int, editMode bool) (bool, string) {
	cbounds := *bounds.cptr()
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	res := C.GuiTextBoxMulti(cbounds, ctext, C.int(textSize), C.bool(editMode))
	return bool(res), C.GoString(ctext)
}
func GuiSlider(bounds Rectangle, textLeft string, textRight string, value float32, min float32, max float32) float32 {
	cbounds := *bounds.cptr()
	ctextLeft := C.CString(textLeft)
	ctextRight := C.CString(textRight)
	defer C.free(unsafe.Pointer(ctextLeft))
	defer C.free(unsafe.Pointer(ctextRight))
	res := C.GuiSlider(cbounds, ctextLeft, ctextRight, C.float(value), C.float(min), C.float(max))
	return float32(res)
}
func GuiSliderBar(bounds Rectangle, textLeft string, textRight string, value float32, min float32, max float32) float32 {
	cbounds := *bounds.cptr()
	ctextLeft := C.CString(textLeft)
	ctextRight := C.CString(textRight)
	defer C.free(unsafe.Pointer(ctextLeft))
	defer C.free(unsafe.Pointer(ctextRight))
	res := C.GuiSliderBar(cbounds, ctextLeft, ctextRight, C.float(value), C.float(min), C.float(max))
	return float32(res)
}
func GuiProgressBar(bounds Rectangle, textLeft string, textRight string, value float32, min float32, max float32) float32 {
	cbounds := *bounds.cptr()
	ctextLeft := C.CString(textLeft)
	ctextRight := C.CString(textRight)
	defer C.free(unsafe.Pointer(ctextLeft))
	defer C.free(unsafe.Pointer(ctextRight))
	res := C.GuiProgressBar(cbounds, ctextLeft, ctextRight, C.float(value), C.float(min), C.float(max))
	return float32(res)
}
func GuiStatusBar(bounds Rectangle, text string) {
	cbounds := *bounds.cptr()
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	C.GuiStatusBar(cbounds, ctext)
}
func GuiDummyRect(bounds Rectangle, text string) {
	cbounds := *bounds.cptr()
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	C.GuiDummyRec(cbounds, ctext)
}
func GuiScrollBar(bounds Rectangle, value, min, max int) int {
	cbounds := *bounds.cptr()
	res := C.GuiScrollBar(cbounds, C.int(value), C.int(min), C.int(max))
	return int(res)
}
func GuiGrid(bounds Rectangle, spacing float32, subdivs int) Vector2 {
	cbounds := *bounds.cptr()
	res := C.GuiGrid(cbounds, C.float(spacing), C.int(subdivs))
	return newVector2FromPointer(unsafe.Pointer(&res))
}
