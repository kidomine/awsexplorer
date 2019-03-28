package view

type RegionListView struct {
	ListView
}

// This function should be called only once due to the fact that the
// coordinates of the widget is fixed.
func NewRegionListView(regionList []string) *RegionListView {
	listView := newListView("Regions", 0, 0, 20, 45, regionList)
	return &RegionListView{*listView}
}
