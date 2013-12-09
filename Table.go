package walkwrap

import (
	. "github.com/joy999/utils"
	"github.com/lxn/walk"
	"github.com/lxn/win"
	"sort"
	"time"
)

type TableData struct {
	Index   int
	data    []interface{}
	checked bool
}

type Table struct {
	*walk.TableView
	model *TableModel
}

func (this *Table) Model() *TableModel {
	return this.model
}

func NewTable(parent walk.Container, width, height, x, y int) *Table {
	t := new(Table)
	t.model = NewTableModel()
	t.TableView, _ = walk.NewTableView(parent)
	t.SetWidth(width)
	t.SetHeight(height)
	t.SetX(x)
	t.SetY(y)
	t.SetModel(t.model)
	t.AddExStyle(win.LVS_EX_GRIDLINES)
	t.SetColumnsOrderable(false)

	menu, _ := walk.NewMenu()
	t.TableView.SetContextMenu(menu)

	return t
}

func (this *Table) AddColumn(title string, width int) {
	c := walk.NewTableViewColumn()
	if width > 0 {
		c.SetWidth(width)
	}
	c.SetTitle(title)
	this.Columns().Add(c)
}

func (this *Table) ClearColumn() {
	this.Columns().Clear()
}

func (this *Table) SelectedIndexes() *walk.IndexList {
	items := make([]int, 0)

	il := this.TableView.SelectedIndexes()
	itemLen := len(this.model.items)
	l := il.Len()
	for i := 0; i < l; i++ {
		index := il.At(i)
		if index < 0 || index >= itemLen {
			continue
		}
		items = append(items, this.model.items[index].Index)
	}

	return walk.NewIndexList(items)
}

func (this *Table) CurrentIndex() int {
	i := this.TableView.CurrentIndex()
	l := len(this.model.items)
	if i < 0 || i >= l {
		return -1
	} else {
		return this.model.items[i].Index
	}
}

func (this *Table) AddRowData(fields ...interface{}) int {
	return this.model.AddRow(fields...)
}

func (this *Table) SetChecked(col_index int, v bool) {
	len := this.model.Len()
	for i := 0; i < len; i++ {
		item := this.model.items[i]
		if item.Index == col_index {
			item.checked = v
			break
		}
	}
}

func (this *Table) GetChecked(col_index int) (checked bool, ok bool) {
	len := this.model.Len()
	ok = false
	checked = false
	for i := 0; i < len; i++ {
		item := this.model.items[i]
		if item.Index == col_index {
			checked = item.checked
			ok = true
			break
		}
	}

	return
}

func (this *Table) ClearRowData() {

	this.model.Clear()

}

func (this *Table) RemoveRowData(index int) {

	this.model.Romove(index)

}

func (this *Table) AddMenuAction(text string, cb func()) *Table {
	act := walk.NewAction()
	act.SetText(text)
	act.Triggered().Attach(cb)
	this.ContextMenu().Actions().Add(act)

	return this
}

func (this *Table) AddSeparatorAction() *Table {
	act := walk.NewSeparatorAction()
	this.ContextMenu().Actions().Add(act)

	return this
}

func (this *Table) AddExStyle(style int) {
	exStyle := this.SendMessage(win.LVM_GETEXTENDEDLISTVIEWSTYLE, 0, 0)
	exStyle |= uintptr(style)
	this.SendMessage(win.LVM_SETEXTENDEDLISTVIEWSTYLE, 0, exStyle)
}

func (this *Table) RemoveExStyle(style int) {
	exStyle := this.SendMessage(win.LVM_GETEXTENDEDLISTVIEWSTYLE, 0, 0)
	exStyle &= ^uintptr(style)
	this.SendMessage(win.LVM_SETEXTENDEDLISTVIEWSTYLE, 0, exStyle)
}

type TableModel struct {
	walk.TableModelBase
	walk.SorterBase
	sortColumn int
	sortOrder  walk.SortOrder
	items      []*TableData
}

func NewTableModel() *TableModel {
	m := new(TableModel)
	m.ResetRows(nil)
	return m
}

// Called by the TableView from SetModel and every time the model publishes a
// RowsReset event.
func (m *TableModel) RowCount() int {
	return len(m.items)
}

// Called by the TableView when it needs the text to display for a given cell.
func (m *TableModel) Value(row, col int) interface{} {
	item := m.items[row]
	return item.data[col]

}

// Called by the TableView to retrieve if a given row is checked.
func (m *TableModel) Checked(row int) bool {
	return m.items[row].checked
}

