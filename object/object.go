package object

import "github.com/ygto/goui/object/transform"

type ObjectInterface interface {
	Drawable() bool
	SetDrawable(bool)
	Shape() []float32
	GetTransform() *transform.Transform
}
type Object struct {
	transform.Transform
	IsDrawable bool
}

func (obj *Object) GetTransform() *transform.Transform {

	return &obj.Transform
}

func (obj *Object) Drawable() bool {

	return obj.IsDrawable
}

func (obj *Object) SetDrawable(b bool) {

	obj.IsDrawable = b
}
