package dataworks_public

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

// ListBaselineConfigs invokes the dataworks_public.ListBaselineConfigs API synchronously
// api document: https://help.aliyun.com/api/dataworks-public/listbaselineconfigs.html
func (client *Client) ListBaselineConfigs(request *ListBaselineConfigsRequest) (response *ListBaselineConfigsResponse, err error) {
	response = CreateListBaselineConfigsResponse()
	err = client.DoAction(request, response)
	return
}

// ListBaselineConfigsWithChan invokes the dataworks_public.ListBaselineConfigs API asynchronously
// api document: https://help.aliyun.com/api/dataworks-public/listbaselineconfigs.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListBaselineConfigsWithChan(request *ListBaselineConfigsRequest) (<-chan *ListBaselineConfigsResponse, <-chan error) {
	responseChan := make(chan *ListBaselineConfigsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListBaselineConfigs(request)
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

// ListBaselineConfigsWithCallback invokes the dataworks_public.ListBaselineConfigs API asynchronously
// api document: https://help.aliyun.com/api/dataworks-public/listbaselineconfigs.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListBaselineConfigsWithCallback(request *ListBaselineConfigsRequest, callback func(response *ListBaselineConfigsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListBaselineConfigsResponse
		var err error
		defer close(result)
		response, err = client.ListBaselineConfigs(request)
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

// ListBaselineConfigsRequest is the request struct for api ListBaselineConfigs
type ListBaselineConfigsRequest struct {
	*requests.RpcRequest
	Owner         string           `position:"Body" name:"Owner"`
	SearchText    string           `position:"Body" name:"SearchText"`
	Useflag       requests.Boolean `position:"Body" name:"Useflag"`
	Priority      string           `position:"Body" name:"Priority"`
	PageNumber    requests.Integer `position:"Body" name:"PageNumber"`
	PageSize      requests.Integer `position:"Body" name:"PageSize"`
	ProjectId     requests.Integer `position:"Body" name:"ProjectId"`
	BaselineTypes string           `position:"Body" name:"BaselineTypes"`
}

// ListBaselineConfigsResponse is the response struct for api ListBaselineConfigs
type ListBaselineConfigsResponse struct {
	*responses.BaseResponse
	Success        bool                      `json:"Success" xml:"Success"`
	ErrorCode      string                    `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage   string                    `json:"ErrorMessage" xml:"ErrorMessage"`
	HttpStatusCode int                       `json:"HttpStatusCode" xml:"HttpStatusCode"`
	RequestId      string                    `json:"RequestId" xml:"RequestId"`
	Data           DataInListBaselineConfigs `json:"Data" xml:"Data"`
}

// CreateListBaselineConfigsRequest creates a request to invoke ListBaselineConfigs API
func CreateListBaselineConfigsRequest() (request *ListBaselineConfigsRequest) {
	request = &ListBaselineConfigsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dataworks-public", "2020-05-18", "ListBaselineConfigs", "dide", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListBaselineConfigsResponse creates a response to parse from ListBaselineConfigs response
func CreateListBaselineConfigsResponse() (response *ListBaselineConfigsResponse) {
	response = &ListBaselineConfigsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}