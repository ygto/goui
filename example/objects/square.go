package objects

import (
	"github.com/ygto/goui/object"
)

type Square struct {
	object.Object
}

func (t *Square) Shape() []float32 {

	return []float32{
		0, 0, 0,
		0, -10, 0,
		-10, -10, 0,

		-10, -10, 0,
		-10, 0, 0,
		0, 0, 0,
	}
}
