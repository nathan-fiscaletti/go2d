package main

import(
    "../go2d"
)

const PLAYER_SIZE  = 48
const PLAYER_SPEED = 10
const PLAYER_RAIL  = 16
const BULLET_SIZE  = 32
const BULLET_SPEED = 10
const PLAYER_LAYER = 1
const BULLET_LAYER = 0
const ENEMY_LAYER  = 0
const ENEMY_SIZE   = PLAYER_SIZE
const ENEMY_RAIL   = PLAYER_RAIL

type Shooter struct {
    *go2d.ImageEntity
}

func (this *Shooter) Constrain(engine *go2d.Engine) []go2d.RectSide {
    return this.Bounds.Constrain(engine.Bounds())
}

func (this *Shooter) KeyDown(scancode int, rn rune, name string) {
    if name == "ArrowLeft" {
        this.Entity.Velocity = go2d.NewVelocityVector(-PLAYER_SPEED, 0, go2d.TICK_DURATION)
    } else if name == "ArrowRight" {
        this.Entity.Velocity = go2d.NewVelocityVector(PLAYER_SPEED, 0, go2d.TICK_DURATION)
    }
}

func (this *Shooter) KeyUp(scancode int, rn rune, name string) {
    if name == "ArrowLeft" || name == "ArrowRight" {
        this.Entity.Velocity = go2d.NewVelocityVector(0, 0, go2d.TICK_DURATION)
    } else if name == "Space" {
        this.Shoot()
    }
}

func (this *Shooter) Shoot() {
    image := go2d.NewCircleImageEntity("#00FF00", BULLET_SIZE)
    bullet := &Projectile{
        ImageEntity: image,
    }
    bullet.Entity.Velocity = go2d.NewVelocityVector(0, -BULLET_SPEED, go2d.TICK_DURATION)
    bullet.MoveTo(go2d.Vector{
        X: this.Bounds.X + (this.Bounds.Width / 2) - BULLET_SIZE / 2,
        Y: this.Bounds.Y,
    })
    bullet.key = go2d.GetActiveScene().AddEntity(BULLET_LAYER, bullet)
}

type Projectile struct {
    *go2d.ImageEntity
    key string
}

func (this *Projectile) Update(engine *go2d.Engine) {
    this.Entity.Update()
    if this.Bounds.Y < -BULLET_SIZE {
        this.Remove()
    }
}

func (this *Projectile) GetCollider() go2d.Rect {
    return this.Bounds
}

func (this *Projectile) CollidedWith(other interface{}) {
    other.(*Enemy).Remove()
    this.Remove()
}

func (this *Projectile) Remove() {
    go2d.GetActiveScene().RemoveEntity(BULLET_LAYER, this.key)
}

// Maintain a list of active enemies
var activeEnemies map[string]*Enemy = map[string]*Enemy{}

type Enemy struct {
    *go2d.ImageEntity
    key string
}

// Remove the enemy if it has gone below the bounds of the screen.
func (this *Enemy) Update(engine *go2d.Engine) {
    this.Entity.Update()
    if this.Bounds.Y > engine.Bounds().Height {
        this.Remove()
    }
}

// Add a collider to enemies
func (this *Enemy) GetCollider() go2d.Rect {
    return this.Bounds
}

// Remove will remove the Enemy from it's owning scene.
func (this *Enemy) Remove() {
    delete(activeEnemies, this.key)
    go2d.GetActiveScene().RemoveEntity(ENEMY_LAYER, this.key)
}

// AddEnemyRow moves all existing enemies down by one grid element and spawns a new
// row of enemies.
func AddEnemyRow(engine *go2d.Engine) {
    for _,e := range activeEnemies {
        e.MoveTo(go2d.Vector{
            X: e.Bounds.X,
            Y: e.Bounds.Y + ENEMY_SIZE,
        })
    }

    enemyRowCount := int(engine.Bounds().Width / (ENEMY_SIZE))
    for i := 0; i < enemyRowCount; i++ {
        enemyImage := go2d.NewCircleImageEntity("#0000FF", ENEMY_SIZE)
        enemy := &Enemy{
            ImageEntity: enemyImage,
        }
        enemy.MoveTo(go2d.Vector{
            X: float64(i * ENEMY_SIZE),
            Y: 0,
        })
        enemy.key = engine.GetScene().AddEntity(ENEMY_LAYER, enemy)
        activeEnemies[enemy.key] = enemy
    }
}

func main() {
    engine := go2d.NewEngine(
        "Shooter",
        go2d.NewAspectRatio(
            16, 9, go2d.AspectRatioControlAxisWidth,
        ).NewDimensions(1200),
    )

    scene := go2d.NewScene(engine, "Level 1")
    scene.Initialize = func(engine *go2d.Engine, scene *go2d.Scene) {
        // Create a rectangle image for the player
        playerImage := go2d.NewRectImageEntity("#FFFFFF", go2d.Dimensions{
            Width: PLAYER_SIZE,
            Height: PLAYER_SIZE,
        })

        // Create the player instance
        player := &Shooter{
            ImageEntity: playerImage,
        }

        // Move them to their dedicated rail at the bottom of the scren
        player.MoveTo(
            go2d.Vector{
                X: 10,
                Y: engine.Bounds().Height - PLAYER_RAIL - PLAYER_SIZE,
            },
        )

        // Add them to the scene
        scene.AddNamedEntity("player", 1, player)
    }

    // Spawn a new row of enemies every two seconds.
    scene.AddTimer("EnemySpawner", &go2d.Timer{
        Seconds: 2,
        Update: func(owner interface{}) {
            AddEnemyRow(engine)
        },
    })

    scene.RenderFPS("../test_resources/font.ttf", 24, "#ff0000")

    engine.SetScene(&scene)
    engine.HideCursor = true
    engine.Run()
}