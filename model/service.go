/*
 * MIT License
 *
 * Copyright (c) 2019 Ian Diaz.
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package model

import (
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"reflect"
)

type Service struct {
	Id        string
	instances []ServiceInstance
}

func newService(serviceId string, awsService *endpoints.Service) *Service {
	return &Service{serviceId, nil}
}

func getServiceIDs(region *endpoints.Region) []string {
	regionalServices := region.Services()
	keys := reflect.ValueOf(regionalServices).MapKeys()

	serviceIds := make([]string, 0)
	for _, serviceId := range keys {
		if checkIfServiceIsSupported(serviceId.String()) == true {
			serviceIds = append(serviceIds, serviceId.String())
		}
	}

	return serviceIds
}

func checkIfServiceIsSupported(serviceId string) bool {
	var supportedServices []string
	supportedServices = make([]string, 0)

	supportedServices = append(supportedServices, "dynamodb")
	supportedServices = append(supportedServices, "ec2")
	supportedServices = append(supportedServices, "lambda")
	supportedServices = append(supportedServices, "s3")

	for _, supportedService := range supportedServices {
		if supportedService == serviceId {
			return true
		}
	}
	return false
}

func (s *Service) getServiceInstances(awsSession *session.Session) {
	factory := getServiceInstanceList(s.Id)
	if factory == nil {
		factory = getDummyInstances
	}

	s.instances = factory(awsSession)
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

func (s *Service) getServiceInstanceIds(awsSession *session.Session) []string {
	if s.instances == nil {
		s.getServiceInstances(awsSession)
	}

	var instanceIds []string
	for _, s := range s.instances {
		instanceIds = append(instanceIds, s.GetId())
	}

	return instanceIds
}

func (s *Service) getServiceInstanceData(awsSession *session.Session) []string {
	if s.instances == nil {
		s.getServiceInstances(awsSession)
	}

	var instanceData []string
	for _, s := range s.instances {
		instanceData = append(instanceData, s.GetData())
	}

	return instanceData
}

func (s *Service) getServiceInstanceById(instanceId string) *ServiceInstance {
	for _, instance := range s.instances {
		if instance.GetId() == instanceId {
			return &instance
		}
	}

	return nil
}
