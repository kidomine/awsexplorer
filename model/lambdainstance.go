package model

type LambdaInstance struct {
	BaseInstance
}

func newLambdaInstance() *LambdaInstance {
	return nil
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
