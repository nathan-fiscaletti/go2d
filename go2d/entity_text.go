package go2d

type TextEntity struct {
    Entity

    textColor     string
    text          string
    font          string
    fontSize      float64

    centeredIn    Rect
    textCentering TextCentering

    isMeasured    bool
}

type TextCentering int
const (
    TEXT_CENTERING_NONE = iota
    TEXT_CENTERING_VERTICAL
    TEXT_CENTERING_HORIZONTAL
)

var defaultFont string
var defaultFontSize float64
var defaultTextColor string

func NewTextEntity(text string, font string, fontSize float64, textColor string) *TextEntity {
    return &TextEntity{
        text:       text,
        font:       font,
        fontSize:   fontSize,
        textColor:  textColor,
        isMeasured: false,
        Entity:     Entity {
                        Bounds: Rect{
                            Vector: Vector{
                                X:0,
                                Y:0,
                            },
                        },
                        Visible: true,
                    },
    }
}

func NewTextEntitySimple(text string) *TextEntity {
    return &TextEntity{
        text:       text,
        font:       defaultFont,
        fontSize:   defaultFontSize,
        textColor:  defaultTextColor,
        isMeasured: false,
        Entity:     Entity {
                        Bounds: Rect{
                            Vector: Vector{
                                X:0,
                                Y:0,
                            },
                        },
                        Visible: true,
                    },
    }
}

func SetDefaultFont(font string, fontSize float64, textColor string) {
    defaultFont = font
    defaultFontSize = fontSize
    defaultTextColor = textColor
}

func (this *TextEntity) SetCenteredIn(r Rect) {
    this.centeredIn = r
}

func (this *TextEntity) SetCentering(c TextCentering) {
    this.textCentering = c
}

func (this *TextEntity) SetText(text string) {
    this.text = text
    this.isMeasured = false
}

func (this *TextEntity) SetFontSize(fontSize float64) {
    this.fontSize = fontSize
    this.isMeasured = false
}

func (this *TextEntity) SetFont(font string) {
    this.font = font
    this.isMeasured = false
}

func (this *TextEntity) SetTextColor(color string) {
    this.textColor = color
}

func (this *TextEntity) Measure(e *Engine) {
    e.Canvas.SetFont(this.font, float64(this.fontSize))
    tm := e.Canvas.MeasureText(this.text)
    this.Bounds.Dimensions = Dimensions{
        Width: tm.Width,
        Height: tm.ActualBoundingBoxAscent + tm.ActualBoundingBoxDescent,
    }

    centeredIn := this.centeredIn
    if this.centeredIn.IsZero() {
        centeredIn = e.Bounds()
    }

    if this.textCentering & TEXT_CENTERING_VERTICAL != 0 {
        this.Bounds.Vector.Y = centeredIn.Height / 2 - this.Bounds.Height / 2
    } 
    
    if this.textCentering & TEXT_CENTERING_HORIZONTAL != 0 {
        this.Bounds.Vector.X = centeredIn.Width / 2 - this.Bounds.Width / 2
    }

    this.isMeasured = true
}

func (this *TextEntity) Render(e *Engine) {
    if !this.isMeasured {
        this.Measure(e)
    }

    e.Canvas.SetFont(this.font, this.fontSize)
    e.Canvas.SetFillStyle(this.textColor)
    e.Canvas.FillText(this.text, this.Bounds.X, this.Bounds.Y + this.Bounds.Height)
}

func (this *TextEntity) Update(e *Engine) {
    this.Entity.Update()
}