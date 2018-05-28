package entity

import (
	"github.com/ygto/goui/vector"
	"github.com/ygto/goui/color"
)

func NewRectangle(x float32, y float32, width float32, height float32) *Entity {
	e := newEntity()
	e.shape = SHAPE_RECTANGLE
	e.size = vector.NewVector2(width, height)
	e.position = vector.NewVector2(x, y)
	e.SetColor(color.White())

	return e
}
func NewCircle(x float32, y float32, radius float32) *Entity {
	e := newEntity()
	e.shape = SHAPE_CIRCLE
	e.size = vector.NewVector2(radius, radius)
	e.position = vector.NewVector2(x, y)
	e.SetColor(color.White())

	return e
}
