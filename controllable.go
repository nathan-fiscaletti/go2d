package go2d

const (
	MOUSE_BUTTON_LEFT   = 1
	MOUSE_BUTTON_MIDDLE = 2
	MOUSE_BUTTON_RIGHT  = 3
	MOUSE_BUTTON_X1     = 4
	MOUSE_BUTTON_X2     = 5
)

type IKeyUp interface {
	KeyUp(scanCode int, rn rune, name string)
}

type IKeyDown interface {
	KeyDown(scanCode int, rn rune, name string)
}

type IKeyChar interface {
	KeyChar(rn rune)
}

type IMouseDown interface {
	MouseDown(button int, pos Vector)
}

type IMouseUp interface {
	MouseUp(button int, pos Vector)
}

type IMouseMove interface {
	MouseMove(pos Vector)
}
