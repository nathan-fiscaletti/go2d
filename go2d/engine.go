package go2d

import (
    "fmt"
    "time"
    "sync"

    "github.com/tfriedel6/canvas"
    "github.com/tfriedel6/canvas/sdlcanvas"
)

type Engine struct {
    Title             string
    TickHz            int
    Dimensions        Dimensions
    OnTickRateUpdated func(*Engine, int)
    OnFPSUpdated      func(*Engine, int)
    Canvas            *canvas.Canvas

    scene             *Scene
    window            *sdlcanvas.Window
    currentHz         int
    currentFps        int
    renderMux         sync.Mutex
    tickMux           sync.Mutex
}

func NewEngine(title string, dimensions Dimensions) *Engine {
    engine := Engine{
        Title: title,
        Dimensions: dimensions,
        TickHz: 60,
        renderMux: sync.Mutex{},
        tickMux: sync.Mutex{},
    }

    wnd, cv, err := sdlcanvas.CreateWindow(engine.Dimensions.Width, engine.Dimensions.Height, engine.Title)
    if err != nil {
        panic(err)
    }

    engine.window = wnd
    engine.Canvas = cv

    engine.window.KeyDown = func(scancode int, rn rune, name string) {
        engine.scene.notifyKeyDown(scancode, rn, name)
    }

    engine.window.KeyUp = func(scancode int, rn rune, name string) {
        engine.scene.notifyKeyUp(scancode, rn, name)
    }

    return &engine
}

func (this *Engine) Bounds() Rect {
    return Rect {
        Vector: Vector {
            X:0, Y:0,
        },
        Dimensions: this.Dimensions,
    }
}

func (this *Engine) runPhysics() {
    frequency := time.Second / time.Duration(this.TickHz)

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
                this.OnTickRateUpdated(this, this.currentHz)
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
                this.OnFPSUpdated(this, this.currentFps)
            }
        }
    })
}

func (this *Engine) render() {
    this.renderMux.Lock()
    
    w, h := float64(this.Canvas.Width()), float64(this.Canvas.Height())
    this.Canvas.SetFillStyle("#000")
    this.Canvas.FillRect(0, 0, w, h)
    this.scene.performUpdate(this)
    this.scene.performRender(this)
    this.renderMux.Unlock()
}

func (this *Engine) tick() {
    this.tickMux.Lock()
    this.scene.performFixedUpdate(this)
    this.tickMux.Unlock()
}

func (this *Engine) GetFPS() int {
    return this.currentFps
}

func (this *Engine) GetHz() int {
    return this.currentHz
}

func (this *Engine) SetScene(scene *Scene) {
    this.renderMux.Lock()
    this.tickMux.Lock()

    if this.scene != nil {
        this.scene.ClearResources()
    }
    this.scene = scene
    this.scene.LoadResources(this, this.scene)
    this.window.Window.SetTitle(fmt.Sprintf("%s - %s", this.Title, this.scene.Name))

    this.tickMux.Unlock()
    this.renderMux.Unlock()
}

func (this *Engine) GetScene() *Scene {
    return this.scene
}

func (this *Engine) Run() {
    go this.runPhysics()
    this.runGraphics()
}
