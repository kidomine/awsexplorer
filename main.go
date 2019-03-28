package main

import (
	"github.com/gizak/termui"
	"github.com/kidomine/awsexplorer/model"
	"github.com/kidomine/awsexplorer/view"
	"log"
)

func main() {
	if err := termui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer termui.Close()

	var serviceList []string
	rs := model.NewRegionList()

	regionListView := view.NewRegionListView(model.GetRegionIds(rs))
	serviceListView := view.NewServiceListView(serviceList)

	currRegion := model.SelectRegionById(rs, regionListView.GetSelectedData())
	serviceListView.SetData(currRegion.GetServiceIds())

	regionListView.Render()
	serviceListView.Render()

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
