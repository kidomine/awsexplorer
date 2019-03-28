package controller

import (
	"github.com/kidomine/awsexplorer/model"
	"github.com/kidomine/awsexplorer/view"
)

type Controller struct {
	regionList              []*model.Region
	views                   []view.View
	regionListView          *view.RegionListView
	serviceListView         *view.ServiceListView
	serviceInstanceListView *view.ServiceInstanceListView
	currViewIndex           int
}

func (c *Controller) Initialize() {
	regionIds := model.GetRegionIds()

	//TODO: temporary
	var tmp []string
	tmp = make([]string, 3)
	tmp[0] = "aaa"
	tmp[1] = "bbb"
	tmp[2] = "ccc"

	for _, regionId := range regionIds {
		c.regionList = append(c.regionList, model.NewRegion(regionId))
	}

	c.views = make([]view.View, 3)
	c.setRegionListView(view.NewRegionListView(regionIds))
	currRegion := c.SelectRegionById(c.regionListView.GetSelectedData())
	c.setServiceListView(view.NewServiceListView(currRegion.GetServiceIds()))
	c.setServiceInstanceListView(view.NewServiceInstanceListView(tmp))
	c.regionListView.Render()
	c.serviceListView.Render()
	c.serviceInstanceListView.Render()

	c.currViewIndex = 0
}

func (c *Controller) setRegionListView(regionListView *view.RegionListView) {
	c.regionListView = regionListView
	c.views[0] = c.regionListView
}

func (c *Controller) setServiceListView(serviceListView *view.ServiceListView) {
	c.serviceListView = serviceListView
	c.views[1] = c.serviceListView
}

func (c *Controller) setServiceInstanceListView(serviceInstanceListView *view.ServiceInstanceListView) {
	c.serviceInstanceListView = serviceInstanceListView
	c.views[2] = c.serviceInstanceListView
}

func (c *Controller) SelectRegionById(regionId string) *model.Region {
	for _, region := range c.regionList {
		if region.Id == regionId {
			return region
		}
	}

	return nil
}

func (c *Controller) HandleEvent(event string) {
	switch event {
	case "h", "<Left>":
		// change to previous view
		c.left()
	case "l", "<Right>":
		// change to next view
		c.right()
	default:
		// let the current view handle the event
		currData := c.getSelectedData()
		c.views[c.currViewIndex].HandleEvent(event)

		if currData != c.getSelectedData() {
			switch c.currViewIndex {
			case 0:
				// region list has been updated
				// no need to Render(), this will be done after the next left() or right() sets
				// the service list to currentView
				newRegion := c.SelectRegionById(c.regionListView.GetSelectedData())
				c.setServiceListView(view.NewServiceListView(newRegion.GetServiceIds()))
			case 1:
				//service list has been updated
			case 2:
				// service instance list has been updated

			}
		}
	}
}

func (c *Controller) Render() {
	c.views[c.currViewIndex].Render()
}

// Helper functions
func (c *Controller) left() {
	if c.currViewIndex > 0 {
		c.currViewIndex -= 1
	}
}

func (c *Controller) right() {
	if c.currViewIndex < len(c.views)-1 {
		c.currViewIndex += 1
	}
}

func (c *Controller) getSelectedData() string {
	return c.views[c.currViewIndex].GetSelectedData()
}
