package main

import (
	"github.com/lxn/walk"

//	"github.com/lxn/win"
)

type TextEdit struct {
	*walk.TextEdit
}

func NewTextEdit(parent walk.Container, text string, w, h, x, y int) *TextEdit {
	t, err := walk.NewTextEdit(parent)
	if err != nil {
		return nil
	}

	font, _ := walk.NewFont("微软雅黑", 9, 0)
	t.SetFont(font)

	t.SetText(text)
	t.SetWidth(w)
	t.SetHeight(h)
	t.SetX(x)
	t.SetY(y)

	te := new(TextEdit)
	te.TextEdit = t

	return te
}

type LineEdit struct {
	*walk.LineEdit
}

func NewLineEdit(parent walk.Container, text string, w, h, x, y int) *LineEdit {
	t, err := walk.NewLineEdit(parent)
	if err != nil {
		return nil
	}

	font, _ := walk.NewFont("微软雅黑", 9, 0)
	t.SetFont(font)

	t.SetText(text)
	t.SetWidth(w)
	t.SetHeight(h)
	t.SetX(x)
	t.SetY(y)

	te := new(LineEdit)
	te.LineEdit = t

	return te
}
