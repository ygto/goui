package entity

type ActionInterface interface {
	UpdateAction(e *Entity)
	GetName() string
	SetTrigger(func(e *Entity))
	Reset()
	Copy() ActionInterface
}
