package transform

type Transform struct {
	Position Vector2
	Scale    Vector2
}

var INIT_TRANSFORM = Transform{Position: Vector2{0, 0}, Scale: Vector2{1, 1}}
