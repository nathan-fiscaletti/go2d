package go2d

type Controllable interface {
    KeyUp(scancode int, rn rune, name string)
    KeyDown(scancode int, rn rune, name string)
}