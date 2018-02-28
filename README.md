# basic Opengl framework

why? because to learn Opengl with go

base libraries

"github.com/go-gl/gl/v4.1-core/gl"

"github.com/go-gl/glfw/v3.2/glfw"

**Usage**


```go
runtime.LockOSThread()

win := window.New(1000, 600)
defer win.Destroy()
win.Title("hello world")
win.SetResizable(window.FIXED)
win.Fps(20)

mainScene := new(scenes.MainScene)

triangle := new(objects.Triangle)
triangle.Transform = transform.INIT_TRANSFORM
triangle.Position = transform.Vector2{100, 100}
triangle.SetDrawable(true)

mainScene.AddObject(triangle)

win.ActiveScene(mainScene)
for win.Run() {
}
```
