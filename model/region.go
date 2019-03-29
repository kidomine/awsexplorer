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
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
)

type Region struct {
	Id       string
	services []*Service
	session  *session.Session
}

func NewRegion(regionId string) *Region {
	allRegions := getAllRegions()
	region := allRegions[regionId]

	return newRegion(regionId, &region)
}

func GetRegionIds() []string {
	var regionIds []string
	regionList := getAllRegions()

	for _, region := range regionList {
		regionIds = append(regionIds, region.ID())
	}

	return regionIds
}

func (r *Region) UpdateServiceInstances(serviceId string) {
	//TODO: connect to AWS and obtain list of service instances
	var serviceInstances []ServiceInstance

	service := r.getServiceById(serviceId)
	if service != nil {
		for _, serviceInstance := range serviceInstances {
			//TODO: update serviceInstance here
			service.UpdateServiceInstance(serviceInstance)
		}
	}
}

func (r *Region) GetServiceIds() []string {
	var serviceIds []string
	for _, service := range r.services {
		serviceIds = append(serviceIds, service.Id)
	}

	return serviceIds
}

func (r *Region) GetServiceInstanceIds(serviceId string) []string {
	var serviceInstanceIds []string
	service := r.getServiceById(serviceId)

	if service != nil {
		return service.getServiceInstanceIds(r.session)
	}

	return serviceInstanceIds
}

func (r *Region) GetServiceInstanceData(serviceId string) []string {
	var serviceInstanceData []string
	service := r.getServiceById(serviceId)

	if service != nil {
		return service.getServiceInstanceData(r.session)
	}

	return serviceInstanceData
}

func (r *Region) GetServiceById(serviceId string) Service {
	return *r.getServiceById(serviceId)
}

func (r *Region) GetServiceInstanceById(serviceId, instanceId string) ServiceInstance {
	return *r.getServiceById(serviceId).getServiceInstanceById(instanceId)
}

func (r *Region) getServiceById(serviceId string) *Service {
	for _, service := range r.services {
		if service.Id == serviceId {
			return service
		}
	}

	return nil
}

// Helper functions
func newRegion(regionId string, awsRegion *endpoints.Region) *Region {

	sessionInstance, _ := session.NewSession(&aws.Config{Region: aws.String(regionId)})

	awsServices := awsRegion.Services()
	var services []*Service

	for _, serviceId := range getServiceIDs(awsRegion) {
		awsService := awsServices[serviceId]
		services = append(services, newService(serviceId, &awsService))
	}

	return &Region{regionId, services, sessionInstance}
}

func appendRegions(allRegions map[string]endpoints.Region, subRegions map[string]endpoints.Region) map[string]endpoints.Region {
	for key, value := range subRegions {
		allRegions[key] = value
	}

	return allRegions
}

func getAllRegions() map[string]endpoints.Region {
	var allRegions map[string]endpoints.Region

	awsRegions := endpoints.AwsPartition().Regions()
	chnRegions := endpoints.AwsCnPartition().Regions()
	govRegions := endpoints.AwsUsGovPartition().Regions()

	allRegions = make(map[string]endpoints.Region, len(awsRegions)+len(chnRegions)+len(govRegions))

	allRegions = appendRegions(allRegions, awsRegions)
	allRegions = appendRegions(allRegions, chnRegions)
	allRegions = appendRegions(allRegions, govRegions)

	return allRegions
}
