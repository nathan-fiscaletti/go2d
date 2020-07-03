package main

import(
    "../go2d"
    "math/rand"
)

type Paddle struct {
    *go2d.ImageEntity
    aiControlled bool
    aiTargetBall bool
    aiSpeed      int
    aiMoveUp     bool
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

func (this *Paddle) Constrained(r go2d.RectSide) {
    if this.aiControlled {
        if r == go2d.RectSideBottom {
            this.aiMoveUp = true
        } else if r == go2d.RectSideTop {
            this.aiMoveUp = false
        }
    }
}

type Ball struct {
    *go2d.ImageEntity
    aiSpeed int
    engine *go2d.Engine
}

func (this *Ball) Update(engine *go2d.Engine) {
    player := engine.GetScene().GetEntity(1, "player").(*Paddle)
    cpu    := engine.GetScene().GetEntity(1, "cpu").(*Paddle)
    if this.Entity.CollidesWith(&player.Entity) || this.Entity.CollidesWith(&cpu.Entity) {
        this.Velocity = this.Velocity.PlusX(-(this.Velocity.X*2))
    }
}

func (this *Ball) Constrain(engine *go2d.Engine) []go2d.RectSide {
    return this.Bounds.Constrain(engine.Bounds())
}

func (this *Ball) Constrained(r go2d.RectSide) {
    if r == go2d.RectSideTop || r == go2d.RectSideBottom {
        this.Velocity = this.Velocity.PlusY(-(this.Velocity.Y*2))
    } else if r == go2d.RectSideLeft || r == go2d.RectSideRight {
        this.Respawn(this.engine)
    }
}

func (this *Ball) Respawn(engine *go2d.Engine) {
    this.MoveTo(engine.Bounds().Center())
    this.Velocity = go2d.Vector{
        X: this.aiSpeed,
        Y: rand.Intn(10 - -10) + -10,
    }
}

func main() {
    go2d.SetDefaultFont("../test_resources/font.ttf", 16, "#fff")
    engine := go2d.NewEngine(
        "Pong",
        go2d.NewAspectRatio(
            16, 9, go2d.AspectRatioControlAxisWidth,
        ).NewDimensions(1200),
    )

    // Create the new Scene
    scene := go2d.NewScene(engine, "Level 1")
    scene.LoadResources = func(engine *go2d.Engine, scene *go2d.Scene) {
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

        ball.Respawn(engine)

        scene.AddNamedEntity("ball", 2, ball)
    }

    engine.SetScene(&scene)
    engine.HideCursor = true
    engine.Run()
}