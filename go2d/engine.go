package go2d

import (
    "fmt"
    "time"
    "sync"

    "github.com/tfriedel6/canvas"
    "github.com/tfriedel6/canvas/sdlcanvas"
    "github.com/veandco/go-sdl2/sdl"
)

type Engine struct {
    Name              string
    Dimensions        Dimensions
    OnFPSUpdated      func(*Engine, int)
    Canvas            *canvas.Canvas
    HideCursor        bool

    scene             *Scene
    window            *sdlcanvas.Window
    currentFps        int
    renderMux         sync.Mutex
}

var runningEngines []*Engine = []*Engine{}

func NewEngine(name string, dimensions Dimensions) *Engine {
    engine := Engine{
        Name: name,
        Dimensions: dimensions,
        renderMux: sync.Mutex{},
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
        engine.scene.notifyMouseUp(b, Vector{X:float64(x), Y:float64(y)})
    }

    engine.window.MouseDown = func(b, x, y int) {
        engine.scene.notifyMouseDown(b, Vector{X:float64(x), Y:float64(y)})
    }

    engine.window.MouseMove = func(x, y int) {
        engine.scene.notifyMouseMove(Vector{X:float64(x), Y:float64(y)})
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

func (this *Engine) GetFPS() int {
    return this.currentFps
}

func (this *Engine) SetScene(scene *Scene) {
    this.renderMux.Lock()

    if this.scene != nil {
        this.scene.ClearResources()
    }
    this.scene = scene
    this.scene.Initialize(this, this.scene)
    this.window.Window.SetTitle(fmt.Sprintf("%s - %s", this.Name, this.scene.Name))

    this.renderMux.Unlock()
}

func (this *Engine) GetScene() *Scene {
    return this.scene
}

func (this *Engine) Run() {
    if this.HideCursor == true {
        sdl.ShowCursor(0)
    }
    runningEngines = append(runningEngines, this)
    this.runGraphics()
}

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

func GetActiveEngines() []*Engine {
    return runningEngines
}
