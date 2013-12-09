package walkwrap

import (
	"github.com/lxn/walk"

//	"github.com/lxn/win"
)

type Label struct {
	*walk.Label
}

func NewLabel(parent walk.Container, txt string, w, h, x, y int) *Label {
	ll := new(Label)
	ll.Label, _ = walk.NewLabel(parent)

	font, _ := walk.NewFont("微软雅黑", 9, 0)
	ll.Label.SetFont(font)
	//WS_EX_TRANSPARENT
	ll.Label.SetText(txt)
	ll.Label.SetWidth(w)
	ll.Label.SetHeight(h)
	ll.Label.SetX(x)
	ll.Label.SetY(y)

	//win.SetWindowLongPtr(ll.Label.Handle(), win.GWL_EXSTYLE, win.WS_EX_TRANSPARENT)

	return ll
}

type LinkLabel struct {
	*walk.Label
}

func NewLinkLabel(parent walk.Container, txt string, w, h, x, y int) *LinkLabel {
	ll := new(LinkLabel)
	ll.Label, _ = walk.NewLabel(parent)

	font, _ := walk.NewFont("微软雅黑", 9, walk.FontUnderline)
	ll.Label.SetFont(font)

	ll.Label.SetText(txt)
	ll.Label.SetWidth(w)
	ll.Label.SetHeight(h)
	ll.Label.SetX(x)
	ll.Label.SetY(y)

	return ll
}
