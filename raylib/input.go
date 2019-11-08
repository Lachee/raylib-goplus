package raylib

type MouseButton int32

const (
	MouseLeftButton MouseButton = iota
	MouseRightButton
	MouseMiddleButton
)

type GamepadNumber int32

const (
	GamepadPlayer1 GamepadNumber = iota
	GamepadPlayer2
	GamepadPlayer3
	GamepadPlayer4
)

type GamepadButton int32

const (
	GamepadButtonUnkown GamepadButton = iota

	GamepadButtonLeftFaceUp
	GamepadButtonLeftFaceRight
	GamepadButtonLeftFaceDown
	GamepadButtonLeftFaceLeft

	GamepadButtonRightFaceUp
	GamepadButtonRightFaceRight
	GamepadButtonRightFaceDown
	GamepadButtonRightFaceLeft

	GamepadButtonLeftTrigger1
	GamepadButtonLeftTrigger2
	GamepadButtonRightTrigger1
	GamepadButtonRightTrigger2

	GamepadButtonMiddleLeft
	GamepadButtonMiddle
	GamepadButtonMiddleRight

	GamepadButtonLeftThumb
	GamepadButtonRightThumb
)

type GamepadAxis int32

const (
	GamepadAxisUnkown GamepadAxis = iota
	GamepadAxisLeftX
	GamepadAxisLeftY
	GamepadAxisRightX
	GamepadAxisRightY
	GamepadAxisLeftTrigger
	GamepadAxisRightTrigger
)
