package objects

import (
	"github.com/ygto/goui/object"
)

type Triangle struct {
	object.Object
}

func (t *Triangle) Shape() []float32 {

	return []float32{
		0, 10, 0,
		-10, -10, 0,
		10, -10, 0,
	}
}
