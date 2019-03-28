package model

type EC2Instance struct {
	BaseInstance
}

func (e *EC2Instance) GetId() string {
	return e.id
}

func (e *EC2Instance) GetData() string {
	return e.data
}

func (e *EC2Instance) SetId(id string) {
	e.id = id
}

func (e *EC2Instance) SetData(data interface{}) {
	if _data, ok := data.(string); ok {
		e.data = _data
	}
}
