package go2d

type Entity struct {
    Layer      int
    Visible    bool
    Bounds     Rect
    Constraint Rect

    Velocity      *Vector
    OnConstrained func(RectSide)
}

func (this *Entity) IsCollidingWith(other *Entity) bool {
    return this.Bounds.IntersectsWith(other.Bounds)
}

func (this *Entity) MoveTo(pos Vector) {
    this.Bounds.Vector = pos
}

func (this *Entity) Push(distance Vector) {
    this.Bounds.Vector = this.Bounds.Vector.Plus(distance)
}

func (this *Entity) FixedUpdate(engine *Engine, scene *Scene) {
    if this.Velocity != nil {
        this.Push(*this.Velocity)
    }
}

func (this *Entity) Constrain(r Rect) {
    out := this.Bounds.Constrain(r)
    if this.OnConstrained != nil {
        for _,side := range out {
            this.OnConstrained(side)
        }
    }
}