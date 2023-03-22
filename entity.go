package go2d

// IEntityRenderer is an interface that can be implemented by entities that
// want to render themselves.
type IEntityRenderer interface {
	Render(c *Engine)
}

// IEntityUpdater is an interface that can be implemented by entities that
// want to update themselves.
type IEntityUpdater interface {
	Update(e *Engine)
}

// IEntityConstraint is an interface that can be implemented by entities
// that want to constrain themselves.
type IEntityConstraint interface {
	Constrain(e *Engine) []RectSide
}

// IEntityConstrainedHandler is an interface that can be implemented by
// entities that want to be notified when they are constrained.
type IEntityConstrainedHandler interface {
	OnConstrained(s RectSide)
}

// IEntityCollider is an interface that can be implemented by entities
// that want to be recognized by other entities with that implement
// IEntityCollisionDetection.
type IEntityCollider interface {
	GetCollider() Rect
}

// IEntityCollisionDetection is an interface that can be implemented by
// entities that want to be notified when they collide with other
// entities that implement IEntityCollider.
type IEntityCollisionDetection interface {
	CollidedWith(other interface{})
}

type IEntity interface {
	GetEntity() *Entity
}

// Entity is a basic entity that can be used to create more complex entities.
type Entity struct {
	// Visible is a flag that can be used to hide the entity.
	Visible bool
	// Bounds is the rectangle that defines the entity's position and size.
	Bounds Rect
	// Velocity is the entity's velocity.
	Velocity VelocityVector
}

// CollidesWith returns true if the entity collides with the other entity. This
// is not using IEntityCollider and IEntityCollisionDetection, but rather
// directly checks the entity's bounds.
func (this *Entity) CollidesWith(other *Entity) bool {
	return this.Bounds.IntersectsWith(other.Bounds)
}

// MoveTo moves the entity to the specified position instantly.
func (this *Entity) MoveTo(pos Vector) {
	this.Bounds.Vector = pos
}

// Push moves the entity by the specified distance instantly.
func (this *Entity) Push(distance Vector) {
	this.Bounds.Vector.X += distance.X
	this.Bounds.Vector.Y += distance.Y
}

// Update updates the entity's position based on its velocity.
func (this *Entity) Update() {
	this.Push(this.Velocity.GetNextMovement())
}
