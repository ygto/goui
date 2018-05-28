package entity

import (
	"github.com/ygto/goui/vector"
	"github.com/ygto/goui/color"
)

const (
	SHAPE_RECTANGLE = iota
	SHAPE_CIRCLE
)

type Entity struct {
	position *vector.Vector2
	size     *vector.Vector2
	color    *color.Color
	shape    int
	actions  map[string]ActionInterface
}

func createEntity() *Entity {
	e := new(Entity)
	e.actions = make(map[string]ActionInterface)
	return e
}

func (e *Entity) GetAction(name string) ActionInterface {

	return e.actions[name]
}
func (e *Entity) SetColor(c *color.Color) {
	e.color = c
}
func (e *Entity) GetPosition() *vector.Vector2 {
	return e.position
}

func (e *Entity) GetSize() *vector.Vector2 {
	return e.size
}

func (e *Entity) GetShape() int {
	return e.shape
}

func (e *Entity) GetColor() *color.Color {
	return e.color
}

func (e *Entity) AddAction(name string, a ActionInterface) {
	e.actions[a.GetName()] = a
}

func (e *Entity) Update() {
	for _, b := range e.actions {
		b.UpdateAction(e)
	}
}
