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

	// dummy variable
	var serviceInstanceList []string

	for _, regionId := range regionIds {
		c.regionList = append(c.regionList, model.NewRegion(regionId))
	}

	c.views = make([]view.View, 3)
	c.setRegionListView(view.NewRegionListView(regionIds))
	currRegion := c.SelectRegionById(c.regionListView.GetSelectedData())
	c.setServiceListView(view.NewServiceListView(currRegion.GetServiceIds()))
	c.setServiceInstanceListView(view.NewServiceInstanceListView(serviceInstanceList))
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
	case "r", "<Enter>":
		c.views[c.currViewIndex].HandleEvent(event)

		switch c.currViewIndex {
		case 0:
			// region list has been updated
			// no need to Render(), this will be done after the next left() or right() sets
			// the service list to currentView
			newRegion := c.SelectRegionById(c.regionListView.GetSelectedData())
			c.setServiceListView(view.NewServiceListView(newRegion.GetServiceIds()))
			c.serviceListView.Render()
		case 1:
			//service list has been updated
			// Render() the service instance list immediately
			currRegion := c.SelectRegionById(c.regionListView.GetSelectedData())
			serviceInstanceIds := currRegion.GetServiceInstanceIds(c.serviceListView.GetSelectedData())
			c.setServiceInstanceListView(view.NewServiceInstanceListView(serviceInstanceIds))
			c.serviceInstanceListView.Render()
		case 2:
			// service instance list has been updated
		}
	default:
		// let the current view handle the event
		currData := c.getSelectedData()
		if currData != "" {
			c.views[c.currViewIndex].HandleEvent(event)

			if currData != c.getSelectedData() {
				switch c.currViewIndex {
				case 0:
					// region id has been changed
					// no need to Render(), this will be done after the next left() or right() sets
					// the service list to currentView
					newRegion := c.SelectRegionById(c.regionListView.GetSelectedData())
					c.setServiceListView(view.NewServiceListView(newRegion.GetServiceIds()))
				case 1:
					//service id has been changed

				case 2:
					// service instance id has been changed

				}
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
