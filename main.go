package main

// brew install sdl2{,_image,_mixer,_ttf,_gfx} pkg-config
// go get https://github.com/tfriedel6/canvas/sdlcanvas

// "github.com/tfriedel6/canvas/sdlcanvas"

import (
    "fmt"

    "./go2d"
)

type Spaceship struct {
    go2d.Entity
    *go2d.ImageEntity //how
}

func (this *Spaceship) KeyDown(scancode int, rn rune, name string) {
    if name == "ArrowDown" {
        this.ImageEntity.Entity.Velocity = go2d.Vector{X:this.Entity.Velocity.X, Y:1}
    } else if name == "ArrowUp" {
        this.ImageEntity.Entity.Velocity = go2d.Vector{X:this.Entity.Velocity.X, Y:-1}
    } else if name == "ArrowLeft" {
        this.ImageEntity.Entity.Velocity = go2d.Vector{X:-1, Y:this.Entity.Velocity.Y}
    } else if name == "ArrowRight" {
        this.ImageEntity.Entity.Velocity = go2d.Vector{X:1, Y:this.Entity.Velocity.Y}
    }
}

func (this *Spaceship) KeyUp(scancode int, rn rune, name string) {
    if name == "ArrowDown" || name == "ArrowUp" {
        this.ImageEntity.Entity.Velocity = go2d.Vector{X:this.Entity.Velocity.X, Y:0}
    } else if name == "ArrowLeft" || name == "ArrowRight"{
        this.ImageEntity.Entity.Velocity = go2d.Vector{X:0, Y:this.Entity.Velocity.Y}
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
        blue, err := go2d.LoadImageEntity("/Users/nathan/Pictures/blue.png")
        if err != nil {
            panic(err)
        }
        red, err := go2d.LoadImageEntity("/Users/nathan/Pictures/red.png")
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

        go2d.SetDefaultFont("/Users/nathan/Downloads/Anonymous/Anonymous.ttf", 40, "#fff")

        fps := go2d.NewTextEntitySimple("FPS: 0")
        fps.Measure(engine)
        scene.AddNamedEntity("fps", 3, fps)

        tps := go2d.NewTextEntitySimple("TPS: 0")
        tps.Bounds.X = fps.Bounds.X + fps.Bounds.Width + 20
        scene.AddNamedEntity("tps", 3, tps)

        spriteSheet,err := go2d.NewSpriteSheet("/Users/nathan/Pictures/sprite_sheet.png", 32, 32)
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
