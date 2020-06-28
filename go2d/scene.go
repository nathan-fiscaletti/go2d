package go2d

type Scene struct {
    engine        *Engine
    resources     map[string]interface{}
    Name          string
    LoadResources func(engine *Engine, scene *Scene)

    Render        func(engine *Engine, scene *Scene)
    Update        func(engine *Engine, scene *Scene)
    FixedUpdate   func(engine *Engine, scene *Scene)
}

func NewScene(engine *Engine, name string) Scene {
    return Scene {
        engine: engine,
        resources: map[string]interface{}{},
        Name: name,
    }
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