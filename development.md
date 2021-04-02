# Go2D - Developer Documentation

## Requirements

1. SDL

   You must install SDL in order to create games in Go2D. See [Installing SDL](./sdl.md) for more information on installing SDL for your particular platform.

2. The Go Programming Language

   Go2D is written in the Go programming language. See [Go: Download and Install](https://golang.org/doc/install) for more information on installing the Go programming language for your particular platform.

3. Go Libraries

   There are several Go libraries that are required to develop Go2D games. You can run the following commands to install these libraries.

   ```sh
   $ go get github.com/tfriedel6/canvas
   $ go get github.com/tfriedel6/canvas/sdlcanvas
   $ go get github.com/tfriedel6/canvas/backend/softwarebackend
   $ go get github.com/veandco/go-sdl2/sdl
   ```

## Developing in Go2D

> Note: This documentation is a work-in-progress

## Engines

One of the core components of Go2D is an Engine. Engines manage the core Canvas for your game, the window in which the canvas is placed as well as rendering and tick events.

You can create an engine using the following code.

```go
engine := go2d.NewEngine(
    "My Engine",

    // Set the engine to 16x9 aspect ratio with 1200 width.
    go2d.NewAspectRatio(
        16, 9, go2d.AspectRatioControlAxisWidth,
    ).NewDimensions(1200),
)
```

> See [`go2d.NewEngine()`]()

This will create a new Engine with a `16:9` aspect ratio, controlled by the Width of the Window, setting the Window to a `1200` width. This means the resulting window will be `1200x675`. We'll go into the Aspect Ratio helper functions, dimensions and other helpful data structures later on in this documentation.

Once you've created your Engine, you can access the raw SDL canvas it uses for displaying scenes as well as some configuration options within it's properties. The full `Engine` structure will list these properties.

Once you have your engine set up, you need to add a Scene to it.

## Scenes

Scenes are a way in which you can display content in the game. They are a handy way to manage all of the active Entities within your game at any given time. You can change the Scene that your engine is displaying using the [`engine.SetScene(...)`]() function.

To create a new Scene, use the [`go2d.NewScene(...)`]() function. After you've created and configured your Scene, you can set it as the active Scene for the Engine.

```go
myScene := go2d.NewScene(engine, "My Scene")

// .. add entities / configure scene ..

engine.SetScene(myScene)
```

* You can manage the Entities on your Scene using the functions provided by the EntityGroup functions attached to the Scene structure. We'll go into more detail regarding EntityGroups later in this documentation.

* You can manage the Timers attached to your Scene using the [`myScene.AddTimer(...)`]() and [`myScene.RemoveTimer(...)`]() functions. We'll go into more detail regarding Timers later in this documentation.

* You can manage the Resources required by your Scene using the [`myScene.SetResource(...)`](), [`myScene.GetResource(...)`]() and [`myScene.ClearResources()`]() functions. We'll go into more detail regarding Resources later in this documentation.

> Hint: You can tell your scene to render the FPS on the screen using the [`myScene.RenderFPS(...)`]() function. Likewise, you can stop rendering the FPS on your scene using the [`myScene.StopRenderingFPS()`]() function.

### Initializing your Scene

You can initialize all of your Scenes Resources, Entities and Timers using the [`Scene.Initialize`]() closure property.

```go
myScene.Initialize = func(engine *go2d.Engine, scene *go2d.Scene) {
    // initialize entities, resources and timers for the scene
    // and add them here.
}
```

### Pre-rendering

You can opt to run some code before each frame is rendered using the [`Scene.PreRender`]() closer property.

```go
myScene.PreRender = func(engine *go2d.Engine, scene *go2d.Scene) {
    // run some pre-render code before a frame is rendered
}
```

### Running code on each tick

You can run code on each game tick using the [`Scene.Update`]() closer property.

```go
myScene.Update = func(engine *go2d.Engine, scene *go2d.Scene) {
    // run some code on each game tick
}
```

> It's recommended that for updating positioning of entities you use the Update method for the entity itself instead of using the Scenes update method.

### Running custom Render code

By default the Scene will render out any entities in it's EntityGroup. However, if you want to add additional Render code you can make use of the [`Scene.Render`]() closer property.

```go
myScene.Render = func(engine *go2d.Engine, scene *go2d.Scene) {
    // run some custom render code
}
```

## Entities

An entity is any given rendered object within a Scene. These each have their own Update method.

There are several built in base Entity types you can choose from when creating an entity. These include

* Image Entities
* Line Entities
* Text Entities

Included in the Image Entities are basic Shape Entities.

You should define your own Entity structure in order to create an Entity.

### Implementing one of the Base Entity types

This is an example of a custom Entity that is meant to represent a ball.

```go
type Ball struct {
    *go2d.ImageEntity
}
```

By default, all Entities have an [`Entity.Velocity`]() property which can be used to move the ball on each tick. So, if we want to move the Ball one pixel to the right on each frame, we would simply update it's Velocity.

```go
// Create a new Image Entity for the Ball as a White image
// with a 32px radius.
ballImageEntity := go2d.NewCircleImageEntity("#FFFFFF", 32)

// Create the ball
ball := &Ball{
    ImageEntity: ballImageEntity,
}

// Set it's velocity
ball.Velocity = go2d.NewVelocityVector(
    1,                 // Move 1 pixel to the right on the X axis
    0,                 // Don't move up or down,
    go2d.TICK_DURATION // Move once per tick
)
```

### Collision

You can add collision to your entity by implementing the [`go2d.ICollisionDetection`]() and [`go2d.ICollidable`]() interfaces.

The [`go2d.ICollidable`]() interface will allow other Entities that implement the [`go2d.ICollisionDetection`]() interface to be notified when they collide with this Entity.

Likewise, the [`go2d.ICollisionDetection`]() interface is used to notify an Entity that it has collided with another Entity. 

> The [`go2d.ICollisionDetection`]() interface requires that the Entity it is applied to also implements the [`go2d.ICollidable`]() interface.

```go
// ICollidable
func (ball *Ball) GetCollider() Rect {
    return ball.Bounds
}

// ICollisionDetection
func (ball *Ball) CollidedWith(other interface{}) {
    // handle the collision
}
```

### Constraints

You can confine an Entity to a given rect by implementing the [`go2d.IConstrain`]() and [`go2d.IConstrained`]() interfaces. This will stop the Entities Velocity from allowing it to Leave the provided Rect. This is useful if you want to keep an Entity from leaving a specific area due to it's Velocity or other factors.

**IConstrain**

You can implement the [`go2d.IConstrain`]() interface to stop an entity from leaving the desired rect.

This function should return an array of sides that the entity was constrained to if any. These values would include one or more of [`go2d.RectSideLeft`](), [`go2d.RectSideRight`](), [`go2d.RectSideTop`]() or [`go2d.RectSideBottom`](). If the array is empty, it indicates that the entity was not constrained.

You should facilitate this by constraining the Bounds of the entity to a specified Rect. The related function automatically returns the RectSides array.

In this instance, we will constrain the ball to the bounds of the Engine itself.

```go
func (ball *Ball) Constrain(engine *go2d.Engine) []go2d.RectSide {
    return ball.Bounds.Constrain(engine.Bounds())
}
```

**IConstrained**

You can implement the [`go2d.IConstrained`]() interface to be notified when the Entity attempted to leave those constraints.

> The [`go2d.IConstrained`]() interface requires that the Entity it is applied to also implement the [`go2d.IConstrain`]() interface.

In this instance, when the Ball is constrained we will reverse it's X Velocity which will keep the ball bouncing between the left and right sides of the Scene.

```go
func (ball *Ball) Constrained(r go2d.RectSide) {
    if r == go2d.RectSideTop || r == go2d.RectSideBottom {
        ball.Velocity.X = -ball.Velocity.X
    }
}
```