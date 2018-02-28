package window

import (
	"fmt"
	"log"
	"time"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/ygto/goui/scene"
	"github.com/ygto/goui/object"
)

const (
	DONTCARE  = iota
	RESIZABLE
	FIXED
)

type window struct {
	window     *glfw.Window
	program    uint32
	fps        uint8
	width      int
	height     int
	resizeable int
	sceneInit  bool
	winInit    bool
	scene      scene.SceneInterface
	objects    []object.Object
}

func New(width int, height int) *window {
	if err := glfw.Init(); err != nil {
		panic(fmt.Errorf("could not initialize glfw: %v", err))
	}
	w := new(window)
	w.width = width
	w.height = height

	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)

	switch w.resizeable {
	case 0:
		glfw.WindowHint(glfw.Resizable, glfw.DontCare)
	case 1:
		glfw.WindowHint(glfw.Resizable, glfw.True)
	default:
		glfw.WindowHint(glfw.Resizable, glfw.False)
	}
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	var err error
	w.window, err = glfw.CreateWindow(width, height, "", nil, nil)
	if err != nil {
		panic(fmt.Errorf("could not create opengl renderer: %v", err))
	}
	w.window.MakeContextCurrent()
	w.program = initOpenGL()
	gl.UseProgram(w.program)

	w.sceneInit = false
	w.fps = 60

	return w
}
func (w *window) Title(t string) {
	w.window.SetTitle(t)
}
func (w *window) Fps(fps uint8) {
	w.fps = fps
}
func (w *window) SetResizable(opt int) {
	w.resizeable = opt
}

func (w *window) Run() bool {

	shouldClose := w.window.ShouldClose()
	if !shouldClose {
		w.Update()
	}
	return !shouldClose
}

func (w *window) Destroy() {
	log.Println(w.program, "closed")
	glfw.Terminate()
}

func (w *window) Update() {

	start := time.Now()

	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	w.updateScene()
	glfw.PollEvents()
	w.window.SwapBuffers()

	//fps
	diff := time.Now().Sub(start)
	time.Sleep((time.Second / time.Duration(w.fps)) - diff)

}
