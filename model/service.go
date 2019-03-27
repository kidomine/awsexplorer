package model

import "github.com/aws/aws-sdk-go/aws/endpoints"

type Service struct {
	Id        string
	instances []ServiceInstance
}

func newService(serviceId string, awsService *endpoints.Service) *Service {
	//serviceInstances := nil
	return &Service{serviceId, nil}
}

func (s *Service) SetServiceInstances(serviceInstances []ServiceInstance) {
	s.instances = serviceInstances
}

func (s *Service) UpdateServiceInstance(serviceInstance ServiceInstance) {
	for _, s := range s.instances {
		if s.GetId() == serviceInstance.GetId() {
			//TODO: temporary
			s.SetData(serviceInstance.GetData())
			break
		}
	}
}

func (s *Service) GetServiceInstanceIds() []string {
	var instanceIds []string
	for _, s := range s.instances {
		instanceIds = append(instanceIds, s.GetId())
	}

	return instanceIds
}

func (s *Service) getServiceInstanceById(instanceId string) *ServiceInstance {
	for _, instance := range s.instances {
		if instance.GetId() == instanceId {
			return &instance
		}
	}

	return nil
}
