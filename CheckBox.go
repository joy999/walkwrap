package main

import (
	"github.com/lxn/walk"
)

type CheckBox struct {
	*walk.CheckBox
}

func NewCheckBox(parent walk.Container, text string, width, height, x, y int) *CheckBox {
	b := new(CheckBox)
	b.CheckBox, _ = walk.NewCheckBox(parent)
	b.SetWidth(width)
	b.SetHeight(height)
	b.SetY(y)
	b.SetX(x)
	b.SetText(text)

	return b
}
