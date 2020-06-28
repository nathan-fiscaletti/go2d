package main

// brew install sdl2{,_image,_mixer,_ttf,_gfx} pkg-config
// go get https://github.com/tfriedel6/canvas/sdlcanvas

// "github.com/tfriedel6/canvas/sdlcanvas"

import (
	"./go2d"
	"./go2d/metrics"
)

func main() {
	engine := go2d.Engine{
		MaxFPS: 60,
		MaxTPS: 50,
		Dimensions: metrics.NewAspectDimensions(
			metrics.AspectRatio{
				Ratio: metrics.Dimensions{
					Width:  16,
					Height: 9,
				},
				ControlAxis: metrics.AspectRatioControlAxisWidth,
			},
			1200,
		),
	}
	engine.Run()

	// lockToHz := 50
	// v := metrics.NewRandomVector(metrics.Vector{
	// 	X: 100,
	// 	Y: 100,
	// })

	// print(fmt.Sprintf("%v, %v\n", v.X, v.Y))

	// TODO: I should keep the physics update in a seprate thread and run it more
	//       frequently, this would allow it to update physics as often as possible
	//       but stick to the regular 60fps limit. Rendering would happen only once
	//       every 60th of a second (ideally). Physics updates happen at 50hz.

	//       I think physics should be based on time.

	// wnd, cv, err := sdlcanvas.CreateWindow(1280, 720, "Hello")
	// if err != nil {
	// 	panic(err)
	// }
	// defer wnd.Destroy()

	// lcnt := 0
	// strt := time.Duration(time.Now().UnixNano())

	// wnd.MainLoop(func() {
	// 	w, h := float64(cv.Width()), float64(cv.Height())
	// 	cv.SetFillStyle("#000")
	// 	cv.FillRect(0, 0, w, h)

	// 	for r := 0.0; r < math.Pi*2; r += math.Pi * 0.1 {
	// 		cv.SetFillStyle(int(r*10), int(r*20), int(r*40))
	// 		cv.BeginPath()
	// 		cv.MoveTo(w*0.5, h*0.5)
	// 		cv.Arc(w*0.5, h*0.5, math.Min(w, h)*0.4, r, r+0.1*math.Pi, false)
	// 		cv.ClosePath()
	// 		cv.Fill()
	// 	}

	// 	cv.SetStrokeStyle("#FFF")
	// 	cv.SetLineWidth(10)
	// 	cv.BeginPath()
	// 	cv.Arc(w*0.5, h*0.5, math.Min(w, h)*0.4, 0, math.Pi*2, false)
	// 	cv.Stroke()

	// 	lcnt += 1

	// 	now := time.Duration(time.Now().UnixNano())
	// 	if now-strt >= time.Second {
	// 		fmt.Printf("%v fps\n", lcnt)
	// 		lcnt = 0
	// 		strt = now
	// 	}
	// })
}
