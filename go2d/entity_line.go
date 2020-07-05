package go2d

import(
    "github.com/tfriedel6/canvas"
)

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

type LineEntity struct {
    Entity

    direction Vector
    capStyle  LineCapStyle
    length    float64
    thickness int
    color     string
}

func NewLineEntitySimple(from Vector, to Vector, thickness int, color string) *LineEntity {
    direction := from.DirectionTo(to)
    length := from.DistanceTo(to)
    e := NewLineEntity(from, direction, length, thickness, color)
    return e
}

func NewLineEntity(from Vector, direction Vector, length float64, thickness int, color string) *LineEntity {
    return &LineEntity {
        Entity: Entity {
            Bounds: Rect {
                Vector: from,
            },
        },
        direction: direction,
        length:    length,
        thickness: thickness,
        color:     color,
    }
}

func (this *LineEntity) SetCapStyle(capStyle LineCapStyle) {
    this.capStyle = capStyle
}

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

func (this *LineEntity) Update(e *Engine) {
    this.Entity.Update()
}

