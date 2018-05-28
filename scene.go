package goui

import (
	"github.com/ygto/goui/entity"
	"fmt"
)

type Scene struct {
	name        string
	goui        *Goui
	layers      map[SceneIndex]*Texture
	entities    map[SceneIndex][]*entity.Entity
	activeLayer SceneIndex //default 0
}

func CreateScene(name string, g *Goui) (*Scene, error) {
	s := new(Scene)
	s.name = name
	s.goui = g
	s.layers = make(map[SceneIndex]*Texture)
	s.entities = make(map[SceneIndex][]*entity.Entity)
	s.SetLayer(0)

	return s, nil
}
func (s *Scene) SetLayer(i SceneIndex) {
	if i < 0 {
		fmt.Errorf("layer can't be lower than 0")
	}
	if _, ok := s.layers[i]; ok {
		return
	}
	s.addLayerAndEntities(i)
	s.activeLayer = i
}

func (s *Scene) addLayerAndEntities(index SceneIndex) error {
	texture, err := NewTexture(s.goui.win, s.goui.renderer)
	if err != nil {

		return err
	}
	s.layers[SceneIndex(index)] = texture
	s.entities[SceneIndex(index)] = make([]*entity.Entity, 0, 0)
	fmt.Println(len(s.layers), len(s.entities))

	return nil
}
func (s *Scene) AddEntity(e *entity.Entity) {
	index := SceneIndex(s.activeLayer)
	s.entities[index] = append(s.entities[index], e)
}

func (s *Scene) GetEntities() map[SceneIndex][]*entity.Entity {
	return s.entities
}

func (s *Scene) Update() {
	for i := 0; i < len(s.entities); i++ {
		for _, entity := range s.entities[SceneIndex(i)] {
			entity.Update()
		}
	}
}

func (s *Scene) Draw() {

	for i := 0; i < len(s.layers); i++ {
		layer, ok := s.layers[SceneIndex(i)]
		if !ok {
			continue
		}
		layer.Clear()
		//fmt.Println("i draw layer :",i)
		for _, entity := range s.entities[SceneIndex(i)] {
			//fmt.Println("i draw entity :",entity.GetColor().R)

			layer.Draw(entity)
		}
		layer.Update()
		s.goui.renderer.Copy(layer)
	}
}
