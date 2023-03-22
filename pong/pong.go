package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/nathan-fiscaletti/go2d"
)

// Game Configuration
var PADDLE_WIDTH float64 = 32
var PADDLE_HEIGHT float64 = PADDLE_WIDTH * 5
var BALL_SIZE float64 = 32
var AI_PADDLE_RATE float64 = 5
var AI_PADDLE_TIME time.Duration = go2d.TICK_DURATION
var AI_BALL_RATE float64 = 20
var AI_BALL_TIME time.Duration = go2d.TICK_DURATION
var PLAYER_CONTROL_MOUSE bool = false
var PLAYER_CONTROL_KEY bool = true
var PLAYER_PADDLE_RATE float64 = AI_PADDLE_RATE * 2
var PLAYER_PADDLE_TIME time.Duration = go2d.TICK_DURATION

// Used by the game during run time
var PLAYER_PADDLE_MULTIPLIER float64 = 0
var CPU_PADDLE_MULTIPLIER float64 = 0
var PLAYER_SCORE int = 0
var CPU_SCORE int = 0
var SCORE_DISPLAY string = "0   0"

type ScoreDisplay struct {
	*go2d.TextEntity
}

func (this *ScoreDisplay) Update(engine *go2d.Engine) {
	this.Entity.Update()
}

func (this *ScoreDisplay) Render(engine *go2d.Engine) {
	this.SetText(SCORE_DISPLAY)
	this.TextEntity.Render(engine)
}

type Paddle struct {
	*go2d.ImageEntity
	aiControlled bool
}

func (this *Paddle) MouseMove(pos go2d.Vector) {
	if !this.aiControlled && PLAYER_CONTROL_MOUSE {
		this.MoveTo(go2d.Vector{
			X: this.Bounds.X,
			Y: pos.Y - (this.Bounds.Height / 2),
		})
	}
}

func (this *Paddle) KeyDown(scancode int, rn rune, name string) {
	if !this.aiControlled && PLAYER_CONTROL_KEY {
		if name == "ArrowUp" {
			this.Entity.Velocity = go2d.NewVelocityVector(0, -PLAYER_PADDLE_RATE, PLAYER_PADDLE_TIME)
		} else if name == "ArrowDown" {
			this.Entity.Velocity = go2d.NewVelocityVector(0, PLAYER_PADDLE_RATE, PLAYER_PADDLE_TIME)
		}
	}
}

func (this *Paddle) KeyUp(scancode int, rn rune, name string) {
	if !this.aiControlled {
		if name == "ArrowUp" || name == "ArrowDown" {
			this.Entity.Velocity = go2d.NewVelocityVector(0, 0, PLAYER_PADDLE_TIME)
		}
	}
}

func (this *Paddle) Update(engine *go2d.Engine) {
	this.Entity.Update()
	if this.aiControlled {
		this.Bounds.Height = PADDLE_HEIGHT - (CPU_PADDLE_MULTIPLIER * PADDLE_HEIGHT)
		ball := engine.GetScene().GetEntity(2, "ball").(*Ball)
		if ball.Bounds.Y < this.Bounds.Y+(this.Bounds.Height/2)-BALL_SIZE {
			this.Velocity = go2d.NewVelocityVector(
				0, -AI_PADDLE_RATE, AI_PADDLE_TIME,
			)
		} else if ball.Bounds.Y > this.Bounds.Y+(this.Bounds.Height/2)+BALL_SIZE {
			this.Velocity = go2d.NewVelocityVector(
				0, AI_PADDLE_RATE, AI_PADDLE_TIME,
			)
		} else {
			this.Velocity = go2d.NewVelocityVector(
				0, 0, AI_PADDLE_TIME,
			)
		}
	} else {
		this.Bounds.Height = PADDLE_HEIGHT - (PLAYER_PADDLE_MULTIPLIER * PADDLE_HEIGHT)
	}
}

func (this *Paddle) Constrain(engine *go2d.Engine) []go2d.RectSide {
	return this.Bounds.Constrain(engine.Bounds())
}

type Ball struct {
	*go2d.ImageEntity
	direction go2d.Vector
}

