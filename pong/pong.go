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
    this.Measure(engine)
    this.MoveTo(go2d.Vector{
        X: (engine.Bounds().Width / 2) - (this.Bounds.Width / 2),
        Y: (engine.Bounds().Height / 2) - (this.Bounds.Height / 2),
    })
}

type Paddle struct {
    *go2d.ImageEntity
    aiControlled bool
    aiSpeed      int
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

type Ball struct {
    *go2d.ImageEntity
    aiSpeed int
    engine *go2d.Engine
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
            Y: rand.Intn(10 - -10) + -10,
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
        this.engine.GetScene().SetResource("cpuScore", this.engine.GetScene().GetResource("cpuScore").(int) + 1)
        this.engine.GetScene().SetResource("score", fmt.Sprintf("%v - %v", this.engine.GetScene().GetResource("playerScore").(int), this.engine.GetScene().GetResource("cpuScore").(int)))
        this.Respawn(this.engine, go2d.DirectionLeft())
    } else if r == go2d.RectSideRight {
        this.engine.GetScene().SetResource("playerScore", this.engine.GetScene().GetResource("playerScore").(int) + 1)
        this.engine.GetScene().SetResource("score", fmt.Sprintf("%v - %v", this.engine.GetScene().GetResource("playerScore").(int), this.engine.GetScene().GetResource("cpuScore").(int)))
        this.Respawn(this.engine, go2d.DirectionRight())
    }
}

func (this *Ball) Respawn(engine *go2d.Engine, direction go2d.Vector) {
    this.direction = direction
    this.MoveTo(engine.Bounds().Center())
    this.Velocity = go2d.Vector{
        X: this.direction.X * this.aiSpeed,
        Y: rand.Intn(10 - -10) + -10,
    }
}

func main() {
    engine := go2d.NewEngine(
        "Pong",
        go2d.NewAspectRatio(
            16, 9, go2d.AspectRatioControlAxisWidth,
        ).NewDimensions(1200),
    )

    // Create the new Scene
    scene := go2d.NewScene(engine, "Level 1")
    scene.LoadResources = func(engine *go2d.Engine, scene *go2d.Scene) {
        // Set default font
        go2d.SetDefaultFont("../test_resources/font.ttf", 24, "#fff")

        score := go2d.NewTextEntitySimple("0 - 0")
        score.Measure(engine)
        score.MoveTo(go2d.Vector{
            X: (engine.Bounds().Width / 2) - (score.Bounds.Width / 2),
            Y: (engine.Bounds().Height / 2) - (score.Bounds.Height / 2),
        })
        scoreDisplay := &ScoreDisplay{
            TextEntity: score,
        }
        scene.SetResource("score", "0 - 0")
        scene.SetResource("playerScore", 0)
        scene.SetResource("cpuScore", 0)
        scene.AddNamedEntity("score", 3, scoreDisplay)

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
            engine: engine,
        }

        ball.Respawn(engine, go2d.DirectionRight())

        scene.AddNamedEntity("ball", 2, ball)
    }

    engine.SetScene(&scene)
    engine.HideCursor = true
    engine.Run()
}