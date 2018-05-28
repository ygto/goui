package goui

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/ygto/goui/vector"
)

type Window struct {
	window   *sdl.Window
	size     *vector.Vector2
	title    string
}

func NewWindow(w int, h int, title string) (*Window, error) {
	win := new(Window)
	win.size = vector.CreateVector2(float32(w), float32(h))
	win.title = title
	var err error
	win.window, err = sdl.CreateWindow(title, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, int32(win.GetSize().GetX()), int32(win.GetSize().GetY()), sdl.WINDOW_SHOWN)
	if err != nil {
		return nil, err
	}

	return win, nil
}

func (w *Window) GetSdlWindow() *sdl.Window {
	return w.window
}

func (w *Window) GetSize() *vector.Vector2 {
	return w.size
}

func (w *Window) Destroy() {
	w.window.Destroy();
}