func (this *Ball) Update(engine *go2d.Engine) {
	this.Entity.Update()
	player := engine.GetScene().GetEntity(1, "player").(*Paddle)
	cpu := engine.GetScene().GetEntity(1, "cpu").(*Paddle)

	// handle paddle collision
	if this.Entity.CollidesWith(&player.Entity) || this.Entity.CollidesWith(&cpu.Entity) {
		this.direction = this.direction.InvertedX()

		var collidingEntity *Paddle

		if this.Entity.CollidesWith(&player.Entity) {
			collidingEntity = player
			PLAYER_PADDLE_MULTIPLIER += 0.05
		} else {
			collidingEntity = cpu
			CPU_PADDLE_MULTIPLIER += 0.05
		}

		yMax := float64(10)
		yMin := float64(-10)

		deadzone := float64(10)
		if this.Bounds.Y > collidingEntity.Bounds.Center().Y-deadzone/2 &&
			this.Bounds.Y < collidingEntity.Bounds.Center().Y+deadzone/2 {
			yMax = 0
			yMin = 0
		}

		if this.Bounds.IsAbove(collidingEntity.Bounds.Center()) {
			yMax = 2
		} else {
			yMin = -2
		}

		this.Velocity = go2d.NewVelocityVector(
			AI_BALL_RATE*this.direction.X,
			yMin+rand.Float64()*(yMax-yMin),
			AI_BALL_TIME,
		)
	} else {
		// handle scoring
		if this.Bounds.X < PADDLE_WIDTH {
			CPU_SCORE += 1
			SCORE_DISPLAY = fmt.Sprintf("%v   %v", PLAYER_SCORE, CPU_SCORE)
			PLAYER_PADDLE_MULTIPLIER = 0
			this.ReSpawn(go2d.GetActiveEngine(), go2d.DirectionLeft())
		} else if this.Bounds.X > engine.Bounds().Width-PADDLE_WIDTH {
			PLAYER_SCORE += 1
			SCORE_DISPLAY = fmt.Sprintf("%v   %v", PLAYER_SCORE, CPU_SCORE)
			CPU_PADDLE_MULTIPLIER = 0
			this.ReSpawn(go2d.GetActiveEngine(), go2d.DirectionRight())
		}
	}
}

func (this *Ball) Constrain(engine *go2d.Engine) []go2d.RectSide {
	return this.Bounds.Constrain(engine.Bounds())
}

func (this *Ball) OnConstrained(r go2d.RectSide) {
	if r == go2d.RectSideTop || r == go2d.RectSideBottom {
		this.Velocity.Y = -this.Velocity.Y
	}
}

func (this *Ball) ReSpawn(engine *go2d.Engine, direction go2d.Vector) {
	this.direction = direction
	this.MoveTo(engine.Bounds().Center())
	this.Velocity = go2d.NewVelocityVector(
		this.direction.X*AI_BALL_RATE,
		-10+rand.Float64()*(10 - -10),
		AI_BALL_TIME,
	)
}

type PongScene struct{}

func (this *PongScene) Initialize(engine *go2d.Engine, scene *go2d.Scene) {
	go2d.SetDefaultFont("../test_resources/font.ttf", 24, "#fff")

	score := go2d.NewTextEntitySimple("0   0")
	score.SetCentering(go2d.TEXT_CENTERING_VERTICAL | go2d.TEXT_CENTERING_HORIZONTAL)
	scoreDisplay := &ScoreDisplay{
		TextEntity: score,
	}
	scene.AddNamedEntity("score", 3, scoreDisplay)

	line := go2d.NewLineEntity(
		go2d.Vector{
			X: engine.Bounds().Width / 2,
		},
		go2d.DirectionDown(),
		engine.Bounds().Height,
		3, "#FFF",
	)
	scene.AddEntity(1, line)

	paddleImage := go2d.NewRectImageEntity("#FFFFFF", go2d.Dimensions{
		Width:  PADDLE_WIDTH,
		Height: PADDLE_HEIGHT,
	})

	paddle := &Paddle{
		ImageEntity:  paddleImage,
		aiControlled: false,
	}

	paddle.MoveTo(go2d.Vector{
		X: 0,
		Y: 0,
	})

	scene.AddNamedEntity("player", 1, paddle)

	paddleImage2 := go2d.NewRectImageEntity("#FFFFFF", go2d.Dimensions{
		Width:  PADDLE_WIDTH,
		Height: PADDLE_HEIGHT,
	})

	cpupaddle := &Paddle{
		ImageEntity:  paddleImage2,
		aiControlled: true,
	}

	cpupaddle.MoveTo(go2d.Vector{
		X: engine.Bounds().Width - cpupaddle.Bounds.Width,
		Y: 0,
	})

	scene.AddNamedEntity("cpu", 1, cpupaddle)

	ballImage := go2d.NewCircleImageEntity("#FFFFFF", int(BALL_SIZE))
	ball := &Ball{
		ImageEntity: ballImage,
	}

	ball.ReSpawn(engine, go2d.DirectionRight())

	scene.AddNamedEntity("ball", 2, ball)
}

func main() {
	engine := go2d.NewEngine(
		"Pong",
		go2d.NewAspectRatio(
			16, 9, go2d.AspectRatioControlAxisWidth,
		).NewDimensions(1200),
	)

	pongScene := &PongScene{}

	scene := go2d.NewScene(engine, "Level 1")
	scene.Initializer = pongScene

	scene.RenderStats("../test_resources/font.ttf", 24, "#ff0000")

	engine.SetScene(&scene)
	engine.HideCursor = true
	engine.Run()
}
