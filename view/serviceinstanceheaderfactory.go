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

package view

type ServiceInstanceHeaderFactory func() []string

func getDummyInstanceHeader() []string {
	return []string{"header 1", "header 2"}
}

func getEC2InstanceHeader() []string {
	return []string{"header 1", "header 2"}
}

func getDynamoDBInstanceHeader() []string {
	return []string{"header 1", "header 2"}
}

func getLambdaInstanceHeader() []string {
	return []string{"Function Name", "Runtime", "Size", "Date"}
}

func getS3InstanceHeader() []string {
	return []string{"header 1", "header 2"}
}

func getServiceInstanceHeader(serviceId string) []string {
	serviceHeaderFactoryMap := map[string]ServiceInstanceHeaderFactory{
		"ec2":      getEC2InstanceHeader,
		"dynamodb": getDynamoDBInstanceHeader,
		"lambda":   getLambdaInstanceHeader,
		"s3":       getS3InstanceHeader,
	}

	factory := serviceHeaderFactoryMap[serviceId]
	if factory == nil {
		factory = getDummyInstanceHeader
	}

	return factory()
}
