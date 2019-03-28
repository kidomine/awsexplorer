package view

import (
	"github.com/gizak/termui"
	"github.com/gizak/termui/widgets"
)

type ListView struct {
	listWidget *widgets.List
}

func newListWidget(title string, x1, y1, x2, y2 int) *widgets.List {
	listWidget := widgets.NewList()
	listWidget.Title = title
	listWidget.Rows = nil
	listWidget.TextStyle = termui.NewStyle(termui.ColorYellow)
	listWidget.WrapText = false
	listWidget.SetRect(x1, y1, x2, y2)

	return listWidget
}

// This function should be called only once due to the fact that the
// coordinates of the widget is fixed.
func newListView(title string, x1, y1, x2, y2 int, listData []string) *ListView {
	listWidget := newListWidget(title, x1, y1, x2, y2)
	listView := &ListView{listWidget}

	listView.SetData(listData)

	return listView
}

func (r *ListView) SetData(listData []string) {
	r.listWidget.Rows = listData
}

func (r *ListView) GetSelectedData() string {
	return r.listWidget.Rows[r.listWidget.SelectedRow]
}

func (r *ListView) Render() {
	termui.Render(r.listWidget)
}

func (r *ListView) ScrollUp() {
	r.listWidget.ScrollUp()
}

func (r *ListView) ScrollDown() {
	r.listWidget.ScrollDown()
}
