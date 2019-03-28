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
		return service.GetServiceInstanceIds()
	}

	return serviceInstanceIds
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
