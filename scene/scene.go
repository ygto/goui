package scene

import (
	"github.com/ygto/goui/object"
)

type SceneInterface interface {
	Init()
	Update()
	GetObjects() []object.ObjectInterface
}
type Scene struct {
	Objects []object.ObjectInterface
}

func (s *Scene) AddObject(obj object.ObjectInterface) {
	s.Objects = append(s.Objects, obj)
}

func (s *Scene) GetObjects() []object.ObjectInterface {
	return s.Objects
}
