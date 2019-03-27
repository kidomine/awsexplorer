package model

import (
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"reflect"
)

type Region struct {
	Id       string
	services []*Service
}

func NewRegionList() []*Region {
	var regions []*Region

	allRegions := getAllRegions()
	regionIds := getRegionIDs(allRegions)

	for _, regionId := range regionIds {
		region := allRegions[regionId]
		regions = append(regions, newRegion(regionId, &region))
	}

	return regions
}

func GetRegionIds(regionList []*Region) []string {
	var regionIds []string
	for _, region := range regionList {
		regionIds = append(regionIds, region.Id)
	}

	return regionIds
}

func newRegion(regionId string, awsRegion *endpoints.Region) *Region {
	awsServices := awsRegion.Services()
	var services []*Service

	for _, serviceId := range getServiceIDs(awsRegion) {
		awsService := awsServices[serviceId]
		services = append(services, newService(serviceId, &awsService))
	}

	return &Region{regionId, services}
}

//func (r *Region) SetRegionServices(services []*Service) {
//	r.services = services
//}

//func (r *Region) SetServiceInstances(serviceId string, serviceInstances []*ServiceInstance) {
//	service := r.getServiceById(serviceId)
//	if service != nil {
//		service.SetServiceInstances(serviceInstances)
//	}
//}

func (r *Region) UpdateServiceInstances(serviceId string) {
	//TODO: connect to AWS and obtain list of service instances
	var serviceInstances []ServiceInstance

	service := r.getServiceById(serviceId)
	if service != nil {
		for _, serviceInstance := range serviceInstances {
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

// Helper functions
func (r *Region) getServiceById(serviceId string) *Service {
	for _, service := range r.services {
		if service.Id == serviceId {
			return service
		}
	}

	return nil
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

func getRegionIDs(allRegions map[string]endpoints.Region) []string {
	keys := reflect.ValueOf(allRegions).MapKeys()
	regionValues := make([]string, len(keys))

	for i := 0; i < len(keys); i++ {
		regionValues[i] = keys[i].String()
	}
	return regionValues
}

func getServiceIDs(region *endpoints.Region) []string {
	regionalServices := region.Services()
	keys := reflect.ValueOf(regionalServices).MapKeys()
	serviceIds := make([]string, len(keys))

	for i := 0; i < len(keys); i++ {
		serviceIds[i] = keys[i].String()
	}
	return serviceIds
}
