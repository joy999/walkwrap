package main

import (
	"github.com/lxn/walk"
)

type DropDownBox struct {
	*walk.ComboBox
	model *DropDownBoxModel
}

func NewDropDownBox(parent walk.Container, text string, width, height, x, y int) *DropDownBox {
	b := new(DropDownBox)

	b.ComboBox, _ = walk.NewDropDownBox(parent)
	b.SetWidth(width)
	b.SetHeight(height)
	b.SetY(y)
	b.SetX(x)

	b.model = NewDropDownBoxModel()
	b.SetModel(b.model)

	return b
}

func (this *DropDownBox) AddItem(name string) {
	this.model.Add(name)
}

type DropDownBoxModel struct {
	walk.ListModelBase
	data []string
}

/*
// ListModel is the interface that a model must implement to support widgets
// like ComboBox.
type ListModel interface {
	// ItemCount returns the number of items in the model.
	ItemCount() int

	// Value returns the value that should be displayed for the given index.
	Value(index int) interface{}

	// ItemsReset returns the event that the model should publish when the
	// number of its items changes.
	ItemsReset() *Event

	// ItemChanged returns the event that the model should publish when an item
	// was changed.
	ItemChanged() *IntEvent
}
*/

func NewDropDownBoxModel() *DropDownBoxModel {
	d := new(DropDownBoxModel)
	d.data = make([]string, 0)

	return d
}

func (this *DropDownBoxModel) Value(index int) interface{} {
	return this.data[index]
}

func (this *DropDownBoxModel) ItemCount() int {
	return len(this.data)
}

func (this *DropDownBoxModel) Add(v string) {
	this.data = append(this.data, v)

	this.PublishItemsReset()
}
