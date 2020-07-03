package go2d

import(
    "fmt"
    "time"
    "math/rand"
    "sort"
)

type Scene struct {
    Name          string
    LoadResources func(engine *Engine, scene *Scene)

    PreRender     func(engine *Engine, scene *Scene)
    Render        func(engine *Engine, scene *Scene)
    Update        func(engine *Engine, scene *Scene)
    FixedUpdate   func(engine *Engine, scene *Scene)

    engine        *Engine
    resources     map[string]interface{}
    entities      map[int]map[string]interface{}
}

type ByLayer []int

func (a ByLayer) Len() int           { return len(a)}
func (a ByLayer) Less(i, j int) bool { return a[i] < a[j] }
func (a ByLayer) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func NewScene(engine *Engine, name string) Scene {
    return Scene {
        engine: engine,
        resources: map[string]interface{}{},
        entities: map[int]map[string]interface{}{},
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

func (this *Scene) GetEntity(layer int, name string) interface{} {
    if _,layerExists := this.entities[layer]; layerExists {
        if _,entityExists := this.entities[layer][name]; entityExists {
            return this.entities[layer][name]
        }
    }
    
    return nil
}

func (this *Scene) AddNamedEntity(name string, layer int, ent interface{}) {
    _, err := entityForInterface(ent)
    if err != nil {
        panic(err)
    }

    if _,layerExists := this.entities[layer]; !layerExists {
        this.entities[layer] = map[string]interface{}{}
    }
    this.entities[layer][name] = ent
}

func (this *Scene) AddEntity(layer int, ent interface{}) string {
    _, err := entityForInterface(ent)
    if err != nil {
        panic(err)
    }

    n := time.Now().UnixNano()
    r := rand.New(rand.NewSource(n))
    id := fmt.Sprintf("entity_%v.%v", n, r.Intn(10000))
    if _,layerExists := this.entities[layer]; !layerExists {
        this.entities[layer] = map[string]interface{}{}
    }
    this.entities[layer][id] = ent
    return id
}

func (this *Scene) RemoveEntity(layer int, name string) {
    delete(this.entities[layer], name)
}

func (this *Scene) ClearEntities() {
    for k := range this.entities {
        delete(this.entities, k)
    }
}

func (this *Scene) notifyKeyUp(scancode int, rn rune, name string) {
    i := 0
    layers := make([]int, len(this.entities))
    for k := range this.entities {
        layers[i] = k
        i++
    }
    sort.Sort(ByLayer(layers))

    for i := 0; i < len(layers); i++ {
        layer := layers[i]
        for _, e := range this.entities[layer] {
            _,isControllable := e.(Controllable)
            if isControllable {
                e.(Controllable).KeyUp(scancode, rn, name)
            }
        }
    }
}

func (this *Scene) notifyKeyDown(scancode int, rn rune, name string) {
    i := 0
    layers := make([]int, len(this.entities))
    for k := range this.entities {
        layers[i] = k
        i++
    }
    sort.Sort(ByLayer(layers))

    for i := 0; i < len(layers); i++ {
        layer := layers[i]
        for _, e := range this.entities[layer] {
            _,isControllable := e.(Controllable)
            if isControllable {
                e.(Controllable).KeyDown(scancode, rn, name)
            }
        }
    }
}

func (this *Scene) performUpdate(engine *Engine) {
    i := 0
    layers := make([]int, len(this.entities))
    for k := range this.entities {
        layers[i] = k
        i++
    }
    sort.Sort(ByLayer(layers))

    for i := 0; i < len(layers); i++ {
        layer := layers[i]
        for _, e := range this.entities[layer] {
            _,isUpdatable := e.(Updatable)
            if isUpdatable {
                e.(Updatable).Update(engine)
            }
        }
    }
    if this.Update != nil {
        this.Update(engine, this)
    }
}

func (this *Scene) performFixedUpdate(engine *Engine) {
    i := 0
    layers := make([]int, len(this.entities))
    for k := range this.entities {
        layers[i] = k
        i++
    }
    sort.Sort(ByLayer(layers))

    for i := 0; i < len(layers); i++ {
        layer := layers[i]
        for _, e := range this.entities[layer] {
            _,isFixedUpdatable := e.(FixedUpdatable)
            if isFixedUpdatable {
                e.(FixedUpdatable).FixedUpdate(engine)
            }
        }
    }
    if this.FixedUpdate != nil {
        this.FixedUpdate(engine, this)
    }
}

func (this *Scene) performRender(engine *Engine) {
    if this.PreRender != nil {
        this.PreRender(engine, this)
    }

    i := 0
    layers := make([]int, len(this.entities))
    for k := range this.entities {
        layers[i] = k
        i++
    }
    sort.Sort(ByLayer(layers))

    for i := 0; i < len(layers); i++ {
        layer := layers[i]
        for _, e := range this.entities[layer] {
            _,isRenderable := e.(Renderable)
            if isRenderable {
                e.(Renderable).Render(this.engine)
            }
        }
    }
    if this.Render != nil {
        this.Render(engine, this)
    }
}