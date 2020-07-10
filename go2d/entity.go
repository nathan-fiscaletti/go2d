package go2d

type IRender interface {
    Render(c *Engine)
}

type IUpdate interface {
    Update(e *Engine)
}

type IConstrain interface {
    Constrain(e *Engine) []RectSide
}

type IConstrained interface {
    Constrained(s RectSide)
}

type Entity struct {
    Visible    bool
    Bounds     Rect
    Velocity   VelocityVector
}

func (this *Entity) CollidesWith(other *Entity) bool {
    return this.Bounds.IntersectsWith(other.Bounds)
}

func (this *Entity) MoveTo(pos Vector) {
    this.Bounds.Vector = pos
}

func (this *Entity) Push(distance Vector) {
    this.Bounds.Vector.X += distance.X
    this.Bounds.Vector.Y += distance.Y
}

func (this *Entity) Update() {
    this.Push(this.Velocity.GetNextMovement())
}