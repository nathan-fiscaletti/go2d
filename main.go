package main

// brew install sdl2{,_image,_mixer,_ttf,_gfx} pkg-config
// go get https://github.com/tfriedel6/canvas/sdlcanvas

// "github.com/tfriedel6/canvas/sdlcanvas"

import (
    "./go2d"
)

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
        blue.Velocity = &go2d.Vector{X: 10, Y: 0}
        red.Velocity = &go2d.Vector{X: 0, Y: 10}
        
        blue.Entity.Constraint = engine.Bounds()
        red.Entity.Constraint = engine.Bounds()

        // Save the Resource in the Scene Resources
        //scene.SetResource("img", im)
        scene.AddNamedEntity("blue", 1, blue)
        scene.AddNamedEntity("red", 2, red)
    }

    // Set the Scene for the Engine
    engine.SetScene(&scene)

    // Run the Engine
    engine.Run()
}
