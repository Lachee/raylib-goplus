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

var guiEnabled bool
var guiLocked bool

//GuiEnabled returns if the gui is currently enabled.
func GuiEnabled() bool {
	return guiEnabled
}

//GuiLocked returns if the gui is currently locked.
func GuiLocked() bool {
	return guiLocked
}
