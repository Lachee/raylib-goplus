package raylib

/*
#include "raylib.h"
#include <stdlib.h>
#include "go.h"
#include "raygui.h"
*/
import "C"

//GuiEnable : Enable gui controls (global state)
func GuiEnable() {
	C.GuiEnable()
}

//GuiDisable : Disable gui controls (global state)
func GuiDisable() {
	C.GuiDisable()
}

//GuiLock : Lock gui controls (global state)
func GuiLock() {
	C.GuiLock()
}

//GuiUnlock : Unlock gui controls (global state)
func GuiUnlock() {
	C.GuiUnlock()
}

//GuiFade : Set gui controls alpha (global state), alpha goes from 0.0f to 1.0f
func GuiFade(alpha float32) {
	C.GuiFade(C.float(alpha))
}

//GuiSetState : Set gui state (global state)
func GuiSetState(state int) {
	C.GuiSetState(C.int(state))
}

//GuiGetState : Get gui state (global state)
func GuiGetState() int {
	res := C.GuiGetState()
	return int(res)
}

//GuiSetFont : Set gui custom font (global state)
func GuiSetFont(font Font) {
	cfont := *font.cptr()
	C.GuiSetFont(cfont)
}

//GuiGetFont : Get gui custom font (global state)
func GuiGetFont() Font {
	res := C.GuiGetFont()
	return newFontFromPointer(unsafe.Pointer(&res))
}

//GuiSetStyle : Set one style property
func GuiSetStyle(control GuiControl, property GuiProperty, value int) {
	C.GuiSetStyle(C.int(control), C.int(property), C.int(value))
}

//GuiGetStyle : Get one style property
func GuiGetStyle(control GuiControl, property GuiProperty) int {
	res := C.GuiGetStyle(C.int(control), C.int(property))
	return int(res)
}

//GuiWindowBox : Window Box control, shows a window that can be closed
func GuiWindowBox(bounds Rectangle, title string) bool {
	ctitle := C.CString(title)
	defer C.free(unsafe.Pointer(ctitle))
	cbounds := *bounds.cptr()
	res := C.GuiWindowBox(cbounds, ctitle)
	return bool(res)
}

//GuiGroupBox : Group Box control with text name
func GuiGroupBox(bounds Rectangle, text string) {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	cbounds := *bounds.cptr()
	C.GuiGroupBox(cbounds, ctext)
}

//GuiLine : Line separator control, could contain text
func GuiLine(bounds Rectangle, text string) {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	cbounds := *bounds.cptr()
	C.GuiLine(cbounds, ctext)
}

//GuiPanel : Panel control, useful to group controls
func GuiPanel(bounds Rectangle) {
	cbounds := *bounds.cptr()
	C.GuiPanel(cbounds)
}

//GuiScrollPanel : Scroll Panel control
func GuiScrollPanel(bounds Rectangle, content Rectangle, scroll Vector2) (Rectangle, Vector2) {
	cscroll := scroll.cptr()
	ccontent := *content.cptr()
	cbounds := *bounds.cptr()
	res := C.GuiScrollPanel(cbounds, ccontent, cscroll)
	return newRectangleFromPointer(unsafe.Pointer(&res)), newVector2FromPointer(unsafe.Pointer(&cscroll))
}

//GuiLabel : Label control, shows text
func GuiLabel(bounds Rectangle, text string) {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	cbounds := *bounds.cptr()
	C.GuiLabel(cbounds, ctext)
}

//GuiButton : Button control, returns true when clicked
func GuiButton(bounds Rectangle, text string) bool {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	cbounds := *bounds.cptr()
	res := C.GuiButton(cbounds, ctext)
	return bool(res)
}

//GuiLabelButton : Label button control, show true when clicked
func GuiLabelButton(bounds Rectangle, text string) bool {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	cbounds := *bounds.cptr()
	res := C.GuiLabelButton(cbounds, ctext)
	return bool(res)
}

//GuiImageButton : Image button control, returns true when clicked
func GuiImageButton(bounds Rectangle, text string, texture Texture2D) bool {
	ctexture := *texture.cptr()
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	cbounds := *bounds.cptr()
	res := C.GuiImageButton(cbounds, ctext, ctexture)
	return bool(res)
}

//GuiImageButtonEx : Image button extended control, returns true when clicked
func GuiImageButtonEx(bounds Rectangle, text string, texture Texture2D, texSource Rectangle) bool {
	ctexSource := *texSource.cptr()
	ctexture := *texture.cptr()
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	cbounds := *bounds.cptr()
	res := C.GuiImageButtonEx(cbounds, ctext, ctexture, ctexSource)
	return bool(res)
}

