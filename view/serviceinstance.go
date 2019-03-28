package view

type ServiceInstanceListView struct {
	ListView
}

// This function should be called only once due to the fact that the
// coordinates of the widget is fixed.
func NewServiceInstanceListView(serviceInstanceList []string) *ServiceInstanceListView {
	listView := newListView("Service Instances", 42, 0, 62, 45, serviceInstanceList)
	return &ServiceInstanceListView{*listView}
}
