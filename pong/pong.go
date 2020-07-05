package main

import(
    "fmt"
    "../go2d"
    "math/rand"
)

type ScoreDisplay struct {
    *go2d.TextEntity
}

func (this *ScoreDisplay) Update(engine *go2d.Engine) {
    this.Entity.Update()
    this.SetText(engine.GetScene().GetResource("score").(string))
}

type Paddle struct {
    *go2d.ImageEntity
    aiControlled bool
    aiSpeed      float64
}

func (this *Paddle) MouseMove(pos go2d.Vector) {
    if !this.aiControlled {
        this.MoveTo(go2d.Vector{
            X: this.Bounds.X,
            Y: pos.Y - (this.Bounds.Height / 2),
        })
    }
}

func (this *Paddle) Update(engine *go2d.Engine) {
    this.Entity.Update()
    if this.aiControlled {
        ball := engine.GetScene().GetEntity(2, "ball").(*Ball)
        if ball.Bounds.Y < this.Bounds.Y + (this.Bounds.Height / 2) {
            this.Velocity = go2d.Vector{
                Y: -this.aiSpeed,
            }
        } else if ball.Bounds.Y > this.Bounds.Y + (this.Bounds.Height / 2) {
            this.Velocity = go2d.Vector{
                Y: this.aiSpeed,
            }
        }
    }
}

func (this *Paddle) Constrain(engine *go2d.Engine) []go2d.RectSide {
    return this.Bounds.Constrain(engine.Bounds())
}

type Ball struct {
    *go2d.ImageEntity
    aiSpeed float64
    direction go2d.Vector
}

func (this *Ball) Update(engine *go2d.Engine) {
    this.Entity.Update()
    player := engine.GetScene().GetEntity(1, "player").(*Paddle)
    cpu    := engine.GetScene().GetEntity(1, "cpu").(*Paddle)
    if this.Entity.CollidesWith(&player.Entity) || this.Entity.CollidesWith(&cpu.Entity) {
        this.direction = this.direction.InvertedX()
        this.Velocity = go2d.Vector {
            X: this.aiSpeed * this.direction.X,
            Y: -10 + rand.Float64() * (10 - -10),
        }
    }
}

func (this *Ball) Constrain(engine *go2d.Engine) []go2d.RectSide {
    return this.Bounds.Constrain(engine.Bounds())
}

func (this *Ball) Constrained(r go2d.RectSide) {
    if r == go2d.RectSideTop || r == go2d.RectSideBottom {
        this.Velocity.Y = -this.Velocity.Y
    } else if r == go2d.RectSideLeft {
        go2d.GetActiveScene().SetResource("cpuScore", go2d.GetActiveScene().GetResource("cpuScore").(int) + 1)
        go2d.GetActiveScene().SetResource("score", fmt.Sprintf("%v   %v", go2d.GetActiveScene().GetResource("playerScore").(int), go2d.GetActiveScene().GetResource("cpuScore").(int)))
        this.Respawn(go2d.GetActiveEngine(), go2d.DirectionLeft())
    } else if r == go2d.RectSideRight {
        go2d.GetActiveScene().SetResource("playerScore", go2d.GetActiveScene().GetResource("playerScore").(int) + 1)
        go2d.GetActiveScene().SetResource("score", fmt.Sprintf("%v   %v", go2d.GetActiveScene().GetResource("playerScore").(int), go2d.GetActiveScene().GetResource("cpuScore").(int)))
        this.Respawn(go2d.GetActiveEngine(), go2d.DirectionRight())
    }
}

func (this *Ball) Respawn(engine *go2d.Engine, direction go2d.Vector) {
    this.direction = direction
    this.MoveTo(engine.Bounds().Center())
    this.Velocity = go2d.Vector{
        X: this.direction.X * this.aiSpeed,
        Y: -10 + rand.Float64() * (10 - -10),
    }
}

func main() {
    engine := go2d.NewEngine(
        "Pong",
        go2d.NewAspectRatio(
            16, 9, go2d.AspectRatioControlAxisWidth,
        ).NewDimensions(1200),
    )

    scene := go2d.NewScene(engine, "Level 1")
    scene.LoadResources = func(engine *go2d.Engine, scene *go2d.Scene) {
        go2d.SetDefaultFont("../test_resources/font.ttf", 24, "#fff")

        score := go2d.NewTextEntitySimple("0   0")
        score.SetCentering(go2d.TEXT_CENTERING_VERTICAL | go2d.TEXT_CENTERING_HORIZONTAL)
        scoreDisplay := &ScoreDisplay{
            TextEntity: score,
        }
        scene.SetResource("score", "0   0")
        scene.SetResource("playerScore", 0)
        scene.SetResource("cpuScore", 0)
        scene.AddNamedEntity("score", 3, scoreDisplay)

        line := go2d.NewLineEntity(
            go2d.Vector{
                X: engine.Bounds().Width / 2,
            }, 
            go2d.DirectionDown(), 
            engine.Bounds().Height, 
            3, "#FFF",
        )
        scene.AddEntity(1, line)

        paddleImage := go2d.NewRectImageEntity("#FFFFFF", go2d.Dimensions{
            Width: 32,
            Height: 96,
        })

        paddle := &Paddle{
            ImageEntity: paddleImage,
            aiControlled: false,
        }

        paddle.MoveTo(go2d.Vector{
            X: 0,
            Y: 0,
        })

        scene.AddNamedEntity("player", 1, paddle)

        paddleImage2 := go2d.NewRectImageEntity("#FFFFFF", go2d.Dimensions{
            Width: 32,
            Height: 96,
        })

        cpupaddle := &Paddle{
            ImageEntity: paddleImage2,
            aiControlled: true,
            aiSpeed: 5,
        }

        cpupaddle.MoveTo(go2d.Vector{
            X: engine.Bounds().Width - cpupaddle.Bounds.Width,
            Y: 0,
        })

        scene.AddNamedEntity("cpu", 1, cpupaddle)

        ballImage := go2d.NewCircleImageEntity("#FFFFFF", 32)
        ball := &Ball{
            ImageEntity: ballImage,
            aiSpeed: 10,
        }

        ball.Respawn(engine, go2d.DirectionRight())

        scene.AddNamedEntity("ball", 2, ball)
    }

    engine.SetScene(&scene)
    engine.HideCursor = true
    engine.Run()
}