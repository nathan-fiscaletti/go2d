package go2d

type Timer struct {
    Seconds float64
    Update func(owner interface{})

    currentTickInSecond int
}

func (this *Timer) notifyUpdate(owner interface{}, scene *Scene) {
    if scene.engine.GetFPS() > 0 && (float64(this.currentTickInSecond) / float64(scene.engine.GetFPS())) > this.Seconds {
        this.Update(owner)
        this.currentTickInSecond = 0
    }

    this.currentTickInSecond += 1
}

