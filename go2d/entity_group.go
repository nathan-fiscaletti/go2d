package go2d

import(
    "fmt"
    "sort"
    "time"
    "math/rand"
)

type EntityGroup struct {
    Entity

    entities map[int]map[string]interface{}
}

type ByLayer []int

func (a ByLayer) Len() int           { return len(a)}
func (a ByLayer) Less(i, j int) bool { return a[i] < a[j] }
func (a ByLayer) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func NewEntityGroup() *EntityGroup {
    return &EntityGroup{
        entities: map[int]map[string]interface{}{},
    }
}

func (this *EntityGroup) Add(layer int, ent interface{}) string {
    n := time.Now().UnixNano()
    r := rand.New(rand.NewSource(n))
    id := fmt.Sprintf("entity_%v.%v", n, r.Intn(10000))
    if _,layerExists := this.entities[layer]; !layerExists {
        this.entities[layer] = map[string]interface{}{}
    }
    this.entities[layer][id] = ent
    return id
}

func (this *EntityGroup) AddNamed(name string, layer int, ent interface{}) {
    if _,layerExists := this.entities[layer]; !layerExists {
        this.entities[layer] = map[string]interface{}{}
    }
    this.entities[layer][name] = ent
}

func (this *EntityGroup) Get(layer int, name string) interface{} {
    if _,layerExists := this.entities[layer]; layerExists {
        if _,entityExists := this.entities[layer][name]; entityExists {
            return this.entities[layer][name]
        }
    }
    
    return nil
}

func (this *EntityGroup) Remove(layer int, name string) {
    delete(this.entities[layer], name)
}

func (this *EntityGroup) Clear() {
    for k := range this.entities {
        delete(this.entities, k)
    }
}

func (this *EntityGroup) IterateEntities(cb func(interface{})) {
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
            cb(e)
        }
    }
}

func (this *EntityGroup) Render(engine *Engine) {
    this.IterateEntities(func (e interface{}) {
        _,isEntity := e.(IEntity)
        if isEntity {
            entity := e.(IEntity).GetEntity()
            originalLoc := entity.Bounds.Vector
            entity.Bounds.Vector = Vector{
                X: this.Bounds.Vector.X + originalLoc.X,
                Y: this.Bounds.Vector.Y + originalLoc.Y,
            }

            _,isRenderable := e.(IRender)
            if isRenderable {
                e.(IRender).Render(engine)
            }

            entity.Bounds.Vector = originalLoc
        }
    })
}

func (this *EntityGroup) Update(engine *Engine) {
    this.Entity.Update()
}

func (this *EntityGroup) GetEntity() *Entity {
    return &this.Entity
}