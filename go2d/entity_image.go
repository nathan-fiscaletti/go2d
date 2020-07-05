package go2d

import (
    "image"
    "os"
    "math"

    "github.com/tfriedel6/canvas"
    "github.com/tfriedel6/canvas/backend/softwarebackend"
)

type ITexture interface {
    GetTexture() image.Image
}

type ImageEntity struct {
    Entity

    gImg   image.Image
    cImg   *canvas.Image
}

func NewImageEntity(img image.Image) *ImageEntity {
    return &ImageEntity{
        gImg:   img,
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

func NewRectImageEntity(color string, dimensions Dimensions) *ImageEntity {
    backend := softwarebackend.New(int(dimensions.Width), int(dimensions.Height))
    cv := canvas.New(backend)

    cv.SetFillStyle(color)
    cv.Rect(0, 0, float64(dimensions.Width), float64(dimensions.Height))
    cv.Fill()
    
    img := cv.GetImageData(0, 0, int(dimensions.Width), int(dimensions.Height))

    return NewImageEntity(img)
}

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

func (this *ImageEntity) GetImage() image.Image {
    return this.gImg
}

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
            float64(this.Bounds.X),
            float64(this.Bounds.Y),
            float64(this.Bounds.Width),
            float64(this.Bounds.Height),
        )
    }
}

func (this *ImageEntity) Update(e *Engine) {
    this.Entity.Update()
}