package go2d

type IRender interface {
    Render(c *Engine)
}

type IFixedUpdate interface {
    FixedUpdate(e *Engine)
}

type IUpdate interface {
    Update(e *Engine)
}

type IConstrained interface {
    Constrain(e *Engine) []RectSide
    Constrained(s RectSide)
}

type Entity struct {
    Visible    bool
    Bounds     Rect
    Velocity   Vector
}

func (this *Entity) CollidesWith(other *Entity) bool {
    return this.Bounds.IntersectsWith(other.Bounds)
}

func (this *Entity) MoveTo(pos Vector) {
    this.Bounds.Vector = pos
}

func (this *Entity) Push(distance Vector) {
    this.Bounds.Vector = this.Bounds.Vector.Plus(distance)
}

func (this *Entity) FixedUpdate() {
    this.Push(this.Velocity)
}