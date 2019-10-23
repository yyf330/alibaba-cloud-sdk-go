package vod

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

// GetCategories invokes the vod.GetCategories API synchronously
// api document: https://help.aliyun.com/api/vod/getcategories.html
func (client *Client) GetCategories(request *GetCategoriesRequest) (response *GetCategoriesResponse, err error) {
	response = CreateGetCategoriesResponse()
	err = client.DoAction(request, response)
	return
}

// GetCategoriesWithChan invokes the vod.GetCategories API asynchronously
// api document: https://help.aliyun.com/api/vod/getcategories.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetCategoriesWithChan(request *GetCategoriesRequest) (<-chan *GetCategoriesResponse, <-chan error) {
	responseChan := make(chan *GetCategoriesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetCategories(request)
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

// GetCategoriesWithCallback invokes the vod.GetCategories API asynchronously
// api document: https://help.aliyun.com/api/vod/getcategories.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetCategoriesWithCallback(request *GetCategoriesRequest, callback func(response *GetCategoriesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetCategoriesResponse
		var err error
		defer close(result)
		response, err = client.GetCategories(request)
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

// GetCategoriesRequest is the request struct for api GetCategories
type GetCategoriesRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      string           `position:"Query" name:"ResourceOwnerId"`
	Type                 string           `position:"Query" name:"Type"`
	CateId               requests.Integer `position:"Query" name:"CateId"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              string           `position:"Query" name:"OwnerId"`
	PageNo               requests.Integer `position:"Query" name:"PageNo"`
	SortBy               string           `position:"Query" name:"SortBy"`
}

// GetCategoriesResponse is the response struct for api GetCategories
type GetCategoriesResponse struct {
	*responses.BaseResponse
	RequestId     string        `json:"RequestId" xml:"RequestId"`
	SubTotal      int64         `json:"SubTotal" xml:"SubTotal"`
	Category      Category      `json:"Category" xml:"Category"`
	SubCategories SubCategories `json:"SubCategories" xml:"SubCategories"`
}

// CreateGetCategoriesRequest creates a request to invoke GetCategories API
func CreateGetCategoriesRequest() (request *GetCategoriesRequest) {
	request = &GetCategoriesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vod", "2017-03-21", "GetCategories", "vod", "openAPI")
	return
}

// CreateGetCategoriesResponse creates a response to parse from GetCategories response
func CreateGetCategoriesResponse() (response *GetCategoriesResponse) {
	response = &GetCategoriesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
