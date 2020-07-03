package go2d


type RectSide int
const (
    RectSideLeft = iota
    RectSideRight
    RectSideTop
    RectSideBottom
)

type Rect struct {
    Vector
    Dimensions
}

func NewZeroRect(w int, h int) Rect {
    return NewRect(0, 0, w, h)
}

func NewRect(x int, y int, w int, h int) Rect {
    return Rect {
        Vector: Vector {
            X: x,
            Y: y,
        },
        Dimensions: Dimensions {
            Width: w,
            Height: h,
        },
    }
}

func (this Rect) Equals(other Rect) bool {
    return this.X == other.X && this.Y == other.Y && 
           this.Width == other.Width && this.Height == other.Height
}

func (this Rect) IntersectsWith(other Rect) bool {
    return this.X < other.X + other.Width &&
           this.X + this.Width > other.X &&
           this.Y < other.Y + other.Height &&
           this.Y + this.Height > other.Y
}

func (this Rect) Contains(v Vector) bool {
    return !(v.X < this.X ||
        v.Y < this.Y ||
        v.X > this.X+this.Width ||
        v.Y > this.Y+this.Height)
}

func (this Rect) Center() Vector {
    return Vector {
        X: this.X + this.Width / 2,
        Y: this.Y + this.Height / 2,
    }
}

func (this *Rect) Constrain(r Rect) []RectSide {
    res := []RectSide{}

    if this.X < r.X {
        this.X = r.X
        res = append(res, RectSideLeft)
    }

    if this.X + this.Width > r.X + r.Width {
        this.X = r.X + r.Width - this.Width
        res = append(res, RectSideRight)
    }

    if this.Y < r.Y {
        this.Y = r.Y
        res = append(res, RectSideTop)
    }

    if this.Y + this.Height > r.Y + r.Height {
        this.Y = r.Y + r.Height - this.Height
        res = append(res, RectSideBottom)
    }

    return res
}