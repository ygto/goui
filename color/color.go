package color

type Color struct {
	R, G, B byte
}

func CreateColor(red byte, green byte, blue byte) *Color {
	c := Color{red, green, blue}

	return &c
}
func White() *Color {
	return CreateColor(255, 255, 255)
}

func Black() *Color {
	return CreateColor(0, 0, 0)
}

func Red() *Color {
	return CreateColor(255, 0, 0)
}
