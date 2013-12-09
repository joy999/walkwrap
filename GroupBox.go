package walkwrap

import (
	"github.com/lxn/walk"
)

type GroupBox struct {
	*walk.GroupBox
}

func NewGroupBox(parent walk.Container, text string, width, height, x, y int) *GroupBox {
	var err error
	b := new(GroupBox)
	b.GroupBox, err = walk.NewGroupBox(parent)

	if err != nil {
		return nil
	}

	b.SetWidth(width)
	b.SetHeight(height)
	b.SetY(y)
	b.SetX(x)
	b.GroupBox.SetTitle(text)

	return b
}
