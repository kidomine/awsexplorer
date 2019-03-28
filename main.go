package main

import (
	"github.com/gizak/termui"
	"github.com/kidomine/awsexplorer/controller"
	"log"
)

func main() {
	if err := termui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer termui.Close()

	ctrl := controller.Controller{}
	ctrl.Initialize()

	uiEvents := termui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		default:
			// let the controller handle the other events
			ctrl.HandleEvent(e.ID)
		}

		ctrl.Render()
	}
}
