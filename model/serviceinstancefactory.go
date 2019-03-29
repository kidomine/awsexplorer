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
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
)

type ServiceInstanceListFactory func(awsSession *session.Session) []ServiceInstance

func getDummyInstances(awsSession *session.Session) []ServiceInstance {
	return nil
}

func getLambdaInstances(awsSession *session.Session) []ServiceInstance {
	return getLambdaList(lambda.New(awsSession))
}

func getServiceInstanceList(serviceId string) ServiceInstanceListFactory {
	var serviceFactoryMap map[string]ServiceInstanceListFactory
	serviceFactoryMap = make(map[string]ServiceInstanceListFactory, 2)

	serviceFactoryMap["lambda"] = getLambdaInstances

	factory := serviceFactoryMap[serviceId]
	if factory == nil {
		factory = getDummyInstances
	}

	return factory
}
