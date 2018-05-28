package goui

import (
	"github.com/veandco/go-sdl2/sdl"
	"time"
	"fmt"
)

type SceneIndex int8
type Goui struct {
	win         *Window
	renderer    *Renderer
	activeScene string
	fps         int64
	scenes      map[string]*Scene
}

func (goui *Goui) SetFps(fps int) {
	goui.fps = int64(1000 / fps)
}

func (goui *Goui) AddScene(s *Scene) {
	goui.scenes[s.name] = s
}

func (g *Goui) getActiveScene() *Scene {
	return g.scenes[g.activeScene]
}

func (g *Goui) SetActiveScene(name string) {
	if ok := g.scenes[name]; ok == nil {
		fmt.Errorf(" %s scene not found", name)
	}
	g.activeScene = name
}

func CreateGoui(width int, height int, title string) (*Goui, error) {

	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {

		return nil, err
	}

	g := new(Goui)

	g.scenes = make(map[string]*Scene)

	g.win, err = NewWindow(width, height, title)
	if err != nil {
		return nil, err
	}

	g.renderer, err = NewRenderer(g.win)
	if err != nil {
		return nil, err
	}

	return g, nil
}

func (g *Goui) Update() bool {
	t := time.Now().Nanosecond()
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch event.(type) {
		case *sdl.QuitEvent:
			return false

		}
	}
	g.getActiveScene().Update()

	diff := int64(time.Now().Nanosecond()-t) / int64(time.Millisecond)
	if diff > 0 && diff < g.fps {
		sdl.Delay(uint32(g.fps - diff))

	}

	return true
}
func (g *Goui) Draw() {
	g.getActiveScene().Draw()
	g.renderer.Present()
}
func (g *Goui) Destroy() {
	defer sdl.Quit()
	g.win.Destroy()
}
