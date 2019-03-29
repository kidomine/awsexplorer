/*
 * MIT License
 *
 * Copyright (c) 2019 Ian Diaz.
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

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

func (r *ListView) HandleEvent(event string) {
	switch event {
	case "j", "<Down>":
		r.ScrollDown()
	case "k", "<Up>":
		r.ScrollUp()
	}
}

func (r *ListView) ScrollUp() {
	r.listWidget.ScrollUp()
}

func (r *ListView) ScrollDown() {
	r.listWidget.ScrollDown()
}
