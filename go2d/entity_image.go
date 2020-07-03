package go2d

import (
    "image"
    "os"

    "github.com/tfriedel6/canvas"
)

type ImageEntity struct {
    Entity

    gImg   image.Image
    canvas *canvas.Canvas
    cImg   *canvas.Image
}

func NewImageEntity(c *canvas.Canvas, img image.Image) *ImageEntity {
    i, err := c.LoadImage(img)
    if err != nil {
        panic(err)
    }
    return &ImageEntity{
        gImg:   img,
        cImg:   i,
        canvas: c,
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

func LoadImageEntity(c *canvas.Canvas, path string) (*ImageEntity, error) {
    imgf, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    i, _, err := image.Decode(imgf)
    if err != nil {
        return nil, err
    }

    return NewImageEntity(c, i), nil
}

func (this *ImageEntity) GetImage() image.Image {
    return this.gImg
}

func (this *ImageEntity) Render() {
    if this.Visible {
        this.canvas.DrawImage(
            this.cImg,
            float64(this.Bounds.X),
            float64(this.Bounds.Y),
            float64(this.Bounds.Width),
            float64(this.Bounds.Height),
        )
    }
}

func (this *ImageEntity) FixedUpdate() {
    this.Entity.FixedUpdate()
}