//GuiToggle : Toggle Button control, returns true when active
func GuiToggle(bounds Rectangle, text string, active bool) bool {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	cbounds := *bounds.cptr()
	res := C.GuiToggle(cbounds, ctext, C.bool(active))
	return bool(res)
}

//GuiToggleGroup : Toggle Group control, returns active toggle index
func GuiToggleGroup(bounds Rectangle, text string, active int) int {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	cbounds := *bounds.cptr()
	res := C.GuiToggleGroup(cbounds, ctext, C.int(active))
	return int(res)
}

//GuiCheckBox : Check Box control, returns true when active
func GuiCheckBox(bounds Rectangle, text string, checked bool) bool {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	cbounds := *bounds.cptr()
	res := C.GuiCheckBox(cbounds, ctext, C.bool(checked))
	return bool(res)
}

//GuiComboBox : Combo Box control, returns selected item index
func GuiComboBox(bounds Rectangle, text string, active int) int {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	cbounds := *bounds.cptr()
	res := C.GuiComboBox(cbounds, ctext, C.int(active))
	return int(res)
}

//GuiDropdownBox : Dropdown Box control, returns selected item
func GuiDropdownBox(bounds Rectangle, text string, active int, editMode bool) (bool, int) {
	cactive := C.int(active)
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	cbounds := *bounds.cptr()
	res := C.GuiDropdownBox(cbounds, ctext, &cactive, C.bool(editMode))
	return bool(res), int(cactive)
}

//GuiSpinner : Spinner control, returns selected value
func GuiSpinner(bounds Rectangle, text string, value int, minValue int, maxValue int, editMode bool) (bool, int) {
	cvalue := C.int(value)
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	cbounds := *bounds.cptr()
	res := C.GuiSpinner(cbounds, ctext, &cvalue, C.int(minValue), C.int(maxValue), C.bool(editMode))
	return bool(res), int(cvalue)
}

//GuiValueBox : Value Box control, updates input text with numbers
func GuiValueBox(bounds Rectangle, text string, value int, minValue int, maxValue int, editMode bool) (bool, int) {
	cvalue := C.int(value)
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	cbounds := *bounds.cptr()
	res := C.GuiValueBox(cbounds, ctext, &cvalue, C.int(minValue), C.int(maxValue), C.bool(editMode))
	return bool(res), int(cvalue)
}

//GuiTextBox : Text Box control, updates input text
func GuiTextBox(bounds Rectangle, text string, textSize int, editMode bool) (bool, string) {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	cbounds := *bounds.cptr()
	res := C.GuiTextBox(cbounds, ctext, C.int(textSize), C.bool(editMode))
	return bool(res), C.GoString(ctext)
}

//GuiTextBoxMulti : Text Box control with multiple lines
func GuiTextBoxMulti(bounds Rectangle, text string, textSize int, editMode bool) (bool, string) {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	cbounds := *bounds.cptr()
	res := C.GuiTextBoxMulti(cbounds, ctext, C.int(textSize), C.bool(editMode))
	return bool(res), C.GoString(ctext)
}

//GuiSlider : Slider control, returns selected value
func GuiSlider(bounds Rectangle, textLeft string, textRight string, value float32, minValue float32, maxValue float32) float32 {
	ctextRight := C.CString(textRight)
	defer C.free(unsafe.Pointer(ctextRight))
	ctextLeft := C.CString(textLeft)
	defer C.free(unsafe.Pointer(ctextLeft))
	cbounds := *bounds.cptr()
	res := C.GuiSlider(cbounds, ctextLeft, ctextRight, C.float(value), C.float(minValue), C.float(maxValue))
	return float32(res)
}

//GuiSliderBar : Slider Bar control, returns selected value
func GuiSliderBar(bounds Rectangle, textLeft string, textRight string, value float32, minValue float32, maxValue float32) float32 {
	ctextRight := C.CString(textRight)
	defer C.free(unsafe.Pointer(ctextRight))
	ctextLeft := C.CString(textLeft)
	defer C.free(unsafe.Pointer(ctextLeft))
	cbounds := *bounds.cptr()
	res := C.GuiSliderBar(cbounds, ctextLeft, ctextRight, C.float(value), C.float(minValue), C.float(maxValue))
	return float32(res)
}

//GuiProgressBar : Progress Bar control, shows current progress value
func GuiProgressBar(bounds Rectangle, textLeft string, textRight string, value float32, minValue float32, maxValue float32) float32 {
	ctextRight := C.CString(textRight)
	defer C.free(unsafe.Pointer(ctextRight))
	ctextLeft := C.CString(textLeft)
	defer C.free(unsafe.Pointer(ctextLeft))
	cbounds := *bounds.cptr()
	res := C.GuiProgressBar(cbounds, ctextLeft, ctextRight, C.float(value), C.float(minValue), C.float(maxValue))
	return float32(res)
}

