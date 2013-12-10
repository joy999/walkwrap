package walkwrap

import (
	"github.com/lxn/walk"
)

type Button struct {
	*walk.PushButton
}

func NewButton(parent walk.Container, text string, width, height, x, y int) *Button {
	b := new(Button)
	b.PushButton, _ = walk.NewPushButton(parent)
	size := walk.Size{width, height}
	b.SetSize(size)
	b.SetMinMaxSize(size, size)
	b.SetY(y)
	b.SetX(x)
	b.SetText(text)

	return b
}
