package main

// brew install sdl2{,_image,_mixer,_ttf,_gfx} pkg-config
// go get https://github.com/tfriedel6/canvas/sdlcanvas

// "github.com/tfriedel6/canvas/sdlcanvas"

import (
    "fmt"

    "./go2d"
)

type Spaceship struct {
    *go2d.ImageEntity
}

func (this *Spaceship) KeyDown(scancode int, rn rune, name string) {
    if name == "ArrowDown" {
        this.Velocity = go2d.Vector{X:this.Velocity.X, Y:1}
    } else if name == "ArrowUp" {
        this.Velocity = go2d.Vector{X:this.Velocity.X, Y:-1}
    } else if name == "ArrowLeft" {
        this.Velocity = go2d.Vector{X:-1, Y:this.Velocity.Y}
    } else if name == "ArrowRight" {
        this.Velocity = go2d.Vector{X:1, Y:this.Velocity.Y}
    }
}

func (this *Spaceship) KeyUp(scancode int, rn rune, name string) {
    if name == "ArrowDown" || name == "ArrowUp" {
        this.Velocity = go2d.Vector{X:this.Velocity.X, Y:0}
    } else if name == "ArrowLeft" || name == "ArrowRight"{
        this.Velocity = go2d.Vector{X:0, Y:this.Velocity.Y}
    }
}

func (this *Spaceship) MouseDown(button int, pos go2d.Vector) {
    if button == go2d.MOUSE_BUTTON_LEFT {
        if this.Bounds.Contains(pos) {
            fmt.Printf("clicked on spaceship")
        }
    }
}

func main() {

    // Create a new Engine
    engine := go2d.NewEngine(
        "My Engine",

        // Set the engine to 16x9 aspect ratio with 1200 width.
        go2d.NewAspectRatio(
            16, 9, go2d.AspectRatioControlAxisWidth,
        ).NewDimensions(1200),
    )

    // Set the Tick Rate of the Engine to ~60hz
    engine.TickHz = 60

    // Create the new Scene
    scene := go2d.NewScene(engine, "Main Scene")

    // Set the LoadResources callback for the Scene
    scene.LoadResources = func(engine *go2d.Engine, scene *go2d.Scene) {
        blue, err := go2d.LoadImageEntity("./test_resources/blue.png")
        if err != nil {
            panic(err)
        }
        red, err := go2d.LoadImageEntity("./test_resources/red.png")
        if err != nil {
            panic(err)
        }

        // Update the images velocity to move one pixel per tick.
        // TODO: Update velocity to be based on PixelsPerSecond
        blue.Velocity = go2d.Vector{X: 10, Y: 0}
        red.Velocity = go2d.Vector{X: 0, Y: 10}
        
        blue.Entity.Constraint = engine.Bounds()
        red.Entity.Constraint = engine.Bounds()

        //scene.AddNamedEntity("blue", 1, blue)
        //scene.AddNamedEntity("red", 2, red)

        go2d.SetDefaultFont("./test_resources/font.ttf", 40, "#fff")

        fps := go2d.NewTextEntitySimple("FPS: 0")
        fps.Measure(engine)
        scene.AddNamedEntity("fps", 3, fps)

        tps := go2d.NewTextEntitySimple("TPS: 0")
        tps.Bounds.X = fps.Bounds.X + fps.Bounds.Width + 20
        scene.AddNamedEntity("tps", 3, tps)

        spriteSheet,err := go2d.NewSpriteSheet("./test_resources/sprite_sheet.png", 32, 32)
        if err != nil {
            panic(err)
        }
        spaceShipImg := spriteSheet.GetSprite(go2d.Vector{X:0, Y:0})

        spaceShip := go2d.NewImageEntity(spaceShipImg)
        spaceShip.MoveTo(go2d.Vector{
            X: 0,
            Y: 0,
        })

        spaceShipFinal := Spaceship{
            ImageEntity: spaceShip,
        }
        scene.AddNamedEntity("ss", 4, &spaceShipFinal)

        //circle := go2d.NewCircleImageEntity("#00FF00", 32)
        circle := go2d.NewRectImageEntity("#00FF00", go2d.Dimensions{Width: 32, Height:96})
        circle.MoveTo(go2d.Vector{
            X:32, Y: 32,
        })

        scene.AddNamedEntity("circle", 5, circle)
    }

    // Set the Scene for the Engine
    engine.SetScene(&scene)

    engine.OnFPSUpdated = func(engine *go2d.Engine, fps int) {
        fpsDisplay := engine.GetScene().GetEntity(3, "fps").(*go2d.TextEntity)
        fpsDisplay.SetText(fmt.Sprintf("FPS: %v", fps))

        if fps < 60 {
            fpsDisplay.SetTextColor("#ff0000")
        } else {
            fpsDisplay.SetTextColor("#fff")
        }

        tpsDisplay := engine.GetScene().GetEntity(3, "tps").(*go2d.TextEntity)
        tpsDisplay.Bounds.X = fpsDisplay.Bounds.X + fpsDisplay.Bounds.Width + 20
    }

    engine.OnTickRateUpdated = func(engine *go2d.Engine, tps int) {
        tpsDisplay := engine.GetScene().GetEntity(3, "tps").(*go2d.TextEntity)
        if tps < engine.TickHz {
            tpsDisplay.SetTextColor("#ff0000")
        } else {
            tpsDisplay.SetTextColor("#fff")
        }
        tpsDisplay.SetText(fmt.Sprintf("TPS: %v", tps))
    }

    // Run the Engine
    engine.Run()
}
