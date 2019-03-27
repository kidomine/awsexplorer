package main

import (
	"github.com/gizak/termui"
	"github.com/gizak/termui/widgets"
	"github.com/kidomine/awsexplorer/model"
	"log"
)

//TODO: temporary
func newListWidget(title string, x1, y1, x2, y2 int) *widgets.List {
	lw := widgets.NewList()
	lw.Title = title
	lw.Rows = nil
	lw.TextStyle = termui.NewStyle(termui.ColorYellow)
	lw.WrapText = false
	lw.SetRect(x1, y1, x2, y2)

	return lw
}

func main() {
	if err := termui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer termui.Close()

	regionListWidget := newListWidget("Regions", 0, 0, 20, 45)
	serviceListWidget := newListWidget("Services", 21, 0, 41, 45)

	rs := model.NewRegionList()

	regionListWidget.Rows = model.GetRegionIds(rs)
	serviceListWidget.Rows = nil

	termui.Render(regionListWidget)

	uiEvents := termui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		case "h", "<Left>":
			//
		case "l", "<Right>":
			//
		}
	}
}