//GuiStatusBar : Status Bar control, shows info text
func GuiStatusBar(bounds Rectangle, text string) {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	cbounds := *bounds.cptr()
	C.GuiStatusBar(cbounds, ctext)
}

//GuiDummyRec : Dummy control for placeholders
func GuiDummyRec(bounds Rectangle, text string) {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	cbounds := *bounds.cptr()
	C.GuiDummyRec(cbounds, ctext)
}

//GuiScrollBar : Scroll Bar control
func GuiScrollBar(bounds Rectangle, value int, minValue int, maxValue int) int {
	cbounds := *bounds.cptr()
	res := C.GuiScrollBar(cbounds, C.int(value), C.int(minValue), C.int(maxValue))
	return int(res)
}

//GuiGrid : Grid control
func GuiGrid(bounds Rectangle, spacing float32, subdivs int) Vector2 {
	cbounds := *bounds.cptr()
	res := C.GuiGrid(cbounds, C.float(spacing), C.int(subdivs))
	return newVector2FromPointer(unsafe.Pointer(&res))
}

//GuiListView : List View control, returns selected list item index
func GuiListView(bounds Rectangle, text string, scrollIndex int, active int) (int, int) {
	cscrollIndex := C.int(scrollIndex)
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	cbounds := *bounds.cptr()
	res := C.GuiListView(cbounds, ctext, &cscrollIndex, C.int(active))
	return int(res), int(cscrollIndex)
}

//GuiListViewEx : List View with extended parameters
func GuiListViewEx(bounds Rectangle, text []string, count int, focus int, scrollIndex int, active int) (int, int, int) {
	cscrollIndex := C.int(scrollIndex)
	cfocus := C.int(focus)
	cbounds := *bounds.cptr()

	//Copies the string into an array in C memory
	cargs := C.makeCharArray(C.int(len(text)))
	defer C.freeCharArray(cargs, C.int(len(text)))
	for i, s := range sargs {
		C.setArrayString(cargs, C.CString(s), C.int(i))
	}

	//Create the thingy
	res := C.GuiListViewEx(cbounds, cargs, C.int(count), &cfocus, &cscrollIndex, C.int(active))
	return int(res), int(cfocus), int(cscrollIndex)
}

//GuiMessageBox : Message Box control, displays a message
func GuiMessageBox(bounds Rectangle, title string, message string, buttons string) int {
	cbuttons := C.CString(buttons)
	defer C.free(unsafe.Pointer(cbuttons))
	cmessage := C.CString(message)
	defer C.free(unsafe.Pointer(cmessage))
	ctitle := C.CString(title)
	defer C.free(unsafe.Pointer(ctitle))
	cbounds := *bounds.cptr()
	res := C.GuiMessageBox(cbounds, ctitle, cmessage, cbuttons)
	return int(res)
}

//GuiTextInputBox : Text Input Box control, ask for text
func GuiTextInputBox(bounds Rectangle, title string, message string, buttons string, text string) (int, string) {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	cbuttons := C.CString(buttons)
	defer C.free(unsafe.Pointer(cbuttons))
	cmessage := C.CString(message)
	defer C.free(unsafe.Pointer(cmessage))
	ctitle := C.CString(title)
	defer C.free(unsafe.Pointer(ctitle))
	cbounds := *bounds.cptr()
	res := C.GuiTextInputBox(cbounds, ctitle, cmessage, cbuttons, ctext)
	return int(res), C.GoString(ctext)
}

//GuiColorPicker : Color Picker control
func GuiColorPicker(bounds Rectangle, color Color) Color {
	ccolor := *color.cptr()
	cbounds := *bounds.cptr()
	res := C.GuiColorPicker(cbounds, ccolor)
	return newColorFromPointer(unsafe.Pointer(&res))
}

//GuiLoadStyle : Load style file (.rgs)
func GuiLoadStyle(fileName string) {
	cfileName := C.CString(fileName)
	defer C.free(unsafe.Pointer(cfileName))
	C.GuiLoadStyle(cfileName)
}

//GuiLoadStyleDefault : Load style default over global style
func GuiLoadStyleDefault() {
	C.GuiLoadStyleDefault()
}

//GuiIconText : Get text with icon id prepended
func GuiIconText(iconId int, text string) string {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))
	res := C.GuiIconText(C.int(iconId), ctext)
	return C.GoString(res)
}
