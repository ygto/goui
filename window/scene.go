package window

import (
	"github.com/ygto/goui/scene"
	"github.com/ygto/goui/object"
)

func (w *window) ActiveScene(s scene.SceneInterface) {
	w.scene = s
	w.scene.Init()
}

func (w *window) updateScene() {
	w.scene.Update()
	w.drawObjects()
}

func (w *window) drawObjects() {
	winHeight := float32(w.height)
	winWidth := float32(w.width)

	var o object.ObjectInterface
	var s float32
	var calc float32 = 0
	var i int = 0
	var tempShape []float32
	for _, o = range w.scene.GetObjects() {
		if o.Drawable() {
			i = 0
			tempShape = []float32{}

			for _, s = range o.Shape() {

				//todo check memory leak!!!!
				switch i % 3 {
				case 0:
					calc = ((s * o.GetTransform().Scale.X) + o.GetTransform().Position.X - winWidth) / winWidth
				case 1:

					calc = ((s * o.GetTransform().Scale.Y) + o.GetTransform().Position.Y - winHeight) / winHeight
				}
				tempShape = append(tempShape, calc)

				i++
			}
			//fmt.Println(tempShape)
			drawOject(tempShape)

		}
	}
}
