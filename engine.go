package go2d

import (
	"fmt"
	"sync"
	"time"

	"github.com/tfriedel6/canvas"
	"github.com/tfriedel6/canvas/sdlcanvas"
	"github.com/veandco/go-sdl2/sdl"
)

type IFPSUpdateHandler interface {
	OnFPSUpdated(*Engine, int)
}

type ITPSUpdateHandler interface {
	OnTPSUpdated(*Engine, int)
}

// Engine is the main engine that handles the game loop and rendering.
type Engine struct {
	// Name is the name of the game window.
	Name string
	// Dimensions is the dimensions of the game window.
	Dimensions Dimensions
	// FPSUpdateHandler is the handler for when the FPS is updated.
	FPSUpdateHandler IFPSUpdateHandler
	// TPSUpdateHandler is the handler for when the TPS is updated.
	TPSUpdateHandler ITPSUpdateHandler
	// Canvas is the canvas that is used to render the game.
	Canvas *canvas.Canvas
	// HideCursor is a flag that determines if the cursor should be hidden.
	HideCursor bool
	// MaxTPS is the maximum TPS that the game should run at.
	MaxTPS int

	scene      *Scene
	window     *sdlcanvas.Window
	currentFps int
	currentTps int
	renderMux  sync.Mutex
}

var runningEngines []*Engine = []*Engine{}

// NewEngine creates a new engine with the given name and dimensions.
func NewEngine(name string, dimensions Dimensions) *Engine {
	engine := Engine{
		Name:       name,
		Dimensions: dimensions,
		MaxTPS:     60,
		renderMux:  sync.Mutex{},
	}

	sdl.SetHint(sdl.HINT_VIDEO_HIGHDPI_DISABLED, "1")
	wnd, cv, err := sdlcanvas.CreateWindow(int(engine.Dimensions.Width), int(engine.Dimensions.Height), engine.Name)
	if err != nil {
		panic(err)
	}

	wnd.Window.SetResizable(false)

	engine.window = wnd
	engine.Canvas = cv

	engine.window.KeyDown = func(scancode int, rn rune, name string) {
		engine.scene.notifyKeyDown(scancode, rn, name)
	}

	engine.window.KeyUp = func(scancode int, rn rune, name string) {
		engine.scene.notifyKeyUp(scancode, rn, name)
	}

	engine.window.MouseUp = func(b, x, y int) {
		engine.scene.notifyMouseUp(b, Vector{X: float64(x), Y: float64(y)})
	}

	engine.window.MouseDown = func(b, x, y int) {
		engine.scene.notifyMouseDown(b, Vector{X: float64(x), Y: float64(y)})
	}

	engine.window.MouseMove = func(x, y int) {
		engine.scene.notifyMouseMove(Vector{X: float64(x), Y: float64(y)})
	}

	return &engine
}

// Bounds returns the bounds of the game window.
func (this *Engine) Bounds() Rect {
	return Rect{
		Vector: Vector{
			X: 0, Y: 0,
		},
		Dimensions: this.Dimensions,
	}
}

// GetFPS returns the current FPS.
func (this *Engine) GetFPS() int {
	return this.currentFps
}

// GetTPS returns the current TPS.
func (this *Engine) GetTPS() int {
	return this.currentTps
}

// SetScene sets the current scene to the given scene.
func (this *Engine) SetScene(scene *Scene) {
	this.renderMux.Lock()

	if this.scene != nil {
		this.scene.ClearResources()
	}
	this.scene = scene
	if this.scene.Initializer != nil {
		this.scene.Initializer.Initialize(this, this.scene)
	}
	this.window.Window.SetTitle(fmt.Sprintf("%s - %s", this.Name, this.scene.Name))

	this.renderMux.Unlock()
}

// GetScene returns the current scene.
func (this *Engine) GetScene() *Scene {
	return this.scene
}

// Run starts the game loop.
func (this *Engine) Run() {
	if this.HideCursor == true {
		sdl.ShowCursor(0)
	}
	runningEngines = append(runningEngines, this)
	this.runUpdates()
	this.runGraphics()
}

// GetActiveEngine returns the active engine.
func GetActiveEngine() *Engine {
	runningEngineCount := len(runningEngines)

	if runningEngineCount < 1 {
		return nil
	}

	if runningEngineCount > 1 {
		panic("More than one running Engine, please use GetActiveEngines() instead.")
	}

	return runningEngines[0]
}

// GetActiveEngines returns all active engines.
func GetActiveEngines() []*Engine {
	return runningEngines
}

func (this *Engine) runGraphics() {
	defer this.window.Destroy()

	timeSliceOpened := time.Now()
	framesThisSecond := 0

	this.window.MainLoop(func() {
		this.render()
		framesThisSecond += 1

		if time.Since(timeSliceOpened) >= time.Second {
			this.currentFps = framesThisSecond
			timeSliceOpened = time.Now()
			framesThisSecond = 0

			if this.FPSUpdateHandler != nil {
				this.FPSUpdateHandler.OnFPSUpdated(this, this.currentFps)
			}
		}
	})
}

func (this *Engine) runUpdates() {
	go func() {
		now := time.Now()

		ticksThisSecond := 0
		timeSliceOpened := now
		lastTickCompletedAt := time.Now()

		minimumTickDuration := time.Second / time.Duration(60)

		for {
			now := time.Now()

			if time.Since(timeSliceOpened) >= time.Second {
				this.currentTps = ticksThisSecond
				ticksThisSecond = 0
				timeSliceOpened = now

				if this.TPSUpdateHandler != nil {
					this.TPSUpdateHandler.OnTPSUpdated(this, this.currentTps)
				}
			}

			this.update()

			lastTickCompletedAt = now
			ticksThisSecond += 1
			timeRemainingForTick := lastTickCompletedAt.Add(minimumTickDuration).Sub(now)
			time.Sleep(timeRemainingForTick)
		}
	}()
}

func (this *Engine) update() {
	this.scene.performUpdate(this)
}

func (this *Engine) render() {
	this.renderMux.Lock()

	w, h := float64(this.Canvas.Width()), float64(this.Canvas.Height())
	this.Canvas.SetFillStyle("#000")
	this.Canvas.FillRect(0, 0, w, h)
	this.scene.performRender(this)
	this.renderMux.Unlock()
}
