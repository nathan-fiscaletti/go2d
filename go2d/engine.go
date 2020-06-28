package go2d

import (
	"fmt"
	"time"

	"./metrics"
	"github.com/tfriedel6/canvas"
	"github.com/tfriedel6/canvas/sdlcanvas"
)

type Engine struct {
	MaxFPS int
	MaxTPS int

	Dimensions metrics.Dimensions
	Title      string

	OnTickRateUpdated func(int)
	OnFPSUpdated      func(int)

	currentHz  int
	currentFps int

	window *sdlcanvas.Window
	canvas *canvas.Canvas
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
	w, h := float64(this.canvas.Width()), float64(this.canvas.Height())
	this.canvas.SetFillStyle("#000")
	this.canvas.FillRect(0, 0, w, h)
	this.canvas.SetFillStyle("#FFF")
	this.canvas.FillText(fmt.Sprintf("FPS: %v, TPS: %v", this.currentFps, this.currentHz), 0, 32)
}

func (this *Engine) tick() {

}

func (this *Engine) Run() {
	font := "/Users/nathan/Downloads/Anonymous/Anonymous.ttf"

	wnd, cv, err := sdlcanvas.CreateWindow(this.Dimensions.Width, this.Dimensions.Height, this.Title)

	if err != nil {
		panic(err)
	}
	cv.SetFont(font, 40)

	this.window = wnd
	this.canvas = cv

	go this.runPhysics()
	this.runGraphics()
}
