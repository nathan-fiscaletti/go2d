package go2d

type Drawable struct {
    Layer      int
    Visible    bool
    Bounds     Rect
    Constraint Rect

    Velocity      *Vector
    OnConstrained func(RectSide)
}

func (this *Drawable) IsCollidingWith(other *Drawable) bool {
    return this.Bounds.IntersectsWith(other.Bounds)
}

func (this *Drawable) MoveTo(pos Vector) {
    this.Bounds.Vector = pos
}

func (this *Drawable) Push(distance Vector) {
    this.Bounds.Vector = this.Bounds.Vector.Plus(distance)
}

func (this *Drawable) FixedUpdate() {
    if this.Velocity != nil {
        this.Push(*this.Velocity)
    }
}

func (this *Drawable) Constrain(r Rect) {
    out := this.Bounds.Constrain(r)
    if this.OnConstrained != nil {
        for _,side := range out {
            this.OnConstrained(side)
        }
    }
}