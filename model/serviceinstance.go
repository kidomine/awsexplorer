package model

type ServiceInstance interface {
	GetId() string
	GetData() string
	SetId(id string)
	SetData(data interface{})
}

type BaseInstance struct {
	id   string
	data string
}
