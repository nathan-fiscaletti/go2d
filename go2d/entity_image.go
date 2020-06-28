package go2d

import(
    "image"

    "github.com/tfriedel6/canvas"
)

type ImageEntity struct {
    Entity

    gImg image.Image

    canvas *canvas.Canvas
    cImg *canvas.Image
}

func NewImageEntity(c *canvas.Canvas, img image.Image) *ImageEntity {
    i,err := c.LoadImage(img)
    if err != nil {
        panic(err)
    }
    return &ImageEntity {
        gImg: img,
        cImg: i,
        canvas: c,
        Entity: Entity {
            Visible: true,
            Bounds: Rect {
                Dimensions: Dimensions {
                    Width: img.Bounds().Dx(),
                    Height: img.Bounds().Dy(),
                },
            },
        },
    }
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
