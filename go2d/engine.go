package go2d

import (
    "time"
    "sync"

    "./metrics"
    //"./graphics"

	"github.com/tfriedel6/canvas"
	"github.com/tfriedel6/canvas/sdlcanvas"
)

type Engine struct {
    Title             string
	MaxTPS            int
    Dimensions        metrics.Dimensions
    OnTickRateUpdated func(int)
    OnFPSUpdated      func(int)
    Canvas            *canvas.Canvas

	scene             *Scene
    window            *sdlcanvas.Window
    currentHz         int
    currentFps        int
    renderMux         sync.Mutex
    tickMux           sync.Mutex
}

func NewEngine(title string, dimensions metrics.Dimensions) *Engine {
    engine := Engine{
        Title: title,
        Dimensions: dimensions,
        MaxTPS: 60,
        renderMux: sync.Mutex{},
        tickMux: sync.Mutex{},
    }

    wnd, cv, err := sdlcanvas.CreateWindow(engine.Dimensions.Width, engine.Dimensions.Height, engine.Title)
    if err != nil {
		panic(err)
    }
    
    engine.window = wnd
    engine.Canvas = cv

    return &engine
}

func (this *Engine) runPhysics() {
	frequency := time.Second / time.Duration(this.MaxTPS)

	timeSliceOpened := time.Now()
	ticksThisSecond := 0
	for true {
		tick := time.Now()

		this.tick()

		ticksThisSecond += 1
		if time.Since(timeSliceOpened) >= time.Second {
			this.currentHz = ticksThisSecond
			timeSliceOpened = time.Now()
			ticksThisSecond = 0

			if this.OnTickRateUpdated != nil {
				this.OnTickRateUpdated(this.currentHz)
			}
		}

		tock := time.Since(tick)
		if frequency-tock > 0 {
			// we just accept that this won't always be accurate
			// due to the resolution of the system clock. But that
			// just means the desired hz may not be perfect.
			time.Sleep(frequency - tock)
		}
	}
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

			if this.OnFPSUpdated != nil {
				this.OnFPSUpdated(this.currentFps)
			}
		}
	})
}

func (this *Engine) render() {
    this.renderMux.Lock()
    
    w, h := float64(this.Canvas.Width()), float64(this.Canvas.Height())
	this.Canvas.SetFillStyle("#000")
    this.Canvas.FillRect(0, 0, w, h)

    if this.scene.Render != nil {
        this.scene.Render(this, this.scene)
    }

    if this.scene.Update != nil {
        this.scene.Update(this, this.scene)
    }

    this.renderMux.Unlock()
}

func (this *Engine) tick() {
    this.tickMux.Lock()

    if this.scene.FixedUpdate != nil {
        this.scene.FixedUpdate(this, this.scene)
    }

    this.tickMux.Unlock()
}

func (this *Engine) SetScene(scene *Scene) {
    this.renderMux.Lock()
    this.tickMux.Lock()

    if this.scene != nil {
        this.scene.ClearResources()
    }
    this.scene = scene
    this.scene.LoadResources(this, this.scene)

    this.tickMux.Unlock()
    this.renderMux.Unlock()
}

func (this *Engine) Run() {
    go this.runPhysics()
	this.runGraphics()
}
