package go2d

import (
    "image"
    "os"

    "github.com/tfriedel6/canvas"
)

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
                    Width:  img.Bounds().Dx(),
                    Height: img.Bounds().Dy(),
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

func (this *ImageEntity) FixedUpdate(e *Engine) {
    this.Entity.FixedUpdate()
}
