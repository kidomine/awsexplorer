package model

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
