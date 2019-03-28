package model

type DynamoDBInstance struct {
	BaseInstance
}

func (d *DynamoDBInstance) GetId() string {
	return d.id
}

func (d *DynamoDBInstance) GetData() string {
	return d.data
}

func (d *DynamoDBInstance) SetId(id string) {
	d.id = id
}

func (d *DynamoDBInstance) SetData(data interface{}) {
	if _data, ok := data.(string); ok {
		d.data = _data
	}
}
