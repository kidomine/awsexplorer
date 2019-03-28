package view

type ServiceListView struct {
	ListView
}

// This function should be called only once due to the fact that the
// coordinates of the widget is fixed.
func NewServiceListView(serviceList []string) *ServiceListView {
	listView := newListView("Services", 21, 0, 41, 45, serviceList)
	return &ServiceListView{*listView}
}
