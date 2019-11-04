package raylib

/*
Function Bindings for raygui.
author: lachee
source: https://github.com/raysan5/raygui/blob/master/src/raygui.h
*/

/*
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

//GuiTextBoxState keeps the state of active textboxes
type GuiTextBoxState struct {
	Cursor int32
	Start  int32
	Index  int32
	Select int32
}

func newGuiTextBoxStateFromPointer(ptr unsafe.Pointer) GuiTextBoxState {
	return *(*GuiTextBoxState)(ptr)
}

/*
func (state *GuiTextBoxState) cptr() *C.GuiTextBoxState {
	return (*C.GuiTextBoxState)(unsafe.Pointer(state))
}
*/
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
