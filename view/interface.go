package view

type View interface {
	Render()
	HandleEvent(event string)
	GetSelectedData() string
}
