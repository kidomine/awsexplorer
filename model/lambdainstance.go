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
	"fmt"
	"github.com/aws/aws-sdk-go/service/lambda"
	"os"
)

type LambdaInstance struct {
	BaseInstance

	// The SHA256 hash of the function's deployment package.
	CodeSha256 *string `type:"string"`

	// The size of the function's deployment package in bytes.
	CodeSize *int64 `type:"long"`

	// The function's dead letter queue.
	//DeadLetterConfig *DeadLetterConfig `type:"structure"`

	// The function's description.
	//Description *string `type:"string"`

	// The function's environment variables.
	//Environment *EnvironmentResponse `type:"structure"`

	// The function's Amazon Resource Name (ARN).
	FunctionArn *string `type:"string"`

	// The name of the function.
	FunctionName *string `min:"1" type:"string"`

	// The function Lambda calls to begin executing your function.
	Handler *string `type:"string"`

	// The KMS key used to encrypt the function's environment variables. Only returned
	// if you've configured a customer managed CMK.
	//KMSKeyArn *string `type:"string"`

	// The date and time that the function was last updated, in ISO-8601 format
	// (https://www.w3.org/TR/NOTE-datetime) (YYYY-MM-DDThh:mm:ss.sTZD).
	LastModified *string `type:"string"`

	// The function's  layers (http://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html)
	//Layers []*Layer `type:"list"`

	// For Lambda@Edge functions, the ARN of the master function.
	MasterArn *string `type:"string"`

	// The memory allocated to the function
	MemorySize *int64 `min:"128" type:"integer"`

	// Represents the latest updated revision of the function or alias.
	//RevisionId *string `type:"string"`

	// The function's execution role.
	Role *string `type:"string"`

	// The runtime environment for the Lambda function.
	Runtime *string `type:"string" enum:"Runtime"`

	// The amount of time that Lambda allows a function to run before terminating
	// it.
	Timeout *int64 `min:"1" type:"integer"`

	// The function's AWS X-Ray tracing configuration.
	//TracingConfig *TracingConfigResponse `type:"structure"`

	// The version of the Lambda function.
	Version *string `min:"1" type:"string"`

	// The function's networking configuration.
	//vpcConfig *VpcConfigResponse `type:"structure"`
}

func newLambdaInstance(lambdaFunction *lambda.FunctionConfiguration) *LambdaInstance {
	lambdaInstance := &LambdaInstance{}
	lambdaInstance.CodeSha256 = lambdaFunction.CodeSha256
	lambdaInstance.CodeSize = lambdaFunction.CodeSize
	lambdaInstance.FunctionArn = lambdaFunction.FunctionArn
	lambdaInstance.FunctionName = lambdaFunction.FunctionName
	lambdaInstance.Handler = lambdaFunction.Handler
	lambdaInstance.LastModified = lambdaFunction.LastModified
	lambdaInstance.MemorySize = lambdaFunction.MemorySize
	lambdaInstance.Role = lambdaFunction.Role
	lambdaInstance.Runtime = lambdaFunction.Runtime
	lambdaInstance.Timeout = lambdaFunction.Timeout
	lambdaInstance.Version = lambdaFunction.Version

	lambdaInstance.id = *lambdaInstance.FunctionName
	lambdaInstance.data = fmt.Sprintf("%s %s %d %s",
		*lambdaInstance.FunctionName, *lambdaInstance.Runtime, lambdaInstance.MemorySize, *lambdaInstance.LastModified)

	return lambdaInstance
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

// Helper functions
func getLambdaList(client *lambda.Lambda) []ServiceInstance {
	//details := []*DynamoDbTableInfo{}

	result, err := client.ListFunctions(&lambda.ListFunctionsInput{})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return getLambdaFunctionNames(result.Functions)
}

func getLambdaFunctionNames(lambdaList []*lambda.FunctionConfiguration) []ServiceInstance {
	var lamdaInstances []ServiceInstance

	for _, lambdaFunction := range lambdaList {
		lamdaInstances = append(lamdaInstances, newLambdaInstance(lambdaFunction))
	}

	return lamdaInstances
}
