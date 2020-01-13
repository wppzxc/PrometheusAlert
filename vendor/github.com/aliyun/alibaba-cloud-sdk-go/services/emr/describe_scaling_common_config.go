package emr

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// DescribeScalingCommonConfig invokes the emr.DescribeScalingCommonConfig API synchronously
// api document: https://help.aliyun.com/api/emr/describescalingcommonconfig.html
func (client *Client) DescribeScalingCommonConfig(request *DescribeScalingCommonConfigRequest) (response *DescribeScalingCommonConfigResponse, err error) {
	response = CreateDescribeScalingCommonConfigResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeScalingCommonConfigWithChan invokes the emr.DescribeScalingCommonConfig API asynchronously
// api document: https://help.aliyun.com/api/emr/describescalingcommonconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeScalingCommonConfigWithChan(request *DescribeScalingCommonConfigRequest) (<-chan *DescribeScalingCommonConfigResponse, <-chan error) {
	responseChan := make(chan *DescribeScalingCommonConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeScalingCommonConfig(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// DescribeScalingCommonConfigWithCallback invokes the emr.DescribeScalingCommonConfig API asynchronously
// api document: https://help.aliyun.com/api/emr/describescalingcommonconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeScalingCommonConfigWithCallback(request *DescribeScalingCommonConfigRequest, callback func(response *DescribeScalingCommonConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeScalingCommonConfigResponse
		var err error
		defer close(result)
		response, err = client.DescribeScalingCommonConfig(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// DescribeScalingCommonConfigRequest is the request struct for api DescribeScalingCommonConfig
type DescribeScalingCommonConfigRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
}

// DescribeScalingCommonConfigResponse is the response struct for api DescribeScalingCommonConfig
type DescribeScalingCommonConfigResponse struct {
	*responses.BaseResponse
	RequestId                    string `json:"RequestId" xml:"RequestId"`
	AutoScalingRuleMinDelayLimit int    `json:"AutoScalingRuleMinDelayLimit" xml:"AutoScalingRuleMinDelayLimit"`
}

// CreateDescribeScalingCommonConfigRequest creates a request to invoke DescribeScalingCommonConfig API
func CreateDescribeScalingCommonConfigRequest() (request *DescribeScalingCommonConfigRequest) {
	request = &DescribeScalingCommonConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "DescribeScalingCommonConfig", "emr", "openAPI")
	return
}

// CreateDescribeScalingCommonConfigResponse creates a response to parse from DescribeScalingCommonConfig response
func CreateDescribeScalingCommonConfigResponse() (response *DescribeScalingCommonConfigResponse) {
	response = &DescribeScalingCommonConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}