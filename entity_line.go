package go2d

import (
	"github.com/tfriedel6/canvas"
)

// LienCapStyle is a type that represents the different line cap styles
type LineCapStyle uint8

const (
	LINECAP_MITER = iota
	LINECAP_BEVEL
	LINECAP_ROUND
	LINECAP_SQUARE
	LINECAP_BUTT
)

func (l LineCapStyle) fillLineCapStyle(c *canvas.Canvas) {
	if l == LINECAP_MITER {
		c.SetLineCap(canvas.Miter)
	} else if l == LINECAP_BEVEL {
		c.SetLineCap(canvas.Bevel)
	} else if l == LINECAP_ROUND {
		c.SetLineCap(canvas.Round)
	} else if l == LINECAP_SQUARE {
		c.SetLineCap(canvas.Square)
	} else if l == LINECAP_BUTT {
		c.SetLineCap(canvas.Butt)
	}
}

// LineEntity is a simple line entity that can be used to draw lines on the screen.
type LineEntity struct {
	Entity

	direction Vector
	capStyle  LineCapStyle
	length    float64
	thickness int
	color     string
}

// NewLineEntitySimple creates a new line entity with the given from and to
// vectors, thickness, and color.
func NewLineEntitySimple(from Vector, to Vector, thickness int, color string) *LineEntity {
	direction := from.DirectionTo(to)
	length := from.DistanceTo(to)
	e := NewLineEntity(from, direction, length, thickness, color)
	return e
}

// NewLineEntity creates a new line entity with the given from vector, direction
// vector, length, thickness, and color.
func NewLineEntity(from Vector, direction Vector, length float64, thickness int, color string) *LineEntity {
	return &LineEntity{
		Entity: Entity{
			Bounds: Rect{
				Vector: from,
			},
		},
		direction: direction,
		length:    length,
		thickness: thickness,
		color:     color,
	}
}

// SetCapStyle sets the cap style of the line.
func (this *LineEntity) SetCapStyle(capStyle LineCapStyle) {
	this.capStyle = capStyle
}

// Render renders the line entity.
func (this *LineEntity) Render(e *Engine) {
	e.Canvas.SetLineWidth(float64(this.thickness))
	e.Canvas.SetStrokeStyle(this.color)
	this.capStyle.fillLineCapStyle(e.Canvas)

	e.Canvas.BeginPath()
	e.Canvas.MoveTo(this.Bounds.X, this.Bounds.Y)

	to := Vector{
		X: this.Bounds.X + (this.direction.X * this.length),
		Y: this.Bounds.Y + (this.direction.Y * this.length),
	}

	e.Canvas.LineTo(to.X, to.Y)

	e.Canvas.ClosePath()
	e.Canvas.Stroke()
}

// Update updates the line entity.
func (this *LineEntity) Update(e *Engine) {
	this.Entity.Update()
}

// GetEntity returns the entity of the line entity.
func (this *LineEntity) GetEntity() *Entity {
	return &this.Entity
}
