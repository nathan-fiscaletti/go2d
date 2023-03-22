package go2d

import (
	"fmt"
	"math/rand"
	"sort"
	"sync"
	"time"
)

// EntityGroup is an entity that represents a group of
// entities that can be rendered in layers.
type EntityGroup struct {
	Entity

	entities *sync.Map
}

type byLayer []int

func (a byLayer) Len() int           { return len(a) }
func (a byLayer) Less(i, j int) bool { return a[i] < a[j] }
func (a byLayer) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

// NewEntityGroup creates a new entity group.
func NewEntityGroup() *EntityGroup {
	return &EntityGroup{
		entities: &sync.Map{},
	}
}

// AddEntity adds an entity to the group.
func (this *EntityGroup) AddEntity(layer int, ent interface{}) string {
	n := time.Now().UnixNano()
	r := rand.New(rand.NewSource(n))
	id := fmt.Sprintf("entity_%v.%v", n, r.Intn(10000))

	layerData, _ := this.entities.LoadOrStore(layer, &sync.Map{})
	(layerData.(*sync.Map)).LoadOrStore(id, ent)

	return id
}

// AddNamedEntity adds an entity to the group with the given name.
func (this *EntityGroup) AddNamedEntity(name string, layer int, ent interface{}) {
	layerData, _ := this.entities.LoadOrStore(layer, &sync.Map{})
	(layerData.(*sync.Map)).LoadOrStore(name, ent)
}

// GetEntity gets an entity from the group.
func (this *EntityGroup) GetEntity(layer int, name string) interface{} {
	layerData, _ := this.entities.LoadOrStore(layer, &sync.Map{})
	res, _ := (layerData.(*sync.Map)).Load(name)
	return res
}

// RemoveEntity removes an entity from the group.
func (this *EntityGroup) RemoveEntity(layer int, name string) {
	layerData, _ := this.entities.LoadOrStore(layer, &sync.Map{})
	(layerData.(*sync.Map)).Delete(name)
}

// ClearEntities clears all entities from the group.
func (this *EntityGroup) ClearEntities() {
	this.entities.Range(func(key, value interface{}) bool {
		this.entities.Delete(key)
		return true
	})
}

// IterateEntities iterates over all entities in the group.
func (this *EntityGroup) IterateEntities(cb func(interface{})) {
	layers := []int{}
	this.entities.Range(func(key, value interface{}) bool {
		layers = append(layers, key.(int))
		return true
	})
	sort.Sort(byLayer(layers))

	for i := 0; i < len(layers); i++ {
		layer := layers[i]
		entities, _ := this.entities.Load(layer)
		if entities != nil {
			entities.(*sync.Map).Range(func(key, value interface{}) bool {
				cb(value)
				return true
			})
		}
	}
}

// Render renders the entity group.
func (this *EntityGroup) Render(engine *Engine) {
	this.IterateEntities(func(e interface{}) {
		_, isEntity := e.(IEntity)
		if isEntity {
			entity := e.(IEntity).GetEntity()
			originalLoc := entity.Bounds.Vector
			entity.Bounds.Vector = Vector{
				X: this.Bounds.Vector.X + originalLoc.X,
				Y: this.Bounds.Vector.Y + originalLoc.Y,
			}

			_, isRenderable := e.(IEntityRenderer)
			if isRenderable {
				e.(IEntityRenderer).Render(engine)
			}

			entity.Bounds.Vector = originalLoc
		}
	})
}

// Update updates the entity group.
func (this *EntityGroup) Update(engine *Engine) {
	this.Entity.Update()
}

// GetEntityGroupEntity returns the entity group entity.
func (this *EntityGroup) GetEntityGroupEntity() *Entity {
	return &this.Entity
}
