package go2d

type TimerTrigger interface {
	OnTriggered(owner interface{})
}

// Timer is a simple timer that can be used to trigger events at a certain interval.
type timer struct {
	// Seconds is the interval in seconds at which the timer will trigger.
	seconds float64

	currentTickInSecond int
	trigger             TimerTrigger
}

// NewTimer creates a new timer that will trigger the given TimerTrigger every given seconds.
func NewTimer(seconds float64, trigger TimerTrigger) *timer {
	return &timer{
		seconds: seconds,
		trigger: trigger,
	}
}

func (this *timer) notifyUpdate(owner interface{}, scene *Scene) {
	if scene.engine.GetFPS() > 0 && (float64(this.currentTickInSecond)/float64(scene.engine.GetTPS())) > this.seconds {
		this.trigger.OnTriggered(owner)
		this.currentTickInSecond = 0
	}

	this.currentTickInSecond += 1
}
