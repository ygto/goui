package vector

type Vector2 struct {
	x, y float32
}

func CreateVector2(x float32, y float32) *Vector2 {
	v2 := new(Vector2)
	v2.x = x
	v2.y = y

	return v2
}

func (v2 *Vector2) GetX() float32 {
	return v2.x
}
func (v2 *Vector2) GetY() float32 {
	return v2.y
}

func (v2 *Vector2) GetXInt() int {
	return int(v2.x)
}
func (v2 *Vector2) GetYInt() int {
	return int(v2.y)
}

func (v2 *Vector2) SetX(x float32) {
	v2.x = x
}
func (v2 *Vector2) SetY(y float32) {
	v2.y = y
}

func (v2 *Vector2) AddX(x float32) {
	v2.x += x
}
func (v2 *Vector2) AddY(y float32) {
	v2.y += y
}
