package main

import (
	"github.com/nathan-fiscaletti/go2d"
)

type TestEntitiesScene struct{}

func (this *TestEntitiesScene) Initialize(engine *go2d.Engine, scene *go2d.Scene) {
	// Create Draggable Entity
	circle := go2d.NewCircleImageEntity("#00FF00", 64)
	draggable := &Draggable{
		ImageEntity: circle,
	}
	draggable.MoveTo(
		go2d.Vector{
			X: (engine.Bounds().Width / 2) - 32,
			Y: (engine.Bounds().Height / 2) - 32,
		},
	)
	draggable.originalPos = draggable.Bounds.Vector
	scene.AddNamedEntity("draggable", 1, draggable)

	// Create Text Entity
	desc := go2d.NewTextEntitySimple("You can drag this around and it will snap back when you let go")
	desc.Measure(engine) // measure the text so that it's Bounds get filled in
	desc.MoveTo(
		go2d.Vector{
			X: (engine.Bounds().Width / 2) - (desc.Bounds.Width / 2),
			Y: (engine.Bounds().Height / 2) - (desc.Bounds.Height / 2) + 32 + 32,
		},
	)
	scene.AddNamedEntity("desc", 1, desc)
}

func main() {
	// Set default font
	go2d.SetDefaultFont("../test_resources/font.ttf", 16, "#fff")

	// Create a new Engine
	engine := go2d.NewEngine(
		"My Engine",

		// Set the engine to 16x9 aspect ratio with 1200 width.
		go2d.NewAspectRatio(
			16, 9, go2d.AspectRatioControlAxisWidth,
		).NewDimensions(1200),
	)

	// Create the new Scene
	scene := go2d.NewScene(engine, "Main Scene")
	scene.Initializer = &TestEntitiesScene{}

	engine.SetScene(&scene)
	engine.Run()
}

// Custom Draggable entity
type Draggable struct {
	// embed an ImageEntity for display
	*go2d.ImageEntity

	// when set to true, this entity will follow the cursor
	followCursor bool

	// The original position for the entity that it should snap back
	// to after being dragged and then let go.
	originalPos go2d.Vector
}

// MouseMove event for the Draggable entity
func (this *Draggable) MouseMove(pos go2d.Vector) {
	if this.followCursor {
		this.MoveTo(go2d.Vector{
			X: pos.X - (this.Bounds.Width / 2),
			Y: pos.Y - (this.Bounds.Height / 2),
		})
	}
}

// Mouse down event for the draggable entity
func (this *Draggable) MouseDown(button int, pos go2d.Vector) {
	if this.Bounds.Contains(pos) {
		this.followCursor = true
	}
}

// Mouse up event for the draggable entity
func (this *Draggable) MouseUp(button int, pos go2d.Vector) {
	this.followCursor = false
	this.MoveTo(this.originalPos)
}

func (this *Draggable) Update(engine *go2d.Engine) {
	this.Entity.Update()
}
