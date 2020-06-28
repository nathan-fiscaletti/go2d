package graphics

import(
    "../metrics"
)

type Drawable struct {
    Layer      int
    Visible    bool
    Bounds     metrics.Rect
    Constraint metrics.Rect
}

type IGameObject interface {
    Update()
    FixedUpdate()
}

type IConstrainable interface {
    OnConstrained(constraintAxis metrics.RectSide)
}

func (this *Drawable) IsCollidingWith(other *Drawable) bool {
    return this.Bounds.IntersectsWith(other.Bounds)
}

func (this *Drawable) MoveTo(pos metrics.Vector) {
    this.Bounds.Vector = pos
}

func (this *Drawable) Push(distance metrics.Vector) {
    this.Bounds.Vector = this.Bounds.Vector.Plus(distance)
}

func (this *Drawable) Constrain(r metrics.Rect) {
    out := this.Bounds.Constrain(r)
    constrainable, isConstrainable := (interface{}(this)).(IConstrainable)
    if isConstrainable {
        for _,side := range out {
            constrainable.OnConstrained(side)
        }
    }
}