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
        us, err := go2d.LoadImageEntity(engine.Canvas, "/Users/nathanf/Pictures/jennyandi.jpeg")
        if err != nil {
            panic(err)
        }
        lg,err := go2d.LoadImageEntity(engine.Canvas, "/Users/nathanf/Pictures/vrazo_logo.png")
        if err != nil {
            panic(err)
        }

        // Update the images velocity to move one pixel per tick.
        // TODO: Update velocity to be based on PixelsPerSecond
        us.Velocity = &go2d.Vector{X: 1, Y: 0}
        lg.Velocity = &go2d.Vector{X: 0, Y: 1}

        // Save the Resource in the Scene Resources
        //scene.SetResource("img", im)
        scene.AddNamedEntity("img", 1, us)
        scene.AddNamedEntity("img2", 2, lg)
    }
    
    // Set the Scene for the Engine
    engine.SetScene(&scene)

    // Run the Engine
	engine.Run()
}
