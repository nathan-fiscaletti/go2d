package go2d

import(
    "reflect" // :(
    "errors"
)

type Renderable interface {
    Render()
}

type FixedUpdatable interface {
    FixedUpdate()
}

type Updatable interface {
    Update()
}

type Entity struct {
    Visible    bool
    Bounds     Rect
    Constraint Rect
    Velocity   *Vector
    layer int

    OnConstrained func(s RectSide)
}

func entityForInterface(iface interface{}) (Entity, error) {
    if _,isEntity := iface.(Entity); isEntity {
        return iface.(Entity), nil
    }

    r := reflect.ValueOf(iface)
    if r.Kind() == reflect.Ptr {
        r = reflect.Indirect(r)
    }

    for i := 0; i < r.NumField(); i++ {
        f := r.Field(i)
        if f.Kind() == reflect.Struct {
            if f.Type() == reflect.TypeOf(Entity{}) {
                return f.Interface().(Entity), nil
            }
        }
    }

    return Entity{}, errors.New("not an instance of Entity nor does it embed Entity")
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

func (this *Entity) FixedUpdate() {
    if this.Velocity != nil {
        this.Push(*this.Velocity)
    }
    this.constrain(this.Constraint)
}

func (this *Entity) constrain(r Rect) {
    out := this.Bounds.Constrain(r)
    if this.OnConstrained != nil {
        for _,side := range out {
            this.OnConstrained(side)
        }
    }
}