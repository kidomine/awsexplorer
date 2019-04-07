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

type TableView struct {
	tableWidget *widgets.Table
}

func newTableWidget(title string, x1, y1, x2, y2 int) *widgets.Table {
	tableWidget := widgets.NewTable()
	tableWidget.Title = title
	tableWidget.Rows = nil
	tableWidget.TextStyle = termui.NewStyle(termui.ColorYellow)
	tableWidget.SetRect(x1, y1, x2, y2)
	tableWidget.TextAlign = termui.AlignCenter

	return tableWidget
}

// This function should be called only once due to the fact that the
// coordinates of the widget is fixed.
func newTableView(title string, x1, y1, x2, y2 int, tableData [][]string) *TableView {
	tableWidget := newTableWidget(title, x1, y1, x2, y2)
	tableView := &TableView{tableWidget}

	tableView.SetData(tableData)

	return tableView
}

func (t *TableView) SetData(listData [][]string) {
	t.tableWidget.Rows = listData
}

func (t *TableView) GetSelectedData() string {
	//return t.tableWidget.Rows[t.tableWidget]
	return ""
}

func (t *TableView) Render() {
	termui.Render(t.tableWidget)
}

func (t *TableView) HandleEvent(event string) {
	switch event {
	}
}
