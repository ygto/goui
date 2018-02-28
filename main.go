package main

import (
	"runtime"
	"github.com/ygto/goui/window"
	"github.com/ygto/goui/example/scenes"
	"github.com/ygto/goui/example/objects"
	"github.com/ygto/goui/object/transform"
)

func main() {

	runtime.LockOSThread()
	win := window.New(1000, 600)
	defer win.Destroy()
	win.Title("hello world")
	win.SetResizable(window.FIXED)
	win.Fps(20)

	//mainScene := &scenes.MainScene{}
	mainScene := new(scenes.MainScene)
	/*
		matrix := [][]int{
			{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		}
		objMatrix := make(map[int]map[int]object.ObjectInterface)
		for y, yVal := range matrix {
			objMatrix[y] = make(map[int]object.ObjectInterface)
			for x, xVal := range yVal {

				obj := new(objects.Square)
				obj.Transform = transform.INIT_TRANSFORM
				obj.Position = transform.Vector2{float32(x*60) + 100, float32(y*60) + 100}
				obj.Scale = transform.Vector2{5, 5}
				if xVal == 0 {
					obj.Object.SetDrawable(false)
				} else {
					obj.Object.SetDrawable(true)
				}
				objMatrix[y][x] = obj

				mainScene.AddObject(obj)
			}
		}*/
	triangle := new(objects.Triangle)
	triangle.Transform = transform.INIT_TRANSFORM
	triangle.Position = transform.Vector2{100, 100}
	triangle.SetDrawable(true)

	mainScene.AddObject(triangle)

	win.ActiveScene(mainScene)
	for win.Run() {

	}

}
