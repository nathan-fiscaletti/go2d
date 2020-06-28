package main

// brew install sdl2{,_image,_mixer,_ttf,_gfx} pkg-config
// go get https://github.com/tfriedel6/canvas/sdlcanvas

// "github.com/tfriedel6/canvas/sdlcanvas"

import (
    "os"
    "image"

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
        // Load the image from a file
        imgf, err := os.Open("/Users/nathanf/Pictures/jennyandi.jpeg")
        if err != nil {
            panic(err)
        }
        i, _, err := image.Decode(imgf)
        if err != nil {
            panic(err)
        }

        // Create an ImageDrawable with the image.Image object
        im := go2d.NewImageDrawable(engine.Canvas, i)

        // Update the images velocity to move one pixel per tick.
        // TODO: Update velocity to be based on PixelsPerSecond
        im.Velocity = &go2d.Vector{X: 1, Y: 0}

        // Save the Resource in the Scene Resources
        scene.SetResource("img", im)
    }

    // Set the Render Callback for the Scene
    scene.Render = func(engine *go2d.Engine, scene *go2d.Scene) {
        // Retrieve the Image from the Scene Resources
        i := scene.GetResource("img").(*go2d.ImageDrawable)

        // Render it to the SCene
        i.Render()
    }

    // Set the Fixed Update callback for the scene
    scene.FixedUpdate = func(engine *go2d.Engine, scene *go2d.Scene) {
        // Retrieve the Image from the Scene Resources
        im := scene.GetResource("img").(*go2d.ImageDrawable)

        // Run the FixedUpdate for the Image
        im.FixedUpdate()
    }
    
    // Set the Scene for the Engine
    engine.SetScene(&scene)

    // Run the Engine
	engine.Run()
}