// Called by the TableView when the user toggled the check box of a given row.
func (m *TableModel) SetChecked(row int, checked bool) error {
	m.items[row].checked = checked

	return nil
}

// Called by the TableView to sort the model.
func (m *TableModel) Sort(col int, order walk.SortOrder) error {
	m.sortColumn, m.sortOrder = col, order

	sort.Sort(m)

	return m.SorterBase.Sort(col, order)
}

func (m *TableModel) Len() int {
	return len(m.items)
}

func (m *TableModel) Less(i, j int) bool {
	a, b := m.items[i].data, m.items[j].data

	c := func(ls bool) bool {

		if m.sortOrder == walk.SortAscending {
			if !ls {
				m.items[i], m.items[j] = m.items[j], m.items[i]
			}
			return ls
		}
		if ls {
			m.items[i], m.items[j] = m.items[j], m.items[i]
		}
		return !ls
	}

	av := a[m.sortColumn]
	bv := b[m.sortColumn]
	switch v := av.(type) {
	case int:
		if _v, ok := bv.(int); ok {
			return c(v < _v)
		}
	case float64:
		if _v, ok := bv.(float64); ok {
			return c(v < _v)
		}
	case time.Time:
		if _v, ok := bv.(time.Time); ok {
			return c(v.Before(_v))
		}
	case int8:
		if _v, ok := bv.(int8); ok {
			return c(v < _v)
		}
	case int16:
		if _v, ok := bv.(int16); ok {
			return c(v < _v)
		}
	case int32:
		if _v, ok := bv.(int32); ok {
			return c(v < _v)
		}

	case int64:
		if _v, ok := bv.(int64); ok {
			return c(v < _v)
		}
	case uint:
		if _v, ok := bv.(uint); ok {
			return c(v < _v)
		}
	case uint8:
		if _v, ok := bv.(uint8); ok {
			return c(v < _v)
		}
	case uint16:
		if _v, ok := bv.(uint16); ok {
			return c(v < _v)
		}
	case uint32:
		if _v, ok := bv.(uint32); ok {
			return c(v < _v)
		}
	case uint64:
		if _v, ok := bv.(uint64); ok {
			return c(v < _v)
		}
	case float32:
		if _v, ok := bv.(float32); ok {
			return c(v < _v)
		}

	case string:
		if _v, ok := bv.(string); ok {
			return c(v < _v)
		}
	case String:
		if _v, ok := bv.(String); ok {
			return c(string(v) < string(_v))
		}
	case Time_t:
		if _v, ok := bv.(Time_t); ok {
			return c(v < _v)
		}

	}
	return c(false)

}

func (m *TableModel) Swap(i, j int) {
	m.items[i], m.items[j] = m.items[j], m.items[i]
}

// Called by the TableView to retrieve an item image.
func (m *TableModel) Image(row int) interface{} {
	return nil
	//if m.items[row].Index%2 == 0 {
	//	return m.evenBitmap
	//}

	//return m.oddIcon
}

func (m *TableModel) ResetRows(data [][]interface{}) {
	// Create some random data.
	if data != nil {
		tdArr := make([]*TableData, len(data))
		for i, row := range data {
			td := new(TableData)
			td.data = make([]interface{}, len(row))
			for t, v := range row {
				td.data[t] = v
			}
			td.checked = false
			tdArr[i] = td
		}
		m.items = tdArr
	} else {
		m.items = make([]*TableData, 0)
	}
	// Notify TableView and other interested parties about the reset.
	m.PublishRowsReset()

	m.Sort(m.sortColumn, m.sortOrder)
}

func (m *TableModel) Clear() {
	m.items = make([]*TableData, 0)
	m.PublishRowsReset()
	m.Sort(m.sortColumn, m.sortOrder)
}

func (m *TableModel) Add(data []interface{}) int {
	td := new(TableData)
	td.data = data
	td.checked = false
	index := len(m.items)
	m.items = append(m.items, td)
	m.PublishRowsReset()
	m.Sort(m.sortColumn, m.sortOrder)
	td.Index = index
	return index
}

func (m *TableModel) AddRow(fields ...interface{}) int {
	td := new(TableData)
	td.data = fields
	td.checked = false
	index := len(m.items)
	m.items = append(m.items, td)
	m.PublishRowsReset()
	m.Sort(m.sortColumn, m.sortOrder)
	td.Index = index
	return index
}

func (m *TableModel) Romove(index int) {
	items := make([]*TableData, 0)
	for i, v := range m.items {
		if i == index {
			continue
		}
		items = append(items, v)
	}
	m.items = items
	m.PublishRowsReset()
	m.Sort(m.sortColumn, m.sortOrder)
}
