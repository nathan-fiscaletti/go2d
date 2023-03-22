package go2d

import (
	"image"
	"math"
	"os"

	"github.com/tfriedel6/canvas"
	"github.com/tfriedel6/canvas/backend/softwarebackend"
)

// ITexture is an interface that represents a texture.
type ITexture interface {
	GetTexture() image.Image
}

// ImageEntity is a simple image entity that can be used to draw images on the screen.
type ImageEntity struct {
	Entity

	gImg image.Image
	cImg *canvas.Image
}

// NewImageEntity creates a new image entity from the given image.
func NewImageEntity(img image.Image) *ImageEntity {
	return &ImageEntity{
		gImg: img,
		Entity: Entity{
			Visible: true,
			Bounds: Rect{
				Dimensions: Dimensions{
					Width:  float64(img.Bounds().Dx()),
					Height: float64(img.Bounds().Dy()),
				},
			},
		},
	}
}

// LoadImageEntity loads an image entity from the given path.
func LoadImageEntity(path string) (*ImageEntity, error) {
	imgf, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	i, _, err := image.Decode(imgf)
	if err != nil {
		return nil, err
	}

	return NewImageEntity(i), nil
}

// NewRectImageEntity creates a new image entity with the given color and dimensions.
func NewRectImageEntity(color string, dimensions Dimensions) *ImageEntity {
	backend := softwarebackend.New(int(dimensions.Width), int(dimensions.Height))
	cv := canvas.New(backend)

	cv.SetFillStyle(color)
	cv.Rect(0, 0, dimensions.Width, dimensions.Height)
	cv.Fill()

	img := cv.GetImageData(0, 0, int(dimensions.Width), int(dimensions.Height))

	return NewImageEntity(img)
}

// NewCircleImageEntity creates a new image entity depicting a circle with the given color and radius.
func NewCircleImageEntity(color string, radius int) *ImageEntity {
	backend := softwarebackend.New(radius*2, radius*2)
	cv := canvas.New(backend)

	cv.SetFillStyle(color)
	cv.BeginPath()
	cv.Arc(float64(radius/2), float64(radius/2), float64(radius/2), 0, math.Pi*2, false)
	cv.Fill()

	img := cv.GetImageData(0, 0, radius, radius)

	return NewImageEntity(img)
}

// GetImage returns the image of the image entity.
func (this *ImageEntity) GetImage() image.Image {
	return this.gImg
}

// Render renders the image entity.
func (this *ImageEntity) Render(e *Engine) {
	if this.cImg == nil {
		i, err := e.Canvas.LoadImage(this.gImg)
		if err != nil {
			panic(err)
		}
		this.cImg = i
	}

	if this.Visible {
		e.Canvas.DrawImage(
			this.cImg,
			this.Bounds.X,
			this.Bounds.Y,
			this.Bounds.Width,
			this.Bounds.Height,
		)
	}
}

// Update updates the image entity.
func (this *ImageEntity) Update(e *Engine) {
	this.Entity.Update()
}

// GetEntity returns the entity of the image entity.
func (this *ImageEntity) GetEntity() *Entity {
	return &this.Entity
}
