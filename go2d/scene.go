package go2d

import(
    "fmt"
)

type Scene struct {
    *EntityGroup

    Name          string
    Initialize    func(engine *Engine, scene *Scene)

    PreRender     func(engine *Engine, scene *Scene)
    Render        func(engine *Engine, scene *Scene)
    Update        func(engine *Engine, scene *Scene)

    renderFPS     bool
    fpsEntity     *TextEntity

    engine        *Engine
    resources     map[string]interface{}
    timers        map[string]*Timer
}

func GetActiveScene() *Scene {
    engineCount := len(GetActiveEngines())
    if engineCount < 1 {
        return nil
    }

    if engineCount > 1 {
        panic("More than one running Engine. Please use GetActiveScenes().")
    }

    return GetActiveEngine().GetScene()
}

func GetActiveScenes() []*Scene {
    scenes := []*Scene{}
    for _,e := range GetActiveEngines() {
        scenes = append(scenes, e.GetScene())
    }

    return scenes
}

func NewScene(engine *Engine, name string) Scene {
    return Scene {
        EntityGroup: NewEntityGroup(),
        engine: engine,
        timers: map[string]*Timer{},
        resources: map[string]interface{}{},
        Name: name,
    }
}

func (this *Scene) RenderFPS(font string, size float64, color string) {
    this.renderFPS = true
    this.fpsEntity = NewTextEntity("FPS: 0", font, size, color)
}

func (this *Scene) StopRenderingFPS() {
    this.renderFPS = false
}

func (this *Scene) AddTimer(name string, t *Timer) {
    this.timers[name] = t
}

func (this *Scene) RemoveTimer(name string) {
    delete(this.timers, name)
}

func (this *Scene) GetResource(name string) interface{} {
    return this.resources[name]
}

func (this *Scene) SetResource(name string, res interface{}) {
    this.resources[name] = res
}

func (this *Scene) ClearResources() {
    for k := range this.resources {
        delete(this.resources, k)
    }
}

func (this *Scene) notifyMouseMove(pos Vector) {
    this.IterateEntities(func (e interface{}) {
        _,isMouseSensitive := e.(IMouseMove)
        if isMouseSensitive {
            e.(IMouseMove).MouseMove(pos)
        }
    })
}

func (this *Scene) notifyMouseUp(button int, pos Vector) {
    this.IterateEntities(func (e interface{}) {
        _,isMouseSensitive := e.(IMouseUp)
        if isMouseSensitive {
            e.(IMouseUp).MouseUp(button, pos)
        }
    })
}

func (this *Scene) notifyMouseDown(button int, pos Vector) {
    this.IterateEntities(func (e interface{}) {
        _,isMouseSensitive := e.(IMouseDown)
        if isMouseSensitive {
            e.(IMouseDown).MouseDown(button, pos)
        }
    })
}

func (this *Scene) notifyKeyUp(scancode int, rn rune, name string) {
    this.IterateEntities(func (e interface{}) {
        _,isKeySensitive := e.(IKeyUp)
        if isKeySensitive {
            e.(IKeyUp).KeyUp(scancode, rn, name)
        }
    })
}

func (this *Scene) notifyKeyDown(scancode int, rn rune, name string) {
    this.IterateEntities(func (e interface{}) {
        _,isKeySensitive := e.(IKeyDown)
        if isKeySensitive {
            e.(IKeyDown).KeyDown(scancode, rn, name)
        }
    })
}

func (this *Scene) notifyKeyChar(rn rune) {
    this.IterateEntities(func (e interface{}) {
        _,isKeySensitive := e.(IKeyChar)
        if isKeySensitive {
            e.(IKeyChar).KeyChar(rn)
        }
    })
}

func (this *Scene) performUpdate(engine *Engine) {
    for _,t := range this.timers {
        t.notifyUpdate(this, this)
    }

    this.EntityGroup.Update(engine)

    // Handle constraint and Update calls
    this.IterateEntities(func (e interface{}) {
        _,isConstrain := e.(IConstrain)
        if isConstrain {
            constrainedSides := e.(IConstrain).Constrain(engine)
            _,isConstrained := e.(IConstrained)
            if isConstrained {
                for _, side := range constrainedSides {
                    e.(IConstrained).Constrained(side)
                }
            }
        }

        _,isUpdatable := e.(IUpdate)
        if isUpdatable {
            e.(IUpdate).Update(engine)
        }
    })

    // Handle Collision
    this.IterateEntities(func (e interface{}) {
        _,isCollisionDetection := e.(ICollisionDetection)
        _,isCollidable := e.(ICollidable)
        if isCollisionDetection && isCollidable {
            this.IterateEntities(func (e2 interface{}) {
                if e2 != e {
                    _,isCollidable2 := e2.(ICollidable)
                    if isCollidable2 {
                        if e.(ICollidable).GetCollider().IntersectsWith(
                            e2.(ICollidable).GetCollider(),
                        ) {
                            e.(ICollisionDetection).CollidedWith(e2)
                        }
                    }
                }
            })
        }
    })
    
    if this.Update != nil {
        this.Update(engine, this)
    }

    if this.renderFPS {
        this.fpsEntity.SetText(fmt.Sprintf("FPS: %v", engine.GetFPS()))
    }
}

func (this *Scene) performRender(engine *Engine) {
    if this.PreRender != nil {
        this.PreRender(engine, this)
    }

    this.EntityGroup.Render(engine)

    if this.Render != nil {
        this.Render(engine, this)
    }

    if this.renderFPS {
        this.fpsEntity.Render(engine)
    }
}

