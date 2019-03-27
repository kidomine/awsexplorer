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

type LambdaInstance struct {
	BaseInstance
}

func (l *LambdaInstance) GetId() string {
	return l.id
}

func (l *LambdaInstance) GetData() string {
	return l.data
}

func (l *LambdaInstance) SetId(id string) {
	l.id = id
}

func (l *LambdaInstance) SetData(data interface{}) {
	if _data, ok := data.(string); ok {
		l.data = _data
	}
}

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

type S3Instance struct {
	BaseInstance
}

func (s *S3Instance) GetId() string {
	return s.id
}

func (s *S3Instance) GetData() string {
	return s.data
}

func (s *S3Instance) SetId(id string) {
	s.id = id
}

func (s *S3Instance) SetData(data interface{}) {
	if _data, ok := data.(string); ok {
		s.data = _data
	}
}
