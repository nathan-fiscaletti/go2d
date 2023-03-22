package go2d

import (
	"fmt"
)

type ISceneInitializer interface {
	Initialize(engine *Engine, scene *Scene)
}

type IScenePreRenderer interface {
	PreRender(engine *Engine, scene *Scene)
}

type ISceneRenderer interface {
	Render(engine *Engine, scene *Scene)
}

type ISceneUpdater interface {
	Update(engine *Engine, scene *Scene)
}

// Scene is a simple scene implementation.
type Scene struct {
	*EntityGroup

	// Name is the name of the scene.
	Name string

	// Initializer is the initializer that will be called when the scene is initialized.
	Initializer ISceneInitializer
	// PreRenderer is the pre-renderer that will be called before the scene is rendered.
	PreRenderer IScenePreRenderer
	// Renderer is the renderer that will be called when the scene is rendered.
	Renderer ISceneRenderer
	// Updater is the updater that will be called when the scene is updated.
	Updater ISceneUpdater

	renderStats bool
	statsEntity *TextEntity
	engine      *Engine
	resources   map[string]interface{}
	timers      map[string]*timer
}

// GetActiveScene returns the active scene. If no scene is active, nil is returned. If you are using
// multiple scenes, you should use GetActiveScenes() otherwise you will get a panic.
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

// GetActiveScenes returns all active scenes.
func GetActiveScenes() []*Scene {
	scenes := []*Scene{}
	for _, e := range GetActiveEngines() {
		scenes = append(scenes, e.GetScene())
	}

	return scenes
}

// NewScene creates a new scene with the given name.
func NewScene(engine *Engine, name string) Scene {
	return Scene{
		EntityGroup: NewEntityGroup(),
		engine:      engine,
		timers:      map[string]*timer{},
		resources:   map[string]interface{}{},
		Name:        name,
	}
}

// RenderStats will enable rendering of the FPS and TPS in the top left corner of the screen. The
// font, size and color can be specified. The stats will be rendered on top of everything else.
func (this *Scene) RenderStats(font string, size float64, color string) {
	this.renderStats = true
	this.statsEntity = NewTextEntity("FPS: 0, TPS: 0", font, size, color)
}

// StopRenderingStats will stop rendering the FPS and TPS.
func (this *Scene) StopRenderingStats() {
	this.renderStats = false
}

// AddTimer adds a timer to the scene.
func (this *Scene) AddTimer(name string, t *timer) {
	this.timers[name] = t
}

// RemoveTimer removes a timer from the scene.
func (this *Scene) RemoveTimer(name string) {
	delete(this.timers, name)
}

// GetResource returns a resource by name.
func (this *Scene) GetResource(name string) interface{} {
	return this.resources[name]
}

// SetResource sets a resource by name.
func (this *Scene) SetResource(name string, res interface{}) {
	this.resources[name] = res
}

// ClearResources clears all resources.
func (this *Scene) ClearResources() {
	for k := range this.resources {
		delete(this.resources, k)
	}
}

func (this *Scene) notifyMouseMove(pos Vector) {
	this.IterateEntities(func(e interface{}) {
		_, isMouseSensitive := e.(IMouseMove)
		if isMouseSensitive {
			e.(IMouseMove).MouseMove(pos)
		}
	})
}

func (this *Scene) notifyMouseUp(button int, pos Vector) {
	this.IterateEntities(func(e interface{}) {
		_, isMouseSensitive := e.(IMouseUp)
		if isMouseSensitive {
			e.(IMouseUp).MouseUp(button, pos)
		}
	})
}

func (this *Scene) notifyMouseDown(button int, pos Vector) {
	this.IterateEntities(func(e interface{}) {
		_, isMouseSensitive := e.(IMouseDown)
		if isMouseSensitive {
			e.(IMouseDown).MouseDown(button, pos)
		}
	})
}

func (this *Scene) notifyKeyUp(scanCode int, rn rune, name string) {
	this.IterateEntities(func(e interface{}) {
		_, isKeySensitive := e.(IKeyUp)
		if isKeySensitive {
			e.(IKeyUp).KeyUp(scanCode, rn, name)
		}
	})
}

func (this *Scene) notifyKeyDown(scanCode int, rn rune, name string) {
	this.IterateEntities(func(e interface{}) {
		_, isKeySensitive := e.(IKeyDown)
		if isKeySensitive {
			e.(IKeyDown).KeyDown(scanCode, rn, name)
		}
	})
}

func (this *Scene) notifyKeyChar(rn rune) {
	this.IterateEntities(func(e interface{}) {
		_, isKeySensitive := e.(IKeyChar)
		if isKeySensitive {
			e.(IKeyChar).KeyChar(rn)
		}
	})
}

func (this *Scene) performUpdate(engine *Engine) {
	for _, t := range this.timers {
		t.notifyUpdate(this, this)
	}

	this.EntityGroup.Update(engine)

	// Handle constraint and Update calls
	this.IterateEntities(func(e interface{}) {
		_, isConstrain := e.(IEntityConstraint)
		if isConstrain {
			constrainedSides := e.(IEntityConstraint).Constrain(engine)
			_, isConstrained := e.(IEntityConstrainedHandler)
			if isConstrained {
				for _, side := range constrainedSides {
					e.(IEntityConstrainedHandler).OnConstrained(side)
				}
			}
		}

		_, isUpdatable := e.(IEntityUpdater)
		if isUpdatable {
			e.(IEntityUpdater).Update(engine)
		}
	})

	// Handle Collision
	this.IterateEntities(func(e interface{}) {
		_, hasCollisionDetection := e.(IEntityCollisionDetection)
		_, hasCollider := e.(IEntityCollider)
		if hasCollisionDetection && hasCollider {
			this.IterateEntities(func(e2 interface{}) {
				if e2 != e {
					_, hasCollider2 := e2.(IEntityCollider)
					if hasCollider2 {
						if e.(IEntityCollider).GetCollider().IntersectsWith(
							e2.(IEntityCollider).GetCollider(),
						) {
							e.(IEntityCollisionDetection).CollidedWith(e2)
						}
					}
				}
			})
		}
	})

	if this.Updater != nil {
		this.Updater.Update(engine, this)
	}

	if this.renderStats {
		this.statsEntity.SetText(fmt.Sprintf("FPS: %v, TPS: %v", engine.GetFPS(), engine.GetTPS()))
	}
}

func (this *Scene) performRender(engine *Engine) {
	if this.PreRenderer != nil {
		this.PreRenderer.PreRender(engine, this)
	}

	this.EntityGroup.Render(engine)

	if this.Renderer != nil {
		this.Renderer.Render(engine, this)
	}

	if this.renderStats {
		this.statsEntity.Render(engine)
	}
}
