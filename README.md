# basic Opengl framework

why? because to learn Opengl with go

base libraries

"github.com/go-gl/gl/v4.1-core/gl"

"github.com/go-gl/glfw/v3.2/glfw"

**Usage**


```go

func main() {
	g, err := goui.CreateGoui(800, 600, "yigit")
	g.SetFps(60)

	mainScene, err := goui.CreateScene("main", g)
	if err != nil {
		panic(err)
	}
	g.AddScene(mainScene)
	g.SetActiveScene("main")

	rect := entity.NewRectangle(50, 300, 100, 30)
	rect.SetColor(color.CreateColor(255, 0, 0))
	moveToStart := actions.MoveTo(vector.NewVector2(50, 300), 0.5, nil)
	moveForward := actions.MoveTo(vector.NewVector2(300, 300), 0.5, func(e *entity.Entity) {
		moveToStart.Reset()
		e.AddAction("move", moveToStart)
	})
	moveToStart.SetTrigger(func(e *entity.Entity) {
		moveForward.Reset()
		e.AddAction("move", moveForward)

	})
	rect.AddAction("move", moveForward)

	mainScene.AddEntity(rect)

	rect2 := entity.NewRectangle(50, 300, 80, 30)
	rect2.SetColor(color.CreateColor(0, 0, 255))
	moveToStart2 := actions.MoveTo(vector.NewVector2(50, 300), 0.2, nil)
	moveForward2 := actions.MoveTo(vector.NewVector2(300, 300), 0.2, func(e *entity.Entity) {
		moveToStart2.Reset()
		e.AddAction("move", moveToStart2)

	})
	moveToStart2.SetTrigger(func(e *entity.Entity) {
		moveForward2.Reset()
		e.AddAction("move", moveForward2)

	})
	rect2.AddAction("move", moveForward2)

	mainScene.AddEntity(rect2)
	mainScene.GetEntities()

	for g.Update() {
		g.Draw()
	}
}
```
