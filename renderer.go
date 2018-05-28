package goui

import "github.com/veandco/go-sdl2/sdl"

type Renderer struct {
	renderer *sdl.Renderer
}

func NewRenderer(win *Window) (*Renderer, error) {
	r := new(Renderer)
	var err error
	r.renderer, err = sdl.CreateRenderer(win.GetSdlWindow(), -1, sdl.RENDERER_ACCELERATED)
	if err != nil {

		return nil, err
	}

	return r, nil
}
func (r *Renderer) Copy(t *Texture) {
	r.renderer.Copy(t.texture, nil, nil)
}

func (r *Renderer) Present() {
	r.renderer.Present()
}

func (r *Renderer) Destroy() {
	r.renderer.Destroy()
}
