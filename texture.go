package goui

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/ygto/goui/color"
	"github.com/ygto/goui/vector"
	"github.com/ygto/goui/entity"
)

type Texture struct {
	size    *vector.Vector2
	texture *sdl.Texture
	pixels  []byte
}

func NewTexture(win *Window, renderer *Renderer) (*Texture, error) {
	t := new(Texture)
	t.size = win.GetSize()
	var err error
	t.texture, err = renderer.renderer.CreateTexture(
		sdl.PIXELFORMAT_ABGR8888,
		sdl.TEXTUREACCESS_STREAMING,
		int32(t.size.GetX()),
		int32(t.size.GetY()),
	)
	t.pixels = make([]byte, win.GetSize().GetXInt()*win.GetSize().GetYInt()*4)
	clearTexture(t.pixels)

	if err != nil {

		return nil, err
	}

	return t, nil
}
func (t *Texture) setAlpha(alpha uint8) {
	t.texture.SetAlphaMod(alpha)
}
func (t *Texture) setPixel(x int, y int, c *color.Color) {
	index := (t.size.GetXInt()*y + x) * 4
	if index+4 < cap(t.pixels) {
		t.pixels[index] = c.R
		t.pixels[index+1] = c.G
		t.pixels[index+2] = c.B
	}
}

func (t *Texture) Clear() {
	for y := 0; y < t.size.GetYInt(); y++ {
		for x := 0; x < t.size.GetXInt(); x++ {
			t.setPixel(x, y, color.Black())
		}
	}
}

func (t *Texture) Draw(e *entity.Entity) {
	switch e.GetShape() {
	case entity.SHAPE_RECTANGLE:
		t.drawRectangle(e)
	case entity.SHAPE_CIRCLE:
		t.drawCircle(e)
	}
}
func (t *Texture) drawRectangle(e *entity.Entity) {
	for y := 0; y < int(e.GetSize().GetY()); y++ {
		for x := 0; x < int(e.GetSize().GetX()); x++ {
			t.setPixel(int(e.GetPosition().GetX())+x, int(e.GetPosition().GetY())+y, e.GetColor())
		}
	}
}

func (t *Texture) drawCircle(e *entity.Entity) {

	for y := -int(e.GetSize().GetY()); y < int(e.GetSize().GetY()); y++ {
		for x := -int(e.GetSize().GetY()); x < int(e.GetSize().GetX())*2; x++ {
			if x*x+y*y < int(e.GetSize().GetX()*e.GetSize().GetX()) {
				t.setPixel(int(e.GetPosition().GetX())+x, int(e.GetPosition().GetY())+y, e.GetColor())
			}
		}
	}
}

func (t *Texture) Update() {
	t.texture.Update(nil, t.pixels, t.size.GetXInt()*4)
}

func (t *Texture) Destroy() {
	t.texture.Destroy()
}

func clearTexture(bytes []byte) {
	bytes = make([]byte, len(bytes))
}
