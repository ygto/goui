package entity

import (
	"github.com/ygto/goui/vector"
	"github.com/ygto/goui/color"
)

func CreateRectangle(x float32, y float32, width float32, height float32) *Entity {
	e := createEntity()
	e.shape = SHAPE_RECTANGLE
	e.size = vector.CreateVector2(width, height)
	e.position = vector.CreateVector2(x, y)
	e.SetColor(color.White())

	return e
}
func CreateCircle(x float32, y float32, radius float32) *Entity {
	e := createEntity()
	e.shape = SHAPE_CIRCLE
	e.size = vector.CreateVector2(radius, radius)
	e.position = vector.CreateVector2(x, y)
	e.SetColor(color.White())

	return e
}
