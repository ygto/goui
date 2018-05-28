
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

	rect := entity.CreateRectangle(50, 300, 100, 30)
	rect.SetColor(color.CreateColor(255, 0, 0))
	moveToStart := actions.MoveTo(vector.CreateVector2(50, 300), 0.5, nil)
	moveForward := actions.MoveTo(vector.CreateVector2(300, 300), 0.5, func(e *entity.Entity) {
		moveToStart.Reset()
		e.AddAction("move", moveToStart)
	})
	moveToStart.SetTrigger(func(e *entity.Entity) {
		moveForward.Reset()
		e.AddAction("move", moveForward)

	})
	rect.AddAction("move", moveForward)

	mainScene.AddEntity(rect)

	rect2 := entity.CreateRectangle(50, 300, 80, 30)
	rect2.SetColor(color.CreateColor(0, 0, 255))
	moveToStart2 := actions.MoveTo(vector.CreateVector2(50, 300), 0.2, nil)
	moveForward2 := actions.MoveTo(vector.CreateVector2(300, 300), 0.2, func(e *entity.Entity) {
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
