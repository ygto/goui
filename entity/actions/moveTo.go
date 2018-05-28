package actions

import (
	"github.com/ygto/goui/entity"
	"github.com/ygto/goui/vector"
	"math"
)

type moveTo struct {
	hasTarget     bool
	targetTrigger func(e *entity.Entity)
	targetPoint   *vector.Vector2
	speed         float32
}

func MoveTo(targetPoint *vector.Vector2, speed float32, f func(e *entity.Entity)) *moveTo {
	b := new(moveTo)
	b.hasTarget = true
	b.targetPoint = targetPoint
	b.speed = speed
	b.targetTrigger = f

	return b
}

func (a *moveTo) GetName() string {
	return "moveTo"
}

func (a *moveTo) Reset() {
	a.hasTarget = true
}
func (a *moveTo) Copy() entity.ActionInterface {
	return MoveTo(a.targetPoint, a.speed, a.targetTrigger)
}
func (a *moveTo) SetTrigger(f func(e *entity.Entity)) {
	a.targetTrigger = f
}
func (a *moveTo) UpdateAction(e *entity.Entity) {

	//fmt.Println(a.targetPoint.GetX(), a.targetPoint.GetY(), e.GetPosition().GetX(), e.GetPosition().GetY())
	if !a.hasTarget {
		return
	}
	if !a.IsReactedToTarget(e) {

		diffX := a.targetPoint.GetX() - e.GetPosition().GetX()
		if math.Abs(float64(diffX)) < float64(a.speed) {
			e.GetPosition().AddX(diffX)
		}
		diffY := a.targetPoint.GetY() - e.GetPosition().GetY()
		if math.Abs(float64(diffY)) < float64(a.speed) {
			e.GetPosition().AddY(diffY)
		}

		if e.GetPosition().GetX() > a.targetPoint.GetX() {
			e.GetPosition().AddX(-a.speed)
		} else if e.GetPosition().GetX() < a.targetPoint.GetX() {
			e.GetPosition().AddX(a.speed)

		}
		if e.GetPosition().GetY() > a.targetPoint.GetY() {
			e.GetPosition().AddY(-a.speed)
		} else if e.GetPosition().GetY() < a.targetPoint.GetY() {
			e.GetPosition().AddY(a.speed)

		}
	} else {
		a.hasTarget = false
		if a.targetTrigger != nil {
			a.targetTrigger(e)
		}
	}

}

func (a *moveTo) IsReactedToTarget(e *entity.Entity) bool {
	if a.targetPoint.GetX() == e.GetPosition().GetX() && a.targetPoint.GetY() == e.GetPosition().GetY() {
		a.hasTarget = false

		return true
	}
	return false
}
