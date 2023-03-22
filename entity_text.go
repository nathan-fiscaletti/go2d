package go2d

// TextEntity is a simple text entity that can be used to display text on the
// screen.
type TextEntity struct {
	Entity

	textColor string
	text      string
	font      string
	fontSize  float64

	centeredIn    Rect
	textCentering TextCentering

	isMeasured bool
}

// TextCentering is a type that represents the different ways text can be
// centered.
type TextCentering int

const (
	TEXT_CENTERING_NONE = iota
	TEXT_CENTERING_VERTICAL
	TEXT_CENTERING_HORIZONTAL
)

var defaultFont string
var defaultFontSize float64
var defaultTextColor string

// NewTextEntity creates a new text entity with the given text, font, font size,
// and text color.
func NewTextEntity(text string, font string, fontSize float64, textColor string) *TextEntity {
	return &TextEntity{
		text:       text,
		font:       font,
		fontSize:   fontSize,
		textColor:  textColor,
		isMeasured: false,
		Entity: Entity{
			Bounds: Rect{
				Vector: Vector{
					X: 0,
					Y: 0,
				},
			},
			Visible: true,
		},
	}
}

// NewTextEntitySimple creates a new text entity with the given text, using the
// default font, font size, and text color.
func NewTextEntitySimple(text string) *TextEntity {
	return &TextEntity{
		text:       text,
		font:       defaultFont,
		fontSize:   defaultFontSize,
		textColor:  defaultTextColor,
		isMeasured: false,
		Entity: Entity{
			Bounds: Rect{
				Vector: Vector{
					X: 0,
					Y: 0,
				},
			},
			Visible: true,
		},
	}
}

// SetDefaultFont sets the default font, font size, and text color for all text
// entities created with NewTextEntitySimple.
func SetDefaultFont(font string, fontSize float64, textColor string) {
	defaultFont = font
	defaultFontSize = fontSize
	defaultTextColor = textColor
}

// SetCenteredIn sets the rectangle that this text entity should be centered in.
func (this *TextEntity) SetCenteredIn(r Rect) {
	this.centeredIn = r
}

// SetCentering sets the centering of this text entity.
func (this *TextEntity) SetCentering(c TextCentering) {
	this.textCentering = c
}

// SetText sets the text of this text entity.
func (this *TextEntity) SetText(text string) {
	this.text = text
	this.isMeasured = false
}

// SetFontSize sets the font size of this text entity.
func (this *TextEntity) SetFontSize(fontSize float64) {
	this.fontSize = fontSize
	this.isMeasured = false
}

// SetFont sets the font of this text entity.
func (this *TextEntity) SetFont(font string) {
	this.font = font
	this.isMeasured = false
}

// SetTextColor sets the text color of this text entity.
func (this *TextEntity) SetTextColor(color string) {
	this.textColor = color
}

// Measure updates the bounds of this text entity based on the text, font, and
// font size also taking into account the centering of the text.
func (this *TextEntity) Measure(e *Engine) {
	e.Canvas.SetFont(this.font, float64(this.fontSize))
	tm := e.Canvas.MeasureText(this.text)
	this.Bounds.Dimensions = Dimensions{
		Width:  tm.Width,
		Height: tm.ActualBoundingBoxAscent + tm.ActualBoundingBoxDescent,
	}

	centeredIn := this.centeredIn
	if this.centeredIn.IsZero() {
		centeredIn = e.Bounds()
	}

	if this.textCentering&TEXT_CENTERING_VERTICAL != 0 {
		this.Bounds.Vector.Y = centeredIn.Height/2 - this.Bounds.Height/2
	}

	if this.textCentering&TEXT_CENTERING_HORIZONTAL != 0 {
		this.Bounds.Vector.X = centeredIn.Width/2 - this.Bounds.Width/2
	}

	this.isMeasured = true
}

// Render renders this text entity to the given engine.
func (this *TextEntity) Render(e *Engine) {
	if !this.isMeasured {
		this.Measure(e)
	}

	e.Canvas.SetFont(this.font, this.fontSize)
	e.Canvas.SetFillStyle(this.textColor)
	e.Canvas.FillText(this.text, this.Bounds.X, this.Bounds.Y+this.Bounds.Height)
}

// Update updates this text entity.
func (this *TextEntity) Update(e *Engine) {
	this.Entity.Update()
}

// GetEntity returns the base entity of this text entity.
func (this *TextEntity) GetEntity() *Entity {
	return &this.Entity
}
